/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"sort"
	"time"

	"kubepack.dev/kubepack/pkg/lib"
	"kubepack.dev/lib-app/pkg/editor"
	"kubepack.dev/lib-app/pkg/handler"
	actionx "kubepack.dev/lib-helm/pkg/action"
	"kubepack.dev/lib-helm/pkg/getter"
	"kubepack.dev/lib-helm/pkg/repo"
	"kubepack.dev/lib-helm/pkg/values"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	"github.com/spf13/pflag"
	"github.com/unrolled/render"
	"go.wandrs.dev/binding"
	httpw "go.wandrs.dev/http"
	"gomodules.xyz/logs"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/klog/v2"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
	meta_util "kmodules.xyz/client-go/meta"
	"kmodules.xyz/client-go/tools/converter"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/yaml"
	chartsapi "x-helm.dev/apimachinery/apis/charts/v1alpha1"
	releasesapi "x-helm.dev/apimachinery/apis/releases/v1alpha1"
)

var HelmRegistry = repo.NewDiskCacheRegistry()

func main() {
	kubeConfigFlags := genericclioptions.NewConfigFlags(true)
	kubeConfigFlags.AddFlags(pflag.CommandLine)
	matchVersionKubeConfigFlags := cmdutil.NewMatchVersionFlags(kubeConfigFlags)
	matchVersionKubeConfigFlags.AddFlags(pflag.CommandLine)

	logs.Init(nil, true)
	defer logs.FlushLogs()

	f := cmdutil.NewFactory(matchVersionKubeConfigFlags)

	m := chi.NewRouter()
	m.Use(middleware.RequestID)
	m.Use(middleware.RealIP)
	m.Use(middleware.Logger) // middlewares.NewLogger()
	m.Use(middleware.Recoverer)
	m.Use(binding.Injector(render.New()))

	// PUBLIC
	m.Route("/bundleview", func(m chi.Router) {
		m.With(binding.JSON(releasesapi.ChartSourceFlatRef{})).Get("/", binding.HandlerFunc(GetBundleViewForChart))

		// Generate Order for a BundleView
		m.With(binding.JSON(releasesapi.BundleView{})).Post("/orders", binding.HandlerFunc(CreateOrderForBundle))
	})

	// PUBLIC
	m.Route("/packageview", func(m chi.Router) {
		m.With(binding.JSON(releasesapi.ChartSourceFlatRef{})).Get("/", binding.HandlerFunc(GetPackageViewForChart))

		// PUBLIC
		m.With(binding.JSON(releasesapi.ChartSourceFlatRef{})).Get("/files", binding.HandlerFunc(ListPackageFiles))

		// PUBLIC
		m.With(binding.JSON(releasesapi.ChartSourceFlatRef{})).Get("/files/*", binding.HandlerFunc(GetPackageFile))

		// PUBLIC
		m.With(binding.JSON(chartsapi.ChartPresetFlatRef{})).Get("/values", binding.HandlerFunc(GetValuesFile))

		// Generate Order for a Editor PackageView / Chart
		m.With(binding.JSON(releasesapi.ChartOrder{})).Post("/orders", binding.HandlerFunc(CreateOrderForPackage))
	})

	m.Route("/editor", func(m chi.Router) {
		m.Use(binding.JSON(map[string]interface{}{}))

		// PUBLIC
		// INITIAL Model (Values)
		// GET vs POST (Get makes more sense, but do we send so much data via query string?)
		// With POST, we can send large payloads without any non-standard limits
		// https://stackoverflow.com/a/812962
		m.Put("/model", binding.HandlerFunc(GenerateEditorModelFromOptions))
		m.Put("/manifest", binding.HandlerFunc(PreviewEditorManifest))
		m.Put("/resources", binding.HandlerFunc(PreviewEditorResources))
	})

	// PUBLIC
	/*
		m.Get("/products", binding.HandlerFunc(func(ctx httpw.ResponseWriter) {
			// /products

			phase := ctx.R().Params("phase")
			var out v1alpha1.ProductList
			fsys := artifacts.Products()
			_ = fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
				if err != nil {
					return err
				}
				if d.IsDir() {
					return nil
				}
				data, err := fs.ReadFile(fsys, path)
				if err != nil {
					ctx.Error(http.StatusInternalServerError, err.Error())
					return
				}
				var p v1alpha1.Product
				err = json.Unmarshal(data, &p)
				if err != nil {
					ctx.Error(http.StatusInternalServerError, err.Error())
					return
				}
				if phase == "" || p.Spec.Phase == v1alpha1.Phase(phase) {
					out.Items = append(out.Items, p)
				}
			})
			ctx.JSON(http.StatusOK, out)
		}))

		// PUBLIC
		m.Get("/products/{owner}/{key}", binding.HandlerFunc(func(ctx httpw.ResponseWriter) {
			// /products/appscode/kubedb
			// TODO: get product by (owner, key)

			data, err := fs.ReadFile(artifacts.Products(), ctx.R().Params("key")+".json")
			if err != nil {
				ctx.Error(http.StatusInternalServerError, err.Error())
				return
			}
			_, _ = ctx.Write(data)
		}))

		// PUBLIC
		m.Get("/products/{owner}/{key}/plans", binding.HandlerFunc(func(ctx httpw.ResponseWriter) {
			// /products/appscode/kubedb
			// TODO: get product by (owner, key)

			dir := "artifacts/products/kubedb-plans"
			files, err := ioutil.ReadDir(dir)
			if err != nil {
				ctx.Error(http.StatusInternalServerError, err.Error())
				return
			}

			phase := ctx.R().Params("phase")
			var out v1alpha1.PlanList
			for _, file := range files {
				data, err := os.ReadFile(filepath.Join(dir, file.Name()))
				if err != nil {
					ctx.Error(http.StatusInternalServerError, err.Error())
					return
				}
				var plaan v1alpha1.Plan
				err = json.Unmarshal(data, &plaan)
				if err != nil {
					ctx.Error(http.StatusInternalServerError, err.Error())
					return
				}
				if phase == "" || plaan.Spec.Phase == v1alpha1.Phase(phase) {
					out.Items = append(out.Items, plaan)
				}
			}
			ctx.JSON(http.StatusOK, out)
		}))

		// PUBLIC
		m.Get("/products/{owner}/{key}/plans/{plan}", binding.HandlerFunc(func(ctx httpw.ResponseWriter) {
			// /products/appscode/kubedb
			// TODO: get product by (owner, key)

			dir := "artifacts/products/" + ctx.R().Params("key") + "-plans"

			data, err := os.ReadFile(filepath.Join(dir, ctx.R().Params("plan")+".json"))
			if err != nil {
				ctx.Error(http.StatusInternalServerError, err.Error())
				return
			}
			var plaan v1alpha1.Plan
			err = json.Unmarshal(data, &plaan)
			if err != nil {
				ctx.Error(http.StatusInternalServerError, err.Error())
				return
			}
			ctx.JSON(http.StatusOK, plaan)
		}))

		// PUBLIC
		m.Get("/products/{owner}/{key}/compare", binding.HandlerFunc(func(ctx httpw.ResponseWriter) {
			// /products/appscode/kubedb
			// TODO: get product by (owner, key)

			phase := ctx.R().Params("phase")

			// product
			data, err := fs.ReadFile(artifacts.Products(), ctx.R().Params("key")+".json")
			if err != nil {
				ctx.Error(http.StatusInternalServerError, err.Error())
				return
			}
			var p v1alpha1.Product
			err = json.Unmarshal(data, &p)
			if err != nil {
				ctx.Error(http.StatusInternalServerError, err.Error())
				return
			}

			var url string
			version := p.Spec.LatestVersion
			var plaans []v1alpha1.Plan

			dir := "artifacts/products/" + ctx.R().Params("key") + "-plans"
			files, err := ioutil.ReadDir(dir)
			if err != nil {
				ctx.Error(http.StatusInternalServerError, err.Error())
				return
			}
			for idx, file := range files {
				data, err := os.ReadFile(filepath.Join(dir, file.Name()))
				if err != nil {
					ctx.Error(http.StatusInternalServerError, err.Error())
					return
				}
				var plaan v1alpha1.Plan
				err = json.Unmarshal(data, &plaan)
				if err != nil {
					ctx.Error(http.StatusInternalServerError, err.Error())
					return
				}
				if idx == 0 {
					url = plaan.Spec.Bundle.URL
				}

				if phase == "" || plaan.Spec.Phase == v1alpha1.Phase(phase) {
					plaans = append(plaans, plaan)
				}
			}

			sort.Slice(plaans, func(i, j int) bool {
				if plaans[i].Spec.Weight == plaans[j].Spec.Weight {
					return plaans[i].Spec.NickName < plaans[j].Spec.NickName
				}
				return plaans[i].Spec.Weight < plaans[j].Spec.Weight
			})

			var names []string
			for _, plaan := range plaans {
				names = append(names, plaan.Spec.Bundle.Name)
			}

			table, err := lib.ComparePlans(HelmRegistry, url, names, version)
			if err != nil {
				ctx.Error(http.StatusInternalServerError, err.Error())
				return
			}
			table.Plans = plaans
			ctx.JSON(http.StatusOK, table)
		}))

		// PUBLIC
		m.Get("/products/{owner}/{key}/plans/{plan}/bundleview", binding.HandlerFunc(func(ctx httpw.ResponseWriter) {
			// /products/appscode/kubedb
			// TODO: get product by (owner, key)

			// product
			data, err := fs.ReadFile(artifacts.Products(), ctx.R().Params("key")+".json")
			if err != nil {
				ctx.Error(http.StatusInternalServerError, err.Error())
				return
			}
			var p v1alpha1.Product
			err = json.Unmarshal(data, &p)
			if err != nil {
				ctx.Error(http.StatusInternalServerError, err.Error())
				return
			}

			// plan
			dir := "artifacts/products/" + ctx.R().Params("key") + "-plans"
			data, err = os.ReadFile(filepath.Join(dir, ctx.R().Params("plan")+".json"))
			if err != nil {
				ctx.Error(http.StatusInternalServerError, err.Error())
				return
			}
			var plaan v1alpha1.Plan
			err = json.Unmarshal(data, &plaan)
			if err != nil {
				ctx.Error(http.StatusInternalServerError, err.Error())
				return
			}

			bv, err := lib.CreateBundleViewForBundle(HelmRegistry, &v1alpha1.ChartRepoRef{
				URL:     plaan.Spec.Bundle.URL,
				Name:    plaan.Spec.Bundle.Name,
				Version: p.Spec.LatestVersion,
			})
			if err != nil {
				ctx.Error(http.StatusInternalServerError, err.Error())
				return
			}
			ctx.JSON(http.StatusOK, bv)
		}))

		// PUBLIC
		m.Get("/product_id/{id}", binding.HandlerFunc(func(ctx httpw.ResponseWriter) {
			// TODO: get product by id

			data, err := fs.ReadFile(artifacts.Products(), "kubedb.json")
			if err != nil {
				ctx.Error(http.StatusInternalServerError, err.Error())
				return
			}
			_, _ = ctx.Write(data)
		}))
	*/

	// PRIVATE
	// Should we store Order UID in a table per User?
	m.Route("/deploy/orders", func(m chi.Router) {
		// Generate Order for a single chart and optional values patch
		m.With(binding.JSON(releasesapi.Order{})).Post("/", binding.HandlerFunc(CreateOrder))
		m.Get("/{id}/render/manifest", binding.HandlerFunc(PreviewOrderManifest))
		m.Get("/{id}/render/resources", binding.HandlerFunc(PreviewOrderResources))
		m.Get("/{id}/helm3", binding.HandlerFunc(func(ctx httpw.ResponseWriter) {
			bs, err := lib.NewTestBlobStore()
			if err != nil {
				ctx.Error(http.StatusInternalServerError, err.Error())
				return
			}

			data, err := bs.ReadFile(ctx.R().Request().Context(), path.Join(ctx.R().Params("id"), "order.yaml"))
			if err != nil {
				ctx.Error(http.StatusInternalServerError, err.Error())
				return
			}

			var order releasesapi.Order
			err = yaml.Unmarshal(data, &order)
			if err != nil {
				ctx.Error(http.StatusInternalServerError, err.Error())
				return
			}

			scripts, err := lib.GenerateHelm3Script(bs, HelmRegistry, order)
			if err != nil {
				ctx.Error(http.StatusInternalServerError, err.Error())
				return
			}

			ctx.JSON(http.StatusOK, scripts)
		}))
		m.Get("/{id}/yaml", binding.HandlerFunc(func(ctx httpw.ResponseWriter) {
			bs, err := lib.NewTestBlobStore()
			if err != nil {
				ctx.Error(http.StatusInternalServerError, err.Error())
				return
			}

			data, err := bs.ReadFile(ctx.R().Request().Context(), path.Join(ctx.R().Params("id"), "order.yaml"))
			if err != nil {
				ctx.Error(http.StatusInternalServerError, err.Error())
				return
			}

			var order releasesapi.Order
			err = yaml.Unmarshal(data, &order)
			if err != nil {
				ctx.Error(http.StatusInternalServerError, err.Error())
				return
			}

			scripts, err := lib.GenerateYAMLScript(bs, HelmRegistry, order)
			if err != nil {
				ctx.Error(http.StatusInternalServerError, err.Error())
				return
			}

			ctx.JSON(http.StatusOK, scripts)
		}))
	})

	// PRIVATE
	m.Route("/clusters/{cluster}", func(m chi.Router) {
		m.Route("/editor", func(m chi.Router) {
			// PUT create / update / apply / install
			m.With(binding.JSON(map[string]interface{}{})).Put("/", binding.HandlerFunc(ApplyResource(f)))

			m.Delete("/namespaces/{namespace}/releases/{releaseName}", binding.HandlerFunc(DeleteResource(f)))

			// PUT Model from Existing Installations
			m.With(binding.JSON(releasesapi.ModelMetadata{})).Put("/model", binding.HandlerFunc(LoadEditorModel))

			// redundant apis
			// can be replaced by getting the model, then using the /editor apis
			m.With(binding.JSON(releasesapi.Model{})).Put("/manifest", binding.HandlerFunc(LoadEditorManifest))

			// redundant apis
			// can be replaced by getting the model, then using the /editor apis
			m.With(binding.JSON(releasesapi.Model{})).Put("/resources", binding.HandlerFunc(LoadEditorResources))
		})

		m.Post("/deploy/{id}", binding.HandlerFunc(func(ctx httpw.ResponseWriter) {
		}))
		m.Delete("/deploy/{id}", binding.HandlerFunc(func(ctx httpw.ResponseWriter) {
		}))
	})

	m.Get("/chartrepositories", binding.HandlerFunc(func(ctx httpw.ResponseWriter) {
		repos, err := repo.DefaultNamer.ListHelmHubRepositories()
		if err != nil {
			ctx.Error(http.StatusInternalServerError, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, repos)
	}))
	m.Get("/chartrepositories/charts", binding.HandlerFunc(func(ctx httpw.ResponseWriter) {
		url := ctx.R().Query("url")
		if url == "" {
			ctx.Error(http.StatusBadRequest, "missing url")
			return
		}

		cfg, _, err := HelmRegistry.Get(url)
		if err != nil {
			ctx.Error(http.StatusInternalServerError, err.Error())
			return
		}
		cr, err := repo.NewChartRepository(cfg, getter.All())
		if err != nil {
			ctx.Error(http.StatusInternalServerError, err.Error())
			return
		}
		err = cr.Load()
		if err != nil {
			ctx.Error(http.StatusInternalServerError, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, cr.ListCharts())
	}))
	m.Get("/chartrepositories/charts/{name}/versions", binding.HandlerFunc(func(ctx httpw.ResponseWriter) {
		url := ctx.R().Query("url")
		name := ctx.R().Params("name")

		if url == "" {
			ctx.Error(http.StatusBadRequest, "missing url")
			return
		}
		if name == "" {
			ctx.Error(http.StatusBadRequest, "missing chart name")
			return
		}

		cfg, _, err := HelmRegistry.Get(url)
		if err != nil {
			ctx.Error(http.StatusInternalServerError, err.Error())
			return
		}
		cr, err := repo.NewChartRepository(cfg, getter.All())
		if err != nil {
			ctx.Error(http.StatusInternalServerError, err.Error())
			return
		}
		err = cr.Load()
		if err != nil {
			ctx.Error(http.StatusInternalServerError, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, cr.ListVersions(name))
	}))

	m.Get("/", binding.HandlerFunc(func() string {
		return "Hello world!"
	}))
	klog.Infoln()
	klog.Infoln("Listening on :4000")
	if err := http.ListenAndServe(":4000", m); err != nil {
		klog.Fatalln(err)
	}
}

func LoadEditorResources(ctx httpw.ResponseWriter, model releasesapi.Model) {
	cfg, err := config.GetConfig()
	if err != nil {
		ctx.Error(http.StatusInternalServerError, err.Error())
		return
	}
	kc, err := actionx.NewUncachedClientForConfig(cfg)
	if err != nil {
		ctx.Error(http.StatusInternalServerError, err.Error())
		return
	}

	tpl, err := editor.LoadResourceEditorModel(kc, HelmRegistry, releasesapi.ModelMetadata{
		Metadata: model.Metadata,
	})
	if err != nil {
		ctx.Error(http.StatusInternalServerError, "LoadEditorModel", err.Error())
		return
	}

	var out releasesapi.ResourceOutput
	format := meta_util.NewDataFormat(ctx.R().QueryTrim("format"), meta_util.YAMLFormat)

	for _, r := range tpl.Resources {
		data, err := meta_util.Marshal(r, format)
		if err != nil {
			ctx.Error(http.StatusInternalServerError, "MarshalResource", err.Error())
			return
		}
		out.Resources = append(out.Resources, releasesapi.ResourceFile{
			Filename: r.Filename + "." + string(format),
			Key:      r.Key,
			Data:     string(data),
		})
	}
	ctx.JSON(http.StatusOK, out)
}

func LoadEditorManifest(ctx httpw.ResponseWriter, model releasesapi.Model) {
	cfg, err := config.GetConfig()
	if err != nil {
		ctx.Error(http.StatusInternalServerError, err.Error())
		return
	}
	kc, err := actionx.NewUncachedClientForConfig(cfg)
	if err != nil {
		ctx.Error(http.StatusInternalServerError, err.Error())
		return
	}

	tpl, err := editor.LoadResourceEditorModel(kc, HelmRegistry, releasesapi.ModelMetadata{
		Metadata: model.Metadata,
	})
	if err != nil {
		ctx.Error(http.StatusInternalServerError, "LoadEditorModel", err.Error())
		return
	}
	_, _ = ctx.Write(tpl.Manifest)
}

func LoadEditorModel(ctx httpw.ResponseWriter, model releasesapi.ModelMetadata) {
	cfg, err := config.GetConfig()
	if err != nil {
		ctx.Error(http.StatusInternalServerError, err.Error())
		return
	}
	kc, err := actionx.NewUncachedClientForConfig(cfg)
	if err != nil {
		ctx.Error(http.StatusInternalServerError, err.Error())
		return
	}

	reg := repo.NewRegistry(kc, repo.DefaultDiskCache())
	tpl, err := editor.LoadResourceEditorModel(kc, reg, model)
	if err != nil {
		ctx.Error(http.StatusInternalServerError, "LoadEditorModel", err.Error())
		return
	}

	format := meta_util.NewDataFormat(ctx.R().QueryTrim("format"), meta_util.JsonFormat)
	out, err := meta_util.Marshal(tpl.Values, format)
	if err != nil {
		ctx.Error(http.StatusInternalServerError, "MarshalModel", err.Error())
		return
	}
	_, _ = ctx.Write(out)
}

func DeleteResource(f cmdutil.Factory) func(ctx httpw.ResponseWriter) {
	return func(ctx httpw.ResponseWriter) {
		release := releasesapi.ObjectMeta{
			Name:      ctx.R().Params("releaseName"),
			Namespace: ctx.R().Params("namespace"),
		}
		rls, err := handler.DeleteResource(f, release)
		if err != nil {
			ctx.Error(http.StatusInternalServerError, "DeleteResource", err.Error())
			return
		}
		_, _ = ctx.Write([]byte(rls.Info))
	}
}

func ApplyResource(f cmdutil.Factory) func(ctx httpw.ResponseWriter, model map[string]interface{}) {
	return func(ctx httpw.ResponseWriter, model map[string]interface{}) {
		kc, err := actionx.NewUncachedClient(f)
		if err != nil {
			ctx.Error(http.StatusInternalServerError, "ApplyResourceEditor", err.Error())
			return
		}
		reg := repo.NewRegistry(kc, repo.DefaultDiskCache())
		rls, err := handler.ApplyResourceEditor(f, reg, model, !ctx.R().QueryBool("installCRDs"))
		if err != nil {
			fmt.Println(err.Error())
			ctx.Error(http.StatusInternalServerError, "ApplyResourceEditor", err.Error())
			return
		}
		ctx.JSON(http.StatusOK, rls.Info)
	}
}

func PreviewOrderResources(ctx httpw.ResponseWriter) {
	bs, err := lib.NewTestBlobStore()
	if err != nil {
		ctx.Error(http.StatusInternalServerError, err.Error())
		return
	}

	data, err := bs.ReadFile(ctx.R().Request().Context(), path.Join(ctx.R().Params("id"), "order.yaml"))
	if err != nil {
		ctx.Error(http.StatusInternalServerError, "BlobStoreReadFile", err.Error())
		return
	}

	var order releasesapi.Order
	err = yaml.Unmarshal(data, &order)
	if err != nil {
		ctx.Error(http.StatusInternalServerError, "UnmarshalOrder", err.Error())
		return
	}

	_, tpls, err := editor.RenderOrderTemplate(bs, HelmRegistry, order)
	if err != nil {
		ctx.Error(http.StatusInternalServerError, "RenderOrderTemplate", err.Error())
		return
	}
	if ctx.R().QueryBool("skipCRDs") {
		for i := range tpls {
			tpls[i].CRDs = nil
		}
	}

	format := meta_util.NewDataFormat(ctx.R().QueryTrim("format"), meta_util.YAMLFormat)
	out, err := editor.ConvertChartTemplates(tpls, format)
	if err != nil {
		ctx.Error(http.StatusInternalServerError, "ConvertChartTemplates", err.Error())
		return
	}
	ctx.JSON(http.StatusOK, out)
}

func PreviewOrderManifest(ctx httpw.ResponseWriter) {
	bs, err := lib.NewTestBlobStore()
	if err != nil {
		ctx.Error(http.StatusInternalServerError, err.Error())
		return
	}

	data, err := bs.ReadFile(ctx.R().Request().Context(), path.Join(ctx.R().Params("id"), "order.yaml"))
	if err != nil {
		ctx.Error(http.StatusInternalServerError, "BlobStoreReadFile", err.Error())
		return
	}

	var order releasesapi.Order
	err = yaml.Unmarshal(data, &order)
	if err != nil {
		ctx.Error(http.StatusInternalServerError, "UnmarshalOrder", err.Error())
		return
	}

	manifest, _, err := editor.RenderOrderTemplate(bs, HelmRegistry, order)
	if err != nil {
		ctx.Error(http.StatusInternalServerError, "RenderOrderTemplate", err.Error())
		return
	}
	_, _ = ctx.Write([]byte(manifest))
}

func CreateOrder(ctx httpw.ResponseWriter, order releasesapi.Order) {
	if len(order.Spec.Packages) == 0 {
		ctx.Error(http.StatusBadRequest, "missing package selection for order")
		return
	}

	order.TypeMeta = metav1.TypeMeta{
		APIVersion: releasesapi.GroupVersion.String(),
		Kind:       releasesapi.ResourceKindOrder,
	}
	order.ObjectMeta = metav1.ObjectMeta{
		UID:               types.UID(uuid.New().String()),
		CreationTimestamp: metav1.NewTime(time.Now()),
	}
	if order.Name == "" {
		order.Name = order.Spec.Packages[0].Chart.ReleaseName
	}

	data, err := yaml.Marshal(order)
	if err != nil {
		ctx.Error(http.StatusInternalServerError, "MarshalOrder", err.Error())
		return
	}

	bs, err := lib.NewTestBlobStore()
	if err != nil {
		ctx.Error(http.StatusInternalServerError, err.Error())
		return
	}

	err = bs.WriteFile(ctx.R().Request().Context(), path.Join(string(order.UID), "order.yaml"), data)
	if err != nil {
		ctx.Error(http.StatusInternalServerError, "BlobStoreWriteFile", err.Error())
		return
	}
	ctx.JSON(http.StatusOK, order)
}

func PreviewEditorResources(ctx httpw.ResponseWriter, opts map[string]interface{}) {
	cfg, err := config.GetConfig()
	if err != nil {
		ctx.Error(http.StatusInternalServerError, err.Error())
		return
	}
	kc, err := actionx.NewUncachedClientForConfig(cfg)
	if err != nil {
		ctx.Error(http.StatusInternalServerError, err.Error())
		return
	}

	_, tpls, err := editor.RenderResourceEditorChart(kc, HelmRegistry, opts)
	if err != nil {
		ctx.Error(http.StatusInternalServerError, "RenderChartTemplate", err.Error())
		return
	}
	if ctx.R().QueryBool("skipCRDs") {
		tpls.CRDs = nil
	}

	var out releasesapi.ResourceOutput
	format := meta_util.NewDataFormat(ctx.R().QueryTrim("format"), meta_util.YAMLFormat)

	for _, crd := range tpls.CRDs {
		data, err := meta_util.Marshal(crd.Data, format)
		if err != nil {
			ctx.Error(http.StatusInternalServerError, "MarshalCRD", err.Error())
			return
		}
		out.CRDs = append(out.CRDs, releasesapi.ResourceFile{
			Filename: crd.Filename + "." + string(format),
			Data:     string(data),
		})
	}
	for _, r := range tpls.Resources {
		data, err := meta_util.Marshal(r.Data, format)
		if err != nil {
			ctx.Error(http.StatusInternalServerError, "MarshalResource", err.Error())
			return
		}
		out.Resources = append(out.Resources, releasesapi.ResourceFile{
			Filename: r.Filename + "." + string(format),
			Key:      r.Key,
			Data:     string(data),
		})
	}
	ctx.JSON(http.StatusOK, &out)
}

func PreviewEditorManifest(ctx httpw.ResponseWriter, opts map[string]interface{}) {
	cfg, err := config.GetConfig()
	if err != nil {
		ctx.Error(http.StatusInternalServerError, err.Error())
		return
	}
	kc, err := actionx.NewUncachedClientForConfig(cfg)
	if err != nil {
		ctx.Error(http.StatusInternalServerError, err.Error())
		return
	}

	manifest, _, err := editor.RenderResourceEditorChart(kc, HelmRegistry, opts)
	if err != nil {
		ctx.Error(http.StatusInternalServerError, "RenderChartTemplate", err.Error())
		return
	}
	_, _ = ctx.Write([]byte(manifest))
}

func GenerateEditorModelFromOptions(ctx httpw.ResponseWriter, opts map[string]interface{}) {
	cfg, err := config.GetConfig()
	if err != nil {
		ctx.Error(http.StatusInternalServerError, err.Error())
		return
	}
	kc, err := actionx.NewUncachedClientForConfig(cfg)
	if err != nil {
		ctx.Error(http.StatusInternalServerError, err.Error())
		return
	}

	model, err := editor.GenerateResourceEditorModel(kc, HelmRegistry, opts)
	if err != nil {
		ctx.Error(http.StatusInternalServerError, "GetChart", err.Error())
		return
	}

	format := meta_util.NewDataFormat(ctx.R().QueryTrim("format"), meta_util.JsonFormat)
	if format == meta_util.YAMLFormat {
		out, err := yaml.Marshal(model)
		if err != nil {
			ctx.Error(http.StatusInternalServerError, err.Error())
			return
		}
		_, _ = ctx.Write(out)
		return
	}
	ctx.JSON(http.StatusOK, model)
}

func GetPackageFile(ctx httpw.ResponseWriter, params releasesapi.ChartSourceFlatRef) {
	chrt, err := HelmRegistry.GetChart(params.ToAPIObject())
	if err != nil {
		ctx.Error(http.StatusInternalServerError, "GetChart", err.Error())
		return
	}

	filename := ctx.R().Params("*")
	format := meta_util.NewDataFormat(ctx.R().QueryTrim("format"), meta_util.KeepFormat)
	for _, f := range chrt.Raw {
		if f.Name == filename {
			out, ct, err := converter.Convert(f.Name, f.Data, format)
			if err != nil {
				ctx.Error(http.StatusInternalServerError, "ConvertFormat", err.Error())
				return
			}

			ctx.Header().Set("Content-Type", ct)
			_, _ = ctx.Write(out)
			return
		}
	}
	ctx.WriteHeader(http.StatusNotFound)
}

func GetValuesFile(ctx httpw.ResponseWriter, params chartsapi.ChartPresetFlatRef) {
	cfg, err := config.GetConfig()
	if err != nil {
		ctx.Error(http.StatusInternalServerError, err.Error())
		return
	}
	kc, err := actionx.NewUncachedClientForConfig(cfg)
	if err != nil {
		ctx.Error(http.StatusInternalServerError, err.Error())
		return
	}
	HelmRegistry = repo.NewRegistry(kc, repo.DefaultDiskCache())

	chrt, err := HelmRegistry.GetChart(params.ToAPIObject())
	if err != nil {
		ctx.Error(http.StatusInternalServerError, "GetChart", err.Error())
		return
	}

	vals, err := values.MergePresetValues(kc, chrt.Chart, params)
	if err != nil {
		ctx.Error(http.StatusInternalServerError, "MergePresetValues", err.Error())
		return
	}

	vals, err = editor.UpdateFeatureValues(kc, chrt.Chart, vals)
	if err != nil {
		ctx.Error(http.StatusInternalServerError, "UpdateFeatureValues", err.Error())
		return
	}

	var data []byte
	var ct string
	format := meta_util.NewDataFormat(ctx.R().QueryTrim("format"), meta_util.KeepFormat)
	if format == meta_util.JsonFormat {
		ct = "application/json"
		data, err = json.Marshal(vals)
	} else if format == meta_util.YAMLFormat {
		ct = "text/yaml"
		data, err = yaml.Marshal(vals)
	}
	if err != nil {
		ctx.Error(http.StatusInternalServerError, "ConvertFormat", err.Error())
		return
	}

	ctx.Header().Set("Content-Type", ct)
	_, _ = ctx.Write(data)
}

func ListPackageFiles(ctx httpw.ResponseWriter, params releasesapi.ChartSourceFlatRef) {
	// TODO: verify params

	chrt, err := HelmRegistry.GetChart(params.ToAPIObject())
	if err != nil {
		ctx.Error(http.StatusInternalServerError, "GetChart", err.Error())
		return
	}

	files := make([]string, 0, len(chrt.Raw))
	for _, f := range chrt.Raw {
		files = append(files, f.Name)
	}
	sort.Strings(files)

	ctx.JSON(http.StatusOK, files)
}

func CreateOrderForPackage(ctx httpw.ResponseWriter, params releasesapi.ChartOrder) {
	order, err := editor.CreateChartOrder(HelmRegistry, params)
	if err != nil {
		ctx.Error(http.StatusInternalServerError, "CreateChartOrder", err.Error())
		return
	}

	data, err := yaml.Marshal(order)
	if err != nil {
		ctx.Error(http.StatusInternalServerError, "MarshalOrder", err.Error())
		return
	}

	bs, err := lib.NewTestBlobStore()
	if err != nil {
		ctx.Error(http.StatusInternalServerError, err.Error())
		return
	}

	err = bs.WriteFile(ctx.R().Request().Context(), path.Join(string(order.UID), "order.yaml"), data)
	if err != nil {
		ctx.Error(http.StatusInternalServerError, "BlobStoreWriteFile", err.Error())
		return
	}
	ctx.JSON(http.StatusOK, order)
}

func GetPackageViewForChart(ctx httpw.ResponseWriter, params releasesapi.ChartSourceFlatRef) {
	// TODO: verify params

	chrt, err := HelmRegistry.GetChart(params.ToAPIObject())
	if err != nil {
		ctx.Error(http.StatusInternalServerError, "GetChart", err.Error())
		return
	}

	pv, err := lib.CreatePackageView(params.ToAPIObject().SourceRef, chrt.Chart)
	if err != nil {
		ctx.Error(http.StatusInternalServerError, "CreatePackageView", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, pv)
}

func CreateOrderForBundle(ctx httpw.ResponseWriter, params releasesapi.BundleView) {
	order, err := lib.CreateOrder(HelmRegistry, params)
	if err != nil {
		ctx.Error(http.StatusInternalServerError, "CreateDeployOrder", err.Error())
		return
	}

	data, err := yaml.Marshal(order)
	if err != nil {
		ctx.Error(http.StatusInternalServerError, "MarshalDeployOrder", err.Error())
		return
	}

	bs, err := lib.NewTestBlobStore()
	if err != nil {
		ctx.Error(http.StatusInternalServerError, err.Error())
		return
	}

	err = bs.WriteFile(ctx.R().Request().Context(), path.Join(string(order.UID), "order.yaml"), data)
	if err != nil {
		ctx.Error(http.StatusInternalServerError, "BlobStoreWriteFile", err.Error())
		return
	}
	ctx.JSON(http.StatusOK, order)
}

func GetBundleViewForChart(ctx httpw.ResponseWriter, params releasesapi.ChartSourceFlatRef) {
	// TODO: verify params

	bv, err := lib.CreateBundleViewForChart(HelmRegistry, params.ToAPIObject())
	if err != nil {
		ctx.Error(http.StatusInternalServerError, "CreateBundleViewForChart", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, bv)
}

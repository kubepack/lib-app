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

package editor

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"kubepack.dev/kubepack/apis/kubepack/v1alpha1"
	"kubepack.dev/kubepack/pkg/lib"
	appapi "kubepack.dev/lib-app/api/v1alpha1"
	"kubepack.dev/lib-helm/pkg/repo"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gomodules.xyz/jsonpatch/v3"
	"helm.sh/helm/v3/pkg/chart"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	"kmodules.xyz/client-go/discovery"
	meta_util "kmodules.xyz/client-go/meta"
	"kmodules.xyz/client-go/tools/parser"
	"kmodules.xyz/resource-metadata/hub/resourceeditors"
	"sigs.k8s.io/application/api/app/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/yaml"
)

var helmRepositories = map[string]string{
	"appscode":        "https://charts.appscode.com/stable/",
	"bytebuilders-ui": "https://bundles.byte.builders/ui/",
}

func RenderOrderTemplate(bs *lib.BlobStore, reg repo.IRegistry, order v1alpha1.Order) (string, []appapi.ChartTemplate, error) {
	var buf bytes.Buffer
	var tpls []appapi.ChartTemplate

	for _, pkg := range order.Spec.Packages {
		if pkg.Chart == nil {
			continue
		}

		f1 := &TemplateRenderer{
			Registry:    reg,
			ChartRef:    pkg.Chart.ChartRef,
			Version:     pkg.Chart.Version,
			ReleaseName: pkg.Chart.ReleaseName,
			Namespace:   pkg.Chart.Namespace,
			KubeVersion: "v1.22.0",
			ValuesFile:  pkg.Chart.ValuesFile,
			ValuesPatch: pkg.Chart.ValuesPatch,
			BucketURL:   bs.Bucket,
			UID:         string(order.UID),
			PublicURL:   bs.Host,
		}
		err := f1.Do()
		if err != nil {
			return "", nil, err
		}

		tpl := appapi.ChartTemplate{
			ChartRef:    pkg.Chart.ChartRef,
			Version:     pkg.Chart.Version,
			ReleaseName: pkg.Chart.ReleaseName,
			Namespace:   pkg.Chart.Namespace,
		}
		crds, manifestFile := f1.Result()
		for _, crd := range crds {
			resources, err := ListResources(pkg.Chart.ReleaseName, crd.Data)
			if err != nil {
				return "", nil, err
			}
			if len(resources) != 1 {
				return "", nil, fmt.Errorf("%d crds found in %s", len(resources), crd.Filename)
			}
			tpl.CRDs = append(tpl.CRDs, appapi.BucketObject{
				URL: crd.URL,
				Key: crd.Key,
				ResourceObject: appapi.ResourceObject{
					Filename: crd.Filename,
					Data:     resources[0].Data,
				},
			})
		}
		if manifestFile != nil {
			tpl.Manifest = &appapi.BucketFileRef{
				URL: manifestFile.URL,
				Key: manifestFile.Key,
			}
			tpl.Resources, err = ListResources(pkg.Chart.ReleaseName, manifestFile.Data)
			if err != nil {
				return "", nil, err
			}
			_, err = fmt.Fprintf(&buf, "---\n# Source: %s - %s@%s\n", f1.ChartRef.URL, f1.ChartRef.Name, f1.Version)
			if err != nil {
				return "", nil, err
			}

			_, err := buf.Write(manifestFile.Data)
			if err != nil {
				return "", nil, err
			}
			_, err = buf.WriteRune('\n')
			if err != nil {
				return "", nil, err
			}
		}
		tpls = append(tpls, tpl)
	}

	return buf.String(), tpls, nil
}

func LoadResourceEditorModel(kc client.Client, reg repo.IRegistry, opts appapi.ModelMetadata) (*appapi.EditorTemplate, error) {
	ed, ok := resourceeditors.LoadByResourceID(kc, &opts.Resource)
	if !ok {
		return nil, fmt.Errorf("failed to load resource editor for %+v", opts.Resource)
	}
	chartRef := v1alpha1.ChartRepoRef{
		URL:     helmRepositories[ed.Spec.UI.Editor.SourceRef.Name],
		Name:    ed.Spec.UI.Editor.Name,
		Version: ed.Spec.UI.Editor.Version,
	}
	return loadEditorModel(kc, reg, chartRef, opts)
}

func LoadEditorModel(kc client.Client, reg repo.IRegistry, chartRef v1alpha1.ChartRepoRef, opts appapi.ModelMetadata) (*appapi.EditorTemplate, error) {
	return loadEditorModel(kc, reg, chartRef, opts)
}

func loadEditorModel(kc client.Client, reg repo.IRegistry, chartRef v1alpha1.ChartRepoRef, opts appapi.ModelMetadata) (*appapi.EditorTemplate, error) {
	chrt, err := reg.GetChart(chartRef.URL, chartRef.Name, chartRef.Version)
	if err != nil {
		return nil, err
	}

	var app v1beta1.Application
	err = kc.Get(context.TODO(), client.ObjectKey{Namespace: opts.Release.Namespace, Name: opts.Release.Name}, &app)
	if err != nil {
		return nil, err
	}

	return EditorChartValueManifest(kc, &app, opts.Metadata.Release, chrt.Chart)
}

func EditorChartValueManifest(kc client.Client, app *v1beta1.Application, mt appapi.ObjectMeta, chrt *chart.Chart) (*appapi.EditorTemplate, error) {
	selector, err := metav1.LabelSelectorAsSelector(app.Spec.Selector)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	resourceMap := map[string]interface{}{}

	// detect apiVersion from defaultValues in chart
	gkToVersion := map[metav1.GroupKind]string{}
	rsKeys := sets.NewString()
	for rsKey, x := range chrt.Values["resources"].(map[string]interface{}) {
		rsKeys.Insert(rsKey)
		var tm metav1.TypeMeta

		err = meta_util.DecodeObject(x.(map[string]interface{}), &tm)
		if err != nil {
			return nil, fmt.Errorf("failed to parse TypeMeta for rsKey %s in chart name=%s version=%s values", rsKey, chrt.Name(), chrt.Metadata.Version)
		}
		gv, err := schema.ParseGroupVersion(tm.APIVersion)
		if err != nil {
			return nil, err
		}
		gkToVersion[metav1.GroupKind{
			Group: gv.Group,
			Kind:  tm.Kind,
		}] = gv.Version
	}

	mapper := discovery.NewResourceMapper(kc.RESTMapper())
	_, usesForm := chrt.Values["form"]

	var resources []*unstructured.Unstructured
	for _, gk := range app.Spec.ComponentGroupKinds {
		version, ok := gkToVersion[gk]
		if !ok {
			if !usesForm {
				return nil, fmt.Errorf("failed to detect version for GK %#v in chart name=%s version=%s values", gk, chrt.Name(), chrt.Metadata.Version)
			} else {
				mapping, err := kc.RESTMapper().RESTMapping(schema.GroupKind{
					Group: gk.Group,
					Kind:  gk.Kind,
				})
				if err != nil {
					return nil, err
				}
				version = mapping.GroupVersionKind.Version
			}
		}

		gvk := schema.GroupVersionKind{
			Group:   gk.Group,
			Version: version,
			Kind:    gk.Kind,
		}
		namespaced, err := mapper.IsGVKNamespaced(gvk)
		if meta.IsNoMatchError(err) {
			continue // CRD type not installed, so skip it
		}
		if err != nil {
			return nil, errors.Wrapf(err, "failed to detect if gvk %v is namespaced", gvk)
		}
		var list unstructured.UnstructuredList
		list.SetGroupVersionKind(gvk)
		opts := []client.ListOption{client.MatchingLabelsSelector{Selector: selector}}
		if namespaced {
			opts = append(opts, client.InNamespace(mt.Namespace))
		}
		err = kc.List(context.TODO(), &list, opts...)
		if err != nil {
			return nil, err
		}
		for idx := range list.Items {
			obj := list.Items[idx]
			// remove status
			delete(obj.Object, "status")

			resources = append(resources, &obj)

			buf.WriteString("\n---\n")
			data, err := yaml.Marshal(&obj)
			if err != nil {
				return nil, err
			}
			buf.Write(data)

			rsKey, err := ResourceKey(obj.GetAPIVersion(), obj.GetKind(), mt.Name, obj.GetName())
			if err != nil {
				return nil, err
			}
			if rsKeys.Has(rsKey) { // ski form objects
				if _, ok := resourceMap[rsKey]; ok {
					return nil, fmt.Errorf("duplicate resource key %s for application %s/%s", rsKey, app.Namespace, app.Name)
				}
				resourceMap[rsKey] = &obj
			}
		}
	}

	s1map := map[string]int{}
	s2map := map[string]int{}
	s3map := map[string]int{}
	for _, obj := range resources {
		s1, s2, s3 := ResourceFilename(obj.GetAPIVersion(), obj.GetKind(), mt.Name, obj.GetName())
		if v, ok := s1map[s1]; !ok {
			s1map[s1] = 1
		} else {
			s1map[s1] = v + 1
		}
		if v, ok := s2map[s2]; !ok {
			s2map[s2] = 1
		} else {
			s2map[s2] = v + 1
		}
		if v, ok := s3map[s3]; !ok {
			s3map[s3] = 1
		} else {
			s3map[s3] = v + 1
		}
	}

	rsfiles := make([]appapi.ResourceObject, 0, len(resources))
	for _, obj := range resources {
		s1, s2, s3 := ResourceFilename(obj.GetAPIVersion(), obj.GetKind(), mt.Name, obj.GetName())
		name := s1
		if s1map[s1] > 1 {
			if s2map[s2] > 1 {
				name = s3
			} else {
				name = s2
			}
		}
		rsfiles = append(rsfiles, appapi.ResourceObject{
			Filename: name,
			Data:     obj,
		})
	}

	tpl := appapi.EditorTemplate{
		Manifest: buf.Bytes(),
		Values: &unstructured.Unstructured{
			Object: map[string]interface{}{
				"metadata": map[string]interface{}{
					"resource": chrt.Values["metadata"].(map[string]interface{})["resource"],
					"release":  mt,
				},
				"resources": resourceMap,
			},
		},
		Resources: rsfiles,
	}
	if usesForm {
		// https://github.com/kubepack/lib-helm/commit/4279003342a502f328f3c9a6334f4ab5bfdf900d
		if f, ok := app.Annotations["form.release.x-helm.dev/"+mt.Name]; ok {
			var form map[string]interface{}
			if err = json.Unmarshal([]byte(f), &form); err == nil {
				tpl.Values.Object["form"] = form
			}
		}
	}

	return &tpl, nil
}

func GenerateResourceEditorModel(kc client.Client, reg repo.IRegistry, opts map[string]interface{}) (*unstructured.Unstructured, error) {
	var spec appapi.ModelMetadata
	err := meta_util.DecodeObject(opts, &spec)
	if err != nil {
		return nil, err
	}

	ed, ok := resourceeditors.LoadByResourceID(kc, &spec.Resource)
	if !ok {
		return nil, fmt.Errorf("failed to load resource editor for %+v", spec.Resource)
	}

	chartRef := v1alpha1.ChartRepoRef{
		URL:     helmRepositories[ed.Spec.UI.Options.SourceRef.Name],
		Name:    ed.Spec.UI.Options.Name,
		Version: ed.Spec.UI.Options.Version,
	}

	return generateEditorModel(kc, reg, chartRef, spec, opts)
}

func GenerateEditorModel(kc client.Client, reg repo.IRegistry, chartRef v1alpha1.ChartRepoRef, opts map[string]interface{}) (*unstructured.Unstructured, error) {
	var spec appapi.ModelMetadata
	err := meta_util.DecodeObject(opts, &spec)
	if err != nil {
		return nil, err
	}
	return generateEditorModel(kc, reg, chartRef, spec, opts)
}

func generateEditorModel(
	kc client.Client,
	reg repo.IRegistry,
	chartRef v1alpha1.ChartRepoRef,
	spec appapi.ModelMetadata,
	opts map[string]interface{},
) (*unstructured.Unstructured, error) {
	_, usesForm := opts["form"]
	rsKeys := sets.NewString()

	if usesForm {
		ref := struct {
			url     string
			name    string
			version string
		}{
			url:     chartRef.URL,
			name:    chartRef.Name,
			version: chartRef.Version,
		}
		chrt, err := reg.GetChart(ref.url, ref.name, ref.version)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to load resource editor chart %+v", ref)
		}
		if data, ok := chrt.Chart.Metadata.Annotations["meta.x-helm.dev/editor"]; ok && data != "" {
			// store reference rsKeys
			if refResources, ok := chrt.Chart.Values["resources"].(map[string]interface{}); ok {
				for key := range refResources {
					rsKeys.Insert(key)
				}
			}
		} else {
			return nil, fmt.Errorf("editor chart %+v is missing annotation key meta.x-helm.dev/editor", ref)
		}
	}

	f1 := &EditorModelGenerator{
		Registry: reg,
		ChartRef: v1alpha1.ChartRef{
			URL:  chartRef.URL,
			Name: chartRef.Name,
		},
		Version:     chartRef.Version,
		ReleaseName: spec.Metadata.Release.Name,
		Namespace:   spec.Metadata.Release.Namespace,
		KubeVersion: "v1.22.0",
		Values:      opts,
	}
	err := f1.Do(kc)
	if err != nil {
		return nil, err
	}

	resourceValues := map[string]interface{}{}
	_, manifest := f1.Result()
	err = parser.ProcessResources(manifest, func(ri parser.ResourceInfo) error {
		rsKey, err := ResourceKey(ri.Object.GetAPIVersion(), ri.Object.GetKind(), spec.Metadata.Release.Name, ri.Object.GetName())
		if err != nil {
			return err
		}

		// values
		if !usesForm || rsKeys.Has(rsKey) {
			resourceValues[rsKey] = ri.Object
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	model := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"metadata":  opts["metadata"],
			"resources": resourceValues,
		},
	}
	if form, ok := opts["form"]; ok {
		model.Object["form"] = form
	}
	return model, err
}

func RenderResourceEditorChart(kc client.Client, reg repo.IRegistry, opts map[string]interface{}) (string, *appapi.ChartTemplate, error) {
	var spec appapi.ModelMetadata
	err := meta_util.DecodeObject(opts, &spec)
	if err != nil {
		return "", nil, err
	}

	ed, ok := resourceeditors.LoadByResourceID(kc, &spec.Resource)
	if !ok {
		return "", nil, fmt.Errorf("failed to load resource editor for %+v", spec.Resource)
	}

	chartRef := v1alpha1.ChartRepoRef{
		URL:     helmRepositories[ed.Spec.UI.Editor.SourceRef.Name],
		Name:    ed.Spec.UI.Editor.Name,
		Version: ed.Spec.UI.Editor.Version,
	}
	return renderChart(kc, reg, chartRef, spec, opts)
}

func RenderChart(kc client.Client, reg repo.IRegistry, charRef v1alpha1.ChartRepoRef, opts map[string]interface{}) (string, *appapi.ChartTemplate, error) {
	var spec appapi.ModelMetadata
	err := meta_util.DecodeObject(opts, &spec)
	if err != nil {
		return "", nil, err
	}

	return renderChart(kc, reg, charRef, spec, opts)
}

func renderChart(
	kc client.Client,
	reg repo.IRegistry,
	charRef v1alpha1.ChartRepoRef,
	spec appapi.ModelMetadata,
	opts map[string]interface{},
) (string, *appapi.ChartTemplate, error) {
	f1 := &EditorModelGenerator{
		Registry: reg,
		ChartRef: v1alpha1.ChartRef{
			URL:  charRef.URL,
			Name: charRef.Name,
		},
		Version:        charRef.Version,
		ReleaseName:    spec.Release.Name,
		Namespace:      spec.Release.Namespace,
		KubeVersion:    "v1.22.0",
		Values:         opts,
		RefillMetadata: true,
	}
	err := f1.Do(kc)
	if err != nil {
		return "", nil, err
	}

	tpl := appapi.ChartTemplate{
		ChartRef:    f1.ChartRef,
		Version:     f1.Version,
		ReleaseName: f1.ReleaseName,
		Namespace:   f1.Namespace,
	}

	crds, manifest := f1.Result()
	for _, crd := range crds {
		resources, err := parser.ListResources(crd.Data)
		if err != nil {
			return "", nil, err
		}
		if len(resources) != 1 {
			return "", nil, fmt.Errorf("%d crds found in %s", len(resources), crd.Name)
		}
		tpl.CRDs = append(tpl.CRDs, appapi.BucketObject{
			ResourceObject: appapi.ResourceObject{
				Filename: crd.Name,
				Data:     resources[0].Object,
			},
		})
	}
	if manifest != nil {
		tpl.Resources, err = ListResources(spec.Release.Name, manifest)
		if err != nil {
			return "", nil, err
		}
	}
	return string(manifest), &tpl, nil
}

func CreateChartOrder(reg repo.IRegistry, opts appapi.ChartOrder) (*v1alpha1.Order, error) {
	// editor chart
	chrt, err := reg.GetChart(opts.URL, opts.Name, opts.Version)
	if err != nil {
		return nil, err
	}
	originalValues, err := json.Marshal(chrt.Values)
	if err != nil {
		return nil, err
	}

	modifiedValues, err := json.Marshal(opts.Values)
	if err != nil {
		return nil, err
	}
	patch, err := jsonpatch.CreatePatch(originalValues, modifiedValues)
	if err != nil {
		return nil, err
	}
	patchData, err := json.Marshal(patch)
	if err != nil {
		return nil, err
	}

	order := v1alpha1.Order{
		TypeMeta: metav1.TypeMeta{
			APIVersion: v1alpha1.SchemeGroupVersion.String(),
			Kind:       v1alpha1.ResourceKindOrder,
		}, ObjectMeta: metav1.ObjectMeta{
			Name:              opts.ReleaseName,
			Namespace:         opts.Namespace,
			UID:               types.UID(uuid.New().String()),
			CreationTimestamp: metav1.NewTime(time.Now()),
		},
		Spec: v1alpha1.OrderSpec{
			Packages: []v1alpha1.PackageSelection{
				{
					Chart: &v1alpha1.ChartSelection{
						ChartRef: v1alpha1.ChartRef{
							URL:  opts.URL,
							Name: opts.Name,
						},
						Version:     opts.Version,
						ReleaseName: opts.ReleaseName,
						Namespace:   opts.Namespace,
						Bundle:      nil,
						ValuesFile:  "values.yaml",
						ValuesPatch: &runtime.RawExtension{
							Raw: patchData,
						},
						Resources: nil,
						WaitFors:  nil,
					},
				},
			},
			KubeVersion: "",
		},
	}
	return &order, err
}

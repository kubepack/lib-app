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

package simple

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	docapi "kubepack.dev/chart-doc-gen/api"
	"kubepack.dev/lib-app/pkg/utils"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/crane"
	"github.com/spf13/cobra"
	"helm.sh/helm/v3/pkg/chart"
	crdv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/util/homedir"
	kmapi "kmodules.xyz/client-go/api/v1"
	rsapi "kmodules.xyz/resource-metadata/apis/meta/v1alpha1"
	uiapi "kmodules.xyz/resource-metadata/apis/ui/v1alpha1"
	"kmodules.xyz/resource-metadata/hub"
	"kmodules.xyz/resource-metadata/hub/resourceeditors"
	"sigs.k8s.io/yaml"
	releasesapi "x-helm.dev/apimachinery/apis/releases/v1alpha1"
)

const repoName = "bytebuilders-ui"

var (
	descriptorDir  = homedir.HomeDir() + "/go/src/kmodules.xyz/resource-metadata/hub/resourcedescriptors/"
	chartDir       = homedir.HomeDir() + "/go/src/go.bytebuilders.dev/ui-wizards/charts"
	chartVersion   = utils.ChartVersion
	chartUseDigest = utils.ChartUseDigest
	gvr            schema.GroupVersionResource
	all            bool
	skipExisting   bool

	registry = hub.NewRegistryOfKnownResources()
)

func NewCmdSimple() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "simple-chart",
		Short:             `Generate simple chart`,
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			descriptorDir = filepath.Clean(descriptorDir)

			if !all {
				return GenerateSimpleEditorChart(chartDir, descriptorDir, gvr, registry, skipExisting)
			}

			registry.Visit(func(key string, rd *rsapi.ResourceDescriptor) {
				if rd.Spec.Resource.Group == "ui.k8s.appscode.com" {
					if rd.Spec.Resource.Kind == "FeatureSet" {
						chartName := fmt.Sprintf("%s-%s-{.metadata.release.name}-editor", safeGroupName(rd.Spec.Resource.Group), strings.ToLower(rd.Spec.Resource.Kind))
						err := UpdateEditor(rd, chartName, descriptorDir)
						if err != nil {
							panic(err)
						}
						return
					} else if rd.Spec.Resource.Kind == "Feature" {
						return
					}
				}

				err := GenerateSimpleEditorChart(chartDir, descriptorDir, rd.Spec.Resource.GroupVersionResource(), registry, skipExisting)
				if err != nil {
					panic(err)
				}
			})
			return nil
		},
	}

	cmd.Flags().StringVar(&chartDir, "chart-dir", chartDir, "Charts dir")
	cmd.Flags().StringVar(&chartVersion, "chart-version", chartVersion, "Chart version")
	cmd.Flags().BoolVar(&chartUseDigest, "chart-use-digest", chartUseDigest, "Use digest instead of tag")
	cmd.Flags().BoolVar(&all, "all", all, "Generate editor charts for all")
	cmd.Flags().BoolVar(&skipExisting, "skipExisting", skipExisting, "Skip existing chart")
	cmd.Flags().StringVar(&descriptorDir, "descriptor-dir", descriptorDir, "Resource descriptor dir")

	cmd.Flags().StringVar(&gvr.Group, "resource.group", gvr.Group, "Resource api group")
	cmd.Flags().StringVar(&gvr.Version, "resource.version", gvr.Version, "Resource api version")
	cmd.Flags().StringVar(&gvr.Resource, "resource.name", gvr.Resource, "Resource plural")

	return cmd
}

func GenerateSimpleEditorChart(chartDir, descriptorDir string, gvr schema.GroupVersionResource, registry *hub.Registry, skipExisting bool) error {
	rd, err := registry.LoadByGVR(gvr)
	if err != nil {
		return err
	}

	chartName := fmt.Sprintf("%s-%s-editor", safeGroupName(rd.Spec.Resource.Group), strings.ToLower(rd.Spec.Resource.Kind))

	if _, err := os.Stat(filepath.Join(chartDir, chartName)); !os.IsNotExist(err) && skipExisting {
		return fmt.Errorf("%s chart already exists", chartName)
	}

	tplDir := filepath.Join(chartDir, chartName, "templates")
	err = os.MkdirAll(tplDir, 0o755)
	if err != nil {
		return err
	}

	crdDir := filepath.Join(chartDir, chartName, "crds")
	err = os.MkdirAll(crdDir, 0o755)
	if err != nil {
		return err
	}

	err = GenerateChartMetadata(chartDir, chartName, rd)
	if err != nil {
		return err
	}

	{
		if rd.Spec.Validation != nil && rd.Spec.Validation.OpenAPIV3Schema != nil {
			data3, err := yaml.Marshal(rd.Spec.Validation.OpenAPIV3Schema)
			if err != nil {
				return err
			}
			schemaFilename := filepath.Join(chartDir, chartName, "values.openapiv3_schema.yaml")
			err = os.WriteFile(schemaFilename, data3, 0o644)
			if err != nil {
				return err
			}
		}
	}

	if IsCRD(gvr.Group) {
		if rd.Spec.Validation != nil && rd.Spec.Validation.OpenAPIV3Schema != nil {
			rd.Spec.Validation.OpenAPIV3Schema.Properties["metadata"] = crdv1.JSONSchemaProps{
				Type: "object",
			}
		}

		crd := crdv1.CustomResourceDefinition{
			TypeMeta: metav1.TypeMeta{
				APIVersion: crdv1.SchemeGroupVersion.String(),
				Kind:       "CustomResourceDefinition",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: fmt.Sprintf("%s.%s", gvr.Resource, gvr.Group),
			},
			Spec: crdv1.CustomResourceDefinitionSpec{
				Group: gvr.Group,
				Names: crdv1.CustomResourceDefinitionNames{
					Plural:   rd.Spec.Resource.Name,
					Singular: strings.ToLower(rd.Spec.Resource.Kind),
					// ShortNames: nil,
					Kind:     rd.Spec.Resource.Kind,
					ListKind: rd.Spec.Resource.Kind + "List",
					// Categories: nil,
				},
				Scope: crdv1.ResourceScope(string(rd.Spec.Resource.Scope)),
				Versions: []crdv1.CustomResourceDefinitionVersion{
					{
						Name:    rd.Spec.Resource.Version,
						Served:  true,
						Storage: true,
						Schema:  rd.Spec.Validation,
						// Subresources:             nil,
						// AdditionalPrinterColumns: nil,
					},
				},
				PreserveUnknownFields: false,
			},
		}
		if strings.HasSuffix(gvr.Group, ".k8s.io") ||
			strings.HasSuffix(gvr.Group, "kubernetes.io") {
			crd.Annotations = map[string]string{
				"api-approved.kubernetes.io": "https://github.com/kubernetes-sigs/appRelease/pull/2",
			}
		}

		filename := filepath.Join(crdDir, fmt.Sprintf("%s_%s.yaml", gvr.Group, gvr.Resource))
		data, err := yaml.Marshal(crd)
		if err != nil {
			return err
		}
		err = os.WriteFile(filename, data, 0o644)
		if err != nil {
			return err
		}
	}

	{
		data := `# Patterns to ignore when building packages.
# This supports shell glob matching, relative path matching, and
# negation (prefixed with !). Only one pattern per line.
.DS_Store
# Common VCS dirs
.git/
.gitignore
.bzr/
.bzrignore
.hg/
.hgignore
.svn/
# Common backup files
*.swp
*.bak
*.tmp
*.orig
*~
# Various IDEs
.project
.idea/
*.tmproj
.vscode/
`
		schemaFilename := filepath.Join(chartDir, chartName, ".helmignore")
		err = os.WriteFile(schemaFilename, []byte(data), 0o644)
		if err != nil {
			return err
		}
	}

	{
		fqdn := gvr.Resource
		if IsCRD(gvr.Group) {
			fqdn = fmt.Sprintf("%s.%s", gvr.Resource, gvr.Group)
		}
		data := fmt.Sprintf(`Get the %s by running the following command:

  kubectl --namespace {{ .Release.Namespace }} get %s {{ .Release.Name }}
`, rd.Spec.Resource.Kind, fqdn)
		schemaFilename := filepath.Join(chartDir, chartName, "templates", "NOTES.txt")
		err = os.WriteFile(schemaFilename, []byte(data), 0o644)
		if err != nil {
			return err
		}
	}

	{
		gv := schema.GroupVersion{
			Group:   rd.Spec.Resource.Group,
			Version: rd.Spec.Resource.Version,
		}
		v := releasesapi.SimpleValue{
			TypeMeta: metav1.TypeMeta{
				APIVersion: gv.String(),
				Kind:       rd.Spec.Resource.Kind,
			},
			ObjectMeta: releasesapi.ObjectMeta{
				Name: strings.ToLower(rd.Spec.Resource.Kind),
			},
		}
		if rd.Spec.Resource.Scope == kmapi.NamespaceScoped {
			v.ObjectMeta.Namespace = "default"
		}

		data, err := yaml.Marshal(v)
		if err != nil {
			panic(err)
		}

		filename := filepath.Join(chartDir, chartName, "values.yaml")
		err = os.WriteFile(filename, data, 0o644)
		if err != nil {
			return err
		}
	}

	{
		desc := fmt.Sprintf("%s Editor", rd.Spec.Resource.Kind)
		doc := docapi.DocInfo{
			Project: docapi.ProjectInfo{
				Name:        fmt.Sprintf("%s by AppsCode", desc),
				ShortName:   desc,
				URL:         "https://byte.builders",
				Description: desc,
				App:         fmt.Sprintf("a %s", desc),
			},
			Repository: docapi.RepositoryInfo{
				URL:  "https://bundles.byte.builders/ui/",
				Name: "bytebuilders-ui",
			},
			Chart: docapi.ChartInfo{
				Name:          chartName,
				Values:        "-- generate from values file --",
				ValuesExample: "-- generate from values file --",
			},
			Prerequisites: []string{
				"Kubernetes 1.20+",
			},
			Release: docapi.ReleaseInfo{
				Name:      chartName,
				Namespace: metav1.NamespaceDefault,
			},
		}

		data, err := yaml.Marshal(&doc)
		if err != nil {
			return err
		}

		filename := filepath.Join(chartDir, chartName, "doc.yaml")
		err = os.WriteFile(filename, data, 0o644)
		if err != nil {
			return err
		}
	}

	return UpdateEditor(rd, chartName, descriptorDir)
}

func UpdateEditor(rd *rsapi.ResourceDescriptor, chartName string, descriptorDir string) error {
	ed, ok := resourceeditors.LoadDefaultByGVR(rd.Spec.Resource.GroupVersionResource())
	if ok {
		if ed.Spec.UI == nil {
			ed.Spec.UI = &uiapi.UIParameters{}
		}
		ed.Spec.UI.Editor = &releasesapi.ChartSourceRef{
			Name: chartName,
			SourceRef: kmapi.TypedObjectReference{
				APIGroup:  releasesapi.SourceGroupHelmRepository,
				Kind:      releasesapi.SourceKindHelmRepository,
				Namespace: "",
				Name:      repoName,
			},
		}
		ed.Spec.UI.Editor.Version = getDigestOrVersion(repoName, ed.Spec.UI.Editor.Name, chartVersion)
		return WriteResourceEditor(ed, filepath.Join(filepath.Dir(descriptorDir), uiapi.ResourceResourceEditors))
	}
	return nil
}

func GenerateChartMetadata(chartDir, chartName string, rd *rsapi.ResourceDescriptor) error {
	gvr := metav1.GroupVersionResource{
		Group:    rd.Spec.Resource.Group,
		Version:  rd.Spec.Resource.Version,
		Resource: rd.Spec.Resource.Name,
	}
	gvrData, err := json.Marshal(gvr)
	if err != nil {
		panic(err)
	}

	chartMeta := chart.Metadata{
		Name:        chartName,
		Home:        "https://byte.builders",
		Sources:     nil,
		Version:     chartVersion,
		AppVersion:  chartVersion,
		Description: fmt.Sprintf("%s Editor", rd.Spec.Resource.Kind),
		Keywords:    []string{"appscode"},
		Maintainers: []*chart.Maintainer{
			{
				Name:  "appscode",
				Email: "support@appscode.com",
			},
		},
		Icon:        "https://cdn.appscode.com/images/products/bytebuilders/bytebuilders-512x512.png",
		APIVersion:  "v2",
		Condition:   "",
		Deprecated:  false,
		KubeVersion: ">= 1.20.0",
		Type:        "application",
		Annotations: map[string]string{
			"meta.x-helm.dev/editor": string(gvrData),
		},
	}
	data4, err := yaml.Marshal(chartMeta)
	if err != nil {
		return err
	}
	filename := filepath.Join(chartDir, chartName, "Chart.yaml")
	return os.WriteFile(filename, data4, 0o644)
}

func safeGroupName(group string) string {
	if group == "" {
		group = "core"
	}
	group = strings.ReplaceAll(group, ".", "")
	group = strings.ReplaceAll(group, "-", "")
	return group
}

// Impefect
func IsCRD(group string) bool {
	if group == "app.k8s.io" {
		return true
	}
	return strings.ContainsRune(group, '.') &&
		group != "" &&
		!strings.HasSuffix(group, ".k8s.io") &&
		!strings.HasSuffix(group, ".kubernetes.io")
}

func WriteResourceEditor(rd *uiapi.ResourceEditor, dir string) error {
	data, err := yaml.Marshal(rd)
	if err != nil {
		return err
	}

	group := rd.Spec.Resource.Group
	if group == "" {
		group = "core"
	}
	baseDir := filepath.Join(dir, group, rd.Spec.Resource.Version)
	filename := filepath.Join(baseDir, rd.Spec.Resource.Name+".yaml")
	err = os.MkdirAll(filepath.Dir(filename), 0o755)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0o644)
}

func getDigestOrVersion(repo, bin, ver string) string {
	if !chartUseDigest {
		return ver
	}
	if repo != repoName {
		return ver
	}
	digest, err := crane.Digest(fmt.Sprintf("r.byte.builders/charts/%s:%s", bin, ver), crane.WithAuthFromKeychain(authn.DefaultKeychain))
	if err == nil {
		return digest
	}
	return ver
}

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
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	docapi "kubepack.dev/chart-doc-gen/api"
	appapi "kubepack.dev/lib-app/api/v1alpha1"

	"github.com/spf13/cobra"
	"helm.sh/helm/v3/pkg/chart"
	crdv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"kmodules.xyz/resource-metadata/apis/meta/v1alpha1"
	"kmodules.xyz/resource-metadata/hub"
	"sigs.k8s.io/yaml"
)

func NewCmdSimple() *cobra.Command {
	var (
		descriptorDir = "/home/tamal/go/src/kmodules.xyz/resource-metadata/hub/resourcedescriptors/"
		chartDir      = "/home/tamal/go/src/go.bytebuilders.dev/ui-wizards/charts"
		gvr           schema.GroupVersionResource
		all           bool
		skipExisting  bool

		registry = hub.NewRegistryOfKnownResources()
	)
	cmd := &cobra.Command{
		Use:               "simple-chart",
		Short:             `Generate simple chart`,
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if !all {
				return GenerateSimpleEditorChart(chartDir, descriptorDir, gvr, registry, skipExisting)
			}

			registry.Visit(func(key string, rd *v1alpha1.ResourceDescriptor) {
				err := GenerateSimpleEditorChart(chartDir, descriptorDir, rd.Spec.Resource.GroupVersionResource(), registry, skipExisting)
				if err != nil {
					panic(err)
				}
			})
			return nil
		},
	}

	cmd.Flags().StringVar(&chartDir, "chart-dir", chartDir, "Charts dir")
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
	err = os.MkdirAll(tplDir, 0755)
	if err != nil {
		return err
	}

	crdDir := filepath.Join(chartDir, chartName, "crds")
	err = os.MkdirAll(crdDir, 0755)
	if err != nil {
		return err
	}

	err = GenerateChartMetadata(chartDir, chartName, rd)
	if err != nil {
		return err
	}

	if IsCRD(gvr.Group) {
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
						//Subresources:             nil,
						//AdditionalPrinterColumns: nil,
					},
				},
				PreserveUnknownFields: false,
			},
		}
		if strings.HasSuffix(gvr.Group, ".k8s.io") ||
			strings.HasSuffix(gvr.Group, "kubernetes.io") {
			crd.Annotations = map[string]string{
				"api-approved.kubernetes.io": "https://github.com/kubernetes-sigs/application/pull/2",
			}
		}

		filename := filepath.Join(crdDir, fmt.Sprintf("%s_%s.yaml", gvr.Group, gvr.Resource))
		data, err := yaml.Marshal(crd)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(filename, data, 0644)
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
		err = ioutil.WriteFile(schemaFilename, []byte(data), 0644)
		if err != nil {
			return err
		}
	}

	{
		data := fmt.Sprintf(`Get the Stash BackupConfiguration by running the following command:

  kubectl --namespace {{ .Release.Namespace }} get %s.%s {{ .Release.Name }}
`, gvr.Resource, gvr.Group)
		schemaFilename := filepath.Join(chartDir, chartName, "templates", "NOTES.txt")
		err = ioutil.WriteFile(schemaFilename, []byte(data), 0644)
		if err != nil {
			return err
		}
	}

	{
		if rd.Spec.Validation != nil && rd.Spec.Validation.OpenAPIV3Schema != nil {
			data3, err := yaml.Marshal(rd.Spec.Validation.OpenAPIV3Schema)
			if err != nil {
				return err
			}
			schemaFilename := filepath.Join(chartDir, chartName, "values.openapiv3_schema.yaml")
			err = ioutil.WriteFile(schemaFilename, data3, 0644)
			if err != nil {
				return err
			}
		}
	}

	{
		v := appapi.SimpleValue{
			TypeMeta: metav1.TypeMeta{
				APIVersion: fmt.Sprintf("%s/%s", rd.Spec.Resource.Group, rd.Spec.Resource.Version),
				Kind:       rd.Spec.Resource.Kind,
			},
			ObjectMeta: appapi.ObjectMeta{
				Name: strings.ToLower(rd.Spec.Resource.Kind),
			},
		}
		if rd.Spec.Resource.Scope == v1alpha1.NamespaceScoped {
			v.ObjectMeta.Namespace = "default"
		}

		data, err := yaml.Marshal(v)
		if err != nil {
			panic(err)
		}

		filename := filepath.Join(chartDir, chartName, "values.yaml")
		err = ioutil.WriteFile(filename, data, 0644)
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
				URL:  "https://bundles.bytebuilders.dev/ui/",
				Name: "bytebuilders-ui",
			},
			Chart: docapi.ChartInfo{
				Name:          chartName,
				Version:       "v0.1.0",
				Values:        "-- generate from values file --",
				ValuesExample: "-- generate from values file --",
			},
			Prerequisites: []string{
				"Kubernetes 1.14+",
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
		err = ioutil.WriteFile(filename, data, 0644)
		if err != nil {
			return err
		}
	}

	if rd.Spec.UI == nil {
		rd.Spec.UI = &v1alpha1.UIParameters{
			Options: nil,
			Editor: &v1alpha1.ChartRepoRef{
				URL:     "https://bundles.byte.builders/ui/",
				Name:    chartName,
				Version: "v0.1.0",
			},
		}
		return UpdateDescriptor(rd, descriptorDir)
	}

	return nil
}

func GenerateChartMetadata(chartDir, chartName string, rd *v1alpha1.ResourceDescriptor) error {
	chartMeta := chart.Metadata{
		Name:        chartName,
		Home:        "https://byte.builders",
		Sources:     nil,
		Version:     "v0.1.0",
		AppVersion:  "v0.1.0",
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
		KubeVersion: ">= 1.14.0",
		Type:        "application",
	}
	data4, err := yaml.Marshal(chartMeta)
	if err != nil {
		return err
	}
	filename := filepath.Join(chartDir, chartName, "Chart.yaml")
	return ioutil.WriteFile(filename, data4, 0644)
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

func UpdateDescriptor(rd *v1alpha1.ResourceDescriptor, dir string) error {
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
	return ioutil.WriteFile(filename, data, 0644)
}

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

package fusion

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	docapi "kubepack.dev/chart-doc-gen/api"
	appapi "kubepack.dev/lib-app/api/v1alpha1"
	"kubepack.dev/lib-app/pkg/editor"

	"github.com/Masterminds/sprig/v3"
	"github.com/spf13/cobra"
	ksets "gomodules.xyz/sets/kubernetes"
	y3 "gopkg.in/yaml.v3"
	"helm.sh/helm/v3/pkg/chart"
	crdv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/sets"
	"kmodules.xyz/client-go/tools/parser"
	"kmodules.xyz/resource-metadata/apis/meta/v1alpha1"
	"kmodules.xyz/resource-metadata/hub"
	"sigs.k8s.io/yaml"
)

var (
	sampleDir      = ""
	sampleName     = ""
	chartDir       = ""
	chartName      = ""
	gvr            schema.GroupVersionResource
	resourceSchema = crdv1.JSONSchemaProps{
		Type:       "object",
		Properties: map[string]crdv1.JSONSchemaProps{},
	}
	resourceValues = map[string]*unstructured.Unstructured{}
	registry       = hub.NewRegistryOfKnownResources()
	resourceKeys   = sets.NewString()
)

func NewCmdFuse() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "fuse-chart",
		Short:             `Fuse YAMLs`,
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			rd, err := registry.LoadByGVR(gvr)
			if err != nil {
				return err
			}

			chartName = fmt.Sprintf("%s-%s-editor", safeGroupName(rd.Spec.Resource.Group), strings.ToLower(rd.Spec.Resource.Kind))

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

			gkSet := ksets.NewGroupKind()

			err = parser.ProcessPath(sampleDir, func(ri parser.ResourceInfo) error {
				gkSet.Insert(ri.Object.GetObjectKind().GroupVersionKind().GroupKind())

				rsKey, err := editor.ResourceKey(ri.Object.GetAPIVersion(), ri.Object.GetKind(), sampleName, ri.Object.GetName())
				if err != nil {
					return err
				}
				resourceKeys.Insert(rsKey)
				_, _, rsFilename := editor.ResourceFilename(ri.Object.GetAPIVersion(), ri.Object.GetKind(), sampleName, ri.Object.GetName())

				// values
				cp := ri.Object.DeepCopy()
				delete(cp.Object, "status")
				cp.Object["metadata"] = appapi.ObjectMeta{
					Name:      ri.Object.GetName(),
					Namespace: ri.Object.GetNamespace(),
				}
				resourceValues[rsKey] = cp

				// schema
				gvr, err := registry.GVR(ri.Object.GetObjectKind().GroupVersionKind())
				if err != nil {
					return err
				}
				descriptor, err := registry.LoadByGVR(gvr)
				if err != nil {
					return err
				}

				if descriptor.Spec.Validation != nil && descriptor.Spec.Validation.OpenAPIV3Schema != nil {
					delete(descriptor.Spec.Validation.OpenAPIV3Schema.Properties, "status")
					resourceSchema.Properties[rsKey] = *descriptor.Spec.Validation.OpenAPIV3Schema.DeepCopy()
				}

				if IsCRD(gvr.Group) {
					// Do not update the hub registry
					descriptor = descriptor.DeepCopy()
					if descriptor.Spec.Validation != nil && descriptor.Spec.Validation.OpenAPIV3Schema != nil {
						descriptor.Spec.Validation.OpenAPIV3Schema.Properties["metadata"] = crdv1.JSONSchemaProps{
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
								Plural:   descriptor.Spec.Resource.Name,
								Singular: strings.ToLower(descriptor.Spec.Resource.Kind),
								// ShortNames: nil,
								Kind:     descriptor.Spec.Resource.Kind,
								ListKind: descriptor.Spec.Resource.Kind + "List",
								// Categories: nil,
							},
							Scope: crdv1.ResourceScope(string(descriptor.Spec.Resource.Scope)),
							Versions: []crdv1.CustomResourceDefinitionVersion{
								{
									Name:    descriptor.Spec.Resource.Version,
									Served:  true,
									Storage: true,
									Schema:  descriptor.Spec.Validation,
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
							"api-approved.kubernetes.io": "https://github.com/kubernetes-sigs/application/pull/2",
						}
					}

					filename := filepath.Join(crdDir, fmt.Sprintf("%s_%s.yaml", gvr.Group, gvr.Resource))
					data, err := yaml.Marshal(crd)
					if err != nil {
						return err
					}
					err = ioutil.WriteFile(filename, data, 0o644)
					if err != nil {
						return err
					}
				}

				// templates
				filename := filepath.Join(tplDir, rsFilename+".yaml")
				f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644)
				if err != nil {
					return err
				}
				defer f.Close()

				objModel := appapi.ObjectModel{
					Key:    rsKey,
					Object: ri.Object,
				}
				modelJSON, err := json.Marshal(objModel)
				if err != nil {
					return err
				}

				var data map[string]interface{}
				err = json.Unmarshal(modelJSON, &data)
				if err != nil {
					panic(err)
				}

				funcMap := sprig.TxtFuncMap()
				funcMap["toYaml"] = toYAML
				funcMap["toJson"] = toJSON
				tpl := template.Must(template.New("resourceTemplate").Funcs(funcMap).Parse(resourceTemplate))
				err = tpl.Execute(f, &data)
				if err != nil {
					return err
				}

				return nil
			})
			if err != nil {
				return err
			}

			gks := make([]metav1.GroupKind, gkSet.Len())
			for i, gk := range gkSet.List() {
				gks[i] = metav1.GroupKind{
					Group: gk.Group,
					Kind:  gk.Kind,
				}
			}
			err = GenerateChartMetadata(rd, gks)
			if err != nil {
				return err
			}

			{
				var chartSchema crdv1.JSONSchemaProps
				err = yaml.Unmarshal([]byte(valuesMetadataSchema), &chartSchema)
				if err != nil {
					return err
				}
				chartSchema.Properties["resources"] = resourceSchema
				removeDescription(&chartSchema)
				data3, err := yaml.Marshal(chartSchema)
				if err != nil {
					return err
				}
				schemaFilename := filepath.Join(chartDir, chartName, "values.openapiv3_schema.yaml")
				err = ioutil.WriteFile(schemaFilename, data3, 0o644)
				if err != nil {
					return err
				}
			}

			{
				data, err := yaml.Marshal(resourceValues)
				if err != nil {
					panic(err)
				}

				var root y3.Node
				err = y3.Unmarshal(data, &root)
				if err != nil {
					return err
				}
				addDocComments(&root)

				rd, err := registry.LoadByGVR(gvr)
				if err != nil {
					return err
				}

				values := map[string]interface{}{
					"metadata": appapi.Metadata{
						Resource: rd.Spec.Resource,
						Release: appapi.ObjectMeta{
							Name:      "RELEASE-NAME",
							Namespace: "default",
						},
					},
					"resources": root.Content[0],
				}

				var buf bytes.Buffer
				enc := y3.NewEncoder(&buf)
				enc.SetIndent(2)
				defer enc.Close()
				err = enc.Encode(&values)
				if err != nil {
					return err
				}

				filename := filepath.Join(chartDir, chartName, "values.yaml")
				err = ioutil.WriteFile(filename, buf.Bytes(), 0o644)
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
						"Kubernetes 1.16+",
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
				err = ioutil.WriteFile(filename, data, 0o644)
				if err != nil {
					return err
				}
			}

			return nil
		},
	}

	cmd.Flags().StringVar(&sampleDir, "sample-dir", sampleDir, "Sample dir")
	cmd.Flags().StringVar(&sampleName, "sample-name", sampleName, "Sample name used in yamls")
	cmd.Flags().StringVar(&chartDir, "chart-dir", chartDir, "Charts dir")

	cmd.Flags().StringVar(&gvr.Group, "resource.group", gvr.Group, "Resource api group")
	cmd.Flags().StringVar(&gvr.Version, "resource.version", gvr.Version, "Resource api version")
	cmd.Flags().StringVar(&gvr.Resource, "resource.name", gvr.Resource, "Resource plural")

	return cmd
}

func GenerateChartMetadata(rd *v1alpha1.ResourceDescriptor, gks []metav1.GroupKind) error {
	gvr := metav1.GroupVersionResource{
		Group:    rd.Spec.Resource.Group,
		Version:  rd.Spec.Resource.Version,
		Resource: rd.Spec.Resource.Name,
	}
	gvrData, err := json.Marshal(gvr)
	if err != nil {
		panic(err)
	}
	//if rd.Spec.Resource.Group == "kubedb.com" {
	//	gks = []metav1.GroupKind{
	//		{
	//			Group: rd.Spec.Resource.Group,
	//			Kind:  rd.Spec.Resource.Kind,
	//		},
	//		{
	//			Group: "",
	//			Kind:  "Secret",
	//		},
	//		{
	//			Group: "cert-manager.io",
	//			Kind:  "Issuer",
	//		},
	//		{
	//			Group: "monitoring.coreos.com",
	//			Kind:  "ServiceMonitor",
	//		},
	//		{
	//			Group: "stash.appscode.com",
	//			Kind:  "Repository",
	//		},
	//		{
	//			Group: "stash.appscode.com",
	//			Kind:  "BackupConfiguration",
	//		},
	//		{
	//			Group: "stash.appscode.com",
	//			Kind:  "RestoreSession",
	//		},
	//	}
	//}
	sort.Slice(gks, func(i, j int) bool {
		if gks[i].Group == gks[j].Group {
			return gks[i].Kind < gks[j].Kind
		}
		return gks[i].Group < gks[j].Group
	})

	gkData, err := yaml.Marshal(gks)
	if err != nil {
		panic(err)
	}

	chartMeta := chart.Metadata{
		Name:        chartName,
		Home:        "https://byte.builders",
		Sources:     nil,
		Version:     "v0.2.0",
		AppVersion:  "v0.2.0",
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
		Annotations: map[string]string{
			"meta.x-helm.dev/editor":    string(gvrData),
			"meta.x-helm.dev/resources": string(gkData),
		},
	}
	data4, err := yaml.Marshal(chartMeta)
	if err != nil {
		return err
	}
	filename := filepath.Join(chartDir, chartName, "Chart.yaml")
	return ioutil.WriteFile(filename, data4, 0o644)
}

// toYAML takes an interface, marshals it to yaml, and returns a string. It will
// always return a string, even on marshal error (empty string).
//
// This is designed to be called from a template.
func toYAML(v interface{}) string {
	data, err := yaml.Marshal(v)
	if err != nil {
		// Swallow errors inside of a template.
		return ""
	}
	return strings.TrimSuffix(string(data), "\n")
}

// toJSON takes an interface, marshals it to json, and returns a string. It will
// always return a string, even on marshal error (empty string).
//
// This is designed to be called from a template.
func toJSON(v interface{}) string {
	data, err := json.Marshal(v)
	if err != nil {
		// Swallow errors inside of a template.
		return ""
	}
	return string(data)
}

func addDocComments(node *y3.Node) {
	if node.Tag == "!!str" && resourceKeys.Has(node.Value) {
		node.LineComment = "# +doc-gen:break"
	}
	for i := range node.Content {
		addDocComments(node.Content[i])
	}
}

// removeDescription removes defaults from apiextensions.k8s.io/v1beta1 CRD definition.
func removeDescription(schema *crdv1.JSONSchemaProps) {
	if schema == nil {
		return
	}

	schema.Description = ""

	if schema.Items != nil {
		removeDescription(schema.Items.Schema)

		for idx := range schema.Items.JSONSchemas {
			removeDescription(&schema.Items.JSONSchemas[idx])
		}
	}

	for idx := range schema.AllOf {
		removeDescription(&schema.AllOf[idx])
	}
	for idx := range schema.OneOf {
		removeDescription(&schema.OneOf[idx])
	}
	for idx := range schema.AnyOf {
		removeDescription(&schema.AnyOf[idx])
	}
	if schema.Not != nil {
		removeDescription(schema.Not)
	}
	for key, prop := range schema.Properties {
		removeDescription(&prop)
		schema.Properties[key] = prop
	}
	if schema.AdditionalProperties != nil {
		removeDescription(schema.AdditionalProperties.Schema)
	}
	for key, prop := range schema.PatternProperties {
		removeDescription(&prop)
		schema.PatternProperties[key] = prop
	}
	for key, prop := range schema.Dependencies {
		removeDescription(prop.Schema)
		schema.Dependencies[key] = prop
	}
	if schema.AdditionalItems != nil {
		removeDescription(schema.AdditionalItems.Schema)
	}
	for key, prop := range schema.Definitions {
		removeDescription(&prop)
		schema.Definitions[key] = prop
	}
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

func safeGroupName(group string) string {
	group = strings.ReplaceAll(group, ".", "")
	group = strings.ReplaceAll(group, "-", "")
	return group
}

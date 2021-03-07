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
	"strings"
	"text/template"

	docapi "kubepack.dev/chart-doc-gen/api"
	appapi "kubepack.dev/lib-app/api/v1alpha1"
	"kubepack.dev/lib-app/pkg/editor"

	"github.com/Masterminds/sprig"
	"github.com/spf13/cobra"
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
			err = os.MkdirAll(tplDir, 0755)
			if err != nil {
				return err
			}

			crdDir := filepath.Join(chartDir, chartName, "crds")
			err = os.MkdirAll(crdDir, 0755)
			if err != nil {
				return err
			}

			err = GenerateChartMetadata(rd)
			if err != nil {
				return err
			}

			err = parser.ProcessDir(sampleDir, func(obj *unstructured.Unstructured) error {
				rsKey, err := editor.ResourceKey(obj.GetAPIVersion(), obj.GetKind(), sampleName, obj.GetName())
				if err != nil {
					return err
				}
				resourceKeys.Insert(rsKey)
				_, _, rsFilename := editor.ResourceFilename(obj.GetAPIVersion(), obj.GetKind(), sampleName, obj.GetName())

				// values
				cp := obj.DeepCopy()
				delete(cp.Object, "status")
				cp.Object["metadata"] = appapi.ObjectMeta{
					Name:      obj.GetName(),
					Namespace: obj.GetNamespace(),
				}
				resourceValues[rsKey] = cp

				// schema
				gvr, err := registry.GVR(obj.GetObjectKind().GroupVersionKind())
				if err != nil {
					return err
				}
				descriptor, err := registry.LoadByGVR(gvr)
				if err != nil {
					return err
				}

				if IsCRD(gvr.Group) {
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

				if descriptor.Spec.Validation != nil && descriptor.Spec.Validation.OpenAPIV3Schema != nil {
					delete(descriptor.Spec.Validation.OpenAPIV3Schema.Properties, "status")
					resourceSchema.Properties[rsKey] = *descriptor.Spec.Validation.OpenAPIV3Schema
				}

				// templates
				filename := filepath.Join(tplDir, rsFilename+".yaml")
				f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
				if err != nil {
					return err
				}
				defer f.Close()

				objModel := appapi.ObjectModel{
					Key:    rsKey,
					Object: obj,
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
				err = ioutil.WriteFile(schemaFilename, data3, 0644)
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
				err = ioutil.WriteFile(filename, buf.Bytes(), 0644)
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

func GenerateChartMetadata(rd *v1alpha1.ResourceDescriptor) error {
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

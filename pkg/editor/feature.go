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
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"helm.sh/helm/v3/pkg/chart"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	uiapi "kmodules.xyz/resource-metadata/apis/ui/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func UpdateFeatureValues(kc client.Client, chrt *chart.Chart, vals map[string]any) (map[string]any, error) {
	var gvr metav1.GroupVersionResource

	if data, ok := chrt.Metadata.Annotations["meta.x-helm.dev/editor"]; ok && data != "" {
		if err := json.Unmarshal([]byte(data), &gvr); err != nil {
			return nil, err
		}
	} else {
		return vals, nil
	}

	fsGVR := metav1.GroupVersionResource{
		Group:    "ui.k8s.appscode.com",
		Version:  "v1alpha1",
		Resource: "featuresets",
	}
	if gvr != fsGVR {
		return vals, nil
	}

	if resources, ok, err := unstructured.NestedMap(vals, "resources"); err == nil && ok {
		for k, o := range resources {
			// helmToolkitFluxcdIoHelmRelease_kubestash
			if !strings.HasPrefix(k, "helmToolkitFluxcdIoHelmRelease_") {
				continue
			}

			obj := o.(map[string]interface{})
			featureName, found, err := unstructured.NestedString(obj, "metadata", "name")
			if err != nil {
				return nil, errors.Wrap(err, "can't detect feature name")
			} else if !found {
				return nil, fmt.Errorf("feature name not found for key %s", k)
			}

			var feature uiapi.Feature
			err = kc.Get(context.TODO(), client.ObjectKey{Name: featureName}, &feature)
			if err == nil {
				vals, err = setChartInfo(&feature, k, vals)
				if err != nil {
					return nil, err
				}
			}
		}
	}
	return vals, nil
}

func setChartInfo(feature *uiapi.Feature, featureKey string, values map[string]interface{}) (map[string]any, error) {
	err := unstructured.SetNestedField(values, feature.Spec.Chart.Name, "resources", featureKey, "spec", "chart", "spec", "chart")
	if err != nil {
		return nil, err
	}
	if feature.Spec.Chart.Version != "" {
		err = unstructured.SetNestedField(values, feature.Spec.Chart.Version, "resources", featureKey, "spec", "chart", "spec", "version")
		if err != nil {
			return nil, err
		}
	} else {
		unstructured.RemoveNestedField(values, "resources", featureKey, "spec", "chart", "spec", "version")
	}
	err = unstructured.SetNestedField(values, feature.Spec.Chart.SourceRef.Kind, "resources", featureKey, "spec", "chart", "spec", "sourceRef", "kind")
	if err != nil {
		return nil, err
	}
	err = unstructured.SetNestedField(values, feature.Spec.Chart.SourceRef.Name, "resources", featureKey, "spec", "chart", "spec", "sourceRef", "name")
	if err != nil {
		return nil, err
	}
	err = unstructured.SetNestedField(values, feature.Spec.Chart.SourceRef.Namespace, "resources", featureKey, "spec", "chart", "spec", "sourceRef", "namespace")
	if err != nil {
		return nil, err
	}

	err = unstructured.SetNestedField(values, feature.Spec.Chart.Namespace, "resources", featureKey, "spec", "targetNamespace")
	if err != nil {
		return nil, err
	}
	err = unstructured.SetNestedField(values, feature.Spec.Chart.Namespace, "resources", featureKey, "spec", "storageNamespace")
	if err != nil {
		return nil, err
	}

	if feature.Spec.Values != nil {
		fvalues := map[string]interface{}{}
		if err = json.Unmarshal(feature.Spec.Values.Raw, &fvalues); err != nil {
			return nil, err
		}
		err = unstructured.SetNestedField(values, fvalues, "resources", featureKey, "spec", "values")
		if err != nil {
			return nil, err
		}
	}

	return values, nil
}

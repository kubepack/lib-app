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

			var feature unstructured.Unstructured
			feature.SetAPIVersion(fsGVR.Group + "/" + fsGVR.Version)
			feature.SetKind("Feature")
			err = kc.Get(context.TODO(), client.ObjectKey{Name: featureName}, &feature)
			if err == nil {
				defVal, found, err := unstructured.NestedFieldNoCopy(feature.UnstructuredContent(), "spec", "values")
				if err == nil && found {
					err = unstructured.SetNestedField(obj, defVal, "spec", "values")
					if err != nil {
						return nil, err
					}
					resources[k] = obj
				}
			}
		}
		err = unstructured.SetNestedField(vals, resources, "resources")
		if err != nil {
			return nil, err
		}
	}
	return vals, nil
}

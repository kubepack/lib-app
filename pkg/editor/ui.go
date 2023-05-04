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
	"fmt"

	"kubepack.dev/lib-helm/pkg/repo"
	"kubepack.dev/lib-helm/pkg/storage/driver"

	"kmodules.xyz/resource-metadata/hub/resourceeditors"
	"sigs.k8s.io/controller-runtime/pkg/client"
	driversapi "x-helm.dev/apimachinery/apis/drivers/v1alpha1"
	releasesapi "x-helm.dev/apimachinery/apis/releases/v1alpha1"
)

func CreateAppReleaseIfMissing(kc client.Client, reg *repo.Registry, model releasesapi.Metadata) (*driversapi.AppRelease, error) {
	var app driversapi.AppRelease
	err := kc.Get(context.TODO(), client.ObjectKey{Namespace: model.Release.Namespace, Name: model.Release.Name}, &app)
	if err == nil || client.IgnoreNotFound(err) != nil {
		return &app, err // err == nil, means AppRelease exists
	}

	// apprelease not found, so create it.

	ed, ok := resourceeditors.LoadByResourceID(kc, &model.Resource)
	if !ok {
		return nil, fmt.Errorf("failed to load resource editor for %+v", model.Resource)
	}

	if ed.Spec.UI.Editor == nil {
		return nil, fmt.Errorf("missing editor chart for %+v", ed.Spec.Resource.GroupVersionKind())
	}
	chartRef := *ed.Spec.UI.Editor

	if chartRef.SourceRef.Namespace == "" {
		ns, err := DefaultSourceRefNamespace(kc, chartRef.SourceRef.Name)
		if err != nil {
			return nil, err
		}
		chartRef.SourceRef.Namespace = ns
	}

	chrt, err := reg.GetChart(chartRef)
	if err != nil {
		return nil, err
	}

	genApp, err := driver.GenerateAppReleaseObject(chrt.Chart, model)
	if err != nil {
		return nil, err
	}

	err = kc.Create(context.TODO(), genApp)
	return genApp, client.IgnoreAlreadyExists(err)
}

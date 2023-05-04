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
	"fmt"

	"kubepack.dev/lib-app/pkg/editor"
	actionx "kubepack.dev/lib-helm/pkg/action"
	"kubepack.dev/lib-helm/pkg/repo"
	"kubepack.dev/lib-helm/pkg/storage/driver"

	kmapi "kmodules.xyz/client-go/api/v1"
	"kmodules.xyz/resource-metadata/hub/resourceeditors"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/yaml"
	releasesapi "x-helm.dev/apimachinery/apis/releases/v1alpha1"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	model := releasesapi.Metadata{
		Resource: kmapi.ResourceID{
			Group:   "kubedb.com",
			Version: "v1alpha2",
			Name:    "",
			Kind:    "MariaDB",
			Scope:   "",
		},
		Release: releasesapi.ObjectMeta{
			Name:      "mariadb-test",
			Namespace: "demo",
		},
	}

	cfg, err := config.GetConfig()
	if err != nil {
		return err
	}
	kc, err := actionx.NewUncachedClientForConfig(cfg)
	if err != nil {
		return err
	}
	reg := repo.NewRegistry(kc, repo.DefaultDiskCache())

	ed, ok := resourceeditors.LoadByResourceID(kc, &model.Resource)
	if !ok {
		return fmt.Errorf("failed to load resource editor for %+v", model.Resource)
	}

	if ed.Spec.UI.Editor == nil {
		return fmt.Errorf("missing editor chart for %+v", ed.Spec.Resource.GroupVersionKind())
	}
	chartRef := *ed.Spec.UI.Editor

	if chartRef.SourceRef.Namespace == "" {
		ns, err := editor.DefaultSourceRefNamespace(kc, chartRef.SourceRef.Name)
		if err != nil {
			return err
		}
		chartRef.SourceRef.Namespace = ns
	}

	chrt, err := reg.GetChart(chartRef)
	if err != nil {
		return err
	}

	appr, err := driver.GenerateAppReleaseObject(chrt.Chart, model)
	if err != nil {
		return err
	}

	data, err := yaml.Marshal(appr)
	if err != nil {
		return err
	}
	fmt.Println(string(data))

	return nil
}

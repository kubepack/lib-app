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

package handler

import (
	"encoding/json"
	"errors"

	"kubepack.dev/kubepack/pkg/lib"
	appapi "kubepack.dev/lib-app/api/v1alpha1"
	"kubepack.dev/lib-app/pkg/editor"
	"kubepack.dev/lib-app/pkg/lib/action"

	jsonpatch "github.com/evanphx/json-patch"
	"helm.sh/helm/v3/pkg/release"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
	meta_util "kmodules.xyz/client-go/meta"
	"kmodules.xyz/resource-metadata/hub"
)

func ApplyResource(f cmdutil.Factory, model unstructured.Unstructured, skipCRds bool) (*release.Release, error) {
	var tm appapi.ModelMetadata
	err := meta_util.DecodeObject(model.Object, &tm)
	if err != nil {
		return nil, errors.New("failed to parse Metadata for values")
	}

	rd, err := hub.NewRegistryOfKnownResources().LoadByGVR(schema.GroupVersionResource{
		Group:    tm.Resource.Group,
		Version:  tm.Resource.Version,
		Resource: tm.Resource.Name,
	})
	if err != nil {
		return nil, err
	}

	applier, err := action.NewApplier(f, tm.Release.Namespace, "applications")
	if err != nil {
		return nil, err
	}

	applier.WithRegistry(lib.DefaultRegistry)
	opts2 := action.NewApplyOptions()
	opts2.ChartURL = rd.Spec.UI.Editor.URL
	opts2.ChartName = rd.Spec.UI.Editor.Name
	opts2.Version = rd.Spec.UI.Editor.Version
	if _, ok := model.Object["patch"]; ok {
		// NOTE: Makes an assumption that this is a "edit" apply
		cfg, err := f.ToRESTConfig()
		if err != nil {
			return nil, err
		}
		tpl, err := editor.LoadEditorModel(cfg, lib.DefaultRegistry, appapi.ModelMetadata{
			Metadata: tm.Metadata,
		})
		if err != nil {
			return nil, err
		}

		p3 := struct {
			Patch jsonpatch.Patch `json:"patch"`
		}{}
		err = meta_util.DecodeObject(model.Object, &p3)
		if err != nil {
			return nil, err
		}

		original, err := json.Marshal(tpl.Values.Object)
		if err != nil {
			return nil, err
		}

		modified, err := p3.Patch.Apply(original)
		if err != nil {
			return nil, err
		}

		var mod map[string]interface{}
		err = json.Unmarshal(modified, &mod)
		if err != nil {
			return nil, err
		}
		opts2.Values = mod
	} else {
		opts2.Values = model.Object
	}

	opts2.CreateNamespace = true // TODO?
	opts2.DryRun = false
	opts2.DisableHooks = false
	opts2.Replace = false
	opts2.Wait = false
	opts2.Timeout = 0
	opts2.Description = "Apply editor"
	opts2.Devel = false
	opts2.Namespace = tm.Release.Namespace
	opts2.ReleaseName = tm.Release.Name
	opts2.Atomic = false
	opts2.SkipCRDs = skipCRds
	opts2.SubNotes = false
	opts2.DisableOpenAPIValidation = false
	opts2.IncludeCRDs = false

	opts2.RefillMetadata = true

	applier.WithOptions(opts2)

	return applier.Run()
}

func DeleteResource(f cmdutil.Factory, release appapi.ObjectMeta) (*release.UninstallReleaseResponse, error) {
	cmd, err := action.NewUninstaller(f, release.Namespace, "applications")
	if err != nil {
		return nil, err
	}

	cmd.WithReleaseName(release.Name)
	cmd.WithOptions(action.UninstallOptions{
		DisableHooks: false,
		DryRun:       false,
		KeepHistory:  false,
		Timeout:      0,
	})
	return cmd.Run()
}

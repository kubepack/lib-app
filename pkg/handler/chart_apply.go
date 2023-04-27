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
	"context"
	"encoding/json"
	"fmt"

	"kubepack.dev/lib-app/pkg/editor"
	actionx "kubepack.dev/lib-helm/pkg/action"
	"kubepack.dev/lib-helm/pkg/repo"
	"kubepack.dev/lib-helm/pkg/storage/driver"
	"kubepack.dev/lib-helm/pkg/values"

	jsonpatch "github.com/evanphx/json-patch/v5"
	"github.com/pkg/errors"
	ha "helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/release"
	"k8s.io/apimachinery/pkg/types"
	apiregistrationapi "k8s.io/kube-aggregator/pkg/apis/apiregistration/v1"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
	meta_util "kmodules.xyz/client-go/meta"
	"kmodules.xyz/resource-metadata/hub/resourceeditors"
	"sigs.k8s.io/controller-runtime/pkg/client"
	releasesapi "x-helm.dev/apimachinery/apis/releases/v1alpha1"
)

func ApplyResourceEditor(f cmdutil.Factory, reg repo.IRegistry, model map[string]interface{}, skipCRds bool, log ...ha.DebugLog) (*release.Release, error) {
	var tm releasesapi.ModelMetadata
	err := meta_util.DecodeObject(model, &tm)
	if err != nil {
		return nil, errors.New("failed to parse Metadata for values")
	}

	kc, err := actionx.NewUncachedClient(f)
	if err != nil {
		return nil, err
	}

	ed, ok := resourceeditors.LoadByResourceID(kc, &tm.Resource)
	if !ok {
		return nil, fmt.Errorf("failed to load resource editor for %+v", tm.Resource)
	}

	deployer, err := actionx.NewDeployer(f, tm.Release.Namespace, driver.AppReleasesDriverName, log...)
	if err != nil {
		return nil, err
	}

	deployer.WithRegistry(reg)

	if ed.Spec.UI.Editor == nil {
		return nil, fmt.Errorf("missing editor chart for %+v", ed.Spec.Resource.GroupVersionKind())
	}

	return applyResource(f, reg, *ed.Spec.UI.Editor, model, skipCRds, log...)
}

func getSourceRefNamespace(kc client.Client, sourceName string) (string, error) {
	// k get apiservices v1alpha1.meta.k8s.appscode.com -o yaml
	var apisvc apiregistrationapi.APIService
	apisvcName := "v1alpha1.meta.k8s.appscode.com"
	err := kc.Get(context.TODO(), types.NamespacedName{Name: apisvcName}, &apisvc)
	if err != nil {
		return "", errors.Wrapf(err, "failed to detect namespace for HelmRepository %s", sourceName)
	}
	if apisvc.Spec.Service == nil {
		return "", errors.Wrapf(err, "failed to detect namespace for HelmRepository %s from Local APIService %s", sourceName, apisvcName)
	}

	return apisvc.Spec.Service.Namespace, nil
}

func ApplyResource(f cmdutil.Factory, reg repo.IRegistry, chartRef releasesapi.ChartSourceRef, model map[string]interface{}, skipCRds bool, log ...ha.DebugLog) (*release.Release, error) {
	return applyResource(f, reg, chartRef, model, skipCRds, log...)
}

func applyResource(f cmdutil.Factory, reg repo.IRegistry, chartRef releasesapi.ChartSourceRef, model map[string]interface{}, skipCRds bool, log ...ha.DebugLog) (*release.Release, error) {
	var tm releasesapi.ModelMetadata
	err := meta_util.DecodeObject(model, &tm)
	if err != nil {
		return nil, errors.New("failed to parse Metadata for values")
	}

	kc, err := actionx.NewUncachedClient(f)
	if err != nil {
		return nil, err
	}

	deployer, err := actionx.NewDeployer(f, tm.Release.Namespace, driver.AppReleasesDriverName, log...)
	if err != nil {
		return nil, err
	}

	deployer.WithRegistry(reg)

	var opts actionx.DeployOptions

	if chartRef.SourceRef.Namespace == "" {
		ns, err := getSourceRefNamespace(kc, chartRef.SourceRef.Name)
		if err != nil {
			return nil, err
		}

		chartRef.SourceRef.Namespace = ns
	}
	opts.ChartSourceFlatRef.FromAPIObject(chartRef)

	var vals map[string]interface{}
	if _, ok := model["patch"]; ok {
		// NOTE: Makes an assumption that this is a "edit" apply
		tpl, err := editor.LoadResourceEditorModel(kc, reg, releasesapi.ModelMetadata{
			Metadata: tm.Metadata,
		})
		if err != nil {
			return nil, err
		}

		p3 := struct {
			Patch jsonpatch.Patch `json:"patch"`
		}{}

		data, err := json.Marshal(model)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(data, &p3)
		if err != nil {
			return nil, err
		}

		original, err := json.Marshal(tpl.Values.Object)
		if err != nil {
			return nil, err
		}

		modified, err := p3.Patch.ApplyWithOptions(original, &jsonpatch.ApplyOptions{
			SupportNegativeIndices:   jsonpatch.SupportNegativeIndices,
			AccumulatedCopySizeLimit: jsonpatch.AccumulatedCopySizeLimit,
			AllowMissingPathOnRemove: true,
			EnsurePathExistsOnAdd:    false,
		})
		if err != nil {
			return nil, err
		}

		var mod map[string]interface{}
		err = json.Unmarshal(modified, &mod)
		if err != nil {
			return nil, err
		}
		vals = mod
	} else {
		vals = model
	}
	opts.Options = values.Options{
		ReplaceValues: vals,
	}

	opts.DryRun = false
	opts.DisableHooks = false
	opts.Replace = false
	opts.Wait = false
	opts.Timeout = 0
	opts.Description = "Apply editor"
	opts.Devel = false
	opts.Namespace = tm.Release.Namespace
	opts.ReleaseName = tm.Release.Name
	opts.Atomic = false
	opts.SkipCRDs = skipCRds
	opts.SubNotes = false
	opts.DisableOpenAPIValidation = false
	opts.IncludeCRDs = false

	deployer.WithOptions(opts)

	rls, _, err := deployer.Run()
	return rls, err
}

func DeleteResource(f cmdutil.Factory, release releasesapi.ObjectMeta, log ...ha.DebugLog) (*release.UninstallReleaseResponse, error) {
	cmd, err := actionx.NewUninstaller(f, release.Namespace, driver.AppReleasesDriverName, log...)
	if err != nil {
		return nil, err
	}

	cmd.WithReleaseName(release.Name)
	cmd.WithOptions(actionx.UninstallOptions{
		DisableHooks: false,
		DryRun:       false,
		KeepHistory:  false,
		Timeout:      0,
	})
	return cmd.Run()
}

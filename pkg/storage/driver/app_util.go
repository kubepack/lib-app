/*
Copyright The Helm Authors.

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

package driver

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"kubepack.dev/kubepack/apis"
	"kubepack.dev/kubepack/pkg/lib"
	api "kubepack.dev/lib-app/api/v1alpha1"
	"kubepack.dev/lib-app/pkg/editor"

	"github.com/gabriel-vasile/mimetype"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/release"
	rspb "helm.sh/helm/v3/pkg/release"
	helmtime "helm.sh/helm/v3/pkg/time"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/restmapper"
	"sigs.k8s.io/application/api/app/v1beta1"
)

// newApplicationSecretsObject constructs a kubernetes Application object
// to store a release. Each configmap data entry is the base64
// encoded gzipped string of a release.
//
// The following labels are used within each configmap:
//
//    "modifiedAt"     - timestamp indicating when this configmap was last modified. (set in Update)
//    "createdAt"      - timestamp indicating when this configmap was created. (set in Create)
//    "version"        - version of the release.
//    "status"         - status of the release (see pkg/release/status.go for variants)
//    "owner"          - owner of the configmap, currently "helm".
//    "name"           - name of the release.
//
func newApplicationObject(rls *rspb.Release) *v1beta1.Application {
	const owner = "helm"

	// create and return configmap object
	obj := &v1beta1.Application{
		ObjectMeta: metav1.ObjectMeta{
			Name:      rls.Name,
			Namespace: rls.Namespace,
			Labels: map[string]string{
				"name":    rls.Name,
				"owner":   owner,
				"status":  release.StatusDeployed.String(),
				"version": strconv.Itoa(rls.Version),
			},
			Annotations: map[string]string{
				apis.LabelChartFirstDeployed: rls.Info.FirstDeployed.UTC().Format(time.RFC3339),
				apis.LabelChartLastDeployed:  rls.Info.LastDeployed.UTC().Format(time.RFC3339),
			},
		},
		Spec: v1beta1.ApplicationSpec{
			Descriptor: v1beta1.Descriptor{
				Type:        rls.Chart.Metadata.Type,
				Version:     rls.Chart.Metadata.AppVersion,
				Description: rls.Info.Description,
				Owners:      nil, // FIX
				Keywords:    rls.Chart.Metadata.Keywords,
				Links: []v1beta1.Link{
					{
						Description: "website",
						URL:         rls.Chart.Metadata.Home,
					},
				},
				Notes: rls.Info.Notes,
			},
			ComponentGroupKinds: nil,
			Selector:            nil,
			AddOwnerRef:         false, // TODO
			AssemblyPhase:       toAssemblyPhase(rls.Info.Status),
		},
	}
	if rls.Chart.Metadata.Icon != "" {
		var imgType string
		if resp, err := http.Get(rls.Chart.Metadata.Icon); err == nil {
			if mime, err := mimetype.DetectReader(resp.Body); err == nil {
				imgType = mime.String()
			}
			_ = resp.Body.Close()
		}
		obj.Spec.Descriptor.Icons = []v1beta1.ImageSpec{
			{
				Source: rls.Chart.Metadata.Icon,
				// TotalSize: "",
				Type: imgType,
			},
		}
	}
	for _, maintainer := range rls.Chart.Metadata.Maintainers {
		obj.Spec.Descriptor.Maintainers = append(obj.Spec.Descriptor.Maintainers, v1beta1.ContactData{
			Name:  maintainer.Name,
			URL:   maintainer.URL,
			Email: maintainer.Email,
		})
	}

	//if len(commonLabels) > 0 {
	//	obj.Spec.Selector = &metav1.LabelSelector{
	//		MatchLabels: commonLabels,
	//	}
	//}

	return obj
}

func toAssemblyPhase(status release.Status) v1beta1.ApplicationAssemblyPhase {
	switch status {
	case release.StatusUnknown, release.StatusUninstalling, release.StatusPendingInstall, release.StatusPendingUpgrade, release.StatusPendingRollback:
		return v1beta1.Pending
	case release.StatusDeployed, release.StatusUninstalled, release.StatusSuperseded:
		return v1beta1.Succeeded
	case release.StatusFailed:
		return v1beta1.Failed
	}
	panic(fmt.Sprintf("unknown status: %s", status.String()))
}

// decodeRelease decodes the bytes of data into a release
// type. Data must contain a base64 encoded gzipped string of a
// valid release, otherwise an error is returned.
func decodeReleaseFromApp(app *v1beta1.Application, di dynamic.Interface, cl discovery.CachedDiscoveryInterface) (*rspb.Release, error) {
	var rls rspb.Release

	rls.Name = app.Labels["name"]
	rls.Namespace = app.Namespace
	rls.Version, _ = strconv.Atoi(app.Labels["version"])

	chartURL, ok := app.Annotations[apis.LabelChartURL]
	if !ok {
		return nil, fmt.Errorf("missing %s annotation on application %s/%s", apis.LabelChartURL, app.Namespace, app.Name)
	}
	chartName, ok := app.Annotations[apis.LabelChartName]
	if !ok {
		return nil, fmt.Errorf("missing %s annotation on application %s/%s", apis.LabelChartName, app.Namespace, app.Name)
	}
	chartVersion, ok := app.Annotations[apis.LabelChartVersion]
	if !ok {
		return nil, fmt.Errorf("missing %s annotation on application %s/%s", apis.LabelChartVersion, app.Namespace, app.Name)
	}
	chrt, err := lib.DefaultRegistry.GetChart(chartURL, chartName, chartVersion)
	if err != nil {
		return nil, err
	}
	rls.Chart = chrt.Chart

	rls.Info = &release.Info{
		Description: app.Spec.Descriptor.Description,
		Status:      release.Status(app.Labels["status"]),
		Notes:       app.Spec.Descriptor.Notes,
	}
	rls.Info.FirstDeployed, _ = helmtime.Parse(time.RFC3339, app.Annotations[apis.LabelChartFirstDeployed])
	rls.Info.LastDeployed, _ = helmtime.Parse(time.RFC3339, app.Annotations[apis.LabelChartLastDeployed])

	rlm := api.ObjectMeta{
		Name:      rls.Name,
		Namespace: rls.Namespace,
	}
	mapper := restmapper.NewDeferredDiscoveryRESTMapper(cl)
	tpl, err := editor.EditorChartValueManifest(app, mapper, di, rlm, rls.Chart)
	if err != nil {
		return nil, err
	}

	rls.Manifest = string(tpl.Manifest)

	if rls.Chart == nil {
		rls.Chart = &chart.Chart{}
	}
	rls.Chart.Values = tpl.Values.Object
	rls.Config = map[string]interface{}{}

	return &rls, nil
}

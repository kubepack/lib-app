/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

const (
	ResourceKindPlatformGrafanaDashboards = "PlatformGrafanaDashboards"
	ResourcePlatformGrafanaDashboards     = "platformgrafanadashboards"
	ResourcePlatformGrafanaDashboardss    = "platformgrafanadashboardss"
)

// PlatformGrafanaDashboards defines the schema for PlatformGrafanaDashboards Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=platformgrafanadashboardss,singular=platformgrafanadashboards,categories={kubeops,appscode}
type PlatformGrafanaDashboards struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PlatformGrafanaDashboardsSpec `json:"spec,omitempty"`
}

// PlatformGrafanaDashboardsSpec is the schema for PlatformGrafanaDashboards values file
type PlatformGrafanaDashboardsSpec struct {
	NameOverride     string `json:"nameOverride,omitempty"`
	FullnameOverride string `json:"fullnameOverride,omitempty"`

	// +optional
	Enabled bool                  `json:"enabled"`
	Values  *runtime.RawExtension `json:"values,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PlatformGrafanaDashboardss is a list of PlatformGrafanaDashboardss
type PlatformGrafanaDashboardss struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PlatformGrafanaDashboards CRD objects
	Items []PlatformGrafanaDashboards `json:"items,omitempty"`
}

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
)

const (
	ResourceKindPlatformOpscenter = "PlatformOpscenter"
	ResourcePlatformOpscenter     = "platformopscenter"
	ResourcePlatformOpscenters    = "platformopscenters"
)

// PlatformOpscenter defines the schema for PlatformOpscenter Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=platformopscenters,singular=platformopscenter,categories={kubeops,appscode}
type PlatformOpscenter struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PlatformOpscentersSpec `json:"spec,omitempty"`
}

// PlatformOpscentersSpec is the schema for PlatformOpscenter values file
type PlatformOpscentersSpec struct {
	NameOverride     string `json:"nameOverride,omitempty"`
	FullnameOverride string `json:"fullnameOverride,omitempty"`

	//+optional
	PlatformGrafanaDashboard PlatformGrafanaDashboard `json:"platform-grafana-dashboards"`
	//+optional
	NatsAlerts NatsAlertsConfig `json:"nats-alerts"`
	//+optional
	OpenfgaAlerts OpenfgaAlertsConfig `json:"openfga-alerts"`
}

type PlatformGrafanaDashboard struct {
	Enabled bool `json:"enabled"`
}

// NatsAlertsConfig configuration for NATS alerts
type NatsAlertsConfig struct {
	Enabled bool `json:"enabled"`
}

// OpenfgaAlertsConfig configuration for OpenFGA alerts
type OpenfgaAlertsConfig struct {
	Enabled bool `json:"enabled"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PlatformOpscenters is a list of PlatformOpscenters
type PlatformOpscenters struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PlatformOpscenter CRD objects
	Items []PlatformOpscenter `json:"items,omitempty"`
}

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
	catgwapi "go.bytebuilders.dev/catalog/api/gateway/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ResourceKindServiceGatewayPresets = "ServiceGatewayPresets"
	ResourceServiceGatewayPresets     = "servicegatewaypresets"
	ResourceServiceGatewayPresetss    = "servicegatewaypresetss"
)

// ServiceGatewayPresets defines the schama for ServiceGatewayPresets chart.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=servicegatewaypresetss,singular=servicegatewaypresets,categories={kubeops,appscode}
type ServiceGatewayPresets struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceGatewayPresetsSpec `json:"spec,omitempty"`
}

type ServiceGatewayPresetsSpec struct {
	NameOverride               string               `json:"nameOverride"`
	FullnameOverride           string               `json:"fullnameOverride"`
	ClusterMetadata            StashClusterMetadata `json:"clusterMetadata"`
	SkipGatewayPreset          bool                 `json:"skipGatewayPreset"`
	catgwapi.GatewayConfigSpec `json:",inline,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceGatewayPresetsList is a list of ServiceGatewayPresetss
type ServiceGatewayPresetsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ServiceGatewayPresets CRD objects
	Items []ServiceGatewayPresets `json:"items,omitempty"`
}

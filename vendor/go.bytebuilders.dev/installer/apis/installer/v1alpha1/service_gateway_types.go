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
	ResourceKindServiceGateway = "ServiceGateway"
	ResourceServiceGateway     = "servicegateways"
	ResourceServiceGateways    = "servicegateways"
)

// ServiceGateway defines the schama for ServiceGateway chart.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=servicegateways,singular=servicegateways,categories={kubeops,appscode}
type ServiceGateway struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceGatewaySpec `json:"spec,omitempty"`
}

type ServiceGatewaySpec struct {
	NameOverride         string               `json:"nameOverride"`
	FullnameOverride     string               `json:"fullnameOverride"`
	ClusterMetadata      StashClusterMetadata `json:"clusterMetadata"`
	catgwapi.GatewaySpec `json:",inline,omitempty"`
}

// +kubebuilder:validation:Enum=ca
type ClusterTLSIssuerType string

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceGatewayList is a list of ServiceGateways
type ServiceGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ServiceGateway CRD objects
	Items []ServiceGateway `json:"items,omitempty"`
}

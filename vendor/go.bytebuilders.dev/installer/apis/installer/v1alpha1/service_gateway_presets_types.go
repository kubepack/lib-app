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
	dnsapi "kubeops.dev/external-dns-operator/apis/external/v1alpha1"
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
	NameOverride     string                 `json:"nameOverride"`
	FullnameOverride string                 `json:"fullnameOverride"`
	Infra            ServiceProviderInfra   `json:"infra"`
	GatewayDns       ServiceGatewayDns      `json:"gateway-dns"`
	Cluster          ServiceProviderCluster `json:"cluster"`
}

type ServiceProviderInfra struct {
	ClusterName string `json:"clusterName"`
	HostInfo    `json:",inline"`
	TLS         InfraTLS   `json:"tls"`
	DNS         GatewayDns `json:"dns"`
}

type ServiceGatewayDns struct {
	Enabled bool                    `json:"enabled"`
	Spec    *dnsapi.ExternalDNSSpec `json:"spec,omitempty"`
}

type ServiceProviderCluster struct {
	TLS ClusterTLS `json:"tls"`
}

type ClusterTLS struct {
	Issuer ClusterTLSIssuerType `json:"issuer"`
	CA     TLSData              `json:"ca"`
}

// +kubebuilder:validation:Enum=ca
type ClusterTLSIssuerType string

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceGatewayPresetsList is a list of ServiceGatewayPresetss
type ServiceGatewayPresetsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ServiceGatewayPresets CRD objects
	Items []ServiceGatewayPresets `json:"items,omitempty"`
}

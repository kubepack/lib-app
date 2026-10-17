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

	core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kmodules.xyz/resource-metadata/apis/shared"
)

const (
	ResourceKindServiceVault = "ServiceVault"
	ResourceServiceVault     = "servicevault"
	ResourceServiceVaults    = "servicevaults"
)

// ServiceVault defines the schama for ServiceVault Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=servicevaults,singular=servicevault,categories={kubeops,appscode}
type ServiceVault struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceVaultSpec `json:"spec,omitempty"`
}

type ServiceVaultSpec struct {
	NameOverride     string                        `json:"nameOverride"`
	FullnameOverride string                        `json:"fullnameOverride"`
	Global           ServiceVaultGlobalValues      `json:"global"`
	ClusterMetadata  StashClusterMetadata          `json:"clusterMetadata"`
	Infra            catgwapi.ServiceProviderInfra `json:"infra"`
	GatewayDns       catgwapi.ServiceGatewayDns    `json:"gateway-dns"`
	// +optional
	VaultServer VaultServerSpec `json:"vaultServer"`
	// +optional
	Distro shared.DistroSpec `json:"distro"`
}

type ServiceVaultGlobalValues struct {
	License           string   `json:"license"`
	LicenseSecretName string   `json:"licenseSecretName"`
	Registry          string   `json:"registry"`
	RegistryFQDN      string   `json:"registryFQDN"`
	ImagePullSecrets  []string `json:"imagePullSecrets"`
	// +optional
	Distro shared.DistroSpec `json:"distro"`
}

type VaultServerSpec struct {
	Name              string               `json:"name"`
	Version           string               `json:"version"`
	Replicas          int32                `json:"replicas"`
	IsolateTenants    bool                 `json:"isolateTenants"`
	ServiceType       core.ServiceType     `json:"serviceType"`
	TerminationPolicy string               `json:"terminationPolicy"`
	Persistence       VaultPersistenceSpec `json:"persistence"`
	Unsealer          VaultUnsealerSpec    `json:"unsealer"`
	TLS               VaultTLSSpec         `json:"tls"`
	Relay             VaultRelaySpec       `json:"relay"`
}

type VaultPersistenceSpec struct {
	StorageClassName string            `json:"storageClassName"`
	Size             resource.Quantity `json:"size"`
}

type VaultUnsealerSpec struct {
	SecretShares    int32 `json:"secretShares"`
	SecretThreshold int32 `json:"secretThreshold"`
}

type VaultTLSSpec struct {
	CASecretName string `json:"caSecretName"`
}

type VaultRelaySpec struct {
	ClusterSet        string `json:"clusterSet"`
	Namespace         string `json:"namespace"`
	BootstrapTokenTTL string `json:"bootstrapTokenTTL"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceVaultList is a list of ServiceVaults
type ServiceVaultList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ServiceVault CRD objects
	Items []ServiceVault `json:"items,omitempty"`
}

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
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kmodules.xyz/resource-metadata/apis/shared"
)

const (
	ResourceKindCatalogManager = "CatalogManager"
	ResourceCatalogManager     = "catalogmanager"
	ResourceCatalogManagers    = "catalogmanagers"
)

// CatalogManager defines the schama for CatalogManager operator installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=catalogmanagers,singular=catalogmanager,categories={kubeops,appscode}
type CatalogManager struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CatalogManagerSpec `json:"spec,omitempty"`
}

// CatalogManagerSpec is the schema for Identity Server values file
type CatalogManagerSpec struct {
	//+optional
	Proxies shared.RegistryProxies `json:"proxies"`
	//+optional
	NameOverride string `json:"nameOverride"`
	//+optional
	FullnameOverride string    `json:"fullnameOverride"`
	ReplicaCount     int       `json:"replicaCount"`
	RegistryFQDN     string    `json:"registryFQDN"`
	Image            Container `json:"image"`
	//+optional
	ImagePullSecrets []string           `json:"imagePullSecrets"`
	ImagePullPolicy  string             `json:"imagePullPolicy"`
	ServiceAccount   ServiceAccountSpec `json:"serviceAccount"`
	//+optional
	PodAnnotations map[string]string `json:"podAnnotations"`
	// PodSecurityContext holds pod-level security attributes and common container settings.
	// Optional: Defaults to empty.  See type description for default values of each field.
	// +optional
	PodSecurityContext *core.PodSecurityContext `json:"podSecurityContext"`
	//+optional
	NodeSelector map[string]string `json:"nodeSelector"`
	// If specified, the pod's tolerations.
	// +optional
	Tolerations []core.Toleration `json:"tolerations"`
	// If specified, the pod's scheduling constraints
	// +optional
	Affinity   *core.Affinity `json:"affinity"`
	Monitoring Monitoring     `json:"monitoring"`

	// +optional
	ServiceProviderServiceAccount ObjectReference `json:"serviceProviderServiceAccount"`
	// +optional
	Keda CatalogManagerKedaSpec `json:"keda"`
	// +optional
	Helmrepo ObjectReference            `json:"helmrepo"`
	Platform CatalogManagerPlatformSpec `json:"platform"`
	// +optional
	Distro shared.DistroSpec `json:"distro"`
}

type CatalogManagerGatewaySpec struct {
	ClassName     string          `json:"className"`
	Name          string          `json:"name"`
	PortRange     string          `json:"portRange"`
	NodeportRange string          `json:"nodeportRange"`
	TlsSecretRef  ObjectReference `json:"tlsSecretRef"`
}

type CatalogManagerKedaSpec struct {
	ProxyService ObjectReference `json:"proxyService"`
}

type CatalogManagerPlatformSpec struct {
	BaseURL string `json:"baseURL"`
	// +optional
	CaBundle string `json:"caBundle"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CatalogManagerList is a list of CatalogManagers
type CatalogManagerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CatalogManager CRD objects
	Items []CatalogManager `json:"items,omitempty"`
}

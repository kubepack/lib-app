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
	ResourceKindRustfs = "Rustfs"
	ResourceRustfs     = "rustfs"
	ResourceRustfss    = "rustfss"
)

// Rustfs defines the schama for Rustfs Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=rustfss,singular=rustfs,categories={kubeops,appscode}
type Rustfs struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RustfsSpec `json:"spec,omitempty"`
}

// RustfsSpec is the schema for Rustfs Operator values file
type RustfsSpec struct {
	ReplicaCount int `json:"replicaCount"`
	//+optional
	RegistryFQDN string         `json:"registryFQDN"`
	Image        ImageReference `json:"image"`
	//+optional
	ImagePullSecrets []string `json:"imagePullSecrets"`
	//+optional
	NameOverride string `json:"nameOverride"`
	//+optional
	FullnameOverride string `json:"fullnameOverride"`
	//+optional
	PodAnnotations map[string]string `json:"podAnnotations"`
	//+optional
	PodSecurityContext *core.PodSecurityContext `json:"podSecurityContext"`
	//+optional
	SecurityContext *core.SecurityContext `json:"securityContext"`
	Service         AceServiceSpec        `json:"service"`
	//+optional
	Resources core.ResourceRequirements `json:"resources"`
	//+optional
	NodeSelector map[string]string `json:"nodeSelector"`
	// If specified, the pod's tolerations.
	// +optional
	Tolerations []core.Toleration `json:"tolerations"`
	// If specified, the pod's scheduling constraints
	// +optional
	Affinity     *core.Affinity       `json:"affinity"`
	Persistence  PersistenceSpec      `json:"persistence"`
	StorageClass LocalObjectReference `json:"storageClass"`
	Ingress      RustfsIngress        `json:"ingress"`
	Gateway      AppGateway           `json:"gateway"`
	Rustfs       RustfsConfig         `json:"rustfs"`
	// +optional
	Distro shared.DistroSpec `json:"distro"`
}

type RustfsIngress struct {
	Enabled     bool              `json:"enabled"`
	ClassName   string            `json:"className"`
	Annotations map[string]string `json:"annotations"`
	Domain      string            `json:"domain"`
}

type RustfsConfig struct {
	Auth RustfsAuth `json:"auth"`
	TLS  RustfsTLS  `json:"tls"`
}

type RustfsAuth struct {
	AccessKeyId     string `json:"accessKeyId"`
	SecretAccessKey string `json:"secretAccessKey"`
}

type RustfsTLS struct {
	Enable bool                 `json:"enable"`
	Mount  bool                 `json:"mount"`
	Issuer CertificateIssuerRef `json:"issuer"`
	Secret LocalObjectReference `json:"secret"`
}

type CertificateIssuerRef struct {
	Name string `json:"name"`
	Kind string `json:"kind"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RustfsList is a list of Rustfss
type RustfsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Rustfs CRD objects
	Items []Rustfs `json:"items,omitempty"`
}

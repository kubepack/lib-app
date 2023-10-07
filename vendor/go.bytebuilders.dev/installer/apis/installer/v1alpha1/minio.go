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
)

const (
	ResourceKindMinio = "Minio"
	ResourceMinio     = "minio"
	ResourceMinios    = "minios"
)

// Minio defines the schama for Minio Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=minios,singular=minio,categories={kubeops,appscode}
type Minio struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MinioSpec `json:"spec,omitempty"`
}

// MinioSpec is the schema for Minio Operator values file
type MinioSpec struct {
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
	Ingress      MinioIngress         `json:"ingress"`
	Minio        MinioConfig          `json:"minio"`
}

type MinioIngress struct {
	Enabled     bool              `json:"enabled"`
	ClassName   string            `json:"className"`
	Annotations map[string]string `json:"annotations"`
	Domain      string            `json:"domain"`
}

type MinioConfig struct {
	Auth MinioAuth `json:"auth"`
	TLS  MinioTLS  `json:"tls"`
}

type MinioAuth struct {
	AccessKeyId     string `json:"accessKeyId"`
	SecretAccessKey string `json:"secretAccessKey"`
}

type MinioTLS struct {
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

// MinioList is a list of Minios
type MinioList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Minio CRD objects
	Items []Minio `json:"items,omitempty"`
}

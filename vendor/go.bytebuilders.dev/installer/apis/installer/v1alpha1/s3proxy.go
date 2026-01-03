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
	ResourceKindS3proxy = "S3proxy"
	ResourceS3proxy     = "s3proxy"
	ResourceS3proxys    = "s3proxys"
)

// S3proxy defines the schama for S3proxy Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=s3proxys,singular=s3proxy,categories={kubeops,appscode}
type S3proxy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              S3proxySpec `json:"spec,omitempty"`
}

// S3proxySpec is the schema for S3proxy Operator values file
type S3proxySpec struct {
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
	Ingress      S3proxyIngress       `json:"ingress"`
	S3proxy      S3proxyConfig        `json:"s3proxy"`
	// +optional
	Distro shared.DistroSpec `json:"distro"`
}

type S3proxyIngress struct {
	Enabled     bool              `json:"enabled"`
	ClassName   string            `json:"className"`
	Annotations map[string]string `json:"annotations"`
	Domain      string            `json:"domain"`
}

type S3proxyConfig struct {
	Auth S3proxyAuth `json:"auth"`
	TLS  S3proxyTLS  `json:"tls"`
}

type S3proxyAuth struct {
	AccessKeyId     string `json:"accessKeyId"`
	SecretAccessKey string `json:"secretAccessKey"`
}

type S3proxyTLS struct {
	Enable               bool                 `json:"enable"`
	Issuer               CertificateIssuerRef `json:"issuer"`
	Secret               LocalObjectReference `json:"secret"`
	JksPasswordSecretRef ConfigKeySelector    `json:"jksPasswordSecretRef"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// S3proxyList is a list of S3proxys
type S3proxyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of S3proxy CRD objects
	Items []S3proxy `json:"items,omitempty"`
}

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
	ResourceKindServiceBackend = "ServiceBackend"
	ResourceServiceBackend     = "servicebackend"
	ResourceServiceBackends    = "servicebackends"
)

// ServiceBackend defines the schama for ServiceBackend Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=servicebackends,singular=servicebackend,categories={kubeops,appscode}
type ServiceBackend struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceBackendSpec `json:"spec,omitempty"`
}

// ServiceBackendSpec is the schema for ServiceBackend Operator values file
type ServiceBackendSpec struct {
	ReplicaCount int `json:"replicaCount"`
	//+optional
	RegistryFQDN string         `json:"registryFQDN"`
	Image        ImageReference `json:"image"`
	//+optional
	ImagePullSecrets []string `json:"imagePullSecrets"`
	//+optional
	NameOverride string `json:"nameOverride"`
	//+optional
	FullnameOverride string               `json:"fullnameOverride"`
	ServiceAccount   LocalObjectReference `json:"serviceAccount"`
	//+optional
	PodAnnotations map[string]string `json:"podAnnotations"`
	//+optional
	PodSecurityContext *core.PodSecurityContext `json:"podSecurityContext"`
	//+optional
	SecurityContext *core.SecurityContext `json:"securityContext"`
	Service         AceServiceSpec        `json:"service"`
	//+optional
	Resources   core.ResourceRequirements `json:"resources"`
	Autoscaling AutoscalingSpec           `json:"autoscaling"`
	//+optional
	NodeSelector map[string]string `json:"nodeSelector"`
	// If specified, the pod's tolerations.
	// +optional
	Tolerations []core.Toleration `json:"tolerations"`
	// If specified, the pod's scheduling constraints
	// +optional
	Affinity     *core.Affinity     `json:"affinity"`
	Volumes      []core.Volume      `json:"volumes"`
	VolumeMounts []core.VolumeMount `json:"volumeMounts"`
	Ingress      PlatformIngress    `json:"ingress"`
	Monitoring   Monitoring         `json:"monitoring"`
	Server       ServerConfig       `json:"server"`
}

type PlatformIngress struct {
	AppIngress `json:",inline,omitempty"`
	// +optional
	DNS AppIngressDns `json:"dns"`
}

type ServerConfig struct {
	OIDC            OIDC   `json:"oidc"`
	NamespacePrefix string `json:"namespacePrefix"`
	ConsumerScope   string `json:"consumerScope"`
	// External           External `json:"external"`
	Cookie Cookie `json:"cookie"`
}

type OIDC struct {
	ClientID     string `json:"clientID"`
	ClientSecret string `json:"clientSecret"`
	BaseURL      string `json:"baseURL"`
	IssuerURL    string `json:"issuerURL"`
	CallbackURL  string `json:"callbackURL"`
}

type Cookie struct {
	SigningKey    string `json:"signingKey"`
	EncryptionKey string `json:"encryptionKey"`
}

type AppIngressDns struct {
	// +optional
	TargetIPs []string `json:"targetIPs"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceBackendList is a list of ServiceBackends
type ServiceBackendList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ServiceBackend CRD objects
	Items []ServiceBackend `json:"items,omitempty"`
}

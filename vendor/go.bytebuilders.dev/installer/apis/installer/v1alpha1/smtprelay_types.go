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
	ResourceKindSmtprelay = "Smtprelay"
	ResourceSmtprelay     = "smtprelay"
	ResourceSmtprelays    = "smtprelays"
)

// Smtprelay defines the schama for Smtprelay Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=smtprelays,singular=smtprelay,categories={kubeops,appscode}
type Smtprelay struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SmtprelaySpec `json:"spec,omitempty"`
}

// SmtprelaySpec is the schema for Smtprelay Operator values file
type SmtprelaySpec struct {
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
	Controller         SmtpController           `json:"controller"`
	//+optional
	HostPort *IngressNginxControllerHostPort `json:"hostPort"`
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
	Affinity   *core.Affinity `json:"affinity"`
	Monitoring Monitoring     `json:"monitoring"`
	Ingress    AppIngress     `json:"ingress"`
	LogLevel   string         `json:"logLevel"`
	Smtp       SMTPConfig     `json:"smtp"`
}

type SMTPConfig struct {
	Addresses []string `json:"addresses"`
	Remotes   []string `json:"remotes"`
	TLS       SmtpTLS  `json:"tls"`
	Auth      SmtpAuth `json:"auth"`
}

type SmtpTLS struct {
	Enable bool                 `json:"enable"`
	Issuer CertificateIssuerRef `json:"issuer"`
	Secret LocalObjectReference `json:"secret"`
}

type SmtpAuth struct {
	// +optional
	Domain string `json:"domain"`
}

// +kubebuilder:validation:Enum=Deployment;DaemonSet
type AppController string

const (
	Deployment AppController = "Deployment"
	DaemonSet  AppController = "DaemonSet"
)

type SmtpController struct {
	Kind AppController `json:"kind"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SmtprelayList is a list of Smtprelays
type SmtprelayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Smtprelay CRD objects
	Items []Smtprelay `json:"items,omitempty"`
}

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
	ResourceKindGhCiWebhook = "GhCiWebhook"
	ResourceGhCiWebhook     = "ghciwebhook"
	ResourceGhCiWebhooks    = "ghciwebhooks"
)

// GhCiWebhook defines the schama for GhCiWebhook Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=dnsproxys,singular=dnsproxy,categories={kubeops,appscode}
type GhCiWebhook struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GhCiWebhookSpec `json:"spec,omitempty"`
}

// GhCiWebhookSpec is the schema for GhCiWebhook Operator values file
type GhCiWebhookSpec struct {
	ReplicaCount int                `json:"replicaCount"`
	Image        HelmImageReference `json:"image"`
	//+optional
	ImagePullSecrets []string `json:"imagePullSecrets"`
	//+optional
	NameOverride string `json:"nameOverride"`
	//+optional
	FullnameOverride string             `json:"fullnameOverride"`
	ServiceAccount   HelmServiceAccount `json:"serviceAccount"`
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
	Ingress      AppIngress         `json:"ingress"`
	Gateway      AppGateway         `json:"gateway"`
	Volumes      []core.Volume      `json:"volumes"`
	VolumeMounts []core.VolumeMount `json:"volumeMounts"`
	Args         []string           `json:"args"`
	EnvVars      map[string]string  `json:"envVars"`
	// +optional
	LivenessProbe *core.Probe `json:"livenessProbe"`
	// +optional
	ReadinessProbe *core.Probe       `json:"readinessProbe"`
	PodLabels      map[string]string `json:"podLabels"`
	// +optional
	Distro shared.DistroSpec `json:"distro"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GhCiWebhookList is a list of GhCiWebhooks
type GhCiWebhookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GhCiWebhook CRD objects
	Items []GhCiWebhook `json:"items,omitempty"`
}

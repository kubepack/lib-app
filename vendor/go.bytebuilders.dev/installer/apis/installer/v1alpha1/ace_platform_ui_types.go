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
	ResourceKindPlatformUi = "PlatformUi"
	ResourcePlatformUi     = "platformui"
	ResourcePlatformUis    = "platformuis"
)

// PlatformUi defines the schama for PlatformUi Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=platformuis,singular=platformui,categories={kubeops,appscode}
type PlatformUi struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PlatformUiSpec `json:"spec,omitempty"`
}

// PlatformUiSpec is the schema for PlatformUi Operator values file
type PlatformUiSpec struct {
	ReplicaCount int `json:"replicaCount"`
	//+optional
	RegistryFQDN string         `json:"registryFQDN"`
	Image        ImageReference `json:"image"`
	//+optional
	ImagePullSecrets []string `json:"imagePullSecrets"`
	//+optional
	NameOverride string `json:"nameOverride"`
	//+optional
	FullnameOverride string                `json:"fullnameOverride"`
	Platform         AcePlatformUISettings `json:"platform"`
	ServiceAccount   LocalObjectReference  `json:"serviceAccount"`
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
	Affinity *core.Affinity `json:"affinity"`
	Settings AceSettings    `json:"settings"`
}

type AcePlatformUISettings struct {
	Domain         string         `json:"domain"`
	DeploymentType DeploymentType `json:"deploymentType"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PlatformUiList is a list of PlatformUis
type PlatformUiList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PlatformUi CRD objects
	Items []PlatformUi `json:"items,omitempty"`
}

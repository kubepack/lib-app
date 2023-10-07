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
	ResourceKindBilling = "Billing"
	ResourceBilling     = "billing"
	ResourceBillings    = "billings"
)

// Billing defines the schama for Billing Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=billings,singular=billing,categories={kubeops,appscode}
type Billing struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BillingSpec `json:"spec,omitempty"`
}

// BillingSpec is the schema for Billing Operator values file
type BillingSpec struct {
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
	Resources core.ResourceRequirements `json:"resources"`
	//+optional
	NodeSelector map[string]string `json:"nodeSelector"`
	// If specified, the pod's tolerations.
	// +optional
	Tolerations []core.Toleration `json:"tolerations"`
	// If specified, the pod's scheduling constraints
	// +optional
	Affinity    *core.Affinity  `json:"affinity"`
	Persistence PersistenceSpec `json:"persistence"`
	Monitoring  Monitoring      `json:"monitoring"`
	Infra       AceInfra        `json:"infra"`
	Settings    AceSettings     `json:"settings"`
	Summary     BillingSummary  `json:"summary"`
}

type BillingSummary struct {
	Schedule string `json:"schedule"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BillingList is a list of Billings
type BillingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Billing CRD objects
	Items []Billing `json:"items,omitempty"`
}

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kmodules.xyz/resource-metadata/apis/shared"
)

const (
	ResourceKindAceOcmAddons = "AceOcmAddons"
	ResourceAceOcmAddons     = "aceocmaddons"
	ResourceAceOcmAddonss    = "aceocmaddonss"
)

// AceOcmAddons defines the schama for AceOcmAddons Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=aceocmaddonss,singular=aceocmaddons,categories={kubeops,appscode}
type AceOcmAddons struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AceOcmAddonsSpec `json:"spec,omitempty"`
}

// AceOcmAddonsSpec is the schema for AceOcmAddons Operator values file
type AceOcmAddonsSpec struct {
	NameOverride     string `json:"nameOverride"`
	FullnameOverride string `json:"fullnameOverride"`

	shared.BootstrapPresets `json:",inline,omitempty"`

	KubeconfigSecretName  string        `json:"kubeconfigSecretName"`
	AddonManagerNamespace string        `json:"addonManagerNamespace"`
	Placement             PlacementSpec `json:"placement"`
	Kubectl               DockerImage   `json:"kubectl"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AceOcmAddonsList is a list of AceOcmAddonss
type AceOcmAddonsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AceOcmAddons CRD objects
	Items []AceOcmAddons `json:"items,omitempty"`
}

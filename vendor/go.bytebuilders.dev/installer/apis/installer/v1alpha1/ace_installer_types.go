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
	configapi "go.bytebuilders.dev/resource-model/apis/config/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kmodules.xyz/resource-metadata/apis/shared"
)

const (
	ResourceKindAceInstaller = "AceInstaller"
	ResourceAceInstaller     = "aceinstaller"
	ResourceAceInstallers    = "aceinstallers"
)

// AceInstaller defines the schama for AceInstaller Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=aceinstallers,singular=aceinstaller,categories={kubeops,appscode}
type AceInstaller struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AceInstallerSpec `json:"spec,omitempty"`
}

// AceInstallerSpec is the schema for AceInstaller Operator values file
type AceInstallerSpec struct {
	NameOverride     string `json:"nameOverride"`
	FullnameOverride string `json:"fullnameOverride"`

	// +optional
	DeploymentType          DeploymentType `json:"deploymentType"`
	shared.BootstrapPresets `json:",inline,omitempty"`
	SelfManagement          configapi.SelfManagement `json:"selfManagement"`
	Precheck                AceInstallerPrecheckSpec `json:"precheck"`
}

type AceInstallerPrecheckSpec struct {
	Enabled            bool                      `json:"enabled"`
	Image              ImageReference            `json:"image"`
	PodAnnotations     map[string]string         `json:"podAnnotations"`
	PodSecurityContext *core.PodSecurityContext  `json:"podSecurityContext"`
	SecurityContext    *core.SecurityContext     `json:"securityContext"`
	Resources          core.ResourceRequirements `json:"resources"`
	//+optional
	NodeSelector map[string]string `json:"nodeSelector"`
	Tolerations  []core.Toleration `json:"tolerations"`
	Affinity     *core.Affinity    `json:"affinity"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AceInstallerList is a list of AceInstallers
type AceInstallerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AceInstaller CRD objects
	Items []AceInstaller `json:"items,omitempty"`
}

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

const (
	ResourceKindAkpInstaller = "AkpInstaller"
	ResourceAkpInstaller     = "akpinstaller"
	ResourceAkpInstallers    = "akpinstallers"
)

// AkpInstaller defines the AKP installer chart values.
// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:resource:path=akpinstallers,singular=akpinstaller,categories={kubeops,appscode}
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
type AkpInstaller struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AceInstallerSpec      `json:"spec,omitempty"`
	Status            *runtime.RawExtension `json:"status,omitempty"`
}

// AkpInstallerList is a list of AKP installers.
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type AkpInstallerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AkpInstaller `json:"items,omitempty"`
}

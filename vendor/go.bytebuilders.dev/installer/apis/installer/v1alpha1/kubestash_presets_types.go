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
	stashv1alpha1 "stash.appscode.dev/apimachinery/apis/stash/v1alpha1"
)

const (
	ResourceKindKubestashPresets = "KubestashPresets"
	ResourceKubestashPresets     = "kubestashpresets"
	ResourceKubestashPresetss    = "kubestashpresetss"
)

// KubestashPresets defines the schama for KubestashPresets Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=kubestashpresetss,singular=kubestashpresets,categories={kubeops,appscode}
type KubestashPresets struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubestashPresetsSpec `json:"spec,omitempty"`
}

type KubestashPresetsSpec struct {
	Kubestash KubeStashInfo `json:"kubestash"`
}

type KubeStashInfo struct {
	// Schedule specifies the schedule for invoking backup sessions
	// +optional
	Schedule string `json:"schedule,omitempty"`
	// RetentionPolicy indicates the policy to follow to clean old backup snapshots
	RetentionPolicy stashv1alpha1.RetentionPolicy `json:"retentionPolicy"`
	AuthSecret      AuthSecret                    `json:"authSecret"`
	Backend         RepositoryBackend             `json:"backend"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubestashPresetsList is a list of KubestashPresetss
type KubestashPresetsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KubestashPresets CRD objects
	Items []KubestashPresets `json:"items,omitempty"`
}

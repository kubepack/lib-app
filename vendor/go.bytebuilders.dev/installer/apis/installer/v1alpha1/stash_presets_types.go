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
	wizardsapi "go.bytebuilders.dev/ui-wizards/apis/wizards/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ResourceKindStashPresets = "StashPresets"
	ResourceStashPresets     = "stashpresets"
	ResourceStashPresetss    = "stashpresetss"
)

// StashPresets defines the schama for StashPresets Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=stashpresetss,singular=stashpresets,categories={kubeops,appscode}
type StashPresets struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StashPresetsSpec `json:"spec,omitempty"`
}

type StashPresetsSpec struct {
	// +kubebuilder:default=KubeStash
	Tool            BackupTool            `json:"tool"`
	ClusterMetadata *StashClusterMetadata `json:"clusterMetadata,omitempty"`
	KubeStash       KubeStashInfo         `json:"kubestash"`
}

// +kubebuilder:validation:Enum=KubeStash;Stash
type BackupTool string

const (
	BackupToolKubeStash BackupTool = "KubeStash"
	BackupToolStash     BackupTool = "Stash"
)

type StashClusterMetadata struct {
	Uid  string `json:"uid"`
	Name string `json:"name"`
}

type KubeStashInfo struct {
	// Schedule specifies the schedule for invoking backup sessions
	Schedule string `json:"schedule,omitempty"`
	// RetentionPolicy indicates the policy to follow to clean old backup snapshots
	// +kubebuilder:default=keep-1mo
	RetentionPolicy  KubeStashRetentionPolicy    `json:"retentionPolicy"`
	EncryptionSecret string                      `json:"encryptionSecret"`
	StorageSecret    wizardsapi.OptionalResource `json:"storageSecret"`
	Backend          wizardsapi.KubeStashBackend `json:"backend"`
}

// +kubebuilder:validation:Enum=keep-1wk;keep-1mo;keep-3mo;keep-1yr
type KubeStashRetentionPolicy string

const (
	KubeStashKeep1W KubeStashRetentionPolicy = "keep-1wk"
	KubeStashKeep1M KubeStashRetentionPolicy = "keep-1mo"
	KubeStashKeep3M KubeStashRetentionPolicy = "keep-3mo"
	KubeStashKeep1Y KubeStashRetentionPolicy = "keep-1yr"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// StashPresetsList is a list of StashPresetss
type StashPresetsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StashPresets CRD objects
	Items []StashPresets `json:"items,omitempty"`
}

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
	ResourceKindRegcache = "Regcache"
	ResourceRegcache     = "regcache"
	ResourceRegcaches    = "regcaches"
)

// Regcache defines the schema for the regcache chart values.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=regcaches,singular=regcache,categories={kubeops,appscode}
type Regcache struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RegcacheSpec `json:"spec,omitempty"`
}

// RegcacheSpec is the schema for the regcache chart values. It deploys an
// in-cluster pull-through Docker registry cache (CNCF Distribution / registry:3)
// that transparently proxies images from an upstream registry (ghcr.io by
// default), so that ghcr.io/appscode-charts images can be served from within
// the cluster.
type RegcacheSpec struct {
	Enabled      bool `json:"enabled"`
	ReplicaCount int  `json:"replicaCount"`
	//+optional
	NameOverride string `json:"nameOverride"`
	//+optional
	FullnameOverride string `json:"fullnameOverride"`
	// RegistryFQDN is the docker registry fqdn used to pull the registry image.
	RegistryFQDN string         `json:"registryFQDN"`
	Image        ImageReference `json:"image"`
	//+optional
	ImagePullSecrets []string `json:"imagePullSecrets"`
	// RemoteURL is the upstream registry that this instance proxies as a
	// pull-through cache.
	RemoteURL string `json:"remoteURL"`
	// TTL is how long a cached blob/manifest is kept before the pull-through
	// cache scheduler purges it. This is the only mechanism that frees disk:
	// the cache has no capacity-based (LRU) eviction, so lower values bound
	// disk usage at the cost of re-fetching from upstream more often.
	//+optional
	TTL string `json:"ttl"`
	// Username is an optional credential for authenticating against the
	// upstream registry. Leave empty for anonymous pulls.
	//+optional
	Username string `json:"username"`
	// Password is an optional credential for authenticating against the
	// upstream registry. Leave empty for anonymous pulls.
	//+optional
	Password string `json:"password"`
	//+optional
	PodAnnotations map[string]string `json:"podAnnotations"`
	//+optional
	PodSecurityContext *core.PodSecurityContext `json:"podSecurityContext"`
	//+optional
	SecurityContext *core.SecurityContext `json:"securityContext"`
	Service         AceServiceSpec        `json:"service"`
	Persistence     RegcachePersistence   `json:"persistence"`
	//+optional
	Resources core.ResourceRequirements `json:"resources"`
	//+optional
	NodeSelector map[string]string `json:"nodeSelector"`
	// If specified, the pod's tolerations.
	//+optional
	Tolerations []core.Toleration `json:"tolerations"`
	// If specified, the pod's scheduling constraints
	//+optional
	Affinity *core.Affinity `json:"affinity"`
	Infra    RegcacheInfra  `json:"infra"`
	//+optional
	Distro shared.DistroSpec `json:"distro"`
}

type RegcacheInfra struct {
	StorageClass LocalObjectReference `json:"storageClass"`
}

type RegcachePersistence struct {
	Enabled      bool   `json:"enabled"`
	StorageClass string `json:"storageClass"`
	AccessMode   string `json:"accessMode"`
	Size         string `json:"size"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RegcacheList is a list of Regcaches
type RegcacheList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Regcache CRD objects
	Items []Regcache `json:"items,omitempty"`
}

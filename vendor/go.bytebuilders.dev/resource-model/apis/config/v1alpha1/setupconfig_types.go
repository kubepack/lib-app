/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	cloudv1alpha1 "go.bytebuilders.dev/resource-model/apis/cloud/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
)

// AceSetupConfig is the Schema for the kubestashconfigs API

// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type AceSetupConfig struct {
	metav1.TypeMeta `json:",inline"`

	DeploymentType         string     `json:"deploymentType,omitempty"`
	Nats                   NatsConfig `json:"nats"`
	ImporterServiceAccount string     `json:"importerServiceAccount,omitempty"`
	AceSetupInlineConfig   `json:",inline"`
}

// NatsConfig holds the NATS-related fields.
type NatsConfig struct {
	Exports              bool `json:"exports"`
	ReloadNatsAccounts   bool `json:"reloadNatsAccounts"`
	CreateNatsStream     bool `json:"createNatsStream,omitempty"`
	RefactorNatsAccounts bool `json:"refactorNatsAccounts,omitempty"`
	Migrate              bool `json:"migrate,omitempty"`
}

type AceSetupInlineConfig struct {
	// +optional
	Admin AcePlatformAdmin `json:"admin"`
	// +optional
	SelfManagement SelfManagement `json:"selfManagement"`
	// +optional
	CloudCredential *cloudv1alpha1.Credential `json:"cloudCredential,omitempty"`
	// +optional
	Cluster *CAPIClusterConfig `json:"cluster,omitempty"`
	// +optional
	Subscription *MarketplaceSubscriptionInfo `json:"subscription,omitempty"`
}

type AceSetupInlineOptions struct {
	// +optional
	Admin AcePlatformAdmin `json:"admin"`
	// +optional
	SelfManagement SelfManagementOptions `json:"selfManagement"`
	// +optional
	CloudCredential *cloudv1alpha1.Credential `json:"cloudCredential,omitempty"`
	// +optional
	Cluster *CAPIClusterConfig `json:"cluster,omitempty"`
	// +optional
	Subscription *MarketplaceSubscriptionInfo `json:"subscription,omitempty"`
}

func (opt AceSetupInlineOptions) ToConfig() AceSetupInlineConfig {
	return AceSetupInlineConfig{
		Admin:           opt.Admin,
		SelfManagement:  opt.SelfManagement.ToConfig(),
		CloudCredential: opt.CloudCredential,
		Cluster:         opt.Cluster,
		Subscription:    opt.Subscription,
	}
}

type AcePlatformAdmin struct {
	// +optional
	Username string `json:"username"`
	// +optional
	Password string `json:"password"`
	// +optional
	Email string `json:"email"`
	// +optional
	DisplayName string `json:"displayName"`

	// +optional
	// Organization name should contain only alphanumeric, dash ('-'), underscore ('_') and dot ('.') characters.
	Orgname string `json:"orgname"`
}

type SelfManagement struct {
	// +optional
	Import bool `json:"import"`
	// +optional
	KubeAPIServer string `json:"kubeAPIServer,omitempty"`
	// +optional
	EnableFeatures []string `json:"enableFeatures"`
	// +optional
	DisableFeatures []string `json:"disableFeatures"`
}

type SelfManagementOptions struct {
	// +optional
	Import bool `json:"import"`
	// +optional
	EnableFeatures map[string][]string `json:"enableFeatures"`
	// +optional
	DisableFeatures []string `json:"disableFeatures"`
}

func (opt SelfManagementOptions) ToConfig() SelfManagement {
	enableFeatures := sets.Set[string]{}
	for _, features := range opt.EnableFeatures {
		enableFeatures.Insert(features...)
	}
	return SelfManagement{
		Import:          opt.Import,
		EnableFeatures:  sets.List(enableFeatures),
		DisableFeatures: sets.List(sets.New[string](opt.DisableFeatures...)),
	}
}

type CAPIClusterConfig struct {
	ClusterName       string        `json:"clusterName,omitempty"`
	Region            string        `json:"region,omitempty"`
	NetworkCIDR       string        `json:"networkCIDR,omitempty"`
	KubernetesVersion string        `json:"kubernetesVersion,omitempty"`
	GoogleProjectID   string        `json:"googleProjectID,omitempty"`
	WorkerPools       []MachinePool `json:"workerPools,omitempty"`
}

type MachinePool struct {
	MachineType  string `json:"machineType"`
	MachineCount int    `json:"machineCount"`
}

type MarketplaceSubscriptionInfo struct {
	AWS   *AWSMarSubscriptionInfo `json:"aws,omitempty"`
	Azure *AzureSubscriptionInfo  `json:"azure,omitempty"`
	GCP   *GCPSubscriptionInfo    `json:"gcp,omitempty"`
}

// https://docs.aws.amazon.com/marketplacemetering/latest/APIReference/API_MeterUsage.html
type AWSMarSubscriptionInfo struct {
	MeteringServiceProxyToken string `json:"meteringServiceProxyToken"`
}

// https://learn.microsoft.com/en-us/azure/azure-resource-manager/managed-applications/publish-notifications
type AzureSubscriptionInfo struct {
	ApplicationID string `json:"applicationId"`
}

// https://cloud.google.com/service-infrastructure/docs/service-control/reference/rest/v2/services/report
type GCPSubscriptionInfo struct {
	ServiceControlProxyToken string `json:"serviceControlProxyToken"`
}

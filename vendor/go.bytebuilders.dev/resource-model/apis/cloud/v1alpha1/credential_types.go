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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ResourceKindCredential = "Credential"
	ResourceCredential     = "credential"
	ResourceCredentials    = "credentials"
)

// +kubebuilder:validation:Enum=Aws;Azure;AzureStorage;DigitalOcean;GoogleCloud;GoogleOAuth;Linode;Packet;Rancher;Scaleway;Vultr;Swift
type CredentialType string

const (
	CredentialTypeAWS          CredentialType = "Aws"
	CredentialTypeAzure        CredentialType = "Azure"
	CredentialTypeAzureStorage CredentialType = "AzureStorage"
	CredentialTypeDigitalOcean CredentialType = "DigitalOcean"
	CredentialTypeGoogleCloud  CredentialType = "GoogleCloud"
	CredentialTypeGoogleOAuth  CredentialType = "GoogleOAuth"
	CredentialTypeLinode       CredentialType = "Linode"
	CredentialTypePacket       CredentialType = "Packet"
	CredentialTypeRancher      CredentialType = "Rancher"
	CredentialTypeScaleway     CredentialType = "Scaleway"
	CredentialTypeVultr        CredentialType = "Vultr"
	CredentialTypeSwift        CredentialType = "Swift"
	CredentialTypeHetzner      CredentialType = "Hetzner"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=credentials,singular=credential,shortName=cred,scope=Cluster,categories={kubernetes,resource-model,appscode}
// +kubebuilder:subresource:status
type Credential struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CredentialSpec   `json:"spec,omitempty"`
	Status            CredentialStatus `json:"status,omitempty"`
}

type CredentialSpec struct {
	Name    string         `json:"name"`
	Type    CredentialType `json:"type"`
	OwnerID int64          `json:"ownerID"`

	//+optional
	AWS *AWSCredential `json:"aws,omitempty"`
	//+optional
	Azure *AzureCredential `json:"azure,omitempty"`
	//+optional
	AzureStorage *AzureStorageCredential `json:"azureStorage,omitempty"`
	//+optional
	DigitalOcean *DigitalOceanCredential `json:"digitalocean,omitempty"`
	//+optional
	GoogleCloud *GoogleCloudCredential `json:"googleCloud,omitempty"`
	//+optional
	GoogleOAuth *GoogleOAuthCredential `json:"googleOAuth,omitempty"`
	//+optional
	Linode *LinodeCredential `json:"linode,omitempty"`
	//+optional
	Packet *PacketCredential `json:"packet,omitempty"`
	//+optional
	Rancher *RancherCredential `json:"rancher,omitempty"`
	//+optional
	Scaleway *ScalewayCredential `json:"scaleway,omitempty"`
	//+optional
	Swift *SwiftCredential `json:"swift,omitempty"`
	//+optional
	Vultr *VultrCredential `json:"vultr,omitempty"`
	//+optional
	Hetzner *HetznerCredential `json:"hetzner,omitempty"`
}

type GoogleOAuthCredential struct {
	AccessToken string `json:"accessToken"`
	// +optional
	RefreshToken string `json:"refreshToken,omitempty"`
	// +optional
	Scopes []string `json:"scopes,omitempty"`
	// +optional
	Expiry int64 `json:"expiry,omitempty"`
}

type GoogleCloudCredential struct {
	ProjectID      string `json:"projectID"`
	ServiceAccount string `json:"serviceAccount"`
}

type DigitalOceanCredential struct {
	Token string `json:"token"`
}

type AzureCredential struct {
	TenantID       string `json:"tenantID"`
	SubscriptionID string `json:"subscriptionID"`
	ClientID       string `json:"clientID"`
	ClientSecret   string `json:"clientSecret"`
}

type AzureStorageCredential struct {
	Account string `json:"account"`
	Key     string `json:"key"`
}

type AWSCredential struct {
	AccessKeyID     string `json:"accessKeyID"`
	SecretAccessKey string `json:"secretAccessKey"`
}

type PacketCredential struct {
	ProjectID string `json:"projectID"`
	APIKey    string `json:"apiKey"`
}

type RancherCredential struct {
	AccessKeyID     string `json:"accessKeyID"`
	SecretAccessKey string `json:"secretAccessKey"`
	Endpoint        string `json:"endpoint"`
}

type ScalewayCredential struct {
	Organization string `json:"organization"`
	Token        string `json:"token"`
}

type LinodeCredential struct {
	Token string `json:"token"`
}

type VultrCredential struct {
	Token string `json:"token"`
}

type SwiftCredential struct {
	Username      string `json:"username"`
	Password      string `json:"password"`
	TenantName    string `json:"tenantName,omitempty"`
	TenantAuthURL string `json:"tenantAuthURL,omitempty"`
	Domain        string `json:"domain,omitempty"`
	Region        string `json:"region,omitempty"`
	TenantId      string `json:"tenantID,omitempty"`
	TenantDomain  string `json:"tenantDomain,omitempty"`
}
type HetznerCredential struct {
	SSHKeyName string `json:"sshKeyName"`
	Token      string `json:"token"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
type CredentialList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Credential `json:"items,omitempty"`
}

type CredentialStatus struct {
	// ObservedGeneration is the most recent generation observed for this resource. It corresponds to the
	// resource's generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
}

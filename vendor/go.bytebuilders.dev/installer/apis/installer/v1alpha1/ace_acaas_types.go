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
)

const (
	ResourceKindAcaas = "Acaas"
	ResourceAcaas     = "acaas"
	ResourceAcaass    = "Acaass"
)

// Acaas defines the schama for ACE Hosting Components.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=aces,singular=ace,categories={kubeops,appscode}
type Acaas struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AcaasSpec `json:"spec,omitempty"`
}

// AcaasSpec is the schema for Ace Operator values file
type AcaasSpec struct {
	Billing        AceBilling        `json:"billing"`
	BillingUi      AceBillingUi      `json:"billing-ui"`
	DeployUi       AceDeployUi       `json:"deploy-ui"`
	MarketplaceApi AceMarketplaceApi `json:"marketplace-api"`
	MarketplaceUi  AceMarketplaceUi  `json:"marketplace-ui"`
	PlatformLinks  AcePlatformLinks  `json:"platform-links"`
	Website        AceWebsite        `json:"website"`

	Global  AcaasGlobalValues `json:"global"`
	Ingress AcaasIngress      `json:"ingress"`
}

type AceBilling struct {
	Enabled      bool `json:"enabled"`
	*BillingSpec `json:",inline,omitempty"`
}

type AceBillingUi struct {
	Enabled        bool `json:"enabled"`
	*BillingUiSpec `json:",inline,omitempty"`
}

type AceDeployUi struct {
	Enabled       bool `json:"enabled"`
	*DeployUiSpec `json:",inline,omitempty"`
}

type AceDnsProxy struct {
	Enabled       bool `json:"enabled"`
	*DnsProxySpec `json:",inline,omitempty"`
}

type AceMarketplaceApi struct {
	Enabled             bool `json:"enabled"`
	*MarketplaceApiSpec `json:",inline,omitempty"`
}

type AceMarketplaceUi struct {
	Enabled            bool `json:"enabled"`
	*MarketplaceUiSpec `json:",inline,omitempty"`
}

type AcePlatformLinks struct {
	Enabled            bool `json:"enabled"`
	*PlatformLinksSpec `json:",inline,omitempty"`
}

type AceSmtprelay struct {
	Enabled        bool `json:"enabled"`
	*SmtprelaySpec `json:",inline,omitempty"`
}

type AceWebsite struct {
	Enabled      bool `json:"enabled"`
	*WebsiteSpec `json:",inline,omitempty"`
}

type AcaasGlobalValues struct {
	NameOverride     string                `json:"nameOverride"`
	FullnameOverride string                `json:"fullnameOverride"`
	Platform         AcaasPlatformSettings `json:"platform"`
	Registry         string                `json:"registry"`
	RegistryFQDN     string                `json:"registryFQDN"`
	Settings         AcaasSettings         `json:"settings"`
}

type AcaasPlatformSettings struct {
	Host string `json:"host"`
}

type AcaasSettings struct {
	CAProviderClass                string `json:"caProviderClass"`
	AcaasSettingsSecretName        `json:"secretName"`
	SpreadsheetCredentialMountPath string `json:"spreadsheetCredentialMountPath"`
}

type AcaasSettingsSecretName struct {
	AceSettingsSecretName `json:",inline,omitempty"`
	Objstore              string `json:"objstore"`
	Spreadsheet           string `json:"spreadsheet"`
}

type AcaasIngress struct {
	ClassName string            `json:"className"`
	TLS       AcaasIngressTLS   `json:"tls"`
	Rules     AcaasIngressRules `json:"rules"`
}

type AcaasIngressTLS struct {
	Enable bool                 `json:"enable"`
	Secret LocalObjectReference `json:"secret"`
}

type AcaasIngressRules struct {
	Blog     ExternalService `json:"blog"`
	Docs     ExternalService `json:"docs"`
	Learn    ExternalService `json:"learn"`
	License  ExternalService `json:"license"`
	Selfhost ExternalService `json:"selfhost"`
}

type ExternalService struct {
	Upstream string `json:"upstream"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AcaasList is a list of Aces
type AcaasList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Ace CRD objects
	Items []Acaas `json:"items,omitempty"`
}

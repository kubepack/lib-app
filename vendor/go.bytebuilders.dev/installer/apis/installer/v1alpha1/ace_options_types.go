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
	"fmt"
	"net/url"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kmodules.xyz/resource-metadata/apis/shared"
)

const (
	ResourceKindAceOptions = "AceOptions"
	ResourceAceOptions     = "aceoptions"
	ResourceAceOptionss    = "aceoptionss"
)

// AceOptions defines the schama for AceOptions Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=aceoptionss,singular=aceoptions,categories={kubeops,appscode}
type AceOptions struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AceOptionsSpec `json:"spec,omitempty"`
}

// AceOptionsSpec is the schema for AceOptions Operator values file
type AceOptionsSpec struct {
	Context       AceDeploymentContext    `json:"context"`
	Release       ObjectReference         `json:"release"`
	Registry      RegistrySpec            `json:"registry"`
	Monitoring    GlobalMonitoring        `json:"monitoring"`
	Infra         AceOptionsPlatformInfra `json:"infra"`
	Settings      AceOptionsSettings      `json:"settings"`
	Billing       AceOptionsComponentSpec `json:"billing"`
	PlatformUi    AceOptionsComponentSpec `json:"platform-ui"`
	AccountsUi    AceOptionsComponentSpec `json:"accounts-ui"`
	ClusterUi     AceOptionsComponentSpec `json:"cluster-ui"`
	DeployUi      AceOptionsComponentSpec `json:"deploy-ui"`
	Grafana       AceOptionsComponentSpec `json:"grafana"`
	KubedbUi      AceOptionsComponentSpec `json:"kubedb-ui"`
	MarketplaceUi AceOptionsComponentSpec `json:"marketplace-ui"`
	PlatformApi   AceOptionsComponentSpec `json:"platform-api"`
	PlatformLinks AceOptionsComponentSpec `json:"platform-links"`
	Ingress       AceOptionsIngressNginx  `json:"ingress"`
	Nats          AceOptionsNatsSettings  `json:"nats"`
	Trickster     AceOptionsComponentSpec `json:"trickster"`
	DNSProxy      AceOptionsComponentSpec `json:"dns-proxy"`
	SMTPRelay     AceOptionsComponentSpec `json:"smtprelay"`
	Minio         AceOptionsComponentSpec `json:"minio"`
	Branding      AceBrandingSpec         `json:"branding"`
}

type RegistrySpec struct {
	//+optional
	Image shared.ImageRegistrySpec `json:"image"`
	//+optional
	Credentials RepositoryCredential `json:"credentials"`
	//+optional
	Helm HelmOptions `json:"helm"`
	//+optional
	AllowNondistributableArtifacts bool `json:"allowNondistributableArtifacts"`
	//+optional
	Insecure bool `json:"insecure"`
}

type RepositoryCredential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type HelmOptions struct {
	//+optional
	Repositories HelmRepositories `json:"repositories"`
}

type HelmRepositories struct {
	//+optional
	AppscodeChartsOci string `json:"appscode-charts-oci"`
}

type AceOptionsComponentSpec struct {
	Enabled bool `json:"enabled"`
	//+optional
	Resources core.ResourceRequirements `json:"resources"`
	// +optional
	NodeSelector map[string]string `json:"nodeSelector"`
}

// +kubebuilder:validation:Enum=LoadBalancer;HostPort
type ServiceType string

const (
	ServiceTypeLoadBalancer ServiceType = "LoadBalancer"
	ServiceTypeHostPort     ServiceType = "HostPort"
)

const (
	DefaultPasswordLength = 16
)

// ref: https://github.com/kubernetes-sigs/external-dns/blob/v0.13.1/pkg/apis/externaldns/types.go#L325
// +kubebuilder:validation:Enum=A;CNAME
type DNSRecordType string

const (
	DNSRecordTypeA     DNSRecordType = "A"
	DNSRecordTypeCNAME DNSRecordType = "CNAME"
)

type AceOptionsIngressNginx struct {
	ExposeVia ServiceType `json:"exposeVia"`
	// DNS record types that will be considered for management
	ManagedRecordTypes []DNSRecordType `json:"managedRecordTypes,omitempty"`
	//+optional
	Resources    core.ResourceRequirements `json:"resources"`
	NodeSelector map[string]string         `json:"nodeSelector"`
}

type AceOptionsNatsSettings struct {
	ExposeVia ServiceType `json:"exposeVia"`
	Replics   int         `json:"replicas"`
	//+optional
	Resources core.ResourceRequirements `json:"resources"`
	//+optional
	NodeSelector map[string]string `json:"nodeSelector"`
}

type AceOptionsPlatformInfra struct {
	StorageClass  LocalObjectReference         `json:"storageClass"`
	Stash         InfraStash                   `json:"stash"`
	TLS           AceOptionsInfraTLS           `json:"tls"`
	DNS           InfraDns                     `json:"dns"`
	CloudServices AceOptionsInfraCloudServices `json:"cloudServices"`
}

type AceOptionsInfraTLS struct {
	Issuer      TLSIssuerType `json:"issuer"`
	Acme        TLSIssuerAcme `json:"acme"`
	Certificate TLSData       `json:"certificate"`
}

type AceOptionsInfraCloudServices struct {
	Provider ObjstoreProvider        `json:"provider"`
	Objstore AceOptionsInfraObjstore `json:"objstore"`
	// +optional
	Kms *AceOptionsInfraKms `json:"kms,omitempty"`
}

func (cs *AceOptionsInfraCloudServices) ObjStoreURL() string {
	if cs.Provider == ObjstoreProviderS3 {
		ur, _ := url.Parse(cs.Objstore.Bucket)
		values := ur.Query()
		values.Set("region", cs.Objstore.Region)
		if cs.Objstore.Endpoint != "" {
			values.Set("endpoint", cs.Objstore.Endpoint)
		}
		ur.RawQuery = values.Encode()
		return ur.String()
	}
	return cs.Objstore.Bucket
}

func (cs *AceOptionsInfraCloudServices) GenerateExportEnv() string {
	switch cs.Provider {
	case ObjstoreProviderS3:
		return fmt.Sprintf(`export AWS_ACCESS_KEY_ID=%s
export AWS_SECRET_ACCESS_KEY=%s`, cs.Objstore.Auth.S3.AwsAccessKeyID, cs.Objstore.Auth.S3.AwsSecretAccessKey)
	case ObjstoreProviderGCS:
		return fmt.Sprintf(`echo %s > google_credentials.json
export GOOGLE_APPLICATION_CREDENTIALS=google_credentials.json`, cs.Objstore.Auth.GCS.GoogleServiceAccountJSONKey)
	case ObjstoreProviderAzure:
		return fmt.Sprintf(`export AZURE_STORAGE_ACCOUNT=%s
export AZURE_STORAGE_KEY=%s`, cs.Objstore.Auth.Azure.AzureAccountName, cs.Objstore.Auth.Azure.AzureAccountKey)
	default:
		return ""
	}
}

type AceOptionsInfraObjstore struct {
	Bucket string `json:"bucket"`
	Prefix string `json:"prefix,omitempty"`
	// Required for s3 type buckets other than AWS s3 buckets
	Endpoint string `json:"endpoint,omitempty"`
	// Required for s3 buckets
	Region string       `json:"region,omitempty"`
	Auth   ObjstoreAuth `json:"auth"`
}

type ObjstoreAuth struct {
	S3    *S3Auth    `json:"s3,omitempty"`
	Azure *AzureAuth `json:"azure,omitempty"`
	GCS   *GCSAuth   `json:"gcs,omitempty"`
	Swift *SwiftAuth `json:"swift,omitempty"`
}

type AceOptionsInfraKms struct {
	MasterKeyURL string `json:"masterKeyURL"`
}

type AceOptionsSettings struct {
	DB              AceOptionsDBSettings    `json:"db"`
	Cache           AceOptionsCacheSettings `json:"cache"`
	SMTP            AceOptionsSMTPSettings  `json:"smtp"`
	DomainWhiteList []string                `json:"domainWhiteList"`
}

type AceOptionsDBSettings struct {
	Persistence PersistenceSpec           `json:"persistence"`
	Resources   core.ResourceRequirements `json:"resources"`
}

type AceOptionsCacheSettings struct {
	Persistence PersistenceSpec           `json:"persistence"`
	Resources   core.ResourceRequirements `json:"resources"`
}

type AceOptionsSMTPSettings struct {
	Enabled    bool   `json:"enabled"`
	Host       string `json:"host"`
	TlsEnabled bool   `json:"tlsEnabled"`
	From       string `json:"from"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	// SubjectPrefix   string `json:"subjectPrefix"`
	SendAsPlainText bool `json:"sendAsPlainText"`
}

// +kubebuilder:validation:Enum=Hosted;SelfHostedProduction;SelfHostedDemo
type DeploymentType string

const (
	HostedDeployment               DeploymentType = "Hosted"
	SelfHostedProductionDeployment DeploymentType = "SelfHostedProduction"
	SelfHostedDemoDeployment       DeploymentType = "SelfHostedDemo"
)

func (dt DeploymentType) Hosted() bool {
	return dt == HostedDeployment
}

func (dt DeploymentType) SelfHosted() bool {
	return dt == SelfHostedProductionDeployment || dt == SelfHostedDemoDeployment
}

func (dt DeploymentType) Demo() bool {
	return dt == SelfHostedDemoDeployment
}

type AceDeploymentContext struct {
	DeploymentType       DeploymentType `json:"deploymentType"`
	RequestedDomain      string         `json:"requestedDomain"`
	HostedDomain         string         `json:"hostedDomain,omitempty"`
	RequesterDisplayName string         `json:"requesterDisplayName,omitempty"`
	RequesterUsername    string         `json:"requesterUsername,omitempty"`
	ProxyServiceDomain   string         `json:"proxyServiceDomain,omitempty"`
	Token                string         `json:"token,omitempty"`
	// +optional
	OfflineInstaller bool `json:"offlineInstaller"`
	// WARNING!!! Update docs in schema/ace-options/patch.yaml
	// +optional
	ClusterID string `json:"clusterID"`
	// +optional
	PublicIPs []string          `json:"publicIPs"`
	Licenses  map[string]string `json:"licenses,omitempty"`
	// +optional
	Admin AcePlatformAdmin `json:"admin"`

	PromotedToProduction bool             `json:"promotedToProduction,omitempty"`
	PromotionValues      *PromotionValues `json:"promotionValues,omitempty"`

	GeneratedValues `json:",inline,omitempty"`
}

type GeneratedValues struct {
	// +optional
	AdminPassword string `json:"adminPassword"`
	// +optional
	BackupPassword string `json:"backupPassword"`
	// +optional
	PostgresPassword string `json:"postgresPassword"`
	// +optional
	RedisPassword string `json:"redisPassword"`
	// +optional
	Oauth2JWTSecret string `json:"oauth2JWTSecret"`
	// +optional
	CsrfSecretKey string `json:"csrfSecretKey"`
	// +optional
	S3AccessKeyID string `json:"s3AccessKeyID"`
	// +optional
	S3AccessKeySecret string `json:"s3AccessKeySecret"`
	// +optional
	Nats map[string]string `json:"nats"`
}

type PromotionValues struct {
	Minio AceOptionsInfraObjstore `json:"minio,omitempty"`
}

type AcePlatformAdmin struct {
	Username    string `json:"username,omitempty"`
	Password    string `json:"password,omitempty"`
	Email       string `json:"email,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AceOptionsList is a list of AceOptionss
type AceOptionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AceOptions CRD objects
	Items []AceOptions `json:"items,omitempty"`
}

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

	catgwapi "go.bytebuilders.dev/catalog/api/gateway/v1alpha1"
	configapi "go.bytebuilders.dev/resource-model/apis/config/v1alpha1"
	wizardsapi "go.bytebuilders.dev/ui-wizards/apis/wizards/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	store "kmodules.xyz/objectstore-api/api/v1"
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
	Context      AceDeploymentContext            `json:"context"`
	Release      ObjectReference                 `json:"release"`
	Registry     RegistrySpec                    `json:"registry"`
	Monitoring   GlobalMonitoring                `json:"monitoring"`
	Infra        AceOptionsPlatformInfra         `json:"infra"`
	Settings     AceOptionsSettings              `json:"settings"`
	PlatformUi   AceOptionsComponentSpec         `json:"platform-ui"`
	ClusterUi    AceOptionsComponentSpec         `json:"cluster-ui"`
	Grafana      AceOptionsComponentSpec         `json:"grafana"`
	KubedbUi     AceOptionsComponentSpec         `json:"kubedb-ui"`
	PlatformApi  AceOptionsComponentSpec         `json:"platform-api"`
	Ingress      AceOptionsIngressNginx          `json:"ingress"`
	Nats         AceOptionsNatsSettings          `json:"nats"`
	Trickster    AceOptionsComponentSpec         `json:"trickster"`
	Openfga      AceOptionsComponentSpec         `json:"openfga"`
	S3proxy      AceOptionsComponentSpec         `json:"s3proxy"`
	Branding     AceBrandingSpec                 `json:"branding"`
	InitialSetup configapi.AceSetupInlineOptions `json:"initialSetup"`
}

func (a *AceOptionsSpec) IsOptionsComplete() bool {
	switch a.Context.DeploymentType {
	case AWSMarketplaceDeployment:
		return a.InitialSetup.Subscription != nil &&
			a.InitialSetup.Subscription.AWS != nil &&
			a.InitialSetup.Subscription.AWS.MeteringServiceProxyToken != ""
	case AzureMarketplaceDeployment:
		return a.InitialSetup.Subscription != nil &&
			a.InitialSetup.Subscription.Azure != nil &&
			a.InitialSetup.Subscription.Azure.ApplicationID != ""
	case GCPMarketplaceDeployment:
		return a.InitialSetup.Subscription != nil &&
			a.InitialSetup.Subscription.GCP != nil &&
			a.InitialSetup.Subscription.GCP.ServiceControlProxyToken != ""
	}
	return true
}

func (a *AceOptionsSpec) Host() string {
	if a.Infra.DNS.Provider == catgwapi.DNSProviderNone && a.IsOptionsComplete() {
		if len(a.Infra.DNS.TargetIPs) == 0 {
			panic("target IPs required when no dns provider is used")
		}
		return a.Infra.DNS.TargetIPs[0]
	}
	return a.Context.HostedDomain
}

func (a *AceOptionsSpec) HostType() catgwapi.HostType {
	if a.Infra.DNS.Provider == catgwapi.DNSProviderNone {
		return catgwapi.HostTypeIP
	}
	return catgwapi.HostTypeDomain
}

type RegistrySpec struct {
	//+optional
	Image shared.ImageRegistrySpec `json:"image"`
	//+optional
	Credentials RepositoryCredential `json:"credentials"`
	//+optional
	Certs RepositoryCertificates `json:"certs"`
	//+optional
	Helm HelmOptions `json:"helm"`
	//+optional
	AllowNondistributableArtifacts bool `json:"allowNondistributableArtifacts"`
	//+optional
	Insecure bool `json:"insecure"`
	//+optional
	ImagePullSecrets []string `json:"imagePullSecrets"`
}

type RepositoryCredential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RepositoryCertificates struct {
	CACert     string `json:"caCert,omitempty"`
	ClientCert string `json:"clientCert,omitempty"`
	ClientKey  string `json:"clientKey,omitempty"`
}

type HelmOptions struct {
	// +optional
	CreateNamespace bool `json:"createNamespace,omitempty"`
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

// +kubebuilder:validation:Enum=LoadBalancer;ClusterIP;HostPort
type ServiceType string

const (
	ServiceTypeLoadBalancer ServiceType = "LoadBalancer"
	ServiceTypeHostPort     ServiceType = "HostPort"
	ServiceTypeClusterIP    ServiceType = "ClusterIP"
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
	// +optional
	ExternalIPs []string `json:"externalIPs"`
}

// +kubebuilder:validation:Enum=Ingress;HostPort
type ExposeNatsVia string

const (
	ExposeNatsViaIngress  ExposeNatsVia = "Ingress"
	ExposeNatsViaHostPort ExposeNatsVia = "HostPort"
)

type AceOptionsNatsSettings struct {
	ExposeVia ExposeNatsVia `json:"exposeVia"`
	Replics   int           `json:"replicas"`
	//+optional
	Resources core.ResourceRequirements `json:"resources"`
	//+optional
	NodeSelector map[string]string `json:"nodeSelector"`
}

type AceOptionsPlatformInfra struct {
	StorageClass  LocalObjectReference         `json:"storageClass"`
	KubeStash     KubeStashOptions             `json:"kubestash"`
	TLS           catgwapi.InfraTLS            `json:"tls"`
	DNS           InfraDns                     `json:"dns"`
	CloudServices AceOptionsInfraCloudServices `json:"cloudServices"`
}

type KubeStashOptions struct {
	// Schedule specifies the schedule for invoking backup sessions
	// +optional
	Schedule string `json:"schedule,omitempty"`
	// RetentionPolicy indicates the policy to follow to clean old backup snapshots
	// +kubebuilder:default=keep-1mo
	RetentionPolicy KubeStashRetentionPolicy    `json:"retentionPolicy"`
	StorageSecret   wizardsapi.OptionalResource `json:"storageSecret"`
	Backend         KubeStashBackendInfra       `json:"backend"`
}

type KubeStashBackendInfra struct {
	Provider string `json:"provider"`
	// +optional
	S3 store.S3Spec `json:"s3"`
	// +optional
	Azure store.AzureSpec `json:"azure"`
	// +optional
	GCS store.GCSSpec `json:"gcs"`
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
	S3    *wizardsapi.S3Auth    `json:"s3,omitempty"`
	Azure *wizardsapi.AzureAuth `json:"azure,omitempty"`
	GCS   *wizardsapi.GCSAuth   `json:"gcs,omitempty"`
}

type AceOptionsInfraKms struct {
	MasterKeyURL string `json:"masterKeyURL"`
}

type AceOptionsSettings struct {
	DB    AceOptionsDBSettings    `json:"db"`
	Cache AceOptionsCacheSettings `json:"cache"`
	SMTP  AceOptionsSMTPSettings  `json:"smtp"`

	// DomainWhiteList is an array of domain names that are allowed.
	// Each domain should be in the format of a fully qualified domain name,
	// such as 'example.com' or 'appscode.com' etc.
	// +optional
	DomainWhiteList []string `json:"domainWhiteList"`
	// +optional
	LogoutURL string `json:"logoutURL"`
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

// +kubebuilder:validation:Enum=Hosted;SelfHostedProduction;OpenShiftProduction;CloudDemo;OnpremDemo;KubeAppDemo;OpenShiftDemo;AWSMarketplace;AzureMarketplace;GoogleCloudMarketplace
type DeploymentType string

const (
	HostedDeployment               DeploymentType = "Hosted"
	SelfHostedProductionDeployment DeploymentType = "SelfHostedProduction"
	OpenShiftProductionDeployment  DeploymentType = "OpenShiftProduction"

	CloudDemoDeployment     DeploymentType = "CloudDemo"
	OnpremDemoDeployment    DeploymentType = "OnpremDemo"
	KubeAppDemoDeployment   DeploymentType = "KubeAppDemo"
	OpenShiftDemoDeployment DeploymentType = "OpenShiftDemo"

	AWSMarketplaceDeployment   DeploymentType = "AWSMarketplace"
	AzureMarketplaceDeployment DeploymentType = "AzureMarketplace"
	GCPMarketplaceDeployment   DeploymentType = "GoogleCloudMarketplace"
)

func (dt DeploymentType) Hosted() bool {
	return dt == HostedDeployment
}

func (dt DeploymentType) Demo() bool {
	return dt == CloudDemoDeployment ||
		dt == OnpremDemoDeployment ||
		dt == KubeAppDemoDeployment ||
		dt == OpenShiftDemoDeployment
}

func (dt DeploymentType) Onprem() bool {
	return dt == OnpremDemoDeployment
}

func (dt DeploymentType) OpenShift() bool {
	return dt == OpenShiftProductionDeployment ||
		dt == OpenShiftDemoDeployment
}

func (dt DeploymentType) MarketplaceDeployment() bool {
	return dt == AWSMarketplaceDeployment ||
		dt == AzureMarketplaceDeployment ||
		dt == GCPMarketplaceDeployment
}

func (dt DeploymentType) UsesVirtualCluster() bool {
	return dt == KubeAppDemoDeployment ||
		dt == GCPMarketplaceDeployment
}

type AceDeploymentContext struct {
	DeploymentType DeploymentType `json:"deploymentType"`
	InstallerName  string         `json:"installerName"`
	UploadID       string         `json:"uploadID"`
	Version        string         `json:"version"`
	// +optional
	RequestedDomain      string `json:"requestedDomain"`
	HostedDomain         string `json:"hostedDomain,omitempty"`
	RequesterDisplayName string `json:"requesterDisplayName,omitempty"`
	RequesterUsername    string `json:"requesterUsername,omitempty"`
	ProxyServiceDomain   string `json:"proxyServiceDomain,omitempty"`
	Token                string `json:"token,omitempty"`
	LicenseServiceDomain string `json:"licenseServiceDomain,omitempty"`
	LicenseServiceToken  string `json:"licenseServiceToken,omitempty"`
	LicenseOwnerID       int64  `json:"licenseOwnerID"`
	LicenseOwnerName     string `json:"licenseOwnerName"`

	// +optional
	OfflineInstaller bool `json:"offlineInstaller"`
	// WARNING!!! Update docs in schema/ace-options/patch.yaml
	// +optional
	ClusterID string            `json:"clusterID"`
	Licenses  map[string]string `json:"licenses,omitempty"`

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
	// +optional
	ServiceBackendCookie Cookie `json:"serviceBackendCookie"`
	// +optional
	ClusterCA catgwapi.TLSData `json:"clusterCA"`
	// JKS password is used to create keystore for s3proxy
	// +optional
	JKSPassword string `json:"jksPassword"`
	// +optional
	GrafanaSecretKey string            `json:"grafanaSecretKey"`
	InboxServer      InboxServerValues `json:"inboxServer"`
	// InstallerSecret used by hosted mode (prod and ninja)
	// to generate and validate marketplace self-hosted installer options
	// +optional
	InstallerSecret string `json:"installerSecret,omitempty"`
}

type InboxServerValues struct {
	AdminJWTPrivateKey string `json:"adminJWTPrivateKey"`
}

type PromotionValues struct {
	S3proxy AceOptionsInfraObjstore `json:"s3proxy,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AceOptionsList is a list of AceOptionss
type AceOptionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AceOptions CRD objects
	Items []AceOptions `json:"items,omitempty"`
}

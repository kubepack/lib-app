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
	"maps"
	"slices"
)

// Dotted values paths kept out of HelmRelease spec.values and passed via valuesFrom Secrets.
// The charts get these lists from cmd/gen-cred-paths (make gen-cred-paths); b3 and catalog-manager import them.

// ACECredPaths are relative to the ace chart values.
var ACECredPaths = []string{
	"global.platform.token",
	"global.infra.kms.masterKeyURL",
	"global.infra.tls.ca.key",
	"global.infra.tls.certificate.key",
	"global.infra.tls.jks.password",
	"global.infra.tls.jks.keystore",
	"global.infra.dns.auth.cloudflare.token",
	"global.infra.dns.auth.route53.AWS_ACCESS_KEY_ID",
	"global.infra.dns.auth.route53.AWS_SECRET_ACCESS_KEY",
	"global.infra.dns.auth.cloudDNS.GOOGLE_SERVICE_ACCOUNT_JSON_KEY",
	"global.infra.dns.auth.azureDNS.servicePrincipalAppPassword",
	"global.infra.objstore.s3.AWS_ACCESS_KEY_ID",
	"global.infra.objstore.s3.AWS_SECRET_ACCESS_KEY",
	"global.infra.objstore.azure.AZURE_ACCOUNT_KEY",
	"global.infra.objstore.gcs.GOOGLE_SERVICE_ACCOUNT_JSON_KEY",
	"s3proxy.s3proxy.auth.accessKeyId",
	"s3proxy.s3proxy.auth.secretAccessKey",
	"openfga.datastoreURI",
	"openfga.authn.preshared.keys",
	"perses.config.database.sql.password",
	"perses.config.security.encryption_key",
	"pgoutbox.app.config.database.password",
	"settings.db.auth.password",
	"settings.cache.auth.password",
	"settings.smtp.password",
	"settings.grafana.secretKey",
	"settings.inboxServer.adminJWTPrivateKey",
	"settings.nats.operatorCreds",
	"settings.nats.systemCreds",
	"settings.nats.adminCreds",
	"settings.nats.systemUserCreds",
	"settings.nats.adminUserCreds",
	"settings.openfga.preSharedKey",
	"settings.perses.encryptionKey",
	"settings.security.csrfSecretKey",
	"settings.security.oauth2JWTSecret",
	"setupJob.config.admin.password",
}

// FeatureCredPaths maps an opscenter-features child release to paths relative to its own values.
var FeatureCredPaths = map[string][]string{
	"stash-presets": {
		"kubestash.encryptionSecret",
		"kubestash.backend.s3.auth.AWS_ACCESS_KEY_ID",
		"kubestash.backend.s3.auth.AWS_SECRET_ACCESS_KEY",
		"kubestash.backend.azure.auth.AZURE_ACCOUNT_KEY",
		"kubestash.backend.gcs.auth.GOOGLE_SERVICE_ACCOUNT_JSON_KEY",
	},
	"service-gateway-presets": GatewayCredPaths,
	"prom-label-proxy": {
		"infra.tls.ca.key",
		"infra.tls.certificate.key",
		"infra.tls.jks.password",
		"infra.tls.jks.keystore",
	},
	"service-backend": {
		"server.oidc.clientSecret",
		"server.cookie.signingKey",
		"server.cookie.encryptionKey",
	},
	"inbox-server": {
		"james.secret.jwtPrivateKey",
		"james.secret.dkimPrivateKey",
		"james.env.jamesKeyStorePassword",
		"postgresql.auth.password",
	},
}

// GatewayCredPaths are relative to the service-gateway-presets values, which catalog-manager also uses for gateway HelmReleases.
var GatewayCredPaths = []string{
	"infra.tls.ca.key",
	"infra.tls.certificate.key",
	"infra.tls.jks.password",
	"infra.tls.jks.keystore",
	"infra.dns.auth.cloudflare.token",
	"infra.dns.auth.route53.AWS_ACCESS_KEY_ID",
	"infra.dns.auth.route53.AWS_SECRET_ACCESS_KEY",
	"infra.dns.auth.cloudDNS.GOOGLE_SERVICE_ACCOUNT_JSON_KEY",
	"infra.dns.auth.azureDNS.servicePrincipalAppPassword",
	"cluster.tls.ca.key",
}

// OpscenterFeaturesCredPaths are relative to the opscenter-features chart values.
func OpscenterFeaturesCredPaths() []string {
	paths := []string{"licenseServer.token"}
	for _, child := range slices.Sorted(maps.Keys(FeatureCredPaths)) {
		for _, p := range FeatureCredPaths[child] {
			paths = append(paths, "helm.releases."+child+".values."+p)
		}
	}
	return paths
}

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
	core "k8s.io/api/core/v1"
)

// CadenceSpec is the schema for the cadence chart values set by the ace chart.
// It covers only the subset of
// https://github.com/cadence-workflow/cadence-charts/blob/cadence-1.6.7/charts/cadence/values.yaml
// that ace overrides; every other value keeps the cadence chart default.
type CadenceSpec struct {
	Global    CadenceGlobalSpec    `json:"global"`
	Web       CadenceWebSpec       `json:"web"`
	Config    CadenceConfigSpec    `json:"config"`
	Cassandra CadenceCassandraSpec `json:"cassandra"`
	Schema    CadenceSchemaSpec    `json:"schema"`
}

// CadenceGlobalSpec holds the values shared by the cadence server services
// (frontend, matching, history and worker).
type CadenceGlobalSpec struct {
	Image CadenceImageReference `json:"image"`
	// Env is added to every cadence server container. The ace chart uses it to
	// hand the database password (POSTGRES_PWD) to the servers.
	Env []core.EnvVar `json:"env"`
}

type CadenceImageReference struct {
	Repository string `json:"repository"`
	Tag        string `json:"tag"`
	PullPolicy string `json:"pullPolicy"`
}

type CadenceWebSpec struct {
	Image   CadenceImageReference `json:"image"`
	Busybox CadenceBusyboxSpec    `json:"busybox"`
}

// CadenceBusyboxSpec toggles the wait-frontend init container of the web
// deployment. The cadence chart renders its pod security context with broken
// indentation, so the init container is rejected whenever fsGroup is set.
type CadenceBusyboxSpec struct {
	Enabled bool `json:"enabled"`
}

// CadenceCassandraSpec toggles the cassandra chart bundled with the cadence
// chart. The ace chart disables it and stores cadence data in its own
// postgres database.
type CadenceCassandraSpec struct {
	Enabled bool `json:"enabled"`
}

type CadenceSchemaSpec struct {
	CheckSchema CadenceCheckSchemaSpec `json:"checkSchema"`
}

type CadenceCheckSchemaSpec struct {
	Postgres CadenceCheckSchemaPostgresSpec `json:"postgres"`
}

// CadenceCheckSchemaPostgresSpec configures the wait-for-database init
// container of the schema server job. The cadence chart defaults to the
// unpinned, root alpine/psql:latest image.
type CadenceCheckSchemaPostgresSpec struct {
	Image                    CadenceImageReference `json:"image"`
	ContainerSecurityContext *core.SecurityContext `json:"containerSecurityContext,omitempty"`
}

type CadenceConfigSpec struct {
	Persistence CadencePersistenceSpec `json:"persistence"`
}

type CadencePersistenceSpec struct {
	Database CadenceDatabaseSpec `json:"database"`
}

type CadenceDatabaseSpec struct {
	// Driver is one of cassandra, mysql or postgres.
	Driver string         `json:"driver"`
	SQL    CadenceSQLSpec `json:"sql"`
}

type CadenceSQLSpec struct {
	Hosts            string `json:"hosts"`
	Dbname           string `json:"dbname"`
	VisibilityDbname string `json:"visibilityDbname"`
	User             string `json:"user"`
	// Password must be non-empty for the cadence chart to read the password
	// from POSTGRES_PWD. The ace chart sets POSTGRES_PWD to
	// settings.db.auth.password, so this value itself is never used.
	Password        string            `json:"password"`
	MaxConns        int               `json:"maxConns"`
	MaxIdleConns    int               `json:"maxIdleConns"`
	MaxConnLifetime string            `json:"maxConnLifetime"`
	TLS             CadenceSQLTLSSpec `json:"tls"`
}

type CadenceSQLTLSSpec struct {
	Enabled bool `json:"enabled"`
}

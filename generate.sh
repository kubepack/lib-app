#!/bin/bash

# Copyright AppsCode Inc. and Contributors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

CHART_VERSION=${CHART_VERSION:-v0.4.18}

go run cmd/gen-simple-editor/main.go --all --skipExisting=false \
    --chart-version=$CHART_VERSION

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/elasticsearch/custom \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --form-templates=monitoring/alert.yaml \
    --form-templates=capi/capa.yaml \
    --form-templates=capi/capg.yaml \
    --form-templates=capi/capz.yaml \
    --sample-name=elasticsearch \
    --resource.group=kubedb.com \
    --resource.version=v1alpha2 \
    --resource.name=elasticsearches

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/kafka/custom \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --sample-name=kafka \
    --resource.group=kubedb.com \
    --resource.version=v1alpha2 \
    --resource.name=kafkas

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/mariadb/cluster/prometheus.io/backupconfiguration/stash/tls/custom-auth/config-file/customize-pod-template \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --form-templates=monitoring/alert.yaml \
    --form-templates=capi/capa.yaml \
    --form-templates=capi/capg.yaml \
    --form-templates=capi/capz.yaml \
    --sample-name=mariadb \
    --resource.group=kubedb.com \
    --resource.version=v1alpha2 \
    --resource.name=mariadbs

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/mongodb/managed \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --form-templates=monitoring/alert.yaml \
    --form-templates=capi/capa.yaml \
    --form-templates=capi/capg.yaml \
    --form-templates=capi/capz.yaml \
    --sample-name=mongodb \
    --resource.group=kubedb.com \
    --resource.version=v1alpha2 \
    --resource.name=mongodbs

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/mysql/group-replication/prometheus.io/backupconfiguration/stash/tls/custom-auth/config-file/customize-pod-template \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --form-templates=monitoring/alert.yaml \
    --form-templates=capi/capa.yaml \
    --form-templates=capi/capg.yaml \
    --form-templates=capi/capz.yaml \
    --sample-name=mysql \
    --resource.group=kubedb.com \
    --resource.version=v1alpha2 \
    --resource.name=mysqls

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/pgbouncer/custom \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --sample-name=pgbouncer \
    --resource.group=kubedb.com \
    --resource.version=v1alpha2 \
    --resource.name=pgbouncers

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/postgres/cluster/prometheus.io/backupconfiguration/stash/tls/custom-auth/custom-auth-mode/custom-config/custom-pg-coordinator/custom-uid \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --form-templates=monitoring/alert.yaml \
    --form-templates=capi/capa.yaml \
    --form-templates=capi/capg.yaml \
    --form-templates=capi/capz.yaml \
    --sample-name=postgres \
    --resource.group=kubedb.com \
    --resource.version=v1alpha2 \
    --resource.name=postgreses

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/proxysql/custom \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --form-templates=monitoring/alert.yaml \
    --form-templates=capi/capa.yaml \
    --form-templates=capi/capg.yaml \
    --form-templates=capi/capz.yaml \
    --sample-name=proxysql \
    --resource.group=kubedb.com \
    --resource.version=v1alpha2 \
    --resource.name=proxysqls

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/redis/sentinel/prometheus.io/backupconfiguration/stash/tls/custom-auth/custom-config \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --form-templates=monitoring/alert.yaml \
    --form-templates=capi/capa.yaml \
    --form-templates=capi/capg.yaml \
    --form-templates=capi/capz.yaml \
    --sample-name=redis \
    --resource.group=kubedb.com \
    --resource.version=v1alpha2 \
    --resource.name=redises

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/redissentinel/prometheus.io/tls/custom-auth \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --sample-name=redissentinel \
    --resource.group=kubedb.com \
    --resource.version=v1alpha2 \
    --resource.name=redissentinels

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/kubevault/vaultserver/custom \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --form-templates=monitoring/alert.yaml \
    --form-templates=capi/capa.yaml \
    --form-templates=capi/capg.yaml \
    --form-templates=capi/capz.yaml \
    --sample-name=vault \
    --resource.group=kubevault.com \
    --resource.version=v1alpha1 \
    --resource.name=vaultservers

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/repository/custom \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --sample-name=s3 \
    --resource.group=stash.appscode.com \
    --resource.version=v1alpha1 \
    --resource.name=repositories

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/restoresession/custom \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --sample-name=restore-app \
    --resource.group=stash.appscode.com \
    --resource.version=v1beta1 \
    --resource.name=restoresessions

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/featuresets/opscenter-core \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --sample-name=not-used \
    --instance-name=opscenter-core \
    --resource.group=ui.k8s.appscode.com \
    --resource.version=v1alpha1 \
    --resource.name=featuresets \
    --gen-crd=false

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/featuresets/opscenter-monitoring \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --sample-name=not-used \
    --instance-name=opscenter-monitoring \
    --resource.group=ui.k8s.appscode.com \
    --resource.version=v1alpha1 \
    --resource.name=featuresets \
    --gen-crd=false

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/featuresets/opscenter-datastore \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --sample-name=not-used \
    --instance-name=opscenter-datastore \
    --resource.group=ui.k8s.appscode.com \
    --resource.version=v1alpha1 \
    --resource.name=featuresets \
    --gen-crd=false

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/featuresets/opscenter-backup \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --sample-name=not-used \
    --instance-name=opscenter-backup \
    --resource.group=ui.k8s.appscode.com \
    --resource.version=v1alpha1 \
    --resource.name=featuresets \
    --gen-crd=false

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/featuresets/opscenter-security \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --sample-name=not-used \
    --instance-name=opscenter-security \
    --resource.group=ui.k8s.appscode.com \
    --resource.version=v1alpha1 \
    --resource.name=featuresets \
    --gen-crd=false

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/featuresets/opscenter-cost \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --sample-name=not-used \
    --instance-name=opscenter-cost \
    --resource.group=ui.k8s.appscode.com \
    --resource.version=v1alpha1 \
    --resource.name=featuresets \
    --gen-crd=false

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/featuresets/opscenter-networking \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --sample-name=not-used \
    --instance-name=opscenter-networking \
    --resource.group=ui.k8s.appscode.com \
    --resource.version=v1alpha1 \
    --resource.name=featuresets \
    --gen-crd=false

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/featuresets/opscenter-tools \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --sample-name=not-used \
    --instance-name=opscenter-tools \
    --resource.group=ui.k8s.appscode.com \
    --resource.version=v1alpha1 \
    --resource.name=featuresets \
    --gen-crd=false

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/featuresets/opscenter-presets \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --sample-name=not-used \
    --instance-name=opscenter-presets \
    --resource.group=ui.k8s.appscode.com \
    --resource.version=v1alpha1 \
    --resource.name=featuresets \
    --gen-crd=false

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/featuresets/dbaas-capa \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --sample-name=not-used \
    --instance-name=dbaas-capa \
    --resource.group=ui.k8s.appscode.com \
    --resource.version=v1alpha1 \
    --resource.name=featuresets \
    --gen-crd=false

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/featuresets/dbaas-capg \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --sample-name=not-used \
    --instance-name=dbaas-capg \
    --resource.group=ui.k8s.appscode.com \
    --resource.version=v1alpha1 \
    --resource.name=featuresets \
    --gen-crd=false

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/featuresets/dbaas-capz \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --sample-name=not-used \
    --instance-name=dbaas-capz \
    --resource.group=ui.k8s.appscode.com \
    --resource.version=v1alpha1 \
    --resource.name=featuresets \
    --gen-crd=false

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/featuresets/crossplane \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --sample-name=not-used \
    --instance-name=crossplane \
    --resource.group=ui.k8s.appscode.com \
    --resource.version=v1alpha1 \
    --resource.name=featuresets \
    --gen-crd=false

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/featuresets/opscenter-storage \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --sample-name=not-used \
    --instance-name=opscenter-storage \
    --resource.group=ui.k8s.appscode.com \
    --resource.version=v1alpha1 \
    --resource.name=featuresets \
    --gen-crd=false

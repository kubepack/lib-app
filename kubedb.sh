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

CHART_VERSION=${CHART_VERSION:-v0.4.21}

go run cmd/gen-simple-editor/main.go --all --skipExisting=false \
    --chart-version=$CHART_VERSION

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/elasticsearch/custom \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --form-templates=monitoring/alert.yaml \
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
    --sample-name=mariadb \
    --resource.group=kubedb.com \
    --resource.version=v1alpha2 \
    --resource.name=mariadbs

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/mongodb/managed \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --form-templates=monitoring/alert.yaml \
    --sample-name=mongodb \
    --resource.group=kubedb.com \
    --resource.version=v1alpha2 \
    --resource.name=mongodbs

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/mysql/group-replication/prometheus.io/backupconfiguration/stash/tls/custom-auth/config-file/customize-pod-template \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --form-templates=monitoring/alert.yaml \
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
    --sample-name=postgres \
    --resource.group=kubedb.com \
    --resource.version=v1alpha2 \
    --resource.name=postgreses

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/proxysql/custom \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --form-templates=monitoring/alert.yaml \
    --sample-name=proxysql \
    --resource.group=kubedb.com \
    --resource.version=v1alpha2 \
    --resource.name=proxysqls

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/redis/sentinel/prometheus.io/backupconfiguration/stash/tls/custom-auth/custom-config \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --form-templates=monitoring/alert.yaml \
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
    --sample-name=vault \
    --resource.group=kubevault.com \
    --resource.version=v1alpha1 \
    --resource.name=vaultservers

# new dbs
#################################################################################

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/druid/druid.yaml \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --form-templates=monitoring/alert.yaml \
    --sample-name=druid \
    --resource.group=kubedb.com \
    --resource.version=v1alpha2 \
    --resource.name=druids

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/ferretdb/ferretdb.yaml \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --form-templates=monitoring/alert.yaml \
    --sample-name=ferret \
    --resource.group=kubedb.com \
    --resource.version=v1alpha2 \
    --resource.name=ferretdbs

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/mssqlserver/sqlserver-ag.yaml \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --form-templates=monitoring/alert.yaml \
    --sample-name=sqlserver-ag \
    --resource.group=kubedb.com \
    --resource.version=v1alpha2 \
    --resource.name=mssqlservers

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/pgpool/pgpool.yaml \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --form-templates=monitoring/alert.yaml \
    --sample-name=pgpool \
    --resource.group=kubedb.com \
    --resource.version=v1alpha2 \
    --resource.name=pgpools

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/rabbitmq/rabbitmq.yaml \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --form-templates=monitoring/alert.yaml \
    --sample-name=rabbitmq \
    --resource.group=kubedb.com \
    --resource.version=v1alpha2 \
    --resource.name=rabbitmqs

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/singlestore/cluster-mode/cluster.yaml \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --form-templates=monitoring/alert.yaml \
    --sample-name=singlestore \
    --resource.group=kubedb.com \
    --resource.version=v1alpha2 \
    --resource.name=singlestores

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/solr/topology.yaml \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --form-templates=monitoring/alert.yaml \
    --sample-name=solr \
    --resource.group=kubedb.com \
    --resource.version=v1alpha2 \
    --resource.name=solrs

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/zookeeper/sample.yaml \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --form-templates=monitoring/alert.yaml \
    --sample-name=zookeeper \
    --resource.group=kubedb.com \
    --resource.version=v1alpha2 \
    --resource.name=zookeepers

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

go run cmd/gen-simple-editor/main.go --all --skipExisting=false

go run cmd/fuse-chart/*.go \
    --sample-dir=/home/tamal/go/src/github.com/appscode/kubedb-samples/mongodb/sharded/prometheus.io/backupconfiguration/stash/tls/custom-auth/sharded \
    --chart-dir=/home/tamal/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --sample-name=mongodb \
    --resource.group=kubedb.com \
    --resource.version=v1alpha2 \
    --resource.name=mongodbs

go run cmd/fuse-chart/*.go \
    --sample-dir=/home/tamal/go/src/github.com/appscode/kubedb-samples/postgres/cluster/prometheus.io/backupconfiguration/stash/tls/custom-auth/custom-auth-mode/custom-config/custom-pg-coordinator/custom-uid \
    --chart-dir=/home/tamal/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --sample-name=postgres \
    --resource.group=kubedb.com \
    --resource.version=v1alpha2 \
    --resource.name=postgreses

go run cmd/fuse-chart/*.go \
    --sample-dir=/home/tamal/go/src/github.com/appscode/kubedb-samples/mariadb/cluster/prometheus.io/backupconfiguration/stash/tls/custom-auth/config-file/customize-pod-template \
    --chart-dir=/home/tamal/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --sample-name=mariadb \
    --resource.group=kubedb.com \
    --resource.version=v1alpha2 \
    --resource.name=mariadbs

go run cmd/fuse-chart/*.go \
    --sample-dir=/home/tamal/go/src/github.com/appscode/kubedb-samples/elasticsearch/custom \
    --chart-dir=/home/tamal/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --sample-name=elasticsearch \
    --resource.group=kubedb.com \
    --resource.version=v1alpha2 \
    --resource.name=elasticsearches

go run cmd/fuse-chart/*.go \
    --sample-dir=/home/tamal/go/src/github.com/appscode/kubedb-samples/mysql/group-replication/prometheus.io/backupconfiguration/stash/tls/custom-auth/config-file/customize-pod-template \
    --chart-dir=/home/tamal/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --sample-name=mysql \
    --resource.group=kubedb.com \
    --resource.version=v1alpha2 \
    --resource.name=mysqls

go run cmd/fuse-chart/*.go \
    --sample-dir=/home/tamal/go/src/github.com/appscode/kubedb-samples/redis/sentinel/prometheus.io/backupconfiguration/stash/tls/custom-auth/custom-config \
    --chart-dir=/home/tamal/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --sample-name=redis \
    --resource.group=kubedb.com \
    --resource.version=v1alpha2 \
    --resource.name=redises

go run cmd/fuse-chart/*.go \
    --sample-dir=/home/tamal/go/src/github.com/appscode/kubedb-samples/redissentinel/prometheus.io/tls/custom-auth/custom-config \
    --chart-dir=/home/tamal/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --sample-name=redissentinel \
    --resource.group=kubedb.com \
    --resource.version=v1alpha2 \
    --resource.name=redissentinels

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

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/aws/documentdb \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --sample-name=example \
    --resource.group=docdb.aws.kubedb.com \
    --resource.version=v1alpha1 \
    --resource.name=clusters

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/aws/dynamodb \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --sample-name=example \
    --resource.group=dynamodb.aws.kubedb.com \
    --resource.version=v1alpha1 \
    --resource.name=tables

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/aws/elasticache \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --sample-name=example \
    --resource.group=elasticache.aws.kubedb.com \
    --resource.version=v1alpha1 \
    --resource.name=clusters

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/aws/rds/postgres \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --sample-name=example \
    --resource.group=rds.aws.kubedb.com \
    --resource.version=v1alpha1 \
    --resource.name=clusters

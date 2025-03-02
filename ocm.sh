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

echo
basename "$0"
echo "---------------"

CHART_VERSION=${CHART_VERSION:-v0.14.0}

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/featuresets/ocm-hub \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --sample-name=not-used \
    --instance-name=ocm-hub \
    --resource.group=ui.k8s.appscode.com \
    --resource.version=v1alpha1 \
    --resource.name=featuresets \
    --gen-crd=false

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/featuresets/ocm-mc \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --sample-name=not-used \
    --instance-name=ocm-mc \
    --resource.group=ui.k8s.appscode.com \
    --resource.version=v1alpha1 \
    --resource.name=featuresets \
    --gen-crd=false

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/featuresets/ocm-spoke \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --sample-name=not-used \
    --instance-name=ocm-spoke \
    --resource.group=ui.k8s.appscode.com \
    --resource.version=v1alpha1 \
    --resource.name=featuresets \
    --gen-crd=false

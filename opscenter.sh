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

CHART_VERSION=${CHART_VERSION:-v0.6.0}

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
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/featuresets/opscenter-observability \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --sample-name=not-used \
    --instance-name=opscenter-observability \
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

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/featuresets/saas-core \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --sample-name=not-used \
    --instance-name=saas-core \
    --resource.group=ui.k8s.appscode.com \
    --resource.version=v1alpha1 \
    --resource.name=featuresets \
    --gen-crd=false

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/featuresets/opscenter-secret-management \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --sample-name=not-used \
    --instance-name=opscenter-secret-management \
    --resource.group=ui.k8s.appscode.com \
    --resource.version=v1alpha1 \
    --resource.name=featuresets \
    --gen-crd=false

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/featuresets/opscenter-policy-management \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --sample-name=not-used \
    --instance-name=opscenter-policy-management \
    --resource.group=ui.k8s.appscode.com \
    --resource.version=v1alpha1 \
    --resource.name=featuresets \
    --gen-crd=false

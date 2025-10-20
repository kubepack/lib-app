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

CHART_VERSION=${CHART_VERSION:-v0.26.0}

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/kubestash/backupstorage/s3 \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --sample-name=s3 \
    --resource.group=storage.kubestash.com \
    --resource.version=v1alpha1 \
    --resource.name=backupstorages

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/kubestash/repository/custom \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --sample-name=ace-repo \
    --resource.group=storage.kubestash.com \
    --resource.version=v1alpha1 \
    --resource.name=repositories

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/kubestash/backupconfiguration/custom \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --sample-name=backup-app \
    --resource.group=core.kubestash.com \
    --resource.version=v1alpha1 \
    --resource.name=backupconfigurations

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/kubestash/backupsession/custom \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --sample-name=backup-app \
    --resource.group=core.kubestash.com \
    --resource.version=v1alpha1 \
    --resource.name=backupsessions

go run cmd/fuse-chart/*.go \
    --sample-dir=$HOME/go/src/go.bytebuilders.dev/ui-samples/kubestash/restoresession/custom \
    --chart-dir=$HOME/go/src/go.bytebuilders.dev/ui-wizards/charts \
    --chart-version=$CHART_VERSION \
    --sample-name=restore-app \
    --resource.group=core.kubestash.com \
    --resource.version=v1alpha1 \
    --resource.name=restoresessions

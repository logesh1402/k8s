#!/bin/bash

# Copyright 2015 The Kubernetes Authors All rights reserved.
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

# Generates updated api-reference docs from the latest swagger spec.
# Usage: ./update-api-reference-docs.sh <absolute output path>

set -o errexit
set -o nounset
set -o pipefail

KUBE_ROOT=$(dirname "${BASH_SOURCE}")/..
DEFAULT_OUTPUT_PATH="$PWD/${KUBE_ROOT}/docs/api-reference"
OUTPUT=${1:-${DEFAULT_OUTPUT_PATH}}

echo "Generating api reference docs at ${OUTPUT}"

V1_PATH="${OUTPUT}/v1/"
V1BETA1_PATH="${OUTPUT}/extensions/v1beta1"
SWAGGER_PATH="$PWD/${KUBE_ROOT}/api/swagger-spec/"

echo "Reading swagger spec from: ${SWAGGER_PATH}"

mkdir -p $V1_PATH
mkdir -p $V1BETA1_PATH

docker run --privileged -v $V1_PATH:/output -v ${SWAGGER_PATH}:/swagger-source gcr.io/google_containers/gen-swagger-docs:v3 \
    v1 \
    https://raw.githubusercontent.com/kubernetes/kubernetes/master/pkg/api/v1/register.go

docker run --privileged -v $V1BETA1_PATH:/output -v ${SWAGGER_PATH}:/swagger-source gcr.io/google_containers/gen-swagger-docs:v3 \
    v1beta1 \
    https://raw.githubusercontent.com/kubernetes/kubernetes/master/pkg/apis/extensions/v1beta1/register.go

# ex: ts=2 sw=2 et filetype=sh

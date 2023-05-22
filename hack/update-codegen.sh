#!/bin/bash
set -x
vendor/k8s.io/code-generator/generate-groups.sh all \
github.com/Neaj-Morshad-101/crd-controller/pkg/client \
github.com/Neaj-Morshad-101/crd-controller/pkg/apis \
neajmorshad.dev:v1beta1 \
--go-header-file /home/appscodepc/go/src/github.com/Neaj-Morshad-101/crd-controller/hack/boilerplate.go.txt
#
controller-gen rbac:roleName=controller-perms crd paths=github.com/Neaj-Morshad-101/crd-controller/pkg/apis/neajmorshad.dev/v1beta1 \
crd:crdVersions=v1 output:crd:dir=/home/appscodepc/go/src/github.com/Neaj-Morshad-101/crd-controller/manifests output:stdout


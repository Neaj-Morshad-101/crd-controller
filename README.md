
# The Custom Controller (crd-controller) is created to manage a Custom Resource (Kluster) #
## Learning Custom Resource Definition and Sample Controller ##


## Goals: ##
- Understand CRs
- Define a CR using CRD
- Understand Custom Controller
- Write a Custom Controller (crd-controller) to manage a Custom Resource (Kluster)

## Info: ##
- Group Name: `neajmorshad.dev`
- Version Name: `v1beta1`
- Resource Name: `Kluster`

## Procedure to run this project: ##
- go to the project path. Example: $HOME/go/src/github.com/Neaj-Morshad-101/crd-controller/
- import `"k8s.io/code-generator"` into `main.go`
- run `go install sigs.k8s.io/controller-tools/cmd/controller-gen@latest`
- run `go mod tidy;go mod vendor`
- run chmod +x `./hack/update-codegen.sh`
- run `chmod +x vendor/k8s.io/code-generator/generate-groups.sh`
- run `hack/update-codegen.sh`
- again run `go mod tidy;go mod vendor`


## To generate controller-gen: ##
- Run `depelopmentDir=$(pwd)`
- Then run: `controller-gen rbac:roleName=controller-perms crd paths=github.com/Neaj-Morshad-101/crd-controller/pkg/apis/neajmorshad.dev/v1beta1 crd:crdVersions=v1 output:crd:dir=$depelopmentDir/manifests output:stdout`


## Deploy Custom resources ##
- kubectl apply -f `neajmorshad.dev_klusters.yaml`
- Create an example yaml file like `manifests/kluster.yaml`.
- Run `kubectl apply -f manifests/kluster.yaml`.
- Run `kubectl get kluster`.


## Resource ##
https://www.linkedin.com/pulse/kubernetes-custom-controllers-part-1-kritik-sachdeva/
https://www.linkedin.com/pulse/kubernetes-custom-controller-part-2-kritik-sachdeva/
Workqueue example:
https://github.com/kubernetes/client-go/tree/master/examples/workqueue
Sample-controller
https://github.com/kubernetes/sample-controller
Code-Generator
https://github.com/kubernetes/code-generator

## This YouTube video will help to understand the basic structure of a Custom Controller ##
https://www.youtube.com/watch?v=89PdRvRUcPU&list=PLh4KH3LtJvRTtFWz1WGlyDa7cKjj2Sns0&index=4&ab_channel=VivekSingh


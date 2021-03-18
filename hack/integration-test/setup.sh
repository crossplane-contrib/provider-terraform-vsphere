#! /usr/bin/env bash

go install github.com/vmware/govmomi/govc

./hack/integration-test/kind-docker-init.sh

docker build . -t=localhost:5000/vcsim:v0.1.0
docker push localhost:5000/vcsim:v0.1.0
kubectl apply -f ./hack/integration-test/cluster

echo 'set up your environment to interact with the vcsim server:'
echo 'export GOVC_URL=https://user:pass@localhost:30443'
echo 'export GOVC_INSECURE=true'

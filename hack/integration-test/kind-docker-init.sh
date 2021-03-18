#!/bin/sh
set -o errexit
set -e

# create registry container unless it already exists
reg_name='kind-registry'
reg_port='5000'
running="$(docker inspect -f '{{.State.Running}}' "${reg_name}" 2>/dev/null || true)"
if [ "${running}" != 'true' ]; then
  docker run \
    -d --restart=always -p "127.0.0.1:${reg_port}:5000" --name "${reg_name}" \
    registry:2
fi

# projectdir="$( cd "$( dirname "${BASH_SOURCE[0]}")" && pwd )"

# get the build environment variables from the special build.vars target in the main makefile
eval $(make --no-print-directory build.vars)

# ------------------------------

HOSTARCH="${HOSTARCH:-amd64}"
BUILD_IMAGE="${BUILD_REGISTRY}/${PROJECT_NAME}-${HOSTARCH}"
CONTROLLER_IMAGE="${BUILD_REGISTRY}/${PROJECT_NAME}-controller-${HOSTARCH}"

version_tag="$(cat ./_output/version)"
# tag as latest version to load into kind cluster
PACKAGE_CONTROLLER_IMAGE="${DOCKER_REGISTRY}/${PROJECT_NAME}-controller:${version_tag}"
K8S_CLUSTER="${K8S_CLUSTER:-${BUILD_REGISTRY}-inttests}"

CROSSPLANE_NAMESPACE="crossplane-system"
PACKAGE_NAME="provider-terraform-vsphere"

CACHE_PATH="./.work/inttest-package-cache"
mkdir -p "${CACHE_PATH}"
echo "created cache dir at ${CACHE_PATH}"
docker save "${BUILD_IMAGE}" -o "${CACHE_PATH}/${PACKAGE_NAME}.xpkg" && chmod 644 "${CACHE_PATH}/${PACKAGE_NAME}.xpkg"

# create a cluster with the local registry enabled in containerd
cat <<EOF | kind create cluster --name=vcsim --config=-
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
  extraPortMappings:
  - containerPort: 30443
    hostPort: 30443
    protocol: TCP
  extraMounts:
  - hostPath: "${CACHE_PATH}/"
    containerPath: /cache
containerdConfigPatches:
- |-
  [plugins."io.containerd.grpc.v1.cri".registry.mirrors."localhost:${reg_port}"]
    endpoint = ["http://${reg_name}:${reg_port}"]
EOF

# connect the registry to the cluster network
# (the network may already be connected)
docker network connect "kind" "${reg_name}" || true

# Document the local registry
# https://github.com/kubernetes/enhancements/tree/master/keps/sig-cluster-lifecycle/generic/1755-communicating-a-local-registry
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: ConfigMap
metadata:
  name: local-registry-hosting
  namespace: kube-public
data:
  localRegistryHosting.v1: |
    host: "localhost:${reg_port}"
    help: "https://kind.sigs.k8s.io/docs/user/local-registry/"
EOF

# tag controller image and load it into kind cluster
docker tag "${CONTROLLER_IMAGE}" "${PACKAGE_CONTROLLER_IMAGE}"
"${KIND}" load docker-image "${PACKAGE_CONTROLLER_IMAGE}" --name=vcsim

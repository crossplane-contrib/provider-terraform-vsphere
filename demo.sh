#!/bin/sh

kubectl create namespace crossplane-system

kubectl apply -f ./examples/pv.yaml
kubectl apply -f ./examples/pvc.yaml

helm repo add crossplane-stable https://charts.crossplane.io/stable
helm repo update

helm install crossplane --namespace crossplane-system crossplane-stable/crossplane --set packageCache.pvc=package-cache

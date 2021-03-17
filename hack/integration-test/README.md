docker-entrypoint.sh borrowed from https://github.com/nimmis/docker-vcsim

setup.sh will:
- install govc (for verification of results)
- setup a kind cluster running a container registry and port 30443 briged to allow ClusterIP services to be accessible on localhost
- build a vcsim docker container and push it to the kind cluster's registry
- start a deployment+service to run vcsim and expose it on 30443

In order to use govc, take note of the export statements provided in setup.sh output. 

The `provider` folder contains a provider config and secret for use with the vcsim server (which expects a username=`user` and password=`pass`). Apply these for use by the provider:

```
kubectl apply -f hack/integration-test/provider
```

The `resources` folder contains example resources that can be used to verify the functionality of the provider. Apply the resources in this folder and then use the following govc commands to verify each type:

```
kubectl apply -f hack/integration-test/resources/tag-category.yaml
govc tags.category.ls --json=true
kubectl apply -f hack/integration-test/resources/tag.yaml
govc tags.ls --json=true
kubectl apply -f hack/integration-test/resources/folder.yaml
govc folder.info root --json=true
kubectl apply -f hack/integration-test/resources/datacenter.yaml
govc datacenter.info /root-dc/dczero
````

docker-entrypoint.sh borrowed from https://github.com/nimmis/docker-vcsim

after running setup.sh you should have a kind cluster running vcsim. in order to use govc, take note of the export statements provided in setup.sh output. The `resources` folder contains example resources that can be used to verify the functionality of the provider. Apply the resources in this folder and then use the following govc commands to verify each type:

```
govc tags.category.ls --json=true
govc tags.ls --json=true
govc folder.info root --json=true
govc datacenter.info /root-dc/dczero
````

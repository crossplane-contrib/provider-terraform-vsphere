module github.com/crossplane-contrib/provider-terraform-vsphere

go 1.14

require (
	cloud.google.com/go/pubsub v1.3.1 // indirect
	github.com/crossplane-contrib/terraform-provider-dl v0.0.0-20210326151831-95dc732e4246 // indirect
	github.com/crossplane-contrib/terraform-provider-gen v0.0.0-20210412223606-81a63c2f0a87 // indirect
	github.com/crossplane-contrib/terraform-runtime v0.0.0-20210317191104-9eb36dba841c
	github.com/crossplane/crossplane-runtime v0.12.0
	github.com/crossplane/crossplane-tools v0.0.0-20210320162312-1baca298c527
	github.com/google/uuid v1.2.0 // indirect
	github.com/hashicorp/terraform v0.13.5
	github.com/kr/pretty v0.2.1 // indirect
	github.com/pkg/errors v0.9.1
	github.com/zclconf/go-cty v1.7.0
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	honnef.co/go/tools v0.0.1-2020.1.5 // indirect
	k8s.io/api v0.18.8
	k8s.io/apimachinery v0.18.8
	sigs.k8s.io/controller-runtime v0.6.2
	sigs.k8s.io/controller-tools v0.4.0
)

replace (
	github.com/aws/aws-sdk-go-v2 => github.com/aws/aws-sdk-go-v2 v0.23.0
	k8s.io/api => k8s.io/api v0.18.6
	k8s.io/apimachinery => k8s.io/apimachinery v0.18.6
	k8s.io/client-go => k8s.io/client-go v0.18.6
)

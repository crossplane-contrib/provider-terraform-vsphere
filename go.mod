module github.com/crossplane-contrib/provider-terraform-vsphere

go 1.14

replace github.com/crossplane-contrib/terraform-runtime => /Users/kasey/src/crossplane-contrib/terraform-provider-runtime

require (
	github.com/crossplane-contrib/terraform-provider-gen v0.0.0-20210211205231-04187a50b2e1 // indirect
	github.com/crossplane-contrib/terraform-runtime v0.0.0-00010101000000-000000000000
	github.com/crossplane/crossplane-runtime v0.12.0
	github.com/crossplane/crossplane-tools v0.0.0-20201201125637-9ddc70edfd0d
	github.com/crossplane/provider-gcp v0.15.0 // indirect
	github.com/google/uuid v1.2.0 // indirect
	github.com/hashicorp/terraform v0.13.5
	github.com/kr/pretty v0.2.1 // indirect
	github.com/pkg/errors v0.9.1
	github.com/vmware/govmomi v0.24.0 // indirect
	github.com/zclconf/go-cty v1.7.0
	go.pedge.io/lion v0.0.0-20190619200210-304b2f426641 // indirect
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
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

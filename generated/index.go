package generated

import (
	"github.com/crossplane-contrib/provider-terraform-vsphere/generated/provider/v1alpha1"
	"github.com/crossplane-contrib/terraform-runtime/pkg/plugin"
)

const ProviderReferenceName string = "vsphere"

func Index(idxr *plugin.Indexer) error {
	for _, impl := range ResourceImplementations {
		err := idxr.Overlay(impl)
		if err != nil {
			return err
		}
	}
	return nil
}

func ProviderInit() *plugin.ProviderInit {
	return v1alpha1.GetProviderInit()
}

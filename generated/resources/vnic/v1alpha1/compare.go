/*
	Copyright 2019 The Crossplane Authors.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	    http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package v1alpha1

import (
	"github.com/crossplane-contrib/terraform-runtime/pkg/plugin"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
)

//mergeManagedResourceEntrypointTemplate
type resourceMerger struct{}

func (r *resourceMerger) MergeResources(kube resource.Managed, prov resource.Managed) plugin.MergeDescription {
	k := kube.(*Vnic)
	p := prov.(*Vnic)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeVnic_DistributedPortGroup(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVnic_DistributedSwitchPort(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVnic_Host(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVnic_Mac(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVnic_Mtu(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVnic_Netstack(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVnic_Portgroup(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVnic_Ipv4(&k.Spec.ForProvider.Ipv4, &p.Spec.ForProvider.Ipv4, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVnic_Ipv6(&k.Spec.ForProvider.Ipv6, &p.Spec.ForProvider.Ipv6, md)
	if updated {
		anyChildUpdated = true
	}

	for key, v := range p.Annotations {
		if k.Annotations[key] != v {
			k.Annotations[key] = v
			md.AnnotationsUpdated = true
		}
	}
	md.AnyFieldUpdated = anyChildUpdated
	return *md
}

//mergePrimitiveTemplateSpec
func MergeVnic_DistributedPortGroup(k *VnicParameters, p *VnicParameters, md *plugin.MergeDescription) bool {
	if k.DistributedPortGroup != p.DistributedPortGroup {
		p.DistributedPortGroup = k.DistributedPortGroup
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVnic_DistributedSwitchPort(k *VnicParameters, p *VnicParameters, md *plugin.MergeDescription) bool {
	if k.DistributedSwitchPort != p.DistributedSwitchPort {
		p.DistributedSwitchPort = k.DistributedSwitchPort
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVnic_Host(k *VnicParameters, p *VnicParameters, md *plugin.MergeDescription) bool {
	if k.Host != p.Host {
		p.Host = k.Host
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVnic_Mac(k *VnicParameters, p *VnicParameters, md *plugin.MergeDescription) bool {
	if k.Mac != p.Mac {
		p.Mac = k.Mac
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVnic_Mtu(k *VnicParameters, p *VnicParameters, md *plugin.MergeDescription) bool {
	if k.Mtu != p.Mtu {
		p.Mtu = k.Mtu
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVnic_Netstack(k *VnicParameters, p *VnicParameters, md *plugin.MergeDescription) bool {
	if k.Netstack != p.Netstack {
		p.Netstack = k.Netstack
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVnic_Portgroup(k *VnicParameters, p *VnicParameters, md *plugin.MergeDescription) bool {
	if k.Portgroup != p.Portgroup {
		p.Portgroup = k.Portgroup
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructTemplateSpec
func MergeVnic_Ipv4(k *Ipv4, p *Ipv4, md *plugin.MergeDescription) bool {
	updated := false
	anyChildUpdated := false
	updated = MergeVnic_Ipv4_Netmask(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVnic_Ipv4_Dhcp(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVnic_Ipv4_Gw(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVnic_Ipv4_Ip(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveTemplateSpec
func MergeVnic_Ipv4_Netmask(k *Ipv4, p *Ipv4, md *plugin.MergeDescription) bool {
	if k.Netmask != p.Netmask {
		p.Netmask = k.Netmask
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVnic_Ipv4_Dhcp(k *Ipv4, p *Ipv4, md *plugin.MergeDescription) bool {
	if k.Dhcp != p.Dhcp {
		p.Dhcp = k.Dhcp
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVnic_Ipv4_Gw(k *Ipv4, p *Ipv4, md *plugin.MergeDescription) bool {
	if k.Gw != p.Gw {
		p.Gw = k.Gw
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVnic_Ipv4_Ip(k *Ipv4, p *Ipv4, md *plugin.MergeDescription) bool {
	if k.Ip != p.Ip {
		p.Ip = k.Ip
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructTemplateSpec
func MergeVnic_Ipv6(k *Ipv6, p *Ipv6, md *plugin.MergeDescription) bool {
	updated := false
	anyChildUpdated := false
	updated = MergeVnic_Ipv6_Autoconfig(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVnic_Ipv6_Dhcp(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVnic_Ipv6_Gw(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVnic_Ipv6_Addresses(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveTemplateSpec
func MergeVnic_Ipv6_Autoconfig(k *Ipv6, p *Ipv6, md *plugin.MergeDescription) bool {
	if k.Autoconfig != p.Autoconfig {
		p.Autoconfig = k.Autoconfig
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVnic_Ipv6_Dhcp(k *Ipv6, p *Ipv6, md *plugin.MergeDescription) bool {
	if k.Dhcp != p.Dhcp {
		p.Dhcp = k.Dhcp
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVnic_Ipv6_Gw(k *Ipv6, p *Ipv6, md *plugin.MergeDescription) bool {
	if k.Gw != p.Gw {
		p.Gw = k.Gw
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeVnic_Ipv6_Addresses(k *Ipv6, p *Ipv6, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.Addresses, p.Addresses) {
		p.Addresses = k.Addresses
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

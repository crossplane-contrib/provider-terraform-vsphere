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
	"fmt"

	ctwhy "github.com/crossplane-contrib/terraform-runtime/pkg/plugin/cty"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
	"github.com/zclconf/go-cty/cty"
)

type ctyDecoder struct{}

func (e *ctyDecoder) DecodeCty(mr resource.Managed, ctyValue cty.Value, schema *providers.Schema) (resource.Managed, error) {
	r, ok := mr.(*Vnic)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeVnic(r, ctyValue)
}

func DecodeVnic(prev *Vnic, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeVnic_DistributedPortGroup(&new.Spec.ForProvider, valMap)
	DecodeVnic_DistributedSwitchPort(&new.Spec.ForProvider, valMap)
	DecodeVnic_Host(&new.Spec.ForProvider, valMap)
	DecodeVnic_Mac(&new.Spec.ForProvider, valMap)
	DecodeVnic_Mtu(&new.Spec.ForProvider, valMap)
	DecodeVnic_Netstack(&new.Spec.ForProvider, valMap)
	DecodeVnic_Portgroup(&new.Spec.ForProvider, valMap)
	DecodeVnic_Ipv4(&new.Spec.ForProvider.Ipv4, valMap)
	DecodeVnic_Ipv6(&new.Spec.ForProvider.Ipv6, valMap)

	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeVnic_DistributedPortGroup(p *VnicParameters, vals map[string]cty.Value) {
	p.DistributedPortGroup = ctwhy.ValueAsString(vals["distributed_port_group"])
}

//primitiveTypeDecodeTemplate
func DecodeVnic_DistributedSwitchPort(p *VnicParameters, vals map[string]cty.Value) {
	p.DistributedSwitchPort = ctwhy.ValueAsString(vals["distributed_switch_port"])
}

//primitiveTypeDecodeTemplate
func DecodeVnic_Host(p *VnicParameters, vals map[string]cty.Value) {
	p.Host = ctwhy.ValueAsString(vals["host"])
}

//primitiveTypeDecodeTemplate
func DecodeVnic_Mac(p *VnicParameters, vals map[string]cty.Value) {
	p.Mac = ctwhy.ValueAsString(vals["mac"])
}

//primitiveTypeDecodeTemplate
func DecodeVnic_Mtu(p *VnicParameters, vals map[string]cty.Value) {
	p.Mtu = ctwhy.ValueAsInt64(vals["mtu"])
}

//primitiveTypeDecodeTemplate
func DecodeVnic_Netstack(p *VnicParameters, vals map[string]cty.Value) {
	p.Netstack = ctwhy.ValueAsString(vals["netstack"])
}

//primitiveTypeDecodeTemplate
func DecodeVnic_Portgroup(p *VnicParameters, vals map[string]cty.Value) {
	p.Portgroup = ctwhy.ValueAsString(vals["portgroup"])
}

//containerCollectionSingletonTypeDecodeTemplate
func DecodeVnic_Ipv4(p *Ipv4, vals map[string]cty.Value) {
	if vals["ipv4"].IsNull() {
		p = nil
		return
	}
	rvals := ctwhy.ValueAsList(vals["ipv4"])
	if len(rvals) == 0 {
		p = nil
		return
	}
	// this template should be used when single dictionary/object values are nested in sets/lists
	// if rvals turns out to be a list with > 1 elements, something has broken with that heuristic
	valMap := rvals[0].AsValueMap()
	DecodeVnic_Ipv4_Dhcp(p, valMap)
	DecodeVnic_Ipv4_Gw(p, valMap)
	DecodeVnic_Ipv4_Ip(p, valMap)
	DecodeVnic_Ipv4_Netmask(p, valMap)
}

//primitiveTypeDecodeTemplate
func DecodeVnic_Ipv4_Dhcp(p *Ipv4, vals map[string]cty.Value) {
	p.Dhcp = ctwhy.ValueAsBool(vals["dhcp"])
}

//primitiveTypeDecodeTemplate
func DecodeVnic_Ipv4_Gw(p *Ipv4, vals map[string]cty.Value) {
	p.Gw = ctwhy.ValueAsString(vals["gw"])
}

//primitiveTypeDecodeTemplate
func DecodeVnic_Ipv4_Ip(p *Ipv4, vals map[string]cty.Value) {
	p.Ip = ctwhy.ValueAsString(vals["ip"])
}

//primitiveTypeDecodeTemplate
func DecodeVnic_Ipv4_Netmask(p *Ipv4, vals map[string]cty.Value) {
	p.Netmask = ctwhy.ValueAsString(vals["netmask"])
}

//containerCollectionSingletonTypeDecodeTemplate
func DecodeVnic_Ipv6(p *Ipv6, vals map[string]cty.Value) {
	if vals["ipv6"].IsNull() {
		p = nil
		return
	}
	rvals := ctwhy.ValueAsList(vals["ipv6"])
	if len(rvals) == 0 {
		p = nil
		return
	}
	// this template should be used when single dictionary/object values are nested in sets/lists
	// if rvals turns out to be a list with > 1 elements, something has broken with that heuristic
	valMap := rvals[0].AsValueMap()
	DecodeVnic_Ipv6_Autoconfig(p, valMap)
	DecodeVnic_Ipv6_Dhcp(p, valMap)
	DecodeVnic_Ipv6_Gw(p, valMap)
	DecodeVnic_Ipv6_Addresses(p, valMap)
}

//primitiveTypeDecodeTemplate
func DecodeVnic_Ipv6_Autoconfig(p *Ipv6, vals map[string]cty.Value) {
	p.Autoconfig = ctwhy.ValueAsBool(vals["autoconfig"])
}

//primitiveTypeDecodeTemplate
func DecodeVnic_Ipv6_Dhcp(p *Ipv6, vals map[string]cty.Value) {
	p.Dhcp = ctwhy.ValueAsBool(vals["dhcp"])
}

//primitiveTypeDecodeTemplate
func DecodeVnic_Ipv6_Gw(p *Ipv6, vals map[string]cty.Value) {
	p.Gw = ctwhy.ValueAsString(vals["gw"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeVnic_Ipv6_Addresses(p *Ipv6, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsList(vals["addresses"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.Addresses = goVals
}

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

	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
	"github.com/zclconf/go-cty/cty"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*Vnic)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a Vnic.")
	}
	return EncodeVnic(*r), nil
}

func EncodeVnic(r Vnic) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeVnic_DistributedPortGroup(r.Spec.ForProvider, ctyVal)
	EncodeVnic_DistributedSwitchPort(r.Spec.ForProvider, ctyVal)
	EncodeVnic_Host(r.Spec.ForProvider, ctyVal)
	EncodeVnic_Ipv4(r.Spec.ForProvider.Ipv4, ctyVal)
	EncodeVnic_Ipv6(r.Spec.ForProvider.Ipv6, ctyVal)
	EncodeVnic_Mac(r.Spec.ForProvider, ctyVal)
	EncodeVnic_Mtu(r.Spec.ForProvider, ctyVal)
	EncodeVnic_Netstack(r.Spec.ForProvider, ctyVal)
	EncodeVnic_Portgroup(r.Spec.ForProvider, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeVnic_DistributedPortGroup(p VnicParameters, vals map[string]cty.Value) {
	vals["distributed_port_group"] = cty.StringVal(p.DistributedPortGroup)
}

func EncodeVnic_DistributedSwitchPort(p VnicParameters, vals map[string]cty.Value) {
	vals["distributed_switch_port"] = cty.StringVal(p.DistributedSwitchPort)
}

func EncodeVnic_Host(p VnicParameters, vals map[string]cty.Value) {
	vals["host"] = cty.StringVal(p.Host)
}

func EncodeVnic_Ipv4(p Ipv4, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeVnic_Ipv4_Dhcp(p, ctyVal)
	EncodeVnic_Ipv4_Gw(p, ctyVal)
	EncodeVnic_Ipv4_Ip(p, ctyVal)
	EncodeVnic_Ipv4_Netmask(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	if len(valsForCollection) == 0 {
		vals["ipv4"] = cty.ListValEmpty(cty.EmptyObject)
	} else {
		vals["ipv4"] = cty.ListVal(valsForCollection)
	}
}

func EncodeVnic_Ipv4_Dhcp(p Ipv4, vals map[string]cty.Value) {
	vals["dhcp"] = cty.BoolVal(p.Dhcp)
}

func EncodeVnic_Ipv4_Gw(p Ipv4, vals map[string]cty.Value) {
	vals["gw"] = cty.StringVal(p.Gw)
}

func EncodeVnic_Ipv4_Ip(p Ipv4, vals map[string]cty.Value) {
	vals["ip"] = cty.StringVal(p.Ip)
}

func EncodeVnic_Ipv4_Netmask(p Ipv4, vals map[string]cty.Value) {
	vals["netmask"] = cty.StringVal(p.Netmask)
}

func EncodeVnic_Ipv6(p Ipv6, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeVnic_Ipv6_Addresses(p, ctyVal)
	EncodeVnic_Ipv6_Autoconfig(p, ctyVal)
	EncodeVnic_Ipv6_Dhcp(p, ctyVal)
	EncodeVnic_Ipv6_Gw(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	if len(valsForCollection) == 0 {
		vals["ipv6"] = cty.ListValEmpty(cty.EmptyObject)
	} else {
		vals["ipv6"] = cty.ListVal(valsForCollection)
	}
}

func EncodeVnic_Ipv6_Addresses(p Ipv6, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Addresses {
		colVals = append(colVals, cty.StringVal(value))
	}
	if len(colVals) == 0 {
		vals["addresses"] = cty.ListValEmpty(cty.String)
	} else {
		vals["addresses"] = cty.ListVal(colVals)
	}
}

func EncodeVnic_Ipv6_Autoconfig(p Ipv6, vals map[string]cty.Value) {
	vals["autoconfig"] = cty.BoolVal(p.Autoconfig)
}

func EncodeVnic_Ipv6_Dhcp(p Ipv6, vals map[string]cty.Value) {
	vals["dhcp"] = cty.BoolVal(p.Dhcp)
}

func EncodeVnic_Ipv6_Gw(p Ipv6, vals map[string]cty.Value) {
	vals["gw"] = cty.StringVal(p.Gw)
}

func EncodeVnic_Mac(p VnicParameters, vals map[string]cty.Value) {
	vals["mac"] = cty.StringVal(p.Mac)
}

func EncodeVnic_Mtu(p VnicParameters, vals map[string]cty.Value) {
	vals["mtu"] = cty.NumberIntVal(p.Mtu)
}

func EncodeVnic_Netstack(p VnicParameters, vals map[string]cty.Value) {
	vals["netstack"] = cty.StringVal(p.Netstack)
}

func EncodeVnic_Portgroup(p VnicParameters, vals map[string]cty.Value) {
	vals["portgroup"] = cty.StringVal(p.Portgroup)
}

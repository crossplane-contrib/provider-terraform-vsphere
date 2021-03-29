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
	r, ok := mr.(*HostPortGroup)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeHostPortGroup(r, ctyValue)
}

func DecodeHostPortGroup(prev *HostPortGroup, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeHostPortGroup_ActiveNics(&new.Spec.ForProvider, valMap)
	DecodeHostPortGroup_AllowForgedTransmits(&new.Spec.ForProvider, valMap)
	DecodeHostPortGroup_AllowMacChanges(&new.Spec.ForProvider, valMap)
	DecodeHostPortGroup_AllowPromiscuous(&new.Spec.ForProvider, valMap)
	DecodeHostPortGroup_CheckBeacon(&new.Spec.ForProvider, valMap)
	DecodeHostPortGroup_Failback(&new.Spec.ForProvider, valMap)
	DecodeHostPortGroup_HostSystemId(&new.Spec.ForProvider, valMap)
	DecodeHostPortGroup_Name(&new.Spec.ForProvider, valMap)
	DecodeHostPortGroup_NotifySwitches(&new.Spec.ForProvider, valMap)
	DecodeHostPortGroup_ShapingAverageBandwidth(&new.Spec.ForProvider, valMap)
	DecodeHostPortGroup_ShapingBurstSize(&new.Spec.ForProvider, valMap)
	DecodeHostPortGroup_ShapingEnabled(&new.Spec.ForProvider, valMap)
	DecodeHostPortGroup_ShapingPeakBandwidth(&new.Spec.ForProvider, valMap)
	DecodeHostPortGroup_StandbyNics(&new.Spec.ForProvider, valMap)
	DecodeHostPortGroup_TeamingPolicy(&new.Spec.ForProvider, valMap)
	DecodeHostPortGroup_VirtualSwitchName(&new.Spec.ForProvider, valMap)
	DecodeHostPortGroup_VlanId(&new.Spec.ForProvider, valMap)
	DecodeHostPortGroup_ComputedPolicy(&new.Status.AtProvider, valMap)
	DecodeHostPortGroup_Key(&new.Status.AtProvider, valMap)
	DecodeHostPortGroup_Ports(&new.Status.AtProvider.Ports, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveCollectionTypeDecodeTemplate
func DecodeHostPortGroup_ActiveNics(p *HostPortGroupParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsList(vals["active_nics"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.ActiveNics = goVals
}

//primitiveTypeDecodeTemplate
func DecodeHostPortGroup_AllowForgedTransmits(p *HostPortGroupParameters, vals map[string]cty.Value) {
	p.AllowForgedTransmits = ctwhy.ValueAsBool(vals["allow_forged_transmits"])
}

//primitiveTypeDecodeTemplate
func DecodeHostPortGroup_AllowMacChanges(p *HostPortGroupParameters, vals map[string]cty.Value) {
	p.AllowMacChanges = ctwhy.ValueAsBool(vals["allow_mac_changes"])
}

//primitiveTypeDecodeTemplate
func DecodeHostPortGroup_AllowPromiscuous(p *HostPortGroupParameters, vals map[string]cty.Value) {
	p.AllowPromiscuous = ctwhy.ValueAsBool(vals["allow_promiscuous"])
}

//primitiveTypeDecodeTemplate
func DecodeHostPortGroup_CheckBeacon(p *HostPortGroupParameters, vals map[string]cty.Value) {
	p.CheckBeacon = ctwhy.ValueAsBool(vals["check_beacon"])
}

//primitiveTypeDecodeTemplate
func DecodeHostPortGroup_Failback(p *HostPortGroupParameters, vals map[string]cty.Value) {
	p.Failback = ctwhy.ValueAsBool(vals["failback"])
}

//primitiveTypeDecodeTemplate
func DecodeHostPortGroup_HostSystemId(p *HostPortGroupParameters, vals map[string]cty.Value) {
	p.HostSystemId = ctwhy.ValueAsString(vals["host_system_id"])
}

//primitiveTypeDecodeTemplate
func DecodeHostPortGroup_Name(p *HostPortGroupParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveTypeDecodeTemplate
func DecodeHostPortGroup_NotifySwitches(p *HostPortGroupParameters, vals map[string]cty.Value) {
	p.NotifySwitches = ctwhy.ValueAsBool(vals["notify_switches"])
}

//primitiveTypeDecodeTemplate
func DecodeHostPortGroup_ShapingAverageBandwidth(p *HostPortGroupParameters, vals map[string]cty.Value) {
	p.ShapingAverageBandwidth = ctwhy.ValueAsInt64(vals["shaping_average_bandwidth"])
}

//primitiveTypeDecodeTemplate
func DecodeHostPortGroup_ShapingBurstSize(p *HostPortGroupParameters, vals map[string]cty.Value) {
	p.ShapingBurstSize = ctwhy.ValueAsInt64(vals["shaping_burst_size"])
}

//primitiveTypeDecodeTemplate
func DecodeHostPortGroup_ShapingEnabled(p *HostPortGroupParameters, vals map[string]cty.Value) {
	p.ShapingEnabled = ctwhy.ValueAsBool(vals["shaping_enabled"])
}

//primitiveTypeDecodeTemplate
func DecodeHostPortGroup_ShapingPeakBandwidth(p *HostPortGroupParameters, vals map[string]cty.Value) {
	p.ShapingPeakBandwidth = ctwhy.ValueAsInt64(vals["shaping_peak_bandwidth"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeHostPortGroup_StandbyNics(p *HostPortGroupParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsList(vals["standby_nics"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.StandbyNics = goVals
}

//primitiveTypeDecodeTemplate
func DecodeHostPortGroup_TeamingPolicy(p *HostPortGroupParameters, vals map[string]cty.Value) {
	p.TeamingPolicy = ctwhy.ValueAsString(vals["teaming_policy"])
}

//primitiveTypeDecodeTemplate
func DecodeHostPortGroup_VirtualSwitchName(p *HostPortGroupParameters, vals map[string]cty.Value) {
	p.VirtualSwitchName = ctwhy.ValueAsString(vals["virtual_switch_name"])
}

//primitiveTypeDecodeTemplate
func DecodeHostPortGroup_VlanId(p *HostPortGroupParameters, vals map[string]cty.Value) {
	p.VlanId = ctwhy.ValueAsInt64(vals["vlan_id"])
}

//primitiveMapTypeDecodeTemplate
func DecodeHostPortGroup_ComputedPolicy(p *HostPortGroupObservation, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	if vals["computed_policy"].IsNull() {
		p.ComputedPolicy = nil
		return
	}
	vMap := make(map[string]string)
	v := vals["computed_policy"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.ComputedPolicy = vMap
}

//primitiveTypeDecodeTemplate
func DecodeHostPortGroup_Key(p *HostPortGroupObservation, vals map[string]cty.Value) {
	p.Key = ctwhy.ValueAsString(vals["key"])
}

//containerCollectionTypeDecodeTemplate
func DecodeHostPortGroup_Ports(pp *[]Ports, vals map[string]cty.Value) {
	if vals["ports"].IsNull() {
		*pp = []Ports{}
		return
	}
	rvals := ctwhy.ValueAsList(vals["ports"])
	if len(rvals) == 0 {
		*pp = []Ports{}
		return
	}
	lval := make([]Ports, 0)
	for _, value := range rvals {
		valMap := value.AsValueMap()
		vi := &Ports{}
		DecodeHostPortGroup_Ports_Key(vi, valMap)
		DecodeHostPortGroup_Ports_MacAddresses(vi, valMap)
		DecodeHostPortGroup_Ports_Type(vi, valMap)
		lval = append(lval, *vi)
	}
	*pp = lval
}

//primitiveTypeDecodeTemplate
func DecodeHostPortGroup_Ports_Key(p *Ports, vals map[string]cty.Value) {
	p.Key = ctwhy.ValueAsString(vals["key"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeHostPortGroup_Ports_MacAddresses(p *Ports, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsList(vals["mac_addresses"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.MacAddresses = goVals
}

//primitiveTypeDecodeTemplate
func DecodeHostPortGroup_Ports_Type(p *Ports, vals map[string]cty.Value) {
	p.Type = ctwhy.ValueAsString(vals["type"])
}

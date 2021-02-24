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

	"github.com/zclconf/go-cty/cty"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*HostPortGroup)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a HostPortGroup.")
	}
	return EncodeHostPortGroup(*r), nil
}

func EncodeHostPortGroup(r HostPortGroup) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeHostPortGroup_ActiveNics(r.Spec.ForProvider, ctyVal)
	EncodeHostPortGroup_AllowForgedTransmits(r.Spec.ForProvider, ctyVal)
	EncodeHostPortGroup_AllowMacChanges(r.Spec.ForProvider, ctyVal)
	EncodeHostPortGroup_AllowPromiscuous(r.Spec.ForProvider, ctyVal)
	EncodeHostPortGroup_CheckBeacon(r.Spec.ForProvider, ctyVal)
	EncodeHostPortGroup_Failback(r.Spec.ForProvider, ctyVal)
	EncodeHostPortGroup_HostSystemId(r.Spec.ForProvider, ctyVal)
	EncodeHostPortGroup_Name(r.Spec.ForProvider, ctyVal)
	EncodeHostPortGroup_NotifySwitches(r.Spec.ForProvider, ctyVal)
	EncodeHostPortGroup_ShapingAverageBandwidth(r.Spec.ForProvider, ctyVal)
	EncodeHostPortGroup_ShapingBurstSize(r.Spec.ForProvider, ctyVal)
	EncodeHostPortGroup_ShapingEnabled(r.Spec.ForProvider, ctyVal)
	EncodeHostPortGroup_ShapingPeakBandwidth(r.Spec.ForProvider, ctyVal)
	EncodeHostPortGroup_StandbyNics(r.Spec.ForProvider, ctyVal)
	EncodeHostPortGroup_TeamingPolicy(r.Spec.ForProvider, ctyVal)
	EncodeHostPortGroup_VirtualSwitchName(r.Spec.ForProvider, ctyVal)
	EncodeHostPortGroup_VlanId(r.Spec.ForProvider, ctyVal)
	EncodeHostPortGroup_ComputedPolicy(r.Status.AtProvider, ctyVal)
	EncodeHostPortGroup_Key(r.Status.AtProvider, ctyVal)
	EncodeHostPortGroup_Ports(r.Status.AtProvider.Ports, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeHostPortGroup_ActiveNics(p HostPortGroupParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.ActiveNics {
		colVals = append(colVals, cty.StringVal(value))
	}
	if len(colVals) == 0 {
		vals["active_nics"] = cty.ListValEmpty(cty.String)
	} else {
		vals["active_nics"] = cty.ListVal(colVals)
    }
}

func EncodeHostPortGroup_AllowForgedTransmits(p HostPortGroupParameters, vals map[string]cty.Value) {
	vals["allow_forged_transmits"] = cty.BoolVal(p.AllowForgedTransmits)
}

func EncodeHostPortGroup_AllowMacChanges(p HostPortGroupParameters, vals map[string]cty.Value) {
	vals["allow_mac_changes"] = cty.BoolVal(p.AllowMacChanges)
}

func EncodeHostPortGroup_AllowPromiscuous(p HostPortGroupParameters, vals map[string]cty.Value) {
	vals["allow_promiscuous"] = cty.BoolVal(p.AllowPromiscuous)
}

func EncodeHostPortGroup_CheckBeacon(p HostPortGroupParameters, vals map[string]cty.Value) {
	vals["check_beacon"] = cty.BoolVal(p.CheckBeacon)
}

func EncodeHostPortGroup_Failback(p HostPortGroupParameters, vals map[string]cty.Value) {
	vals["failback"] = cty.BoolVal(p.Failback)
}

func EncodeHostPortGroup_HostSystemId(p HostPortGroupParameters, vals map[string]cty.Value) {
	vals["host_system_id"] = cty.StringVal(p.HostSystemId)
}

func EncodeHostPortGroup_Name(p HostPortGroupParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeHostPortGroup_NotifySwitches(p HostPortGroupParameters, vals map[string]cty.Value) {
	vals["notify_switches"] = cty.BoolVal(p.NotifySwitches)
}

func EncodeHostPortGroup_ShapingAverageBandwidth(p HostPortGroupParameters, vals map[string]cty.Value) {
	vals["shaping_average_bandwidth"] = cty.NumberIntVal(p.ShapingAverageBandwidth)
}

func EncodeHostPortGroup_ShapingBurstSize(p HostPortGroupParameters, vals map[string]cty.Value) {
	vals["shaping_burst_size"] = cty.NumberIntVal(p.ShapingBurstSize)
}

func EncodeHostPortGroup_ShapingEnabled(p HostPortGroupParameters, vals map[string]cty.Value) {
	vals["shaping_enabled"] = cty.BoolVal(p.ShapingEnabled)
}

func EncodeHostPortGroup_ShapingPeakBandwidth(p HostPortGroupParameters, vals map[string]cty.Value) {
	vals["shaping_peak_bandwidth"] = cty.NumberIntVal(p.ShapingPeakBandwidth)
}

func EncodeHostPortGroup_StandbyNics(p HostPortGroupParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.StandbyNics {
		colVals = append(colVals, cty.StringVal(value))
	}
	if len(colVals) == 0 {
		vals["standby_nics"] = cty.ListValEmpty(cty.String)
	} else {
		vals["standby_nics"] = cty.ListVal(colVals)
    }
}

func EncodeHostPortGroup_TeamingPolicy(p HostPortGroupParameters, vals map[string]cty.Value) {
	vals["teaming_policy"] = cty.StringVal(p.TeamingPolicy)
}

func EncodeHostPortGroup_VirtualSwitchName(p HostPortGroupParameters, vals map[string]cty.Value) {
	vals["virtual_switch_name"] = cty.StringVal(p.VirtualSwitchName)
}

func EncodeHostPortGroup_VlanId(p HostPortGroupParameters, vals map[string]cty.Value) {
	vals["vlan_id"] = cty.NumberIntVal(p.VlanId)
}

func EncodeHostPortGroup_ComputedPolicy(p HostPortGroupObservation, vals map[string]cty.Value) {
	if len(p.ComputedPolicy) == 0 {
		vals["computed_policy"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.ComputedPolicy {
		mVals[key] = cty.StringVal(value)
	}
	vals["computed_policy"] = cty.MapVal(mVals)
}

func EncodeHostPortGroup_Key(p HostPortGroupObservation, vals map[string]cty.Value) {
	vals["key"] = cty.StringVal(p.Key)
}

func EncodeHostPortGroup_Ports(p []Ports, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeHostPortGroup_Ports_Key(v, ctyVal)
		EncodeHostPortGroup_Ports_MacAddresses(v, ctyVal)
		EncodeHostPortGroup_Ports_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
    if len(valsForCollection) == 0 {
		vals["ports"] = cty.ListValEmpty(cty.Object(map[string]cty.Type{"key":cty.String, "mac_addresses":cty.List(cty.String), "type":cty.String}))
	} else {
		vals["ports"] = cty.ListVal(valsForCollection)
	}
}

func EncodeHostPortGroup_Ports_Key(p Ports, vals map[string]cty.Value) {
	vals["key"] = cty.StringVal(p.Key)
}

func EncodeHostPortGroup_Ports_MacAddresses(p Ports, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.MacAddresses {
		colVals = append(colVals, cty.StringVal(value))
	}
	if len(colVals) == 0 {
		vals["mac_addresses"] = cty.ListValEmpty(cty.String)
	} else {
		vals["mac_addresses"] = cty.ListVal(colVals)
    }
}

func EncodeHostPortGroup_Ports_Type(p Ports, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}
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
	r, ok := mr.(*HostVirtualSwitch)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a HostVirtualSwitch.")
	}
	return EncodeHostVirtualSwitch(*r), nil
}

func EncodeHostVirtualSwitch(r HostVirtualSwitch) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeHostVirtualSwitch_ActiveNics(r.Spec.ForProvider, ctyVal)
	EncodeHostVirtualSwitch_AllowForgedTransmits(r.Spec.ForProvider, ctyVal)
	EncodeHostVirtualSwitch_AllowMacChanges(r.Spec.ForProvider, ctyVal)
	EncodeHostVirtualSwitch_AllowPromiscuous(r.Spec.ForProvider, ctyVal)
	EncodeHostVirtualSwitch_BeaconInterval(r.Spec.ForProvider, ctyVal)
	EncodeHostVirtualSwitch_CheckBeacon(r.Spec.ForProvider, ctyVal)
	EncodeHostVirtualSwitch_Failback(r.Spec.ForProvider, ctyVal)
	EncodeHostVirtualSwitch_HostSystemId(r.Spec.ForProvider, ctyVal)
	EncodeHostVirtualSwitch_LinkDiscoveryOperation(r.Spec.ForProvider, ctyVal)
	EncodeHostVirtualSwitch_LinkDiscoveryProtocol(r.Spec.ForProvider, ctyVal)
	EncodeHostVirtualSwitch_Mtu(r.Spec.ForProvider, ctyVal)
	EncodeHostVirtualSwitch_Name(r.Spec.ForProvider, ctyVal)
	EncodeHostVirtualSwitch_NetworkAdapters(r.Spec.ForProvider, ctyVal)
	EncodeHostVirtualSwitch_NotifySwitches(r.Spec.ForProvider, ctyVal)
	EncodeHostVirtualSwitch_NumberOfPorts(r.Spec.ForProvider, ctyVal)
	EncodeHostVirtualSwitch_ShapingAverageBandwidth(r.Spec.ForProvider, ctyVal)
	EncodeHostVirtualSwitch_ShapingBurstSize(r.Spec.ForProvider, ctyVal)
	EncodeHostVirtualSwitch_ShapingEnabled(r.Spec.ForProvider, ctyVal)
	EncodeHostVirtualSwitch_ShapingPeakBandwidth(r.Spec.ForProvider, ctyVal)
	EncodeHostVirtualSwitch_StandbyNics(r.Spec.ForProvider, ctyVal)
	EncodeHostVirtualSwitch_TeamingPolicy(r.Spec.ForProvider, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeHostVirtualSwitch_ActiveNics(p HostVirtualSwitchParameters, vals map[string]cty.Value) {
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

func EncodeHostVirtualSwitch_AllowForgedTransmits(p HostVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["allow_forged_transmits"] = cty.BoolVal(p.AllowForgedTransmits)
}

func EncodeHostVirtualSwitch_AllowMacChanges(p HostVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["allow_mac_changes"] = cty.BoolVal(p.AllowMacChanges)
}

func EncodeHostVirtualSwitch_AllowPromiscuous(p HostVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["allow_promiscuous"] = cty.BoolVal(p.AllowPromiscuous)
}

func EncodeHostVirtualSwitch_BeaconInterval(p HostVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["beacon_interval"] = cty.NumberIntVal(p.BeaconInterval)
}

func EncodeHostVirtualSwitch_CheckBeacon(p HostVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["check_beacon"] = cty.BoolVal(p.CheckBeacon)
}

func EncodeHostVirtualSwitch_Failback(p HostVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["failback"] = cty.BoolVal(p.Failback)
}

func EncodeHostVirtualSwitch_HostSystemId(p HostVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["host_system_id"] = cty.StringVal(p.HostSystemId)
}

func EncodeHostVirtualSwitch_LinkDiscoveryOperation(p HostVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["link_discovery_operation"] = cty.StringVal(p.LinkDiscoveryOperation)
}

func EncodeHostVirtualSwitch_LinkDiscoveryProtocol(p HostVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["link_discovery_protocol"] = cty.StringVal(p.LinkDiscoveryProtocol)
}

func EncodeHostVirtualSwitch_Mtu(p HostVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["mtu"] = cty.NumberIntVal(p.Mtu)
}

func EncodeHostVirtualSwitch_Name(p HostVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeHostVirtualSwitch_NetworkAdapters(p HostVirtualSwitchParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.NetworkAdapters {
		colVals = append(colVals, cty.StringVal(value))
	}
	if len(colVals) == 0 {
		vals["network_adapters"] = cty.ListValEmpty(cty.String)
	} else {
		vals["network_adapters"] = cty.ListVal(colVals)
	}
}

func EncodeHostVirtualSwitch_NotifySwitches(p HostVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["notify_switches"] = cty.BoolVal(p.NotifySwitches)
}

func EncodeHostVirtualSwitch_NumberOfPorts(p HostVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["number_of_ports"] = cty.NumberIntVal(p.NumberOfPorts)
}

func EncodeHostVirtualSwitch_ShapingAverageBandwidth(p HostVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["shaping_average_bandwidth"] = cty.NumberIntVal(p.ShapingAverageBandwidth)
}

func EncodeHostVirtualSwitch_ShapingBurstSize(p HostVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["shaping_burst_size"] = cty.NumberIntVal(p.ShapingBurstSize)
}

func EncodeHostVirtualSwitch_ShapingEnabled(p HostVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["shaping_enabled"] = cty.BoolVal(p.ShapingEnabled)
}

func EncodeHostVirtualSwitch_ShapingPeakBandwidth(p HostVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["shaping_peak_bandwidth"] = cty.NumberIntVal(p.ShapingPeakBandwidth)
}

func EncodeHostVirtualSwitch_StandbyNics(p HostVirtualSwitchParameters, vals map[string]cty.Value) {
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

func EncodeHostVirtualSwitch_TeamingPolicy(p HostVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["teaming_policy"] = cty.StringVal(p.TeamingPolicy)
}

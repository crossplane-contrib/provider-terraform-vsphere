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
	ctwhy "github.com/crossplane-contrib/terraform-runtime/pkg/plugin/cty"
)

type ctyDecoder struct{}

func (e *ctyDecoder) DecodeCty(mr resource.Managed, ctyValue cty.Value, schema *providers.Schema) (resource.Managed, error) {
	r, ok := mr.(*HostVirtualSwitch)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeHostVirtualSwitch(r, ctyValue)
}

func DecodeHostVirtualSwitch(prev *HostVirtualSwitch, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeHostVirtualSwitch_ActiveNics(&new.Spec.ForProvider, valMap)
	DecodeHostVirtualSwitch_AllowForgedTransmits(&new.Spec.ForProvider, valMap)
	DecodeHostVirtualSwitch_AllowMacChanges(&new.Spec.ForProvider, valMap)
	DecodeHostVirtualSwitch_AllowPromiscuous(&new.Spec.ForProvider, valMap)
	DecodeHostVirtualSwitch_BeaconInterval(&new.Spec.ForProvider, valMap)
	DecodeHostVirtualSwitch_CheckBeacon(&new.Spec.ForProvider, valMap)
	DecodeHostVirtualSwitch_Failback(&new.Spec.ForProvider, valMap)
	DecodeHostVirtualSwitch_HostSystemId(&new.Spec.ForProvider, valMap)
	DecodeHostVirtualSwitch_LinkDiscoveryOperation(&new.Spec.ForProvider, valMap)
	DecodeHostVirtualSwitch_LinkDiscoveryProtocol(&new.Spec.ForProvider, valMap)
	DecodeHostVirtualSwitch_Mtu(&new.Spec.ForProvider, valMap)
	DecodeHostVirtualSwitch_Name(&new.Spec.ForProvider, valMap)
	DecodeHostVirtualSwitch_NetworkAdapters(&new.Spec.ForProvider, valMap)
	DecodeHostVirtualSwitch_NotifySwitches(&new.Spec.ForProvider, valMap)
	DecodeHostVirtualSwitch_NumberOfPorts(&new.Spec.ForProvider, valMap)
	DecodeHostVirtualSwitch_ShapingAverageBandwidth(&new.Spec.ForProvider, valMap)
	DecodeHostVirtualSwitch_ShapingBurstSize(&new.Spec.ForProvider, valMap)
	DecodeHostVirtualSwitch_ShapingEnabled(&new.Spec.ForProvider, valMap)
	DecodeHostVirtualSwitch_ShapingPeakBandwidth(&new.Spec.ForProvider, valMap)
	DecodeHostVirtualSwitch_StandbyNics(&new.Spec.ForProvider, valMap)
	DecodeHostVirtualSwitch_TeamingPolicy(&new.Spec.ForProvider, valMap)

	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveCollectionTypeDecodeTemplate
func DecodeHostVirtualSwitch_ActiveNics(p *HostVirtualSwitchParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsList(vals["active_nics"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.ActiveNics = goVals
}

//primitiveTypeDecodeTemplate
func DecodeHostVirtualSwitch_AllowForgedTransmits(p *HostVirtualSwitchParameters, vals map[string]cty.Value) {
	p.AllowForgedTransmits = ctwhy.ValueAsBool(vals["allow_forged_transmits"])
}

//primitiveTypeDecodeTemplate
func DecodeHostVirtualSwitch_AllowMacChanges(p *HostVirtualSwitchParameters, vals map[string]cty.Value) {
	p.AllowMacChanges = ctwhy.ValueAsBool(vals["allow_mac_changes"])
}

//primitiveTypeDecodeTemplate
func DecodeHostVirtualSwitch_AllowPromiscuous(p *HostVirtualSwitchParameters, vals map[string]cty.Value) {
	p.AllowPromiscuous = ctwhy.ValueAsBool(vals["allow_promiscuous"])
}

//primitiveTypeDecodeTemplate
func DecodeHostVirtualSwitch_BeaconInterval(p *HostVirtualSwitchParameters, vals map[string]cty.Value) {
	p.BeaconInterval = ctwhy.ValueAsInt64(vals["beacon_interval"])
}

//primitiveTypeDecodeTemplate
func DecodeHostVirtualSwitch_CheckBeacon(p *HostVirtualSwitchParameters, vals map[string]cty.Value) {
	p.CheckBeacon = ctwhy.ValueAsBool(vals["check_beacon"])
}

//primitiveTypeDecodeTemplate
func DecodeHostVirtualSwitch_Failback(p *HostVirtualSwitchParameters, vals map[string]cty.Value) {
	p.Failback = ctwhy.ValueAsBool(vals["failback"])
}

//primitiveTypeDecodeTemplate
func DecodeHostVirtualSwitch_HostSystemId(p *HostVirtualSwitchParameters, vals map[string]cty.Value) {
	p.HostSystemId = ctwhy.ValueAsString(vals["host_system_id"])
}

//primitiveTypeDecodeTemplate
func DecodeHostVirtualSwitch_LinkDiscoveryOperation(p *HostVirtualSwitchParameters, vals map[string]cty.Value) {
	p.LinkDiscoveryOperation = ctwhy.ValueAsString(vals["link_discovery_operation"])
}

//primitiveTypeDecodeTemplate
func DecodeHostVirtualSwitch_LinkDiscoveryProtocol(p *HostVirtualSwitchParameters, vals map[string]cty.Value) {
	p.LinkDiscoveryProtocol = ctwhy.ValueAsString(vals["link_discovery_protocol"])
}

//primitiveTypeDecodeTemplate
func DecodeHostVirtualSwitch_Mtu(p *HostVirtualSwitchParameters, vals map[string]cty.Value) {
	p.Mtu = ctwhy.ValueAsInt64(vals["mtu"])
}

//primitiveTypeDecodeTemplate
func DecodeHostVirtualSwitch_Name(p *HostVirtualSwitchParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeHostVirtualSwitch_NetworkAdapters(p *HostVirtualSwitchParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsList(vals["network_adapters"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.NetworkAdapters = goVals
}

//primitiveTypeDecodeTemplate
func DecodeHostVirtualSwitch_NotifySwitches(p *HostVirtualSwitchParameters, vals map[string]cty.Value) {
	p.NotifySwitches = ctwhy.ValueAsBool(vals["notify_switches"])
}

//primitiveTypeDecodeTemplate
func DecodeHostVirtualSwitch_NumberOfPorts(p *HostVirtualSwitchParameters, vals map[string]cty.Value) {
	p.NumberOfPorts = ctwhy.ValueAsInt64(vals["number_of_ports"])
}

//primitiveTypeDecodeTemplate
func DecodeHostVirtualSwitch_ShapingAverageBandwidth(p *HostVirtualSwitchParameters, vals map[string]cty.Value) {
	p.ShapingAverageBandwidth = ctwhy.ValueAsInt64(vals["shaping_average_bandwidth"])
}

//primitiveTypeDecodeTemplate
func DecodeHostVirtualSwitch_ShapingBurstSize(p *HostVirtualSwitchParameters, vals map[string]cty.Value) {
	p.ShapingBurstSize = ctwhy.ValueAsInt64(vals["shaping_burst_size"])
}

//primitiveTypeDecodeTemplate
func DecodeHostVirtualSwitch_ShapingEnabled(p *HostVirtualSwitchParameters, vals map[string]cty.Value) {
	p.ShapingEnabled = ctwhy.ValueAsBool(vals["shaping_enabled"])
}

//primitiveTypeDecodeTemplate
func DecodeHostVirtualSwitch_ShapingPeakBandwidth(p *HostVirtualSwitchParameters, vals map[string]cty.Value) {
	p.ShapingPeakBandwidth = ctwhy.ValueAsInt64(vals["shaping_peak_bandwidth"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeHostVirtualSwitch_StandbyNics(p *HostVirtualSwitchParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsList(vals["standby_nics"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.StandbyNics = goVals
}

//primitiveTypeDecodeTemplate
func DecodeHostVirtualSwitch_TeamingPolicy(p *HostVirtualSwitchParameters, vals map[string]cty.Value) {
	p.TeamingPolicy = ctwhy.ValueAsString(vals["teaming_policy"])
}
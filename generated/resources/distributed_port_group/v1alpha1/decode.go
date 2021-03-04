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
	r, ok := mr.(*DistributedPortGroup)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeDistributedPortGroup(r, ctyValue)
}

func DecodeDistributedPortGroup(prev *DistributedPortGroup, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeDistributedPortGroup_ActiveUplinks(&new.Spec.ForProvider, valMap)
	DecodeDistributedPortGroup_AllowForgedTransmits(&new.Spec.ForProvider, valMap)
	DecodeDistributedPortGroup_AllowMacChanges(&new.Spec.ForProvider, valMap)
	DecodeDistributedPortGroup_AllowPromiscuous(&new.Spec.ForProvider, valMap)
	DecodeDistributedPortGroup_AutoExpand(&new.Spec.ForProvider, valMap)
	DecodeDistributedPortGroup_BlockAllPorts(&new.Spec.ForProvider, valMap)
	DecodeDistributedPortGroup_BlockOverrideAllowed(&new.Spec.ForProvider, valMap)
	DecodeDistributedPortGroup_CheckBeacon(&new.Spec.ForProvider, valMap)
	DecodeDistributedPortGroup_CustomAttributes(&new.Spec.ForProvider, valMap)
	DecodeDistributedPortGroup_Description(&new.Spec.ForProvider, valMap)
	DecodeDistributedPortGroup_DirectpathGen2Allowed(&new.Spec.ForProvider, valMap)
	DecodeDistributedPortGroup_DistributedVirtualSwitchUuid(&new.Spec.ForProvider, valMap)
	DecodeDistributedPortGroup_EgressShapingAverageBandwidth(&new.Spec.ForProvider, valMap)
	DecodeDistributedPortGroup_EgressShapingBurstSize(&new.Spec.ForProvider, valMap)
	DecodeDistributedPortGroup_EgressShapingEnabled(&new.Spec.ForProvider, valMap)
	DecodeDistributedPortGroup_EgressShapingPeakBandwidth(&new.Spec.ForProvider, valMap)
	DecodeDistributedPortGroup_Failback(&new.Spec.ForProvider, valMap)
	DecodeDistributedPortGroup_IngressShapingAverageBandwidth(&new.Spec.ForProvider, valMap)
	DecodeDistributedPortGroup_IngressShapingBurstSize(&new.Spec.ForProvider, valMap)
	DecodeDistributedPortGroup_IngressShapingEnabled(&new.Spec.ForProvider, valMap)
	DecodeDistributedPortGroup_IngressShapingPeakBandwidth(&new.Spec.ForProvider, valMap)
	DecodeDistributedPortGroup_LacpEnabled(&new.Spec.ForProvider, valMap)
	DecodeDistributedPortGroup_LacpMode(&new.Spec.ForProvider, valMap)
	DecodeDistributedPortGroup_LivePortMovingAllowed(&new.Spec.ForProvider, valMap)
	DecodeDistributedPortGroup_Name(&new.Spec.ForProvider, valMap)
	DecodeDistributedPortGroup_NetflowEnabled(&new.Spec.ForProvider, valMap)
	DecodeDistributedPortGroup_NetflowOverrideAllowed(&new.Spec.ForProvider, valMap)
	DecodeDistributedPortGroup_NetworkResourcePoolKey(&new.Spec.ForProvider, valMap)
	DecodeDistributedPortGroup_NetworkResourcePoolOverrideAllowed(&new.Spec.ForProvider, valMap)
	DecodeDistributedPortGroup_NotifySwitches(&new.Spec.ForProvider, valMap)
	DecodeDistributedPortGroup_NumberOfPorts(&new.Spec.ForProvider, valMap)
	DecodeDistributedPortGroup_PortConfigResetAtDisconnect(&new.Spec.ForProvider, valMap)
	DecodeDistributedPortGroup_PortNameFormat(&new.Spec.ForProvider, valMap)
	DecodeDistributedPortGroup_PortPrivateSecondaryVlanId(&new.Spec.ForProvider, valMap)
	DecodeDistributedPortGroup_SecurityPolicyOverrideAllowed(&new.Spec.ForProvider, valMap)
	DecodeDistributedPortGroup_ShapingOverrideAllowed(&new.Spec.ForProvider, valMap)
	DecodeDistributedPortGroup_StandbyUplinks(&new.Spec.ForProvider, valMap)
	DecodeDistributedPortGroup_Tags(&new.Spec.ForProvider, valMap)
	DecodeDistributedPortGroup_TeamingPolicy(&new.Spec.ForProvider, valMap)
	DecodeDistributedPortGroup_TrafficFilterOverrideAllowed(&new.Spec.ForProvider, valMap)
	DecodeDistributedPortGroup_TxUplink(&new.Spec.ForProvider, valMap)
	DecodeDistributedPortGroup_Type(&new.Spec.ForProvider, valMap)
	DecodeDistributedPortGroup_UplinkTeamingOverrideAllowed(&new.Spec.ForProvider, valMap)
	DecodeDistributedPortGroup_VlanId(&new.Spec.ForProvider, valMap)
	DecodeDistributedPortGroup_VlanOverrideAllowed(&new.Spec.ForProvider, valMap)
	DecodeDistributedPortGroup_VlanRange(&new.Spec.ForProvider.VlanRange, valMap)
	DecodeDistributedPortGroup_ConfigVersion(&new.Status.AtProvider, valMap)
	DecodeDistributedPortGroup_Key(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveCollectionTypeDecodeTemplate
func DecodeDistributedPortGroup_ActiveUplinks(p *DistributedPortGroupParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsList(vals["active_uplinks"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.ActiveUplinks = goVals
}

//primitiveTypeDecodeTemplate
func DecodeDistributedPortGroup_AllowForgedTransmits(p *DistributedPortGroupParameters, vals map[string]cty.Value) {
	p.AllowForgedTransmits = ctwhy.ValueAsBool(vals["allow_forged_transmits"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedPortGroup_AllowMacChanges(p *DistributedPortGroupParameters, vals map[string]cty.Value) {
	p.AllowMacChanges = ctwhy.ValueAsBool(vals["allow_mac_changes"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedPortGroup_AllowPromiscuous(p *DistributedPortGroupParameters, vals map[string]cty.Value) {
	p.AllowPromiscuous = ctwhy.ValueAsBool(vals["allow_promiscuous"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedPortGroup_AutoExpand(p *DistributedPortGroupParameters, vals map[string]cty.Value) {
	p.AutoExpand = ctwhy.ValueAsBool(vals["auto_expand"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedPortGroup_BlockAllPorts(p *DistributedPortGroupParameters, vals map[string]cty.Value) {
	p.BlockAllPorts = ctwhy.ValueAsBool(vals["block_all_ports"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedPortGroup_BlockOverrideAllowed(p *DistributedPortGroupParameters, vals map[string]cty.Value) {
	p.BlockOverrideAllowed = ctwhy.ValueAsBool(vals["block_override_allowed"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedPortGroup_CheckBeacon(p *DistributedPortGroupParameters, vals map[string]cty.Value) {
	p.CheckBeacon = ctwhy.ValueAsBool(vals["check_beacon"])
}

//primitiveMapTypeDecodeTemplate
func DecodeDistributedPortGroup_CustomAttributes(p *DistributedPortGroupParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	if vals["custom_attributes"].IsNull() {
		p.CustomAttributes = nil
		return
	}
	vMap := make(map[string]string)
	v := vals["custom_attributes"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.CustomAttributes = vMap
}

//primitiveTypeDecodeTemplate
func DecodeDistributedPortGroup_Description(p *DistributedPortGroupParameters, vals map[string]cty.Value) {
	p.Description = ctwhy.ValueAsString(vals["description"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedPortGroup_DirectpathGen2Allowed(p *DistributedPortGroupParameters, vals map[string]cty.Value) {
	p.DirectpathGen2Allowed = ctwhy.ValueAsBool(vals["directpath_gen2_allowed"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedPortGroup_DistributedVirtualSwitchUuid(p *DistributedPortGroupParameters, vals map[string]cty.Value) {
	p.DistributedVirtualSwitchUuid = ctwhy.ValueAsString(vals["distributed_virtual_switch_uuid"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedPortGroup_EgressShapingAverageBandwidth(p *DistributedPortGroupParameters, vals map[string]cty.Value) {
	p.EgressShapingAverageBandwidth = ctwhy.ValueAsInt64(vals["egress_shaping_average_bandwidth"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedPortGroup_EgressShapingBurstSize(p *DistributedPortGroupParameters, vals map[string]cty.Value) {
	p.EgressShapingBurstSize = ctwhy.ValueAsInt64(vals["egress_shaping_burst_size"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedPortGroup_EgressShapingEnabled(p *DistributedPortGroupParameters, vals map[string]cty.Value) {
	p.EgressShapingEnabled = ctwhy.ValueAsBool(vals["egress_shaping_enabled"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedPortGroup_EgressShapingPeakBandwidth(p *DistributedPortGroupParameters, vals map[string]cty.Value) {
	p.EgressShapingPeakBandwidth = ctwhy.ValueAsInt64(vals["egress_shaping_peak_bandwidth"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedPortGroup_Failback(p *DistributedPortGroupParameters, vals map[string]cty.Value) {
	p.Failback = ctwhy.ValueAsBool(vals["failback"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedPortGroup_IngressShapingAverageBandwidth(p *DistributedPortGroupParameters, vals map[string]cty.Value) {
	p.IngressShapingAverageBandwidth = ctwhy.ValueAsInt64(vals["ingress_shaping_average_bandwidth"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedPortGroup_IngressShapingBurstSize(p *DistributedPortGroupParameters, vals map[string]cty.Value) {
	p.IngressShapingBurstSize = ctwhy.ValueAsInt64(vals["ingress_shaping_burst_size"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedPortGroup_IngressShapingEnabled(p *DistributedPortGroupParameters, vals map[string]cty.Value) {
	p.IngressShapingEnabled = ctwhy.ValueAsBool(vals["ingress_shaping_enabled"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedPortGroup_IngressShapingPeakBandwidth(p *DistributedPortGroupParameters, vals map[string]cty.Value) {
	p.IngressShapingPeakBandwidth = ctwhy.ValueAsInt64(vals["ingress_shaping_peak_bandwidth"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedPortGroup_LacpEnabled(p *DistributedPortGroupParameters, vals map[string]cty.Value) {
	p.LacpEnabled = ctwhy.ValueAsBool(vals["lacp_enabled"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedPortGroup_LacpMode(p *DistributedPortGroupParameters, vals map[string]cty.Value) {
	p.LacpMode = ctwhy.ValueAsString(vals["lacp_mode"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedPortGroup_LivePortMovingAllowed(p *DistributedPortGroupParameters, vals map[string]cty.Value) {
	p.LivePortMovingAllowed = ctwhy.ValueAsBool(vals["live_port_moving_allowed"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedPortGroup_Name(p *DistributedPortGroupParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedPortGroup_NetflowEnabled(p *DistributedPortGroupParameters, vals map[string]cty.Value) {
	p.NetflowEnabled = ctwhy.ValueAsBool(vals["netflow_enabled"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedPortGroup_NetflowOverrideAllowed(p *DistributedPortGroupParameters, vals map[string]cty.Value) {
	p.NetflowOverrideAllowed = ctwhy.ValueAsBool(vals["netflow_override_allowed"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedPortGroup_NetworkResourcePoolKey(p *DistributedPortGroupParameters, vals map[string]cty.Value) {
	p.NetworkResourcePoolKey = ctwhy.ValueAsString(vals["network_resource_pool_key"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedPortGroup_NetworkResourcePoolOverrideAllowed(p *DistributedPortGroupParameters, vals map[string]cty.Value) {
	p.NetworkResourcePoolOverrideAllowed = ctwhy.ValueAsBool(vals["network_resource_pool_override_allowed"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedPortGroup_NotifySwitches(p *DistributedPortGroupParameters, vals map[string]cty.Value) {
	p.NotifySwitches = ctwhy.ValueAsBool(vals["notify_switches"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedPortGroup_NumberOfPorts(p *DistributedPortGroupParameters, vals map[string]cty.Value) {
	p.NumberOfPorts = ctwhy.ValueAsInt64(vals["number_of_ports"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedPortGroup_PortConfigResetAtDisconnect(p *DistributedPortGroupParameters, vals map[string]cty.Value) {
	p.PortConfigResetAtDisconnect = ctwhy.ValueAsBool(vals["port_config_reset_at_disconnect"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedPortGroup_PortNameFormat(p *DistributedPortGroupParameters, vals map[string]cty.Value) {
	p.PortNameFormat = ctwhy.ValueAsString(vals["port_name_format"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedPortGroup_PortPrivateSecondaryVlanId(p *DistributedPortGroupParameters, vals map[string]cty.Value) {
	p.PortPrivateSecondaryVlanId = ctwhy.ValueAsInt64(vals["port_private_secondary_vlan_id"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedPortGroup_SecurityPolicyOverrideAllowed(p *DistributedPortGroupParameters, vals map[string]cty.Value) {
	p.SecurityPolicyOverrideAllowed = ctwhy.ValueAsBool(vals["security_policy_override_allowed"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedPortGroup_ShapingOverrideAllowed(p *DistributedPortGroupParameters, vals map[string]cty.Value) {
	p.ShapingOverrideAllowed = ctwhy.ValueAsBool(vals["shaping_override_allowed"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeDistributedPortGroup_StandbyUplinks(p *DistributedPortGroupParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsList(vals["standby_uplinks"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.StandbyUplinks = goVals
}

//primitiveCollectionTypeDecodeTemplate
func DecodeDistributedPortGroup_Tags(p *DistributedPortGroupParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["tags"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.Tags = goVals
}

//primitiveTypeDecodeTemplate
func DecodeDistributedPortGroup_TeamingPolicy(p *DistributedPortGroupParameters, vals map[string]cty.Value) {
	p.TeamingPolicy = ctwhy.ValueAsString(vals["teaming_policy"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedPortGroup_TrafficFilterOverrideAllowed(p *DistributedPortGroupParameters, vals map[string]cty.Value) {
	p.TrafficFilterOverrideAllowed = ctwhy.ValueAsBool(vals["traffic_filter_override_allowed"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedPortGroup_TxUplink(p *DistributedPortGroupParameters, vals map[string]cty.Value) {
	p.TxUplink = ctwhy.ValueAsBool(vals["tx_uplink"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedPortGroup_Type(p *DistributedPortGroupParameters, vals map[string]cty.Value) {
	p.Type = ctwhy.ValueAsString(vals["type"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedPortGroup_UplinkTeamingOverrideAllowed(p *DistributedPortGroupParameters, vals map[string]cty.Value) {
	p.UplinkTeamingOverrideAllowed = ctwhy.ValueAsBool(vals["uplink_teaming_override_allowed"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedPortGroup_VlanId(p *DistributedPortGroupParameters, vals map[string]cty.Value) {
	p.VlanId = ctwhy.ValueAsInt64(vals["vlan_id"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedPortGroup_VlanOverrideAllowed(p *DistributedPortGroupParameters, vals map[string]cty.Value) {
	p.VlanOverrideAllowed = ctwhy.ValueAsBool(vals["vlan_override_allowed"])
}

//containerCollectionSingletonTypeDecodeTemplate
func DecodeDistributedPortGroup_VlanRange(p *VlanRange, vals map[string]cty.Value) {
	if vals["vlan_range"].IsNull() {
		*p = VlanRange{}
		return
	}
	rvals := ctwhy.ValueAsSet(vals["vlan_range"])
	if len(rvals) == 0 {
		*p = VlanRange{}
		return
	}
	// this template should be used when single dictionary/object values are nested in sets/lists
	// if rvals turns out to be a list with > 1 elements, something has broken with that heuristic
	valMap := rvals[0].AsValueMap()
	DecodeDistributedPortGroup_VlanRange_MinVlan(p, valMap)
	DecodeDistributedPortGroup_VlanRange_MaxVlan(p, valMap)
}

//primitiveTypeDecodeTemplate
func DecodeDistributedPortGroup_VlanRange_MinVlan(p *VlanRange, vals map[string]cty.Value) {
	p.MinVlan = ctwhy.ValueAsInt64(vals["min_vlan"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedPortGroup_VlanRange_MaxVlan(p *VlanRange, vals map[string]cty.Value) {
	p.MaxVlan = ctwhy.ValueAsInt64(vals["max_vlan"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedPortGroup_ConfigVersion(p *DistributedPortGroupObservation, vals map[string]cty.Value) {
	p.ConfigVersion = ctwhy.ValueAsString(vals["config_version"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedPortGroup_Key(p *DistributedPortGroupObservation, vals map[string]cty.Value) {
	p.Key = ctwhy.ValueAsString(vals["key"])
}

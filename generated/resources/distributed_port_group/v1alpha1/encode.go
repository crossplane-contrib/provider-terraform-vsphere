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
	r, ok := mr.(*DistributedPortGroup)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a DistributedPortGroup.")
	}
	return EncodeDistributedPortGroup(*r), nil
}

func EncodeDistributedPortGroup(r DistributedPortGroup) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDistributedPortGroup_ActiveUplinks(r.Spec.ForProvider, ctyVal)
	EncodeDistributedPortGroup_AllowForgedTransmits(r.Spec.ForProvider, ctyVal)
	EncodeDistributedPortGroup_AllowMacChanges(r.Spec.ForProvider, ctyVal)
	EncodeDistributedPortGroup_AllowPromiscuous(r.Spec.ForProvider, ctyVal)
	EncodeDistributedPortGroup_AutoExpand(r.Spec.ForProvider, ctyVal)
	EncodeDistributedPortGroup_BlockAllPorts(r.Spec.ForProvider, ctyVal)
	EncodeDistributedPortGroup_BlockOverrideAllowed(r.Spec.ForProvider, ctyVal)
	EncodeDistributedPortGroup_CheckBeacon(r.Spec.ForProvider, ctyVal)
	EncodeDistributedPortGroup_CustomAttributes(r.Spec.ForProvider, ctyVal)
	EncodeDistributedPortGroup_Description(r.Spec.ForProvider, ctyVal)
	EncodeDistributedPortGroup_DirectpathGen2Allowed(r.Spec.ForProvider, ctyVal)
	EncodeDistributedPortGroup_DistributedVirtualSwitchUuid(r.Spec.ForProvider, ctyVal)
	EncodeDistributedPortGroup_EgressShapingAverageBandwidth(r.Spec.ForProvider, ctyVal)
	EncodeDistributedPortGroup_EgressShapingBurstSize(r.Spec.ForProvider, ctyVal)
	EncodeDistributedPortGroup_EgressShapingEnabled(r.Spec.ForProvider, ctyVal)
	EncodeDistributedPortGroup_EgressShapingPeakBandwidth(r.Spec.ForProvider, ctyVal)
	EncodeDistributedPortGroup_Failback(r.Spec.ForProvider, ctyVal)
	EncodeDistributedPortGroup_IngressShapingAverageBandwidth(r.Spec.ForProvider, ctyVal)
	EncodeDistributedPortGroup_IngressShapingBurstSize(r.Spec.ForProvider, ctyVal)
	EncodeDistributedPortGroup_IngressShapingEnabled(r.Spec.ForProvider, ctyVal)
	EncodeDistributedPortGroup_IngressShapingPeakBandwidth(r.Spec.ForProvider, ctyVal)
	EncodeDistributedPortGroup_LacpEnabled(r.Spec.ForProvider, ctyVal)
	EncodeDistributedPortGroup_LacpMode(r.Spec.ForProvider, ctyVal)
	EncodeDistributedPortGroup_LivePortMovingAllowed(r.Spec.ForProvider, ctyVal)
	EncodeDistributedPortGroup_Name(r.Spec.ForProvider, ctyVal)
	EncodeDistributedPortGroup_NetflowEnabled(r.Spec.ForProvider, ctyVal)
	EncodeDistributedPortGroup_NetflowOverrideAllowed(r.Spec.ForProvider, ctyVal)
	EncodeDistributedPortGroup_NetworkResourcePoolKey(r.Spec.ForProvider, ctyVal)
	EncodeDistributedPortGroup_NetworkResourcePoolOverrideAllowed(r.Spec.ForProvider, ctyVal)
	EncodeDistributedPortGroup_NotifySwitches(r.Spec.ForProvider, ctyVal)
	EncodeDistributedPortGroup_NumberOfPorts(r.Spec.ForProvider, ctyVal)
	EncodeDistributedPortGroup_PortConfigResetAtDisconnect(r.Spec.ForProvider, ctyVal)
	EncodeDistributedPortGroup_PortNameFormat(r.Spec.ForProvider, ctyVal)
	EncodeDistributedPortGroup_PortPrivateSecondaryVlanId(r.Spec.ForProvider, ctyVal)
	EncodeDistributedPortGroup_SecurityPolicyOverrideAllowed(r.Spec.ForProvider, ctyVal)
	EncodeDistributedPortGroup_ShapingOverrideAllowed(r.Spec.ForProvider, ctyVal)
	EncodeDistributedPortGroup_StandbyUplinks(r.Spec.ForProvider, ctyVal)
	EncodeDistributedPortGroup_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDistributedPortGroup_TeamingPolicy(r.Spec.ForProvider, ctyVal)
	EncodeDistributedPortGroup_TrafficFilterOverrideAllowed(r.Spec.ForProvider, ctyVal)
	EncodeDistributedPortGroup_TxUplink(r.Spec.ForProvider, ctyVal)
	EncodeDistributedPortGroup_Type(r.Spec.ForProvider, ctyVal)
	EncodeDistributedPortGroup_UplinkTeamingOverrideAllowed(r.Spec.ForProvider, ctyVal)
	EncodeDistributedPortGroup_VlanId(r.Spec.ForProvider, ctyVal)
	EncodeDistributedPortGroup_VlanOverrideAllowed(r.Spec.ForProvider, ctyVal)
	EncodeDistributedPortGroup_VlanRange(r.Spec.ForProvider.VlanRange, ctyVal)
	EncodeDistributedPortGroup_ConfigVersion(r.Status.AtProvider, ctyVal)
	EncodeDistributedPortGroup_Key(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeDistributedPortGroup_ActiveUplinks(p DistributedPortGroupParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.ActiveUplinks {
		colVals = append(colVals, cty.StringVal(value))
	}
	if len(colVals) == 0 {
		vals["active_uplinks"] = cty.ListValEmpty(cty.String)
	} else {
		vals["active_uplinks"] = cty.ListVal(colVals)
    }
}

func EncodeDistributedPortGroup_AllowForgedTransmits(p DistributedPortGroupParameters, vals map[string]cty.Value) {
	vals["allow_forged_transmits"] = cty.BoolVal(p.AllowForgedTransmits)
}

func EncodeDistributedPortGroup_AllowMacChanges(p DistributedPortGroupParameters, vals map[string]cty.Value) {
	vals["allow_mac_changes"] = cty.BoolVal(p.AllowMacChanges)
}

func EncodeDistributedPortGroup_AllowPromiscuous(p DistributedPortGroupParameters, vals map[string]cty.Value) {
	vals["allow_promiscuous"] = cty.BoolVal(p.AllowPromiscuous)
}

func EncodeDistributedPortGroup_AutoExpand(p DistributedPortGroupParameters, vals map[string]cty.Value) {
	vals["auto_expand"] = cty.BoolVal(p.AutoExpand)
}

func EncodeDistributedPortGroup_BlockAllPorts(p DistributedPortGroupParameters, vals map[string]cty.Value) {
	vals["block_all_ports"] = cty.BoolVal(p.BlockAllPorts)
}

func EncodeDistributedPortGroup_BlockOverrideAllowed(p DistributedPortGroupParameters, vals map[string]cty.Value) {
	vals["block_override_allowed"] = cty.BoolVal(p.BlockOverrideAllowed)
}

func EncodeDistributedPortGroup_CheckBeacon(p DistributedPortGroupParameters, vals map[string]cty.Value) {
	vals["check_beacon"] = cty.BoolVal(p.CheckBeacon)
}

func EncodeDistributedPortGroup_CustomAttributes(p DistributedPortGroupParameters, vals map[string]cty.Value) {
	if len(p.CustomAttributes) == 0 {
		vals["custom_attributes"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.CustomAttributes {
		mVals[key] = cty.StringVal(value)
	}
	vals["custom_attributes"] = cty.MapVal(mVals)
}

func EncodeDistributedPortGroup_Description(p DistributedPortGroupParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeDistributedPortGroup_DirectpathGen2Allowed(p DistributedPortGroupParameters, vals map[string]cty.Value) {
	vals["directpath_gen2_allowed"] = cty.BoolVal(p.DirectpathGen2Allowed)
}

func EncodeDistributedPortGroup_DistributedVirtualSwitchUuid(p DistributedPortGroupParameters, vals map[string]cty.Value) {
	vals["distributed_virtual_switch_uuid"] = cty.StringVal(p.DistributedVirtualSwitchUuid)
}

func EncodeDistributedPortGroup_EgressShapingAverageBandwidth(p DistributedPortGroupParameters, vals map[string]cty.Value) {
	vals["egress_shaping_average_bandwidth"] = cty.NumberIntVal(p.EgressShapingAverageBandwidth)
}

func EncodeDistributedPortGroup_EgressShapingBurstSize(p DistributedPortGroupParameters, vals map[string]cty.Value) {
	vals["egress_shaping_burst_size"] = cty.NumberIntVal(p.EgressShapingBurstSize)
}

func EncodeDistributedPortGroup_EgressShapingEnabled(p DistributedPortGroupParameters, vals map[string]cty.Value) {
	vals["egress_shaping_enabled"] = cty.BoolVal(p.EgressShapingEnabled)
}

func EncodeDistributedPortGroup_EgressShapingPeakBandwidth(p DistributedPortGroupParameters, vals map[string]cty.Value) {
	vals["egress_shaping_peak_bandwidth"] = cty.NumberIntVal(p.EgressShapingPeakBandwidth)
}

func EncodeDistributedPortGroup_Failback(p DistributedPortGroupParameters, vals map[string]cty.Value) {
	vals["failback"] = cty.BoolVal(p.Failback)
}

func EncodeDistributedPortGroup_IngressShapingAverageBandwidth(p DistributedPortGroupParameters, vals map[string]cty.Value) {
	vals["ingress_shaping_average_bandwidth"] = cty.NumberIntVal(p.IngressShapingAverageBandwidth)
}

func EncodeDistributedPortGroup_IngressShapingBurstSize(p DistributedPortGroupParameters, vals map[string]cty.Value) {
	vals["ingress_shaping_burst_size"] = cty.NumberIntVal(p.IngressShapingBurstSize)
}

func EncodeDistributedPortGroup_IngressShapingEnabled(p DistributedPortGroupParameters, vals map[string]cty.Value) {
	vals["ingress_shaping_enabled"] = cty.BoolVal(p.IngressShapingEnabled)
}

func EncodeDistributedPortGroup_IngressShapingPeakBandwidth(p DistributedPortGroupParameters, vals map[string]cty.Value) {
	vals["ingress_shaping_peak_bandwidth"] = cty.NumberIntVal(p.IngressShapingPeakBandwidth)
}

func EncodeDistributedPortGroup_LacpEnabled(p DistributedPortGroupParameters, vals map[string]cty.Value) {
	vals["lacp_enabled"] = cty.BoolVal(p.LacpEnabled)
}

func EncodeDistributedPortGroup_LacpMode(p DistributedPortGroupParameters, vals map[string]cty.Value) {
	vals["lacp_mode"] = cty.StringVal(p.LacpMode)
}

func EncodeDistributedPortGroup_LivePortMovingAllowed(p DistributedPortGroupParameters, vals map[string]cty.Value) {
	vals["live_port_moving_allowed"] = cty.BoolVal(p.LivePortMovingAllowed)
}

func EncodeDistributedPortGroup_Name(p DistributedPortGroupParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeDistributedPortGroup_NetflowEnabled(p DistributedPortGroupParameters, vals map[string]cty.Value) {
	vals["netflow_enabled"] = cty.BoolVal(p.NetflowEnabled)
}

func EncodeDistributedPortGroup_NetflowOverrideAllowed(p DistributedPortGroupParameters, vals map[string]cty.Value) {
	vals["netflow_override_allowed"] = cty.BoolVal(p.NetflowOverrideAllowed)
}

func EncodeDistributedPortGroup_NetworkResourcePoolKey(p DistributedPortGroupParameters, vals map[string]cty.Value) {
	vals["network_resource_pool_key"] = cty.StringVal(p.NetworkResourcePoolKey)
}

func EncodeDistributedPortGroup_NetworkResourcePoolOverrideAllowed(p DistributedPortGroupParameters, vals map[string]cty.Value) {
	vals["network_resource_pool_override_allowed"] = cty.BoolVal(p.NetworkResourcePoolOverrideAllowed)
}

func EncodeDistributedPortGroup_NotifySwitches(p DistributedPortGroupParameters, vals map[string]cty.Value) {
	vals["notify_switches"] = cty.BoolVal(p.NotifySwitches)
}

func EncodeDistributedPortGroup_NumberOfPorts(p DistributedPortGroupParameters, vals map[string]cty.Value) {
	vals["number_of_ports"] = cty.NumberIntVal(p.NumberOfPorts)
}

func EncodeDistributedPortGroup_PortConfigResetAtDisconnect(p DistributedPortGroupParameters, vals map[string]cty.Value) {
	vals["port_config_reset_at_disconnect"] = cty.BoolVal(p.PortConfigResetAtDisconnect)
}

func EncodeDistributedPortGroup_PortNameFormat(p DistributedPortGroupParameters, vals map[string]cty.Value) {
	vals["port_name_format"] = cty.StringVal(p.PortNameFormat)
}

func EncodeDistributedPortGroup_PortPrivateSecondaryVlanId(p DistributedPortGroupParameters, vals map[string]cty.Value) {
	vals["port_private_secondary_vlan_id"] = cty.NumberIntVal(p.PortPrivateSecondaryVlanId)
}

func EncodeDistributedPortGroup_SecurityPolicyOverrideAllowed(p DistributedPortGroupParameters, vals map[string]cty.Value) {
	vals["security_policy_override_allowed"] = cty.BoolVal(p.SecurityPolicyOverrideAllowed)
}

func EncodeDistributedPortGroup_ShapingOverrideAllowed(p DistributedPortGroupParameters, vals map[string]cty.Value) {
	vals["shaping_override_allowed"] = cty.BoolVal(p.ShapingOverrideAllowed)
}

func EncodeDistributedPortGroup_StandbyUplinks(p DistributedPortGroupParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.StandbyUplinks {
		colVals = append(colVals, cty.StringVal(value))
	}
	if len(colVals) == 0 {
		vals["standby_uplinks"] = cty.ListValEmpty(cty.String)
	} else {
		vals["standby_uplinks"] = cty.ListVal(colVals)
    }
}

func EncodeDistributedPortGroup_Tags(p DistributedPortGroupParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Tags {
		colVals = append(colVals, cty.StringVal(value))
	}
	if len(colVals) == 0 {
		vals["tags"] = cty.SetValEmpty(cty.String)
	} else {
		vals["tags"] = cty.SetVal(colVals)
    }
}

func EncodeDistributedPortGroup_TeamingPolicy(p DistributedPortGroupParameters, vals map[string]cty.Value) {
	vals["teaming_policy"] = cty.StringVal(p.TeamingPolicy)
}

func EncodeDistributedPortGroup_TrafficFilterOverrideAllowed(p DistributedPortGroupParameters, vals map[string]cty.Value) {
	vals["traffic_filter_override_allowed"] = cty.BoolVal(p.TrafficFilterOverrideAllowed)
}

func EncodeDistributedPortGroup_TxUplink(p DistributedPortGroupParameters, vals map[string]cty.Value) {
	vals["tx_uplink"] = cty.BoolVal(p.TxUplink)
}

func EncodeDistributedPortGroup_Type(p DistributedPortGroupParameters, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeDistributedPortGroup_UplinkTeamingOverrideAllowed(p DistributedPortGroupParameters, vals map[string]cty.Value) {
	vals["uplink_teaming_override_allowed"] = cty.BoolVal(p.UplinkTeamingOverrideAllowed)
}

func EncodeDistributedPortGroup_VlanId(p DistributedPortGroupParameters, vals map[string]cty.Value) {
	vals["vlan_id"] = cty.NumberIntVal(p.VlanId)
}

func EncodeDistributedPortGroup_VlanOverrideAllowed(p DistributedPortGroupParameters, vals map[string]cty.Value) {
	vals["vlan_override_allowed"] = cty.BoolVal(p.VlanOverrideAllowed)
}

func EncodeDistributedPortGroup_VlanRange(p VlanRange, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeDistributedPortGroup_VlanRange_MaxVlan(p, ctyVal)
	EncodeDistributedPortGroup_VlanRange_MinVlan(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
    if len(valsForCollection) == 0 {
		vals["vlan_range"] = cty.SetValEmpty(cty.EmptyObject)
    } else {
		vals["vlan_range"] = cty.SetVal(valsForCollection)
    }
}

func EncodeDistributedPortGroup_VlanRange_MaxVlan(p VlanRange, vals map[string]cty.Value) {
	vals["max_vlan"] = cty.NumberIntVal(p.MaxVlan)
}

func EncodeDistributedPortGroup_VlanRange_MinVlan(p VlanRange, vals map[string]cty.Value) {
	vals["min_vlan"] = cty.NumberIntVal(p.MinVlan)
}

func EncodeDistributedPortGroup_ConfigVersion(p DistributedPortGroupObservation, vals map[string]cty.Value) {
	vals["config_version"] = cty.StringVal(p.ConfigVersion)
}

func EncodeDistributedPortGroup_Key(p DistributedPortGroupObservation, vals map[string]cty.Value) {
	vals["key"] = cty.StringVal(p.Key)
}
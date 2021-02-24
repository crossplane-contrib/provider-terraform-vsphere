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
	r, ok := mr.(*DistributedVirtualSwitch)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a DistributedVirtualSwitch.")
	}
	return EncodeDistributedVirtualSwitch(*r), nil
}

func EncodeDistributedVirtualSwitch(r DistributedVirtualSwitch) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDistributedVirtualSwitch_ActiveUplinks(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_AllowForgedTransmits(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_AllowMacChanges(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_AllowPromiscuous(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_BlockAllPorts(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_CheckBeacon(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_ContactDetail(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_ContactName(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_CustomAttributes(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_DatacenterId(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_Description(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_DirectpathGen2Allowed(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_EgressShapingAverageBandwidth(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_EgressShapingBurstSize(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_EgressShapingEnabled(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_EgressShapingPeakBandwidth(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_Failback(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_FaulttoleranceMaximumMbit(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_FaulttoleranceReservationMbit(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_FaulttoleranceShareCount(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_FaulttoleranceShareLevel(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_Folder(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_HbrMaximumMbit(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_HbrReservationMbit(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_HbrShareCount(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_HbrShareLevel(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_IgnoreOtherPvlanMappings(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_IngressShapingAverageBandwidth(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_IngressShapingBurstSize(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_IngressShapingEnabled(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_IngressShapingPeakBandwidth(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_Ipv4Address(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_IscsiMaximumMbit(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_IscsiReservationMbit(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_IscsiShareCount(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_IscsiShareLevel(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_LacpApiVersion(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_LacpEnabled(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_LacpMode(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_LinkDiscoveryOperation(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_LinkDiscoveryProtocol(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_ManagementMaximumMbit(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_ManagementReservationMbit(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_ManagementShareCount(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_ManagementShareLevel(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_MaxMtu(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_MulticastFilteringMode(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_Name(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_NetflowActiveFlowTimeout(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_NetflowCollectorIpAddress(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_NetflowCollectorPort(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_NetflowEnabled(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_NetflowIdleFlowTimeout(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_NetflowInternalFlowsOnly(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_NetflowObservationDomainId(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_NetflowSamplingRate(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_NetworkResourceControlEnabled(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_NetworkResourceControlVersion(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_NfsMaximumMbit(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_NfsReservationMbit(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_NfsShareCount(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_NfsShareLevel(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_NotifySwitches(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_PortPrivateSecondaryVlanId(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_StandbyUplinks(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_TeamingPolicy(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_TxUplink(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_Uplinks(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_VdpMaximumMbit(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_VdpReservationMbit(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_VdpShareCount(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_VdpShareLevel(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_Version(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_VirtualmachineMaximumMbit(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_VirtualmachineReservationMbit(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_VirtualmachineShareCount(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_VirtualmachineShareLevel(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_VlanId(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_VmotionMaximumMbit(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_VmotionReservationMbit(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_VmotionShareCount(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_VmotionShareLevel(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_VsanMaximumMbit(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_VsanReservationMbit(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_VsanShareCount(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_VsanShareLevel(r.Spec.ForProvider, ctyVal)
	EncodeDistributedVirtualSwitch_Host(r.Spec.ForProvider.Host, ctyVal)
	EncodeDistributedVirtualSwitch_PvlanMapping(r.Spec.ForProvider.PvlanMapping, ctyVal)
	EncodeDistributedVirtualSwitch_VlanRange(r.Spec.ForProvider.VlanRange, ctyVal)
	EncodeDistributedVirtualSwitch_ConfigVersion(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeDistributedVirtualSwitch_ActiveUplinks(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
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

func EncodeDistributedVirtualSwitch_AllowForgedTransmits(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["allow_forged_transmits"] = cty.BoolVal(p.AllowForgedTransmits)
}

func EncodeDistributedVirtualSwitch_AllowMacChanges(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["allow_mac_changes"] = cty.BoolVal(p.AllowMacChanges)
}

func EncodeDistributedVirtualSwitch_AllowPromiscuous(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["allow_promiscuous"] = cty.BoolVal(p.AllowPromiscuous)
}

func EncodeDistributedVirtualSwitch_BlockAllPorts(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["block_all_ports"] = cty.BoolVal(p.BlockAllPorts)
}

func EncodeDistributedVirtualSwitch_CheckBeacon(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["check_beacon"] = cty.BoolVal(p.CheckBeacon)
}

func EncodeDistributedVirtualSwitch_ContactDetail(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["contact_detail"] = cty.StringVal(p.ContactDetail)
}

func EncodeDistributedVirtualSwitch_ContactName(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["contact_name"] = cty.StringVal(p.ContactName)
}

func EncodeDistributedVirtualSwitch_CustomAttributes(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
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

func EncodeDistributedVirtualSwitch_DatacenterId(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["datacenter_id"] = cty.StringVal(p.DatacenterId)
}

func EncodeDistributedVirtualSwitch_Description(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeDistributedVirtualSwitch_DirectpathGen2Allowed(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["directpath_gen2_allowed"] = cty.BoolVal(p.DirectpathGen2Allowed)
}

func EncodeDistributedVirtualSwitch_EgressShapingAverageBandwidth(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["egress_shaping_average_bandwidth"] = cty.NumberIntVal(p.EgressShapingAverageBandwidth)
}

func EncodeDistributedVirtualSwitch_EgressShapingBurstSize(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["egress_shaping_burst_size"] = cty.NumberIntVal(p.EgressShapingBurstSize)
}

func EncodeDistributedVirtualSwitch_EgressShapingEnabled(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["egress_shaping_enabled"] = cty.BoolVal(p.EgressShapingEnabled)
}

func EncodeDistributedVirtualSwitch_EgressShapingPeakBandwidth(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["egress_shaping_peak_bandwidth"] = cty.NumberIntVal(p.EgressShapingPeakBandwidth)
}

func EncodeDistributedVirtualSwitch_Failback(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["failback"] = cty.BoolVal(p.Failback)
}

func EncodeDistributedVirtualSwitch_FaulttoleranceMaximumMbit(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["faulttolerance_maximum_mbit"] = cty.NumberIntVal(p.FaulttoleranceMaximumMbit)
}

func EncodeDistributedVirtualSwitch_FaulttoleranceReservationMbit(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["faulttolerance_reservation_mbit"] = cty.NumberIntVal(p.FaulttoleranceReservationMbit)
}

func EncodeDistributedVirtualSwitch_FaulttoleranceShareCount(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["faulttolerance_share_count"] = cty.NumberIntVal(p.FaulttoleranceShareCount)
}

func EncodeDistributedVirtualSwitch_FaulttoleranceShareLevel(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["faulttolerance_share_level"] = cty.StringVal(p.FaulttoleranceShareLevel)
}

func EncodeDistributedVirtualSwitch_Folder(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["folder"] = cty.StringVal(p.Folder)
}

func EncodeDistributedVirtualSwitch_HbrMaximumMbit(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["hbr_maximum_mbit"] = cty.NumberIntVal(p.HbrMaximumMbit)
}

func EncodeDistributedVirtualSwitch_HbrReservationMbit(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["hbr_reservation_mbit"] = cty.NumberIntVal(p.HbrReservationMbit)
}

func EncodeDistributedVirtualSwitch_HbrShareCount(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["hbr_share_count"] = cty.NumberIntVal(p.HbrShareCount)
}

func EncodeDistributedVirtualSwitch_HbrShareLevel(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["hbr_share_level"] = cty.StringVal(p.HbrShareLevel)
}

func EncodeDistributedVirtualSwitch_IgnoreOtherPvlanMappings(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["ignore_other_pvlan_mappings"] = cty.BoolVal(p.IgnoreOtherPvlanMappings)
}

func EncodeDistributedVirtualSwitch_IngressShapingAverageBandwidth(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["ingress_shaping_average_bandwidth"] = cty.NumberIntVal(p.IngressShapingAverageBandwidth)
}

func EncodeDistributedVirtualSwitch_IngressShapingBurstSize(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["ingress_shaping_burst_size"] = cty.NumberIntVal(p.IngressShapingBurstSize)
}

func EncodeDistributedVirtualSwitch_IngressShapingEnabled(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["ingress_shaping_enabled"] = cty.BoolVal(p.IngressShapingEnabled)
}

func EncodeDistributedVirtualSwitch_IngressShapingPeakBandwidth(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["ingress_shaping_peak_bandwidth"] = cty.NumberIntVal(p.IngressShapingPeakBandwidth)
}

func EncodeDistributedVirtualSwitch_Ipv4Address(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["ipv4_address"] = cty.StringVal(p.Ipv4Address)
}

func EncodeDistributedVirtualSwitch_IscsiMaximumMbit(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["iscsi_maximum_mbit"] = cty.NumberIntVal(p.IscsiMaximumMbit)
}

func EncodeDistributedVirtualSwitch_IscsiReservationMbit(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["iscsi_reservation_mbit"] = cty.NumberIntVal(p.IscsiReservationMbit)
}

func EncodeDistributedVirtualSwitch_IscsiShareCount(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["iscsi_share_count"] = cty.NumberIntVal(p.IscsiShareCount)
}

func EncodeDistributedVirtualSwitch_IscsiShareLevel(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["iscsi_share_level"] = cty.StringVal(p.IscsiShareLevel)
}

func EncodeDistributedVirtualSwitch_LacpApiVersion(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["lacp_api_version"] = cty.StringVal(p.LacpApiVersion)
}

func EncodeDistributedVirtualSwitch_LacpEnabled(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["lacp_enabled"] = cty.BoolVal(p.LacpEnabled)
}

func EncodeDistributedVirtualSwitch_LacpMode(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["lacp_mode"] = cty.StringVal(p.LacpMode)
}

func EncodeDistributedVirtualSwitch_LinkDiscoveryOperation(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["link_discovery_operation"] = cty.StringVal(p.LinkDiscoveryOperation)
}

func EncodeDistributedVirtualSwitch_LinkDiscoveryProtocol(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["link_discovery_protocol"] = cty.StringVal(p.LinkDiscoveryProtocol)
}

func EncodeDistributedVirtualSwitch_ManagementMaximumMbit(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["management_maximum_mbit"] = cty.NumberIntVal(p.ManagementMaximumMbit)
}

func EncodeDistributedVirtualSwitch_ManagementReservationMbit(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["management_reservation_mbit"] = cty.NumberIntVal(p.ManagementReservationMbit)
}

func EncodeDistributedVirtualSwitch_ManagementShareCount(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["management_share_count"] = cty.NumberIntVal(p.ManagementShareCount)
}

func EncodeDistributedVirtualSwitch_ManagementShareLevel(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["management_share_level"] = cty.StringVal(p.ManagementShareLevel)
}

func EncodeDistributedVirtualSwitch_MaxMtu(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["max_mtu"] = cty.NumberIntVal(p.MaxMtu)
}

func EncodeDistributedVirtualSwitch_MulticastFilteringMode(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["multicast_filtering_mode"] = cty.StringVal(p.MulticastFilteringMode)
}

func EncodeDistributedVirtualSwitch_Name(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeDistributedVirtualSwitch_NetflowActiveFlowTimeout(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["netflow_active_flow_timeout"] = cty.NumberIntVal(p.NetflowActiveFlowTimeout)
}

func EncodeDistributedVirtualSwitch_NetflowCollectorIpAddress(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["netflow_collector_ip_address"] = cty.StringVal(p.NetflowCollectorIpAddress)
}

func EncodeDistributedVirtualSwitch_NetflowCollectorPort(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["netflow_collector_port"] = cty.NumberIntVal(p.NetflowCollectorPort)
}

func EncodeDistributedVirtualSwitch_NetflowEnabled(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["netflow_enabled"] = cty.BoolVal(p.NetflowEnabled)
}

func EncodeDistributedVirtualSwitch_NetflowIdleFlowTimeout(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["netflow_idle_flow_timeout"] = cty.NumberIntVal(p.NetflowIdleFlowTimeout)
}

func EncodeDistributedVirtualSwitch_NetflowInternalFlowsOnly(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["netflow_internal_flows_only"] = cty.BoolVal(p.NetflowInternalFlowsOnly)
}

func EncodeDistributedVirtualSwitch_NetflowObservationDomainId(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["netflow_observation_domain_id"] = cty.NumberIntVal(p.NetflowObservationDomainId)
}

func EncodeDistributedVirtualSwitch_NetflowSamplingRate(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["netflow_sampling_rate"] = cty.NumberIntVal(p.NetflowSamplingRate)
}

func EncodeDistributedVirtualSwitch_NetworkResourceControlEnabled(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["network_resource_control_enabled"] = cty.BoolVal(p.NetworkResourceControlEnabled)
}

func EncodeDistributedVirtualSwitch_NetworkResourceControlVersion(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["network_resource_control_version"] = cty.StringVal(p.NetworkResourceControlVersion)
}

func EncodeDistributedVirtualSwitch_NfsMaximumMbit(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["nfs_maximum_mbit"] = cty.NumberIntVal(p.NfsMaximumMbit)
}

func EncodeDistributedVirtualSwitch_NfsReservationMbit(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["nfs_reservation_mbit"] = cty.NumberIntVal(p.NfsReservationMbit)
}

func EncodeDistributedVirtualSwitch_NfsShareCount(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["nfs_share_count"] = cty.NumberIntVal(p.NfsShareCount)
}

func EncodeDistributedVirtualSwitch_NfsShareLevel(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["nfs_share_level"] = cty.StringVal(p.NfsShareLevel)
}

func EncodeDistributedVirtualSwitch_NotifySwitches(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["notify_switches"] = cty.BoolVal(p.NotifySwitches)
}

func EncodeDistributedVirtualSwitch_PortPrivateSecondaryVlanId(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["port_private_secondary_vlan_id"] = cty.NumberIntVal(p.PortPrivateSecondaryVlanId)
}

func EncodeDistributedVirtualSwitch_StandbyUplinks(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
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

func EncodeDistributedVirtualSwitch_Tags(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
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

func EncodeDistributedVirtualSwitch_TeamingPolicy(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["teaming_policy"] = cty.StringVal(p.TeamingPolicy)
}

func EncodeDistributedVirtualSwitch_TxUplink(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["tx_uplink"] = cty.BoolVal(p.TxUplink)
}

func EncodeDistributedVirtualSwitch_Uplinks(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Uplinks {
		colVals = append(colVals, cty.StringVal(value))
	}
	if len(colVals) == 0 {
		vals["uplinks"] = cty.ListValEmpty(cty.String)
	} else {
		vals["uplinks"] = cty.ListVal(colVals)
    }
}

func EncodeDistributedVirtualSwitch_VdpMaximumMbit(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["vdp_maximum_mbit"] = cty.NumberIntVal(p.VdpMaximumMbit)
}

func EncodeDistributedVirtualSwitch_VdpReservationMbit(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["vdp_reservation_mbit"] = cty.NumberIntVal(p.VdpReservationMbit)
}

func EncodeDistributedVirtualSwitch_VdpShareCount(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["vdp_share_count"] = cty.NumberIntVal(p.VdpShareCount)
}

func EncodeDistributedVirtualSwitch_VdpShareLevel(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["vdp_share_level"] = cty.StringVal(p.VdpShareLevel)
}

func EncodeDistributedVirtualSwitch_Version(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["version"] = cty.StringVal(p.Version)
}

func EncodeDistributedVirtualSwitch_VirtualmachineMaximumMbit(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["virtualmachine_maximum_mbit"] = cty.NumberIntVal(p.VirtualmachineMaximumMbit)
}

func EncodeDistributedVirtualSwitch_VirtualmachineReservationMbit(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["virtualmachine_reservation_mbit"] = cty.NumberIntVal(p.VirtualmachineReservationMbit)
}

func EncodeDistributedVirtualSwitch_VirtualmachineShareCount(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["virtualmachine_share_count"] = cty.NumberIntVal(p.VirtualmachineShareCount)
}

func EncodeDistributedVirtualSwitch_VirtualmachineShareLevel(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["virtualmachine_share_level"] = cty.StringVal(p.VirtualmachineShareLevel)
}

func EncodeDistributedVirtualSwitch_VlanId(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["vlan_id"] = cty.NumberIntVal(p.VlanId)
}

func EncodeDistributedVirtualSwitch_VmotionMaximumMbit(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["vmotion_maximum_mbit"] = cty.NumberIntVal(p.VmotionMaximumMbit)
}

func EncodeDistributedVirtualSwitch_VmotionReservationMbit(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["vmotion_reservation_mbit"] = cty.NumberIntVal(p.VmotionReservationMbit)
}

func EncodeDistributedVirtualSwitch_VmotionShareCount(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["vmotion_share_count"] = cty.NumberIntVal(p.VmotionShareCount)
}

func EncodeDistributedVirtualSwitch_VmotionShareLevel(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["vmotion_share_level"] = cty.StringVal(p.VmotionShareLevel)
}

func EncodeDistributedVirtualSwitch_VsanMaximumMbit(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["vsan_maximum_mbit"] = cty.NumberIntVal(p.VsanMaximumMbit)
}

func EncodeDistributedVirtualSwitch_VsanReservationMbit(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["vsan_reservation_mbit"] = cty.NumberIntVal(p.VsanReservationMbit)
}

func EncodeDistributedVirtualSwitch_VsanShareCount(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["vsan_share_count"] = cty.NumberIntVal(p.VsanShareCount)
}

func EncodeDistributedVirtualSwitch_VsanShareLevel(p DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	vals["vsan_share_level"] = cty.StringVal(p.VsanShareLevel)
}

func EncodeDistributedVirtualSwitch_Host(p Host, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeDistributedVirtualSwitch_Host_Devices(p, ctyVal)
	EncodeDistributedVirtualSwitch_Host_HostSystemId(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
    if len(valsForCollection) == 0 {
		vals["host"] = cty.SetValEmpty(cty.EmptyObject)
    } else {
		vals["host"] = cty.SetVal(valsForCollection)
    }
}

func EncodeDistributedVirtualSwitch_Host_Devices(p Host, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Devices {
		colVals = append(colVals, cty.StringVal(value))
	}
	if len(colVals) == 0 {
		vals["devices"] = cty.ListValEmpty(cty.String)
	} else {
		vals["devices"] = cty.ListVal(colVals)
    }
}

func EncodeDistributedVirtualSwitch_Host_HostSystemId(p Host, vals map[string]cty.Value) {
	vals["host_system_id"] = cty.StringVal(p.HostSystemId)
}

func EncodeDistributedVirtualSwitch_PvlanMapping(p PvlanMapping, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeDistributedVirtualSwitch_PvlanMapping_PrimaryVlanId(p, ctyVal)
	EncodeDistributedVirtualSwitch_PvlanMapping_PvlanType(p, ctyVal)
	EncodeDistributedVirtualSwitch_PvlanMapping_SecondaryVlanId(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
    if len(valsForCollection) == 0 {
		vals["pvlan_mapping"] = cty.SetValEmpty(cty.EmptyObject)
    } else {
		vals["pvlan_mapping"] = cty.SetVal(valsForCollection)
    }
}

func EncodeDistributedVirtualSwitch_PvlanMapping_PrimaryVlanId(p PvlanMapping, vals map[string]cty.Value) {
	vals["primary_vlan_id"] = cty.NumberIntVal(p.PrimaryVlanId)
}

func EncodeDistributedVirtualSwitch_PvlanMapping_PvlanType(p PvlanMapping, vals map[string]cty.Value) {
	vals["pvlan_type"] = cty.StringVal(p.PvlanType)
}

func EncodeDistributedVirtualSwitch_PvlanMapping_SecondaryVlanId(p PvlanMapping, vals map[string]cty.Value) {
	vals["secondary_vlan_id"] = cty.NumberIntVal(p.SecondaryVlanId)
}

func EncodeDistributedVirtualSwitch_VlanRange(p VlanRange, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeDistributedVirtualSwitch_VlanRange_MaxVlan(p, ctyVal)
	EncodeDistributedVirtualSwitch_VlanRange_MinVlan(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
    if len(valsForCollection) == 0 {
		vals["vlan_range"] = cty.SetValEmpty(cty.EmptyObject)
    } else {
		vals["vlan_range"] = cty.SetVal(valsForCollection)
    }
}

func EncodeDistributedVirtualSwitch_VlanRange_MaxVlan(p VlanRange, vals map[string]cty.Value) {
	vals["max_vlan"] = cty.NumberIntVal(p.MaxVlan)
}

func EncodeDistributedVirtualSwitch_VlanRange_MinVlan(p VlanRange, vals map[string]cty.Value) {
	vals["min_vlan"] = cty.NumberIntVal(p.MinVlan)
}

func EncodeDistributedVirtualSwitch_ConfigVersion(p DistributedVirtualSwitchObservation, vals map[string]cty.Value) {
	vals["config_version"] = cty.StringVal(p.ConfigVersion)
}
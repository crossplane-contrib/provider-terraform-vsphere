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
	r, ok := mr.(*DistributedVirtualSwitch)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeDistributedVirtualSwitch(r, ctyValue)
}

func DecodeDistributedVirtualSwitch(prev *DistributedVirtualSwitch, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeDistributedVirtualSwitch_ActiveUplinks(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_AllowForgedTransmits(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_AllowMacChanges(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_AllowPromiscuous(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_BlockAllPorts(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_CheckBeacon(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_ContactDetail(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_ContactName(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_CustomAttributes(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_DatacenterId(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_Description(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_DirectpathGen2Allowed(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_EgressShapingAverageBandwidth(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_EgressShapingBurstSize(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_EgressShapingEnabled(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_EgressShapingPeakBandwidth(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_Failback(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_FaulttoleranceMaximumMbit(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_FaulttoleranceReservationMbit(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_FaulttoleranceShareCount(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_FaulttoleranceShareLevel(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_Folder(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_HbrMaximumMbit(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_HbrReservationMbit(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_HbrShareCount(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_HbrShareLevel(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_IgnoreOtherPvlanMappings(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_IngressShapingAverageBandwidth(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_IngressShapingBurstSize(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_IngressShapingEnabled(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_IngressShapingPeakBandwidth(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_Ipv4Address(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_IscsiMaximumMbit(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_IscsiReservationMbit(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_IscsiShareCount(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_IscsiShareLevel(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_LacpApiVersion(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_LacpEnabled(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_LacpMode(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_LinkDiscoveryOperation(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_LinkDiscoveryProtocol(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_ManagementMaximumMbit(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_ManagementReservationMbit(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_ManagementShareCount(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_ManagementShareLevel(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_MaxMtu(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_MulticastFilteringMode(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_Name(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_NetflowActiveFlowTimeout(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_NetflowCollectorIpAddress(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_NetflowCollectorPort(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_NetflowEnabled(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_NetflowIdleFlowTimeout(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_NetflowInternalFlowsOnly(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_NetflowObservationDomainId(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_NetflowSamplingRate(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_NetworkResourceControlEnabled(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_NetworkResourceControlVersion(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_NfsMaximumMbit(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_NfsReservationMbit(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_NfsShareCount(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_NfsShareLevel(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_NotifySwitches(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_PortPrivateSecondaryVlanId(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_StandbyUplinks(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_Tags(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_TeamingPolicy(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_TxUplink(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_Uplinks(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_VdpMaximumMbit(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_VdpReservationMbit(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_VdpShareCount(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_VdpShareLevel(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_Version(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_VirtualmachineMaximumMbit(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_VirtualmachineReservationMbit(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_VirtualmachineShareCount(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_VirtualmachineShareLevel(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_VlanId(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_VmotionMaximumMbit(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_VmotionReservationMbit(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_VmotionShareCount(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_VmotionShareLevel(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_VsanMaximumMbit(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_VsanReservationMbit(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_VsanShareCount(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_VsanShareLevel(&new.Spec.ForProvider, valMap)
	DecodeDistributedVirtualSwitch_Host(&new.Spec.ForProvider.Host, valMap)
	DecodeDistributedVirtualSwitch_PvlanMapping(&new.Spec.ForProvider.PvlanMapping, valMap)
	DecodeDistributedVirtualSwitch_VlanRange(&new.Spec.ForProvider.VlanRange, valMap)
	DecodeDistributedVirtualSwitch_ConfigVersion(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveCollectionTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_ActiveUplinks(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsList(vals["active_uplinks"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.ActiveUplinks = goVals
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_AllowForgedTransmits(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.AllowForgedTransmits = ctwhy.ValueAsBool(vals["allow_forged_transmits"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_AllowMacChanges(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.AllowMacChanges = ctwhy.ValueAsBool(vals["allow_mac_changes"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_AllowPromiscuous(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.AllowPromiscuous = ctwhy.ValueAsBool(vals["allow_promiscuous"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_BlockAllPorts(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.BlockAllPorts = ctwhy.ValueAsBool(vals["block_all_ports"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_CheckBeacon(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.CheckBeacon = ctwhy.ValueAsBool(vals["check_beacon"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_ContactDetail(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.ContactDetail = ctwhy.ValueAsString(vals["contact_detail"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_ContactName(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.ContactName = ctwhy.ValueAsString(vals["contact_name"])
}

//primitiveMapTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_CustomAttributes(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
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
func DecodeDistributedVirtualSwitch_DatacenterId(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.DatacenterId = ctwhy.ValueAsString(vals["datacenter_id"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_Description(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.Description = ctwhy.ValueAsString(vals["description"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_DirectpathGen2Allowed(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.DirectpathGen2Allowed = ctwhy.ValueAsBool(vals["directpath_gen2_allowed"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_EgressShapingAverageBandwidth(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.EgressShapingAverageBandwidth = ctwhy.ValueAsInt64(vals["egress_shaping_average_bandwidth"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_EgressShapingBurstSize(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.EgressShapingBurstSize = ctwhy.ValueAsInt64(vals["egress_shaping_burst_size"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_EgressShapingEnabled(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.EgressShapingEnabled = ctwhy.ValueAsBool(vals["egress_shaping_enabled"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_EgressShapingPeakBandwidth(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.EgressShapingPeakBandwidth = ctwhy.ValueAsInt64(vals["egress_shaping_peak_bandwidth"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_Failback(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.Failback = ctwhy.ValueAsBool(vals["failback"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_FaulttoleranceMaximumMbit(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.FaulttoleranceMaximumMbit = ctwhy.ValueAsInt64(vals["faulttolerance_maximum_mbit"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_FaulttoleranceReservationMbit(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.FaulttoleranceReservationMbit = ctwhy.ValueAsInt64(vals["faulttolerance_reservation_mbit"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_FaulttoleranceShareCount(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.FaulttoleranceShareCount = ctwhy.ValueAsInt64(vals["faulttolerance_share_count"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_FaulttoleranceShareLevel(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.FaulttoleranceShareLevel = ctwhy.ValueAsString(vals["faulttolerance_share_level"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_Folder(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.Folder = ctwhy.ValueAsString(vals["folder"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_HbrMaximumMbit(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.HbrMaximumMbit = ctwhy.ValueAsInt64(vals["hbr_maximum_mbit"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_HbrReservationMbit(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.HbrReservationMbit = ctwhy.ValueAsInt64(vals["hbr_reservation_mbit"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_HbrShareCount(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.HbrShareCount = ctwhy.ValueAsInt64(vals["hbr_share_count"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_HbrShareLevel(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.HbrShareLevel = ctwhy.ValueAsString(vals["hbr_share_level"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_IgnoreOtherPvlanMappings(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.IgnoreOtherPvlanMappings = ctwhy.ValueAsBool(vals["ignore_other_pvlan_mappings"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_IngressShapingAverageBandwidth(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.IngressShapingAverageBandwidth = ctwhy.ValueAsInt64(vals["ingress_shaping_average_bandwidth"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_IngressShapingBurstSize(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.IngressShapingBurstSize = ctwhy.ValueAsInt64(vals["ingress_shaping_burst_size"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_IngressShapingEnabled(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.IngressShapingEnabled = ctwhy.ValueAsBool(vals["ingress_shaping_enabled"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_IngressShapingPeakBandwidth(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.IngressShapingPeakBandwidth = ctwhy.ValueAsInt64(vals["ingress_shaping_peak_bandwidth"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_Ipv4Address(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.Ipv4Address = ctwhy.ValueAsString(vals["ipv4_address"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_IscsiMaximumMbit(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.IscsiMaximumMbit = ctwhy.ValueAsInt64(vals["iscsi_maximum_mbit"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_IscsiReservationMbit(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.IscsiReservationMbit = ctwhy.ValueAsInt64(vals["iscsi_reservation_mbit"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_IscsiShareCount(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.IscsiShareCount = ctwhy.ValueAsInt64(vals["iscsi_share_count"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_IscsiShareLevel(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.IscsiShareLevel = ctwhy.ValueAsString(vals["iscsi_share_level"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_LacpApiVersion(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.LacpApiVersion = ctwhy.ValueAsString(vals["lacp_api_version"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_LacpEnabled(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.LacpEnabled = ctwhy.ValueAsBool(vals["lacp_enabled"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_LacpMode(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.LacpMode = ctwhy.ValueAsString(vals["lacp_mode"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_LinkDiscoveryOperation(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.LinkDiscoveryOperation = ctwhy.ValueAsString(vals["link_discovery_operation"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_LinkDiscoveryProtocol(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.LinkDiscoveryProtocol = ctwhy.ValueAsString(vals["link_discovery_protocol"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_ManagementMaximumMbit(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.ManagementMaximumMbit = ctwhy.ValueAsInt64(vals["management_maximum_mbit"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_ManagementReservationMbit(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.ManagementReservationMbit = ctwhy.ValueAsInt64(vals["management_reservation_mbit"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_ManagementShareCount(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.ManagementShareCount = ctwhy.ValueAsInt64(vals["management_share_count"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_ManagementShareLevel(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.ManagementShareLevel = ctwhy.ValueAsString(vals["management_share_level"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_MaxMtu(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.MaxMtu = ctwhy.ValueAsInt64(vals["max_mtu"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_MulticastFilteringMode(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.MulticastFilteringMode = ctwhy.ValueAsString(vals["multicast_filtering_mode"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_Name(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_NetflowActiveFlowTimeout(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.NetflowActiveFlowTimeout = ctwhy.ValueAsInt64(vals["netflow_active_flow_timeout"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_NetflowCollectorIpAddress(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.NetflowCollectorIpAddress = ctwhy.ValueAsString(vals["netflow_collector_ip_address"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_NetflowCollectorPort(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.NetflowCollectorPort = ctwhy.ValueAsInt64(vals["netflow_collector_port"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_NetflowEnabled(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.NetflowEnabled = ctwhy.ValueAsBool(vals["netflow_enabled"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_NetflowIdleFlowTimeout(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.NetflowIdleFlowTimeout = ctwhy.ValueAsInt64(vals["netflow_idle_flow_timeout"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_NetflowInternalFlowsOnly(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.NetflowInternalFlowsOnly = ctwhy.ValueAsBool(vals["netflow_internal_flows_only"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_NetflowObservationDomainId(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.NetflowObservationDomainId = ctwhy.ValueAsInt64(vals["netflow_observation_domain_id"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_NetflowSamplingRate(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.NetflowSamplingRate = ctwhy.ValueAsInt64(vals["netflow_sampling_rate"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_NetworkResourceControlEnabled(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.NetworkResourceControlEnabled = ctwhy.ValueAsBool(vals["network_resource_control_enabled"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_NetworkResourceControlVersion(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.NetworkResourceControlVersion = ctwhy.ValueAsString(vals["network_resource_control_version"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_NfsMaximumMbit(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.NfsMaximumMbit = ctwhy.ValueAsInt64(vals["nfs_maximum_mbit"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_NfsReservationMbit(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.NfsReservationMbit = ctwhy.ValueAsInt64(vals["nfs_reservation_mbit"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_NfsShareCount(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.NfsShareCount = ctwhy.ValueAsInt64(vals["nfs_share_count"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_NfsShareLevel(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.NfsShareLevel = ctwhy.ValueAsString(vals["nfs_share_level"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_NotifySwitches(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.NotifySwitches = ctwhy.ValueAsBool(vals["notify_switches"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_PortPrivateSecondaryVlanId(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.PortPrivateSecondaryVlanId = ctwhy.ValueAsInt64(vals["port_private_secondary_vlan_id"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_StandbyUplinks(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsList(vals["standby_uplinks"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.StandbyUplinks = goVals
}

//primitiveCollectionTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_Tags(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["tags"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.Tags = goVals
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_TeamingPolicy(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.TeamingPolicy = ctwhy.ValueAsString(vals["teaming_policy"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_TxUplink(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.TxUplink = ctwhy.ValueAsBool(vals["tx_uplink"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_Uplinks(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsList(vals["uplinks"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.Uplinks = goVals
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_VdpMaximumMbit(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.VdpMaximumMbit = ctwhy.ValueAsInt64(vals["vdp_maximum_mbit"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_VdpReservationMbit(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.VdpReservationMbit = ctwhy.ValueAsInt64(vals["vdp_reservation_mbit"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_VdpShareCount(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.VdpShareCount = ctwhy.ValueAsInt64(vals["vdp_share_count"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_VdpShareLevel(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.VdpShareLevel = ctwhy.ValueAsString(vals["vdp_share_level"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_Version(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.Version = ctwhy.ValueAsString(vals["version"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_VirtualmachineMaximumMbit(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.VirtualmachineMaximumMbit = ctwhy.ValueAsInt64(vals["virtualmachine_maximum_mbit"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_VirtualmachineReservationMbit(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.VirtualmachineReservationMbit = ctwhy.ValueAsInt64(vals["virtualmachine_reservation_mbit"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_VirtualmachineShareCount(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.VirtualmachineShareCount = ctwhy.ValueAsInt64(vals["virtualmachine_share_count"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_VirtualmachineShareLevel(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.VirtualmachineShareLevel = ctwhy.ValueAsString(vals["virtualmachine_share_level"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_VlanId(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.VlanId = ctwhy.ValueAsInt64(vals["vlan_id"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_VmotionMaximumMbit(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.VmotionMaximumMbit = ctwhy.ValueAsInt64(vals["vmotion_maximum_mbit"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_VmotionReservationMbit(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.VmotionReservationMbit = ctwhy.ValueAsInt64(vals["vmotion_reservation_mbit"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_VmotionShareCount(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.VmotionShareCount = ctwhy.ValueAsInt64(vals["vmotion_share_count"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_VmotionShareLevel(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.VmotionShareLevel = ctwhy.ValueAsString(vals["vmotion_share_level"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_VsanMaximumMbit(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.VsanMaximumMbit = ctwhy.ValueAsInt64(vals["vsan_maximum_mbit"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_VsanReservationMbit(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.VsanReservationMbit = ctwhy.ValueAsInt64(vals["vsan_reservation_mbit"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_VsanShareCount(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.VsanShareCount = ctwhy.ValueAsInt64(vals["vsan_share_count"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_VsanShareLevel(p *DistributedVirtualSwitchParameters, vals map[string]cty.Value) {
	p.VsanShareLevel = ctwhy.ValueAsString(vals["vsan_share_level"])
}

//containerCollectionSingletonTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_Host(p *Host, vals map[string]cty.Value) {
	if vals["host"].IsNull() {
		p = nil
		return
	}
	rvals := ctwhy.ValueAsSet(vals["host"])
	if len(rvals) == 0 {
		p = nil
		return
	}
	// this template should be used when single dictionary/object values are nested in sets/lists
	// if rvals turns out to be a list with > 1 elements, something has broken with that heuristic
	valMap := rvals[0].AsValueMap()
	DecodeDistributedVirtualSwitch_Host_Devices(p, valMap)
	DecodeDistributedVirtualSwitch_Host_HostSystemId(p, valMap)
}

//primitiveCollectionTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_Host_Devices(p *Host, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsList(vals["devices"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.Devices = goVals
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_Host_HostSystemId(p *Host, vals map[string]cty.Value) {
	p.HostSystemId = ctwhy.ValueAsString(vals["host_system_id"])
}

//containerCollectionSingletonTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_PvlanMapping(p *PvlanMapping, vals map[string]cty.Value) {
	if vals["pvlan_mapping"].IsNull() {
		p = nil
		return
	}
	rvals := ctwhy.ValueAsSet(vals["pvlan_mapping"])
	if len(rvals) == 0 {
		p = nil
		return
	}
	// this template should be used when single dictionary/object values are nested in sets/lists
	// if rvals turns out to be a list with > 1 elements, something has broken with that heuristic
	valMap := rvals[0].AsValueMap()
	DecodeDistributedVirtualSwitch_PvlanMapping_PrimaryVlanId(p, valMap)
	DecodeDistributedVirtualSwitch_PvlanMapping_PvlanType(p, valMap)
	DecodeDistributedVirtualSwitch_PvlanMapping_SecondaryVlanId(p, valMap)
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_PvlanMapping_PrimaryVlanId(p *PvlanMapping, vals map[string]cty.Value) {
	p.PrimaryVlanId = ctwhy.ValueAsInt64(vals["primary_vlan_id"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_PvlanMapping_PvlanType(p *PvlanMapping, vals map[string]cty.Value) {
	p.PvlanType = ctwhy.ValueAsString(vals["pvlan_type"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_PvlanMapping_SecondaryVlanId(p *PvlanMapping, vals map[string]cty.Value) {
	p.SecondaryVlanId = ctwhy.ValueAsInt64(vals["secondary_vlan_id"])
}

//containerCollectionSingletonTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_VlanRange(p *VlanRange, vals map[string]cty.Value) {
	if vals["vlan_range"].IsNull() {
		p = nil
		return
	}
	rvals := ctwhy.ValueAsSet(vals["vlan_range"])
	if len(rvals) == 0 {
		p = nil
		return
	}
	// this template should be used when single dictionary/object values are nested in sets/lists
	// if rvals turns out to be a list with > 1 elements, something has broken with that heuristic
	valMap := rvals[0].AsValueMap()
	DecodeDistributedVirtualSwitch_VlanRange_MaxVlan(p, valMap)
	DecodeDistributedVirtualSwitch_VlanRange_MinVlan(p, valMap)
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_VlanRange_MaxVlan(p *VlanRange, vals map[string]cty.Value) {
	p.MaxVlan = ctwhy.ValueAsInt64(vals["max_vlan"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_VlanRange_MinVlan(p *VlanRange, vals map[string]cty.Value) {
	p.MinVlan = ctwhy.ValueAsInt64(vals["min_vlan"])
}

//primitiveTypeDecodeTemplate
func DecodeDistributedVirtualSwitch_ConfigVersion(p *DistributedVirtualSwitchObservation, vals map[string]cty.Value) {
	p.ConfigVersion = ctwhy.ValueAsString(vals["config_version"])
}
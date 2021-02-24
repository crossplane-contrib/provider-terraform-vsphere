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
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane-contrib/terraform-runtime/pkg/plugin"
)

//mergeManagedResourceEntrypointTemplate
type resourceMerger struct{}

func (r *resourceMerger) MergeResources(kube resource.Managed, prov resource.Managed) plugin.MergeDescription {
	k := kube.(*DistributedVirtualSwitch)
	p := prov.(*DistributedVirtualSwitch)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeDistributedVirtualSwitch_ActiveUplinks(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_AllowForgedTransmits(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_AllowMacChanges(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_AllowPromiscuous(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_BlockAllPorts(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_CheckBeacon(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_ContactDetail(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_ContactName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_CustomAttributes(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_DatacenterId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_Description(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_DirectpathGen2Allowed(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_EgressShapingAverageBandwidth(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_EgressShapingBurstSize(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_EgressShapingEnabled(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_EgressShapingPeakBandwidth(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_Failback(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_FaulttoleranceMaximumMbit(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_FaulttoleranceReservationMbit(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_FaulttoleranceShareCount(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_FaulttoleranceShareLevel(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_Folder(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_HbrMaximumMbit(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_HbrReservationMbit(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_HbrShareCount(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_HbrShareLevel(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_IgnoreOtherPvlanMappings(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_IngressShapingAverageBandwidth(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_IngressShapingBurstSize(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_IngressShapingEnabled(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_IngressShapingPeakBandwidth(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_Ipv4Address(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_IscsiMaximumMbit(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_IscsiReservationMbit(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_IscsiShareCount(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_IscsiShareLevel(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_LacpApiVersion(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_LacpEnabled(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_LacpMode(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_LinkDiscoveryOperation(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_LinkDiscoveryProtocol(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_ManagementMaximumMbit(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_ManagementReservationMbit(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_ManagementShareCount(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_ManagementShareLevel(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_MaxMtu(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_MulticastFilteringMode(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_NetflowActiveFlowTimeout(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_NetflowCollectorIpAddress(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_NetflowCollectorPort(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_NetflowEnabled(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_NetflowIdleFlowTimeout(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_NetflowInternalFlowsOnly(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_NetflowObservationDomainId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_NetflowSamplingRate(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_NetworkResourceControlEnabled(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_NetworkResourceControlVersion(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_NfsMaximumMbit(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_NfsReservationMbit(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_NfsShareCount(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_NfsShareLevel(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_NotifySwitches(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_PortPrivateSecondaryVlanId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_StandbyUplinks(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_TeamingPolicy(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_TxUplink(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_Uplinks(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_VdpMaximumMbit(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_VdpReservationMbit(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_VdpShareCount(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_VdpShareLevel(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_Version(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_VirtualmachineMaximumMbit(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_VirtualmachineReservationMbit(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_VirtualmachineShareCount(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_VirtualmachineShareLevel(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_VlanId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_VmotionMaximumMbit(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_VmotionReservationMbit(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_VmotionShareCount(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_VmotionShareLevel(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_VsanMaximumMbit(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_VsanReservationMbit(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_VsanShareCount(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_VsanShareLevel(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_Host(&k.Spec.ForProvider.Host, &p.Spec.ForProvider.Host, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_PvlanMapping(&k.Spec.ForProvider.PvlanMapping, &p.Spec.ForProvider.PvlanMapping, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_VlanRange(&k.Spec.ForProvider.VlanRange, &p.Spec.ForProvider.VlanRange, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_ConfigVersion(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	for key, v := range p.Annotations {
		if k.Annotations[key] != v {
			k.Annotations[key] = v
			md.AnnotationsUpdated = true
		}
	}
	md.AnyFieldUpdated = anyChildUpdated
	return *md
}

//mergePrimitiveContainerTemplateSpec
func MergeDistributedVirtualSwitch_ActiveUplinks(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.ActiveUplinks, p.ActiveUplinks) {
		p.ActiveUplinks = k.ActiveUplinks
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_AllowForgedTransmits(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.AllowForgedTransmits != p.AllowForgedTransmits {
		p.AllowForgedTransmits = k.AllowForgedTransmits
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_AllowMacChanges(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.AllowMacChanges != p.AllowMacChanges {
		p.AllowMacChanges = k.AllowMacChanges
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_AllowPromiscuous(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.AllowPromiscuous != p.AllowPromiscuous {
		p.AllowPromiscuous = k.AllowPromiscuous
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_BlockAllPorts(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.BlockAllPorts != p.BlockAllPorts {
		p.BlockAllPorts = k.BlockAllPorts
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_CheckBeacon(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.CheckBeacon != p.CheckBeacon {
		p.CheckBeacon = k.CheckBeacon
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_ContactDetail(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.ContactDetail != p.ContactDetail {
		p.ContactDetail = k.ContactDetail
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_ContactName(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.ContactName != p.ContactName {
		p.ContactName = k.ContactName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeDistributedVirtualSwitch_CustomAttributes(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.CustomAttributes, p.CustomAttributes) {
		p.CustomAttributes = k.CustomAttributes
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_DatacenterId(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.DatacenterId != p.DatacenterId {
		p.DatacenterId = k.DatacenterId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_Description(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.Description != p.Description {
		p.Description = k.Description
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_DirectpathGen2Allowed(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.DirectpathGen2Allowed != p.DirectpathGen2Allowed {
		p.DirectpathGen2Allowed = k.DirectpathGen2Allowed
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_EgressShapingAverageBandwidth(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.EgressShapingAverageBandwidth != p.EgressShapingAverageBandwidth {
		p.EgressShapingAverageBandwidth = k.EgressShapingAverageBandwidth
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_EgressShapingBurstSize(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.EgressShapingBurstSize != p.EgressShapingBurstSize {
		p.EgressShapingBurstSize = k.EgressShapingBurstSize
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_EgressShapingEnabled(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.EgressShapingEnabled != p.EgressShapingEnabled {
		p.EgressShapingEnabled = k.EgressShapingEnabled
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_EgressShapingPeakBandwidth(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.EgressShapingPeakBandwidth != p.EgressShapingPeakBandwidth {
		p.EgressShapingPeakBandwidth = k.EgressShapingPeakBandwidth
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_Failback(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.Failback != p.Failback {
		p.Failback = k.Failback
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_FaulttoleranceMaximumMbit(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.FaulttoleranceMaximumMbit != p.FaulttoleranceMaximumMbit {
		p.FaulttoleranceMaximumMbit = k.FaulttoleranceMaximumMbit
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_FaulttoleranceReservationMbit(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.FaulttoleranceReservationMbit != p.FaulttoleranceReservationMbit {
		p.FaulttoleranceReservationMbit = k.FaulttoleranceReservationMbit
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_FaulttoleranceShareCount(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.FaulttoleranceShareCount != p.FaulttoleranceShareCount {
		p.FaulttoleranceShareCount = k.FaulttoleranceShareCount
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_FaulttoleranceShareLevel(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.FaulttoleranceShareLevel != p.FaulttoleranceShareLevel {
		p.FaulttoleranceShareLevel = k.FaulttoleranceShareLevel
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_Folder(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.Folder != p.Folder {
		p.Folder = k.Folder
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_HbrMaximumMbit(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.HbrMaximumMbit != p.HbrMaximumMbit {
		p.HbrMaximumMbit = k.HbrMaximumMbit
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_HbrReservationMbit(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.HbrReservationMbit != p.HbrReservationMbit {
		p.HbrReservationMbit = k.HbrReservationMbit
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_HbrShareCount(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.HbrShareCount != p.HbrShareCount {
		p.HbrShareCount = k.HbrShareCount
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_HbrShareLevel(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.HbrShareLevel != p.HbrShareLevel {
		p.HbrShareLevel = k.HbrShareLevel
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_IgnoreOtherPvlanMappings(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.IgnoreOtherPvlanMappings != p.IgnoreOtherPvlanMappings {
		p.IgnoreOtherPvlanMappings = k.IgnoreOtherPvlanMappings
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_IngressShapingAverageBandwidth(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.IngressShapingAverageBandwidth != p.IngressShapingAverageBandwidth {
		p.IngressShapingAverageBandwidth = k.IngressShapingAverageBandwidth
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_IngressShapingBurstSize(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.IngressShapingBurstSize != p.IngressShapingBurstSize {
		p.IngressShapingBurstSize = k.IngressShapingBurstSize
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_IngressShapingEnabled(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.IngressShapingEnabled != p.IngressShapingEnabled {
		p.IngressShapingEnabled = k.IngressShapingEnabled
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_IngressShapingPeakBandwidth(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.IngressShapingPeakBandwidth != p.IngressShapingPeakBandwidth {
		p.IngressShapingPeakBandwidth = k.IngressShapingPeakBandwidth
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_Ipv4Address(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.Ipv4Address != p.Ipv4Address {
		p.Ipv4Address = k.Ipv4Address
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_IscsiMaximumMbit(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.IscsiMaximumMbit != p.IscsiMaximumMbit {
		p.IscsiMaximumMbit = k.IscsiMaximumMbit
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_IscsiReservationMbit(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.IscsiReservationMbit != p.IscsiReservationMbit {
		p.IscsiReservationMbit = k.IscsiReservationMbit
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_IscsiShareCount(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.IscsiShareCount != p.IscsiShareCount {
		p.IscsiShareCount = k.IscsiShareCount
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_IscsiShareLevel(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.IscsiShareLevel != p.IscsiShareLevel {
		p.IscsiShareLevel = k.IscsiShareLevel
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_LacpApiVersion(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.LacpApiVersion != p.LacpApiVersion {
		p.LacpApiVersion = k.LacpApiVersion
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_LacpEnabled(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.LacpEnabled != p.LacpEnabled {
		p.LacpEnabled = k.LacpEnabled
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_LacpMode(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.LacpMode != p.LacpMode {
		p.LacpMode = k.LacpMode
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_LinkDiscoveryOperation(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.LinkDiscoveryOperation != p.LinkDiscoveryOperation {
		p.LinkDiscoveryOperation = k.LinkDiscoveryOperation
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_LinkDiscoveryProtocol(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.LinkDiscoveryProtocol != p.LinkDiscoveryProtocol {
		p.LinkDiscoveryProtocol = k.LinkDiscoveryProtocol
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_ManagementMaximumMbit(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.ManagementMaximumMbit != p.ManagementMaximumMbit {
		p.ManagementMaximumMbit = k.ManagementMaximumMbit
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_ManagementReservationMbit(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.ManagementReservationMbit != p.ManagementReservationMbit {
		p.ManagementReservationMbit = k.ManagementReservationMbit
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_ManagementShareCount(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.ManagementShareCount != p.ManagementShareCount {
		p.ManagementShareCount = k.ManagementShareCount
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_ManagementShareLevel(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.ManagementShareLevel != p.ManagementShareLevel {
		p.ManagementShareLevel = k.ManagementShareLevel
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_MaxMtu(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.MaxMtu != p.MaxMtu {
		p.MaxMtu = k.MaxMtu
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_MulticastFilteringMode(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.MulticastFilteringMode != p.MulticastFilteringMode {
		p.MulticastFilteringMode = k.MulticastFilteringMode
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_Name(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_NetflowActiveFlowTimeout(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.NetflowActiveFlowTimeout != p.NetflowActiveFlowTimeout {
		p.NetflowActiveFlowTimeout = k.NetflowActiveFlowTimeout
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_NetflowCollectorIpAddress(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.NetflowCollectorIpAddress != p.NetflowCollectorIpAddress {
		p.NetflowCollectorIpAddress = k.NetflowCollectorIpAddress
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_NetflowCollectorPort(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.NetflowCollectorPort != p.NetflowCollectorPort {
		p.NetflowCollectorPort = k.NetflowCollectorPort
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_NetflowEnabled(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.NetflowEnabled != p.NetflowEnabled {
		p.NetflowEnabled = k.NetflowEnabled
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_NetflowIdleFlowTimeout(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.NetflowIdleFlowTimeout != p.NetflowIdleFlowTimeout {
		p.NetflowIdleFlowTimeout = k.NetflowIdleFlowTimeout
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_NetflowInternalFlowsOnly(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.NetflowInternalFlowsOnly != p.NetflowInternalFlowsOnly {
		p.NetflowInternalFlowsOnly = k.NetflowInternalFlowsOnly
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_NetflowObservationDomainId(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.NetflowObservationDomainId != p.NetflowObservationDomainId {
		p.NetflowObservationDomainId = k.NetflowObservationDomainId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_NetflowSamplingRate(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.NetflowSamplingRate != p.NetflowSamplingRate {
		p.NetflowSamplingRate = k.NetflowSamplingRate
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_NetworkResourceControlEnabled(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.NetworkResourceControlEnabled != p.NetworkResourceControlEnabled {
		p.NetworkResourceControlEnabled = k.NetworkResourceControlEnabled
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_NetworkResourceControlVersion(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.NetworkResourceControlVersion != p.NetworkResourceControlVersion {
		p.NetworkResourceControlVersion = k.NetworkResourceControlVersion
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_NfsMaximumMbit(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.NfsMaximumMbit != p.NfsMaximumMbit {
		p.NfsMaximumMbit = k.NfsMaximumMbit
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_NfsReservationMbit(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.NfsReservationMbit != p.NfsReservationMbit {
		p.NfsReservationMbit = k.NfsReservationMbit
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_NfsShareCount(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.NfsShareCount != p.NfsShareCount {
		p.NfsShareCount = k.NfsShareCount
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_NfsShareLevel(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.NfsShareLevel != p.NfsShareLevel {
		p.NfsShareLevel = k.NfsShareLevel
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_NotifySwitches(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.NotifySwitches != p.NotifySwitches {
		p.NotifySwitches = k.NotifySwitches
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_PortPrivateSecondaryVlanId(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.PortPrivateSecondaryVlanId != p.PortPrivateSecondaryVlanId {
		p.PortPrivateSecondaryVlanId = k.PortPrivateSecondaryVlanId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeDistributedVirtualSwitch_StandbyUplinks(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.StandbyUplinks, p.StandbyUplinks) {
		p.StandbyUplinks = k.StandbyUplinks
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeDistributedVirtualSwitch_Tags(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_TeamingPolicy(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.TeamingPolicy != p.TeamingPolicy {
		p.TeamingPolicy = k.TeamingPolicy
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_TxUplink(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.TxUplink != p.TxUplink {
		p.TxUplink = k.TxUplink
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeDistributedVirtualSwitch_Uplinks(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.Uplinks, p.Uplinks) {
		p.Uplinks = k.Uplinks
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_VdpMaximumMbit(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.VdpMaximumMbit != p.VdpMaximumMbit {
		p.VdpMaximumMbit = k.VdpMaximumMbit
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_VdpReservationMbit(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.VdpReservationMbit != p.VdpReservationMbit {
		p.VdpReservationMbit = k.VdpReservationMbit
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_VdpShareCount(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.VdpShareCount != p.VdpShareCount {
		p.VdpShareCount = k.VdpShareCount
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_VdpShareLevel(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.VdpShareLevel != p.VdpShareLevel {
		p.VdpShareLevel = k.VdpShareLevel
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_Version(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.Version != p.Version {
		p.Version = k.Version
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_VirtualmachineMaximumMbit(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.VirtualmachineMaximumMbit != p.VirtualmachineMaximumMbit {
		p.VirtualmachineMaximumMbit = k.VirtualmachineMaximumMbit
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_VirtualmachineReservationMbit(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.VirtualmachineReservationMbit != p.VirtualmachineReservationMbit {
		p.VirtualmachineReservationMbit = k.VirtualmachineReservationMbit
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_VirtualmachineShareCount(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.VirtualmachineShareCount != p.VirtualmachineShareCount {
		p.VirtualmachineShareCount = k.VirtualmachineShareCount
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_VirtualmachineShareLevel(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.VirtualmachineShareLevel != p.VirtualmachineShareLevel {
		p.VirtualmachineShareLevel = k.VirtualmachineShareLevel
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_VlanId(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.VlanId != p.VlanId {
		p.VlanId = k.VlanId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_VmotionMaximumMbit(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.VmotionMaximumMbit != p.VmotionMaximumMbit {
		p.VmotionMaximumMbit = k.VmotionMaximumMbit
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_VmotionReservationMbit(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.VmotionReservationMbit != p.VmotionReservationMbit {
		p.VmotionReservationMbit = k.VmotionReservationMbit
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_VmotionShareCount(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.VmotionShareCount != p.VmotionShareCount {
		p.VmotionShareCount = k.VmotionShareCount
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_VmotionShareLevel(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.VmotionShareLevel != p.VmotionShareLevel {
		p.VmotionShareLevel = k.VmotionShareLevel
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_VsanMaximumMbit(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.VsanMaximumMbit != p.VsanMaximumMbit {
		p.VsanMaximumMbit = k.VsanMaximumMbit
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_VsanReservationMbit(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.VsanReservationMbit != p.VsanReservationMbit {
		p.VsanReservationMbit = k.VsanReservationMbit
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_VsanShareCount(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.VsanShareCount != p.VsanShareCount {
		p.VsanShareCount = k.VsanShareCount
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_VsanShareLevel(k *DistributedVirtualSwitchParameters, p *DistributedVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.VsanShareLevel != p.VsanShareLevel {
		p.VsanShareLevel = k.VsanShareLevel
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructTemplateSpec
func MergeDistributedVirtualSwitch_Host(k *Host, p *Host, md *plugin.MergeDescription) bool {
	updated := false
	anyChildUpdated := false
	updated = MergeDistributedVirtualSwitch_Host_Devices(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_Host_HostSystemId(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveContainerTemplateSpec
func MergeDistributedVirtualSwitch_Host_Devices(k *Host, p *Host, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.Devices, p.Devices) {
		p.Devices = k.Devices
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_Host_HostSystemId(k *Host, p *Host, md *plugin.MergeDescription) bool {
	if k.HostSystemId != p.HostSystemId {
		p.HostSystemId = k.HostSystemId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructTemplateSpec
func MergeDistributedVirtualSwitch_PvlanMapping(k *PvlanMapping, p *PvlanMapping, md *plugin.MergeDescription) bool {
	updated := false
	anyChildUpdated := false
	updated = MergeDistributedVirtualSwitch_PvlanMapping_PrimaryVlanId(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_PvlanMapping_PvlanType(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_PvlanMapping_SecondaryVlanId(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_PvlanMapping_PrimaryVlanId(k *PvlanMapping, p *PvlanMapping, md *plugin.MergeDescription) bool {
	if k.PrimaryVlanId != p.PrimaryVlanId {
		p.PrimaryVlanId = k.PrimaryVlanId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_PvlanMapping_PvlanType(k *PvlanMapping, p *PvlanMapping, md *plugin.MergeDescription) bool {
	if k.PvlanType != p.PvlanType {
		p.PvlanType = k.PvlanType
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_PvlanMapping_SecondaryVlanId(k *PvlanMapping, p *PvlanMapping, md *plugin.MergeDescription) bool {
	if k.SecondaryVlanId != p.SecondaryVlanId {
		p.SecondaryVlanId = k.SecondaryVlanId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructTemplateSpec
func MergeDistributedVirtualSwitch_VlanRange(k *VlanRange, p *VlanRange, md *plugin.MergeDescription) bool {
	updated := false
	anyChildUpdated := false
	updated = MergeDistributedVirtualSwitch_VlanRange_MaxVlan(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedVirtualSwitch_VlanRange_MinVlan(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_VlanRange_MaxVlan(k *VlanRange, p *VlanRange, md *plugin.MergeDescription) bool {
	if k.MaxVlan != p.MaxVlan {
		p.MaxVlan = k.MaxVlan
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedVirtualSwitch_VlanRange_MinVlan(k *VlanRange, p *VlanRange, md *plugin.MergeDescription) bool {
	if k.MinVlan != p.MinVlan {
		p.MinVlan = k.MinVlan
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDistributedVirtualSwitch_ConfigVersion(k *DistributedVirtualSwitchObservation, p *DistributedVirtualSwitchObservation, md *plugin.MergeDescription) bool {
	if k.ConfigVersion != p.ConfigVersion {
		k.ConfigVersion = p.ConfigVersion
		md.StatusUpdated = true
		return true
	}
	return false
}
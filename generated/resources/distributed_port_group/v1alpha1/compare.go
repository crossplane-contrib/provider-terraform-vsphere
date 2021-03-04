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
	"github.com/crossplane-contrib/terraform-runtime/pkg/plugin"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
)

//mergeManagedResourceEntrypointTemplate
type resourceMerger struct{}

func (r *resourceMerger) MergeResources(kube resource.Managed, prov resource.Managed) plugin.MergeDescription {
	k := kube.(*DistributedPortGroup)
	p := prov.(*DistributedPortGroup)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeDistributedPortGroup_ActiveUplinks(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedPortGroup_AllowForgedTransmits(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedPortGroup_AllowMacChanges(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedPortGroup_AllowPromiscuous(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedPortGroup_AutoExpand(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedPortGroup_BlockAllPorts(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedPortGroup_BlockOverrideAllowed(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedPortGroup_CheckBeacon(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedPortGroup_CustomAttributes(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedPortGroup_Description(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedPortGroup_DirectpathGen2Allowed(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedPortGroup_DistributedVirtualSwitchUuid(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedPortGroup_EgressShapingAverageBandwidth(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedPortGroup_EgressShapingBurstSize(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedPortGroup_EgressShapingEnabled(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedPortGroup_EgressShapingPeakBandwidth(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedPortGroup_Failback(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedPortGroup_IngressShapingAverageBandwidth(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedPortGroup_IngressShapingBurstSize(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedPortGroup_IngressShapingEnabled(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedPortGroup_IngressShapingPeakBandwidth(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedPortGroup_LacpEnabled(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedPortGroup_LacpMode(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedPortGroup_LivePortMovingAllowed(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedPortGroup_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedPortGroup_NetflowEnabled(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedPortGroup_NetflowOverrideAllowed(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedPortGroup_NetworkResourcePoolKey(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedPortGroup_NetworkResourcePoolOverrideAllowed(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedPortGroup_NotifySwitches(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedPortGroup_NumberOfPorts(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedPortGroup_PortConfigResetAtDisconnect(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedPortGroup_PortNameFormat(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedPortGroup_PortPrivateSecondaryVlanId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedPortGroup_SecurityPolicyOverrideAllowed(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedPortGroup_ShapingOverrideAllowed(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedPortGroup_StandbyUplinks(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedPortGroup_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedPortGroup_TeamingPolicy(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedPortGroup_TrafficFilterOverrideAllowed(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedPortGroup_TxUplink(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedPortGroup_Type(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedPortGroup_UplinkTeamingOverrideAllowed(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedPortGroup_VlanId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedPortGroup_VlanOverrideAllowed(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedPortGroup_VlanRange(&k.Spec.ForProvider.VlanRange, &p.Spec.ForProvider.VlanRange, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedPortGroup_ConfigVersion(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedPortGroup_Key(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeDistributedPortGroup_ActiveUplinks(k *DistributedPortGroupParameters, p *DistributedPortGroupParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.ActiveUplinks, p.ActiveUplinks) {
		p.ActiveUplinks = k.ActiveUplinks
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedPortGroup_AllowForgedTransmits(k *DistributedPortGroupParameters, p *DistributedPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.AllowForgedTransmits != p.AllowForgedTransmits {
		p.AllowForgedTransmits = k.AllowForgedTransmits
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedPortGroup_AllowMacChanges(k *DistributedPortGroupParameters, p *DistributedPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.AllowMacChanges != p.AllowMacChanges {
		p.AllowMacChanges = k.AllowMacChanges
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedPortGroup_AllowPromiscuous(k *DistributedPortGroupParameters, p *DistributedPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.AllowPromiscuous != p.AllowPromiscuous {
		p.AllowPromiscuous = k.AllowPromiscuous
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedPortGroup_AutoExpand(k *DistributedPortGroupParameters, p *DistributedPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.AutoExpand != p.AutoExpand {
		p.AutoExpand = k.AutoExpand
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedPortGroup_BlockAllPorts(k *DistributedPortGroupParameters, p *DistributedPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.BlockAllPorts != p.BlockAllPorts {
		p.BlockAllPorts = k.BlockAllPorts
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedPortGroup_BlockOverrideAllowed(k *DistributedPortGroupParameters, p *DistributedPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.BlockOverrideAllowed != p.BlockOverrideAllowed {
		p.BlockOverrideAllowed = k.BlockOverrideAllowed
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedPortGroup_CheckBeacon(k *DistributedPortGroupParameters, p *DistributedPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.CheckBeacon != p.CheckBeacon {
		p.CheckBeacon = k.CheckBeacon
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeDistributedPortGroup_CustomAttributes(k *DistributedPortGroupParameters, p *DistributedPortGroupParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.CustomAttributes, p.CustomAttributes) {
		p.CustomAttributes = k.CustomAttributes
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedPortGroup_Description(k *DistributedPortGroupParameters, p *DistributedPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.Description != p.Description {
		p.Description = k.Description
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedPortGroup_DirectpathGen2Allowed(k *DistributedPortGroupParameters, p *DistributedPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.DirectpathGen2Allowed != p.DirectpathGen2Allowed {
		p.DirectpathGen2Allowed = k.DirectpathGen2Allowed
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedPortGroup_DistributedVirtualSwitchUuid(k *DistributedPortGroupParameters, p *DistributedPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.DistributedVirtualSwitchUuid != p.DistributedVirtualSwitchUuid {
		p.DistributedVirtualSwitchUuid = k.DistributedVirtualSwitchUuid
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedPortGroup_EgressShapingAverageBandwidth(k *DistributedPortGroupParameters, p *DistributedPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.EgressShapingAverageBandwidth != p.EgressShapingAverageBandwidth {
		p.EgressShapingAverageBandwidth = k.EgressShapingAverageBandwidth
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedPortGroup_EgressShapingBurstSize(k *DistributedPortGroupParameters, p *DistributedPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.EgressShapingBurstSize != p.EgressShapingBurstSize {
		p.EgressShapingBurstSize = k.EgressShapingBurstSize
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedPortGroup_EgressShapingEnabled(k *DistributedPortGroupParameters, p *DistributedPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.EgressShapingEnabled != p.EgressShapingEnabled {
		p.EgressShapingEnabled = k.EgressShapingEnabled
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedPortGroup_EgressShapingPeakBandwidth(k *DistributedPortGroupParameters, p *DistributedPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.EgressShapingPeakBandwidth != p.EgressShapingPeakBandwidth {
		p.EgressShapingPeakBandwidth = k.EgressShapingPeakBandwidth
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedPortGroup_Failback(k *DistributedPortGroupParameters, p *DistributedPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.Failback != p.Failback {
		p.Failback = k.Failback
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedPortGroup_IngressShapingAverageBandwidth(k *DistributedPortGroupParameters, p *DistributedPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.IngressShapingAverageBandwidth != p.IngressShapingAverageBandwidth {
		p.IngressShapingAverageBandwidth = k.IngressShapingAverageBandwidth
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedPortGroup_IngressShapingBurstSize(k *DistributedPortGroupParameters, p *DistributedPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.IngressShapingBurstSize != p.IngressShapingBurstSize {
		p.IngressShapingBurstSize = k.IngressShapingBurstSize
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedPortGroup_IngressShapingEnabled(k *DistributedPortGroupParameters, p *DistributedPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.IngressShapingEnabled != p.IngressShapingEnabled {
		p.IngressShapingEnabled = k.IngressShapingEnabled
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedPortGroup_IngressShapingPeakBandwidth(k *DistributedPortGroupParameters, p *DistributedPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.IngressShapingPeakBandwidth != p.IngressShapingPeakBandwidth {
		p.IngressShapingPeakBandwidth = k.IngressShapingPeakBandwidth
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedPortGroup_LacpEnabled(k *DistributedPortGroupParameters, p *DistributedPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.LacpEnabled != p.LacpEnabled {
		p.LacpEnabled = k.LacpEnabled
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedPortGroup_LacpMode(k *DistributedPortGroupParameters, p *DistributedPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.LacpMode != p.LacpMode {
		p.LacpMode = k.LacpMode
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedPortGroup_LivePortMovingAllowed(k *DistributedPortGroupParameters, p *DistributedPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.LivePortMovingAllowed != p.LivePortMovingAllowed {
		p.LivePortMovingAllowed = k.LivePortMovingAllowed
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedPortGroup_Name(k *DistributedPortGroupParameters, p *DistributedPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedPortGroup_NetflowEnabled(k *DistributedPortGroupParameters, p *DistributedPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.NetflowEnabled != p.NetflowEnabled {
		p.NetflowEnabled = k.NetflowEnabled
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedPortGroup_NetflowOverrideAllowed(k *DistributedPortGroupParameters, p *DistributedPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.NetflowOverrideAllowed != p.NetflowOverrideAllowed {
		p.NetflowOverrideAllowed = k.NetflowOverrideAllowed
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedPortGroup_NetworkResourcePoolKey(k *DistributedPortGroupParameters, p *DistributedPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.NetworkResourcePoolKey != p.NetworkResourcePoolKey {
		p.NetworkResourcePoolKey = k.NetworkResourcePoolKey
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedPortGroup_NetworkResourcePoolOverrideAllowed(k *DistributedPortGroupParameters, p *DistributedPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.NetworkResourcePoolOverrideAllowed != p.NetworkResourcePoolOverrideAllowed {
		p.NetworkResourcePoolOverrideAllowed = k.NetworkResourcePoolOverrideAllowed
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedPortGroup_NotifySwitches(k *DistributedPortGroupParameters, p *DistributedPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.NotifySwitches != p.NotifySwitches {
		p.NotifySwitches = k.NotifySwitches
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedPortGroup_NumberOfPorts(k *DistributedPortGroupParameters, p *DistributedPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.NumberOfPorts != p.NumberOfPorts {
		p.NumberOfPorts = k.NumberOfPorts
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedPortGroup_PortConfigResetAtDisconnect(k *DistributedPortGroupParameters, p *DistributedPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.PortConfigResetAtDisconnect != p.PortConfigResetAtDisconnect {
		p.PortConfigResetAtDisconnect = k.PortConfigResetAtDisconnect
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedPortGroup_PortNameFormat(k *DistributedPortGroupParameters, p *DistributedPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.PortNameFormat != p.PortNameFormat {
		p.PortNameFormat = k.PortNameFormat
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedPortGroup_PortPrivateSecondaryVlanId(k *DistributedPortGroupParameters, p *DistributedPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.PortPrivateSecondaryVlanId != p.PortPrivateSecondaryVlanId {
		p.PortPrivateSecondaryVlanId = k.PortPrivateSecondaryVlanId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedPortGroup_SecurityPolicyOverrideAllowed(k *DistributedPortGroupParameters, p *DistributedPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.SecurityPolicyOverrideAllowed != p.SecurityPolicyOverrideAllowed {
		p.SecurityPolicyOverrideAllowed = k.SecurityPolicyOverrideAllowed
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedPortGroup_ShapingOverrideAllowed(k *DistributedPortGroupParameters, p *DistributedPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.ShapingOverrideAllowed != p.ShapingOverrideAllowed {
		p.ShapingOverrideAllowed = k.ShapingOverrideAllowed
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeDistributedPortGroup_StandbyUplinks(k *DistributedPortGroupParameters, p *DistributedPortGroupParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.StandbyUplinks, p.StandbyUplinks) {
		p.StandbyUplinks = k.StandbyUplinks
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeDistributedPortGroup_Tags(k *DistributedPortGroupParameters, p *DistributedPortGroupParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedPortGroup_TeamingPolicy(k *DistributedPortGroupParameters, p *DistributedPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.TeamingPolicy != p.TeamingPolicy {
		p.TeamingPolicy = k.TeamingPolicy
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedPortGroup_TrafficFilterOverrideAllowed(k *DistributedPortGroupParameters, p *DistributedPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.TrafficFilterOverrideAllowed != p.TrafficFilterOverrideAllowed {
		p.TrafficFilterOverrideAllowed = k.TrafficFilterOverrideAllowed
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedPortGroup_TxUplink(k *DistributedPortGroupParameters, p *DistributedPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.TxUplink != p.TxUplink {
		p.TxUplink = k.TxUplink
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedPortGroup_Type(k *DistributedPortGroupParameters, p *DistributedPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.Type != p.Type {
		p.Type = k.Type
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedPortGroup_UplinkTeamingOverrideAllowed(k *DistributedPortGroupParameters, p *DistributedPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.UplinkTeamingOverrideAllowed != p.UplinkTeamingOverrideAllowed {
		p.UplinkTeamingOverrideAllowed = k.UplinkTeamingOverrideAllowed
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedPortGroup_VlanId(k *DistributedPortGroupParameters, p *DistributedPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.VlanId != p.VlanId {
		p.VlanId = k.VlanId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedPortGroup_VlanOverrideAllowed(k *DistributedPortGroupParameters, p *DistributedPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.VlanOverrideAllowed != p.VlanOverrideAllowed {
		p.VlanOverrideAllowed = k.VlanOverrideAllowed
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructTemplateSpec
func MergeDistributedPortGroup_VlanRange(k *VlanRange, p *VlanRange, md *plugin.MergeDescription) bool {
	updated := false
	anyChildUpdated := false
	updated = MergeDistributedPortGroup_VlanRange_MinVlan(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDistributedPortGroup_VlanRange_MaxVlan(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveTemplateSpec
func MergeDistributedPortGroup_VlanRange_MinVlan(k *VlanRange, p *VlanRange, md *plugin.MergeDescription) bool {
	if k.MinVlan != p.MinVlan {
		p.MinVlan = k.MinVlan
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDistributedPortGroup_VlanRange_MaxVlan(k *VlanRange, p *VlanRange, md *plugin.MergeDescription) bool {
	if k.MaxVlan != p.MaxVlan {
		p.MaxVlan = k.MaxVlan
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDistributedPortGroup_ConfigVersion(k *DistributedPortGroupObservation, p *DistributedPortGroupObservation, md *plugin.MergeDescription) bool {
	if k.ConfigVersion != p.ConfigVersion {
		k.ConfigVersion = p.ConfigVersion
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeDistributedPortGroup_Key(k *DistributedPortGroupObservation, p *DistributedPortGroupObservation, md *plugin.MergeDescription) bool {
	if k.Key != p.Key {
		k.Key = p.Key
		md.StatusUpdated = true
		return true
	}
	return false
}

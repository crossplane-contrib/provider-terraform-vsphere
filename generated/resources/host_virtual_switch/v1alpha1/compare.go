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
	k := kube.(*HostVirtualSwitch)
	p := prov.(*HostVirtualSwitch)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeHostVirtualSwitch_AllowForgedTransmits(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHostVirtualSwitch_AllowMacChanges(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHostVirtualSwitch_AllowPromiscuous(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHostVirtualSwitch_BeaconInterval(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHostVirtualSwitch_Mtu(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHostVirtualSwitch_ShapingAverageBandwidth(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHostVirtualSwitch_ActiveNics(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHostVirtualSwitch_NotifySwitches(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHostVirtualSwitch_NumberOfPorts(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHostVirtualSwitch_TeamingPolicy(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHostVirtualSwitch_CheckBeacon(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHostVirtualSwitch_LinkDiscoveryOperation(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHostVirtualSwitch_NetworkAdapters(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHostVirtualSwitch_ShapingPeakBandwidth(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHostVirtualSwitch_Failback(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHostVirtualSwitch_LinkDiscoveryProtocol(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHostVirtualSwitch_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHostVirtualSwitch_ShapingBurstSize(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHostVirtualSwitch_ShapingEnabled(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHostVirtualSwitch_StandbyNics(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHostVirtualSwitch_HostSystemId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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

//mergePrimitiveTemplateSpec
func MergeHostVirtualSwitch_AllowForgedTransmits(k *HostVirtualSwitchParameters, p *HostVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.AllowForgedTransmits != p.AllowForgedTransmits {
		p.AllowForgedTransmits = k.AllowForgedTransmits
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHostVirtualSwitch_AllowMacChanges(k *HostVirtualSwitchParameters, p *HostVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.AllowMacChanges != p.AllowMacChanges {
		p.AllowMacChanges = k.AllowMacChanges
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHostVirtualSwitch_AllowPromiscuous(k *HostVirtualSwitchParameters, p *HostVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.AllowPromiscuous != p.AllowPromiscuous {
		p.AllowPromiscuous = k.AllowPromiscuous
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHostVirtualSwitch_BeaconInterval(k *HostVirtualSwitchParameters, p *HostVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.BeaconInterval != p.BeaconInterval {
		p.BeaconInterval = k.BeaconInterval
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHostVirtualSwitch_Mtu(k *HostVirtualSwitchParameters, p *HostVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.Mtu != p.Mtu {
		p.Mtu = k.Mtu
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHostVirtualSwitch_ShapingAverageBandwidth(k *HostVirtualSwitchParameters, p *HostVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.ShapingAverageBandwidth != p.ShapingAverageBandwidth {
		p.ShapingAverageBandwidth = k.ShapingAverageBandwidth
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeHostVirtualSwitch_ActiveNics(k *HostVirtualSwitchParameters, p *HostVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.ActiveNics, p.ActiveNics) {
		p.ActiveNics = k.ActiveNics
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHostVirtualSwitch_NotifySwitches(k *HostVirtualSwitchParameters, p *HostVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.NotifySwitches != p.NotifySwitches {
		p.NotifySwitches = k.NotifySwitches
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHostVirtualSwitch_NumberOfPorts(k *HostVirtualSwitchParameters, p *HostVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.NumberOfPorts != p.NumberOfPorts {
		p.NumberOfPorts = k.NumberOfPorts
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHostVirtualSwitch_TeamingPolicy(k *HostVirtualSwitchParameters, p *HostVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.TeamingPolicy != p.TeamingPolicy {
		p.TeamingPolicy = k.TeamingPolicy
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHostVirtualSwitch_CheckBeacon(k *HostVirtualSwitchParameters, p *HostVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.CheckBeacon != p.CheckBeacon {
		p.CheckBeacon = k.CheckBeacon
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHostVirtualSwitch_LinkDiscoveryOperation(k *HostVirtualSwitchParameters, p *HostVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.LinkDiscoveryOperation != p.LinkDiscoveryOperation {
		p.LinkDiscoveryOperation = k.LinkDiscoveryOperation
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeHostVirtualSwitch_NetworkAdapters(k *HostVirtualSwitchParameters, p *HostVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.NetworkAdapters, p.NetworkAdapters) {
		p.NetworkAdapters = k.NetworkAdapters
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHostVirtualSwitch_ShapingPeakBandwidth(k *HostVirtualSwitchParameters, p *HostVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.ShapingPeakBandwidth != p.ShapingPeakBandwidth {
		p.ShapingPeakBandwidth = k.ShapingPeakBandwidth
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHostVirtualSwitch_Failback(k *HostVirtualSwitchParameters, p *HostVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.Failback != p.Failback {
		p.Failback = k.Failback
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHostVirtualSwitch_LinkDiscoveryProtocol(k *HostVirtualSwitchParameters, p *HostVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.LinkDiscoveryProtocol != p.LinkDiscoveryProtocol {
		p.LinkDiscoveryProtocol = k.LinkDiscoveryProtocol
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHostVirtualSwitch_Name(k *HostVirtualSwitchParameters, p *HostVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHostVirtualSwitch_ShapingBurstSize(k *HostVirtualSwitchParameters, p *HostVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.ShapingBurstSize != p.ShapingBurstSize {
		p.ShapingBurstSize = k.ShapingBurstSize
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHostVirtualSwitch_ShapingEnabled(k *HostVirtualSwitchParameters, p *HostVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.ShapingEnabled != p.ShapingEnabled {
		p.ShapingEnabled = k.ShapingEnabled
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeHostVirtualSwitch_StandbyNics(k *HostVirtualSwitchParameters, p *HostVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.StandbyNics, p.StandbyNics) {
		p.StandbyNics = k.StandbyNics
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHostVirtualSwitch_HostSystemId(k *HostVirtualSwitchParameters, p *HostVirtualSwitchParameters, md *plugin.MergeDescription) bool {
	if k.HostSystemId != p.HostSystemId {
		p.HostSystemId = k.HostSystemId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}
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
	k := kube.(*HostPortGroup)
	p := prov.(*HostPortGroup)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeHostPortGroup_ActiveNics(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHostPortGroup_AllowForgedTransmits(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHostPortGroup_AllowMacChanges(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHostPortGroup_AllowPromiscuous(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHostPortGroup_CheckBeacon(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHostPortGroup_Failback(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHostPortGroup_HostSystemId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHostPortGroup_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHostPortGroup_NotifySwitches(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHostPortGroup_ShapingAverageBandwidth(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHostPortGroup_ShapingBurstSize(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHostPortGroup_ShapingEnabled(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHostPortGroup_ShapingPeakBandwidth(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHostPortGroup_StandbyNics(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHostPortGroup_TeamingPolicy(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHostPortGroup_VirtualSwitchName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHostPortGroup_VlanId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHostPortGroup_ComputedPolicy(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHostPortGroup_Key(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHostPortGroup_Ports(&k.Status.AtProvider.Ports, &p.Status.AtProvider.Ports, md)
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
func MergeHostPortGroup_ActiveNics(k *HostPortGroupParameters, p *HostPortGroupParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.ActiveNics, p.ActiveNics) {
		p.ActiveNics = k.ActiveNics
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHostPortGroup_AllowForgedTransmits(k *HostPortGroupParameters, p *HostPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.AllowForgedTransmits != p.AllowForgedTransmits {
		p.AllowForgedTransmits = k.AllowForgedTransmits
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHostPortGroup_AllowMacChanges(k *HostPortGroupParameters, p *HostPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.AllowMacChanges != p.AllowMacChanges {
		p.AllowMacChanges = k.AllowMacChanges
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHostPortGroup_AllowPromiscuous(k *HostPortGroupParameters, p *HostPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.AllowPromiscuous != p.AllowPromiscuous {
		p.AllowPromiscuous = k.AllowPromiscuous
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHostPortGroup_CheckBeacon(k *HostPortGroupParameters, p *HostPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.CheckBeacon != p.CheckBeacon {
		p.CheckBeacon = k.CheckBeacon
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHostPortGroup_Failback(k *HostPortGroupParameters, p *HostPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.Failback != p.Failback {
		p.Failback = k.Failback
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHostPortGroup_HostSystemId(k *HostPortGroupParameters, p *HostPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.HostSystemId != p.HostSystemId {
		p.HostSystemId = k.HostSystemId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHostPortGroup_Name(k *HostPortGroupParameters, p *HostPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHostPortGroup_NotifySwitches(k *HostPortGroupParameters, p *HostPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.NotifySwitches != p.NotifySwitches {
		p.NotifySwitches = k.NotifySwitches
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHostPortGroup_ShapingAverageBandwidth(k *HostPortGroupParameters, p *HostPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.ShapingAverageBandwidth != p.ShapingAverageBandwidth {
		p.ShapingAverageBandwidth = k.ShapingAverageBandwidth
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHostPortGroup_ShapingBurstSize(k *HostPortGroupParameters, p *HostPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.ShapingBurstSize != p.ShapingBurstSize {
		p.ShapingBurstSize = k.ShapingBurstSize
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHostPortGroup_ShapingEnabled(k *HostPortGroupParameters, p *HostPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.ShapingEnabled != p.ShapingEnabled {
		p.ShapingEnabled = k.ShapingEnabled
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHostPortGroup_ShapingPeakBandwidth(k *HostPortGroupParameters, p *HostPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.ShapingPeakBandwidth != p.ShapingPeakBandwidth {
		p.ShapingPeakBandwidth = k.ShapingPeakBandwidth
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeHostPortGroup_StandbyNics(k *HostPortGroupParameters, p *HostPortGroupParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.StandbyNics, p.StandbyNics) {
		p.StandbyNics = k.StandbyNics
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHostPortGroup_TeamingPolicy(k *HostPortGroupParameters, p *HostPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.TeamingPolicy != p.TeamingPolicy {
		p.TeamingPolicy = k.TeamingPolicy
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHostPortGroup_VirtualSwitchName(k *HostPortGroupParameters, p *HostPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.VirtualSwitchName != p.VirtualSwitchName {
		p.VirtualSwitchName = k.VirtualSwitchName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHostPortGroup_VlanId(k *HostPortGroupParameters, p *HostPortGroupParameters, md *plugin.MergeDescription) bool {
	if k.VlanId != p.VlanId {
		p.VlanId = k.VlanId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateStatus
func MergeHostPortGroup_ComputedPolicy(k *HostPortGroupObservation, p *HostPortGroupObservation, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.ComputedPolicy, p.ComputedPolicy) {
		k.ComputedPolicy = p.ComputedPolicy
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeHostPortGroup_Key(k *HostPortGroupObservation, p *HostPortGroupObservation, md *plugin.MergeDescription) bool {
	if k.Key != p.Key {
		k.Key = p.Key
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergeStructSliceTemplateStatus
func MergeHostPortGroup_Ports(ksp *[]Ports, psp *[]Ports, md *plugin.MergeDescription) bool {
	ks := *ksp
	ps := *psp
	if len(ks) != len(ps) {
		ks = ps
		md.NeedsProviderUpdate = true
		return true
	}
	anyChildUpdated := false
	for i, _ := range ps {
		updated := false
		k := &ks[i]
		p := &ps[i]
		updated = MergeHostPortGroup_Ports_Key(k, p, md)
		if updated {
			anyChildUpdated = true
		}

		updated = MergeHostPortGroup_Ports_MacAddresses(k, p, md)
		if updated {
			anyChildUpdated = true
		}

		updated = MergeHostPortGroup_Ports_Type(k, p, md)
		if updated {
			anyChildUpdated = true
		}

	}
	if anyChildUpdated {
		md.StatusUpdated = true
	}
	return anyChildUpdated
}

//mergePrimitiveTemplateStatus
func MergeHostPortGroup_Ports_Key(k *Ports, p *Ports, md *plugin.MergeDescription) bool {
	if k.Key != p.Key {
		k.Key = p.Key
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateStatus
func MergeHostPortGroup_Ports_MacAddresses(k *Ports, p *Ports, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.MacAddresses, p.MacAddresses) {
		k.MacAddresses = p.MacAddresses
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeHostPortGroup_Ports_Type(k *Ports, p *Ports, md *plugin.MergeDescription) bool {
	if k.Type != p.Type {
		k.Type = p.Type
		md.StatusUpdated = true
		return true
	}
	return false
}
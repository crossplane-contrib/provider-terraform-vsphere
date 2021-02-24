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
	k := kube.(*VappContainer)
	p := prov.(*VappContainer)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeVappContainer_CpuExpandable(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVappContainer_CpuLimit(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVappContainer_CpuReservation(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVappContainer_CpuShareLevel(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVappContainer_CpuShares(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVappContainer_CustomAttributes(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVappContainer_MemoryExpandable(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVappContainer_MemoryLimit(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVappContainer_MemoryReservation(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVappContainer_MemoryShareLevel(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVappContainer_MemoryShares(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVappContainer_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVappContainer_ParentFolderId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVappContainer_ParentResourcePoolId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVappContainer_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeVappContainer_CpuExpandable(k *VappContainerParameters, p *VappContainerParameters, md *plugin.MergeDescription) bool {
	if k.CpuExpandable != p.CpuExpandable {
		p.CpuExpandable = k.CpuExpandable
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVappContainer_CpuLimit(k *VappContainerParameters, p *VappContainerParameters, md *plugin.MergeDescription) bool {
	if k.CpuLimit != p.CpuLimit {
		p.CpuLimit = k.CpuLimit
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVappContainer_CpuReservation(k *VappContainerParameters, p *VappContainerParameters, md *plugin.MergeDescription) bool {
	if k.CpuReservation != p.CpuReservation {
		p.CpuReservation = k.CpuReservation
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVappContainer_CpuShareLevel(k *VappContainerParameters, p *VappContainerParameters, md *plugin.MergeDescription) bool {
	if k.CpuShareLevel != p.CpuShareLevel {
		p.CpuShareLevel = k.CpuShareLevel
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVappContainer_CpuShares(k *VappContainerParameters, p *VappContainerParameters, md *plugin.MergeDescription) bool {
	if k.CpuShares != p.CpuShares {
		p.CpuShares = k.CpuShares
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeVappContainer_CustomAttributes(k *VappContainerParameters, p *VappContainerParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.CustomAttributes, p.CustomAttributes) {
		p.CustomAttributes = k.CustomAttributes
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVappContainer_MemoryExpandable(k *VappContainerParameters, p *VappContainerParameters, md *plugin.MergeDescription) bool {
	if k.MemoryExpandable != p.MemoryExpandable {
		p.MemoryExpandable = k.MemoryExpandable
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVappContainer_MemoryLimit(k *VappContainerParameters, p *VappContainerParameters, md *plugin.MergeDescription) bool {
	if k.MemoryLimit != p.MemoryLimit {
		p.MemoryLimit = k.MemoryLimit
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVappContainer_MemoryReservation(k *VappContainerParameters, p *VappContainerParameters, md *plugin.MergeDescription) bool {
	if k.MemoryReservation != p.MemoryReservation {
		p.MemoryReservation = k.MemoryReservation
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVappContainer_MemoryShareLevel(k *VappContainerParameters, p *VappContainerParameters, md *plugin.MergeDescription) bool {
	if k.MemoryShareLevel != p.MemoryShareLevel {
		p.MemoryShareLevel = k.MemoryShareLevel
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVappContainer_MemoryShares(k *VappContainerParameters, p *VappContainerParameters, md *plugin.MergeDescription) bool {
	if k.MemoryShares != p.MemoryShares {
		p.MemoryShares = k.MemoryShares
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVappContainer_Name(k *VappContainerParameters, p *VappContainerParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVappContainer_ParentFolderId(k *VappContainerParameters, p *VappContainerParameters, md *plugin.MergeDescription) bool {
	if k.ParentFolderId != p.ParentFolderId {
		p.ParentFolderId = k.ParentFolderId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVappContainer_ParentResourcePoolId(k *VappContainerParameters, p *VappContainerParameters, md *plugin.MergeDescription) bool {
	if k.ParentResourcePoolId != p.ParentResourcePoolId {
		p.ParentResourcePoolId = k.ParentResourcePoolId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeVappContainer_Tags(k *VappContainerParameters, p *VappContainerParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}
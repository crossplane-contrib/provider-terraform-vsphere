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
	k := kube.(*VirtualMachineSnapshot)
	p := prov.(*VirtualMachineSnapshot)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeVirtualMachineSnapshot_Consolidate(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachineSnapshot_Description(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachineSnapshot_Memory(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachineSnapshot_Quiesce(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachineSnapshot_RemoveChildren(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachineSnapshot_SnapshotName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachineSnapshot_VirtualMachineUuid(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeVirtualMachineSnapshot_Consolidate(k *VirtualMachineSnapshotParameters, p *VirtualMachineSnapshotParameters, md *plugin.MergeDescription) bool {
	if k.Consolidate != p.Consolidate {
		p.Consolidate = k.Consolidate
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachineSnapshot_Description(k *VirtualMachineSnapshotParameters, p *VirtualMachineSnapshotParameters, md *plugin.MergeDescription) bool {
	if k.Description != p.Description {
		p.Description = k.Description
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachineSnapshot_Memory(k *VirtualMachineSnapshotParameters, p *VirtualMachineSnapshotParameters, md *plugin.MergeDescription) bool {
	if k.Memory != p.Memory {
		p.Memory = k.Memory
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachineSnapshot_Quiesce(k *VirtualMachineSnapshotParameters, p *VirtualMachineSnapshotParameters, md *plugin.MergeDescription) bool {
	if k.Quiesce != p.Quiesce {
		p.Quiesce = k.Quiesce
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachineSnapshot_RemoveChildren(k *VirtualMachineSnapshotParameters, p *VirtualMachineSnapshotParameters, md *plugin.MergeDescription) bool {
	if k.RemoveChildren != p.RemoveChildren {
		p.RemoveChildren = k.RemoveChildren
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachineSnapshot_SnapshotName(k *VirtualMachineSnapshotParameters, p *VirtualMachineSnapshotParameters, md *plugin.MergeDescription) bool {
	if k.SnapshotName != p.SnapshotName {
		p.SnapshotName = k.SnapshotName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachineSnapshot_VirtualMachineUuid(k *VirtualMachineSnapshotParameters, p *VirtualMachineSnapshotParameters, md *plugin.MergeDescription) bool {
	if k.VirtualMachineUuid != p.VirtualMachineUuid {
		p.VirtualMachineUuid = k.VirtualMachineUuid
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

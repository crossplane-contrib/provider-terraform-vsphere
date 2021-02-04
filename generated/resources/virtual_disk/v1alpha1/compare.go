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
	k := kube.(*VirtualDisk)
	p := prov.(*VirtualDisk)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeVirtualDisk_AdapterType(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualDisk_CreateDirectories(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualDisk_Datacenter(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualDisk_Datastore(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualDisk_Size(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualDisk_Type(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualDisk_VmdkPath(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeVirtualDisk_AdapterType(k *VirtualDiskParameters, p *VirtualDiskParameters, md *plugin.MergeDescription) bool {
	if k.AdapterType != p.AdapterType {
		p.AdapterType = k.AdapterType
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualDisk_CreateDirectories(k *VirtualDiskParameters, p *VirtualDiskParameters, md *plugin.MergeDescription) bool {
	if k.CreateDirectories != p.CreateDirectories {
		p.CreateDirectories = k.CreateDirectories
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualDisk_Datacenter(k *VirtualDiskParameters, p *VirtualDiskParameters, md *plugin.MergeDescription) bool {
	if k.Datacenter != p.Datacenter {
		p.Datacenter = k.Datacenter
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualDisk_Datastore(k *VirtualDiskParameters, p *VirtualDiskParameters, md *plugin.MergeDescription) bool {
	if k.Datastore != p.Datastore {
		p.Datastore = k.Datastore
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualDisk_Size(k *VirtualDiskParameters, p *VirtualDiskParameters, md *plugin.MergeDescription) bool {
	if k.Size != p.Size {
		p.Size = k.Size
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualDisk_Type(k *VirtualDiskParameters, p *VirtualDiskParameters, md *plugin.MergeDescription) bool {
	if k.Type != p.Type {
		p.Type = k.Type
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualDisk_VmdkPath(k *VirtualDiskParameters, p *VirtualDiskParameters, md *plugin.MergeDescription) bool {
	if k.VmdkPath != p.VmdkPath {
		p.VmdkPath = k.VmdkPath
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}
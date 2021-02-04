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
	k := kube.(*VmfsDatastore)
	p := prov.(*VmfsDatastore)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeVmfsDatastore_CustomAttributes(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVmfsDatastore_Folder(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVmfsDatastore_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVmfsDatastore_DatastoreClusterId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVmfsDatastore_Disks(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVmfsDatastore_HostSystemId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVmfsDatastore_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVmfsDatastore_Capacity(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVmfsDatastore_FreeSpace(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVmfsDatastore_UncommittedSpace(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVmfsDatastore_Accessible(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVmfsDatastore_MaintenanceMode(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVmfsDatastore_MultipleHostAccess(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVmfsDatastore_Url(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeVmfsDatastore_CustomAttributes(k *VmfsDatastoreParameters, p *VmfsDatastoreParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.CustomAttributes, p.CustomAttributes) {
		p.CustomAttributes = k.CustomAttributes
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVmfsDatastore_Folder(k *VmfsDatastoreParameters, p *VmfsDatastoreParameters, md *plugin.MergeDescription) bool {
	if k.Folder != p.Folder {
		p.Folder = k.Folder
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeVmfsDatastore_Tags(k *VmfsDatastoreParameters, p *VmfsDatastoreParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVmfsDatastore_DatastoreClusterId(k *VmfsDatastoreParameters, p *VmfsDatastoreParameters, md *plugin.MergeDescription) bool {
	if k.DatastoreClusterId != p.DatastoreClusterId {
		p.DatastoreClusterId = k.DatastoreClusterId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeVmfsDatastore_Disks(k *VmfsDatastoreParameters, p *VmfsDatastoreParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.Disks, p.Disks) {
		p.Disks = k.Disks
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVmfsDatastore_HostSystemId(k *VmfsDatastoreParameters, p *VmfsDatastoreParameters, md *plugin.MergeDescription) bool {
	if k.HostSystemId != p.HostSystemId {
		p.HostSystemId = k.HostSystemId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVmfsDatastore_Name(k *VmfsDatastoreParameters, p *VmfsDatastoreParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeVmfsDatastore_Capacity(k *VmfsDatastoreObservation, p *VmfsDatastoreObservation, md *plugin.MergeDescription) bool {
	if k.Capacity != p.Capacity {
		k.Capacity = p.Capacity
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeVmfsDatastore_FreeSpace(k *VmfsDatastoreObservation, p *VmfsDatastoreObservation, md *plugin.MergeDescription) bool {
	if k.FreeSpace != p.FreeSpace {
		k.FreeSpace = p.FreeSpace
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeVmfsDatastore_UncommittedSpace(k *VmfsDatastoreObservation, p *VmfsDatastoreObservation, md *plugin.MergeDescription) bool {
	if k.UncommittedSpace != p.UncommittedSpace {
		k.UncommittedSpace = p.UncommittedSpace
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeVmfsDatastore_Accessible(k *VmfsDatastoreObservation, p *VmfsDatastoreObservation, md *plugin.MergeDescription) bool {
	if k.Accessible != p.Accessible {
		k.Accessible = p.Accessible
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeVmfsDatastore_MaintenanceMode(k *VmfsDatastoreObservation, p *VmfsDatastoreObservation, md *plugin.MergeDescription) bool {
	if k.MaintenanceMode != p.MaintenanceMode {
		k.MaintenanceMode = p.MaintenanceMode
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeVmfsDatastore_MultipleHostAccess(k *VmfsDatastoreObservation, p *VmfsDatastoreObservation, md *plugin.MergeDescription) bool {
	if k.MultipleHostAccess != p.MultipleHostAccess {
		k.MultipleHostAccess = p.MultipleHostAccess
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeVmfsDatastore_Url(k *VmfsDatastoreObservation, p *VmfsDatastoreObservation, md *plugin.MergeDescription) bool {
	if k.Url != p.Url {
		k.Url = p.Url
		md.StatusUpdated = true
		return true
	}
	return false
}
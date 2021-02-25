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
	k := kube.(*NasDatastore)
	p := prov.(*NasDatastore)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeNasDatastore_AccessMode(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNasDatastore_CustomAttributes(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNasDatastore_DatastoreClusterId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNasDatastore_Folder(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNasDatastore_HostSystemIds(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNasDatastore_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNasDatastore_RemoteHosts(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNasDatastore_RemotePath(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNasDatastore_SecurityType(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNasDatastore_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNasDatastore_Type(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNasDatastore_Accessible(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNasDatastore_Capacity(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNasDatastore_FreeSpace(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNasDatastore_MaintenanceMode(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNasDatastore_MultipleHostAccess(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNasDatastore_ProtocolEndpoint(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNasDatastore_UncommittedSpace(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeNasDatastore_Url(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeNasDatastore_AccessMode(k *NasDatastoreParameters, p *NasDatastoreParameters, md *plugin.MergeDescription) bool {
	if k.AccessMode != p.AccessMode {
		p.AccessMode = k.AccessMode
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeNasDatastore_CustomAttributes(k *NasDatastoreParameters, p *NasDatastoreParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.CustomAttributes, p.CustomAttributes) {
		p.CustomAttributes = k.CustomAttributes
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNasDatastore_DatastoreClusterId(k *NasDatastoreParameters, p *NasDatastoreParameters, md *plugin.MergeDescription) bool {
	if k.DatastoreClusterId != p.DatastoreClusterId {
		p.DatastoreClusterId = k.DatastoreClusterId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNasDatastore_Folder(k *NasDatastoreParameters, p *NasDatastoreParameters, md *plugin.MergeDescription) bool {
	if k.Folder != p.Folder {
		p.Folder = k.Folder
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeNasDatastore_HostSystemIds(k *NasDatastoreParameters, p *NasDatastoreParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.HostSystemIds, p.HostSystemIds) {
		p.HostSystemIds = k.HostSystemIds
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNasDatastore_Name(k *NasDatastoreParameters, p *NasDatastoreParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeNasDatastore_RemoteHosts(k *NasDatastoreParameters, p *NasDatastoreParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.RemoteHosts, p.RemoteHosts) {
		p.RemoteHosts = k.RemoteHosts
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNasDatastore_RemotePath(k *NasDatastoreParameters, p *NasDatastoreParameters, md *plugin.MergeDescription) bool {
	if k.RemotePath != p.RemotePath {
		p.RemotePath = k.RemotePath
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNasDatastore_SecurityType(k *NasDatastoreParameters, p *NasDatastoreParameters, md *plugin.MergeDescription) bool {
	if k.SecurityType != p.SecurityType {
		p.SecurityType = k.SecurityType
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeNasDatastore_Tags(k *NasDatastoreParameters, p *NasDatastoreParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeNasDatastore_Type(k *NasDatastoreParameters, p *NasDatastoreParameters, md *plugin.MergeDescription) bool {
	if k.Type != p.Type {
		p.Type = k.Type
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeNasDatastore_Accessible(k *NasDatastoreObservation, p *NasDatastoreObservation, md *plugin.MergeDescription) bool {
	if k.Accessible != p.Accessible {
		k.Accessible = p.Accessible
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeNasDatastore_Capacity(k *NasDatastoreObservation, p *NasDatastoreObservation, md *plugin.MergeDescription) bool {
	if k.Capacity != p.Capacity {
		k.Capacity = p.Capacity
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeNasDatastore_FreeSpace(k *NasDatastoreObservation, p *NasDatastoreObservation, md *plugin.MergeDescription) bool {
	if k.FreeSpace != p.FreeSpace {
		k.FreeSpace = p.FreeSpace
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeNasDatastore_MaintenanceMode(k *NasDatastoreObservation, p *NasDatastoreObservation, md *plugin.MergeDescription) bool {
	if k.MaintenanceMode != p.MaintenanceMode {
		k.MaintenanceMode = p.MaintenanceMode
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeNasDatastore_MultipleHostAccess(k *NasDatastoreObservation, p *NasDatastoreObservation, md *plugin.MergeDescription) bool {
	if k.MultipleHostAccess != p.MultipleHostAccess {
		k.MultipleHostAccess = p.MultipleHostAccess
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeNasDatastore_ProtocolEndpoint(k *NasDatastoreObservation, p *NasDatastoreObservation, md *plugin.MergeDescription) bool {
	if k.ProtocolEndpoint != p.ProtocolEndpoint {
		k.ProtocolEndpoint = p.ProtocolEndpoint
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeNasDatastore_UncommittedSpace(k *NasDatastoreObservation, p *NasDatastoreObservation, md *plugin.MergeDescription) bool {
	if k.UncommittedSpace != p.UncommittedSpace {
		k.UncommittedSpace = p.UncommittedSpace
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeNasDatastore_Url(k *NasDatastoreObservation, p *NasDatastoreObservation, md *plugin.MergeDescription) bool {
	if k.Url != p.Url {
		k.Url = p.Url
		md.StatusUpdated = true
		return true
	}
	return false
}

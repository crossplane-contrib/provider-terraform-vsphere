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
	k := kube.(*StorageDrsVmOverride)
	p := prov.(*StorageDrsVmOverride)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeStorageDrsVmOverride_DatastoreClusterId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeStorageDrsVmOverride_SdrsAutomationLevel(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeStorageDrsVmOverride_SdrsEnabled(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeStorageDrsVmOverride_SdrsIntraVmAffinity(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeStorageDrsVmOverride_VirtualMachineId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeStorageDrsVmOverride_DatastoreClusterId(k *StorageDrsVmOverrideParameters, p *StorageDrsVmOverrideParameters, md *plugin.MergeDescription) bool {
	if k.DatastoreClusterId != p.DatastoreClusterId {
		p.DatastoreClusterId = k.DatastoreClusterId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeStorageDrsVmOverride_SdrsAutomationLevel(k *StorageDrsVmOverrideParameters, p *StorageDrsVmOverrideParameters, md *plugin.MergeDescription) bool {
	if k.SdrsAutomationLevel != p.SdrsAutomationLevel {
		p.SdrsAutomationLevel = k.SdrsAutomationLevel
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeStorageDrsVmOverride_SdrsEnabled(k *StorageDrsVmOverrideParameters, p *StorageDrsVmOverrideParameters, md *plugin.MergeDescription) bool {
	if k.SdrsEnabled != p.SdrsEnabled {
		p.SdrsEnabled = k.SdrsEnabled
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeStorageDrsVmOverride_SdrsIntraVmAffinity(k *StorageDrsVmOverrideParameters, p *StorageDrsVmOverrideParameters, md *plugin.MergeDescription) bool {
	if k.SdrsIntraVmAffinity != p.SdrsIntraVmAffinity {
		p.SdrsIntraVmAffinity = k.SdrsIntraVmAffinity
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeStorageDrsVmOverride_VirtualMachineId(k *StorageDrsVmOverrideParameters, p *StorageDrsVmOverrideParameters, md *plugin.MergeDescription) bool {
	if k.VirtualMachineId != p.VirtualMachineId {
		p.VirtualMachineId = k.VirtualMachineId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}
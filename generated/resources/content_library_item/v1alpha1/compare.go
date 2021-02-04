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
	k := kube.(*ContentLibraryItem)
	p := prov.(*ContentLibraryItem)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeContentLibraryItem_SourceUuid(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeContentLibraryItem_Type(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeContentLibraryItem_Description(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeContentLibraryItem_FileUrl(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeContentLibraryItem_LibraryId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeContentLibraryItem_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeContentLibraryItem_SourceUuid(k *ContentLibraryItemParameters, p *ContentLibraryItemParameters, md *plugin.MergeDescription) bool {
	if k.SourceUuid != p.SourceUuid {
		p.SourceUuid = k.SourceUuid
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeContentLibraryItem_Type(k *ContentLibraryItemParameters, p *ContentLibraryItemParameters, md *plugin.MergeDescription) bool {
	if k.Type != p.Type {
		p.Type = k.Type
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeContentLibraryItem_Description(k *ContentLibraryItemParameters, p *ContentLibraryItemParameters, md *plugin.MergeDescription) bool {
	if k.Description != p.Description {
		p.Description = k.Description
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeContentLibraryItem_FileUrl(k *ContentLibraryItemParameters, p *ContentLibraryItemParameters, md *plugin.MergeDescription) bool {
	if k.FileUrl != p.FileUrl {
		p.FileUrl = k.FileUrl
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeContentLibraryItem_LibraryId(k *ContentLibraryItemParameters, p *ContentLibraryItemParameters, md *plugin.MergeDescription) bool {
	if k.LibraryId != p.LibraryId {
		p.LibraryId = k.LibraryId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeContentLibraryItem_Name(k *ContentLibraryItemParameters, p *ContentLibraryItemParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}
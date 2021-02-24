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
	k := kube.(*VmStoragePolicy)
	p := prov.(*VmStoragePolicy)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeVmStoragePolicy_Description(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVmStoragePolicy_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVmStoragePolicy_TagRules(&k.Spec.ForProvider.TagRules, &p.Spec.ForProvider.TagRules, md)
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
func MergeVmStoragePolicy_Description(k *VmStoragePolicyParameters, p *VmStoragePolicyParameters, md *plugin.MergeDescription) bool {
	if k.Description != p.Description {
		p.Description = k.Description
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVmStoragePolicy_Name(k *VmStoragePolicyParameters, p *VmStoragePolicyParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructSliceTemplateSpec
func MergeVmStoragePolicy_TagRules(ksp *[]TagRules, psp *[]TagRules, md *plugin.MergeDescription) bool {
	ks := *ksp
	ps := *psp
	if len(ks) != len(ps) {
		ps = ks
		md.NeedsProviderUpdate = true
		return true
	}
	anyChildUpdated := false
	for i, _ := range ps {
		updated := false
		k := &ks[i]
		p := &ps[i]
		updated = MergeVmStoragePolicy_TagRules_IncludeDatastoresWithTags(k, p, md)
		if updated {
			anyChildUpdated = true
		}

		updated = MergeVmStoragePolicy_TagRules_TagCategory(k, p, md)
		if updated {
			anyChildUpdated = true
		}

		updated = MergeVmStoragePolicy_TagRules_Tags(k, p, md)
		if updated {
			anyChildUpdated = true
		}

	}
	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveTemplateSpec
func MergeVmStoragePolicy_TagRules_IncludeDatastoresWithTags(k *TagRules, p *TagRules, md *plugin.MergeDescription) bool {
	if k.IncludeDatastoresWithTags != p.IncludeDatastoresWithTags {
		p.IncludeDatastoresWithTags = k.IncludeDatastoresWithTags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVmStoragePolicy_TagRules_TagCategory(k *TagRules, p *TagRules, md *plugin.MergeDescription) bool {
	if k.TagCategory != p.TagCategory {
		p.TagCategory = k.TagCategory
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeVmStoragePolicy_TagRules_Tags(k *TagRules, p *TagRules, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}
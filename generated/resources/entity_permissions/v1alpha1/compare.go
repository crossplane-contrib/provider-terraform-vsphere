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
	k := kube.(*EntityPermissions)
	p := prov.(*EntityPermissions)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeEntityPermissions_EntityId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEntityPermissions_EntityType(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeEntityPermissions_Permissions(&k.Spec.ForProvider.Permissions, &p.Spec.ForProvider.Permissions, md)
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
func MergeEntityPermissions_EntityId(k *EntityPermissionsParameters, p *EntityPermissionsParameters, md *plugin.MergeDescription) bool {
	if k.EntityId != p.EntityId {
		p.EntityId = k.EntityId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEntityPermissions_EntityType(k *EntityPermissionsParameters, p *EntityPermissionsParameters, md *plugin.MergeDescription) bool {
	if k.EntityType != p.EntityType {
		p.EntityType = k.EntityType
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructSliceTemplateSpec
func MergeEntityPermissions_Permissions(ksp *[]Permissions, psp *[]Permissions, md *plugin.MergeDescription) bool {
	if len(*ksp) != len(*psp) {
		*psp = *ksp
		md.NeedsProviderUpdate = true
		return true
	}
	ks := *ksp
	ps := *psp
	anyChildUpdated := false
	for i := range ps {
		updated := false
		k := &ks[i]
		p := &ps[i]
		updated = MergeEntityPermissions_Permissions_IsGroup(k, p, md)
		if updated {
			anyChildUpdated = true
		}

		updated = MergeEntityPermissions_Permissions_Propagate(k, p, md)
		if updated {
			anyChildUpdated = true
		}

		updated = MergeEntityPermissions_Permissions_RoleId(k, p, md)
		if updated {
			anyChildUpdated = true
		}

		updated = MergeEntityPermissions_Permissions_UserOrGroup(k, p, md)
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
func MergeEntityPermissions_Permissions_IsGroup(k *Permissions, p *Permissions, md *plugin.MergeDescription) bool {
	if k.IsGroup != p.IsGroup {
		p.IsGroup = k.IsGroup
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEntityPermissions_Permissions_Propagate(k *Permissions, p *Permissions, md *plugin.MergeDescription) bool {
	if k.Propagate != p.Propagate {
		p.Propagate = k.Propagate
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEntityPermissions_Permissions_RoleId(k *Permissions, p *Permissions, md *plugin.MergeDescription) bool {
	if k.RoleId != p.RoleId {
		p.RoleId = k.RoleId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeEntityPermissions_Permissions_UserOrGroup(k *Permissions, p *Permissions, md *plugin.MergeDescription) bool {
	if k.UserOrGroup != p.UserOrGroup {
		p.UserOrGroup = k.UserOrGroup
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

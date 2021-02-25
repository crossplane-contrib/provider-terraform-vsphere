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
	k := kube.(*VappEntity)
	p := prov.(*VappEntity)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeVappEntity_ContainerId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVappEntity_CustomAttributes(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVappEntity_StartAction(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVappEntity_StartDelay(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVappEntity_StartOrder(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVappEntity_StopAction(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVappEntity_StopDelay(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVappEntity_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVappEntity_TargetId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVappEntity_WaitForGuest(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeVappEntity_ContainerId(k *VappEntityParameters, p *VappEntityParameters, md *plugin.MergeDescription) bool {
	if k.ContainerId != p.ContainerId {
		p.ContainerId = k.ContainerId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeVappEntity_CustomAttributes(k *VappEntityParameters, p *VappEntityParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.CustomAttributes, p.CustomAttributes) {
		p.CustomAttributes = k.CustomAttributes
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVappEntity_StartAction(k *VappEntityParameters, p *VappEntityParameters, md *plugin.MergeDescription) bool {
	if k.StartAction != p.StartAction {
		p.StartAction = k.StartAction
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVappEntity_StartDelay(k *VappEntityParameters, p *VappEntityParameters, md *plugin.MergeDescription) bool {
	if k.StartDelay != p.StartDelay {
		p.StartDelay = k.StartDelay
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVappEntity_StartOrder(k *VappEntityParameters, p *VappEntityParameters, md *plugin.MergeDescription) bool {
	if k.StartOrder != p.StartOrder {
		p.StartOrder = k.StartOrder
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVappEntity_StopAction(k *VappEntityParameters, p *VappEntityParameters, md *plugin.MergeDescription) bool {
	if k.StopAction != p.StopAction {
		p.StopAction = k.StopAction
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVappEntity_StopDelay(k *VappEntityParameters, p *VappEntityParameters, md *plugin.MergeDescription) bool {
	if k.StopDelay != p.StopDelay {
		p.StopDelay = k.StopDelay
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeVappEntity_Tags(k *VappEntityParameters, p *VappEntityParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVappEntity_TargetId(k *VappEntityParameters, p *VappEntityParameters, md *plugin.MergeDescription) bool {
	if k.TargetId != p.TargetId {
		p.TargetId = k.TargetId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVappEntity_WaitForGuest(k *VappEntityParameters, p *VappEntityParameters, md *plugin.MergeDescription) bool {
	if k.WaitForGuest != p.WaitForGuest {
		p.WaitForGuest = k.WaitForGuest
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

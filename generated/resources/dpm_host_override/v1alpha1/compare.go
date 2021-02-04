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
	k := kube.(*DpmHostOverride)
	p := prov.(*DpmHostOverride)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeDpmHostOverride_ComputeClusterId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDpmHostOverride_DpmAutomationLevel(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDpmHostOverride_DpmEnabled(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDpmHostOverride_HostSystemId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeDpmHostOverride_ComputeClusterId(k *DpmHostOverrideParameters, p *DpmHostOverrideParameters, md *plugin.MergeDescription) bool {
	if k.ComputeClusterId != p.ComputeClusterId {
		p.ComputeClusterId = k.ComputeClusterId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDpmHostOverride_DpmAutomationLevel(k *DpmHostOverrideParameters, p *DpmHostOverrideParameters, md *plugin.MergeDescription) bool {
	if k.DpmAutomationLevel != p.DpmAutomationLevel {
		p.DpmAutomationLevel = k.DpmAutomationLevel
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDpmHostOverride_DpmEnabled(k *DpmHostOverrideParameters, p *DpmHostOverrideParameters, md *plugin.MergeDescription) bool {
	if k.DpmEnabled != p.DpmEnabled {
		p.DpmEnabled = k.DpmEnabled
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDpmHostOverride_HostSystemId(k *DpmHostOverrideParameters, p *DpmHostOverrideParameters, md *plugin.MergeDescription) bool {
	if k.HostSystemId != p.HostSystemId {
		p.HostSystemId = k.HostSystemId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}
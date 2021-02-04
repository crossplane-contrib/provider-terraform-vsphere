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
	k := kube.(*DrsVmOverride)
	p := prov.(*DrsVmOverride)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeDrsVmOverride_DrsAutomationLevel(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDrsVmOverride_DrsEnabled(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDrsVmOverride_VirtualMachineId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDrsVmOverride_ComputeClusterId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeDrsVmOverride_DrsAutomationLevel(k *DrsVmOverrideParameters, p *DrsVmOverrideParameters, md *plugin.MergeDescription) bool {
	if k.DrsAutomationLevel != p.DrsAutomationLevel {
		p.DrsAutomationLevel = k.DrsAutomationLevel
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDrsVmOverride_DrsEnabled(k *DrsVmOverrideParameters, p *DrsVmOverrideParameters, md *plugin.MergeDescription) bool {
	if k.DrsEnabled != p.DrsEnabled {
		p.DrsEnabled = k.DrsEnabled
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDrsVmOverride_VirtualMachineId(k *DrsVmOverrideParameters, p *DrsVmOverrideParameters, md *plugin.MergeDescription) bool {
	if k.VirtualMachineId != p.VirtualMachineId {
		p.VirtualMachineId = k.VirtualMachineId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDrsVmOverride_ComputeClusterId(k *DrsVmOverrideParameters, p *DrsVmOverrideParameters, md *plugin.MergeDescription) bool {
	if k.ComputeClusterId != p.ComputeClusterId {
		p.ComputeClusterId = k.ComputeClusterId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}
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
	k := kube.(*ComputeClusterVmDependencyRule)
	p := prov.(*ComputeClusterVmDependencyRule)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeComputeClusterVmDependencyRule_ComputeClusterId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeClusterVmDependencyRule_DependencyVmGroupName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeClusterVmDependencyRule_Enabled(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeClusterVmDependencyRule_Mandatory(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeClusterVmDependencyRule_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeClusterVmDependencyRule_VmGroupName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeComputeClusterVmDependencyRule_ComputeClusterId(k *ComputeClusterVmDependencyRuleParameters, p *ComputeClusterVmDependencyRuleParameters, md *plugin.MergeDescription) bool {
	if k.ComputeClusterId != p.ComputeClusterId {
		p.ComputeClusterId = k.ComputeClusterId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeClusterVmDependencyRule_DependencyVmGroupName(k *ComputeClusterVmDependencyRuleParameters, p *ComputeClusterVmDependencyRuleParameters, md *plugin.MergeDescription) bool {
	if k.DependencyVmGroupName != p.DependencyVmGroupName {
		p.DependencyVmGroupName = k.DependencyVmGroupName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeClusterVmDependencyRule_Enabled(k *ComputeClusterVmDependencyRuleParameters, p *ComputeClusterVmDependencyRuleParameters, md *plugin.MergeDescription) bool {
	if k.Enabled != p.Enabled {
		p.Enabled = k.Enabled
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeClusterVmDependencyRule_Mandatory(k *ComputeClusterVmDependencyRuleParameters, p *ComputeClusterVmDependencyRuleParameters, md *plugin.MergeDescription) bool {
	if k.Mandatory != p.Mandatory {
		p.Mandatory = k.Mandatory
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeClusterVmDependencyRule_Name(k *ComputeClusterVmDependencyRuleParameters, p *ComputeClusterVmDependencyRuleParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeClusterVmDependencyRule_VmGroupName(k *ComputeClusterVmDependencyRuleParameters, p *ComputeClusterVmDependencyRuleParameters, md *plugin.MergeDescription) bool {
	if k.VmGroupName != p.VmGroupName {
		p.VmGroupName = k.VmGroupName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

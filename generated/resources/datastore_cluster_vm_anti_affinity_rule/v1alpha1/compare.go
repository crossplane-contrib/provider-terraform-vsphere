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
	k := kube.(*DatastoreClusterVmAntiAffinityRule)
	p := prov.(*DatastoreClusterVmAntiAffinityRule)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeDatastoreClusterVmAntiAffinityRule_DatastoreClusterId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDatastoreClusterVmAntiAffinityRule_Enabled(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDatastoreClusterVmAntiAffinityRule_Mandatory(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDatastoreClusterVmAntiAffinityRule_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDatastoreClusterVmAntiAffinityRule_VirtualMachineIds(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeDatastoreClusterVmAntiAffinityRule_DatastoreClusterId(k *DatastoreClusterVmAntiAffinityRuleParameters, p *DatastoreClusterVmAntiAffinityRuleParameters, md *plugin.MergeDescription) bool {
	if k.DatastoreClusterId != p.DatastoreClusterId {
		p.DatastoreClusterId = k.DatastoreClusterId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDatastoreClusterVmAntiAffinityRule_Enabled(k *DatastoreClusterVmAntiAffinityRuleParameters, p *DatastoreClusterVmAntiAffinityRuleParameters, md *plugin.MergeDescription) bool {
	if k.Enabled != p.Enabled {
		p.Enabled = k.Enabled
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDatastoreClusterVmAntiAffinityRule_Mandatory(k *DatastoreClusterVmAntiAffinityRuleParameters, p *DatastoreClusterVmAntiAffinityRuleParameters, md *plugin.MergeDescription) bool {
	if k.Mandatory != p.Mandatory {
		p.Mandatory = k.Mandatory
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDatastoreClusterVmAntiAffinityRule_Name(k *DatastoreClusterVmAntiAffinityRuleParameters, p *DatastoreClusterVmAntiAffinityRuleParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeDatastoreClusterVmAntiAffinityRule_VirtualMachineIds(k *DatastoreClusterVmAntiAffinityRuleParameters, p *DatastoreClusterVmAntiAffinityRuleParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.VirtualMachineIds, p.VirtualMachineIds) {
		p.VirtualMachineIds = k.VirtualMachineIds
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}
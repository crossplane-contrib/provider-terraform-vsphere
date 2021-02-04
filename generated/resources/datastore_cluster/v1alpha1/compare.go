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
	k := kube.(*DatastoreCluster)
	p := prov.(*DatastoreCluster)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeDatastoreCluster_SdrsDefaultIntraVmAffinity(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDatastoreCluster_SdrsIoLoadImbalanceThreshold(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDatastoreCluster_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDatastoreCluster_SdrsIoReservablePercentThreshold(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDatastoreCluster_SdrsIoReservableThresholdMode(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDatastoreCluster_SdrsRuleEnforcementAutomationLevel(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDatastoreCluster_CustomAttributes(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDatastoreCluster_DatacenterId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDatastoreCluster_SdrsFreeSpaceUtilizationDifference(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDatastoreCluster_SdrsIoLatencyThreshold(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDatastoreCluster_SdrsSpaceBalanceAutomationLevel(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDatastoreCluster_SdrsSpaceUtilizationThreshold(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDatastoreCluster_SdrsVmEvacuationAutomationLevel(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDatastoreCluster_SdrsPolicyEnforcementAutomationLevel(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDatastoreCluster_Folder(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDatastoreCluster_SdrsAutomationLevel(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDatastoreCluster_SdrsFreeSpaceThreshold(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDatastoreCluster_SdrsIoBalanceAutomationLevel(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDatastoreCluster_SdrsLoadBalanceInterval(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDatastoreCluster_SdrsIoReservableIopsThreshold(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDatastoreCluster_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDatastoreCluster_SdrsAdvancedOptions(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDatastoreCluster_SdrsEnabled(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDatastoreCluster_SdrsFreeSpaceThresholdMode(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeDatastoreCluster_SdrsIoLoadBalanceEnabled(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeDatastoreCluster_SdrsDefaultIntraVmAffinity(k *DatastoreClusterParameters, p *DatastoreClusterParameters, md *plugin.MergeDescription) bool {
	if k.SdrsDefaultIntraVmAffinity != p.SdrsDefaultIntraVmAffinity {
		p.SdrsDefaultIntraVmAffinity = k.SdrsDefaultIntraVmAffinity
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDatastoreCluster_SdrsIoLoadImbalanceThreshold(k *DatastoreClusterParameters, p *DatastoreClusterParameters, md *plugin.MergeDescription) bool {
	if k.SdrsIoLoadImbalanceThreshold != p.SdrsIoLoadImbalanceThreshold {
		p.SdrsIoLoadImbalanceThreshold = k.SdrsIoLoadImbalanceThreshold
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeDatastoreCluster_Tags(k *DatastoreClusterParameters, p *DatastoreClusterParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDatastoreCluster_SdrsIoReservablePercentThreshold(k *DatastoreClusterParameters, p *DatastoreClusterParameters, md *plugin.MergeDescription) bool {
	if k.SdrsIoReservablePercentThreshold != p.SdrsIoReservablePercentThreshold {
		p.SdrsIoReservablePercentThreshold = k.SdrsIoReservablePercentThreshold
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDatastoreCluster_SdrsIoReservableThresholdMode(k *DatastoreClusterParameters, p *DatastoreClusterParameters, md *plugin.MergeDescription) bool {
	if k.SdrsIoReservableThresholdMode != p.SdrsIoReservableThresholdMode {
		p.SdrsIoReservableThresholdMode = k.SdrsIoReservableThresholdMode
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDatastoreCluster_SdrsRuleEnforcementAutomationLevel(k *DatastoreClusterParameters, p *DatastoreClusterParameters, md *plugin.MergeDescription) bool {
	if k.SdrsRuleEnforcementAutomationLevel != p.SdrsRuleEnforcementAutomationLevel {
		p.SdrsRuleEnforcementAutomationLevel = k.SdrsRuleEnforcementAutomationLevel
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeDatastoreCluster_CustomAttributes(k *DatastoreClusterParameters, p *DatastoreClusterParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.CustomAttributes, p.CustomAttributes) {
		p.CustomAttributes = k.CustomAttributes
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDatastoreCluster_DatacenterId(k *DatastoreClusterParameters, p *DatastoreClusterParameters, md *plugin.MergeDescription) bool {
	if k.DatacenterId != p.DatacenterId {
		p.DatacenterId = k.DatacenterId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDatastoreCluster_SdrsFreeSpaceUtilizationDifference(k *DatastoreClusterParameters, p *DatastoreClusterParameters, md *plugin.MergeDescription) bool {
	if k.SdrsFreeSpaceUtilizationDifference != p.SdrsFreeSpaceUtilizationDifference {
		p.SdrsFreeSpaceUtilizationDifference = k.SdrsFreeSpaceUtilizationDifference
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDatastoreCluster_SdrsIoLatencyThreshold(k *DatastoreClusterParameters, p *DatastoreClusterParameters, md *plugin.MergeDescription) bool {
	if k.SdrsIoLatencyThreshold != p.SdrsIoLatencyThreshold {
		p.SdrsIoLatencyThreshold = k.SdrsIoLatencyThreshold
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDatastoreCluster_SdrsSpaceBalanceAutomationLevel(k *DatastoreClusterParameters, p *DatastoreClusterParameters, md *plugin.MergeDescription) bool {
	if k.SdrsSpaceBalanceAutomationLevel != p.SdrsSpaceBalanceAutomationLevel {
		p.SdrsSpaceBalanceAutomationLevel = k.SdrsSpaceBalanceAutomationLevel
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDatastoreCluster_SdrsSpaceUtilizationThreshold(k *DatastoreClusterParameters, p *DatastoreClusterParameters, md *plugin.MergeDescription) bool {
	if k.SdrsSpaceUtilizationThreshold != p.SdrsSpaceUtilizationThreshold {
		p.SdrsSpaceUtilizationThreshold = k.SdrsSpaceUtilizationThreshold
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDatastoreCluster_SdrsVmEvacuationAutomationLevel(k *DatastoreClusterParameters, p *DatastoreClusterParameters, md *plugin.MergeDescription) bool {
	if k.SdrsVmEvacuationAutomationLevel != p.SdrsVmEvacuationAutomationLevel {
		p.SdrsVmEvacuationAutomationLevel = k.SdrsVmEvacuationAutomationLevel
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDatastoreCluster_SdrsPolicyEnforcementAutomationLevel(k *DatastoreClusterParameters, p *DatastoreClusterParameters, md *plugin.MergeDescription) bool {
	if k.SdrsPolicyEnforcementAutomationLevel != p.SdrsPolicyEnforcementAutomationLevel {
		p.SdrsPolicyEnforcementAutomationLevel = k.SdrsPolicyEnforcementAutomationLevel
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDatastoreCluster_Folder(k *DatastoreClusterParameters, p *DatastoreClusterParameters, md *plugin.MergeDescription) bool {
	if k.Folder != p.Folder {
		p.Folder = k.Folder
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDatastoreCluster_SdrsAutomationLevel(k *DatastoreClusterParameters, p *DatastoreClusterParameters, md *plugin.MergeDescription) bool {
	if k.SdrsAutomationLevel != p.SdrsAutomationLevel {
		p.SdrsAutomationLevel = k.SdrsAutomationLevel
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDatastoreCluster_SdrsFreeSpaceThreshold(k *DatastoreClusterParameters, p *DatastoreClusterParameters, md *plugin.MergeDescription) bool {
	if k.SdrsFreeSpaceThreshold != p.SdrsFreeSpaceThreshold {
		p.SdrsFreeSpaceThreshold = k.SdrsFreeSpaceThreshold
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDatastoreCluster_SdrsIoBalanceAutomationLevel(k *DatastoreClusterParameters, p *DatastoreClusterParameters, md *plugin.MergeDescription) bool {
	if k.SdrsIoBalanceAutomationLevel != p.SdrsIoBalanceAutomationLevel {
		p.SdrsIoBalanceAutomationLevel = k.SdrsIoBalanceAutomationLevel
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDatastoreCluster_SdrsLoadBalanceInterval(k *DatastoreClusterParameters, p *DatastoreClusterParameters, md *plugin.MergeDescription) bool {
	if k.SdrsLoadBalanceInterval != p.SdrsLoadBalanceInterval {
		p.SdrsLoadBalanceInterval = k.SdrsLoadBalanceInterval
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDatastoreCluster_SdrsIoReservableIopsThreshold(k *DatastoreClusterParameters, p *DatastoreClusterParameters, md *plugin.MergeDescription) bool {
	if k.SdrsIoReservableIopsThreshold != p.SdrsIoReservableIopsThreshold {
		p.SdrsIoReservableIopsThreshold = k.SdrsIoReservableIopsThreshold
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDatastoreCluster_Name(k *DatastoreClusterParameters, p *DatastoreClusterParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeDatastoreCluster_SdrsAdvancedOptions(k *DatastoreClusterParameters, p *DatastoreClusterParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.SdrsAdvancedOptions, p.SdrsAdvancedOptions) {
		p.SdrsAdvancedOptions = k.SdrsAdvancedOptions
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDatastoreCluster_SdrsEnabled(k *DatastoreClusterParameters, p *DatastoreClusterParameters, md *plugin.MergeDescription) bool {
	if k.SdrsEnabled != p.SdrsEnabled {
		p.SdrsEnabled = k.SdrsEnabled
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDatastoreCluster_SdrsFreeSpaceThresholdMode(k *DatastoreClusterParameters, p *DatastoreClusterParameters, md *plugin.MergeDescription) bool {
	if k.SdrsFreeSpaceThresholdMode != p.SdrsFreeSpaceThresholdMode {
		p.SdrsFreeSpaceThresholdMode = k.SdrsFreeSpaceThresholdMode
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeDatastoreCluster_SdrsIoLoadBalanceEnabled(k *DatastoreClusterParameters, p *DatastoreClusterParameters, md *plugin.MergeDescription) bool {
	if k.SdrsIoLoadBalanceEnabled != p.SdrsIoLoadBalanceEnabled {
		p.SdrsIoLoadBalanceEnabled = k.SdrsIoLoadBalanceEnabled
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}
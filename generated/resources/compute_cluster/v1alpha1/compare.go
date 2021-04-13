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
	k := kube.(*ComputeCluster)
	p := prov.(*ComputeCluster)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeComputeCluster_CustomAttributes(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_DatacenterId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_DpmAutomationLevel(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_DpmEnabled(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_DpmThreshold(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_DrsAdvancedOptions(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_DrsAutomationLevel(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_DrsEnablePredictiveDrs(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_DrsEnableVmOverrides(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_DrsEnabled(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_DrsMigrationThreshold(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_Folder(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_ForceEvacuateOnDestroy(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_HaAdmissionControlFailoverHostSystemIds(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_HaAdmissionControlHostFailureTolerance(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_HaAdmissionControlPerformanceTolerance(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_HaAdmissionControlPolicy(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_HaAdmissionControlResourcePercentageAutoCompute(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_HaAdmissionControlResourcePercentageCpu(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_HaAdmissionControlResourcePercentageMemory(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_HaAdmissionControlSlotPolicyExplicitCpu(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_HaAdmissionControlSlotPolicyExplicitMemory(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_HaAdmissionControlSlotPolicyUseExplicitSize(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_HaAdvancedOptions(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_HaDatastoreApdRecoveryAction(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_HaDatastoreApdResponse(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_HaDatastoreApdResponseDelay(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_HaDatastorePdlResponse(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_HaEnabled(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_HaHeartbeatDatastoreIds(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_HaHeartbeatDatastorePolicy(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_HaHostIsolationResponse(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_HaHostMonitoring(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_HaVmComponentProtection(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_HaVmDependencyRestartCondition(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_HaVmFailureInterval(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_HaVmMaximumFailureWindow(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_HaVmMaximumResets(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_HaVmMinimumUptime(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_HaVmMonitoring(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_HaVmRestartAdditionalDelay(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_HaVmRestartPriority(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_HaVmRestartTimeout(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_HostClusterExitTimeout(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_HostManaged(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_HostSystemIds(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_ProactiveHaAutomationLevel(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_ProactiveHaEnabled(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_ProactiveHaModerateRemediation(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_ProactiveHaProviderIds(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_ProactiveHaSevereRemediation(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_VsanDiskGroup(&k.Spec.ForProvider.VsanDiskGroup, &p.Spec.ForProvider.VsanDiskGroup, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_VsanEnabled(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_ResourcePoolId(&k.Status.AtProvider, &p.Status.AtProvider, md)
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

//mergePrimitiveContainerTemplateSpec
func MergeComputeCluster_CustomAttributes(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.CustomAttributes, p.CustomAttributes) {
		p.CustomAttributes = k.CustomAttributes
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeCluster_DatacenterId(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if k.DatacenterId != p.DatacenterId {
		p.DatacenterId = k.DatacenterId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeCluster_DpmAutomationLevel(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if k.DpmAutomationLevel != p.DpmAutomationLevel {
		p.DpmAutomationLevel = k.DpmAutomationLevel
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeCluster_DpmEnabled(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if k.DpmEnabled != p.DpmEnabled {
		p.DpmEnabled = k.DpmEnabled
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeCluster_DpmThreshold(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if k.DpmThreshold != p.DpmThreshold {
		p.DpmThreshold = k.DpmThreshold
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeComputeCluster_DrsAdvancedOptions(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.DrsAdvancedOptions, p.DrsAdvancedOptions) {
		p.DrsAdvancedOptions = k.DrsAdvancedOptions
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeCluster_DrsAutomationLevel(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if k.DrsAutomationLevel != p.DrsAutomationLevel {
		p.DrsAutomationLevel = k.DrsAutomationLevel
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeCluster_DrsEnablePredictiveDrs(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if k.DrsEnablePredictiveDrs != p.DrsEnablePredictiveDrs {
		p.DrsEnablePredictiveDrs = k.DrsEnablePredictiveDrs
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeCluster_DrsEnableVmOverrides(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if k.DrsEnableVmOverrides != p.DrsEnableVmOverrides {
		p.DrsEnableVmOverrides = k.DrsEnableVmOverrides
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeCluster_DrsEnabled(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if k.DrsEnabled != p.DrsEnabled {
		p.DrsEnabled = k.DrsEnabled
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeCluster_DrsMigrationThreshold(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if k.DrsMigrationThreshold != p.DrsMigrationThreshold {
		p.DrsMigrationThreshold = k.DrsMigrationThreshold
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeCluster_Folder(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if k.Folder != p.Folder {
		p.Folder = k.Folder
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeCluster_ForceEvacuateOnDestroy(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if k.ForceEvacuateOnDestroy != p.ForceEvacuateOnDestroy {
		p.ForceEvacuateOnDestroy = k.ForceEvacuateOnDestroy
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeComputeCluster_HaAdmissionControlFailoverHostSystemIds(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.HaAdmissionControlFailoverHostSystemIds, p.HaAdmissionControlFailoverHostSystemIds) {
		p.HaAdmissionControlFailoverHostSystemIds = k.HaAdmissionControlFailoverHostSystemIds
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeCluster_HaAdmissionControlHostFailureTolerance(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if k.HaAdmissionControlHostFailureTolerance != p.HaAdmissionControlHostFailureTolerance {
		p.HaAdmissionControlHostFailureTolerance = k.HaAdmissionControlHostFailureTolerance
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeCluster_HaAdmissionControlPerformanceTolerance(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if k.HaAdmissionControlPerformanceTolerance != p.HaAdmissionControlPerformanceTolerance {
		p.HaAdmissionControlPerformanceTolerance = k.HaAdmissionControlPerformanceTolerance
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeCluster_HaAdmissionControlPolicy(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if k.HaAdmissionControlPolicy != p.HaAdmissionControlPolicy {
		p.HaAdmissionControlPolicy = k.HaAdmissionControlPolicy
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeCluster_HaAdmissionControlResourcePercentageAutoCompute(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if k.HaAdmissionControlResourcePercentageAutoCompute != p.HaAdmissionControlResourcePercentageAutoCompute {
		p.HaAdmissionControlResourcePercentageAutoCompute = k.HaAdmissionControlResourcePercentageAutoCompute
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeCluster_HaAdmissionControlResourcePercentageCpu(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if k.HaAdmissionControlResourcePercentageCpu != p.HaAdmissionControlResourcePercentageCpu {
		p.HaAdmissionControlResourcePercentageCpu = k.HaAdmissionControlResourcePercentageCpu
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeCluster_HaAdmissionControlResourcePercentageMemory(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if k.HaAdmissionControlResourcePercentageMemory != p.HaAdmissionControlResourcePercentageMemory {
		p.HaAdmissionControlResourcePercentageMemory = k.HaAdmissionControlResourcePercentageMemory
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeCluster_HaAdmissionControlSlotPolicyExplicitCpu(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if k.HaAdmissionControlSlotPolicyExplicitCpu != p.HaAdmissionControlSlotPolicyExplicitCpu {
		p.HaAdmissionControlSlotPolicyExplicitCpu = k.HaAdmissionControlSlotPolicyExplicitCpu
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeCluster_HaAdmissionControlSlotPolicyExplicitMemory(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if k.HaAdmissionControlSlotPolicyExplicitMemory != p.HaAdmissionControlSlotPolicyExplicitMemory {
		p.HaAdmissionControlSlotPolicyExplicitMemory = k.HaAdmissionControlSlotPolicyExplicitMemory
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeCluster_HaAdmissionControlSlotPolicyUseExplicitSize(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if k.HaAdmissionControlSlotPolicyUseExplicitSize != p.HaAdmissionControlSlotPolicyUseExplicitSize {
		p.HaAdmissionControlSlotPolicyUseExplicitSize = k.HaAdmissionControlSlotPolicyUseExplicitSize
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeComputeCluster_HaAdvancedOptions(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.HaAdvancedOptions, p.HaAdvancedOptions) {
		p.HaAdvancedOptions = k.HaAdvancedOptions
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeCluster_HaDatastoreApdRecoveryAction(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if k.HaDatastoreApdRecoveryAction != p.HaDatastoreApdRecoveryAction {
		p.HaDatastoreApdRecoveryAction = k.HaDatastoreApdRecoveryAction
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeCluster_HaDatastoreApdResponse(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if k.HaDatastoreApdResponse != p.HaDatastoreApdResponse {
		p.HaDatastoreApdResponse = k.HaDatastoreApdResponse
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeCluster_HaDatastoreApdResponseDelay(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if k.HaDatastoreApdResponseDelay != p.HaDatastoreApdResponseDelay {
		p.HaDatastoreApdResponseDelay = k.HaDatastoreApdResponseDelay
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeCluster_HaDatastorePdlResponse(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if k.HaDatastorePdlResponse != p.HaDatastorePdlResponse {
		p.HaDatastorePdlResponse = k.HaDatastorePdlResponse
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeCluster_HaEnabled(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if k.HaEnabled != p.HaEnabled {
		p.HaEnabled = k.HaEnabled
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeComputeCluster_HaHeartbeatDatastoreIds(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.HaHeartbeatDatastoreIds, p.HaHeartbeatDatastoreIds) {
		p.HaHeartbeatDatastoreIds = k.HaHeartbeatDatastoreIds
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeCluster_HaHeartbeatDatastorePolicy(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if k.HaHeartbeatDatastorePolicy != p.HaHeartbeatDatastorePolicy {
		p.HaHeartbeatDatastorePolicy = k.HaHeartbeatDatastorePolicy
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeCluster_HaHostIsolationResponse(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if k.HaHostIsolationResponse != p.HaHostIsolationResponse {
		p.HaHostIsolationResponse = k.HaHostIsolationResponse
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeCluster_HaHostMonitoring(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if k.HaHostMonitoring != p.HaHostMonitoring {
		p.HaHostMonitoring = k.HaHostMonitoring
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeCluster_HaVmComponentProtection(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if k.HaVmComponentProtection != p.HaVmComponentProtection {
		p.HaVmComponentProtection = k.HaVmComponentProtection
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeCluster_HaVmDependencyRestartCondition(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if k.HaVmDependencyRestartCondition != p.HaVmDependencyRestartCondition {
		p.HaVmDependencyRestartCondition = k.HaVmDependencyRestartCondition
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeCluster_HaVmFailureInterval(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if k.HaVmFailureInterval != p.HaVmFailureInterval {
		p.HaVmFailureInterval = k.HaVmFailureInterval
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeCluster_HaVmMaximumFailureWindow(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if k.HaVmMaximumFailureWindow != p.HaVmMaximumFailureWindow {
		p.HaVmMaximumFailureWindow = k.HaVmMaximumFailureWindow
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeCluster_HaVmMaximumResets(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if k.HaVmMaximumResets != p.HaVmMaximumResets {
		p.HaVmMaximumResets = k.HaVmMaximumResets
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeCluster_HaVmMinimumUptime(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if k.HaVmMinimumUptime != p.HaVmMinimumUptime {
		p.HaVmMinimumUptime = k.HaVmMinimumUptime
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeCluster_HaVmMonitoring(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if k.HaVmMonitoring != p.HaVmMonitoring {
		p.HaVmMonitoring = k.HaVmMonitoring
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeCluster_HaVmRestartAdditionalDelay(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if k.HaVmRestartAdditionalDelay != p.HaVmRestartAdditionalDelay {
		p.HaVmRestartAdditionalDelay = k.HaVmRestartAdditionalDelay
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeCluster_HaVmRestartPriority(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if k.HaVmRestartPriority != p.HaVmRestartPriority {
		p.HaVmRestartPriority = k.HaVmRestartPriority
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeCluster_HaVmRestartTimeout(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if k.HaVmRestartTimeout != p.HaVmRestartTimeout {
		p.HaVmRestartTimeout = k.HaVmRestartTimeout
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeCluster_HostClusterExitTimeout(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if k.HostClusterExitTimeout != p.HostClusterExitTimeout {
		p.HostClusterExitTimeout = k.HostClusterExitTimeout
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeCluster_HostManaged(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if k.HostManaged != p.HostManaged {
		p.HostManaged = k.HostManaged
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeComputeCluster_HostSystemIds(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.HostSystemIds, p.HostSystemIds) {
		p.HostSystemIds = k.HostSystemIds
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeCluster_Name(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeCluster_ProactiveHaAutomationLevel(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if k.ProactiveHaAutomationLevel != p.ProactiveHaAutomationLevel {
		p.ProactiveHaAutomationLevel = k.ProactiveHaAutomationLevel
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeCluster_ProactiveHaEnabled(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if k.ProactiveHaEnabled != p.ProactiveHaEnabled {
		p.ProactiveHaEnabled = k.ProactiveHaEnabled
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeCluster_ProactiveHaModerateRemediation(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if k.ProactiveHaModerateRemediation != p.ProactiveHaModerateRemediation {
		p.ProactiveHaModerateRemediation = k.ProactiveHaModerateRemediation
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeComputeCluster_ProactiveHaProviderIds(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.ProactiveHaProviderIds, p.ProactiveHaProviderIds) {
		p.ProactiveHaProviderIds = k.ProactiveHaProviderIds
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeCluster_ProactiveHaSevereRemediation(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if k.ProactiveHaSevereRemediation != p.ProactiveHaSevereRemediation {
		p.ProactiveHaSevereRemediation = k.ProactiveHaSevereRemediation
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeComputeCluster_Tags(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructTemplateSpec
func MergeComputeCluster_VsanDiskGroup(k *VsanDiskGroup, p *VsanDiskGroup, md *plugin.MergeDescription) bool {
	updated := false
	anyChildUpdated := false
	updated = MergeComputeCluster_VsanDiskGroup_Cache(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeComputeCluster_VsanDiskGroup_Storage(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveTemplateSpec
func MergeComputeCluster_VsanDiskGroup_Cache(k *VsanDiskGroup, p *VsanDiskGroup, md *plugin.MergeDescription) bool {
	if k.Cache != p.Cache {
		p.Cache = k.Cache
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeComputeCluster_VsanDiskGroup_Storage(k *VsanDiskGroup, p *VsanDiskGroup, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.Storage, p.Storage) {
		p.Storage = k.Storage
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeComputeCluster_VsanEnabled(k *ComputeClusterParameters, p *ComputeClusterParameters, md *plugin.MergeDescription) bool {
	if k.VsanEnabled != p.VsanEnabled {
		p.VsanEnabled = k.VsanEnabled
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeComputeCluster_ResourcePoolId(k *ComputeClusterObservation, p *ComputeClusterObservation, md *plugin.MergeDescription) bool {
	if k.ResourcePoolId != p.ResourcePoolId {
		k.ResourcePoolId = p.ResourcePoolId
		md.StatusUpdated = true
		return true
	}
	return false
}

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
	"fmt"

	ctwhy "github.com/crossplane-contrib/terraform-runtime/pkg/plugin/cty"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
	"github.com/zclconf/go-cty/cty"
)

type ctyDecoder struct{}

func (e *ctyDecoder) DecodeCty(mr resource.Managed, ctyValue cty.Value, schema *providers.Schema) (resource.Managed, error) {
	r, ok := mr.(*ComputeCluster)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeComputeCluster(r, ctyValue)
}

func DecodeComputeCluster(prev *ComputeCluster, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeComputeCluster_CustomAttributes(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_DatacenterId(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_DpmAutomationLevel(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_DpmEnabled(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_DpmThreshold(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_DrsAdvancedOptions(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_DrsAutomationLevel(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_DrsEnablePredictiveDrs(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_DrsEnableVmOverrides(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_DrsEnabled(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_DrsMigrationThreshold(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_Folder(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_ForceEvacuateOnDestroy(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_HaAdmissionControlFailoverHostSystemIds(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_HaAdmissionControlHostFailureTolerance(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_HaAdmissionControlPerformanceTolerance(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_HaAdmissionControlPolicy(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_HaAdmissionControlResourcePercentageAutoCompute(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_HaAdmissionControlResourcePercentageCpu(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_HaAdmissionControlResourcePercentageMemory(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_HaAdmissionControlSlotPolicyExplicitCpu(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_HaAdmissionControlSlotPolicyExplicitMemory(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_HaAdmissionControlSlotPolicyUseExplicitSize(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_HaAdvancedOptions(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_HaDatastoreApdRecoveryAction(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_HaDatastoreApdResponse(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_HaDatastoreApdResponseDelay(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_HaDatastorePdlResponse(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_HaEnabled(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_HaHeartbeatDatastoreIds(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_HaHeartbeatDatastorePolicy(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_HaHostIsolationResponse(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_HaHostMonitoring(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_HaVmComponentProtection(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_HaVmDependencyRestartCondition(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_HaVmFailureInterval(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_HaVmMaximumFailureWindow(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_HaVmMaximumResets(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_HaVmMinimumUptime(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_HaVmMonitoring(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_HaVmRestartAdditionalDelay(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_HaVmRestartPriority(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_HaVmRestartTimeout(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_HostClusterExitTimeout(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_HostManaged(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_HostSystemIds(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_Name(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_ProactiveHaAutomationLevel(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_ProactiveHaEnabled(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_ProactiveHaModerateRemediation(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_ProactiveHaProviderIds(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_ProactiveHaSevereRemediation(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_Tags(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_VsanDiskGroup(&new.Spec.ForProvider.VsanDiskGroup, valMap)
	DecodeComputeCluster_VsanEnabled(&new.Spec.ForProvider, valMap)
	DecodeComputeCluster_ResourcePoolId(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveMapTypeDecodeTemplate
func DecodeComputeCluster_CustomAttributes(p *ComputeClusterParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	if vals["custom_attributes"].IsNull() {
		p.CustomAttributes = nil
		return
	}
	vMap := make(map[string]string)
	v := vals["custom_attributes"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.CustomAttributes = vMap
}

//primitiveTypeDecodeTemplate
func DecodeComputeCluster_DatacenterId(p *ComputeClusterParameters, vals map[string]cty.Value) {
	p.DatacenterId = ctwhy.ValueAsString(vals["datacenter_id"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeCluster_DpmAutomationLevel(p *ComputeClusterParameters, vals map[string]cty.Value) {
	p.DpmAutomationLevel = ctwhy.ValueAsString(vals["dpm_automation_level"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeCluster_DpmEnabled(p *ComputeClusterParameters, vals map[string]cty.Value) {
	p.DpmEnabled = ctwhy.ValueAsBool(vals["dpm_enabled"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeCluster_DpmThreshold(p *ComputeClusterParameters, vals map[string]cty.Value) {
	p.DpmThreshold = ctwhy.ValueAsInt64(vals["dpm_threshold"])
}

//primitiveMapTypeDecodeTemplate
func DecodeComputeCluster_DrsAdvancedOptions(p *ComputeClusterParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	if vals["drs_advanced_options"].IsNull() {
		p.DrsAdvancedOptions = nil
		return
	}
	vMap := make(map[string]string)
	v := vals["drs_advanced_options"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.DrsAdvancedOptions = vMap
}

//primitiveTypeDecodeTemplate
func DecodeComputeCluster_DrsAutomationLevel(p *ComputeClusterParameters, vals map[string]cty.Value) {
	p.DrsAutomationLevel = ctwhy.ValueAsString(vals["drs_automation_level"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeCluster_DrsEnablePredictiveDrs(p *ComputeClusterParameters, vals map[string]cty.Value) {
	p.DrsEnablePredictiveDrs = ctwhy.ValueAsBool(vals["drs_enable_predictive_drs"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeCluster_DrsEnableVmOverrides(p *ComputeClusterParameters, vals map[string]cty.Value) {
	p.DrsEnableVmOverrides = ctwhy.ValueAsBool(vals["drs_enable_vm_overrides"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeCluster_DrsEnabled(p *ComputeClusterParameters, vals map[string]cty.Value) {
	p.DrsEnabled = ctwhy.ValueAsBool(vals["drs_enabled"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeCluster_DrsMigrationThreshold(p *ComputeClusterParameters, vals map[string]cty.Value) {
	p.DrsMigrationThreshold = ctwhy.ValueAsInt64(vals["drs_migration_threshold"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeCluster_Folder(p *ComputeClusterParameters, vals map[string]cty.Value) {
	p.Folder = ctwhy.ValueAsString(vals["folder"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeCluster_ForceEvacuateOnDestroy(p *ComputeClusterParameters, vals map[string]cty.Value) {
	p.ForceEvacuateOnDestroy = ctwhy.ValueAsBool(vals["force_evacuate_on_destroy"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeComputeCluster_HaAdmissionControlFailoverHostSystemIds(p *ComputeClusterParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["ha_admission_control_failover_host_system_ids"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.HaAdmissionControlFailoverHostSystemIds = goVals
}

//primitiveTypeDecodeTemplate
func DecodeComputeCluster_HaAdmissionControlHostFailureTolerance(p *ComputeClusterParameters, vals map[string]cty.Value) {
	p.HaAdmissionControlHostFailureTolerance = ctwhy.ValueAsInt64(vals["ha_admission_control_host_failure_tolerance"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeCluster_HaAdmissionControlPerformanceTolerance(p *ComputeClusterParameters, vals map[string]cty.Value) {
	p.HaAdmissionControlPerformanceTolerance = ctwhy.ValueAsInt64(vals["ha_admission_control_performance_tolerance"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeCluster_HaAdmissionControlPolicy(p *ComputeClusterParameters, vals map[string]cty.Value) {
	p.HaAdmissionControlPolicy = ctwhy.ValueAsString(vals["ha_admission_control_policy"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeCluster_HaAdmissionControlResourcePercentageAutoCompute(p *ComputeClusterParameters, vals map[string]cty.Value) {
	p.HaAdmissionControlResourcePercentageAutoCompute = ctwhy.ValueAsBool(vals["ha_admission_control_resource_percentage_auto_compute"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeCluster_HaAdmissionControlResourcePercentageCpu(p *ComputeClusterParameters, vals map[string]cty.Value) {
	p.HaAdmissionControlResourcePercentageCpu = ctwhy.ValueAsInt64(vals["ha_admission_control_resource_percentage_cpu"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeCluster_HaAdmissionControlResourcePercentageMemory(p *ComputeClusterParameters, vals map[string]cty.Value) {
	p.HaAdmissionControlResourcePercentageMemory = ctwhy.ValueAsInt64(vals["ha_admission_control_resource_percentage_memory"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeCluster_HaAdmissionControlSlotPolicyExplicitCpu(p *ComputeClusterParameters, vals map[string]cty.Value) {
	p.HaAdmissionControlSlotPolicyExplicitCpu = ctwhy.ValueAsInt64(vals["ha_admission_control_slot_policy_explicit_cpu"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeCluster_HaAdmissionControlSlotPolicyExplicitMemory(p *ComputeClusterParameters, vals map[string]cty.Value) {
	p.HaAdmissionControlSlotPolicyExplicitMemory = ctwhy.ValueAsInt64(vals["ha_admission_control_slot_policy_explicit_memory"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeCluster_HaAdmissionControlSlotPolicyUseExplicitSize(p *ComputeClusterParameters, vals map[string]cty.Value) {
	p.HaAdmissionControlSlotPolicyUseExplicitSize = ctwhy.ValueAsBool(vals["ha_admission_control_slot_policy_use_explicit_size"])
}

//primitiveMapTypeDecodeTemplate
func DecodeComputeCluster_HaAdvancedOptions(p *ComputeClusterParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	if vals["ha_advanced_options"].IsNull() {
		p.HaAdvancedOptions = nil
		return
	}
	vMap := make(map[string]string)
	v := vals["ha_advanced_options"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.HaAdvancedOptions = vMap
}

//primitiveTypeDecodeTemplate
func DecodeComputeCluster_HaDatastoreApdRecoveryAction(p *ComputeClusterParameters, vals map[string]cty.Value) {
	p.HaDatastoreApdRecoveryAction = ctwhy.ValueAsString(vals["ha_datastore_apd_recovery_action"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeCluster_HaDatastoreApdResponse(p *ComputeClusterParameters, vals map[string]cty.Value) {
	p.HaDatastoreApdResponse = ctwhy.ValueAsString(vals["ha_datastore_apd_response"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeCluster_HaDatastoreApdResponseDelay(p *ComputeClusterParameters, vals map[string]cty.Value) {
	p.HaDatastoreApdResponseDelay = ctwhy.ValueAsInt64(vals["ha_datastore_apd_response_delay"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeCluster_HaDatastorePdlResponse(p *ComputeClusterParameters, vals map[string]cty.Value) {
	p.HaDatastorePdlResponse = ctwhy.ValueAsString(vals["ha_datastore_pdl_response"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeCluster_HaEnabled(p *ComputeClusterParameters, vals map[string]cty.Value) {
	p.HaEnabled = ctwhy.ValueAsBool(vals["ha_enabled"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeComputeCluster_HaHeartbeatDatastoreIds(p *ComputeClusterParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["ha_heartbeat_datastore_ids"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.HaHeartbeatDatastoreIds = goVals
}

//primitiveTypeDecodeTemplate
func DecodeComputeCluster_HaHeartbeatDatastorePolicy(p *ComputeClusterParameters, vals map[string]cty.Value) {
	p.HaHeartbeatDatastorePolicy = ctwhy.ValueAsString(vals["ha_heartbeat_datastore_policy"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeCluster_HaHostIsolationResponse(p *ComputeClusterParameters, vals map[string]cty.Value) {
	p.HaHostIsolationResponse = ctwhy.ValueAsString(vals["ha_host_isolation_response"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeCluster_HaHostMonitoring(p *ComputeClusterParameters, vals map[string]cty.Value) {
	p.HaHostMonitoring = ctwhy.ValueAsString(vals["ha_host_monitoring"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeCluster_HaVmComponentProtection(p *ComputeClusterParameters, vals map[string]cty.Value) {
	p.HaVmComponentProtection = ctwhy.ValueAsString(vals["ha_vm_component_protection"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeCluster_HaVmDependencyRestartCondition(p *ComputeClusterParameters, vals map[string]cty.Value) {
	p.HaVmDependencyRestartCondition = ctwhy.ValueAsString(vals["ha_vm_dependency_restart_condition"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeCluster_HaVmFailureInterval(p *ComputeClusterParameters, vals map[string]cty.Value) {
	p.HaVmFailureInterval = ctwhy.ValueAsInt64(vals["ha_vm_failure_interval"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeCluster_HaVmMaximumFailureWindow(p *ComputeClusterParameters, vals map[string]cty.Value) {
	p.HaVmMaximumFailureWindow = ctwhy.ValueAsInt64(vals["ha_vm_maximum_failure_window"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeCluster_HaVmMaximumResets(p *ComputeClusterParameters, vals map[string]cty.Value) {
	p.HaVmMaximumResets = ctwhy.ValueAsInt64(vals["ha_vm_maximum_resets"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeCluster_HaVmMinimumUptime(p *ComputeClusterParameters, vals map[string]cty.Value) {
	p.HaVmMinimumUptime = ctwhy.ValueAsInt64(vals["ha_vm_minimum_uptime"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeCluster_HaVmMonitoring(p *ComputeClusterParameters, vals map[string]cty.Value) {
	p.HaVmMonitoring = ctwhy.ValueAsString(vals["ha_vm_monitoring"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeCluster_HaVmRestartAdditionalDelay(p *ComputeClusterParameters, vals map[string]cty.Value) {
	p.HaVmRestartAdditionalDelay = ctwhy.ValueAsInt64(vals["ha_vm_restart_additional_delay"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeCluster_HaVmRestartPriority(p *ComputeClusterParameters, vals map[string]cty.Value) {
	p.HaVmRestartPriority = ctwhy.ValueAsString(vals["ha_vm_restart_priority"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeCluster_HaVmRestartTimeout(p *ComputeClusterParameters, vals map[string]cty.Value) {
	p.HaVmRestartTimeout = ctwhy.ValueAsInt64(vals["ha_vm_restart_timeout"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeCluster_HostClusterExitTimeout(p *ComputeClusterParameters, vals map[string]cty.Value) {
	p.HostClusterExitTimeout = ctwhy.ValueAsInt64(vals["host_cluster_exit_timeout"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeCluster_HostManaged(p *ComputeClusterParameters, vals map[string]cty.Value) {
	p.HostManaged = ctwhy.ValueAsBool(vals["host_managed"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeComputeCluster_HostSystemIds(p *ComputeClusterParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["host_system_ids"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.HostSystemIds = goVals
}

//primitiveTypeDecodeTemplate
func DecodeComputeCluster_Name(p *ComputeClusterParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeCluster_ProactiveHaAutomationLevel(p *ComputeClusterParameters, vals map[string]cty.Value) {
	p.ProactiveHaAutomationLevel = ctwhy.ValueAsString(vals["proactive_ha_automation_level"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeCluster_ProactiveHaEnabled(p *ComputeClusterParameters, vals map[string]cty.Value) {
	p.ProactiveHaEnabled = ctwhy.ValueAsBool(vals["proactive_ha_enabled"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeCluster_ProactiveHaModerateRemediation(p *ComputeClusterParameters, vals map[string]cty.Value) {
	p.ProactiveHaModerateRemediation = ctwhy.ValueAsString(vals["proactive_ha_moderate_remediation"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeComputeCluster_ProactiveHaProviderIds(p *ComputeClusterParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["proactive_ha_provider_ids"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.ProactiveHaProviderIds = goVals
}

//primitiveTypeDecodeTemplate
func DecodeComputeCluster_ProactiveHaSevereRemediation(p *ComputeClusterParameters, vals map[string]cty.Value) {
	p.ProactiveHaSevereRemediation = ctwhy.ValueAsString(vals["proactive_ha_severe_remediation"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeComputeCluster_Tags(p *ComputeClusterParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["tags"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.Tags = goVals
}

//containerCollectionSingletonTypeDecodeTemplate
func DecodeComputeCluster_VsanDiskGroup(p *VsanDiskGroup, vals map[string]cty.Value) {
	if vals["vsan_disk_group"].IsNull() {
		*p = VsanDiskGroup{}
		return
	}
	rvals := ctwhy.ValueAsList(vals["vsan_disk_group"])
	if len(rvals) == 0 {
		*p = VsanDiskGroup{}
		return
	}
	// this template should be used when single dictionary/object values are nested in sets/lists
	// if rvals turns out to be a list with > 1 elements, something has broken with that heuristic
	valMap := rvals[0].AsValueMap()
	DecodeComputeCluster_VsanDiskGroup_Cache(p, valMap)
	DecodeComputeCluster_VsanDiskGroup_Storage(p, valMap)
}

//primitiveTypeDecodeTemplate
func DecodeComputeCluster_VsanDiskGroup_Cache(p *VsanDiskGroup, vals map[string]cty.Value) {
	p.Cache = ctwhy.ValueAsString(vals["cache"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeComputeCluster_VsanDiskGroup_Storage(p *VsanDiskGroup, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["storage"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.Storage = goVals
}

//primitiveTypeDecodeTemplate
func DecodeComputeCluster_VsanEnabled(p *ComputeClusterParameters, vals map[string]cty.Value) {
	p.VsanEnabled = ctwhy.ValueAsBool(vals["vsan_enabled"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeCluster_ResourcePoolId(p *ComputeClusterObservation, vals map[string]cty.Value) {
	p.ResourcePoolId = ctwhy.ValueAsString(vals["resource_pool_id"])
}

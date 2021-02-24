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

	"github.com/zclconf/go-cty/cty"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*ComputeCluster)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a ComputeCluster.")
	}
	return EncodeComputeCluster(*r), nil
}

func EncodeComputeCluster(r ComputeCluster) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeComputeCluster_CustomAttributes(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_DatacenterId(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_DpmAutomationLevel(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_DpmEnabled(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_DpmThreshold(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_DrsAdvancedOptions(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_DrsAutomationLevel(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_DrsEnablePredictiveDrs(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_DrsEnableVmOverrides(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_DrsEnabled(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_DrsMigrationThreshold(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_Folder(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_ForceEvacuateOnDestroy(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_HaAdmissionControlFailoverHostSystemIds(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_HaAdmissionControlHostFailureTolerance(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_HaAdmissionControlPerformanceTolerance(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_HaAdmissionControlPolicy(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_HaAdmissionControlResourcePercentageAutoCompute(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_HaAdmissionControlResourcePercentageCpu(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_HaAdmissionControlResourcePercentageMemory(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_HaAdmissionControlSlotPolicyExplicitCpu(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_HaAdmissionControlSlotPolicyExplicitMemory(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_HaAdmissionControlSlotPolicyUseExplicitSize(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_HaAdvancedOptions(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_HaDatastoreApdRecoveryAction(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_HaDatastoreApdResponse(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_HaDatastoreApdResponseDelay(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_HaDatastorePdlResponse(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_HaEnabled(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_HaHeartbeatDatastoreIds(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_HaHeartbeatDatastorePolicy(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_HaHostIsolationResponse(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_HaHostMonitoring(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_HaVmComponentProtection(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_HaVmDependencyRestartCondition(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_HaVmFailureInterval(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_HaVmMaximumFailureWindow(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_HaVmMaximumResets(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_HaVmMinimumUptime(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_HaVmMonitoring(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_HaVmRestartAdditionalDelay(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_HaVmRestartPriority(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_HaVmRestartTimeout(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_HostClusterExitTimeout(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_HostManaged(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_HostSystemIds(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_Name(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_ProactiveHaAutomationLevel(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_ProactiveHaEnabled(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_ProactiveHaModerateRemediation(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_ProactiveHaProviderIds(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_ProactiveHaSevereRemediation(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_Tags(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_VsanEnabled(r.Spec.ForProvider, ctyVal)
	EncodeComputeCluster_VsanDiskGroup(r.Spec.ForProvider.VsanDiskGroup, ctyVal)
	EncodeComputeCluster_ResourcePoolId(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeComputeCluster_CustomAttributes(p ComputeClusterParameters, vals map[string]cty.Value) {
	if len(p.CustomAttributes) == 0 {
		vals["custom_attributes"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.CustomAttributes {
		mVals[key] = cty.StringVal(value)
	}
	vals["custom_attributes"] = cty.MapVal(mVals)
}

func EncodeComputeCluster_DatacenterId(p ComputeClusterParameters, vals map[string]cty.Value) {
	vals["datacenter_id"] = cty.StringVal(p.DatacenterId)
}

func EncodeComputeCluster_DpmAutomationLevel(p ComputeClusterParameters, vals map[string]cty.Value) {
	vals["dpm_automation_level"] = cty.StringVal(p.DpmAutomationLevel)
}

func EncodeComputeCluster_DpmEnabled(p ComputeClusterParameters, vals map[string]cty.Value) {
	vals["dpm_enabled"] = cty.BoolVal(p.DpmEnabled)
}

func EncodeComputeCluster_DpmThreshold(p ComputeClusterParameters, vals map[string]cty.Value) {
	vals["dpm_threshold"] = cty.NumberIntVal(p.DpmThreshold)
}

func EncodeComputeCluster_DrsAdvancedOptions(p ComputeClusterParameters, vals map[string]cty.Value) {
	if len(p.DrsAdvancedOptions) == 0 {
		vals["drs_advanced_options"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.DrsAdvancedOptions {
		mVals[key] = cty.StringVal(value)
	}
	vals["drs_advanced_options"] = cty.MapVal(mVals)
}

func EncodeComputeCluster_DrsAutomationLevel(p ComputeClusterParameters, vals map[string]cty.Value) {
	vals["drs_automation_level"] = cty.StringVal(p.DrsAutomationLevel)
}

func EncodeComputeCluster_DrsEnablePredictiveDrs(p ComputeClusterParameters, vals map[string]cty.Value) {
	vals["drs_enable_predictive_drs"] = cty.BoolVal(p.DrsEnablePredictiveDrs)
}

func EncodeComputeCluster_DrsEnableVmOverrides(p ComputeClusterParameters, vals map[string]cty.Value) {
	vals["drs_enable_vm_overrides"] = cty.BoolVal(p.DrsEnableVmOverrides)
}

func EncodeComputeCluster_DrsEnabled(p ComputeClusterParameters, vals map[string]cty.Value) {
	vals["drs_enabled"] = cty.BoolVal(p.DrsEnabled)
}

func EncodeComputeCluster_DrsMigrationThreshold(p ComputeClusterParameters, vals map[string]cty.Value) {
	vals["drs_migration_threshold"] = cty.NumberIntVal(p.DrsMigrationThreshold)
}

func EncodeComputeCluster_Folder(p ComputeClusterParameters, vals map[string]cty.Value) {
	vals["folder"] = cty.StringVal(p.Folder)
}

func EncodeComputeCluster_ForceEvacuateOnDestroy(p ComputeClusterParameters, vals map[string]cty.Value) {
	vals["force_evacuate_on_destroy"] = cty.BoolVal(p.ForceEvacuateOnDestroy)
}

func EncodeComputeCluster_HaAdmissionControlFailoverHostSystemIds(p ComputeClusterParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.HaAdmissionControlFailoverHostSystemIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	if len(colVals) == 0 {
		vals["ha_admission_control_failover_host_system_ids"] = cty.SetValEmpty(cty.String)
	} else {
		vals["ha_admission_control_failover_host_system_ids"] = cty.SetVal(colVals)
    }
}

func EncodeComputeCluster_HaAdmissionControlHostFailureTolerance(p ComputeClusterParameters, vals map[string]cty.Value) {
	vals["ha_admission_control_host_failure_tolerance"] = cty.NumberIntVal(p.HaAdmissionControlHostFailureTolerance)
}

func EncodeComputeCluster_HaAdmissionControlPerformanceTolerance(p ComputeClusterParameters, vals map[string]cty.Value) {
	vals["ha_admission_control_performance_tolerance"] = cty.NumberIntVal(p.HaAdmissionControlPerformanceTolerance)
}

func EncodeComputeCluster_HaAdmissionControlPolicy(p ComputeClusterParameters, vals map[string]cty.Value) {
	vals["ha_admission_control_policy"] = cty.StringVal(p.HaAdmissionControlPolicy)
}

func EncodeComputeCluster_HaAdmissionControlResourcePercentageAutoCompute(p ComputeClusterParameters, vals map[string]cty.Value) {
	vals["ha_admission_control_resource_percentage_auto_compute"] = cty.BoolVal(p.HaAdmissionControlResourcePercentageAutoCompute)
}

func EncodeComputeCluster_HaAdmissionControlResourcePercentageCpu(p ComputeClusterParameters, vals map[string]cty.Value) {
	vals["ha_admission_control_resource_percentage_cpu"] = cty.NumberIntVal(p.HaAdmissionControlResourcePercentageCpu)
}

func EncodeComputeCluster_HaAdmissionControlResourcePercentageMemory(p ComputeClusterParameters, vals map[string]cty.Value) {
	vals["ha_admission_control_resource_percentage_memory"] = cty.NumberIntVal(p.HaAdmissionControlResourcePercentageMemory)
}

func EncodeComputeCluster_HaAdmissionControlSlotPolicyExplicitCpu(p ComputeClusterParameters, vals map[string]cty.Value) {
	vals["ha_admission_control_slot_policy_explicit_cpu"] = cty.NumberIntVal(p.HaAdmissionControlSlotPolicyExplicitCpu)
}

func EncodeComputeCluster_HaAdmissionControlSlotPolicyExplicitMemory(p ComputeClusterParameters, vals map[string]cty.Value) {
	vals["ha_admission_control_slot_policy_explicit_memory"] = cty.NumberIntVal(p.HaAdmissionControlSlotPolicyExplicitMemory)
}

func EncodeComputeCluster_HaAdmissionControlSlotPolicyUseExplicitSize(p ComputeClusterParameters, vals map[string]cty.Value) {
	vals["ha_admission_control_slot_policy_use_explicit_size"] = cty.BoolVal(p.HaAdmissionControlSlotPolicyUseExplicitSize)
}

func EncodeComputeCluster_HaAdvancedOptions(p ComputeClusterParameters, vals map[string]cty.Value) {
	if len(p.HaAdvancedOptions) == 0 {
		vals["ha_advanced_options"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.HaAdvancedOptions {
		mVals[key] = cty.StringVal(value)
	}
	vals["ha_advanced_options"] = cty.MapVal(mVals)
}

func EncodeComputeCluster_HaDatastoreApdRecoveryAction(p ComputeClusterParameters, vals map[string]cty.Value) {
	vals["ha_datastore_apd_recovery_action"] = cty.StringVal(p.HaDatastoreApdRecoveryAction)
}

func EncodeComputeCluster_HaDatastoreApdResponse(p ComputeClusterParameters, vals map[string]cty.Value) {
	vals["ha_datastore_apd_response"] = cty.StringVal(p.HaDatastoreApdResponse)
}

func EncodeComputeCluster_HaDatastoreApdResponseDelay(p ComputeClusterParameters, vals map[string]cty.Value) {
	vals["ha_datastore_apd_response_delay"] = cty.NumberIntVal(p.HaDatastoreApdResponseDelay)
}

func EncodeComputeCluster_HaDatastorePdlResponse(p ComputeClusterParameters, vals map[string]cty.Value) {
	vals["ha_datastore_pdl_response"] = cty.StringVal(p.HaDatastorePdlResponse)
}

func EncodeComputeCluster_HaEnabled(p ComputeClusterParameters, vals map[string]cty.Value) {
	vals["ha_enabled"] = cty.BoolVal(p.HaEnabled)
}

func EncodeComputeCluster_HaHeartbeatDatastoreIds(p ComputeClusterParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.HaHeartbeatDatastoreIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	if len(colVals) == 0 {
		vals["ha_heartbeat_datastore_ids"] = cty.SetValEmpty(cty.String)
	} else {
		vals["ha_heartbeat_datastore_ids"] = cty.SetVal(colVals)
    }
}

func EncodeComputeCluster_HaHeartbeatDatastorePolicy(p ComputeClusterParameters, vals map[string]cty.Value) {
	vals["ha_heartbeat_datastore_policy"] = cty.StringVal(p.HaHeartbeatDatastorePolicy)
}

func EncodeComputeCluster_HaHostIsolationResponse(p ComputeClusterParameters, vals map[string]cty.Value) {
	vals["ha_host_isolation_response"] = cty.StringVal(p.HaHostIsolationResponse)
}

func EncodeComputeCluster_HaHostMonitoring(p ComputeClusterParameters, vals map[string]cty.Value) {
	vals["ha_host_monitoring"] = cty.StringVal(p.HaHostMonitoring)
}

func EncodeComputeCluster_HaVmComponentProtection(p ComputeClusterParameters, vals map[string]cty.Value) {
	vals["ha_vm_component_protection"] = cty.StringVal(p.HaVmComponentProtection)
}

func EncodeComputeCluster_HaVmDependencyRestartCondition(p ComputeClusterParameters, vals map[string]cty.Value) {
	vals["ha_vm_dependency_restart_condition"] = cty.StringVal(p.HaVmDependencyRestartCondition)
}

func EncodeComputeCluster_HaVmFailureInterval(p ComputeClusterParameters, vals map[string]cty.Value) {
	vals["ha_vm_failure_interval"] = cty.NumberIntVal(p.HaVmFailureInterval)
}

func EncodeComputeCluster_HaVmMaximumFailureWindow(p ComputeClusterParameters, vals map[string]cty.Value) {
	vals["ha_vm_maximum_failure_window"] = cty.NumberIntVal(p.HaVmMaximumFailureWindow)
}

func EncodeComputeCluster_HaVmMaximumResets(p ComputeClusterParameters, vals map[string]cty.Value) {
	vals["ha_vm_maximum_resets"] = cty.NumberIntVal(p.HaVmMaximumResets)
}

func EncodeComputeCluster_HaVmMinimumUptime(p ComputeClusterParameters, vals map[string]cty.Value) {
	vals["ha_vm_minimum_uptime"] = cty.NumberIntVal(p.HaVmMinimumUptime)
}

func EncodeComputeCluster_HaVmMonitoring(p ComputeClusterParameters, vals map[string]cty.Value) {
	vals["ha_vm_monitoring"] = cty.StringVal(p.HaVmMonitoring)
}

func EncodeComputeCluster_HaVmRestartAdditionalDelay(p ComputeClusterParameters, vals map[string]cty.Value) {
	vals["ha_vm_restart_additional_delay"] = cty.NumberIntVal(p.HaVmRestartAdditionalDelay)
}

func EncodeComputeCluster_HaVmRestartPriority(p ComputeClusterParameters, vals map[string]cty.Value) {
	vals["ha_vm_restart_priority"] = cty.StringVal(p.HaVmRestartPriority)
}

func EncodeComputeCluster_HaVmRestartTimeout(p ComputeClusterParameters, vals map[string]cty.Value) {
	vals["ha_vm_restart_timeout"] = cty.NumberIntVal(p.HaVmRestartTimeout)
}

func EncodeComputeCluster_HostClusterExitTimeout(p ComputeClusterParameters, vals map[string]cty.Value) {
	vals["host_cluster_exit_timeout"] = cty.NumberIntVal(p.HostClusterExitTimeout)
}

func EncodeComputeCluster_HostManaged(p ComputeClusterParameters, vals map[string]cty.Value) {
	vals["host_managed"] = cty.BoolVal(p.HostManaged)
}

func EncodeComputeCluster_HostSystemIds(p ComputeClusterParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.HostSystemIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	if len(colVals) == 0 {
		vals["host_system_ids"] = cty.SetValEmpty(cty.String)
	} else {
		vals["host_system_ids"] = cty.SetVal(colVals)
    }
}

func EncodeComputeCluster_Name(p ComputeClusterParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeComputeCluster_ProactiveHaAutomationLevel(p ComputeClusterParameters, vals map[string]cty.Value) {
	vals["proactive_ha_automation_level"] = cty.StringVal(p.ProactiveHaAutomationLevel)
}

func EncodeComputeCluster_ProactiveHaEnabled(p ComputeClusterParameters, vals map[string]cty.Value) {
	vals["proactive_ha_enabled"] = cty.BoolVal(p.ProactiveHaEnabled)
}

func EncodeComputeCluster_ProactiveHaModerateRemediation(p ComputeClusterParameters, vals map[string]cty.Value) {
	vals["proactive_ha_moderate_remediation"] = cty.StringVal(p.ProactiveHaModerateRemediation)
}

func EncodeComputeCluster_ProactiveHaProviderIds(p ComputeClusterParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.ProactiveHaProviderIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	if len(colVals) == 0 {
		vals["proactive_ha_provider_ids"] = cty.SetValEmpty(cty.String)
	} else {
		vals["proactive_ha_provider_ids"] = cty.SetVal(colVals)
    }
}

func EncodeComputeCluster_ProactiveHaSevereRemediation(p ComputeClusterParameters, vals map[string]cty.Value) {
	vals["proactive_ha_severe_remediation"] = cty.StringVal(p.ProactiveHaSevereRemediation)
}

func EncodeComputeCluster_Tags(p ComputeClusterParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Tags {
		colVals = append(colVals, cty.StringVal(value))
	}
	if len(colVals) == 0 {
		vals["tags"] = cty.SetValEmpty(cty.String)
	} else {
		vals["tags"] = cty.SetVal(colVals)
    }
}

func EncodeComputeCluster_VsanEnabled(p ComputeClusterParameters, vals map[string]cty.Value) {
	vals["vsan_enabled"] = cty.BoolVal(p.VsanEnabled)
}

func EncodeComputeCluster_VsanDiskGroup(p VsanDiskGroup, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeComputeCluster_VsanDiskGroup_Cache(p, ctyVal)
	EncodeComputeCluster_VsanDiskGroup_Storage(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
    if len(valsForCollection) == 0 {
		vals["vsan_disk_group"] = cty.ListValEmpty(cty.EmptyObject)
    } else {
		vals["vsan_disk_group"] = cty.ListVal(valsForCollection)
    }
}

func EncodeComputeCluster_VsanDiskGroup_Cache(p VsanDiskGroup, vals map[string]cty.Value) {
	vals["cache"] = cty.StringVal(p.Cache)
}

func EncodeComputeCluster_VsanDiskGroup_Storage(p VsanDiskGroup, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Storage {
		colVals = append(colVals, cty.StringVal(value))
	}
	if len(colVals) == 0 {
		vals["storage"] = cty.SetValEmpty(cty.String)
	} else {
		vals["storage"] = cty.SetVal(colVals)
    }
}

func EncodeComputeCluster_ResourcePoolId(p ComputeClusterObservation, vals map[string]cty.Value) {
	vals["resource_pool_id"] = cty.StringVal(p.ResourcePoolId)
}
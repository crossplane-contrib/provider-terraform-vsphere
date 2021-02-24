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
	r, ok := mr.(*HaVmOverride)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a HaVmOverride.")
	}
	return EncodeHaVmOverride(*r), nil
}

func EncodeHaVmOverride(r HaVmOverride) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeHaVmOverride_ComputeClusterId(r.Spec.ForProvider, ctyVal)
	EncodeHaVmOverride_HaDatastoreApdRecoveryAction(r.Spec.ForProvider, ctyVal)
	EncodeHaVmOverride_HaDatastoreApdResponse(r.Spec.ForProvider, ctyVal)
	EncodeHaVmOverride_HaDatastoreApdResponseDelay(r.Spec.ForProvider, ctyVal)
	EncodeHaVmOverride_HaDatastorePdlResponse(r.Spec.ForProvider, ctyVal)
	EncodeHaVmOverride_HaHostIsolationResponse(r.Spec.ForProvider, ctyVal)
	EncodeHaVmOverride_HaVmFailureInterval(r.Spec.ForProvider, ctyVal)
	EncodeHaVmOverride_HaVmMaximumFailureWindow(r.Spec.ForProvider, ctyVal)
	EncodeHaVmOverride_HaVmMaximumResets(r.Spec.ForProvider, ctyVal)
	EncodeHaVmOverride_HaVmMinimumUptime(r.Spec.ForProvider, ctyVal)
	EncodeHaVmOverride_HaVmMonitoring(r.Spec.ForProvider, ctyVal)
	EncodeHaVmOverride_HaVmMonitoringUseClusterDefaults(r.Spec.ForProvider, ctyVal)
	EncodeHaVmOverride_HaVmRestartPriority(r.Spec.ForProvider, ctyVal)
	EncodeHaVmOverride_HaVmRestartTimeout(r.Spec.ForProvider, ctyVal)
	EncodeHaVmOverride_VirtualMachineId(r.Spec.ForProvider, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeHaVmOverride_ComputeClusterId(p HaVmOverrideParameters, vals map[string]cty.Value) {
	vals["compute_cluster_id"] = cty.StringVal(p.ComputeClusterId)
}

func EncodeHaVmOverride_HaDatastoreApdRecoveryAction(p HaVmOverrideParameters, vals map[string]cty.Value) {
	vals["ha_datastore_apd_recovery_action"] = cty.StringVal(p.HaDatastoreApdRecoveryAction)
}

func EncodeHaVmOverride_HaDatastoreApdResponse(p HaVmOverrideParameters, vals map[string]cty.Value) {
	vals["ha_datastore_apd_response"] = cty.StringVal(p.HaDatastoreApdResponse)
}

func EncodeHaVmOverride_HaDatastoreApdResponseDelay(p HaVmOverrideParameters, vals map[string]cty.Value) {
	vals["ha_datastore_apd_response_delay"] = cty.NumberIntVal(p.HaDatastoreApdResponseDelay)
}

func EncodeHaVmOverride_HaDatastorePdlResponse(p HaVmOverrideParameters, vals map[string]cty.Value) {
	vals["ha_datastore_pdl_response"] = cty.StringVal(p.HaDatastorePdlResponse)
}

func EncodeHaVmOverride_HaHostIsolationResponse(p HaVmOverrideParameters, vals map[string]cty.Value) {
	vals["ha_host_isolation_response"] = cty.StringVal(p.HaHostIsolationResponse)
}

func EncodeHaVmOverride_HaVmFailureInterval(p HaVmOverrideParameters, vals map[string]cty.Value) {
	vals["ha_vm_failure_interval"] = cty.NumberIntVal(p.HaVmFailureInterval)
}

func EncodeHaVmOverride_HaVmMaximumFailureWindow(p HaVmOverrideParameters, vals map[string]cty.Value) {
	vals["ha_vm_maximum_failure_window"] = cty.NumberIntVal(p.HaVmMaximumFailureWindow)
}

func EncodeHaVmOverride_HaVmMaximumResets(p HaVmOverrideParameters, vals map[string]cty.Value) {
	vals["ha_vm_maximum_resets"] = cty.NumberIntVal(p.HaVmMaximumResets)
}

func EncodeHaVmOverride_HaVmMinimumUptime(p HaVmOverrideParameters, vals map[string]cty.Value) {
	vals["ha_vm_minimum_uptime"] = cty.NumberIntVal(p.HaVmMinimumUptime)
}

func EncodeHaVmOverride_HaVmMonitoring(p HaVmOverrideParameters, vals map[string]cty.Value) {
	vals["ha_vm_monitoring"] = cty.StringVal(p.HaVmMonitoring)
}

func EncodeHaVmOverride_HaVmMonitoringUseClusterDefaults(p HaVmOverrideParameters, vals map[string]cty.Value) {
	vals["ha_vm_monitoring_use_cluster_defaults"] = cty.BoolVal(p.HaVmMonitoringUseClusterDefaults)
}

func EncodeHaVmOverride_HaVmRestartPriority(p HaVmOverrideParameters, vals map[string]cty.Value) {
	vals["ha_vm_restart_priority"] = cty.StringVal(p.HaVmRestartPriority)
}

func EncodeHaVmOverride_HaVmRestartTimeout(p HaVmOverrideParameters, vals map[string]cty.Value) {
	vals["ha_vm_restart_timeout"] = cty.NumberIntVal(p.HaVmRestartTimeout)
}

func EncodeHaVmOverride_VirtualMachineId(p HaVmOverrideParameters, vals map[string]cty.Value) {
	vals["virtual_machine_id"] = cty.StringVal(p.VirtualMachineId)
}
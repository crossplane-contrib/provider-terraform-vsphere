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

	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
	"github.com/zclconf/go-cty/cty"
	ctwhy "github.com/crossplane-contrib/terraform-runtime/pkg/plugin/cty"
)

type ctyDecoder struct{}

func (e *ctyDecoder) DecodeCty(mr resource.Managed, ctyValue cty.Value, schema *providers.Schema) (resource.Managed, error) {
	r, ok := mr.(*HaVmOverride)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeHaVmOverride(r, ctyValue)
}

func DecodeHaVmOverride(prev *HaVmOverride, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeHaVmOverride_HaVmMaximumResets(&new.Spec.ForProvider, valMap)
	DecodeHaVmOverride_ComputeClusterId(&new.Spec.ForProvider, valMap)
	DecodeHaVmOverride_HaDatastoreApdResponseDelay(&new.Spec.ForProvider, valMap)
	DecodeHaVmOverride_HaDatastoreApdRecoveryAction(&new.Spec.ForProvider, valMap)
	DecodeHaVmOverride_HaVmMaximumFailureWindow(&new.Spec.ForProvider, valMap)
	DecodeHaVmOverride_HaVmMonitoringUseClusterDefaults(&new.Spec.ForProvider, valMap)
	DecodeHaVmOverride_HaVmRestartTimeout(&new.Spec.ForProvider, valMap)
	DecodeHaVmOverride_VirtualMachineId(&new.Spec.ForProvider, valMap)
	DecodeHaVmOverride_HaDatastorePdlResponse(&new.Spec.ForProvider, valMap)
	DecodeHaVmOverride_HaHostIsolationResponse(&new.Spec.ForProvider, valMap)
	DecodeHaVmOverride_HaVmFailureInterval(&new.Spec.ForProvider, valMap)
	DecodeHaVmOverride_HaVmMinimumUptime(&new.Spec.ForProvider, valMap)
	DecodeHaVmOverride_HaVmMonitoring(&new.Spec.ForProvider, valMap)
	DecodeHaVmOverride_HaVmRestartPriority(&new.Spec.ForProvider, valMap)
	DecodeHaVmOverride_HaDatastoreApdResponse(&new.Spec.ForProvider, valMap)

	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeHaVmOverride_HaVmMaximumResets(p *HaVmOverrideParameters, vals map[string]cty.Value) {
	p.HaVmMaximumResets = ctwhy.ValueAsInt64(vals["ha_vm_maximum_resets"])
}

//primitiveTypeDecodeTemplate
func DecodeHaVmOverride_ComputeClusterId(p *HaVmOverrideParameters, vals map[string]cty.Value) {
	p.ComputeClusterId = ctwhy.ValueAsString(vals["compute_cluster_id"])
}

//primitiveTypeDecodeTemplate
func DecodeHaVmOverride_HaDatastoreApdResponseDelay(p *HaVmOverrideParameters, vals map[string]cty.Value) {
	p.HaDatastoreApdResponseDelay = ctwhy.ValueAsInt64(vals["ha_datastore_apd_response_delay"])
}

//primitiveTypeDecodeTemplate
func DecodeHaVmOverride_HaDatastoreApdRecoveryAction(p *HaVmOverrideParameters, vals map[string]cty.Value) {
	p.HaDatastoreApdRecoveryAction = ctwhy.ValueAsString(vals["ha_datastore_apd_recovery_action"])
}

//primitiveTypeDecodeTemplate
func DecodeHaVmOverride_HaVmMaximumFailureWindow(p *HaVmOverrideParameters, vals map[string]cty.Value) {
	p.HaVmMaximumFailureWindow = ctwhy.ValueAsInt64(vals["ha_vm_maximum_failure_window"])
}

//primitiveTypeDecodeTemplate
func DecodeHaVmOverride_HaVmMonitoringUseClusterDefaults(p *HaVmOverrideParameters, vals map[string]cty.Value) {
	p.HaVmMonitoringUseClusterDefaults = ctwhy.ValueAsBool(vals["ha_vm_monitoring_use_cluster_defaults"])
}

//primitiveTypeDecodeTemplate
func DecodeHaVmOverride_HaVmRestartTimeout(p *HaVmOverrideParameters, vals map[string]cty.Value) {
	p.HaVmRestartTimeout = ctwhy.ValueAsInt64(vals["ha_vm_restart_timeout"])
}

//primitiveTypeDecodeTemplate
func DecodeHaVmOverride_VirtualMachineId(p *HaVmOverrideParameters, vals map[string]cty.Value) {
	p.VirtualMachineId = ctwhy.ValueAsString(vals["virtual_machine_id"])
}

//primitiveTypeDecodeTemplate
func DecodeHaVmOverride_HaDatastorePdlResponse(p *HaVmOverrideParameters, vals map[string]cty.Value) {
	p.HaDatastorePdlResponse = ctwhy.ValueAsString(vals["ha_datastore_pdl_response"])
}

//primitiveTypeDecodeTemplate
func DecodeHaVmOverride_HaHostIsolationResponse(p *HaVmOverrideParameters, vals map[string]cty.Value) {
	p.HaHostIsolationResponse = ctwhy.ValueAsString(vals["ha_host_isolation_response"])
}

//primitiveTypeDecodeTemplate
func DecodeHaVmOverride_HaVmFailureInterval(p *HaVmOverrideParameters, vals map[string]cty.Value) {
	p.HaVmFailureInterval = ctwhy.ValueAsInt64(vals["ha_vm_failure_interval"])
}

//primitiveTypeDecodeTemplate
func DecodeHaVmOverride_HaVmMinimumUptime(p *HaVmOverrideParameters, vals map[string]cty.Value) {
	p.HaVmMinimumUptime = ctwhy.ValueAsInt64(vals["ha_vm_minimum_uptime"])
}

//primitiveTypeDecodeTemplate
func DecodeHaVmOverride_HaVmMonitoring(p *HaVmOverrideParameters, vals map[string]cty.Value) {
	p.HaVmMonitoring = ctwhy.ValueAsString(vals["ha_vm_monitoring"])
}

//primitiveTypeDecodeTemplate
func DecodeHaVmOverride_HaVmRestartPriority(p *HaVmOverrideParameters, vals map[string]cty.Value) {
	p.HaVmRestartPriority = ctwhy.ValueAsString(vals["ha_vm_restart_priority"])
}

//primitiveTypeDecodeTemplate
func DecodeHaVmOverride_HaDatastoreApdResponse(p *HaVmOverrideParameters, vals map[string]cty.Value) {
	p.HaDatastoreApdResponse = ctwhy.ValueAsString(vals["ha_datastore_apd_response"])
}
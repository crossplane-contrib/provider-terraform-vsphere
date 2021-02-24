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
	r, ok := mr.(*DatastoreCluster)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a DatastoreCluster.")
	}
	return EncodeDatastoreCluster(*r), nil
}

func EncodeDatastoreCluster(r DatastoreCluster) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDatastoreCluster_CustomAttributes(r.Spec.ForProvider, ctyVal)
	EncodeDatastoreCluster_DatacenterId(r.Spec.ForProvider, ctyVal)
	EncodeDatastoreCluster_Folder(r.Spec.ForProvider, ctyVal)
	EncodeDatastoreCluster_Name(r.Spec.ForProvider, ctyVal)
	EncodeDatastoreCluster_SdrsAdvancedOptions(r.Spec.ForProvider, ctyVal)
	EncodeDatastoreCluster_SdrsAutomationLevel(r.Spec.ForProvider, ctyVal)
	EncodeDatastoreCluster_SdrsDefaultIntraVmAffinity(r.Spec.ForProvider, ctyVal)
	EncodeDatastoreCluster_SdrsEnabled(r.Spec.ForProvider, ctyVal)
	EncodeDatastoreCluster_SdrsFreeSpaceThreshold(r.Spec.ForProvider, ctyVal)
	EncodeDatastoreCluster_SdrsFreeSpaceThresholdMode(r.Spec.ForProvider, ctyVal)
	EncodeDatastoreCluster_SdrsFreeSpaceUtilizationDifference(r.Spec.ForProvider, ctyVal)
	EncodeDatastoreCluster_SdrsIoBalanceAutomationLevel(r.Spec.ForProvider, ctyVal)
	EncodeDatastoreCluster_SdrsIoLatencyThreshold(r.Spec.ForProvider, ctyVal)
	EncodeDatastoreCluster_SdrsIoLoadBalanceEnabled(r.Spec.ForProvider, ctyVal)
	EncodeDatastoreCluster_SdrsIoLoadImbalanceThreshold(r.Spec.ForProvider, ctyVal)
	EncodeDatastoreCluster_SdrsIoReservableIopsThreshold(r.Spec.ForProvider, ctyVal)
	EncodeDatastoreCluster_SdrsIoReservablePercentThreshold(r.Spec.ForProvider, ctyVal)
	EncodeDatastoreCluster_SdrsIoReservableThresholdMode(r.Spec.ForProvider, ctyVal)
	EncodeDatastoreCluster_SdrsLoadBalanceInterval(r.Spec.ForProvider, ctyVal)
	EncodeDatastoreCluster_SdrsPolicyEnforcementAutomationLevel(r.Spec.ForProvider, ctyVal)
	EncodeDatastoreCluster_SdrsRuleEnforcementAutomationLevel(r.Spec.ForProvider, ctyVal)
	EncodeDatastoreCluster_SdrsSpaceBalanceAutomationLevel(r.Spec.ForProvider, ctyVal)
	EncodeDatastoreCluster_SdrsSpaceUtilizationThreshold(r.Spec.ForProvider, ctyVal)
	EncodeDatastoreCluster_SdrsVmEvacuationAutomationLevel(r.Spec.ForProvider, ctyVal)
	EncodeDatastoreCluster_Tags(r.Spec.ForProvider, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeDatastoreCluster_CustomAttributes(p DatastoreClusterParameters, vals map[string]cty.Value) {
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

func EncodeDatastoreCluster_DatacenterId(p DatastoreClusterParameters, vals map[string]cty.Value) {
	vals["datacenter_id"] = cty.StringVal(p.DatacenterId)
}

func EncodeDatastoreCluster_Folder(p DatastoreClusterParameters, vals map[string]cty.Value) {
	vals["folder"] = cty.StringVal(p.Folder)
}

func EncodeDatastoreCluster_Name(p DatastoreClusterParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeDatastoreCluster_SdrsAdvancedOptions(p DatastoreClusterParameters, vals map[string]cty.Value) {
	if len(p.SdrsAdvancedOptions) == 0 {
		vals["sdrs_advanced_options"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.SdrsAdvancedOptions {
		mVals[key] = cty.StringVal(value)
	}
	vals["sdrs_advanced_options"] = cty.MapVal(mVals)
}

func EncodeDatastoreCluster_SdrsAutomationLevel(p DatastoreClusterParameters, vals map[string]cty.Value) {
	vals["sdrs_automation_level"] = cty.StringVal(p.SdrsAutomationLevel)
}

func EncodeDatastoreCluster_SdrsDefaultIntraVmAffinity(p DatastoreClusterParameters, vals map[string]cty.Value) {
	vals["sdrs_default_intra_vm_affinity"] = cty.BoolVal(p.SdrsDefaultIntraVmAffinity)
}

func EncodeDatastoreCluster_SdrsEnabled(p DatastoreClusterParameters, vals map[string]cty.Value) {
	vals["sdrs_enabled"] = cty.BoolVal(p.SdrsEnabled)
}

func EncodeDatastoreCluster_SdrsFreeSpaceThreshold(p DatastoreClusterParameters, vals map[string]cty.Value) {
	vals["sdrs_free_space_threshold"] = cty.NumberIntVal(p.SdrsFreeSpaceThreshold)
}

func EncodeDatastoreCluster_SdrsFreeSpaceThresholdMode(p DatastoreClusterParameters, vals map[string]cty.Value) {
	vals["sdrs_free_space_threshold_mode"] = cty.StringVal(p.SdrsFreeSpaceThresholdMode)
}

func EncodeDatastoreCluster_SdrsFreeSpaceUtilizationDifference(p DatastoreClusterParameters, vals map[string]cty.Value) {
	vals["sdrs_free_space_utilization_difference"] = cty.NumberIntVal(p.SdrsFreeSpaceUtilizationDifference)
}

func EncodeDatastoreCluster_SdrsIoBalanceAutomationLevel(p DatastoreClusterParameters, vals map[string]cty.Value) {
	vals["sdrs_io_balance_automation_level"] = cty.StringVal(p.SdrsIoBalanceAutomationLevel)
}

func EncodeDatastoreCluster_SdrsIoLatencyThreshold(p DatastoreClusterParameters, vals map[string]cty.Value) {
	vals["sdrs_io_latency_threshold"] = cty.NumberIntVal(p.SdrsIoLatencyThreshold)
}

func EncodeDatastoreCluster_SdrsIoLoadBalanceEnabled(p DatastoreClusterParameters, vals map[string]cty.Value) {
	vals["sdrs_io_load_balance_enabled"] = cty.BoolVal(p.SdrsIoLoadBalanceEnabled)
}

func EncodeDatastoreCluster_SdrsIoLoadImbalanceThreshold(p DatastoreClusterParameters, vals map[string]cty.Value) {
	vals["sdrs_io_load_imbalance_threshold"] = cty.NumberIntVal(p.SdrsIoLoadImbalanceThreshold)
}

func EncodeDatastoreCluster_SdrsIoReservableIopsThreshold(p DatastoreClusterParameters, vals map[string]cty.Value) {
	vals["sdrs_io_reservable_iops_threshold"] = cty.NumberIntVal(p.SdrsIoReservableIopsThreshold)
}

func EncodeDatastoreCluster_SdrsIoReservablePercentThreshold(p DatastoreClusterParameters, vals map[string]cty.Value) {
	vals["sdrs_io_reservable_percent_threshold"] = cty.NumberIntVal(p.SdrsIoReservablePercentThreshold)
}

func EncodeDatastoreCluster_SdrsIoReservableThresholdMode(p DatastoreClusterParameters, vals map[string]cty.Value) {
	vals["sdrs_io_reservable_threshold_mode"] = cty.StringVal(p.SdrsIoReservableThresholdMode)
}

func EncodeDatastoreCluster_SdrsLoadBalanceInterval(p DatastoreClusterParameters, vals map[string]cty.Value) {
	vals["sdrs_load_balance_interval"] = cty.NumberIntVal(p.SdrsLoadBalanceInterval)
}

func EncodeDatastoreCluster_SdrsPolicyEnforcementAutomationLevel(p DatastoreClusterParameters, vals map[string]cty.Value) {
	vals["sdrs_policy_enforcement_automation_level"] = cty.StringVal(p.SdrsPolicyEnforcementAutomationLevel)
}

func EncodeDatastoreCluster_SdrsRuleEnforcementAutomationLevel(p DatastoreClusterParameters, vals map[string]cty.Value) {
	vals["sdrs_rule_enforcement_automation_level"] = cty.StringVal(p.SdrsRuleEnforcementAutomationLevel)
}

func EncodeDatastoreCluster_SdrsSpaceBalanceAutomationLevel(p DatastoreClusterParameters, vals map[string]cty.Value) {
	vals["sdrs_space_balance_automation_level"] = cty.StringVal(p.SdrsSpaceBalanceAutomationLevel)
}

func EncodeDatastoreCluster_SdrsSpaceUtilizationThreshold(p DatastoreClusterParameters, vals map[string]cty.Value) {
	vals["sdrs_space_utilization_threshold"] = cty.NumberIntVal(p.SdrsSpaceUtilizationThreshold)
}

func EncodeDatastoreCluster_SdrsVmEvacuationAutomationLevel(p DatastoreClusterParameters, vals map[string]cty.Value) {
	vals["sdrs_vm_evacuation_automation_level"] = cty.StringVal(p.SdrsVmEvacuationAutomationLevel)
}

func EncodeDatastoreCluster_Tags(p DatastoreClusterParameters, vals map[string]cty.Value) {
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
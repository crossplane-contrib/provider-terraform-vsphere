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
	r, ok := mr.(*DatastoreCluster)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeDatastoreCluster(r, ctyValue)
}

func DecodeDatastoreCluster(prev *DatastoreCluster, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeDatastoreCluster_CustomAttributes(&new.Spec.ForProvider, valMap)
	DecodeDatastoreCluster_DatacenterId(&new.Spec.ForProvider, valMap)
	DecodeDatastoreCluster_Folder(&new.Spec.ForProvider, valMap)
	DecodeDatastoreCluster_Name(&new.Spec.ForProvider, valMap)
	DecodeDatastoreCluster_SdrsAdvancedOptions(&new.Spec.ForProvider, valMap)
	DecodeDatastoreCluster_SdrsAutomationLevel(&new.Spec.ForProvider, valMap)
	DecodeDatastoreCluster_SdrsDefaultIntraVmAffinity(&new.Spec.ForProvider, valMap)
	DecodeDatastoreCluster_SdrsEnabled(&new.Spec.ForProvider, valMap)
	DecodeDatastoreCluster_SdrsFreeSpaceThreshold(&new.Spec.ForProvider, valMap)
	DecodeDatastoreCluster_SdrsFreeSpaceThresholdMode(&new.Spec.ForProvider, valMap)
	DecodeDatastoreCluster_SdrsFreeSpaceUtilizationDifference(&new.Spec.ForProvider, valMap)
	DecodeDatastoreCluster_SdrsIoBalanceAutomationLevel(&new.Spec.ForProvider, valMap)
	DecodeDatastoreCluster_SdrsIoLatencyThreshold(&new.Spec.ForProvider, valMap)
	DecodeDatastoreCluster_SdrsIoLoadBalanceEnabled(&new.Spec.ForProvider, valMap)
	DecodeDatastoreCluster_SdrsIoLoadImbalanceThreshold(&new.Spec.ForProvider, valMap)
	DecodeDatastoreCluster_SdrsIoReservableIopsThreshold(&new.Spec.ForProvider, valMap)
	DecodeDatastoreCluster_SdrsIoReservablePercentThreshold(&new.Spec.ForProvider, valMap)
	DecodeDatastoreCluster_SdrsIoReservableThresholdMode(&new.Spec.ForProvider, valMap)
	DecodeDatastoreCluster_SdrsLoadBalanceInterval(&new.Spec.ForProvider, valMap)
	DecodeDatastoreCluster_SdrsPolicyEnforcementAutomationLevel(&new.Spec.ForProvider, valMap)
	DecodeDatastoreCluster_SdrsRuleEnforcementAutomationLevel(&new.Spec.ForProvider, valMap)
	DecodeDatastoreCluster_SdrsSpaceBalanceAutomationLevel(&new.Spec.ForProvider, valMap)
	DecodeDatastoreCluster_SdrsSpaceUtilizationThreshold(&new.Spec.ForProvider, valMap)
	DecodeDatastoreCluster_SdrsVmEvacuationAutomationLevel(&new.Spec.ForProvider, valMap)
	DecodeDatastoreCluster_Tags(&new.Spec.ForProvider, valMap)

	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveMapTypeDecodeTemplate
func DecodeDatastoreCluster_CustomAttributes(p *DatastoreClusterParameters, vals map[string]cty.Value) {
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
func DecodeDatastoreCluster_DatacenterId(p *DatastoreClusterParameters, vals map[string]cty.Value) {
	p.DatacenterId = ctwhy.ValueAsString(vals["datacenter_id"])
}

//primitiveTypeDecodeTemplate
func DecodeDatastoreCluster_Folder(p *DatastoreClusterParameters, vals map[string]cty.Value) {
	p.Folder = ctwhy.ValueAsString(vals["folder"])
}

//primitiveTypeDecodeTemplate
func DecodeDatastoreCluster_Name(p *DatastoreClusterParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveMapTypeDecodeTemplate
func DecodeDatastoreCluster_SdrsAdvancedOptions(p *DatastoreClusterParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	if vals["sdrs_advanced_options"].IsNull() {
		p.SdrsAdvancedOptions = nil
		return
	}
	vMap := make(map[string]string)
	v := vals["sdrs_advanced_options"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.SdrsAdvancedOptions = vMap
}

//primitiveTypeDecodeTemplate
func DecodeDatastoreCluster_SdrsAutomationLevel(p *DatastoreClusterParameters, vals map[string]cty.Value) {
	p.SdrsAutomationLevel = ctwhy.ValueAsString(vals["sdrs_automation_level"])
}

//primitiveTypeDecodeTemplate
func DecodeDatastoreCluster_SdrsDefaultIntraVmAffinity(p *DatastoreClusterParameters, vals map[string]cty.Value) {
	p.SdrsDefaultIntraVmAffinity = ctwhy.ValueAsBool(vals["sdrs_default_intra_vm_affinity"])
}

//primitiveTypeDecodeTemplate
func DecodeDatastoreCluster_SdrsEnabled(p *DatastoreClusterParameters, vals map[string]cty.Value) {
	p.SdrsEnabled = ctwhy.ValueAsBool(vals["sdrs_enabled"])
}

//primitiveTypeDecodeTemplate
func DecodeDatastoreCluster_SdrsFreeSpaceThreshold(p *DatastoreClusterParameters, vals map[string]cty.Value) {
	p.SdrsFreeSpaceThreshold = ctwhy.ValueAsInt64(vals["sdrs_free_space_threshold"])
}

//primitiveTypeDecodeTemplate
func DecodeDatastoreCluster_SdrsFreeSpaceThresholdMode(p *DatastoreClusterParameters, vals map[string]cty.Value) {
	p.SdrsFreeSpaceThresholdMode = ctwhy.ValueAsString(vals["sdrs_free_space_threshold_mode"])
}

//primitiveTypeDecodeTemplate
func DecodeDatastoreCluster_SdrsFreeSpaceUtilizationDifference(p *DatastoreClusterParameters, vals map[string]cty.Value) {
	p.SdrsFreeSpaceUtilizationDifference = ctwhy.ValueAsInt64(vals["sdrs_free_space_utilization_difference"])
}

//primitiveTypeDecodeTemplate
func DecodeDatastoreCluster_SdrsIoBalanceAutomationLevel(p *DatastoreClusterParameters, vals map[string]cty.Value) {
	p.SdrsIoBalanceAutomationLevel = ctwhy.ValueAsString(vals["sdrs_io_balance_automation_level"])
}

//primitiveTypeDecodeTemplate
func DecodeDatastoreCluster_SdrsIoLatencyThreshold(p *DatastoreClusterParameters, vals map[string]cty.Value) {
	p.SdrsIoLatencyThreshold = ctwhy.ValueAsInt64(vals["sdrs_io_latency_threshold"])
}

//primitiveTypeDecodeTemplate
func DecodeDatastoreCluster_SdrsIoLoadBalanceEnabled(p *DatastoreClusterParameters, vals map[string]cty.Value) {
	p.SdrsIoLoadBalanceEnabled = ctwhy.ValueAsBool(vals["sdrs_io_load_balance_enabled"])
}

//primitiveTypeDecodeTemplate
func DecodeDatastoreCluster_SdrsIoLoadImbalanceThreshold(p *DatastoreClusterParameters, vals map[string]cty.Value) {
	p.SdrsIoLoadImbalanceThreshold = ctwhy.ValueAsInt64(vals["sdrs_io_load_imbalance_threshold"])
}

//primitiveTypeDecodeTemplate
func DecodeDatastoreCluster_SdrsIoReservableIopsThreshold(p *DatastoreClusterParameters, vals map[string]cty.Value) {
	p.SdrsIoReservableIopsThreshold = ctwhy.ValueAsInt64(vals["sdrs_io_reservable_iops_threshold"])
}

//primitiveTypeDecodeTemplate
func DecodeDatastoreCluster_SdrsIoReservablePercentThreshold(p *DatastoreClusterParameters, vals map[string]cty.Value) {
	p.SdrsIoReservablePercentThreshold = ctwhy.ValueAsInt64(vals["sdrs_io_reservable_percent_threshold"])
}

//primitiveTypeDecodeTemplate
func DecodeDatastoreCluster_SdrsIoReservableThresholdMode(p *DatastoreClusterParameters, vals map[string]cty.Value) {
	p.SdrsIoReservableThresholdMode = ctwhy.ValueAsString(vals["sdrs_io_reservable_threshold_mode"])
}

//primitiveTypeDecodeTemplate
func DecodeDatastoreCluster_SdrsLoadBalanceInterval(p *DatastoreClusterParameters, vals map[string]cty.Value) {
	p.SdrsLoadBalanceInterval = ctwhy.ValueAsInt64(vals["sdrs_load_balance_interval"])
}

//primitiveTypeDecodeTemplate
func DecodeDatastoreCluster_SdrsPolicyEnforcementAutomationLevel(p *DatastoreClusterParameters, vals map[string]cty.Value) {
	p.SdrsPolicyEnforcementAutomationLevel = ctwhy.ValueAsString(vals["sdrs_policy_enforcement_automation_level"])
}

//primitiveTypeDecodeTemplate
func DecodeDatastoreCluster_SdrsRuleEnforcementAutomationLevel(p *DatastoreClusterParameters, vals map[string]cty.Value) {
	p.SdrsRuleEnforcementAutomationLevel = ctwhy.ValueAsString(vals["sdrs_rule_enforcement_automation_level"])
}

//primitiveTypeDecodeTemplate
func DecodeDatastoreCluster_SdrsSpaceBalanceAutomationLevel(p *DatastoreClusterParameters, vals map[string]cty.Value) {
	p.SdrsSpaceBalanceAutomationLevel = ctwhy.ValueAsString(vals["sdrs_space_balance_automation_level"])
}

//primitiveTypeDecodeTemplate
func DecodeDatastoreCluster_SdrsSpaceUtilizationThreshold(p *DatastoreClusterParameters, vals map[string]cty.Value) {
	p.SdrsSpaceUtilizationThreshold = ctwhy.ValueAsInt64(vals["sdrs_space_utilization_threshold"])
}

//primitiveTypeDecodeTemplate
func DecodeDatastoreCluster_SdrsVmEvacuationAutomationLevel(p *DatastoreClusterParameters, vals map[string]cty.Value) {
	p.SdrsVmEvacuationAutomationLevel = ctwhy.ValueAsString(vals["sdrs_vm_evacuation_automation_level"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeDatastoreCluster_Tags(p *DatastoreClusterParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["tags"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.Tags = goVals
}

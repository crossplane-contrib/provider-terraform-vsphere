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
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*ComputeClusterVmAntiAffinityRule)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a ComputeClusterVmAntiAffinityRule.")
	}
	return EncodeComputeClusterVmAntiAffinityRule(*r), nil
}

func EncodeComputeClusterVmAntiAffinityRule(r ComputeClusterVmAntiAffinityRule) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeComputeClusterVmAntiAffinityRule_ComputeClusterId(r.Spec.ForProvider, ctyVal)
	EncodeComputeClusterVmAntiAffinityRule_Enabled(r.Spec.ForProvider, ctyVal)
	EncodeComputeClusterVmAntiAffinityRule_Mandatory(r.Spec.ForProvider, ctyVal)
	EncodeComputeClusterVmAntiAffinityRule_Name(r.Spec.ForProvider, ctyVal)
	EncodeComputeClusterVmAntiAffinityRule_VirtualMachineIds(r.Spec.ForProvider, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeComputeClusterVmAntiAffinityRule_ComputeClusterId(p ComputeClusterVmAntiAffinityRuleParameters, vals map[string]cty.Value) {
	vals["compute_cluster_id"] = cty.StringVal(p.ComputeClusterId)
}

func EncodeComputeClusterVmAntiAffinityRule_Enabled(p ComputeClusterVmAntiAffinityRuleParameters, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeComputeClusterVmAntiAffinityRule_Mandatory(p ComputeClusterVmAntiAffinityRuleParameters, vals map[string]cty.Value) {
	vals["mandatory"] = cty.BoolVal(p.Mandatory)
}

func EncodeComputeClusterVmAntiAffinityRule_Name(p ComputeClusterVmAntiAffinityRuleParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeComputeClusterVmAntiAffinityRule_VirtualMachineIds(p ComputeClusterVmAntiAffinityRuleParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.VirtualMachineIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	if len(colVals) == 0 {
		vals["virtual_machine_ids"] = cty.SetValEmpty(cty.String)
	} else {
		vals["virtual_machine_ids"] = cty.SetVal(colVals)
	}
}

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
	r, ok := mr.(*DrsVmOverride)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a DrsVmOverride.")
	}
	return EncodeDrsVmOverride(*r), nil
}

func EncodeDrsVmOverride(r DrsVmOverride) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDrsVmOverride_DrsAutomationLevel(r.Spec.ForProvider, ctyVal)
	EncodeDrsVmOverride_DrsEnabled(r.Spec.ForProvider, ctyVal)
	EncodeDrsVmOverride_VirtualMachineId(r.Spec.ForProvider, ctyVal)
	EncodeDrsVmOverride_ComputeClusterId(r.Spec.ForProvider, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeDrsVmOverride_DrsAutomationLevel(p DrsVmOverrideParameters, vals map[string]cty.Value) {
	vals["drs_automation_level"] = cty.StringVal(p.DrsAutomationLevel)
}

func EncodeDrsVmOverride_DrsEnabled(p DrsVmOverrideParameters, vals map[string]cty.Value) {
	vals["drs_enabled"] = cty.BoolVal(p.DrsEnabled)
}

func EncodeDrsVmOverride_VirtualMachineId(p DrsVmOverrideParameters, vals map[string]cty.Value) {
	vals["virtual_machine_id"] = cty.StringVal(p.VirtualMachineId)
}

func EncodeDrsVmOverride_ComputeClusterId(p DrsVmOverrideParameters, vals map[string]cty.Value) {
	vals["compute_cluster_id"] = cty.StringVal(p.ComputeClusterId)
}
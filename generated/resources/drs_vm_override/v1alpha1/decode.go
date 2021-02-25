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
	r, ok := mr.(*DrsVmOverride)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeDrsVmOverride(r, ctyValue)
}

func DecodeDrsVmOverride(prev *DrsVmOverride, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeDrsVmOverride_ComputeClusterId(&new.Spec.ForProvider, valMap)
	DecodeDrsVmOverride_DrsAutomationLevel(&new.Spec.ForProvider, valMap)
	DecodeDrsVmOverride_DrsEnabled(&new.Spec.ForProvider, valMap)
	DecodeDrsVmOverride_VirtualMachineId(&new.Spec.ForProvider, valMap)

	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeDrsVmOverride_ComputeClusterId(p *DrsVmOverrideParameters, vals map[string]cty.Value) {
	p.ComputeClusterId = ctwhy.ValueAsString(vals["compute_cluster_id"])
}

//primitiveTypeDecodeTemplate
func DecodeDrsVmOverride_DrsAutomationLevel(p *DrsVmOverrideParameters, vals map[string]cty.Value) {
	p.DrsAutomationLevel = ctwhy.ValueAsString(vals["drs_automation_level"])
}

//primitiveTypeDecodeTemplate
func DecodeDrsVmOverride_DrsEnabled(p *DrsVmOverrideParameters, vals map[string]cty.Value) {
	p.DrsEnabled = ctwhy.ValueAsBool(vals["drs_enabled"])
}

//primitiveTypeDecodeTemplate
func DecodeDrsVmOverride_VirtualMachineId(p *DrsVmOverrideParameters, vals map[string]cty.Value) {
	p.VirtualMachineId = ctwhy.ValueAsString(vals["virtual_machine_id"])
}

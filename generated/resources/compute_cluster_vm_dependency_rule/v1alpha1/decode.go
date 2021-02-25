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
	r, ok := mr.(*ComputeClusterVmDependencyRule)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeComputeClusterVmDependencyRule(r, ctyValue)
}

func DecodeComputeClusterVmDependencyRule(prev *ComputeClusterVmDependencyRule, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeComputeClusterVmDependencyRule_ComputeClusterId(&new.Spec.ForProvider, valMap)
	DecodeComputeClusterVmDependencyRule_DependencyVmGroupName(&new.Spec.ForProvider, valMap)
	DecodeComputeClusterVmDependencyRule_Enabled(&new.Spec.ForProvider, valMap)
	DecodeComputeClusterVmDependencyRule_Mandatory(&new.Spec.ForProvider, valMap)
	DecodeComputeClusterVmDependencyRule_Name(&new.Spec.ForProvider, valMap)
	DecodeComputeClusterVmDependencyRule_VmGroupName(&new.Spec.ForProvider, valMap)

	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeComputeClusterVmDependencyRule_ComputeClusterId(p *ComputeClusterVmDependencyRuleParameters, vals map[string]cty.Value) {
	p.ComputeClusterId = ctwhy.ValueAsString(vals["compute_cluster_id"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeClusterVmDependencyRule_DependencyVmGroupName(p *ComputeClusterVmDependencyRuleParameters, vals map[string]cty.Value) {
	p.DependencyVmGroupName = ctwhy.ValueAsString(vals["dependency_vm_group_name"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeClusterVmDependencyRule_Enabled(p *ComputeClusterVmDependencyRuleParameters, vals map[string]cty.Value) {
	p.Enabled = ctwhy.ValueAsBool(vals["enabled"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeClusterVmDependencyRule_Mandatory(p *ComputeClusterVmDependencyRuleParameters, vals map[string]cty.Value) {
	p.Mandatory = ctwhy.ValueAsBool(vals["mandatory"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeClusterVmDependencyRule_Name(p *ComputeClusterVmDependencyRuleParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeClusterVmDependencyRule_VmGroupName(p *ComputeClusterVmDependencyRuleParameters, vals map[string]cty.Value) {
	p.VmGroupName = ctwhy.ValueAsString(vals["vm_group_name"])
}

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
	r, ok := mr.(*ComputeClusterVmHostRule)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a ComputeClusterVmHostRule.")
	}
	return EncodeComputeClusterVmHostRule(*r), nil
}

func EncodeComputeClusterVmHostRule(r ComputeClusterVmHostRule) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeComputeClusterVmHostRule_AntiAffinityHostGroupName(r.Spec.ForProvider, ctyVal)
	EncodeComputeClusterVmHostRule_ComputeClusterId(r.Spec.ForProvider, ctyVal)
	EncodeComputeClusterVmHostRule_Enabled(r.Spec.ForProvider, ctyVal)
	EncodeComputeClusterVmHostRule_Mandatory(r.Spec.ForProvider, ctyVal)
	EncodeComputeClusterVmHostRule_Name(r.Spec.ForProvider, ctyVal)
	EncodeComputeClusterVmHostRule_VmGroupName(r.Spec.ForProvider, ctyVal)
	EncodeComputeClusterVmHostRule_AffinityHostGroupName(r.Spec.ForProvider, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeComputeClusterVmHostRule_AntiAffinityHostGroupName(p ComputeClusterVmHostRuleParameters, vals map[string]cty.Value) {
	vals["anti_affinity_host_group_name"] = cty.StringVal(p.AntiAffinityHostGroupName)
}

func EncodeComputeClusterVmHostRule_ComputeClusterId(p ComputeClusterVmHostRuleParameters, vals map[string]cty.Value) {
	vals["compute_cluster_id"] = cty.StringVal(p.ComputeClusterId)
}

func EncodeComputeClusterVmHostRule_Enabled(p ComputeClusterVmHostRuleParameters, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeComputeClusterVmHostRule_Mandatory(p ComputeClusterVmHostRuleParameters, vals map[string]cty.Value) {
	vals["mandatory"] = cty.BoolVal(p.Mandatory)
}

func EncodeComputeClusterVmHostRule_Name(p ComputeClusterVmHostRuleParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeComputeClusterVmHostRule_VmGroupName(p ComputeClusterVmHostRuleParameters, vals map[string]cty.Value) {
	vals["vm_group_name"] = cty.StringVal(p.VmGroupName)
}

func EncodeComputeClusterVmHostRule_AffinityHostGroupName(p ComputeClusterVmHostRuleParameters, vals map[string]cty.Value) {
	vals["affinity_host_group_name"] = cty.StringVal(p.AffinityHostGroupName)
}
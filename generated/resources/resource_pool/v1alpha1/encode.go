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
	r, ok := mr.(*ResourcePool)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a ResourcePool.")
	}
	return EncodeResourcePool(*r), nil
}

func EncodeResourcePool(r ResourcePool) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeResourcePool_CpuExpandable(r.Spec.ForProvider, ctyVal)
	EncodeResourcePool_CpuLimit(r.Spec.ForProvider, ctyVal)
	EncodeResourcePool_CpuReservation(r.Spec.ForProvider, ctyVal)
	EncodeResourcePool_CpuShareLevel(r.Spec.ForProvider, ctyVal)
	EncodeResourcePool_CpuShares(r.Spec.ForProvider, ctyVal)
	EncodeResourcePool_CustomAttributes(r.Spec.ForProvider, ctyVal)
	EncodeResourcePool_MemoryExpandable(r.Spec.ForProvider, ctyVal)
	EncodeResourcePool_MemoryLimit(r.Spec.ForProvider, ctyVal)
	EncodeResourcePool_MemoryReservation(r.Spec.ForProvider, ctyVal)
	EncodeResourcePool_MemoryShareLevel(r.Spec.ForProvider, ctyVal)
	EncodeResourcePool_MemoryShares(r.Spec.ForProvider, ctyVal)
	EncodeResourcePool_Name(r.Spec.ForProvider, ctyVal)
	EncodeResourcePool_ParentResourcePoolId(r.Spec.ForProvider, ctyVal)
	EncodeResourcePool_Tags(r.Spec.ForProvider, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeResourcePool_CpuExpandable(p ResourcePoolParameters, vals map[string]cty.Value) {
	vals["cpu_expandable"] = cty.BoolVal(p.CpuExpandable)
}

func EncodeResourcePool_CpuLimit(p ResourcePoolParameters, vals map[string]cty.Value) {
	vals["cpu_limit"] = cty.NumberIntVal(p.CpuLimit)
}

func EncodeResourcePool_CpuReservation(p ResourcePoolParameters, vals map[string]cty.Value) {
	vals["cpu_reservation"] = cty.NumberIntVal(p.CpuReservation)
}

func EncodeResourcePool_CpuShareLevel(p ResourcePoolParameters, vals map[string]cty.Value) {
	vals["cpu_share_level"] = cty.StringVal(p.CpuShareLevel)
}

func EncodeResourcePool_CpuShares(p ResourcePoolParameters, vals map[string]cty.Value) {
	vals["cpu_shares"] = cty.NumberIntVal(p.CpuShares)
}

func EncodeResourcePool_CustomAttributes(p ResourcePoolParameters, vals map[string]cty.Value) {
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

func EncodeResourcePool_MemoryExpandable(p ResourcePoolParameters, vals map[string]cty.Value) {
	vals["memory_expandable"] = cty.BoolVal(p.MemoryExpandable)
}

func EncodeResourcePool_MemoryLimit(p ResourcePoolParameters, vals map[string]cty.Value) {
	vals["memory_limit"] = cty.NumberIntVal(p.MemoryLimit)
}

func EncodeResourcePool_MemoryReservation(p ResourcePoolParameters, vals map[string]cty.Value) {
	vals["memory_reservation"] = cty.NumberIntVal(p.MemoryReservation)
}

func EncodeResourcePool_MemoryShareLevel(p ResourcePoolParameters, vals map[string]cty.Value) {
	vals["memory_share_level"] = cty.StringVal(p.MemoryShareLevel)
}

func EncodeResourcePool_MemoryShares(p ResourcePoolParameters, vals map[string]cty.Value) {
	vals["memory_shares"] = cty.NumberIntVal(p.MemoryShares)
}

func EncodeResourcePool_Name(p ResourcePoolParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeResourcePool_ParentResourcePoolId(p ResourcePoolParameters, vals map[string]cty.Value) {
	vals["parent_resource_pool_id"] = cty.StringVal(p.ParentResourcePoolId)
}

func EncodeResourcePool_Tags(p ResourcePoolParameters, vals map[string]cty.Value) {
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
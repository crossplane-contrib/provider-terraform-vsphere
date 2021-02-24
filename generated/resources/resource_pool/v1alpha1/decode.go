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
	r, ok := mr.(*ResourcePool)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeResourcePool(r, ctyValue)
}

func DecodeResourcePool(prev *ResourcePool, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeResourcePool_CpuExpandable(&new.Spec.ForProvider, valMap)
	DecodeResourcePool_CpuLimit(&new.Spec.ForProvider, valMap)
	DecodeResourcePool_CpuReservation(&new.Spec.ForProvider, valMap)
	DecodeResourcePool_CpuShareLevel(&new.Spec.ForProvider, valMap)
	DecodeResourcePool_CpuShares(&new.Spec.ForProvider, valMap)
	DecodeResourcePool_CustomAttributes(&new.Spec.ForProvider, valMap)
	DecodeResourcePool_MemoryExpandable(&new.Spec.ForProvider, valMap)
	DecodeResourcePool_MemoryLimit(&new.Spec.ForProvider, valMap)
	DecodeResourcePool_MemoryReservation(&new.Spec.ForProvider, valMap)
	DecodeResourcePool_MemoryShareLevel(&new.Spec.ForProvider, valMap)
	DecodeResourcePool_MemoryShares(&new.Spec.ForProvider, valMap)
	DecodeResourcePool_Name(&new.Spec.ForProvider, valMap)
	DecodeResourcePool_ParentResourcePoolId(&new.Spec.ForProvider, valMap)
	DecodeResourcePool_Tags(&new.Spec.ForProvider, valMap)

	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeResourcePool_CpuExpandable(p *ResourcePoolParameters, vals map[string]cty.Value) {
	p.CpuExpandable = ctwhy.ValueAsBool(vals["cpu_expandable"])
}

//primitiveTypeDecodeTemplate
func DecodeResourcePool_CpuLimit(p *ResourcePoolParameters, vals map[string]cty.Value) {
	p.CpuLimit = ctwhy.ValueAsInt64(vals["cpu_limit"])
}

//primitiveTypeDecodeTemplate
func DecodeResourcePool_CpuReservation(p *ResourcePoolParameters, vals map[string]cty.Value) {
	p.CpuReservation = ctwhy.ValueAsInt64(vals["cpu_reservation"])
}

//primitiveTypeDecodeTemplate
func DecodeResourcePool_CpuShareLevel(p *ResourcePoolParameters, vals map[string]cty.Value) {
	p.CpuShareLevel = ctwhy.ValueAsString(vals["cpu_share_level"])
}

//primitiveTypeDecodeTemplate
func DecodeResourcePool_CpuShares(p *ResourcePoolParameters, vals map[string]cty.Value) {
	p.CpuShares = ctwhy.ValueAsInt64(vals["cpu_shares"])
}

//primitiveMapTypeDecodeTemplate
func DecodeResourcePool_CustomAttributes(p *ResourcePoolParameters, vals map[string]cty.Value) {
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
func DecodeResourcePool_MemoryExpandable(p *ResourcePoolParameters, vals map[string]cty.Value) {
	p.MemoryExpandable = ctwhy.ValueAsBool(vals["memory_expandable"])
}

//primitiveTypeDecodeTemplate
func DecodeResourcePool_MemoryLimit(p *ResourcePoolParameters, vals map[string]cty.Value) {
	p.MemoryLimit = ctwhy.ValueAsInt64(vals["memory_limit"])
}

//primitiveTypeDecodeTemplate
func DecodeResourcePool_MemoryReservation(p *ResourcePoolParameters, vals map[string]cty.Value) {
	p.MemoryReservation = ctwhy.ValueAsInt64(vals["memory_reservation"])
}

//primitiveTypeDecodeTemplate
func DecodeResourcePool_MemoryShareLevel(p *ResourcePoolParameters, vals map[string]cty.Value) {
	p.MemoryShareLevel = ctwhy.ValueAsString(vals["memory_share_level"])
}

//primitiveTypeDecodeTemplate
func DecodeResourcePool_MemoryShares(p *ResourcePoolParameters, vals map[string]cty.Value) {
	p.MemoryShares = ctwhy.ValueAsInt64(vals["memory_shares"])
}

//primitiveTypeDecodeTemplate
func DecodeResourcePool_Name(p *ResourcePoolParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveTypeDecodeTemplate
func DecodeResourcePool_ParentResourcePoolId(p *ResourcePoolParameters, vals map[string]cty.Value) {
	p.ParentResourcePoolId = ctwhy.ValueAsString(vals["parent_resource_pool_id"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeResourcePool_Tags(p *ResourcePoolParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["tags"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.Tags = goVals
}
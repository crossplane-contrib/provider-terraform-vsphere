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
	r, ok := mr.(*VappContainer)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeVappContainer(r, ctyValue)
}

func DecodeVappContainer(prev *VappContainer, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeVappContainer_CpuExpandable(&new.Spec.ForProvider, valMap)
	DecodeVappContainer_CpuLimit(&new.Spec.ForProvider, valMap)
	DecodeVappContainer_CpuReservation(&new.Spec.ForProvider, valMap)
	DecodeVappContainer_CpuShareLevel(&new.Spec.ForProvider, valMap)
	DecodeVappContainer_CpuShares(&new.Spec.ForProvider, valMap)
	DecodeVappContainer_CustomAttributes(&new.Spec.ForProvider, valMap)
	DecodeVappContainer_MemoryExpandable(&new.Spec.ForProvider, valMap)
	DecodeVappContainer_MemoryLimit(&new.Spec.ForProvider, valMap)
	DecodeVappContainer_MemoryReservation(&new.Spec.ForProvider, valMap)
	DecodeVappContainer_MemoryShareLevel(&new.Spec.ForProvider, valMap)
	DecodeVappContainer_MemoryShares(&new.Spec.ForProvider, valMap)
	DecodeVappContainer_Name(&new.Spec.ForProvider, valMap)
	DecodeVappContainer_ParentFolderId(&new.Spec.ForProvider, valMap)
	DecodeVappContainer_ParentResourcePoolId(&new.Spec.ForProvider, valMap)
	DecodeVappContainer_Tags(&new.Spec.ForProvider, valMap)

	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeVappContainer_CpuExpandable(p *VappContainerParameters, vals map[string]cty.Value) {
	p.CpuExpandable = ctwhy.ValueAsBool(vals["cpu_expandable"])
}

//primitiveTypeDecodeTemplate
func DecodeVappContainer_CpuLimit(p *VappContainerParameters, vals map[string]cty.Value) {
	p.CpuLimit = ctwhy.ValueAsInt64(vals["cpu_limit"])
}

//primitiveTypeDecodeTemplate
func DecodeVappContainer_CpuReservation(p *VappContainerParameters, vals map[string]cty.Value) {
	p.CpuReservation = ctwhy.ValueAsInt64(vals["cpu_reservation"])
}

//primitiveTypeDecodeTemplate
func DecodeVappContainer_CpuShareLevel(p *VappContainerParameters, vals map[string]cty.Value) {
	p.CpuShareLevel = ctwhy.ValueAsString(vals["cpu_share_level"])
}

//primitiveTypeDecodeTemplate
func DecodeVappContainer_CpuShares(p *VappContainerParameters, vals map[string]cty.Value) {
	p.CpuShares = ctwhy.ValueAsInt64(vals["cpu_shares"])
}

//primitiveMapTypeDecodeTemplate
func DecodeVappContainer_CustomAttributes(p *VappContainerParameters, vals map[string]cty.Value) {
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
func DecodeVappContainer_MemoryExpandable(p *VappContainerParameters, vals map[string]cty.Value) {
	p.MemoryExpandable = ctwhy.ValueAsBool(vals["memory_expandable"])
}

//primitiveTypeDecodeTemplate
func DecodeVappContainer_MemoryLimit(p *VappContainerParameters, vals map[string]cty.Value) {
	p.MemoryLimit = ctwhy.ValueAsInt64(vals["memory_limit"])
}

//primitiveTypeDecodeTemplate
func DecodeVappContainer_MemoryReservation(p *VappContainerParameters, vals map[string]cty.Value) {
	p.MemoryReservation = ctwhy.ValueAsInt64(vals["memory_reservation"])
}

//primitiveTypeDecodeTemplate
func DecodeVappContainer_MemoryShareLevel(p *VappContainerParameters, vals map[string]cty.Value) {
	p.MemoryShareLevel = ctwhy.ValueAsString(vals["memory_share_level"])
}

//primitiveTypeDecodeTemplate
func DecodeVappContainer_MemoryShares(p *VappContainerParameters, vals map[string]cty.Value) {
	p.MemoryShares = ctwhy.ValueAsInt64(vals["memory_shares"])
}

//primitiveTypeDecodeTemplate
func DecodeVappContainer_Name(p *VappContainerParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveTypeDecodeTemplate
func DecodeVappContainer_ParentFolderId(p *VappContainerParameters, vals map[string]cty.Value) {
	p.ParentFolderId = ctwhy.ValueAsString(vals["parent_folder_id"])
}

//primitiveTypeDecodeTemplate
func DecodeVappContainer_ParentResourcePoolId(p *VappContainerParameters, vals map[string]cty.Value) {
	p.ParentResourcePoolId = ctwhy.ValueAsString(vals["parent_resource_pool_id"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeVappContainer_Tags(p *VappContainerParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["tags"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.Tags = goVals
}

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
	r, ok := mr.(*VappEntity)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeVappEntity(r, ctyValue)
}

func DecodeVappEntity(prev *VappEntity, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeVappEntity_StopAction(&new.Spec.ForProvider, valMap)
	DecodeVappEntity_WaitForGuest(&new.Spec.ForProvider, valMap)
	DecodeVappEntity_ContainerId(&new.Spec.ForProvider, valMap)
	DecodeVappEntity_CustomAttributes(&new.Spec.ForProvider, valMap)
	DecodeVappEntity_StartAction(&new.Spec.ForProvider, valMap)
	DecodeVappEntity_TargetId(&new.Spec.ForProvider, valMap)
	DecodeVappEntity_StartDelay(&new.Spec.ForProvider, valMap)
	DecodeVappEntity_StartOrder(&new.Spec.ForProvider, valMap)
	DecodeVappEntity_StopDelay(&new.Spec.ForProvider, valMap)
	DecodeVappEntity_Tags(&new.Spec.ForProvider, valMap)

	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeVappEntity_StopAction(p *VappEntityParameters, vals map[string]cty.Value) {
	p.StopAction = ctwhy.ValueAsString(vals["stop_action"])
}

//primitiveTypeDecodeTemplate
func DecodeVappEntity_WaitForGuest(p *VappEntityParameters, vals map[string]cty.Value) {
	p.WaitForGuest = ctwhy.ValueAsBool(vals["wait_for_guest"])
}

//primitiveTypeDecodeTemplate
func DecodeVappEntity_ContainerId(p *VappEntityParameters, vals map[string]cty.Value) {
	p.ContainerId = ctwhy.ValueAsString(vals["container_id"])
}

//primitiveMapTypeDecodeTemplate
func DecodeVappEntity_CustomAttributes(p *VappEntityParameters, vals map[string]cty.Value) {
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
func DecodeVappEntity_StartAction(p *VappEntityParameters, vals map[string]cty.Value) {
	p.StartAction = ctwhy.ValueAsString(vals["start_action"])
}

//primitiveTypeDecodeTemplate
func DecodeVappEntity_TargetId(p *VappEntityParameters, vals map[string]cty.Value) {
	p.TargetId = ctwhy.ValueAsString(vals["target_id"])
}

//primitiveTypeDecodeTemplate
func DecodeVappEntity_StartDelay(p *VappEntityParameters, vals map[string]cty.Value) {
	p.StartDelay = ctwhy.ValueAsInt64(vals["start_delay"])
}

//primitiveTypeDecodeTemplate
func DecodeVappEntity_StartOrder(p *VappEntityParameters, vals map[string]cty.Value) {
	p.StartOrder = ctwhy.ValueAsInt64(vals["start_order"])
}

//primitiveTypeDecodeTemplate
func DecodeVappEntity_StopDelay(p *VappEntityParameters, vals map[string]cty.Value) {
	p.StopDelay = ctwhy.ValueAsInt64(vals["stop_delay"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeVappEntity_Tags(p *VappEntityParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["tags"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.Tags = goVals
}
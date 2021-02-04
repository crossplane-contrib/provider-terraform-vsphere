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
	r, ok := mr.(*VappEntity)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a VappEntity.")
	}
	return EncodeVappEntity(*r), nil
}

func EncodeVappEntity(r VappEntity) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeVappEntity_StopAction(r.Spec.ForProvider, ctyVal)
	EncodeVappEntity_WaitForGuest(r.Spec.ForProvider, ctyVal)
	EncodeVappEntity_ContainerId(r.Spec.ForProvider, ctyVal)
	EncodeVappEntity_CustomAttributes(r.Spec.ForProvider, ctyVal)
	EncodeVappEntity_StartAction(r.Spec.ForProvider, ctyVal)
	EncodeVappEntity_TargetId(r.Spec.ForProvider, ctyVal)
	EncodeVappEntity_StartDelay(r.Spec.ForProvider, ctyVal)
	EncodeVappEntity_StartOrder(r.Spec.ForProvider, ctyVal)
	EncodeVappEntity_StopDelay(r.Spec.ForProvider, ctyVal)
	EncodeVappEntity_Tags(r.Spec.ForProvider, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeVappEntity_StopAction(p VappEntityParameters, vals map[string]cty.Value) {
	vals["stop_action"] = cty.StringVal(p.StopAction)
}

func EncodeVappEntity_WaitForGuest(p VappEntityParameters, vals map[string]cty.Value) {
	vals["wait_for_guest"] = cty.BoolVal(p.WaitForGuest)
}

func EncodeVappEntity_ContainerId(p VappEntityParameters, vals map[string]cty.Value) {
	vals["container_id"] = cty.StringVal(p.ContainerId)
}

func EncodeVappEntity_CustomAttributes(p VappEntityParameters, vals map[string]cty.Value) {
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

func EncodeVappEntity_StartAction(p VappEntityParameters, vals map[string]cty.Value) {
	vals["start_action"] = cty.StringVal(p.StartAction)
}

func EncodeVappEntity_TargetId(p VappEntityParameters, vals map[string]cty.Value) {
	vals["target_id"] = cty.StringVal(p.TargetId)
}

func EncodeVappEntity_StartDelay(p VappEntityParameters, vals map[string]cty.Value) {
	vals["start_delay"] = cty.NumberIntVal(p.StartDelay)
}

func EncodeVappEntity_StartOrder(p VappEntityParameters, vals map[string]cty.Value) {
	vals["start_order"] = cty.NumberIntVal(p.StartOrder)
}

func EncodeVappEntity_StopDelay(p VappEntityParameters, vals map[string]cty.Value) {
	vals["stop_delay"] = cty.NumberIntVal(p.StopDelay)
}

func EncodeVappEntity_Tags(p VappEntityParameters, vals map[string]cty.Value) {
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
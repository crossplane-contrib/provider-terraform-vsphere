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
	r, ok := mr.(*Datacenter)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeDatacenter(r, ctyValue)
}

func DecodeDatacenter(prev *Datacenter, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeDatacenter_CustomAttributes(&new.Spec.ForProvider, valMap)
	DecodeDatacenter_Folder(&new.Spec.ForProvider, valMap)
	DecodeDatacenter_Name(&new.Spec.ForProvider, valMap)
	DecodeDatacenter_Tags(&new.Spec.ForProvider, valMap)
	DecodeDatacenter_Moid(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveMapTypeDecodeTemplate
func DecodeDatacenter_CustomAttributes(p *DatacenterParameters, vals map[string]cty.Value) {
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
func DecodeDatacenter_Folder(p *DatacenterParameters, vals map[string]cty.Value) {
	p.Folder = ctwhy.ValueAsString(vals["folder"])
}

//primitiveTypeDecodeTemplate
func DecodeDatacenter_Name(p *DatacenterParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeDatacenter_Tags(p *DatacenterParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["tags"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.Tags = goVals
}

//primitiveTypeDecodeTemplate
func DecodeDatacenter_Moid(p *DatacenterObservation, vals map[string]cty.Value) {
	p.Moid = ctwhy.ValueAsString(vals["moid"])
}
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
	r, ok := mr.(*License)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeLicense(r, ctyValue)
}

func DecodeLicense(prev *License, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeLicense_Labels(&new.Spec.ForProvider, valMap)
	DecodeLicense_LicenseKey(&new.Spec.ForProvider, valMap)
	DecodeLicense_Used(&new.Status.AtProvider, valMap)
	DecodeLicense_EditionKey(&new.Status.AtProvider, valMap)
	DecodeLicense_Name(&new.Status.AtProvider, valMap)
	DecodeLicense_Total(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveMapTypeDecodeTemplate
func DecodeLicense_Labels(p *LicenseParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	if vals["labels"].IsNull() {
		p.Labels = nil
        return
    }
	vMap := make(map[string]string)
	v := vals["labels"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Labels = vMap
}

//primitiveTypeDecodeTemplate
func DecodeLicense_LicenseKey(p *LicenseParameters, vals map[string]cty.Value) {
	p.LicenseKey = ctwhy.ValueAsString(vals["license_key"])
}

//primitiveTypeDecodeTemplate
func DecodeLicense_Used(p *LicenseObservation, vals map[string]cty.Value) {
	p.Used = ctwhy.ValueAsInt64(vals["used"])
}

//primitiveTypeDecodeTemplate
func DecodeLicense_EditionKey(p *LicenseObservation, vals map[string]cty.Value) {
	p.EditionKey = ctwhy.ValueAsString(vals["edition_key"])
}

//primitiveTypeDecodeTemplate
func DecodeLicense_Name(p *LicenseObservation, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveTypeDecodeTemplate
func DecodeLicense_Total(p *LicenseObservation, vals map[string]cty.Value) {
	p.Total = ctwhy.ValueAsInt64(vals["total"])
}
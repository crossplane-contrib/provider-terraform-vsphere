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
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*License)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a License.")
	}
	return EncodeLicense(*r), nil
}

func EncodeLicense(r License) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeLicense_Labels(r.Spec.ForProvider, ctyVal)
	EncodeLicense_LicenseKey(r.Spec.ForProvider, ctyVal)
	EncodeLicense_EditionKey(r.Status.AtProvider, ctyVal)
	EncodeLicense_Name(r.Status.AtProvider, ctyVal)
	EncodeLicense_Total(r.Status.AtProvider, ctyVal)
	EncodeLicense_Used(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeLicense_Labels(p LicenseParameters, vals map[string]cty.Value) {
	if len(p.Labels) == 0 {
		vals["labels"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.Labels {
		mVals[key] = cty.StringVal(value)
	}
	vals["labels"] = cty.MapVal(mVals)
}

func EncodeLicense_LicenseKey(p LicenseParameters, vals map[string]cty.Value) {
	vals["license_key"] = cty.StringVal(p.LicenseKey)
}

func EncodeLicense_EditionKey(p LicenseObservation, vals map[string]cty.Value) {
	vals["edition_key"] = cty.StringVal(p.EditionKey)
}

func EncodeLicense_Name(p LicenseObservation, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeLicense_Total(p LicenseObservation, vals map[string]cty.Value) {
	vals["total"] = cty.NumberIntVal(p.Total)
}

func EncodeLicense_Used(p LicenseObservation, vals map[string]cty.Value) {
	vals["used"] = cty.NumberIntVal(p.Used)
}

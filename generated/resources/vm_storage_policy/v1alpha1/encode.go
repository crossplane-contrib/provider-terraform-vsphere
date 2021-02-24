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
	r, ok := mr.(*VmStoragePolicy)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a VmStoragePolicy.")
	}
	return EncodeVmStoragePolicy(*r), nil
}

func EncodeVmStoragePolicy(r VmStoragePolicy) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeVmStoragePolicy_Description(r.Spec.ForProvider, ctyVal)
	EncodeVmStoragePolicy_Name(r.Spec.ForProvider, ctyVal)
	EncodeVmStoragePolicy_TagRules(r.Spec.ForProvider.TagRules, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeVmStoragePolicy_Description(p VmStoragePolicyParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeVmStoragePolicy_Name(p VmStoragePolicyParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeVmStoragePolicy_TagRules(p []TagRules, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeVmStoragePolicy_TagRules_IncludeDatastoresWithTags(v, ctyVal)
		EncodeVmStoragePolicy_TagRules_TagCategory(v, ctyVal)
		EncodeVmStoragePolicy_TagRules_Tags(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
    if len(valsForCollection) == 0 {
		vals["tag_rules"] = cty.ListValEmpty(cty.EmptyObject)
	} else {
		vals["tag_rules"] = cty.ListVal(valsForCollection)
	}
}

func EncodeVmStoragePolicy_TagRules_IncludeDatastoresWithTags(p TagRules, vals map[string]cty.Value) {
	vals["include_datastores_with_tags"] = cty.BoolVal(p.IncludeDatastoresWithTags)
}

func EncodeVmStoragePolicy_TagRules_TagCategory(p TagRules, vals map[string]cty.Value) {
	vals["tag_category"] = cty.StringVal(p.TagCategory)
}

func EncodeVmStoragePolicy_TagRules_Tags(p TagRules, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Tags {
		colVals = append(colVals, cty.StringVal(value))
	}
	if len(colVals) == 0 {
		vals["tags"] = cty.ListValEmpty(cty.String)
	} else {
		vals["tags"] = cty.ListVal(colVals)
    }
}
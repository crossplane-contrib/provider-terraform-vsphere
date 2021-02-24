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
	r, ok := mr.(*VmStoragePolicy)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeVmStoragePolicy(r, ctyValue)
}

func DecodeVmStoragePolicy(prev *VmStoragePolicy, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeVmStoragePolicy_Description(&new.Spec.ForProvider, valMap)
	DecodeVmStoragePolicy_Name(&new.Spec.ForProvider, valMap)
	DecodeVmStoragePolicy_TagRules(&new.Spec.ForProvider.TagRules, valMap)

	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeVmStoragePolicy_Description(p *VmStoragePolicyParameters, vals map[string]cty.Value) {
	p.Description = ctwhy.ValueAsString(vals["description"])
}

//primitiveTypeDecodeTemplate
func DecodeVmStoragePolicy_Name(p *VmStoragePolicyParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//containerCollectionTypeDecodeTemplate
func DecodeVmStoragePolicy_TagRules(pp *[]TagRules, vals map[string]cty.Value) {
    if vals["tag_rules"].IsNull() {
		pp = nil
		return
	}
	rvals := ctwhy.ValueAsList(vals["tag_rules"])
	if len(rvals) == 0 {
		pp = nil
		return
	}
	lval := make([]TagRules, 0)
	for _, value := range rvals {
		valMap := value.AsValueMap()
		vi := &TagRules{}
		DecodeVmStoragePolicy_TagRules_IncludeDatastoresWithTags(vi, valMap)
		DecodeVmStoragePolicy_TagRules_TagCategory(vi, valMap)
		DecodeVmStoragePolicy_TagRules_Tags(vi, valMap)
		lval = append(lval, *vi)
	}
	pp = &lval
}

//primitiveTypeDecodeTemplate
func DecodeVmStoragePolicy_TagRules_IncludeDatastoresWithTags(p *TagRules, vals map[string]cty.Value) {
	p.IncludeDatastoresWithTags = ctwhy.ValueAsBool(vals["include_datastores_with_tags"])
}

//primitiveTypeDecodeTemplate
func DecodeVmStoragePolicy_TagRules_TagCategory(p *TagRules, vals map[string]cty.Value) {
	p.TagCategory = ctwhy.ValueAsString(vals["tag_category"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeVmStoragePolicy_TagRules_Tags(p *TagRules, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsList(vals["tags"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.Tags = goVals
}
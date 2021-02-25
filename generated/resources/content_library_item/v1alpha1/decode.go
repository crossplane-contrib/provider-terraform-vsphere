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
	r, ok := mr.(*ContentLibraryItem)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeContentLibraryItem(r, ctyValue)
}

func DecodeContentLibraryItem(prev *ContentLibraryItem, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeContentLibraryItem_Description(&new.Spec.ForProvider, valMap)
	DecodeContentLibraryItem_FileUrl(&new.Spec.ForProvider, valMap)
	DecodeContentLibraryItem_LibraryId(&new.Spec.ForProvider, valMap)
	DecodeContentLibraryItem_Name(&new.Spec.ForProvider, valMap)
	DecodeContentLibraryItem_SourceUuid(&new.Spec.ForProvider, valMap)
	DecodeContentLibraryItem_Type(&new.Spec.ForProvider, valMap)

	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeContentLibraryItem_Description(p *ContentLibraryItemParameters, vals map[string]cty.Value) {
	p.Description = ctwhy.ValueAsString(vals["description"])
}

//primitiveTypeDecodeTemplate
func DecodeContentLibraryItem_FileUrl(p *ContentLibraryItemParameters, vals map[string]cty.Value) {
	p.FileUrl = ctwhy.ValueAsString(vals["file_url"])
}

//primitiveTypeDecodeTemplate
func DecodeContentLibraryItem_LibraryId(p *ContentLibraryItemParameters, vals map[string]cty.Value) {
	p.LibraryId = ctwhy.ValueAsString(vals["library_id"])
}

//primitiveTypeDecodeTemplate
func DecodeContentLibraryItem_Name(p *ContentLibraryItemParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveTypeDecodeTemplate
func DecodeContentLibraryItem_SourceUuid(p *ContentLibraryItemParameters, vals map[string]cty.Value) {
	p.SourceUuid = ctwhy.ValueAsString(vals["source_uuid"])
}

//primitiveTypeDecodeTemplate
func DecodeContentLibraryItem_Type(p *ContentLibraryItemParameters, vals map[string]cty.Value) {
	p.Type = ctwhy.ValueAsString(vals["type"])
}

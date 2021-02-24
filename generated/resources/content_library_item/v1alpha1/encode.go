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
	r, ok := mr.(*ContentLibraryItem)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a ContentLibraryItem.")
	}
	return EncodeContentLibraryItem(*r), nil
}

func EncodeContentLibraryItem(r ContentLibraryItem) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeContentLibraryItem_Description(r.Spec.ForProvider, ctyVal)
	EncodeContentLibraryItem_FileUrl(r.Spec.ForProvider, ctyVal)
	EncodeContentLibraryItem_LibraryId(r.Spec.ForProvider, ctyVal)
	EncodeContentLibraryItem_Name(r.Spec.ForProvider, ctyVal)
	EncodeContentLibraryItem_SourceUuid(r.Spec.ForProvider, ctyVal)
	EncodeContentLibraryItem_Type(r.Spec.ForProvider, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeContentLibraryItem_Description(p ContentLibraryItemParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeContentLibraryItem_FileUrl(p ContentLibraryItemParameters, vals map[string]cty.Value) {
	vals["file_url"] = cty.StringVal(p.FileUrl)
}

func EncodeContentLibraryItem_LibraryId(p ContentLibraryItemParameters, vals map[string]cty.Value) {
	vals["library_id"] = cty.StringVal(p.LibraryId)
}

func EncodeContentLibraryItem_Name(p ContentLibraryItemParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeContentLibraryItem_SourceUuid(p ContentLibraryItemParameters, vals map[string]cty.Value) {
	vals["source_uuid"] = cty.StringVal(p.SourceUuid)
}

func EncodeContentLibraryItem_Type(p ContentLibraryItemParameters, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}
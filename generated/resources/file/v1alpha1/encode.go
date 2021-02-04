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
	r, ok := mr.(*File)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a File.")
	}
	return EncodeFile(*r), nil
}

func EncodeFile(r File) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeFile_Datacenter(r.Spec.ForProvider, ctyVal)
	EncodeFile_Datastore(r.Spec.ForProvider, ctyVal)
	EncodeFile_DestinationFile(r.Spec.ForProvider, ctyVal)
	EncodeFile_SourceDatacenter(r.Spec.ForProvider, ctyVal)
	EncodeFile_SourceDatastore(r.Spec.ForProvider, ctyVal)
	EncodeFile_SourceFile(r.Spec.ForProvider, ctyVal)
	EncodeFile_CreateDirectories(r.Spec.ForProvider, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeFile_Datacenter(p FileParameters, vals map[string]cty.Value) {
	vals["datacenter"] = cty.StringVal(p.Datacenter)
}

func EncodeFile_Datastore(p FileParameters, vals map[string]cty.Value) {
	vals["datastore"] = cty.StringVal(p.Datastore)
}

func EncodeFile_DestinationFile(p FileParameters, vals map[string]cty.Value) {
	vals["destination_file"] = cty.StringVal(p.DestinationFile)
}

func EncodeFile_SourceDatacenter(p FileParameters, vals map[string]cty.Value) {
	vals["source_datacenter"] = cty.StringVal(p.SourceDatacenter)
}

func EncodeFile_SourceDatastore(p FileParameters, vals map[string]cty.Value) {
	vals["source_datastore"] = cty.StringVal(p.SourceDatastore)
}

func EncodeFile_SourceFile(p FileParameters, vals map[string]cty.Value) {
	vals["source_file"] = cty.StringVal(p.SourceFile)
}

func EncodeFile_CreateDirectories(p FileParameters, vals map[string]cty.Value) {
	vals["create_directories"] = cty.BoolVal(p.CreateDirectories)
}
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
	r, ok := mr.(*File)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeFile(r, ctyValue)
}

func DecodeFile(prev *File, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeFile_Datacenter(&new.Spec.ForProvider, valMap)
	DecodeFile_Datastore(&new.Spec.ForProvider, valMap)
	DecodeFile_DestinationFile(&new.Spec.ForProvider, valMap)
	DecodeFile_SourceDatacenter(&new.Spec.ForProvider, valMap)
	DecodeFile_SourceDatastore(&new.Spec.ForProvider, valMap)
	DecodeFile_SourceFile(&new.Spec.ForProvider, valMap)
	DecodeFile_CreateDirectories(&new.Spec.ForProvider, valMap)

	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeFile_Datacenter(p *FileParameters, vals map[string]cty.Value) {
	p.Datacenter = ctwhy.ValueAsString(vals["datacenter"])
}

//primitiveTypeDecodeTemplate
func DecodeFile_Datastore(p *FileParameters, vals map[string]cty.Value) {
	p.Datastore = ctwhy.ValueAsString(vals["datastore"])
}

//primitiveTypeDecodeTemplate
func DecodeFile_DestinationFile(p *FileParameters, vals map[string]cty.Value) {
	p.DestinationFile = ctwhy.ValueAsString(vals["destination_file"])
}

//primitiveTypeDecodeTemplate
func DecodeFile_SourceDatacenter(p *FileParameters, vals map[string]cty.Value) {
	p.SourceDatacenter = ctwhy.ValueAsString(vals["source_datacenter"])
}

//primitiveTypeDecodeTemplate
func DecodeFile_SourceDatastore(p *FileParameters, vals map[string]cty.Value) {
	p.SourceDatastore = ctwhy.ValueAsString(vals["source_datastore"])
}

//primitiveTypeDecodeTemplate
func DecodeFile_SourceFile(p *FileParameters, vals map[string]cty.Value) {
	p.SourceFile = ctwhy.ValueAsString(vals["source_file"])
}

//primitiveTypeDecodeTemplate
func DecodeFile_CreateDirectories(p *FileParameters, vals map[string]cty.Value) {
	p.CreateDirectories = ctwhy.ValueAsBool(vals["create_directories"])
}
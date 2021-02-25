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
	r, ok := mr.(*VirtualDisk)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeVirtualDisk(r, ctyValue)
}

func DecodeVirtualDisk(prev *VirtualDisk, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeVirtualDisk_AdapterType(&new.Spec.ForProvider, valMap)
	DecodeVirtualDisk_CreateDirectories(&new.Spec.ForProvider, valMap)
	DecodeVirtualDisk_Datacenter(&new.Spec.ForProvider, valMap)
	DecodeVirtualDisk_Datastore(&new.Spec.ForProvider, valMap)
	DecodeVirtualDisk_Size(&new.Spec.ForProvider, valMap)
	DecodeVirtualDisk_Type(&new.Spec.ForProvider, valMap)
	DecodeVirtualDisk_VmdkPath(&new.Spec.ForProvider, valMap)

	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeVirtualDisk_AdapterType(p *VirtualDiskParameters, vals map[string]cty.Value) {
	p.AdapterType = ctwhy.ValueAsString(vals["adapter_type"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualDisk_CreateDirectories(p *VirtualDiskParameters, vals map[string]cty.Value) {
	p.CreateDirectories = ctwhy.ValueAsBool(vals["create_directories"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualDisk_Datacenter(p *VirtualDiskParameters, vals map[string]cty.Value) {
	p.Datacenter = ctwhy.ValueAsString(vals["datacenter"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualDisk_Datastore(p *VirtualDiskParameters, vals map[string]cty.Value) {
	p.Datastore = ctwhy.ValueAsString(vals["datastore"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualDisk_Size(p *VirtualDiskParameters, vals map[string]cty.Value) {
	p.Size = ctwhy.ValueAsInt64(vals["size"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualDisk_Type(p *VirtualDiskParameters, vals map[string]cty.Value) {
	p.Type = ctwhy.ValueAsString(vals["type"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualDisk_VmdkPath(p *VirtualDiskParameters, vals map[string]cty.Value) {
	p.VmdkPath = ctwhy.ValueAsString(vals["vmdk_path"])
}

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
	r, ok := mr.(*VirtualMachineSnapshot)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeVirtualMachineSnapshot(r, ctyValue)
}

func DecodeVirtualMachineSnapshot(prev *VirtualMachineSnapshot, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeVirtualMachineSnapshot_Consolidate(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachineSnapshot_Description(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachineSnapshot_Memory(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachineSnapshot_Quiesce(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachineSnapshot_RemoveChildren(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachineSnapshot_SnapshotName(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachineSnapshot_VirtualMachineUuid(&new.Spec.ForProvider, valMap)

	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachineSnapshot_Consolidate(p *VirtualMachineSnapshotParameters, vals map[string]cty.Value) {
	p.Consolidate = ctwhy.ValueAsBool(vals["consolidate"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachineSnapshot_Description(p *VirtualMachineSnapshotParameters, vals map[string]cty.Value) {
	p.Description = ctwhy.ValueAsString(vals["description"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachineSnapshot_Memory(p *VirtualMachineSnapshotParameters, vals map[string]cty.Value) {
	p.Memory = ctwhy.ValueAsBool(vals["memory"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachineSnapshot_Quiesce(p *VirtualMachineSnapshotParameters, vals map[string]cty.Value) {
	p.Quiesce = ctwhy.ValueAsBool(vals["quiesce"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachineSnapshot_RemoveChildren(p *VirtualMachineSnapshotParameters, vals map[string]cty.Value) {
	p.RemoveChildren = ctwhy.ValueAsBool(vals["remove_children"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachineSnapshot_SnapshotName(p *VirtualMachineSnapshotParameters, vals map[string]cty.Value) {
	p.SnapshotName = ctwhy.ValueAsString(vals["snapshot_name"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachineSnapshot_VirtualMachineUuid(p *VirtualMachineSnapshotParameters, vals map[string]cty.Value) {
	p.VirtualMachineUuid = ctwhy.ValueAsString(vals["virtual_machine_uuid"])
}

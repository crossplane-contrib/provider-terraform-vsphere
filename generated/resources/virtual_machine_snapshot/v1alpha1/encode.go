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
	r, ok := mr.(*VirtualMachineSnapshot)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a VirtualMachineSnapshot.")
	}
	return EncodeVirtualMachineSnapshot(*r), nil
}

func EncodeVirtualMachineSnapshot(r VirtualMachineSnapshot) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeVirtualMachineSnapshot_Consolidate(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachineSnapshot_Description(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachineSnapshot_Memory(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachineSnapshot_Quiesce(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachineSnapshot_RemoveChildren(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachineSnapshot_SnapshotName(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachineSnapshot_VirtualMachineUuid(r.Spec.ForProvider, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeVirtualMachineSnapshot_Consolidate(p VirtualMachineSnapshotParameters, vals map[string]cty.Value) {
	vals["consolidate"] = cty.BoolVal(p.Consolidate)
}

func EncodeVirtualMachineSnapshot_Description(p VirtualMachineSnapshotParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeVirtualMachineSnapshot_Memory(p VirtualMachineSnapshotParameters, vals map[string]cty.Value) {
	vals["memory"] = cty.BoolVal(p.Memory)
}

func EncodeVirtualMachineSnapshot_Quiesce(p VirtualMachineSnapshotParameters, vals map[string]cty.Value) {
	vals["quiesce"] = cty.BoolVal(p.Quiesce)
}

func EncodeVirtualMachineSnapshot_RemoveChildren(p VirtualMachineSnapshotParameters, vals map[string]cty.Value) {
	vals["remove_children"] = cty.BoolVal(p.RemoveChildren)
}

func EncodeVirtualMachineSnapshot_SnapshotName(p VirtualMachineSnapshotParameters, vals map[string]cty.Value) {
	vals["snapshot_name"] = cty.StringVal(p.SnapshotName)
}

func EncodeVirtualMachineSnapshot_VirtualMachineUuid(p VirtualMachineSnapshotParameters, vals map[string]cty.Value) {
	vals["virtual_machine_uuid"] = cty.StringVal(p.VirtualMachineUuid)
}

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
	r, ok := mr.(*VirtualDisk)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a VirtualDisk.")
	}
	return EncodeVirtualDisk(*r), nil
}

func EncodeVirtualDisk(r VirtualDisk) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeVirtualDisk_AdapterType(r.Spec.ForProvider, ctyVal)
	EncodeVirtualDisk_CreateDirectories(r.Spec.ForProvider, ctyVal)
	EncodeVirtualDisk_Datacenter(r.Spec.ForProvider, ctyVal)
	EncodeVirtualDisk_Datastore(r.Spec.ForProvider, ctyVal)
	EncodeVirtualDisk_Size(r.Spec.ForProvider, ctyVal)
	EncodeVirtualDisk_Type(r.Spec.ForProvider, ctyVal)
	EncodeVirtualDisk_VmdkPath(r.Spec.ForProvider, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeVirtualDisk_AdapterType(p VirtualDiskParameters, vals map[string]cty.Value) {
	vals["adapter_type"] = cty.StringVal(p.AdapterType)
}

func EncodeVirtualDisk_CreateDirectories(p VirtualDiskParameters, vals map[string]cty.Value) {
	vals["create_directories"] = cty.BoolVal(p.CreateDirectories)
}

func EncodeVirtualDisk_Datacenter(p VirtualDiskParameters, vals map[string]cty.Value) {
	vals["datacenter"] = cty.StringVal(p.Datacenter)
}

func EncodeVirtualDisk_Datastore(p VirtualDiskParameters, vals map[string]cty.Value) {
	vals["datastore"] = cty.StringVal(p.Datastore)
}

func EncodeVirtualDisk_Size(p VirtualDiskParameters, vals map[string]cty.Value) {
	vals["size"] = cty.NumberIntVal(p.Size)
}

func EncodeVirtualDisk_Type(p VirtualDiskParameters, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeVirtualDisk_VmdkPath(p VirtualDiskParameters, vals map[string]cty.Value) {
	vals["vmdk_path"] = cty.StringVal(p.VmdkPath)
}

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
	r, ok := mr.(*VmfsDatastore)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a VmfsDatastore.")
	}
	return EncodeVmfsDatastore(*r), nil
}

func EncodeVmfsDatastore(r VmfsDatastore) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeVmfsDatastore_CustomAttributes(r.Spec.ForProvider, ctyVal)
	EncodeVmfsDatastore_DatastoreClusterId(r.Spec.ForProvider, ctyVal)
	EncodeVmfsDatastore_Disks(r.Spec.ForProvider, ctyVal)
	EncodeVmfsDatastore_Folder(r.Spec.ForProvider, ctyVal)
	EncodeVmfsDatastore_HostSystemId(r.Spec.ForProvider, ctyVal)
	EncodeVmfsDatastore_Name(r.Spec.ForProvider, ctyVal)
	EncodeVmfsDatastore_Tags(r.Spec.ForProvider, ctyVal)
	EncodeVmfsDatastore_Accessible(r.Status.AtProvider, ctyVal)
	EncodeVmfsDatastore_Capacity(r.Status.AtProvider, ctyVal)
	EncodeVmfsDatastore_FreeSpace(r.Status.AtProvider, ctyVal)
	EncodeVmfsDatastore_MaintenanceMode(r.Status.AtProvider, ctyVal)
	EncodeVmfsDatastore_MultipleHostAccess(r.Status.AtProvider, ctyVal)
	EncodeVmfsDatastore_UncommittedSpace(r.Status.AtProvider, ctyVal)
	EncodeVmfsDatastore_Url(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeVmfsDatastore_CustomAttributes(p VmfsDatastoreParameters, vals map[string]cty.Value) {
	if len(p.CustomAttributes) == 0 {
		vals["custom_attributes"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.CustomAttributes {
		mVals[key] = cty.StringVal(value)
	}
	vals["custom_attributes"] = cty.MapVal(mVals)
}

func EncodeVmfsDatastore_DatastoreClusterId(p VmfsDatastoreParameters, vals map[string]cty.Value) {
	vals["datastore_cluster_id"] = cty.StringVal(p.DatastoreClusterId)
}

func EncodeVmfsDatastore_Disks(p VmfsDatastoreParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Disks {
		colVals = append(colVals, cty.StringVal(value))
	}
	if len(colVals) == 0 {
		vals["disks"] = cty.ListValEmpty(cty.String)
	} else {
		vals["disks"] = cty.ListVal(colVals)
	}
}

func EncodeVmfsDatastore_Folder(p VmfsDatastoreParameters, vals map[string]cty.Value) {
	vals["folder"] = cty.StringVal(p.Folder)
}

func EncodeVmfsDatastore_HostSystemId(p VmfsDatastoreParameters, vals map[string]cty.Value) {
	vals["host_system_id"] = cty.StringVal(p.HostSystemId)
}

func EncodeVmfsDatastore_Name(p VmfsDatastoreParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeVmfsDatastore_Tags(p VmfsDatastoreParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Tags {
		colVals = append(colVals, cty.StringVal(value))
	}
	if len(colVals) == 0 {
		vals["tags"] = cty.SetValEmpty(cty.String)
	} else {
		vals["tags"] = cty.SetVal(colVals)
	}
}

func EncodeVmfsDatastore_Accessible(p VmfsDatastoreObservation, vals map[string]cty.Value) {
	vals["accessible"] = cty.BoolVal(p.Accessible)
}

func EncodeVmfsDatastore_Capacity(p VmfsDatastoreObservation, vals map[string]cty.Value) {
	vals["capacity"] = cty.NumberIntVal(p.Capacity)
}

func EncodeVmfsDatastore_FreeSpace(p VmfsDatastoreObservation, vals map[string]cty.Value) {
	vals["free_space"] = cty.NumberIntVal(p.FreeSpace)
}

func EncodeVmfsDatastore_MaintenanceMode(p VmfsDatastoreObservation, vals map[string]cty.Value) {
	vals["maintenance_mode"] = cty.StringVal(p.MaintenanceMode)
}

func EncodeVmfsDatastore_MultipleHostAccess(p VmfsDatastoreObservation, vals map[string]cty.Value) {
	vals["multiple_host_access"] = cty.BoolVal(p.MultipleHostAccess)
}

func EncodeVmfsDatastore_UncommittedSpace(p VmfsDatastoreObservation, vals map[string]cty.Value) {
	vals["uncommitted_space"] = cty.NumberIntVal(p.UncommittedSpace)
}

func EncodeVmfsDatastore_Url(p VmfsDatastoreObservation, vals map[string]cty.Value) {
	vals["url"] = cty.StringVal(p.Url)
}

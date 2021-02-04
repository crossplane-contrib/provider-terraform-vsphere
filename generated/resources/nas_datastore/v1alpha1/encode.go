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
	r, ok := mr.(*NasDatastore)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a NasDatastore.")
	}
	return EncodeNasDatastore(*r), nil
}

func EncodeNasDatastore(r NasDatastore) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeNasDatastore_AccessMode(r.Spec.ForProvider, ctyVal)
	EncodeNasDatastore_CustomAttributes(r.Spec.ForProvider, ctyVal)
	EncodeNasDatastore_SecurityType(r.Spec.ForProvider, ctyVal)
	EncodeNasDatastore_Tags(r.Spec.ForProvider, ctyVal)
	EncodeNasDatastore_Folder(r.Spec.ForProvider, ctyVal)
	EncodeNasDatastore_RemoteHosts(r.Spec.ForProvider, ctyVal)
	EncodeNasDatastore_RemotePath(r.Spec.ForProvider, ctyVal)
	EncodeNasDatastore_DatastoreClusterId(r.Spec.ForProvider, ctyVal)
	EncodeNasDatastore_HostSystemIds(r.Spec.ForProvider, ctyVal)
	EncodeNasDatastore_Name(r.Spec.ForProvider, ctyVal)
	EncodeNasDatastore_Type(r.Spec.ForProvider, ctyVal)
	EncodeNasDatastore_Capacity(r.Status.AtProvider, ctyVal)
	EncodeNasDatastore_ProtocolEndpoint(r.Status.AtProvider, ctyVal)
	EncodeNasDatastore_Accessible(r.Status.AtProvider, ctyVal)
	EncodeNasDatastore_FreeSpace(r.Status.AtProvider, ctyVal)
	EncodeNasDatastore_Url(r.Status.AtProvider, ctyVal)
	EncodeNasDatastore_MaintenanceMode(r.Status.AtProvider, ctyVal)
	EncodeNasDatastore_MultipleHostAccess(r.Status.AtProvider, ctyVal)
	EncodeNasDatastore_UncommittedSpace(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeNasDatastore_AccessMode(p NasDatastoreParameters, vals map[string]cty.Value) {
	vals["access_mode"] = cty.StringVal(p.AccessMode)
}

func EncodeNasDatastore_CustomAttributes(p NasDatastoreParameters, vals map[string]cty.Value) {
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

func EncodeNasDatastore_SecurityType(p NasDatastoreParameters, vals map[string]cty.Value) {
	vals["security_type"] = cty.StringVal(p.SecurityType)
}

func EncodeNasDatastore_Tags(p NasDatastoreParameters, vals map[string]cty.Value) {
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

func EncodeNasDatastore_Folder(p NasDatastoreParameters, vals map[string]cty.Value) {
	vals["folder"] = cty.StringVal(p.Folder)
}

func EncodeNasDatastore_RemoteHosts(p NasDatastoreParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.RemoteHosts {
		colVals = append(colVals, cty.StringVal(value))
	}
	if len(colVals) == 0 {
		vals["remote_hosts"] = cty.ListValEmpty(cty.String)
	} else {
		vals["remote_hosts"] = cty.ListVal(colVals)
    }
}

func EncodeNasDatastore_RemotePath(p NasDatastoreParameters, vals map[string]cty.Value) {
	vals["remote_path"] = cty.StringVal(p.RemotePath)
}

func EncodeNasDatastore_DatastoreClusterId(p NasDatastoreParameters, vals map[string]cty.Value) {
	vals["datastore_cluster_id"] = cty.StringVal(p.DatastoreClusterId)
}

func EncodeNasDatastore_HostSystemIds(p NasDatastoreParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.HostSystemIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	if len(colVals) == 0 {
		vals["host_system_ids"] = cty.SetValEmpty(cty.String)
	} else {
		vals["host_system_ids"] = cty.SetVal(colVals)
    }
}

func EncodeNasDatastore_Name(p NasDatastoreParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeNasDatastore_Type(p NasDatastoreParameters, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeNasDatastore_Capacity(p NasDatastoreObservation, vals map[string]cty.Value) {
	vals["capacity"] = cty.NumberIntVal(p.Capacity)
}

func EncodeNasDatastore_ProtocolEndpoint(p NasDatastoreObservation, vals map[string]cty.Value) {
	vals["protocol_endpoint"] = cty.StringVal(p.ProtocolEndpoint)
}

func EncodeNasDatastore_Accessible(p NasDatastoreObservation, vals map[string]cty.Value) {
	vals["accessible"] = cty.BoolVal(p.Accessible)
}

func EncodeNasDatastore_FreeSpace(p NasDatastoreObservation, vals map[string]cty.Value) {
	vals["free_space"] = cty.NumberIntVal(p.FreeSpace)
}

func EncodeNasDatastore_Url(p NasDatastoreObservation, vals map[string]cty.Value) {
	vals["url"] = cty.StringVal(p.Url)
}

func EncodeNasDatastore_MaintenanceMode(p NasDatastoreObservation, vals map[string]cty.Value) {
	vals["maintenance_mode"] = cty.StringVal(p.MaintenanceMode)
}

func EncodeNasDatastore_MultipleHostAccess(p NasDatastoreObservation, vals map[string]cty.Value) {
	vals["multiple_host_access"] = cty.BoolVal(p.MultipleHostAccess)
}

func EncodeNasDatastore_UncommittedSpace(p NasDatastoreObservation, vals map[string]cty.Value) {
	vals["uncommitted_space"] = cty.NumberIntVal(p.UncommittedSpace)
}
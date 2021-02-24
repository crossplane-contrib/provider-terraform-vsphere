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
	r, ok := mr.(*NasDatastore)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeNasDatastore(r, ctyValue)
}

func DecodeNasDatastore(prev *NasDatastore, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeNasDatastore_AccessMode(&new.Spec.ForProvider, valMap)
	DecodeNasDatastore_CustomAttributes(&new.Spec.ForProvider, valMap)
	DecodeNasDatastore_DatastoreClusterId(&new.Spec.ForProvider, valMap)
	DecodeNasDatastore_Folder(&new.Spec.ForProvider, valMap)
	DecodeNasDatastore_HostSystemIds(&new.Spec.ForProvider, valMap)
	DecodeNasDatastore_Name(&new.Spec.ForProvider, valMap)
	DecodeNasDatastore_RemoteHosts(&new.Spec.ForProvider, valMap)
	DecodeNasDatastore_RemotePath(&new.Spec.ForProvider, valMap)
	DecodeNasDatastore_SecurityType(&new.Spec.ForProvider, valMap)
	DecodeNasDatastore_Tags(&new.Spec.ForProvider, valMap)
	DecodeNasDatastore_Type(&new.Spec.ForProvider, valMap)
	DecodeNasDatastore_Accessible(&new.Status.AtProvider, valMap)
	DecodeNasDatastore_Capacity(&new.Status.AtProvider, valMap)
	DecodeNasDatastore_FreeSpace(&new.Status.AtProvider, valMap)
	DecodeNasDatastore_MaintenanceMode(&new.Status.AtProvider, valMap)
	DecodeNasDatastore_MultipleHostAccess(&new.Status.AtProvider, valMap)
	DecodeNasDatastore_ProtocolEndpoint(&new.Status.AtProvider, valMap)
	DecodeNasDatastore_UncommittedSpace(&new.Status.AtProvider, valMap)
	DecodeNasDatastore_Url(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeNasDatastore_AccessMode(p *NasDatastoreParameters, vals map[string]cty.Value) {
	p.AccessMode = ctwhy.ValueAsString(vals["access_mode"])
}

//primitiveMapTypeDecodeTemplate
func DecodeNasDatastore_CustomAttributes(p *NasDatastoreParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	if vals["custom_attributes"].IsNull() {
		p.CustomAttributes = nil
        return
    }
	vMap := make(map[string]string)
	v := vals["custom_attributes"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.CustomAttributes = vMap
}

//primitiveTypeDecodeTemplate
func DecodeNasDatastore_DatastoreClusterId(p *NasDatastoreParameters, vals map[string]cty.Value) {
	p.DatastoreClusterId = ctwhy.ValueAsString(vals["datastore_cluster_id"])
}

//primitiveTypeDecodeTemplate
func DecodeNasDatastore_Folder(p *NasDatastoreParameters, vals map[string]cty.Value) {
	p.Folder = ctwhy.ValueAsString(vals["folder"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeNasDatastore_HostSystemIds(p *NasDatastoreParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["host_system_ids"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.HostSystemIds = goVals
}

//primitiveTypeDecodeTemplate
func DecodeNasDatastore_Name(p *NasDatastoreParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeNasDatastore_RemoteHosts(p *NasDatastoreParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsList(vals["remote_hosts"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.RemoteHosts = goVals
}

//primitiveTypeDecodeTemplate
func DecodeNasDatastore_RemotePath(p *NasDatastoreParameters, vals map[string]cty.Value) {
	p.RemotePath = ctwhy.ValueAsString(vals["remote_path"])
}

//primitiveTypeDecodeTemplate
func DecodeNasDatastore_SecurityType(p *NasDatastoreParameters, vals map[string]cty.Value) {
	p.SecurityType = ctwhy.ValueAsString(vals["security_type"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeNasDatastore_Tags(p *NasDatastoreParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["tags"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.Tags = goVals
}

//primitiveTypeDecodeTemplate
func DecodeNasDatastore_Type(p *NasDatastoreParameters, vals map[string]cty.Value) {
	p.Type = ctwhy.ValueAsString(vals["type"])
}

//primitiveTypeDecodeTemplate
func DecodeNasDatastore_Accessible(p *NasDatastoreObservation, vals map[string]cty.Value) {
	p.Accessible = ctwhy.ValueAsBool(vals["accessible"])
}

//primitiveTypeDecodeTemplate
func DecodeNasDatastore_Capacity(p *NasDatastoreObservation, vals map[string]cty.Value) {
	p.Capacity = ctwhy.ValueAsInt64(vals["capacity"])
}

//primitiveTypeDecodeTemplate
func DecodeNasDatastore_FreeSpace(p *NasDatastoreObservation, vals map[string]cty.Value) {
	p.FreeSpace = ctwhy.ValueAsInt64(vals["free_space"])
}

//primitiveTypeDecodeTemplate
func DecodeNasDatastore_MaintenanceMode(p *NasDatastoreObservation, vals map[string]cty.Value) {
	p.MaintenanceMode = ctwhy.ValueAsString(vals["maintenance_mode"])
}

//primitiveTypeDecodeTemplate
func DecodeNasDatastore_MultipleHostAccess(p *NasDatastoreObservation, vals map[string]cty.Value) {
	p.MultipleHostAccess = ctwhy.ValueAsBool(vals["multiple_host_access"])
}

//primitiveTypeDecodeTemplate
func DecodeNasDatastore_ProtocolEndpoint(p *NasDatastoreObservation, vals map[string]cty.Value) {
	p.ProtocolEndpoint = ctwhy.ValueAsString(vals["protocol_endpoint"])
}

//primitiveTypeDecodeTemplate
func DecodeNasDatastore_UncommittedSpace(p *NasDatastoreObservation, vals map[string]cty.Value) {
	p.UncommittedSpace = ctwhy.ValueAsInt64(vals["uncommitted_space"])
}

//primitiveTypeDecodeTemplate
func DecodeNasDatastore_Url(p *NasDatastoreObservation, vals map[string]cty.Value) {
	p.Url = ctwhy.ValueAsString(vals["url"])
}
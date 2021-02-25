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
	r, ok := mr.(*VmfsDatastore)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeVmfsDatastore(r, ctyValue)
}

func DecodeVmfsDatastore(prev *VmfsDatastore, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeVmfsDatastore_CustomAttributes(&new.Spec.ForProvider, valMap)
	DecodeVmfsDatastore_DatastoreClusterId(&new.Spec.ForProvider, valMap)
	DecodeVmfsDatastore_Disks(&new.Spec.ForProvider, valMap)
	DecodeVmfsDatastore_Folder(&new.Spec.ForProvider, valMap)
	DecodeVmfsDatastore_HostSystemId(&new.Spec.ForProvider, valMap)
	DecodeVmfsDatastore_Name(&new.Spec.ForProvider, valMap)
	DecodeVmfsDatastore_Tags(&new.Spec.ForProvider, valMap)
	DecodeVmfsDatastore_Accessible(&new.Status.AtProvider, valMap)
	DecodeVmfsDatastore_Capacity(&new.Status.AtProvider, valMap)
	DecodeVmfsDatastore_FreeSpace(&new.Status.AtProvider, valMap)
	DecodeVmfsDatastore_MaintenanceMode(&new.Status.AtProvider, valMap)
	DecodeVmfsDatastore_MultipleHostAccess(&new.Status.AtProvider, valMap)
	DecodeVmfsDatastore_UncommittedSpace(&new.Status.AtProvider, valMap)
	DecodeVmfsDatastore_Url(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveMapTypeDecodeTemplate
func DecodeVmfsDatastore_CustomAttributes(p *VmfsDatastoreParameters, vals map[string]cty.Value) {
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
func DecodeVmfsDatastore_DatastoreClusterId(p *VmfsDatastoreParameters, vals map[string]cty.Value) {
	p.DatastoreClusterId = ctwhy.ValueAsString(vals["datastore_cluster_id"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeVmfsDatastore_Disks(p *VmfsDatastoreParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsList(vals["disks"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.Disks = goVals
}

//primitiveTypeDecodeTemplate
func DecodeVmfsDatastore_Folder(p *VmfsDatastoreParameters, vals map[string]cty.Value) {
	p.Folder = ctwhy.ValueAsString(vals["folder"])
}

//primitiveTypeDecodeTemplate
func DecodeVmfsDatastore_HostSystemId(p *VmfsDatastoreParameters, vals map[string]cty.Value) {
	p.HostSystemId = ctwhy.ValueAsString(vals["host_system_id"])
}

//primitiveTypeDecodeTemplate
func DecodeVmfsDatastore_Name(p *VmfsDatastoreParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeVmfsDatastore_Tags(p *VmfsDatastoreParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["tags"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.Tags = goVals
}

//primitiveTypeDecodeTemplate
func DecodeVmfsDatastore_Accessible(p *VmfsDatastoreObservation, vals map[string]cty.Value) {
	p.Accessible = ctwhy.ValueAsBool(vals["accessible"])
}

//primitiveTypeDecodeTemplate
func DecodeVmfsDatastore_Capacity(p *VmfsDatastoreObservation, vals map[string]cty.Value) {
	p.Capacity = ctwhy.ValueAsInt64(vals["capacity"])
}

//primitiveTypeDecodeTemplate
func DecodeVmfsDatastore_FreeSpace(p *VmfsDatastoreObservation, vals map[string]cty.Value) {
	p.FreeSpace = ctwhy.ValueAsInt64(vals["free_space"])
}

//primitiveTypeDecodeTemplate
func DecodeVmfsDatastore_MaintenanceMode(p *VmfsDatastoreObservation, vals map[string]cty.Value) {
	p.MaintenanceMode = ctwhy.ValueAsString(vals["maintenance_mode"])
}

//primitiveTypeDecodeTemplate
func DecodeVmfsDatastore_MultipleHostAccess(p *VmfsDatastoreObservation, vals map[string]cty.Value) {
	p.MultipleHostAccess = ctwhy.ValueAsBool(vals["multiple_host_access"])
}

//primitiveTypeDecodeTemplate
func DecodeVmfsDatastore_UncommittedSpace(p *VmfsDatastoreObservation, vals map[string]cty.Value) {
	p.UncommittedSpace = ctwhy.ValueAsInt64(vals["uncommitted_space"])
}

//primitiveTypeDecodeTemplate
func DecodeVmfsDatastore_Url(p *VmfsDatastoreObservation, vals map[string]cty.Value) {
	p.Url = ctwhy.ValueAsString(vals["url"])
}

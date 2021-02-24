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
	r, ok := mr.(*StorageDrsVmOverride)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a StorageDrsVmOverride.")
	}
	return EncodeStorageDrsVmOverride(*r), nil
}

func EncodeStorageDrsVmOverride(r StorageDrsVmOverride) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeStorageDrsVmOverride_DatastoreClusterId(r.Spec.ForProvider, ctyVal)
	EncodeStorageDrsVmOverride_SdrsAutomationLevel(r.Spec.ForProvider, ctyVal)
	EncodeStorageDrsVmOverride_SdrsEnabled(r.Spec.ForProvider, ctyVal)
	EncodeStorageDrsVmOverride_SdrsIntraVmAffinity(r.Spec.ForProvider, ctyVal)
	EncodeStorageDrsVmOverride_VirtualMachineId(r.Spec.ForProvider, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeStorageDrsVmOverride_DatastoreClusterId(p StorageDrsVmOverrideParameters, vals map[string]cty.Value) {
	vals["datastore_cluster_id"] = cty.StringVal(p.DatastoreClusterId)
}

func EncodeStorageDrsVmOverride_SdrsAutomationLevel(p StorageDrsVmOverrideParameters, vals map[string]cty.Value) {
	vals["sdrs_automation_level"] = cty.StringVal(p.SdrsAutomationLevel)
}

func EncodeStorageDrsVmOverride_SdrsEnabled(p StorageDrsVmOverrideParameters, vals map[string]cty.Value) {
	vals["sdrs_enabled"] = cty.StringVal(p.SdrsEnabled)
}

func EncodeStorageDrsVmOverride_SdrsIntraVmAffinity(p StorageDrsVmOverrideParameters, vals map[string]cty.Value) {
	vals["sdrs_intra_vm_affinity"] = cty.StringVal(p.SdrsIntraVmAffinity)
}

func EncodeStorageDrsVmOverride_VirtualMachineId(p StorageDrsVmOverrideParameters, vals map[string]cty.Value) {
	vals["virtual_machine_id"] = cty.StringVal(p.VirtualMachineId)
}
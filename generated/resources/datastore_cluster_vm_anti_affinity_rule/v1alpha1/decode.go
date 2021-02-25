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
	r, ok := mr.(*DatastoreClusterVmAntiAffinityRule)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeDatastoreClusterVmAntiAffinityRule(r, ctyValue)
}

func DecodeDatastoreClusterVmAntiAffinityRule(prev *DatastoreClusterVmAntiAffinityRule, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeDatastoreClusterVmAntiAffinityRule_DatastoreClusterId(&new.Spec.ForProvider, valMap)
	DecodeDatastoreClusterVmAntiAffinityRule_Enabled(&new.Spec.ForProvider, valMap)
	DecodeDatastoreClusterVmAntiAffinityRule_Mandatory(&new.Spec.ForProvider, valMap)
	DecodeDatastoreClusterVmAntiAffinityRule_Name(&new.Spec.ForProvider, valMap)
	DecodeDatastoreClusterVmAntiAffinityRule_VirtualMachineIds(&new.Spec.ForProvider, valMap)

	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeDatastoreClusterVmAntiAffinityRule_DatastoreClusterId(p *DatastoreClusterVmAntiAffinityRuleParameters, vals map[string]cty.Value) {
	p.DatastoreClusterId = ctwhy.ValueAsString(vals["datastore_cluster_id"])
}

//primitiveTypeDecodeTemplate
func DecodeDatastoreClusterVmAntiAffinityRule_Enabled(p *DatastoreClusterVmAntiAffinityRuleParameters, vals map[string]cty.Value) {
	p.Enabled = ctwhy.ValueAsBool(vals["enabled"])
}

//primitiveTypeDecodeTemplate
func DecodeDatastoreClusterVmAntiAffinityRule_Mandatory(p *DatastoreClusterVmAntiAffinityRuleParameters, vals map[string]cty.Value) {
	p.Mandatory = ctwhy.ValueAsBool(vals["mandatory"])
}

//primitiveTypeDecodeTemplate
func DecodeDatastoreClusterVmAntiAffinityRule_Name(p *DatastoreClusterVmAntiAffinityRuleParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeDatastoreClusterVmAntiAffinityRule_VirtualMachineIds(p *DatastoreClusterVmAntiAffinityRuleParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["virtual_machine_ids"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.VirtualMachineIds = goVals
}

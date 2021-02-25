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
	r, ok := mr.(*EntityPermissions)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeEntityPermissions(r, ctyValue)
}

func DecodeEntityPermissions(prev *EntityPermissions, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeEntityPermissions_EntityId(&new.Spec.ForProvider, valMap)
	DecodeEntityPermissions_EntityType(&new.Spec.ForProvider, valMap)
	DecodeEntityPermissions_Permissions(&new.Spec.ForProvider.Permissions, valMap)

	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeEntityPermissions_EntityId(p *EntityPermissionsParameters, vals map[string]cty.Value) {
	p.EntityId = ctwhy.ValueAsString(vals["entity_id"])
}

//primitiveTypeDecodeTemplate
func DecodeEntityPermissions_EntityType(p *EntityPermissionsParameters, vals map[string]cty.Value) {
	p.EntityType = ctwhy.ValueAsString(vals["entity_type"])
}

//containerCollectionTypeDecodeTemplate
func DecodeEntityPermissions_Permissions(pp *[]Permissions, vals map[string]cty.Value) {
	if vals["permissions"].IsNull() {
		pp = nil
		return
	}
	rvals := ctwhy.ValueAsList(vals["permissions"])
	if len(rvals) == 0 {
		pp = nil
		return
	}
	lval := make([]Permissions, 0)
	for _, value := range rvals {
		valMap := value.AsValueMap()
		vi := &Permissions{}
		DecodeEntityPermissions_Permissions_IsGroup(vi, valMap)
		DecodeEntityPermissions_Permissions_Propagate(vi, valMap)
		DecodeEntityPermissions_Permissions_RoleId(vi, valMap)
		DecodeEntityPermissions_Permissions_UserOrGroup(vi, valMap)
		lval = append(lval, *vi)
	}
	pp = &lval
}

//primitiveTypeDecodeTemplate
func DecodeEntityPermissions_Permissions_IsGroup(p *Permissions, vals map[string]cty.Value) {
	p.IsGroup = ctwhy.ValueAsBool(vals["is_group"])
}

//primitiveTypeDecodeTemplate
func DecodeEntityPermissions_Permissions_Propagate(p *Permissions, vals map[string]cty.Value) {
	p.Propagate = ctwhy.ValueAsBool(vals["propagate"])
}

//primitiveTypeDecodeTemplate
func DecodeEntityPermissions_Permissions_RoleId(p *Permissions, vals map[string]cty.Value) {
	p.RoleId = ctwhy.ValueAsString(vals["role_id"])
}

//primitiveTypeDecodeTemplate
func DecodeEntityPermissions_Permissions_UserOrGroup(p *Permissions, vals map[string]cty.Value) {
	p.UserOrGroup = ctwhy.ValueAsString(vals["user_or_group"])
}

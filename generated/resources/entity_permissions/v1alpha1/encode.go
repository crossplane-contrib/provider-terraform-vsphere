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
	r, ok := mr.(*EntityPermissions)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a EntityPermissions.")
	}
	return EncodeEntityPermissions(*r), nil
}

func EncodeEntityPermissions(r EntityPermissions) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeEntityPermissions_EntityId(r.Spec.ForProvider, ctyVal)
	EncodeEntityPermissions_EntityType(r.Spec.ForProvider, ctyVal)
	EncodeEntityPermissions_Permissions(r.Spec.ForProvider.Permissions, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeEntityPermissions_EntityId(p EntityPermissionsParameters, vals map[string]cty.Value) {
	vals["entity_id"] = cty.StringVal(p.EntityId)
}

func EncodeEntityPermissions_EntityType(p EntityPermissionsParameters, vals map[string]cty.Value) {
	vals["entity_type"] = cty.StringVal(p.EntityType)
}

func EncodeEntityPermissions_Permissions(p []Permissions, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeEntityPermissions_Permissions_IsGroup(v, ctyVal)
		EncodeEntityPermissions_Permissions_Propagate(v, ctyVal)
		EncodeEntityPermissions_Permissions_RoleId(v, ctyVal)
		EncodeEntityPermissions_Permissions_UserOrGroup(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	if len(valsForCollection) == 0 {
		vals["permissions"] = cty.ListValEmpty(cty.EmptyObject)
	} else {
		vals["permissions"] = cty.ListVal(valsForCollection)
	}
}

func EncodeEntityPermissions_Permissions_IsGroup(p Permissions, vals map[string]cty.Value) {
	vals["is_group"] = cty.BoolVal(p.IsGroup)
}

func EncodeEntityPermissions_Permissions_Propagate(p Permissions, vals map[string]cty.Value) {
	vals["propagate"] = cty.BoolVal(p.Propagate)
}

func EncodeEntityPermissions_Permissions_RoleId(p Permissions, vals map[string]cty.Value) {
	vals["role_id"] = cty.StringVal(p.RoleId)
}

func EncodeEntityPermissions_Permissions_UserOrGroup(p Permissions, vals map[string]cty.Value) {
	vals["user_or_group"] = cty.StringVal(p.UserOrGroup)
}

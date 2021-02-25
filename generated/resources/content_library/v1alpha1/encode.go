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
	r, ok := mr.(*ContentLibrary)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a ContentLibrary.")
	}
	return EncodeContentLibrary(*r), nil
}

func EncodeContentLibrary(r ContentLibrary) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeContentLibrary_Description(r.Spec.ForProvider, ctyVal)
	EncodeContentLibrary_Name(r.Spec.ForProvider, ctyVal)
	EncodeContentLibrary_StorageBacking(r.Spec.ForProvider, ctyVal)
	EncodeContentLibrary_Publication(r.Spec.ForProvider.Publication, ctyVal)
	EncodeContentLibrary_Subscription(r.Spec.ForProvider.Subscription, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeContentLibrary_Description(p ContentLibraryParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeContentLibrary_Name(p ContentLibraryParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeContentLibrary_StorageBacking(p ContentLibraryParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.StorageBacking {
		colVals = append(colVals, cty.StringVal(value))
	}
	if len(colVals) == 0 {
		vals["storage_backing"] = cty.SetValEmpty(cty.String)
	} else {
		vals["storage_backing"] = cty.SetVal(colVals)
	}
}

func EncodeContentLibrary_Publication(p Publication, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeContentLibrary_Publication_Username(p, ctyVal)
	EncodeContentLibrary_Publication_AuthenticationMethod(p, ctyVal)
	EncodeContentLibrary_Publication_Password(p, ctyVal)
	EncodeContentLibrary_Publication_PublishUrl(p, ctyVal)
	EncodeContentLibrary_Publication_Published(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	if len(valsForCollection) == 0 {
		vals["publication"] = cty.ListValEmpty(cty.EmptyObject)
	} else {
		vals["publication"] = cty.ListVal(valsForCollection)
	}
}

func EncodeContentLibrary_Publication_Username(p Publication, vals map[string]cty.Value) {
	vals["username"] = cty.StringVal(p.Username)
}

func EncodeContentLibrary_Publication_AuthenticationMethod(p Publication, vals map[string]cty.Value) {
	vals["authentication_method"] = cty.StringVal(p.AuthenticationMethod)
}

func EncodeContentLibrary_Publication_Password(p Publication, vals map[string]cty.Value) {
	vals["password"] = cty.StringVal(p.Password)
}

func EncodeContentLibrary_Publication_PublishUrl(p Publication, vals map[string]cty.Value) {
	vals["publish_url"] = cty.StringVal(p.PublishUrl)
}

func EncodeContentLibrary_Publication_Published(p Publication, vals map[string]cty.Value) {
	vals["published"] = cty.BoolVal(p.Published)
}

func EncodeContentLibrary_Subscription(p Subscription, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeContentLibrary_Subscription_Username(p, ctyVal)
	EncodeContentLibrary_Subscription_AuthenticationMethod(p, ctyVal)
	EncodeContentLibrary_Subscription_AutomaticSync(p, ctyVal)
	EncodeContentLibrary_Subscription_OnDemand(p, ctyVal)
	EncodeContentLibrary_Subscription_Password(p, ctyVal)
	EncodeContentLibrary_Subscription_SubscriptionUrl(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	if len(valsForCollection) == 0 {
		vals["subscription"] = cty.ListValEmpty(cty.EmptyObject)
	} else {
		vals["subscription"] = cty.ListVal(valsForCollection)
	}
}

func EncodeContentLibrary_Subscription_Username(p Subscription, vals map[string]cty.Value) {
	vals["username"] = cty.StringVal(p.Username)
}

func EncodeContentLibrary_Subscription_AuthenticationMethod(p Subscription, vals map[string]cty.Value) {
	vals["authentication_method"] = cty.StringVal(p.AuthenticationMethod)
}

func EncodeContentLibrary_Subscription_AutomaticSync(p Subscription, vals map[string]cty.Value) {
	vals["automatic_sync"] = cty.BoolVal(p.AutomaticSync)
}

func EncodeContentLibrary_Subscription_OnDemand(p Subscription, vals map[string]cty.Value) {
	vals["on_demand"] = cty.BoolVal(p.OnDemand)
}

func EncodeContentLibrary_Subscription_Password(p Subscription, vals map[string]cty.Value) {
	vals["password"] = cty.StringVal(p.Password)
}

func EncodeContentLibrary_Subscription_SubscriptionUrl(p Subscription, vals map[string]cty.Value) {
	vals["subscription_url"] = cty.StringVal(p.SubscriptionUrl)
}

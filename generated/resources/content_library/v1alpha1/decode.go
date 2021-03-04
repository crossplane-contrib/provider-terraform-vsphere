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
	r, ok := mr.(*ContentLibrary)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeContentLibrary(r, ctyValue)
}

func DecodeContentLibrary(prev *ContentLibrary, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeContentLibrary_Description(&new.Spec.ForProvider, valMap)
	DecodeContentLibrary_Name(&new.Spec.ForProvider, valMap)
	DecodeContentLibrary_StorageBacking(&new.Spec.ForProvider, valMap)
	DecodeContentLibrary_Publication(&new.Spec.ForProvider.Publication, valMap)
	DecodeContentLibrary_Subscription(&new.Spec.ForProvider.Subscription, valMap)

	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeContentLibrary_Description(p *ContentLibraryParameters, vals map[string]cty.Value) {
	p.Description = ctwhy.ValueAsString(vals["description"])
}

//primitiveTypeDecodeTemplate
func DecodeContentLibrary_Name(p *ContentLibraryParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeContentLibrary_StorageBacking(p *ContentLibraryParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["storage_backing"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.StorageBacking = goVals
}

//containerCollectionSingletonTypeDecodeTemplate
func DecodeContentLibrary_Publication(p *Publication, vals map[string]cty.Value) {
	if vals["publication"].IsNull() {
		*p = Publication{}
		return
	}
	rvals := ctwhy.ValueAsList(vals["publication"])
	if len(rvals) == 0 {
		*p = Publication{}
		return
	}
	// this template should be used when single dictionary/object values are nested in sets/lists
	// if rvals turns out to be a list with > 1 elements, something has broken with that heuristic
	valMap := rvals[0].AsValueMap()
	DecodeContentLibrary_Publication_Published(p, valMap)
	DecodeContentLibrary_Publication_Username(p, valMap)
	DecodeContentLibrary_Publication_AuthenticationMethod(p, valMap)
	DecodeContentLibrary_Publication_Password(p, valMap)
	DecodeContentLibrary_Publication_PublishUrl(p, valMap)
}

//primitiveTypeDecodeTemplate
func DecodeContentLibrary_Publication_Published(p *Publication, vals map[string]cty.Value) {
	p.Published = ctwhy.ValueAsBool(vals["published"])
}

//primitiveTypeDecodeTemplate
func DecodeContentLibrary_Publication_Username(p *Publication, vals map[string]cty.Value) {
	p.Username = ctwhy.ValueAsString(vals["username"])
}

//primitiveTypeDecodeTemplate
func DecodeContentLibrary_Publication_AuthenticationMethod(p *Publication, vals map[string]cty.Value) {
	p.AuthenticationMethod = ctwhy.ValueAsString(vals["authentication_method"])
}

//primitiveTypeDecodeTemplate
func DecodeContentLibrary_Publication_Password(p *Publication, vals map[string]cty.Value) {
	p.Password = ctwhy.ValueAsString(vals["password"])
}

//primitiveTypeDecodeTemplate
func DecodeContentLibrary_Publication_PublishUrl(p *Publication, vals map[string]cty.Value) {
	p.PublishUrl = ctwhy.ValueAsString(vals["publish_url"])
}

//containerCollectionSingletonTypeDecodeTemplate
func DecodeContentLibrary_Subscription(p *Subscription, vals map[string]cty.Value) {
	if vals["subscription"].IsNull() {
		*p = Subscription{}
		return
	}
	rvals := ctwhy.ValueAsList(vals["subscription"])
	if len(rvals) == 0 {
		*p = Subscription{}
		return
	}
	// this template should be used when single dictionary/object values are nested in sets/lists
	// if rvals turns out to be a list with > 1 elements, something has broken with that heuristic
	valMap := rvals[0].AsValueMap()
	DecodeContentLibrary_Subscription_OnDemand(p, valMap)
	DecodeContentLibrary_Subscription_Password(p, valMap)
	DecodeContentLibrary_Subscription_SubscriptionUrl(p, valMap)
	DecodeContentLibrary_Subscription_Username(p, valMap)
	DecodeContentLibrary_Subscription_AuthenticationMethod(p, valMap)
	DecodeContentLibrary_Subscription_AutomaticSync(p, valMap)
}

//primitiveTypeDecodeTemplate
func DecodeContentLibrary_Subscription_OnDemand(p *Subscription, vals map[string]cty.Value) {
	p.OnDemand = ctwhy.ValueAsBool(vals["on_demand"])
}

//primitiveTypeDecodeTemplate
func DecodeContentLibrary_Subscription_Password(p *Subscription, vals map[string]cty.Value) {
	p.Password = ctwhy.ValueAsString(vals["password"])
}

//primitiveTypeDecodeTemplate
func DecodeContentLibrary_Subscription_SubscriptionUrl(p *Subscription, vals map[string]cty.Value) {
	p.SubscriptionUrl = ctwhy.ValueAsString(vals["subscription_url"])
}

//primitiveTypeDecodeTemplate
func DecodeContentLibrary_Subscription_Username(p *Subscription, vals map[string]cty.Value) {
	p.Username = ctwhy.ValueAsString(vals["username"])
}

//primitiveTypeDecodeTemplate
func DecodeContentLibrary_Subscription_AuthenticationMethod(p *Subscription, vals map[string]cty.Value) {
	p.AuthenticationMethod = ctwhy.ValueAsString(vals["authentication_method"])
}

//primitiveTypeDecodeTemplate
func DecodeContentLibrary_Subscription_AutomaticSync(p *Subscription, vals map[string]cty.Value) {
	p.AutomaticSync = ctwhy.ValueAsBool(vals["automatic_sync"])
}

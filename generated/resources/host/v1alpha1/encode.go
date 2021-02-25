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
	r, ok := mr.(*Host)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a Host.")
	}
	return EncodeHost(*r), nil
}

func EncodeHost(r Host) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeHost_Cluster(r.Spec.ForProvider, ctyVal)
	EncodeHost_ClusterManaged(r.Spec.ForProvider, ctyVal)
	EncodeHost_Connected(r.Spec.ForProvider, ctyVal)
	EncodeHost_Datacenter(r.Spec.ForProvider, ctyVal)
	EncodeHost_Force(r.Spec.ForProvider, ctyVal)
	EncodeHost_Hostname(r.Spec.ForProvider, ctyVal)
	EncodeHost_License(r.Spec.ForProvider, ctyVal)
	EncodeHost_Lockdown(r.Spec.ForProvider, ctyVal)
	EncodeHost_Maintenance(r.Spec.ForProvider, ctyVal)
	EncodeHost_Password(r.Spec.ForProvider, ctyVal)
	EncodeHost_Thumbprint(r.Spec.ForProvider, ctyVal)
	EncodeHost_Username(r.Spec.ForProvider, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeHost_Cluster(p HostParameters, vals map[string]cty.Value) {
	vals["cluster"] = cty.StringVal(p.Cluster)
}

func EncodeHost_ClusterManaged(p HostParameters, vals map[string]cty.Value) {
	vals["cluster_managed"] = cty.BoolVal(p.ClusterManaged)
}

func EncodeHost_Connected(p HostParameters, vals map[string]cty.Value) {
	vals["connected"] = cty.BoolVal(p.Connected)
}

func EncodeHost_Datacenter(p HostParameters, vals map[string]cty.Value) {
	vals["datacenter"] = cty.StringVal(p.Datacenter)
}

func EncodeHost_Force(p HostParameters, vals map[string]cty.Value) {
	vals["force"] = cty.BoolVal(p.Force)
}

func EncodeHost_Hostname(p HostParameters, vals map[string]cty.Value) {
	vals["hostname"] = cty.StringVal(p.Hostname)
}

func EncodeHost_License(p HostParameters, vals map[string]cty.Value) {
	vals["license"] = cty.StringVal(p.License)
}

func EncodeHost_Lockdown(p HostParameters, vals map[string]cty.Value) {
	vals["lockdown"] = cty.StringVal(p.Lockdown)
}

func EncodeHost_Maintenance(p HostParameters, vals map[string]cty.Value) {
	vals["maintenance"] = cty.BoolVal(p.Maintenance)
}

func EncodeHost_Password(p HostParameters, vals map[string]cty.Value) {
	vals["password"] = cty.StringVal(p.Password)
}

func EncodeHost_Thumbprint(p HostParameters, vals map[string]cty.Value) {
	vals["thumbprint"] = cty.StringVal(p.Thumbprint)
}

func EncodeHost_Username(p HostParameters, vals map[string]cty.Value) {
	vals["username"] = cty.StringVal(p.Username)
}

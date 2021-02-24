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
	r, ok := mr.(*Host)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeHost(r, ctyValue)
}

func DecodeHost(prev *Host, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeHost_Cluster(&new.Spec.ForProvider, valMap)
	DecodeHost_ClusterManaged(&new.Spec.ForProvider, valMap)
	DecodeHost_Connected(&new.Spec.ForProvider, valMap)
	DecodeHost_Datacenter(&new.Spec.ForProvider, valMap)
	DecodeHost_Force(&new.Spec.ForProvider, valMap)
	DecodeHost_Hostname(&new.Spec.ForProvider, valMap)
	DecodeHost_License(&new.Spec.ForProvider, valMap)
	DecodeHost_Lockdown(&new.Spec.ForProvider, valMap)
	DecodeHost_Maintenance(&new.Spec.ForProvider, valMap)
	DecodeHost_Password(&new.Spec.ForProvider, valMap)
	DecodeHost_Thumbprint(&new.Spec.ForProvider, valMap)
	DecodeHost_Username(&new.Spec.ForProvider, valMap)

	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeHost_Cluster(p *HostParameters, vals map[string]cty.Value) {
	p.Cluster = ctwhy.ValueAsString(vals["cluster"])
}

//primitiveTypeDecodeTemplate
func DecodeHost_ClusterManaged(p *HostParameters, vals map[string]cty.Value) {
	p.ClusterManaged = ctwhy.ValueAsBool(vals["cluster_managed"])
}

//primitiveTypeDecodeTemplate
func DecodeHost_Connected(p *HostParameters, vals map[string]cty.Value) {
	p.Connected = ctwhy.ValueAsBool(vals["connected"])
}

//primitiveTypeDecodeTemplate
func DecodeHost_Datacenter(p *HostParameters, vals map[string]cty.Value) {
	p.Datacenter = ctwhy.ValueAsString(vals["datacenter"])
}

//primitiveTypeDecodeTemplate
func DecodeHost_Force(p *HostParameters, vals map[string]cty.Value) {
	p.Force = ctwhy.ValueAsBool(vals["force"])
}

//primitiveTypeDecodeTemplate
func DecodeHost_Hostname(p *HostParameters, vals map[string]cty.Value) {
	p.Hostname = ctwhy.ValueAsString(vals["hostname"])
}

//primitiveTypeDecodeTemplate
func DecodeHost_License(p *HostParameters, vals map[string]cty.Value) {
	p.License = ctwhy.ValueAsString(vals["license"])
}

//primitiveTypeDecodeTemplate
func DecodeHost_Lockdown(p *HostParameters, vals map[string]cty.Value) {
	p.Lockdown = ctwhy.ValueAsString(vals["lockdown"])
}

//primitiveTypeDecodeTemplate
func DecodeHost_Maintenance(p *HostParameters, vals map[string]cty.Value) {
	p.Maintenance = ctwhy.ValueAsBool(vals["maintenance"])
}

//primitiveTypeDecodeTemplate
func DecodeHost_Password(p *HostParameters, vals map[string]cty.Value) {
	p.Password = ctwhy.ValueAsString(vals["password"])
}

//primitiveTypeDecodeTemplate
func DecodeHost_Thumbprint(p *HostParameters, vals map[string]cty.Value) {
	p.Thumbprint = ctwhy.ValueAsString(vals["thumbprint"])
}

//primitiveTypeDecodeTemplate
func DecodeHost_Username(p *HostParameters, vals map[string]cty.Value) {
	p.Username = ctwhy.ValueAsString(vals["username"])
}
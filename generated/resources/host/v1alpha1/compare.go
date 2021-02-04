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
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane-contrib/terraform-runtime/pkg/plugin"
)

//mergeManagedResourceEntrypointTemplate
type resourceMerger struct{}

func (r *resourceMerger) MergeResources(kube resource.Managed, prov resource.Managed) plugin.MergeDescription {
	k := kube.(*Host)
	p := prov.(*Host)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeHost_Datacenter(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHost_Username(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHost_Cluster(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHost_Connected(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHost_Hostname(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHost_License(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHost_Lockdown(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHost_Maintenance(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHost_Password(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHost_ClusterManaged(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHost_Force(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHost_Thumbprint(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}


	for key, v := range p.Annotations {
		if k.Annotations[key] != v {
			k.Annotations[key] = v
			md.AnnotationsUpdated = true
		}
	}
	md.AnyFieldUpdated = anyChildUpdated
	return *md
}

//mergePrimitiveTemplateSpec
func MergeHost_Datacenter(k *HostParameters, p *HostParameters, md *plugin.MergeDescription) bool {
	if k.Datacenter != p.Datacenter {
		p.Datacenter = k.Datacenter
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHost_Username(k *HostParameters, p *HostParameters, md *plugin.MergeDescription) bool {
	if k.Username != p.Username {
		p.Username = k.Username
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHost_Cluster(k *HostParameters, p *HostParameters, md *plugin.MergeDescription) bool {
	if k.Cluster != p.Cluster {
		p.Cluster = k.Cluster
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHost_Connected(k *HostParameters, p *HostParameters, md *plugin.MergeDescription) bool {
	if k.Connected != p.Connected {
		p.Connected = k.Connected
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHost_Hostname(k *HostParameters, p *HostParameters, md *plugin.MergeDescription) bool {
	if k.Hostname != p.Hostname {
		p.Hostname = k.Hostname
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHost_License(k *HostParameters, p *HostParameters, md *plugin.MergeDescription) bool {
	if k.License != p.License {
		p.License = k.License
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHost_Lockdown(k *HostParameters, p *HostParameters, md *plugin.MergeDescription) bool {
	if k.Lockdown != p.Lockdown {
		p.Lockdown = k.Lockdown
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHost_Maintenance(k *HostParameters, p *HostParameters, md *plugin.MergeDescription) bool {
	if k.Maintenance != p.Maintenance {
		p.Maintenance = k.Maintenance
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHost_Password(k *HostParameters, p *HostParameters, md *plugin.MergeDescription) bool {
	if k.Password != p.Password {
		p.Password = k.Password
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHost_ClusterManaged(k *HostParameters, p *HostParameters, md *plugin.MergeDescription) bool {
	if k.ClusterManaged != p.ClusterManaged {
		p.ClusterManaged = k.ClusterManaged
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHost_Force(k *HostParameters, p *HostParameters, md *plugin.MergeDescription) bool {
	if k.Force != p.Force {
		p.Force = k.Force
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHost_Thumbprint(k *HostParameters, p *HostParameters, md *plugin.MergeDescription) bool {
	if k.Thumbprint != p.Thumbprint {
		p.Thumbprint = k.Thumbprint
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}
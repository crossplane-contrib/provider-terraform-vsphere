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
	"github.com/crossplane-contrib/terraform-runtime/pkg/plugin"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
)

//mergeManagedResourceEntrypointTemplate
type resourceMerger struct{}

func (r *resourceMerger) MergeResources(kube resource.Managed, prov resource.Managed) plugin.MergeDescription {
	k := kube.(*ContentLibrary)
	p := prov.(*ContentLibrary)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeContentLibrary_Description(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeContentLibrary_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeContentLibrary_StorageBacking(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeContentLibrary_Publication(&k.Spec.ForProvider.Publication, &p.Spec.ForProvider.Publication, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeContentLibrary_Subscription(&k.Spec.ForProvider.Subscription, &p.Spec.ForProvider.Subscription, md)
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
func MergeContentLibrary_Description(k *ContentLibraryParameters, p *ContentLibraryParameters, md *plugin.MergeDescription) bool {
	if k.Description != p.Description {
		p.Description = k.Description
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeContentLibrary_Name(k *ContentLibraryParameters, p *ContentLibraryParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeContentLibrary_StorageBacking(k *ContentLibraryParameters, p *ContentLibraryParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.StorageBacking, p.StorageBacking) {
		p.StorageBacking = k.StorageBacking
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructTemplateSpec
func MergeContentLibrary_Publication(k *Publication, p *Publication, md *plugin.MergeDescription) bool {
	updated := false
	anyChildUpdated := false
	updated = MergeContentLibrary_Publication_Published(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeContentLibrary_Publication_Username(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeContentLibrary_Publication_AuthenticationMethod(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeContentLibrary_Publication_Password(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeContentLibrary_Publication_PublishUrl(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveTemplateSpec
func MergeContentLibrary_Publication_Published(k *Publication, p *Publication, md *plugin.MergeDescription) bool {
	if k.Published != p.Published {
		p.Published = k.Published
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeContentLibrary_Publication_Username(k *Publication, p *Publication, md *plugin.MergeDescription) bool {
	if k.Username != p.Username {
		p.Username = k.Username
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeContentLibrary_Publication_AuthenticationMethod(k *Publication, p *Publication, md *plugin.MergeDescription) bool {
	if k.AuthenticationMethod != p.AuthenticationMethod {
		p.AuthenticationMethod = k.AuthenticationMethod
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeContentLibrary_Publication_Password(k *Publication, p *Publication, md *plugin.MergeDescription) bool {
	if k.Password != p.Password {
		p.Password = k.Password
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeContentLibrary_Publication_PublishUrl(k *Publication, p *Publication, md *plugin.MergeDescription) bool {
	if k.PublishUrl != p.PublishUrl {
		p.PublishUrl = k.PublishUrl
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructTemplateSpec
func MergeContentLibrary_Subscription(k *Subscription, p *Subscription, md *plugin.MergeDescription) bool {
	updated := false
	anyChildUpdated := false
	updated = MergeContentLibrary_Subscription_OnDemand(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeContentLibrary_Subscription_Password(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeContentLibrary_Subscription_SubscriptionUrl(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeContentLibrary_Subscription_Username(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeContentLibrary_Subscription_AuthenticationMethod(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeContentLibrary_Subscription_AutomaticSync(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveTemplateSpec
func MergeContentLibrary_Subscription_OnDemand(k *Subscription, p *Subscription, md *plugin.MergeDescription) bool {
	if k.OnDemand != p.OnDemand {
		p.OnDemand = k.OnDemand
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeContentLibrary_Subscription_Password(k *Subscription, p *Subscription, md *plugin.MergeDescription) bool {
	if k.Password != p.Password {
		p.Password = k.Password
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeContentLibrary_Subscription_SubscriptionUrl(k *Subscription, p *Subscription, md *plugin.MergeDescription) bool {
	if k.SubscriptionUrl != p.SubscriptionUrl {
		p.SubscriptionUrl = k.SubscriptionUrl
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeContentLibrary_Subscription_Username(k *Subscription, p *Subscription, md *plugin.MergeDescription) bool {
	if k.Username != p.Username {
		p.Username = k.Username
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeContentLibrary_Subscription_AuthenticationMethod(k *Subscription, p *Subscription, md *plugin.MergeDescription) bool {
	if k.AuthenticationMethod != p.AuthenticationMethod {
		p.AuthenticationMethod = k.AuthenticationMethod
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeContentLibrary_Subscription_AutomaticSync(k *Subscription, p *Subscription, md *plugin.MergeDescription) bool {
	if k.AutomaticSync != p.AutomaticSync {
		p.AutomaticSync = k.AutomaticSync
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

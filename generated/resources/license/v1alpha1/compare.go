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
	k := kube.(*License)
	p := prov.(*License)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeLicense_Labels(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLicense_LicenseKey(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLicense_Used(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLicense_EditionKey(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLicense_Name(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeLicense_Total(&k.Status.AtProvider, &p.Status.AtProvider, md)
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

//mergePrimitiveContainerTemplateSpec
func MergeLicense_Labels(k *LicenseParameters, p *LicenseParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Labels, p.Labels) {
		p.Labels = k.Labels
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeLicense_LicenseKey(k *LicenseParameters, p *LicenseParameters, md *plugin.MergeDescription) bool {
	if k.LicenseKey != p.LicenseKey {
		p.LicenseKey = k.LicenseKey
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeLicense_Used(k *LicenseObservation, p *LicenseObservation, md *plugin.MergeDescription) bool {
	if k.Used != p.Used {
		k.Used = p.Used
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeLicense_EditionKey(k *LicenseObservation, p *LicenseObservation, md *plugin.MergeDescription) bool {
	if k.EditionKey != p.EditionKey {
		k.EditionKey = p.EditionKey
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeLicense_Name(k *LicenseObservation, p *LicenseObservation, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		k.Name = p.Name
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeLicense_Total(k *LicenseObservation, p *LicenseObservation, md *plugin.MergeDescription) bool {
	if k.Total != p.Total {
		k.Total = p.Total
		md.StatusUpdated = true
		return true
	}
	return false
}
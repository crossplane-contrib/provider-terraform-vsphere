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
	k := kube.(*File)
	p := prov.(*File)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeFile_Datacenter(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeFile_Datastore(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeFile_DestinationFile(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeFile_SourceDatacenter(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeFile_SourceDatastore(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeFile_SourceFile(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeFile_CreateDirectories(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeFile_Datacenter(k *FileParameters, p *FileParameters, md *plugin.MergeDescription) bool {
	if k.Datacenter != p.Datacenter {
		p.Datacenter = k.Datacenter
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeFile_Datastore(k *FileParameters, p *FileParameters, md *plugin.MergeDescription) bool {
	if k.Datastore != p.Datastore {
		p.Datastore = k.Datastore
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeFile_DestinationFile(k *FileParameters, p *FileParameters, md *plugin.MergeDescription) bool {
	if k.DestinationFile != p.DestinationFile {
		p.DestinationFile = k.DestinationFile
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeFile_SourceDatacenter(k *FileParameters, p *FileParameters, md *plugin.MergeDescription) bool {
	if k.SourceDatacenter != p.SourceDatacenter {
		p.SourceDatacenter = k.SourceDatacenter
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeFile_SourceDatastore(k *FileParameters, p *FileParameters, md *plugin.MergeDescription) bool {
	if k.SourceDatastore != p.SourceDatastore {
		p.SourceDatastore = k.SourceDatastore
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeFile_SourceFile(k *FileParameters, p *FileParameters, md *plugin.MergeDescription) bool {
	if k.SourceFile != p.SourceFile {
		p.SourceFile = k.SourceFile
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeFile_CreateDirectories(k *FileParameters, p *FileParameters, md *plugin.MergeDescription) bool {
	if k.CreateDirectories != p.CreateDirectories {
		p.CreateDirectories = k.CreateDirectories
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}
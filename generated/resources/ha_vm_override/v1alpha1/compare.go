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
	k := kube.(*HaVmOverride)
	p := prov.(*HaVmOverride)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeHaVmOverride_ComputeClusterId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHaVmOverride_HaDatastoreApdRecoveryAction(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHaVmOverride_HaDatastoreApdResponse(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHaVmOverride_HaDatastoreApdResponseDelay(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHaVmOverride_HaDatastorePdlResponse(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHaVmOverride_HaHostIsolationResponse(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHaVmOverride_HaVmFailureInterval(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHaVmOverride_HaVmMaximumFailureWindow(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHaVmOverride_HaVmMaximumResets(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHaVmOverride_HaVmMinimumUptime(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHaVmOverride_HaVmMonitoring(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHaVmOverride_HaVmMonitoringUseClusterDefaults(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHaVmOverride_HaVmRestartPriority(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHaVmOverride_HaVmRestartTimeout(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeHaVmOverride_VirtualMachineId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
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
func MergeHaVmOverride_ComputeClusterId(k *HaVmOverrideParameters, p *HaVmOverrideParameters, md *plugin.MergeDescription) bool {
	if k.ComputeClusterId != p.ComputeClusterId {
		p.ComputeClusterId = k.ComputeClusterId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHaVmOverride_HaDatastoreApdRecoveryAction(k *HaVmOverrideParameters, p *HaVmOverrideParameters, md *plugin.MergeDescription) bool {
	if k.HaDatastoreApdRecoveryAction != p.HaDatastoreApdRecoveryAction {
		p.HaDatastoreApdRecoveryAction = k.HaDatastoreApdRecoveryAction
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHaVmOverride_HaDatastoreApdResponse(k *HaVmOverrideParameters, p *HaVmOverrideParameters, md *plugin.MergeDescription) bool {
	if k.HaDatastoreApdResponse != p.HaDatastoreApdResponse {
		p.HaDatastoreApdResponse = k.HaDatastoreApdResponse
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHaVmOverride_HaDatastoreApdResponseDelay(k *HaVmOverrideParameters, p *HaVmOverrideParameters, md *plugin.MergeDescription) bool {
	if k.HaDatastoreApdResponseDelay != p.HaDatastoreApdResponseDelay {
		p.HaDatastoreApdResponseDelay = k.HaDatastoreApdResponseDelay
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHaVmOverride_HaDatastorePdlResponse(k *HaVmOverrideParameters, p *HaVmOverrideParameters, md *plugin.MergeDescription) bool {
	if k.HaDatastorePdlResponse != p.HaDatastorePdlResponse {
		p.HaDatastorePdlResponse = k.HaDatastorePdlResponse
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHaVmOverride_HaHostIsolationResponse(k *HaVmOverrideParameters, p *HaVmOverrideParameters, md *plugin.MergeDescription) bool {
	if k.HaHostIsolationResponse != p.HaHostIsolationResponse {
		p.HaHostIsolationResponse = k.HaHostIsolationResponse
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHaVmOverride_HaVmFailureInterval(k *HaVmOverrideParameters, p *HaVmOverrideParameters, md *plugin.MergeDescription) bool {
	if k.HaVmFailureInterval != p.HaVmFailureInterval {
		p.HaVmFailureInterval = k.HaVmFailureInterval
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHaVmOverride_HaVmMaximumFailureWindow(k *HaVmOverrideParameters, p *HaVmOverrideParameters, md *plugin.MergeDescription) bool {
	if k.HaVmMaximumFailureWindow != p.HaVmMaximumFailureWindow {
		p.HaVmMaximumFailureWindow = k.HaVmMaximumFailureWindow
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHaVmOverride_HaVmMaximumResets(k *HaVmOverrideParameters, p *HaVmOverrideParameters, md *plugin.MergeDescription) bool {
	if k.HaVmMaximumResets != p.HaVmMaximumResets {
		p.HaVmMaximumResets = k.HaVmMaximumResets
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHaVmOverride_HaVmMinimumUptime(k *HaVmOverrideParameters, p *HaVmOverrideParameters, md *plugin.MergeDescription) bool {
	if k.HaVmMinimumUptime != p.HaVmMinimumUptime {
		p.HaVmMinimumUptime = k.HaVmMinimumUptime
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHaVmOverride_HaVmMonitoring(k *HaVmOverrideParameters, p *HaVmOverrideParameters, md *plugin.MergeDescription) bool {
	if k.HaVmMonitoring != p.HaVmMonitoring {
		p.HaVmMonitoring = k.HaVmMonitoring
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHaVmOverride_HaVmMonitoringUseClusterDefaults(k *HaVmOverrideParameters, p *HaVmOverrideParameters, md *plugin.MergeDescription) bool {
	if k.HaVmMonitoringUseClusterDefaults != p.HaVmMonitoringUseClusterDefaults {
		p.HaVmMonitoringUseClusterDefaults = k.HaVmMonitoringUseClusterDefaults
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHaVmOverride_HaVmRestartPriority(k *HaVmOverrideParameters, p *HaVmOverrideParameters, md *plugin.MergeDescription) bool {
	if k.HaVmRestartPriority != p.HaVmRestartPriority {
		p.HaVmRestartPriority = k.HaVmRestartPriority
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHaVmOverride_HaVmRestartTimeout(k *HaVmOverrideParameters, p *HaVmOverrideParameters, md *plugin.MergeDescription) bool {
	if k.HaVmRestartTimeout != p.HaVmRestartTimeout {
		p.HaVmRestartTimeout = k.HaVmRestartTimeout
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeHaVmOverride_VirtualMachineId(k *HaVmOverrideParameters, p *HaVmOverrideParameters, md *plugin.MergeDescription) bool {
	if k.VirtualMachineId != p.VirtualMachineId {
		p.VirtualMachineId = k.VirtualMachineId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}
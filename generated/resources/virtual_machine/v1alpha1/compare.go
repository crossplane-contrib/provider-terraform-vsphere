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
	k := kube.(*VirtualMachine)
	p := prov.(*VirtualMachine)
	md := &plugin.MergeDescription{}
	updated := false
	anyChildUpdated := false

	updated = MergeVirtualMachine_AlternateGuestName(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_Annotation(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_BootDelay(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_BootRetryDelay(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_BootRetryEnabled(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_CpuHotAddEnabled(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_CpuHotRemoveEnabled(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_CpuLimit(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_CpuPerformanceCountersEnabled(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_CpuReservation(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_CpuShareCount(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_CpuShareLevel(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_CustomAttributes(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_DatacenterId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_DatastoreClusterId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_DatastoreId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_EfiSecureBootEnabled(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_EnableDiskUuid(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_EnableLogging(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_EptRviMode(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_ExtraConfig(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_Firmware(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_Folder(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_ForcePowerOff(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_GuestId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_HardwareVersion(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_HostSystemId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_HvMode(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_IdeControllerCount(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_IgnoredGuestIps(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_LatencySensitivity(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_Memory(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_MemoryHotAddEnabled(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_MemoryLimit(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_MemoryReservation(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_MemoryShareCount(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_MemoryShareLevel(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_MigrateWaitTimeout(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_Name(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_NestedHvEnabled(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_NumCoresPerSocket(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_NumCpus(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_PciDeviceId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_PoweronTimeout(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_ResourcePoolId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_RunToolsScriptsAfterPowerOn(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_RunToolsScriptsAfterResume(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_RunToolsScriptsBeforeGuestReboot(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_RunToolsScriptsBeforeGuestShutdown(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_RunToolsScriptsBeforeGuestStandby(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_SataControllerCount(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_ScsiBusSharing(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_ScsiControllerCount(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_ScsiType(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_ShutdownWaitTimeout(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_StoragePolicyId(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_SwapPlacementPolicy(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_SyncTimeWithHost(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_Tags(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_WaitForGuestIpTimeout(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_WaitForGuestNetRoutable(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_WaitForGuestNetTimeout(&k.Spec.ForProvider, &p.Spec.ForProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_Cdrom(&k.Spec.ForProvider.Cdrom, &p.Spec.ForProvider.Cdrom, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_Clone(&k.Spec.ForProvider.Clone, &p.Spec.ForProvider.Clone, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_Disk(&k.Spec.ForProvider.Disk, &p.Spec.ForProvider.Disk, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_NetworkInterface(&k.Spec.ForProvider.NetworkInterface, &p.Spec.ForProvider.NetworkInterface, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_OvfDeploy(&k.Spec.ForProvider.OvfDeploy, &p.Spec.ForProvider.OvfDeploy, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_Vapp(&k.Spec.ForProvider.Vapp, &p.Spec.ForProvider.Vapp, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_ChangeVersion(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_DefaultIpAddress(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_GuestIpAddresses(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_Imported(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_Moid(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_RebootRequired(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_Uuid(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_VappTransport(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_VmwareToolsStatus(&k.Status.AtProvider, &p.Status.AtProvider, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_VmxPath(&k.Status.AtProvider, &p.Status.AtProvider, md)
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
func MergeVirtualMachine_AlternateGuestName(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.AlternateGuestName != p.AlternateGuestName {
		p.AlternateGuestName = k.AlternateGuestName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Annotation(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.Annotation != p.Annotation {
		p.Annotation = k.Annotation
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_BootDelay(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.BootDelay != p.BootDelay {
		p.BootDelay = k.BootDelay
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_BootRetryDelay(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.BootRetryDelay != p.BootRetryDelay {
		p.BootRetryDelay = k.BootRetryDelay
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_BootRetryEnabled(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.BootRetryEnabled != p.BootRetryEnabled {
		p.BootRetryEnabled = k.BootRetryEnabled
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_CpuHotAddEnabled(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.CpuHotAddEnabled != p.CpuHotAddEnabled {
		p.CpuHotAddEnabled = k.CpuHotAddEnabled
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_CpuHotRemoveEnabled(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.CpuHotRemoveEnabled != p.CpuHotRemoveEnabled {
		p.CpuHotRemoveEnabled = k.CpuHotRemoveEnabled
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_CpuLimit(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.CpuLimit != p.CpuLimit {
		p.CpuLimit = k.CpuLimit
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_CpuPerformanceCountersEnabled(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.CpuPerformanceCountersEnabled != p.CpuPerformanceCountersEnabled {
		p.CpuPerformanceCountersEnabled = k.CpuPerformanceCountersEnabled
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_CpuReservation(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.CpuReservation != p.CpuReservation {
		p.CpuReservation = k.CpuReservation
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_CpuShareCount(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.CpuShareCount != p.CpuShareCount {
		p.CpuShareCount = k.CpuShareCount
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_CpuShareLevel(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.CpuShareLevel != p.CpuShareLevel {
		p.CpuShareLevel = k.CpuShareLevel
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeVirtualMachine_CustomAttributes(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.CustomAttributes, p.CustomAttributes) {
		p.CustomAttributes = k.CustomAttributes
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_DatacenterId(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.DatacenterId != p.DatacenterId {
		p.DatacenterId = k.DatacenterId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_DatastoreClusterId(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.DatastoreClusterId != p.DatastoreClusterId {
		p.DatastoreClusterId = k.DatastoreClusterId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_DatastoreId(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.DatastoreId != p.DatastoreId {
		p.DatastoreId = k.DatastoreId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_EfiSecureBootEnabled(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.EfiSecureBootEnabled != p.EfiSecureBootEnabled {
		p.EfiSecureBootEnabled = k.EfiSecureBootEnabled
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_EnableDiskUuid(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.EnableDiskUuid != p.EnableDiskUuid {
		p.EnableDiskUuid = k.EnableDiskUuid
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_EnableLogging(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.EnableLogging != p.EnableLogging {
		p.EnableLogging = k.EnableLogging
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_EptRviMode(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.EptRviMode != p.EptRviMode {
		p.EptRviMode = k.EptRviMode
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeVirtualMachine_ExtraConfig(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.ExtraConfig, p.ExtraConfig) {
		p.ExtraConfig = k.ExtraConfig
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Firmware(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.Firmware != p.Firmware {
		p.Firmware = k.Firmware
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Folder(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.Folder != p.Folder {
		p.Folder = k.Folder
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_ForcePowerOff(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.ForcePowerOff != p.ForcePowerOff {
		p.ForcePowerOff = k.ForcePowerOff
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_GuestId(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.GuestId != p.GuestId {
		p.GuestId = k.GuestId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_HardwareVersion(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.HardwareVersion != p.HardwareVersion {
		p.HardwareVersion = k.HardwareVersion
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_HostSystemId(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.HostSystemId != p.HostSystemId {
		p.HostSystemId = k.HostSystemId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_HvMode(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.HvMode != p.HvMode {
		p.HvMode = k.HvMode
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_IdeControllerCount(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.IdeControllerCount != p.IdeControllerCount {
		p.IdeControllerCount = k.IdeControllerCount
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeVirtualMachine_IgnoredGuestIps(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.IgnoredGuestIps, p.IgnoredGuestIps) {
		p.IgnoredGuestIps = k.IgnoredGuestIps
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_LatencySensitivity(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.LatencySensitivity != p.LatencySensitivity {
		p.LatencySensitivity = k.LatencySensitivity
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Memory(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.Memory != p.Memory {
		p.Memory = k.Memory
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_MemoryHotAddEnabled(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.MemoryHotAddEnabled != p.MemoryHotAddEnabled {
		p.MemoryHotAddEnabled = k.MemoryHotAddEnabled
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_MemoryLimit(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.MemoryLimit != p.MemoryLimit {
		p.MemoryLimit = k.MemoryLimit
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_MemoryReservation(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.MemoryReservation != p.MemoryReservation {
		p.MemoryReservation = k.MemoryReservation
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_MemoryShareCount(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.MemoryShareCount != p.MemoryShareCount {
		p.MemoryShareCount = k.MemoryShareCount
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_MemoryShareLevel(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.MemoryShareLevel != p.MemoryShareLevel {
		p.MemoryShareLevel = k.MemoryShareLevel
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_MigrateWaitTimeout(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.MigrateWaitTimeout != p.MigrateWaitTimeout {
		p.MigrateWaitTimeout = k.MigrateWaitTimeout
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Name(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_NestedHvEnabled(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.NestedHvEnabled != p.NestedHvEnabled {
		p.NestedHvEnabled = k.NestedHvEnabled
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_NumCoresPerSocket(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.NumCoresPerSocket != p.NumCoresPerSocket {
		p.NumCoresPerSocket = k.NumCoresPerSocket
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_NumCpus(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.NumCpus != p.NumCpus {
		p.NumCpus = k.NumCpus
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeVirtualMachine_PciDeviceId(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.PciDeviceId, p.PciDeviceId) {
		p.PciDeviceId = k.PciDeviceId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_PoweronTimeout(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.PoweronTimeout != p.PoweronTimeout {
		p.PoweronTimeout = k.PoweronTimeout
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_ResourcePoolId(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.ResourcePoolId != p.ResourcePoolId {
		p.ResourcePoolId = k.ResourcePoolId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_RunToolsScriptsAfterPowerOn(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.RunToolsScriptsAfterPowerOn != p.RunToolsScriptsAfterPowerOn {
		p.RunToolsScriptsAfterPowerOn = k.RunToolsScriptsAfterPowerOn
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_RunToolsScriptsAfterResume(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.RunToolsScriptsAfterResume != p.RunToolsScriptsAfterResume {
		p.RunToolsScriptsAfterResume = k.RunToolsScriptsAfterResume
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_RunToolsScriptsBeforeGuestReboot(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.RunToolsScriptsBeforeGuestReboot != p.RunToolsScriptsBeforeGuestReboot {
		p.RunToolsScriptsBeforeGuestReboot = k.RunToolsScriptsBeforeGuestReboot
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_RunToolsScriptsBeforeGuestShutdown(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.RunToolsScriptsBeforeGuestShutdown != p.RunToolsScriptsBeforeGuestShutdown {
		p.RunToolsScriptsBeforeGuestShutdown = k.RunToolsScriptsBeforeGuestShutdown
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_RunToolsScriptsBeforeGuestStandby(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.RunToolsScriptsBeforeGuestStandby != p.RunToolsScriptsBeforeGuestStandby {
		p.RunToolsScriptsBeforeGuestStandby = k.RunToolsScriptsBeforeGuestStandby
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_SataControllerCount(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.SataControllerCount != p.SataControllerCount {
		p.SataControllerCount = k.SataControllerCount
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_ScsiBusSharing(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.ScsiBusSharing != p.ScsiBusSharing {
		p.ScsiBusSharing = k.ScsiBusSharing
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_ScsiControllerCount(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.ScsiControllerCount != p.ScsiControllerCount {
		p.ScsiControllerCount = k.ScsiControllerCount
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_ScsiType(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.ScsiType != p.ScsiType {
		p.ScsiType = k.ScsiType
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_ShutdownWaitTimeout(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.ShutdownWaitTimeout != p.ShutdownWaitTimeout {
		p.ShutdownWaitTimeout = k.ShutdownWaitTimeout
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_StoragePolicyId(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.StoragePolicyId != p.StoragePolicyId {
		p.StoragePolicyId = k.StoragePolicyId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_SwapPlacementPolicy(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.SwapPlacementPolicy != p.SwapPlacementPolicy {
		p.SwapPlacementPolicy = k.SwapPlacementPolicy
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_SyncTimeWithHost(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.SyncTimeWithHost != p.SyncTimeWithHost {
		p.SyncTimeWithHost = k.SyncTimeWithHost
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeVirtualMachine_Tags(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.Tags, p.Tags) {
		p.Tags = k.Tags
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_WaitForGuestIpTimeout(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.WaitForGuestIpTimeout != p.WaitForGuestIpTimeout {
		p.WaitForGuestIpTimeout = k.WaitForGuestIpTimeout
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_WaitForGuestNetRoutable(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.WaitForGuestNetRoutable != p.WaitForGuestNetRoutable {
		p.WaitForGuestNetRoutable = k.WaitForGuestNetRoutable
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_WaitForGuestNetTimeout(k *VirtualMachineParameters, p *VirtualMachineParameters, md *plugin.MergeDescription) bool {
	if k.WaitForGuestNetTimeout != p.WaitForGuestNetTimeout {
		p.WaitForGuestNetTimeout = k.WaitForGuestNetTimeout
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructTemplateSpec
func MergeVirtualMachine_Cdrom(k *Cdrom, p *Cdrom, md *plugin.MergeDescription) bool {
	updated := false
	anyChildUpdated := false
	updated = MergeVirtualMachine_Cdrom_ClientDevice(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_Cdrom_DatastoreId(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_Cdrom_DeviceAddress(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_Cdrom_Key(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_Cdrom_Path(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Cdrom_ClientDevice(k *Cdrom, p *Cdrom, md *plugin.MergeDescription) bool {
	if k.ClientDevice != p.ClientDevice {
		p.ClientDevice = k.ClientDevice
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Cdrom_DatastoreId(k *Cdrom, p *Cdrom, md *plugin.MergeDescription) bool {
	if k.DatastoreId != p.DatastoreId {
		p.DatastoreId = k.DatastoreId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Cdrom_DeviceAddress(k *Cdrom, p *Cdrom, md *plugin.MergeDescription) bool {
	if k.DeviceAddress != p.DeviceAddress {
		p.DeviceAddress = k.DeviceAddress
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Cdrom_Key(k *Cdrom, p *Cdrom, md *plugin.MergeDescription) bool {
	if k.Key != p.Key {
		p.Key = k.Key
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Cdrom_Path(k *Cdrom, p *Cdrom, md *plugin.MergeDescription) bool {
	if k.Path != p.Path {
		p.Path = k.Path
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructTemplateSpec
func MergeVirtualMachine_Clone(k *Clone, p *Clone, md *plugin.MergeDescription) bool {
	updated := false
	anyChildUpdated := false
	updated = MergeVirtualMachine_Clone_LinkedClone(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_Clone_OvfNetworkMap(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_Clone_OvfStorageMap(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_Clone_TemplateUuid(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_Clone_Timeout(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_Clone_Customize(&k.Customize, &p.Customize, md)
	if updated {
		anyChildUpdated = true
	}

	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Clone_LinkedClone(k *Clone, p *Clone, md *plugin.MergeDescription) bool {
	if k.LinkedClone != p.LinkedClone {
		p.LinkedClone = k.LinkedClone
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeVirtualMachine_Clone_OvfNetworkMap(k *Clone, p *Clone, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.OvfNetworkMap, p.OvfNetworkMap) {
		p.OvfNetworkMap = k.OvfNetworkMap
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeVirtualMachine_Clone_OvfStorageMap(k *Clone, p *Clone, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.OvfStorageMap, p.OvfStorageMap) {
		p.OvfStorageMap = k.OvfStorageMap
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Clone_TemplateUuid(k *Clone, p *Clone, md *plugin.MergeDescription) bool {
	if k.TemplateUuid != p.TemplateUuid {
		p.TemplateUuid = k.TemplateUuid
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Clone_Timeout(k *Clone, p *Clone, md *plugin.MergeDescription) bool {
	if k.Timeout != p.Timeout {
		p.Timeout = k.Timeout
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructTemplateSpec
func MergeVirtualMachine_Clone_Customize(k *Customize, p *Customize, md *plugin.MergeDescription) bool {
	updated := false
	anyChildUpdated := false
	updated = MergeVirtualMachine_Clone_Customize_DnsServerList(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_Clone_Customize_DnsSuffixList(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_Clone_Customize_Ipv4Gateway(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_Clone_Customize_Ipv6Gateway(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_Clone_Customize_Timeout(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_Clone_Customize_WindowsSysprepText(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_Clone_Customize_LinuxOptions(&k.LinuxOptions, &p.LinuxOptions, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_Clone_Customize_NetworkInterface(&k.NetworkInterface, &p.NetworkInterface, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_Clone_Customize_WindowsOptions(&k.WindowsOptions, &p.WindowsOptions, md)
	if updated {
		anyChildUpdated = true
	}

	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveContainerTemplateSpec
func MergeVirtualMachine_Clone_Customize_DnsServerList(k *Customize, p *Customize, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.DnsServerList, p.DnsServerList) {
		p.DnsServerList = k.DnsServerList
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeVirtualMachine_Clone_Customize_DnsSuffixList(k *Customize, p *Customize, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.DnsSuffixList, p.DnsSuffixList) {
		p.DnsSuffixList = k.DnsSuffixList
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Clone_Customize_Ipv4Gateway(k *Customize, p *Customize, md *plugin.MergeDescription) bool {
	if k.Ipv4Gateway != p.Ipv4Gateway {
		p.Ipv4Gateway = k.Ipv4Gateway
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Clone_Customize_Ipv6Gateway(k *Customize, p *Customize, md *plugin.MergeDescription) bool {
	if k.Ipv6Gateway != p.Ipv6Gateway {
		p.Ipv6Gateway = k.Ipv6Gateway
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Clone_Customize_Timeout(k *Customize, p *Customize, md *plugin.MergeDescription) bool {
	if k.Timeout != p.Timeout {
		p.Timeout = k.Timeout
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Clone_Customize_WindowsSysprepText(k *Customize, p *Customize, md *plugin.MergeDescription) bool {
	if k.WindowsSysprepText != p.WindowsSysprepText {
		p.WindowsSysprepText = k.WindowsSysprepText
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructTemplateSpec
func MergeVirtualMachine_Clone_Customize_LinuxOptions(k *LinuxOptions, p *LinuxOptions, md *plugin.MergeDescription) bool {
	updated := false
	anyChildUpdated := false
	updated = MergeVirtualMachine_Clone_Customize_LinuxOptions_Domain(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_Clone_Customize_LinuxOptions_HostName(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_Clone_Customize_LinuxOptions_HwClockUtc(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_Clone_Customize_LinuxOptions_TimeZone(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Clone_Customize_LinuxOptions_Domain(k *LinuxOptions, p *LinuxOptions, md *plugin.MergeDescription) bool {
	if k.Domain != p.Domain {
		p.Domain = k.Domain
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Clone_Customize_LinuxOptions_HostName(k *LinuxOptions, p *LinuxOptions, md *plugin.MergeDescription) bool {
	if k.HostName != p.HostName {
		p.HostName = k.HostName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Clone_Customize_LinuxOptions_HwClockUtc(k *LinuxOptions, p *LinuxOptions, md *plugin.MergeDescription) bool {
	if k.HwClockUtc != p.HwClockUtc {
		p.HwClockUtc = k.HwClockUtc
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Clone_Customize_LinuxOptions_TimeZone(k *LinuxOptions, p *LinuxOptions, md *plugin.MergeDescription) bool {
	if k.TimeZone != p.TimeZone {
		p.TimeZone = k.TimeZone
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructTemplateSpec
func MergeVirtualMachine_Clone_Customize_NetworkInterface(k *NetworkInterface, p *NetworkInterface, md *plugin.MergeDescription) bool {
	updated := false
	anyChildUpdated := false
	updated = MergeVirtualMachine_Clone_Customize_NetworkInterface_DnsDomain(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_Clone_Customize_NetworkInterface_DnsServerList(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_Clone_Customize_NetworkInterface_Ipv4Address(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_Clone_Customize_NetworkInterface_Ipv4Netmask(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_Clone_Customize_NetworkInterface_Ipv6Address(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_Clone_Customize_NetworkInterface_Ipv6Netmask(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Clone_Customize_NetworkInterface_DnsDomain(k *NetworkInterface, p *NetworkInterface, md *plugin.MergeDescription) bool {
	if k.DnsDomain != p.DnsDomain {
		p.DnsDomain = k.DnsDomain
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeVirtualMachine_Clone_Customize_NetworkInterface_DnsServerList(k *NetworkInterface, p *NetworkInterface, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.DnsServerList, p.DnsServerList) {
		p.DnsServerList = k.DnsServerList
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Clone_Customize_NetworkInterface_Ipv4Address(k *NetworkInterface, p *NetworkInterface, md *plugin.MergeDescription) bool {
	if k.Ipv4Address != p.Ipv4Address {
		p.Ipv4Address = k.Ipv4Address
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Clone_Customize_NetworkInterface_Ipv4Netmask(k *NetworkInterface, p *NetworkInterface, md *plugin.MergeDescription) bool {
	if k.Ipv4Netmask != p.Ipv4Netmask {
		p.Ipv4Netmask = k.Ipv4Netmask
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Clone_Customize_NetworkInterface_Ipv6Address(k *NetworkInterface, p *NetworkInterface, md *plugin.MergeDescription) bool {
	if k.Ipv6Address != p.Ipv6Address {
		p.Ipv6Address = k.Ipv6Address
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Clone_Customize_NetworkInterface_Ipv6Netmask(k *NetworkInterface, p *NetworkInterface, md *plugin.MergeDescription) bool {
	if k.Ipv6Netmask != p.Ipv6Netmask {
		p.Ipv6Netmask = k.Ipv6Netmask
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructTemplateSpec
func MergeVirtualMachine_Clone_Customize_WindowsOptions(k *WindowsOptions, p *WindowsOptions, md *plugin.MergeDescription) bool {
	updated := false
	anyChildUpdated := false
	updated = MergeVirtualMachine_Clone_Customize_WindowsOptions_DomainAdminPassword(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_Clone_Customize_WindowsOptions_DomainAdminUser(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_Clone_Customize_WindowsOptions_JoinDomain(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_Clone_Customize_WindowsOptions_ProductKey(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_Clone_Customize_WindowsOptions_AdminPassword(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_Clone_Customize_WindowsOptions_AutoLogon(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_Clone_Customize_WindowsOptions_AutoLogonCount(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_Clone_Customize_WindowsOptions_ComputerName(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_Clone_Customize_WindowsOptions_RunOnceCommandList(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_Clone_Customize_WindowsOptions_Workgroup(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_Clone_Customize_WindowsOptions_FullName(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_Clone_Customize_WindowsOptions_OrganizationName(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_Clone_Customize_WindowsOptions_TimeZone(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Clone_Customize_WindowsOptions_DomainAdminPassword(k *WindowsOptions, p *WindowsOptions, md *plugin.MergeDescription) bool {
	if k.DomainAdminPassword != p.DomainAdminPassword {
		p.DomainAdminPassword = k.DomainAdminPassword
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Clone_Customize_WindowsOptions_DomainAdminUser(k *WindowsOptions, p *WindowsOptions, md *plugin.MergeDescription) bool {
	if k.DomainAdminUser != p.DomainAdminUser {
		p.DomainAdminUser = k.DomainAdminUser
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Clone_Customize_WindowsOptions_JoinDomain(k *WindowsOptions, p *WindowsOptions, md *plugin.MergeDescription) bool {
	if k.JoinDomain != p.JoinDomain {
		p.JoinDomain = k.JoinDomain
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Clone_Customize_WindowsOptions_ProductKey(k *WindowsOptions, p *WindowsOptions, md *plugin.MergeDescription) bool {
	if k.ProductKey != p.ProductKey {
		p.ProductKey = k.ProductKey
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Clone_Customize_WindowsOptions_AdminPassword(k *WindowsOptions, p *WindowsOptions, md *plugin.MergeDescription) bool {
	if k.AdminPassword != p.AdminPassword {
		p.AdminPassword = k.AdminPassword
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Clone_Customize_WindowsOptions_AutoLogon(k *WindowsOptions, p *WindowsOptions, md *plugin.MergeDescription) bool {
	if k.AutoLogon != p.AutoLogon {
		p.AutoLogon = k.AutoLogon
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Clone_Customize_WindowsOptions_AutoLogonCount(k *WindowsOptions, p *WindowsOptions, md *plugin.MergeDescription) bool {
	if k.AutoLogonCount != p.AutoLogonCount {
		p.AutoLogonCount = k.AutoLogonCount
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Clone_Customize_WindowsOptions_ComputerName(k *WindowsOptions, p *WindowsOptions, md *plugin.MergeDescription) bool {
	if k.ComputerName != p.ComputerName {
		p.ComputerName = k.ComputerName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeVirtualMachine_Clone_Customize_WindowsOptions_RunOnceCommandList(k *WindowsOptions, p *WindowsOptions, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.RunOnceCommandList, p.RunOnceCommandList) {
		p.RunOnceCommandList = k.RunOnceCommandList
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Clone_Customize_WindowsOptions_Workgroup(k *WindowsOptions, p *WindowsOptions, md *plugin.MergeDescription) bool {
	if k.Workgroup != p.Workgroup {
		p.Workgroup = k.Workgroup
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Clone_Customize_WindowsOptions_FullName(k *WindowsOptions, p *WindowsOptions, md *plugin.MergeDescription) bool {
	if k.FullName != p.FullName {
		p.FullName = k.FullName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Clone_Customize_WindowsOptions_OrganizationName(k *WindowsOptions, p *WindowsOptions, md *plugin.MergeDescription) bool {
	if k.OrganizationName != p.OrganizationName {
		p.OrganizationName = k.OrganizationName
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Clone_Customize_WindowsOptions_TimeZone(k *WindowsOptions, p *WindowsOptions, md *plugin.MergeDescription) bool {
	if k.TimeZone != p.TimeZone {
		p.TimeZone = k.TimeZone
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructSliceTemplateSpec
func MergeVirtualMachine_Disk(ksp *[]Disk, psp *[]Disk, md *plugin.MergeDescription) bool {
	ks := *ksp
	ps := *psp
	if len(ks) != len(ps) {
		ps = ks
		md.NeedsProviderUpdate = true
		return true
	}
	anyChildUpdated := false
	for i, _ := range ps {
		updated := false
		k := &ks[i]
		p := &ps[i]
		updated = MergeVirtualMachine_Disk_WriteThrough(k, p, md)
		if updated {
			anyChildUpdated = true
		}

		updated = MergeVirtualMachine_Disk_IoShareLevel(k, p, md)
		if updated {
			anyChildUpdated = true
		}

		updated = MergeVirtualMachine_Disk_Label(k, p, md)
		if updated {
			anyChildUpdated = true
		}

		updated = MergeVirtualMachine_Disk_Path(k, p, md)
		if updated {
			anyChildUpdated = true
		}

		updated = MergeVirtualMachine_Disk_Key(k, p, md)
		if updated {
			anyChildUpdated = true
		}

		updated = MergeVirtualMachine_Disk_Name(k, p, md)
		if updated {
			anyChildUpdated = true
		}

		updated = MergeVirtualMachine_Disk_ThinProvisioned(k, p, md)
		if updated {
			anyChildUpdated = true
		}

		updated = MergeVirtualMachine_Disk_Uuid(k, p, md)
		if updated {
			anyChildUpdated = true
		}

		updated = MergeVirtualMachine_Disk_DatastoreId(k, p, md)
		if updated {
			anyChildUpdated = true
		}

		updated = MergeVirtualMachine_Disk_DiskMode(k, p, md)
		if updated {
			anyChildUpdated = true
		}

		updated = MergeVirtualMachine_Disk_EagerlyScrub(k, p, md)
		if updated {
			anyChildUpdated = true
		}

		updated = MergeVirtualMachine_Disk_KeepOnRemove(k, p, md)
		if updated {
			anyChildUpdated = true
		}

		updated = MergeVirtualMachine_Disk_Attach(k, p, md)
		if updated {
			anyChildUpdated = true
		}

		updated = MergeVirtualMachine_Disk_ControllerType(k, p, md)
		if updated {
			anyChildUpdated = true
		}

		updated = MergeVirtualMachine_Disk_IoShareCount(k, p, md)
		if updated {
			anyChildUpdated = true
		}

		updated = MergeVirtualMachine_Disk_IoReservation(k, p, md)
		if updated {
			anyChildUpdated = true
		}

		updated = MergeVirtualMachine_Disk_Size(k, p, md)
		if updated {
			anyChildUpdated = true
		}

		updated = MergeVirtualMachine_Disk_StoragePolicyId(k, p, md)
		if updated {
			anyChildUpdated = true
		}

		updated = MergeVirtualMachine_Disk_UnitNumber(k, p, md)
		if updated {
			anyChildUpdated = true
		}

		updated = MergeVirtualMachine_Disk_DeviceAddress(k, p, md)
		if updated {
			anyChildUpdated = true
		}

		updated = MergeVirtualMachine_Disk_DiskSharing(k, p, md)
		if updated {
			anyChildUpdated = true
		}

		updated = MergeVirtualMachine_Disk_IoLimit(k, p, md)
		if updated {
			anyChildUpdated = true
		}

	}
	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Disk_WriteThrough(k *Disk, p *Disk, md *plugin.MergeDescription) bool {
	if k.WriteThrough != p.WriteThrough {
		p.WriteThrough = k.WriteThrough
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Disk_IoShareLevel(k *Disk, p *Disk, md *plugin.MergeDescription) bool {
	if k.IoShareLevel != p.IoShareLevel {
		p.IoShareLevel = k.IoShareLevel
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Disk_Label(k *Disk, p *Disk, md *plugin.MergeDescription) bool {
	if k.Label != p.Label {
		p.Label = k.Label
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Disk_Path(k *Disk, p *Disk, md *plugin.MergeDescription) bool {
	if k.Path != p.Path {
		p.Path = k.Path
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Disk_Key(k *Disk, p *Disk, md *plugin.MergeDescription) bool {
	if k.Key != p.Key {
		p.Key = k.Key
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Disk_Name(k *Disk, p *Disk, md *plugin.MergeDescription) bool {
	if k.Name != p.Name {
		p.Name = k.Name
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Disk_ThinProvisioned(k *Disk, p *Disk, md *plugin.MergeDescription) bool {
	if k.ThinProvisioned != p.ThinProvisioned {
		p.ThinProvisioned = k.ThinProvisioned
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Disk_Uuid(k *Disk, p *Disk, md *plugin.MergeDescription) bool {
	if k.Uuid != p.Uuid {
		p.Uuid = k.Uuid
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Disk_DatastoreId(k *Disk, p *Disk, md *plugin.MergeDescription) bool {
	if k.DatastoreId != p.DatastoreId {
		p.DatastoreId = k.DatastoreId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Disk_DiskMode(k *Disk, p *Disk, md *plugin.MergeDescription) bool {
	if k.DiskMode != p.DiskMode {
		p.DiskMode = k.DiskMode
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Disk_EagerlyScrub(k *Disk, p *Disk, md *plugin.MergeDescription) bool {
	if k.EagerlyScrub != p.EagerlyScrub {
		p.EagerlyScrub = k.EagerlyScrub
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Disk_KeepOnRemove(k *Disk, p *Disk, md *plugin.MergeDescription) bool {
	if k.KeepOnRemove != p.KeepOnRemove {
		p.KeepOnRemove = k.KeepOnRemove
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Disk_Attach(k *Disk, p *Disk, md *plugin.MergeDescription) bool {
	if k.Attach != p.Attach {
		p.Attach = k.Attach
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Disk_ControllerType(k *Disk, p *Disk, md *plugin.MergeDescription) bool {
	if k.ControllerType != p.ControllerType {
		p.ControllerType = k.ControllerType
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Disk_IoShareCount(k *Disk, p *Disk, md *plugin.MergeDescription) bool {
	if k.IoShareCount != p.IoShareCount {
		p.IoShareCount = k.IoShareCount
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Disk_IoReservation(k *Disk, p *Disk, md *plugin.MergeDescription) bool {
	if k.IoReservation != p.IoReservation {
		p.IoReservation = k.IoReservation
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Disk_Size(k *Disk, p *Disk, md *plugin.MergeDescription) bool {
	if k.Size != p.Size {
		p.Size = k.Size
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Disk_StoragePolicyId(k *Disk, p *Disk, md *plugin.MergeDescription) bool {
	if k.StoragePolicyId != p.StoragePolicyId {
		p.StoragePolicyId = k.StoragePolicyId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Disk_UnitNumber(k *Disk, p *Disk, md *plugin.MergeDescription) bool {
	if k.UnitNumber != p.UnitNumber {
		p.UnitNumber = k.UnitNumber
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Disk_DeviceAddress(k *Disk, p *Disk, md *plugin.MergeDescription) bool {
	if k.DeviceAddress != p.DeviceAddress {
		p.DeviceAddress = k.DeviceAddress
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Disk_DiskSharing(k *Disk, p *Disk, md *plugin.MergeDescription) bool {
	if k.DiskSharing != p.DiskSharing {
		p.DiskSharing = k.DiskSharing
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_Disk_IoLimit(k *Disk, p *Disk, md *plugin.MergeDescription) bool {
	if k.IoLimit != p.IoLimit {
		p.IoLimit = k.IoLimit
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructSliceTemplateSpec
func MergeVirtualMachine_NetworkInterface(ksp *[]NetworkInterface0, psp *[]NetworkInterface0, md *plugin.MergeDescription) bool {
	ks := *ksp
	ps := *psp
	if len(ks) != len(ps) {
		ps = ks
		md.NeedsProviderUpdate = true
		return true
	}
	anyChildUpdated := false
	for i, _ := range ps {
		updated := false
		k := &ks[i]
		p := &ps[i]
		updated = MergeVirtualMachine_NetworkInterface_Key(k, p, md)
		if updated {
			anyChildUpdated = true
		}

		updated = MergeVirtualMachine_NetworkInterface_AdapterType(k, p, md)
		if updated {
			anyChildUpdated = true
		}

		updated = MergeVirtualMachine_NetworkInterface_BandwidthLimit(k, p, md)
		if updated {
			anyChildUpdated = true
		}

		updated = MergeVirtualMachine_NetworkInterface_DeviceAddress(k, p, md)
		if updated {
			anyChildUpdated = true
		}

		updated = MergeVirtualMachine_NetworkInterface_MacAddress(k, p, md)
		if updated {
			anyChildUpdated = true
		}

		updated = MergeVirtualMachine_NetworkInterface_NetworkId(k, p, md)
		if updated {
			anyChildUpdated = true
		}

		updated = MergeVirtualMachine_NetworkInterface_OvfMapping(k, p, md)
		if updated {
			anyChildUpdated = true
		}

		updated = MergeVirtualMachine_NetworkInterface_UseStaticMac(k, p, md)
		if updated {
			anyChildUpdated = true
		}

		updated = MergeVirtualMachine_NetworkInterface_BandwidthReservation(k, p, md)
		if updated {
			anyChildUpdated = true
		}

		updated = MergeVirtualMachine_NetworkInterface_BandwidthShareCount(k, p, md)
		if updated {
			anyChildUpdated = true
		}

		updated = MergeVirtualMachine_NetworkInterface_BandwidthShareLevel(k, p, md)
		if updated {
			anyChildUpdated = true
		}

	}
	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_NetworkInterface_Key(k *NetworkInterface0, p *NetworkInterface0, md *plugin.MergeDescription) bool {
	if k.Key != p.Key {
		p.Key = k.Key
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_NetworkInterface_AdapterType(k *NetworkInterface0, p *NetworkInterface0, md *plugin.MergeDescription) bool {
	if k.AdapterType != p.AdapterType {
		p.AdapterType = k.AdapterType
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_NetworkInterface_BandwidthLimit(k *NetworkInterface0, p *NetworkInterface0, md *plugin.MergeDescription) bool {
	if k.BandwidthLimit != p.BandwidthLimit {
		p.BandwidthLimit = k.BandwidthLimit
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_NetworkInterface_DeviceAddress(k *NetworkInterface0, p *NetworkInterface0, md *plugin.MergeDescription) bool {
	if k.DeviceAddress != p.DeviceAddress {
		p.DeviceAddress = k.DeviceAddress
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_NetworkInterface_MacAddress(k *NetworkInterface0, p *NetworkInterface0, md *plugin.MergeDescription) bool {
	if k.MacAddress != p.MacAddress {
		p.MacAddress = k.MacAddress
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_NetworkInterface_NetworkId(k *NetworkInterface0, p *NetworkInterface0, md *plugin.MergeDescription) bool {
	if k.NetworkId != p.NetworkId {
		p.NetworkId = k.NetworkId
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_NetworkInterface_OvfMapping(k *NetworkInterface0, p *NetworkInterface0, md *plugin.MergeDescription) bool {
	if k.OvfMapping != p.OvfMapping {
		p.OvfMapping = k.OvfMapping
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_NetworkInterface_UseStaticMac(k *NetworkInterface0, p *NetworkInterface0, md *plugin.MergeDescription) bool {
	if k.UseStaticMac != p.UseStaticMac {
		p.UseStaticMac = k.UseStaticMac
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_NetworkInterface_BandwidthReservation(k *NetworkInterface0, p *NetworkInterface0, md *plugin.MergeDescription) bool {
	if k.BandwidthReservation != p.BandwidthReservation {
		p.BandwidthReservation = k.BandwidthReservation
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_NetworkInterface_BandwidthShareCount(k *NetworkInterface0, p *NetworkInterface0, md *plugin.MergeDescription) bool {
	if k.BandwidthShareCount != p.BandwidthShareCount {
		p.BandwidthShareCount = k.BandwidthShareCount
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_NetworkInterface_BandwidthShareLevel(k *NetworkInterface0, p *NetworkInterface0, md *plugin.MergeDescription) bool {
	if k.BandwidthShareLevel != p.BandwidthShareLevel {
		p.BandwidthShareLevel = k.BandwidthShareLevel
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructTemplateSpec
func MergeVirtualMachine_OvfDeploy(k *OvfDeploy, p *OvfDeploy, md *plugin.MergeDescription) bool {
	updated := false
	anyChildUpdated := false
	updated = MergeVirtualMachine_OvfDeploy_RemoteOvfUrl(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_OvfDeploy_AllowUnverifiedSslCert(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_OvfDeploy_DiskProvisioning(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_OvfDeploy_IpAllocationPolicy(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_OvfDeploy_IpProtocol(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_OvfDeploy_LocalOvfPath(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	updated = MergeVirtualMachine_OvfDeploy_OvfNetworkMap(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_OvfDeploy_RemoteOvfUrl(k *OvfDeploy, p *OvfDeploy, md *plugin.MergeDescription) bool {
	if k.RemoteOvfUrl != p.RemoteOvfUrl {
		p.RemoteOvfUrl = k.RemoteOvfUrl
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_OvfDeploy_AllowUnverifiedSslCert(k *OvfDeploy, p *OvfDeploy, md *plugin.MergeDescription) bool {
	if k.AllowUnverifiedSslCert != p.AllowUnverifiedSslCert {
		p.AllowUnverifiedSslCert = k.AllowUnverifiedSslCert
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_OvfDeploy_DiskProvisioning(k *OvfDeploy, p *OvfDeploy, md *plugin.MergeDescription) bool {
	if k.DiskProvisioning != p.DiskProvisioning {
		p.DiskProvisioning = k.DiskProvisioning
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_OvfDeploy_IpAllocationPolicy(k *OvfDeploy, p *OvfDeploy, md *plugin.MergeDescription) bool {
	if k.IpAllocationPolicy != p.IpAllocationPolicy {
		p.IpAllocationPolicy = k.IpAllocationPolicy
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_OvfDeploy_IpProtocol(k *OvfDeploy, p *OvfDeploy, md *plugin.MergeDescription) bool {
	if k.IpProtocol != p.IpProtocol {
		p.IpProtocol = k.IpProtocol
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateSpec
func MergeVirtualMachine_OvfDeploy_LocalOvfPath(k *OvfDeploy, p *OvfDeploy, md *plugin.MergeDescription) bool {
	if k.LocalOvfPath != p.LocalOvfPath {
		p.LocalOvfPath = k.LocalOvfPath
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateSpec
func MergeVirtualMachine_OvfDeploy_OvfNetworkMap(k *OvfDeploy, p *OvfDeploy, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.OvfNetworkMap, p.OvfNetworkMap) {
		p.OvfNetworkMap = k.OvfNetworkMap
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergeStructTemplateSpec
func MergeVirtualMachine_Vapp(k *Vapp, p *Vapp, md *plugin.MergeDescription) bool {
	updated := false
	anyChildUpdated := false
	updated = MergeVirtualMachine_Vapp_Properties(k, p, md)
	if updated {
		anyChildUpdated = true
	}

	if anyChildUpdated {
		md.NeedsProviderUpdate = true
	}
	return anyChildUpdated
}

//mergePrimitiveContainerTemplateSpec
func MergeVirtualMachine_Vapp_Properties(k *Vapp, p *Vapp, md *plugin.MergeDescription) bool {
	if !plugin.CompareMapString(k.Properties, p.Properties) {
		p.Properties = k.Properties
		md.NeedsProviderUpdate = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeVirtualMachine_ChangeVersion(k *VirtualMachineObservation, p *VirtualMachineObservation, md *plugin.MergeDescription) bool {
	if k.ChangeVersion != p.ChangeVersion {
		k.ChangeVersion = p.ChangeVersion
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeVirtualMachine_DefaultIpAddress(k *VirtualMachineObservation, p *VirtualMachineObservation, md *plugin.MergeDescription) bool {
	if k.DefaultIpAddress != p.DefaultIpAddress {
		k.DefaultIpAddress = p.DefaultIpAddress
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateStatus
func MergeVirtualMachine_GuestIpAddresses(k *VirtualMachineObservation, p *VirtualMachineObservation, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.GuestIpAddresses, p.GuestIpAddresses) {
		k.GuestIpAddresses = p.GuestIpAddresses
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeVirtualMachine_Imported(k *VirtualMachineObservation, p *VirtualMachineObservation, md *plugin.MergeDescription) bool {
	if k.Imported != p.Imported {
		k.Imported = p.Imported
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeVirtualMachine_Moid(k *VirtualMachineObservation, p *VirtualMachineObservation, md *plugin.MergeDescription) bool {
	if k.Moid != p.Moid {
		k.Moid = p.Moid
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeVirtualMachine_RebootRequired(k *VirtualMachineObservation, p *VirtualMachineObservation, md *plugin.MergeDescription) bool {
	if k.RebootRequired != p.RebootRequired {
		k.RebootRequired = p.RebootRequired
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeVirtualMachine_Uuid(k *VirtualMachineObservation, p *VirtualMachineObservation, md *plugin.MergeDescription) bool {
	if k.Uuid != p.Uuid {
		k.Uuid = p.Uuid
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveContainerTemplateStatus
func MergeVirtualMachine_VappTransport(k *VirtualMachineObservation, p *VirtualMachineObservation, md *plugin.MergeDescription) bool {
	if !plugin.CompareStringSlices(k.VappTransport, p.VappTransport) {
		k.VappTransport = p.VappTransport
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeVirtualMachine_VmwareToolsStatus(k *VirtualMachineObservation, p *VirtualMachineObservation, md *plugin.MergeDescription) bool {
	if k.VmwareToolsStatus != p.VmwareToolsStatus {
		k.VmwareToolsStatus = p.VmwareToolsStatus
		md.StatusUpdated = true
		return true
	}
	return false
}

//mergePrimitiveTemplateStatus
func MergeVirtualMachine_VmxPath(k *VirtualMachineObservation, p *VirtualMachineObservation, md *plugin.MergeDescription) bool {
	if k.VmxPath != p.VmxPath {
		k.VmxPath = p.VmxPath
		md.StatusUpdated = true
		return true
	}
	return false
}
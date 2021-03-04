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
	r, ok := mr.(*VirtualMachine)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeVirtualMachine(r, ctyValue)
}

func DecodeVirtualMachine(prev *VirtualMachine, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeVirtualMachine_AlternateGuestName(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_Annotation(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_BootDelay(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_BootRetryDelay(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_BootRetryEnabled(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_CpuHotAddEnabled(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_CpuHotRemoveEnabled(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_CpuLimit(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_CpuPerformanceCountersEnabled(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_CpuReservation(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_CpuShareCount(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_CpuShareLevel(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_CustomAttributes(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_DatacenterId(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_DatastoreClusterId(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_DatastoreId(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_EfiSecureBootEnabled(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_EnableDiskUuid(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_EnableLogging(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_EptRviMode(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_ExtraConfig(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_Firmware(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_Folder(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_ForcePowerOff(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_GuestId(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_HardwareVersion(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_HostSystemId(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_HvMode(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_IdeControllerCount(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_IgnoredGuestIps(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_LatencySensitivity(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_Memory(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_MemoryHotAddEnabled(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_MemoryLimit(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_MemoryReservation(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_MemoryShareCount(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_MemoryShareLevel(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_MigrateWaitTimeout(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_Name(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_NestedHvEnabled(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_NumCoresPerSocket(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_NumCpus(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_PciDeviceId(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_PoweronTimeout(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_ResourcePoolId(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_RunToolsScriptsAfterPowerOn(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_RunToolsScriptsAfterResume(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_RunToolsScriptsBeforeGuestReboot(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_RunToolsScriptsBeforeGuestShutdown(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_RunToolsScriptsBeforeGuestStandby(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_SataControllerCount(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_ScsiBusSharing(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_ScsiControllerCount(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_ScsiType(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_ShutdownWaitTimeout(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_StoragePolicyId(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_SwapPlacementPolicy(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_SyncTimeWithHost(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_Tags(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_WaitForGuestIpTimeout(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_WaitForGuestNetRoutable(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_WaitForGuestNetTimeout(&new.Spec.ForProvider, valMap)
	DecodeVirtualMachine_Cdrom(&new.Spec.ForProvider.Cdrom, valMap)
	DecodeVirtualMachine_Clone(&new.Spec.ForProvider.Clone, valMap)
	DecodeVirtualMachine_Disk(&new.Spec.ForProvider.Disk, valMap)
	DecodeVirtualMachine_NetworkInterface(&new.Spec.ForProvider.NetworkInterface, valMap)
	DecodeVirtualMachine_OvfDeploy(&new.Spec.ForProvider.OvfDeploy, valMap)
	DecodeVirtualMachine_Vapp(&new.Spec.ForProvider.Vapp, valMap)
	DecodeVirtualMachine_ChangeVersion(&new.Status.AtProvider, valMap)
	DecodeVirtualMachine_DefaultIpAddress(&new.Status.AtProvider, valMap)
	DecodeVirtualMachine_GuestIpAddresses(&new.Status.AtProvider, valMap)
	DecodeVirtualMachine_Imported(&new.Status.AtProvider, valMap)
	DecodeVirtualMachine_Moid(&new.Status.AtProvider, valMap)
	DecodeVirtualMachine_RebootRequired(&new.Status.AtProvider, valMap)
	DecodeVirtualMachine_Uuid(&new.Status.AtProvider, valMap)
	DecodeVirtualMachine_VappTransport(&new.Status.AtProvider, valMap)
	DecodeVirtualMachine_VmwareToolsStatus(&new.Status.AtProvider, valMap)
	DecodeVirtualMachine_VmxPath(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_AlternateGuestName(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.AlternateGuestName = ctwhy.ValueAsString(vals["alternate_guest_name"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Annotation(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.Annotation = ctwhy.ValueAsString(vals["annotation"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_BootDelay(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.BootDelay = ctwhy.ValueAsInt64(vals["boot_delay"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_BootRetryDelay(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.BootRetryDelay = ctwhy.ValueAsInt64(vals["boot_retry_delay"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_BootRetryEnabled(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.BootRetryEnabled = ctwhy.ValueAsBool(vals["boot_retry_enabled"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_CpuHotAddEnabled(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.CpuHotAddEnabled = ctwhy.ValueAsBool(vals["cpu_hot_add_enabled"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_CpuHotRemoveEnabled(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.CpuHotRemoveEnabled = ctwhy.ValueAsBool(vals["cpu_hot_remove_enabled"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_CpuLimit(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.CpuLimit = ctwhy.ValueAsInt64(vals["cpu_limit"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_CpuPerformanceCountersEnabled(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.CpuPerformanceCountersEnabled = ctwhy.ValueAsBool(vals["cpu_performance_counters_enabled"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_CpuReservation(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.CpuReservation = ctwhy.ValueAsInt64(vals["cpu_reservation"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_CpuShareCount(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.CpuShareCount = ctwhy.ValueAsInt64(vals["cpu_share_count"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_CpuShareLevel(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.CpuShareLevel = ctwhy.ValueAsString(vals["cpu_share_level"])
}

//primitiveMapTypeDecodeTemplate
func DecodeVirtualMachine_CustomAttributes(p *VirtualMachineParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	if vals["custom_attributes"].IsNull() {
		p.CustomAttributes = nil
		return
	}
	vMap := make(map[string]string)
	v := vals["custom_attributes"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.CustomAttributes = vMap
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_DatacenterId(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.DatacenterId = ctwhy.ValueAsString(vals["datacenter_id"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_DatastoreClusterId(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.DatastoreClusterId = ctwhy.ValueAsString(vals["datastore_cluster_id"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_DatastoreId(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.DatastoreId = ctwhy.ValueAsString(vals["datastore_id"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_EfiSecureBootEnabled(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.EfiSecureBootEnabled = ctwhy.ValueAsBool(vals["efi_secure_boot_enabled"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_EnableDiskUuid(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.EnableDiskUuid = ctwhy.ValueAsBool(vals["enable_disk_uuid"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_EnableLogging(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.EnableLogging = ctwhy.ValueAsBool(vals["enable_logging"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_EptRviMode(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.EptRviMode = ctwhy.ValueAsString(vals["ept_rvi_mode"])
}

//primitiveMapTypeDecodeTemplate
func DecodeVirtualMachine_ExtraConfig(p *VirtualMachineParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	if vals["extra_config"].IsNull() {
		p.ExtraConfig = nil
		return
	}
	vMap := make(map[string]string)
	v := vals["extra_config"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.ExtraConfig = vMap
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Firmware(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.Firmware = ctwhy.ValueAsString(vals["firmware"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Folder(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.Folder = ctwhy.ValueAsString(vals["folder"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_ForcePowerOff(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.ForcePowerOff = ctwhy.ValueAsBool(vals["force_power_off"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_GuestId(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.GuestId = ctwhy.ValueAsString(vals["guest_id"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_HardwareVersion(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.HardwareVersion = ctwhy.ValueAsInt64(vals["hardware_version"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_HostSystemId(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.HostSystemId = ctwhy.ValueAsString(vals["host_system_id"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_HvMode(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.HvMode = ctwhy.ValueAsString(vals["hv_mode"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_IdeControllerCount(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.IdeControllerCount = ctwhy.ValueAsInt64(vals["ide_controller_count"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeVirtualMachine_IgnoredGuestIps(p *VirtualMachineParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsList(vals["ignored_guest_ips"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.IgnoredGuestIps = goVals
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_LatencySensitivity(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.LatencySensitivity = ctwhy.ValueAsString(vals["latency_sensitivity"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Memory(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.Memory = ctwhy.ValueAsInt64(vals["memory"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_MemoryHotAddEnabled(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.MemoryHotAddEnabled = ctwhy.ValueAsBool(vals["memory_hot_add_enabled"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_MemoryLimit(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.MemoryLimit = ctwhy.ValueAsInt64(vals["memory_limit"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_MemoryReservation(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.MemoryReservation = ctwhy.ValueAsInt64(vals["memory_reservation"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_MemoryShareCount(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.MemoryShareCount = ctwhy.ValueAsInt64(vals["memory_share_count"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_MemoryShareLevel(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.MemoryShareLevel = ctwhy.ValueAsString(vals["memory_share_level"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_MigrateWaitTimeout(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.MigrateWaitTimeout = ctwhy.ValueAsInt64(vals["migrate_wait_timeout"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Name(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_NestedHvEnabled(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.NestedHvEnabled = ctwhy.ValueAsBool(vals["nested_hv_enabled"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_NumCoresPerSocket(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.NumCoresPerSocket = ctwhy.ValueAsInt64(vals["num_cores_per_socket"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_NumCpus(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.NumCpus = ctwhy.ValueAsInt64(vals["num_cpus"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeVirtualMachine_PciDeviceId(p *VirtualMachineParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["pci_device_id"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.PciDeviceId = goVals
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_PoweronTimeout(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.PoweronTimeout = ctwhy.ValueAsInt64(vals["poweron_timeout"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_ResourcePoolId(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.ResourcePoolId = ctwhy.ValueAsString(vals["resource_pool_id"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_RunToolsScriptsAfterPowerOn(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.RunToolsScriptsAfterPowerOn = ctwhy.ValueAsBool(vals["run_tools_scripts_after_power_on"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_RunToolsScriptsAfterResume(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.RunToolsScriptsAfterResume = ctwhy.ValueAsBool(vals["run_tools_scripts_after_resume"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_RunToolsScriptsBeforeGuestReboot(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.RunToolsScriptsBeforeGuestReboot = ctwhy.ValueAsBool(vals["run_tools_scripts_before_guest_reboot"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_RunToolsScriptsBeforeGuestShutdown(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.RunToolsScriptsBeforeGuestShutdown = ctwhy.ValueAsBool(vals["run_tools_scripts_before_guest_shutdown"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_RunToolsScriptsBeforeGuestStandby(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.RunToolsScriptsBeforeGuestStandby = ctwhy.ValueAsBool(vals["run_tools_scripts_before_guest_standby"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_SataControllerCount(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.SataControllerCount = ctwhy.ValueAsInt64(vals["sata_controller_count"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_ScsiBusSharing(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.ScsiBusSharing = ctwhy.ValueAsString(vals["scsi_bus_sharing"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_ScsiControllerCount(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.ScsiControllerCount = ctwhy.ValueAsInt64(vals["scsi_controller_count"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_ScsiType(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.ScsiType = ctwhy.ValueAsString(vals["scsi_type"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_ShutdownWaitTimeout(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.ShutdownWaitTimeout = ctwhy.ValueAsInt64(vals["shutdown_wait_timeout"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_StoragePolicyId(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.StoragePolicyId = ctwhy.ValueAsString(vals["storage_policy_id"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_SwapPlacementPolicy(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.SwapPlacementPolicy = ctwhy.ValueAsString(vals["swap_placement_policy"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_SyncTimeWithHost(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.SyncTimeWithHost = ctwhy.ValueAsBool(vals["sync_time_with_host"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeVirtualMachine_Tags(p *VirtualMachineParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["tags"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.Tags = goVals
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_WaitForGuestIpTimeout(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.WaitForGuestIpTimeout = ctwhy.ValueAsInt64(vals["wait_for_guest_ip_timeout"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_WaitForGuestNetRoutable(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.WaitForGuestNetRoutable = ctwhy.ValueAsBool(vals["wait_for_guest_net_routable"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_WaitForGuestNetTimeout(p *VirtualMachineParameters, vals map[string]cty.Value) {
	p.WaitForGuestNetTimeout = ctwhy.ValueAsInt64(vals["wait_for_guest_net_timeout"])
}

//containerCollectionSingletonTypeDecodeTemplate
func DecodeVirtualMachine_Cdrom(p *Cdrom, vals map[string]cty.Value) {
	if vals["cdrom"].IsNull() {
		*p = Cdrom{}
		return
	}
	rvals := ctwhy.ValueAsList(vals["cdrom"])
	if len(rvals) == 0 {
		*p = Cdrom{}
		return
	}
	// this template should be used when single dictionary/object values are nested in sets/lists
	// if rvals turns out to be a list with > 1 elements, something has broken with that heuristic
	valMap := rvals[0].AsValueMap()
	DecodeVirtualMachine_Cdrom_ClientDevice(p, valMap)
	DecodeVirtualMachine_Cdrom_DatastoreId(p, valMap)
	DecodeVirtualMachine_Cdrom_DeviceAddress(p, valMap)
	DecodeVirtualMachine_Cdrom_Key(p, valMap)
	DecodeVirtualMachine_Cdrom_Path(p, valMap)
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Cdrom_ClientDevice(p *Cdrom, vals map[string]cty.Value) {
	p.ClientDevice = ctwhy.ValueAsBool(vals["client_device"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Cdrom_DatastoreId(p *Cdrom, vals map[string]cty.Value) {
	p.DatastoreId = ctwhy.ValueAsString(vals["datastore_id"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Cdrom_DeviceAddress(p *Cdrom, vals map[string]cty.Value) {
	p.DeviceAddress = ctwhy.ValueAsString(vals["device_address"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Cdrom_Key(p *Cdrom, vals map[string]cty.Value) {
	p.Key = ctwhy.ValueAsInt64(vals["key"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Cdrom_Path(p *Cdrom, vals map[string]cty.Value) {
	p.Path = ctwhy.ValueAsString(vals["path"])
}

//containerCollectionSingletonTypeDecodeTemplate
func DecodeVirtualMachine_Clone(p *Clone, vals map[string]cty.Value) {
	if vals["clone"].IsNull() {
		*p = Clone{}
		return
	}
	rvals := ctwhy.ValueAsList(vals["clone"])
	if len(rvals) == 0 {
		*p = Clone{}
		return
	}
	// this template should be used when single dictionary/object values are nested in sets/lists
	// if rvals turns out to be a list with > 1 elements, something has broken with that heuristic
	valMap := rvals[0].AsValueMap()
	DecodeVirtualMachine_Clone_LinkedClone(p, valMap)
	DecodeVirtualMachine_Clone_OvfNetworkMap(p, valMap)
	DecodeVirtualMachine_Clone_OvfStorageMap(p, valMap)
	DecodeVirtualMachine_Clone_TemplateUuid(p, valMap)
	DecodeVirtualMachine_Clone_Timeout(p, valMap)
	DecodeVirtualMachine_Clone_Customize(&p.Customize, valMap)
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Clone_LinkedClone(p *Clone, vals map[string]cty.Value) {
	p.LinkedClone = ctwhy.ValueAsBool(vals["linked_clone"])
}

//primitiveMapTypeDecodeTemplate
func DecodeVirtualMachine_Clone_OvfNetworkMap(p *Clone, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	if vals["ovf_network_map"].IsNull() {
		p.OvfNetworkMap = nil
		return
	}
	vMap := make(map[string]string)
	v := vals["ovf_network_map"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.OvfNetworkMap = vMap
}

//primitiveMapTypeDecodeTemplate
func DecodeVirtualMachine_Clone_OvfStorageMap(p *Clone, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	if vals["ovf_storage_map"].IsNull() {
		p.OvfStorageMap = nil
		return
	}
	vMap := make(map[string]string)
	v := vals["ovf_storage_map"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.OvfStorageMap = vMap
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Clone_TemplateUuid(p *Clone, vals map[string]cty.Value) {
	p.TemplateUuid = ctwhy.ValueAsString(vals["template_uuid"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Clone_Timeout(p *Clone, vals map[string]cty.Value) {
	p.Timeout = ctwhy.ValueAsInt64(vals["timeout"])
}

//containerCollectionSingletonTypeDecodeTemplate
func DecodeVirtualMachine_Clone_Customize(p *Customize, vals map[string]cty.Value) {
	if vals["customize"].IsNull() {
		*p = Customize{}
		return
	}
	rvals := ctwhy.ValueAsList(vals["customize"])
	if len(rvals) == 0 {
		*p = Customize{}
		return
	}
	// this template should be used when single dictionary/object values are nested in sets/lists
	// if rvals turns out to be a list with > 1 elements, something has broken with that heuristic
	valMap := rvals[0].AsValueMap()
	DecodeVirtualMachine_Clone_Customize_DnsServerList(p, valMap)
	DecodeVirtualMachine_Clone_Customize_DnsSuffixList(p, valMap)
	DecodeVirtualMachine_Clone_Customize_Ipv4Gateway(p, valMap)
	DecodeVirtualMachine_Clone_Customize_Ipv6Gateway(p, valMap)
	DecodeVirtualMachine_Clone_Customize_Timeout(p, valMap)
	DecodeVirtualMachine_Clone_Customize_WindowsSysprepText(p, valMap)
	DecodeVirtualMachine_Clone_Customize_LinuxOptions(&p.LinuxOptions, valMap)
	DecodeVirtualMachine_Clone_Customize_NetworkInterface(&p.NetworkInterface, valMap)
	DecodeVirtualMachine_Clone_Customize_WindowsOptions(&p.WindowsOptions, valMap)
}

//primitiveCollectionTypeDecodeTemplate
func DecodeVirtualMachine_Clone_Customize_DnsServerList(p *Customize, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsList(vals["dns_server_list"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.DnsServerList = goVals
}

//primitiveCollectionTypeDecodeTemplate
func DecodeVirtualMachine_Clone_Customize_DnsSuffixList(p *Customize, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsList(vals["dns_suffix_list"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.DnsSuffixList = goVals
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Clone_Customize_Ipv4Gateway(p *Customize, vals map[string]cty.Value) {
	p.Ipv4Gateway = ctwhy.ValueAsString(vals["ipv4_gateway"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Clone_Customize_Ipv6Gateway(p *Customize, vals map[string]cty.Value) {
	p.Ipv6Gateway = ctwhy.ValueAsString(vals["ipv6_gateway"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Clone_Customize_Timeout(p *Customize, vals map[string]cty.Value) {
	p.Timeout = ctwhy.ValueAsInt64(vals["timeout"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Clone_Customize_WindowsSysprepText(p *Customize, vals map[string]cty.Value) {
	p.WindowsSysprepText = ctwhy.ValueAsString(vals["windows_sysprep_text"])
}

//containerCollectionSingletonTypeDecodeTemplate
func DecodeVirtualMachine_Clone_Customize_LinuxOptions(p *LinuxOptions, vals map[string]cty.Value) {
	if vals["linux_options"].IsNull() {
		*p = LinuxOptions{}
		return
	}
	rvals := ctwhy.ValueAsList(vals["linux_options"])
	if len(rvals) == 0 {
		*p = LinuxOptions{}
		return
	}
	// this template should be used when single dictionary/object values are nested in sets/lists
	// if rvals turns out to be a list with > 1 elements, something has broken with that heuristic
	valMap := rvals[0].AsValueMap()
	DecodeVirtualMachine_Clone_Customize_LinuxOptions_HwClockUtc(p, valMap)
	DecodeVirtualMachine_Clone_Customize_LinuxOptions_TimeZone(p, valMap)
	DecodeVirtualMachine_Clone_Customize_LinuxOptions_Domain(p, valMap)
	DecodeVirtualMachine_Clone_Customize_LinuxOptions_HostName(p, valMap)
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Clone_Customize_LinuxOptions_HwClockUtc(p *LinuxOptions, vals map[string]cty.Value) {
	p.HwClockUtc = ctwhy.ValueAsBool(vals["hw_clock_utc"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Clone_Customize_LinuxOptions_TimeZone(p *LinuxOptions, vals map[string]cty.Value) {
	p.TimeZone = ctwhy.ValueAsString(vals["time_zone"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Clone_Customize_LinuxOptions_Domain(p *LinuxOptions, vals map[string]cty.Value) {
	p.Domain = ctwhy.ValueAsString(vals["domain"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Clone_Customize_LinuxOptions_HostName(p *LinuxOptions, vals map[string]cty.Value) {
	p.HostName = ctwhy.ValueAsString(vals["host_name"])
}

//containerCollectionSingletonTypeDecodeTemplate
func DecodeVirtualMachine_Clone_Customize_NetworkInterface(p *NetworkInterface, vals map[string]cty.Value) {
	if vals["network_interface"].IsNull() {
		*p = NetworkInterface{}
		return
	}
	rvals := ctwhy.ValueAsList(vals["network_interface"])
	if len(rvals) == 0 {
		*p = NetworkInterface{}
		return
	}
	// this template should be used when single dictionary/object values are nested in sets/lists
	// if rvals turns out to be a list with > 1 elements, something has broken with that heuristic
	valMap := rvals[0].AsValueMap()
	DecodeVirtualMachine_Clone_Customize_NetworkInterface_DnsDomain(p, valMap)
	DecodeVirtualMachine_Clone_Customize_NetworkInterface_DnsServerList(p, valMap)
	DecodeVirtualMachine_Clone_Customize_NetworkInterface_Ipv4Address(p, valMap)
	DecodeVirtualMachine_Clone_Customize_NetworkInterface_Ipv4Netmask(p, valMap)
	DecodeVirtualMachine_Clone_Customize_NetworkInterface_Ipv6Address(p, valMap)
	DecodeVirtualMachine_Clone_Customize_NetworkInterface_Ipv6Netmask(p, valMap)
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Clone_Customize_NetworkInterface_DnsDomain(p *NetworkInterface, vals map[string]cty.Value) {
	p.DnsDomain = ctwhy.ValueAsString(vals["dns_domain"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeVirtualMachine_Clone_Customize_NetworkInterface_DnsServerList(p *NetworkInterface, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsList(vals["dns_server_list"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.DnsServerList = goVals
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Clone_Customize_NetworkInterface_Ipv4Address(p *NetworkInterface, vals map[string]cty.Value) {
	p.Ipv4Address = ctwhy.ValueAsString(vals["ipv4_address"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Clone_Customize_NetworkInterface_Ipv4Netmask(p *NetworkInterface, vals map[string]cty.Value) {
	p.Ipv4Netmask = ctwhy.ValueAsInt64(vals["ipv4_netmask"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Clone_Customize_NetworkInterface_Ipv6Address(p *NetworkInterface, vals map[string]cty.Value) {
	p.Ipv6Address = ctwhy.ValueAsString(vals["ipv6_address"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Clone_Customize_NetworkInterface_Ipv6Netmask(p *NetworkInterface, vals map[string]cty.Value) {
	p.Ipv6Netmask = ctwhy.ValueAsInt64(vals["ipv6_netmask"])
}

//containerCollectionSingletonTypeDecodeTemplate
func DecodeVirtualMachine_Clone_Customize_WindowsOptions(p *WindowsOptions, vals map[string]cty.Value) {
	if vals["windows_options"].IsNull() {
		*p = WindowsOptions{}
		return
	}
	rvals := ctwhy.ValueAsList(vals["windows_options"])
	if len(rvals) == 0 {
		*p = WindowsOptions{}
		return
	}
	// this template should be used when single dictionary/object values are nested in sets/lists
	// if rvals turns out to be a list with > 1 elements, something has broken with that heuristic
	valMap := rvals[0].AsValueMap()
	DecodeVirtualMachine_Clone_Customize_WindowsOptions_TimeZone(p, valMap)
	DecodeVirtualMachine_Clone_Customize_WindowsOptions_AutoLogonCount(p, valMap)
	DecodeVirtualMachine_Clone_Customize_WindowsOptions_ComputerName(p, valMap)
	DecodeVirtualMachine_Clone_Customize_WindowsOptions_FullName(p, valMap)
	DecodeVirtualMachine_Clone_Customize_WindowsOptions_JoinDomain(p, valMap)
	DecodeVirtualMachine_Clone_Customize_WindowsOptions_OrganizationName(p, valMap)
	DecodeVirtualMachine_Clone_Customize_WindowsOptions_RunOnceCommandList(p, valMap)
	DecodeVirtualMachine_Clone_Customize_WindowsOptions_AdminPassword(p, valMap)
	DecodeVirtualMachine_Clone_Customize_WindowsOptions_AutoLogon(p, valMap)
	DecodeVirtualMachine_Clone_Customize_WindowsOptions_DomainAdminPassword(p, valMap)
	DecodeVirtualMachine_Clone_Customize_WindowsOptions_DomainAdminUser(p, valMap)
	DecodeVirtualMachine_Clone_Customize_WindowsOptions_ProductKey(p, valMap)
	DecodeVirtualMachine_Clone_Customize_WindowsOptions_Workgroup(p, valMap)
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Clone_Customize_WindowsOptions_TimeZone(p *WindowsOptions, vals map[string]cty.Value) {
	p.TimeZone = ctwhy.ValueAsInt64(vals["time_zone"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Clone_Customize_WindowsOptions_AutoLogonCount(p *WindowsOptions, vals map[string]cty.Value) {
	p.AutoLogonCount = ctwhy.ValueAsInt64(vals["auto_logon_count"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Clone_Customize_WindowsOptions_ComputerName(p *WindowsOptions, vals map[string]cty.Value) {
	p.ComputerName = ctwhy.ValueAsString(vals["computer_name"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Clone_Customize_WindowsOptions_FullName(p *WindowsOptions, vals map[string]cty.Value) {
	p.FullName = ctwhy.ValueAsString(vals["full_name"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Clone_Customize_WindowsOptions_JoinDomain(p *WindowsOptions, vals map[string]cty.Value) {
	p.JoinDomain = ctwhy.ValueAsString(vals["join_domain"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Clone_Customize_WindowsOptions_OrganizationName(p *WindowsOptions, vals map[string]cty.Value) {
	p.OrganizationName = ctwhy.ValueAsString(vals["organization_name"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeVirtualMachine_Clone_Customize_WindowsOptions_RunOnceCommandList(p *WindowsOptions, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsList(vals["run_once_command_list"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.RunOnceCommandList = goVals
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Clone_Customize_WindowsOptions_AdminPassword(p *WindowsOptions, vals map[string]cty.Value) {
	p.AdminPassword = ctwhy.ValueAsString(vals["admin_password"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Clone_Customize_WindowsOptions_AutoLogon(p *WindowsOptions, vals map[string]cty.Value) {
	p.AutoLogon = ctwhy.ValueAsBool(vals["auto_logon"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Clone_Customize_WindowsOptions_DomainAdminPassword(p *WindowsOptions, vals map[string]cty.Value) {
	p.DomainAdminPassword = ctwhy.ValueAsString(vals["domain_admin_password"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Clone_Customize_WindowsOptions_DomainAdminUser(p *WindowsOptions, vals map[string]cty.Value) {
	p.DomainAdminUser = ctwhy.ValueAsString(vals["domain_admin_user"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Clone_Customize_WindowsOptions_ProductKey(p *WindowsOptions, vals map[string]cty.Value) {
	p.ProductKey = ctwhy.ValueAsString(vals["product_key"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Clone_Customize_WindowsOptions_Workgroup(p *WindowsOptions, vals map[string]cty.Value) {
	p.Workgroup = ctwhy.ValueAsString(vals["workgroup"])
}

//containerCollectionTypeDecodeTemplate
func DecodeVirtualMachine_Disk(pp *[]Disk, vals map[string]cty.Value) {
	if vals["disk"].IsNull() {
		*pp = []Disk{}
		return
	}
	rvals := ctwhy.ValueAsList(vals["disk"])
	if len(rvals) == 0 {
		*pp = []Disk{}
		return
	}
	lval := make([]Disk, 0)
	for _, value := range rvals {
		valMap := value.AsValueMap()
		vi := &Disk{}
		DecodeVirtualMachine_Disk_WriteThrough(vi, valMap)
		DecodeVirtualMachine_Disk_IoLimit(vi, valMap)
		DecodeVirtualMachine_Disk_StoragePolicyId(vi, valMap)
		DecodeVirtualMachine_Disk_IoShareLevel(vi, valMap)
		DecodeVirtualMachine_Disk_KeepOnRemove(vi, valMap)
		DecodeVirtualMachine_Disk_Label(vi, valMap)
		DecodeVirtualMachine_Disk_Name(vi, valMap)
		DecodeVirtualMachine_Disk_Size(vi, valMap)
		DecodeVirtualMachine_Disk_Uuid(vi, valMap)
		DecodeVirtualMachine_Disk_ControllerType(vi, valMap)
		DecodeVirtualMachine_Disk_DiskMode(vi, valMap)
		DecodeVirtualMachine_Disk_EagerlyScrub(vi, valMap)
		DecodeVirtualMachine_Disk_IoReservation(vi, valMap)
		DecodeVirtualMachine_Disk_IoShareCount(vi, valMap)
		DecodeVirtualMachine_Disk_Key(vi, valMap)
		DecodeVirtualMachine_Disk_UnitNumber(vi, valMap)
		DecodeVirtualMachine_Disk_DeviceAddress(vi, valMap)
		DecodeVirtualMachine_Disk_DiskSharing(vi, valMap)
		DecodeVirtualMachine_Disk_Path(vi, valMap)
		DecodeVirtualMachine_Disk_ThinProvisioned(vi, valMap)
		DecodeVirtualMachine_Disk_Attach(vi, valMap)
		DecodeVirtualMachine_Disk_DatastoreId(vi, valMap)
		lval = append(lval, *vi)
	}
	*pp = lval
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Disk_WriteThrough(p *Disk, vals map[string]cty.Value) {
	p.WriteThrough = ctwhy.ValueAsBool(vals["write_through"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Disk_IoLimit(p *Disk, vals map[string]cty.Value) {
	p.IoLimit = ctwhy.ValueAsInt64(vals["io_limit"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Disk_StoragePolicyId(p *Disk, vals map[string]cty.Value) {
	p.StoragePolicyId = ctwhy.ValueAsString(vals["storage_policy_id"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Disk_IoShareLevel(p *Disk, vals map[string]cty.Value) {
	p.IoShareLevel = ctwhy.ValueAsString(vals["io_share_level"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Disk_KeepOnRemove(p *Disk, vals map[string]cty.Value) {
	p.KeepOnRemove = ctwhy.ValueAsBool(vals["keep_on_remove"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Disk_Label(p *Disk, vals map[string]cty.Value) {
	p.Label = ctwhy.ValueAsString(vals["label"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Disk_Name(p *Disk, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Disk_Size(p *Disk, vals map[string]cty.Value) {
	p.Size = ctwhy.ValueAsInt64(vals["size"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Disk_Uuid(p *Disk, vals map[string]cty.Value) {
	p.Uuid = ctwhy.ValueAsString(vals["uuid"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Disk_ControllerType(p *Disk, vals map[string]cty.Value) {
	p.ControllerType = ctwhy.ValueAsString(vals["controller_type"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Disk_DiskMode(p *Disk, vals map[string]cty.Value) {
	p.DiskMode = ctwhy.ValueAsString(vals["disk_mode"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Disk_EagerlyScrub(p *Disk, vals map[string]cty.Value) {
	p.EagerlyScrub = ctwhy.ValueAsBool(vals["eagerly_scrub"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Disk_IoReservation(p *Disk, vals map[string]cty.Value) {
	p.IoReservation = ctwhy.ValueAsInt64(vals["io_reservation"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Disk_IoShareCount(p *Disk, vals map[string]cty.Value) {
	p.IoShareCount = ctwhy.ValueAsInt64(vals["io_share_count"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Disk_Key(p *Disk, vals map[string]cty.Value) {
	p.Key = ctwhy.ValueAsInt64(vals["key"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Disk_UnitNumber(p *Disk, vals map[string]cty.Value) {
	p.UnitNumber = ctwhy.ValueAsInt64(vals["unit_number"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Disk_DeviceAddress(p *Disk, vals map[string]cty.Value) {
	p.DeviceAddress = ctwhy.ValueAsString(vals["device_address"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Disk_DiskSharing(p *Disk, vals map[string]cty.Value) {
	p.DiskSharing = ctwhy.ValueAsString(vals["disk_sharing"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Disk_Path(p *Disk, vals map[string]cty.Value) {
	p.Path = ctwhy.ValueAsString(vals["path"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Disk_ThinProvisioned(p *Disk, vals map[string]cty.Value) {
	p.ThinProvisioned = ctwhy.ValueAsBool(vals["thin_provisioned"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Disk_Attach(p *Disk, vals map[string]cty.Value) {
	p.Attach = ctwhy.ValueAsBool(vals["attach"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Disk_DatastoreId(p *Disk, vals map[string]cty.Value) {
	p.DatastoreId = ctwhy.ValueAsString(vals["datastore_id"])
}

//containerCollectionTypeDecodeTemplate
func DecodeVirtualMachine_NetworkInterface(pp *[]NetworkInterface0, vals map[string]cty.Value) {
	if vals["network_interface"].IsNull() {
		*pp = []NetworkInterface0{}
		return
	}
	rvals := ctwhy.ValueAsList(vals["network_interface"])
	if len(rvals) == 0 {
		*pp = []NetworkInterface0{}
		return
	}
	lval := make([]NetworkInterface0, 0)
	for _, value := range rvals {
		valMap := value.AsValueMap()
		vi := &NetworkInterface0{}
		DecodeVirtualMachine_NetworkInterface_BandwidthShareLevel(vi, valMap)
		DecodeVirtualMachine_NetworkInterface_DeviceAddress(vi, valMap)
		DecodeVirtualMachine_NetworkInterface_Key(vi, valMap)
		DecodeVirtualMachine_NetworkInterface_NetworkId(vi, valMap)
		DecodeVirtualMachine_NetworkInterface_OvfMapping(vi, valMap)
		DecodeVirtualMachine_NetworkInterface_UseStaticMac(vi, valMap)
		DecodeVirtualMachine_NetworkInterface_BandwidthReservation(vi, valMap)
		DecodeVirtualMachine_NetworkInterface_BandwidthShareCount(vi, valMap)
		DecodeVirtualMachine_NetworkInterface_MacAddress(vi, valMap)
		DecodeVirtualMachine_NetworkInterface_AdapterType(vi, valMap)
		DecodeVirtualMachine_NetworkInterface_BandwidthLimit(vi, valMap)
		lval = append(lval, *vi)
	}
	*pp = lval
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_NetworkInterface_BandwidthShareLevel(p *NetworkInterface0, vals map[string]cty.Value) {
	p.BandwidthShareLevel = ctwhy.ValueAsString(vals["bandwidth_share_level"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_NetworkInterface_DeviceAddress(p *NetworkInterface0, vals map[string]cty.Value) {
	p.DeviceAddress = ctwhy.ValueAsString(vals["device_address"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_NetworkInterface_Key(p *NetworkInterface0, vals map[string]cty.Value) {
	p.Key = ctwhy.ValueAsInt64(vals["key"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_NetworkInterface_NetworkId(p *NetworkInterface0, vals map[string]cty.Value) {
	p.NetworkId = ctwhy.ValueAsString(vals["network_id"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_NetworkInterface_OvfMapping(p *NetworkInterface0, vals map[string]cty.Value) {
	p.OvfMapping = ctwhy.ValueAsString(vals["ovf_mapping"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_NetworkInterface_UseStaticMac(p *NetworkInterface0, vals map[string]cty.Value) {
	p.UseStaticMac = ctwhy.ValueAsBool(vals["use_static_mac"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_NetworkInterface_BandwidthReservation(p *NetworkInterface0, vals map[string]cty.Value) {
	p.BandwidthReservation = ctwhy.ValueAsInt64(vals["bandwidth_reservation"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_NetworkInterface_BandwidthShareCount(p *NetworkInterface0, vals map[string]cty.Value) {
	p.BandwidthShareCount = ctwhy.ValueAsInt64(vals["bandwidth_share_count"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_NetworkInterface_MacAddress(p *NetworkInterface0, vals map[string]cty.Value) {
	p.MacAddress = ctwhy.ValueAsString(vals["mac_address"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_NetworkInterface_AdapterType(p *NetworkInterface0, vals map[string]cty.Value) {
	p.AdapterType = ctwhy.ValueAsString(vals["adapter_type"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_NetworkInterface_BandwidthLimit(p *NetworkInterface0, vals map[string]cty.Value) {
	p.BandwidthLimit = ctwhy.ValueAsInt64(vals["bandwidth_limit"])
}

//containerCollectionSingletonTypeDecodeTemplate
func DecodeVirtualMachine_OvfDeploy(p *OvfDeploy, vals map[string]cty.Value) {
	if vals["ovf_deploy"].IsNull() {
		*p = OvfDeploy{}
		return
	}
	rvals := ctwhy.ValueAsList(vals["ovf_deploy"])
	if len(rvals) == 0 {
		*p = OvfDeploy{}
		return
	}
	// this template should be used when single dictionary/object values are nested in sets/lists
	// if rvals turns out to be a list with > 1 elements, something has broken with that heuristic
	valMap := rvals[0].AsValueMap()
	DecodeVirtualMachine_OvfDeploy_LocalOvfPath(p, valMap)
	DecodeVirtualMachine_OvfDeploy_OvfNetworkMap(p, valMap)
	DecodeVirtualMachine_OvfDeploy_RemoteOvfUrl(p, valMap)
	DecodeVirtualMachine_OvfDeploy_AllowUnverifiedSslCert(p, valMap)
	DecodeVirtualMachine_OvfDeploy_DiskProvisioning(p, valMap)
	DecodeVirtualMachine_OvfDeploy_IpAllocationPolicy(p, valMap)
	DecodeVirtualMachine_OvfDeploy_IpProtocol(p, valMap)
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_OvfDeploy_LocalOvfPath(p *OvfDeploy, vals map[string]cty.Value) {
	p.LocalOvfPath = ctwhy.ValueAsString(vals["local_ovf_path"])
}

//primitiveMapTypeDecodeTemplate
func DecodeVirtualMachine_OvfDeploy_OvfNetworkMap(p *OvfDeploy, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	if vals["ovf_network_map"].IsNull() {
		p.OvfNetworkMap = nil
		return
	}
	vMap := make(map[string]string)
	v := vals["ovf_network_map"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.OvfNetworkMap = vMap
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_OvfDeploy_RemoteOvfUrl(p *OvfDeploy, vals map[string]cty.Value) {
	p.RemoteOvfUrl = ctwhy.ValueAsString(vals["remote_ovf_url"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_OvfDeploy_AllowUnverifiedSslCert(p *OvfDeploy, vals map[string]cty.Value) {
	p.AllowUnverifiedSslCert = ctwhy.ValueAsBool(vals["allow_unverified_ssl_cert"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_OvfDeploy_DiskProvisioning(p *OvfDeploy, vals map[string]cty.Value) {
	p.DiskProvisioning = ctwhy.ValueAsString(vals["disk_provisioning"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_OvfDeploy_IpAllocationPolicy(p *OvfDeploy, vals map[string]cty.Value) {
	p.IpAllocationPolicy = ctwhy.ValueAsString(vals["ip_allocation_policy"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_OvfDeploy_IpProtocol(p *OvfDeploy, vals map[string]cty.Value) {
	p.IpProtocol = ctwhy.ValueAsString(vals["ip_protocol"])
}

//containerCollectionSingletonTypeDecodeTemplate
func DecodeVirtualMachine_Vapp(p *Vapp, vals map[string]cty.Value) {
	if vals["vapp"].IsNull() {
		*p = Vapp{}
		return
	}
	rvals := ctwhy.ValueAsList(vals["vapp"])
	if len(rvals) == 0 {
		*p = Vapp{}
		return
	}
	// this template should be used when single dictionary/object values are nested in sets/lists
	// if rvals turns out to be a list with > 1 elements, something has broken with that heuristic
	valMap := rvals[0].AsValueMap()
	DecodeVirtualMachine_Vapp_Properties(p, valMap)
}

//primitiveMapTypeDecodeTemplate
func DecodeVirtualMachine_Vapp_Properties(p *Vapp, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	if vals["properties"].IsNull() {
		p.Properties = nil
		return
	}
	vMap := make(map[string]string)
	v := vals["properties"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Properties = vMap
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_ChangeVersion(p *VirtualMachineObservation, vals map[string]cty.Value) {
	p.ChangeVersion = ctwhy.ValueAsString(vals["change_version"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_DefaultIpAddress(p *VirtualMachineObservation, vals map[string]cty.Value) {
	p.DefaultIpAddress = ctwhy.ValueAsString(vals["default_ip_address"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeVirtualMachine_GuestIpAddresses(p *VirtualMachineObservation, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsList(vals["guest_ip_addresses"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.GuestIpAddresses = goVals
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Imported(p *VirtualMachineObservation, vals map[string]cty.Value) {
	p.Imported = ctwhy.ValueAsBool(vals["imported"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Moid(p *VirtualMachineObservation, vals map[string]cty.Value) {
	p.Moid = ctwhy.ValueAsString(vals["moid"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_RebootRequired(p *VirtualMachineObservation, vals map[string]cty.Value) {
	p.RebootRequired = ctwhy.ValueAsBool(vals["reboot_required"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_Uuid(p *VirtualMachineObservation, vals map[string]cty.Value) {
	p.Uuid = ctwhy.ValueAsString(vals["uuid"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeVirtualMachine_VappTransport(p *VirtualMachineObservation, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsList(vals["vapp_transport"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.VappTransport = goVals
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_VmwareToolsStatus(p *VirtualMachineObservation, vals map[string]cty.Value) {
	p.VmwareToolsStatus = ctwhy.ValueAsString(vals["vmware_tools_status"])
}

//primitiveTypeDecodeTemplate
func DecodeVirtualMachine_VmxPath(p *VirtualMachineObservation, vals map[string]cty.Value) {
	p.VmxPath = ctwhy.ValueAsString(vals["vmx_path"])
}

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
	r, ok := mr.(*VirtualMachine)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a VirtualMachine.")
	}
	return EncodeVirtualMachine(*r), nil
}

func EncodeVirtualMachine(r VirtualMachine) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeVirtualMachine_AlternateGuestName(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_Annotation(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_BootDelay(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_BootRetryDelay(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_BootRetryEnabled(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_CpuHotAddEnabled(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_CpuHotRemoveEnabled(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_CpuLimit(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_CpuPerformanceCountersEnabled(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_CpuReservation(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_CpuShareCount(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_CpuShareLevel(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_CustomAttributes(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_DatacenterId(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_DatastoreClusterId(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_DatastoreId(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_EfiSecureBootEnabled(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_EnableDiskUuid(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_EnableLogging(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_EptRviMode(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_ExtraConfig(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_Firmware(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_Folder(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_ForcePowerOff(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_GuestId(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_HardwareVersion(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_HostSystemId(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_HvMode(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_IdeControllerCount(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_IgnoredGuestIps(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_LatencySensitivity(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_Memory(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_MemoryHotAddEnabled(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_MemoryLimit(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_MemoryReservation(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_MemoryShareCount(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_MemoryShareLevel(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_MigrateWaitTimeout(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_Name(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_NestedHvEnabled(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_NumCoresPerSocket(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_NumCpus(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_PciDeviceId(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_PoweronTimeout(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_ResourcePoolId(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_RunToolsScriptsAfterPowerOn(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_RunToolsScriptsAfterResume(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_RunToolsScriptsBeforeGuestReboot(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_RunToolsScriptsBeforeGuestShutdown(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_RunToolsScriptsBeforeGuestStandby(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_SataControllerCount(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_ScsiBusSharing(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_ScsiControllerCount(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_ScsiType(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_ShutdownWaitTimeout(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_StoragePolicyId(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_SwapPlacementPolicy(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_SyncTimeWithHost(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_Tags(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_WaitForGuestIpTimeout(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_WaitForGuestNetRoutable(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_WaitForGuestNetTimeout(r.Spec.ForProvider, ctyVal)
	EncodeVirtualMachine_Cdrom(r.Spec.ForProvider.Cdrom, ctyVal)
	EncodeVirtualMachine_Clone(r.Spec.ForProvider.Clone, ctyVal)
	EncodeVirtualMachine_Disk(r.Spec.ForProvider.Disk, ctyVal)
	EncodeVirtualMachine_NetworkInterface(r.Spec.ForProvider.NetworkInterface, ctyVal)
	EncodeVirtualMachine_OvfDeploy(r.Spec.ForProvider.OvfDeploy, ctyVal)
	EncodeVirtualMachine_Vapp(r.Spec.ForProvider.Vapp, ctyVal)
	EncodeVirtualMachine_ChangeVersion(r.Status.AtProvider, ctyVal)
	EncodeVirtualMachine_DefaultIpAddress(r.Status.AtProvider, ctyVal)
	EncodeVirtualMachine_GuestIpAddresses(r.Status.AtProvider, ctyVal)
	EncodeVirtualMachine_Imported(r.Status.AtProvider, ctyVal)
	EncodeVirtualMachine_Moid(r.Status.AtProvider, ctyVal)
	EncodeVirtualMachine_RebootRequired(r.Status.AtProvider, ctyVal)
	EncodeVirtualMachine_Uuid(r.Status.AtProvider, ctyVal)
	EncodeVirtualMachine_VappTransport(r.Status.AtProvider, ctyVal)
	EncodeVirtualMachine_VmwareToolsStatus(r.Status.AtProvider, ctyVal)
	EncodeVirtualMachine_VmxPath(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeVirtualMachine_AlternateGuestName(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["alternate_guest_name"] = cty.StringVal(p.AlternateGuestName)
}

func EncodeVirtualMachine_Annotation(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["annotation"] = cty.StringVal(p.Annotation)
}

func EncodeVirtualMachine_BootDelay(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["boot_delay"] = cty.NumberIntVal(p.BootDelay)
}

func EncodeVirtualMachine_BootRetryDelay(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["boot_retry_delay"] = cty.NumberIntVal(p.BootRetryDelay)
}

func EncodeVirtualMachine_BootRetryEnabled(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["boot_retry_enabled"] = cty.BoolVal(p.BootRetryEnabled)
}

func EncodeVirtualMachine_CpuHotAddEnabled(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["cpu_hot_add_enabled"] = cty.BoolVal(p.CpuHotAddEnabled)
}

func EncodeVirtualMachine_CpuHotRemoveEnabled(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["cpu_hot_remove_enabled"] = cty.BoolVal(p.CpuHotRemoveEnabled)
}

func EncodeVirtualMachine_CpuLimit(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["cpu_limit"] = cty.NumberIntVal(p.CpuLimit)
}

func EncodeVirtualMachine_CpuPerformanceCountersEnabled(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["cpu_performance_counters_enabled"] = cty.BoolVal(p.CpuPerformanceCountersEnabled)
}

func EncodeVirtualMachine_CpuReservation(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["cpu_reservation"] = cty.NumberIntVal(p.CpuReservation)
}

func EncodeVirtualMachine_CpuShareCount(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["cpu_share_count"] = cty.NumberIntVal(p.CpuShareCount)
}

func EncodeVirtualMachine_CpuShareLevel(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["cpu_share_level"] = cty.StringVal(p.CpuShareLevel)
}

func EncodeVirtualMachine_CustomAttributes(p VirtualMachineParameters, vals map[string]cty.Value) {
	if len(p.CustomAttributes) == 0 {
		vals["custom_attributes"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.CustomAttributes {
		mVals[key] = cty.StringVal(value)
	}
	vals["custom_attributes"] = cty.MapVal(mVals)
}

func EncodeVirtualMachine_DatacenterId(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["datacenter_id"] = cty.StringVal(p.DatacenterId)
}

func EncodeVirtualMachine_DatastoreClusterId(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["datastore_cluster_id"] = cty.StringVal(p.DatastoreClusterId)
}

func EncodeVirtualMachine_DatastoreId(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["datastore_id"] = cty.StringVal(p.DatastoreId)
}

func EncodeVirtualMachine_EfiSecureBootEnabled(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["efi_secure_boot_enabled"] = cty.BoolVal(p.EfiSecureBootEnabled)
}

func EncodeVirtualMachine_EnableDiskUuid(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["enable_disk_uuid"] = cty.BoolVal(p.EnableDiskUuid)
}

func EncodeVirtualMachine_EnableLogging(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["enable_logging"] = cty.BoolVal(p.EnableLogging)
}

func EncodeVirtualMachine_EptRviMode(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["ept_rvi_mode"] = cty.StringVal(p.EptRviMode)
}

func EncodeVirtualMachine_ExtraConfig(p VirtualMachineParameters, vals map[string]cty.Value) {
	if len(p.ExtraConfig) == 0 {
		vals["extra_config"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.ExtraConfig {
		mVals[key] = cty.StringVal(value)
	}
	vals["extra_config"] = cty.MapVal(mVals)
}

func EncodeVirtualMachine_Firmware(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["firmware"] = cty.StringVal(p.Firmware)
}

func EncodeVirtualMachine_Folder(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["folder"] = cty.StringVal(p.Folder)
}

func EncodeVirtualMachine_ForcePowerOff(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["force_power_off"] = cty.BoolVal(p.ForcePowerOff)
}

func EncodeVirtualMachine_GuestId(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["guest_id"] = cty.StringVal(p.GuestId)
}

func EncodeVirtualMachine_HardwareVersion(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["hardware_version"] = cty.NumberIntVal(p.HardwareVersion)
}

func EncodeVirtualMachine_HostSystemId(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["host_system_id"] = cty.StringVal(p.HostSystemId)
}

func EncodeVirtualMachine_HvMode(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["hv_mode"] = cty.StringVal(p.HvMode)
}

func EncodeVirtualMachine_IdeControllerCount(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["ide_controller_count"] = cty.NumberIntVal(p.IdeControllerCount)
}

func EncodeVirtualMachine_IgnoredGuestIps(p VirtualMachineParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.IgnoredGuestIps {
		colVals = append(colVals, cty.StringVal(value))
	}
	if len(colVals) == 0 {
		vals["ignored_guest_ips"] = cty.ListValEmpty(cty.String)
	} else {
		vals["ignored_guest_ips"] = cty.ListVal(colVals)
	}
}

func EncodeVirtualMachine_LatencySensitivity(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["latency_sensitivity"] = cty.StringVal(p.LatencySensitivity)
}

func EncodeVirtualMachine_Memory(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["memory"] = cty.NumberIntVal(p.Memory)
}

func EncodeVirtualMachine_MemoryHotAddEnabled(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["memory_hot_add_enabled"] = cty.BoolVal(p.MemoryHotAddEnabled)
}

func EncodeVirtualMachine_MemoryLimit(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["memory_limit"] = cty.NumberIntVal(p.MemoryLimit)
}

func EncodeVirtualMachine_MemoryReservation(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["memory_reservation"] = cty.NumberIntVal(p.MemoryReservation)
}

func EncodeVirtualMachine_MemoryShareCount(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["memory_share_count"] = cty.NumberIntVal(p.MemoryShareCount)
}

func EncodeVirtualMachine_MemoryShareLevel(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["memory_share_level"] = cty.StringVal(p.MemoryShareLevel)
}

func EncodeVirtualMachine_MigrateWaitTimeout(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["migrate_wait_timeout"] = cty.NumberIntVal(p.MigrateWaitTimeout)
}

func EncodeVirtualMachine_Name(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeVirtualMachine_NestedHvEnabled(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["nested_hv_enabled"] = cty.BoolVal(p.NestedHvEnabled)
}

func EncodeVirtualMachine_NumCoresPerSocket(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["num_cores_per_socket"] = cty.NumberIntVal(p.NumCoresPerSocket)
}

func EncodeVirtualMachine_NumCpus(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["num_cpus"] = cty.NumberIntVal(p.NumCpus)
}

func EncodeVirtualMachine_PciDeviceId(p VirtualMachineParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.PciDeviceId {
		colVals = append(colVals, cty.StringVal(value))
	}
	if len(colVals) == 0 {
		vals["pci_device_id"] = cty.SetValEmpty(cty.String)
	} else {
		vals["pci_device_id"] = cty.SetVal(colVals)
	}
}

func EncodeVirtualMachine_PoweronTimeout(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["poweron_timeout"] = cty.NumberIntVal(p.PoweronTimeout)
}

func EncodeVirtualMachine_ResourcePoolId(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["resource_pool_id"] = cty.StringVal(p.ResourcePoolId)
}

func EncodeVirtualMachine_RunToolsScriptsAfterPowerOn(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["run_tools_scripts_after_power_on"] = cty.BoolVal(p.RunToolsScriptsAfterPowerOn)
}

func EncodeVirtualMachine_RunToolsScriptsAfterResume(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["run_tools_scripts_after_resume"] = cty.BoolVal(p.RunToolsScriptsAfterResume)
}

func EncodeVirtualMachine_RunToolsScriptsBeforeGuestReboot(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["run_tools_scripts_before_guest_reboot"] = cty.BoolVal(p.RunToolsScriptsBeforeGuestReboot)
}

func EncodeVirtualMachine_RunToolsScriptsBeforeGuestShutdown(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["run_tools_scripts_before_guest_shutdown"] = cty.BoolVal(p.RunToolsScriptsBeforeGuestShutdown)
}

func EncodeVirtualMachine_RunToolsScriptsBeforeGuestStandby(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["run_tools_scripts_before_guest_standby"] = cty.BoolVal(p.RunToolsScriptsBeforeGuestStandby)
}

func EncodeVirtualMachine_SataControllerCount(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["sata_controller_count"] = cty.NumberIntVal(p.SataControllerCount)
}

func EncodeVirtualMachine_ScsiBusSharing(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["scsi_bus_sharing"] = cty.StringVal(p.ScsiBusSharing)
}

func EncodeVirtualMachine_ScsiControllerCount(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["scsi_controller_count"] = cty.NumberIntVal(p.ScsiControllerCount)
}

func EncodeVirtualMachine_ScsiType(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["scsi_type"] = cty.StringVal(p.ScsiType)
}

func EncodeVirtualMachine_ShutdownWaitTimeout(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["shutdown_wait_timeout"] = cty.NumberIntVal(p.ShutdownWaitTimeout)
}

func EncodeVirtualMachine_StoragePolicyId(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["storage_policy_id"] = cty.StringVal(p.StoragePolicyId)
}

func EncodeVirtualMachine_SwapPlacementPolicy(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["swap_placement_policy"] = cty.StringVal(p.SwapPlacementPolicy)
}

func EncodeVirtualMachine_SyncTimeWithHost(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["sync_time_with_host"] = cty.BoolVal(p.SyncTimeWithHost)
}

func EncodeVirtualMachine_Tags(p VirtualMachineParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Tags {
		colVals = append(colVals, cty.StringVal(value))
	}
	if len(colVals) == 0 {
		vals["tags"] = cty.SetValEmpty(cty.String)
	} else {
		vals["tags"] = cty.SetVal(colVals)
	}
}

func EncodeVirtualMachine_WaitForGuestIpTimeout(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["wait_for_guest_ip_timeout"] = cty.NumberIntVal(p.WaitForGuestIpTimeout)
}

func EncodeVirtualMachine_WaitForGuestNetRoutable(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["wait_for_guest_net_routable"] = cty.BoolVal(p.WaitForGuestNetRoutable)
}

func EncodeVirtualMachine_WaitForGuestNetTimeout(p VirtualMachineParameters, vals map[string]cty.Value) {
	vals["wait_for_guest_net_timeout"] = cty.NumberIntVal(p.WaitForGuestNetTimeout)
}

func EncodeVirtualMachine_Cdrom(p Cdrom, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeVirtualMachine_Cdrom_ClientDevice(p, ctyVal)
	EncodeVirtualMachine_Cdrom_DatastoreId(p, ctyVal)
	EncodeVirtualMachine_Cdrom_DeviceAddress(p, ctyVal)
	EncodeVirtualMachine_Cdrom_Key(p, ctyVal)
	EncodeVirtualMachine_Cdrom_Path(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	if len(valsForCollection) == 0 {
		vals["cdrom"] = cty.ListValEmpty(cty.EmptyObject)
	} else {
		vals["cdrom"] = cty.ListVal(valsForCollection)
	}
}

func EncodeVirtualMachine_Cdrom_ClientDevice(p Cdrom, vals map[string]cty.Value) {
	vals["client_device"] = cty.BoolVal(p.ClientDevice)
}

func EncodeVirtualMachine_Cdrom_DatastoreId(p Cdrom, vals map[string]cty.Value) {
	vals["datastore_id"] = cty.StringVal(p.DatastoreId)
}

func EncodeVirtualMachine_Cdrom_DeviceAddress(p Cdrom, vals map[string]cty.Value) {
	vals["device_address"] = cty.StringVal(p.DeviceAddress)
}

func EncodeVirtualMachine_Cdrom_Key(p Cdrom, vals map[string]cty.Value) {
	vals["key"] = cty.NumberIntVal(p.Key)
}

func EncodeVirtualMachine_Cdrom_Path(p Cdrom, vals map[string]cty.Value) {
	vals["path"] = cty.StringVal(p.Path)
}

func EncodeVirtualMachine_Clone(p Clone, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeVirtualMachine_Clone_TemplateUuid(p, ctyVal)
	EncodeVirtualMachine_Clone_Timeout(p, ctyVal)
	EncodeVirtualMachine_Clone_LinkedClone(p, ctyVal)
	EncodeVirtualMachine_Clone_OvfNetworkMap(p, ctyVal)
	EncodeVirtualMachine_Clone_OvfStorageMap(p, ctyVal)
	EncodeVirtualMachine_Clone_Customize(p.Customize, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	if len(valsForCollection) == 0 {
		vals["clone"] = cty.ListValEmpty(cty.EmptyObject)
	} else {
		vals["clone"] = cty.ListVal(valsForCollection)
	}
}

func EncodeVirtualMachine_Clone_TemplateUuid(p Clone, vals map[string]cty.Value) {
	vals["template_uuid"] = cty.StringVal(p.TemplateUuid)
}

func EncodeVirtualMachine_Clone_Timeout(p Clone, vals map[string]cty.Value) {
	vals["timeout"] = cty.NumberIntVal(p.Timeout)
}

func EncodeVirtualMachine_Clone_LinkedClone(p Clone, vals map[string]cty.Value) {
	vals["linked_clone"] = cty.BoolVal(p.LinkedClone)
}

func EncodeVirtualMachine_Clone_OvfNetworkMap(p Clone, vals map[string]cty.Value) {
	if len(p.OvfNetworkMap) == 0 {
		vals["ovf_network_map"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.OvfNetworkMap {
		mVals[key] = cty.StringVal(value)
	}
	vals["ovf_network_map"] = cty.MapVal(mVals)
}

func EncodeVirtualMachine_Clone_OvfStorageMap(p Clone, vals map[string]cty.Value) {
	if len(p.OvfStorageMap) == 0 {
		vals["ovf_storage_map"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.OvfStorageMap {
		mVals[key] = cty.StringVal(value)
	}
	vals["ovf_storage_map"] = cty.MapVal(mVals)
}

func EncodeVirtualMachine_Clone_Customize(p Customize, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeVirtualMachine_Clone_Customize_DnsSuffixList(p, ctyVal)
	EncodeVirtualMachine_Clone_Customize_Ipv4Gateway(p, ctyVal)
	EncodeVirtualMachine_Clone_Customize_Ipv6Gateway(p, ctyVal)
	EncodeVirtualMachine_Clone_Customize_Timeout(p, ctyVal)
	EncodeVirtualMachine_Clone_Customize_WindowsSysprepText(p, ctyVal)
	EncodeVirtualMachine_Clone_Customize_DnsServerList(p, ctyVal)
	EncodeVirtualMachine_Clone_Customize_LinuxOptions(p.LinuxOptions, ctyVal)
	EncodeVirtualMachine_Clone_Customize_NetworkInterface(p.NetworkInterface, ctyVal)
	EncodeVirtualMachine_Clone_Customize_WindowsOptions(p.WindowsOptions, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	if len(valsForCollection) == 0 {
		vals["customize"] = cty.ListValEmpty(cty.EmptyObject)
	} else {
		vals["customize"] = cty.ListVal(valsForCollection)
	}
}

func EncodeVirtualMachine_Clone_Customize_DnsSuffixList(p Customize, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.DnsSuffixList {
		colVals = append(colVals, cty.StringVal(value))
	}
	if len(colVals) == 0 {
		vals["dns_suffix_list"] = cty.ListValEmpty(cty.String)
	} else {
		vals["dns_suffix_list"] = cty.ListVal(colVals)
	}
}

func EncodeVirtualMachine_Clone_Customize_Ipv4Gateway(p Customize, vals map[string]cty.Value) {
	vals["ipv4_gateway"] = cty.StringVal(p.Ipv4Gateway)
}

func EncodeVirtualMachine_Clone_Customize_Ipv6Gateway(p Customize, vals map[string]cty.Value) {
	vals["ipv6_gateway"] = cty.StringVal(p.Ipv6Gateway)
}

func EncodeVirtualMachine_Clone_Customize_Timeout(p Customize, vals map[string]cty.Value) {
	vals["timeout"] = cty.NumberIntVal(p.Timeout)
}

func EncodeVirtualMachine_Clone_Customize_WindowsSysprepText(p Customize, vals map[string]cty.Value) {
	vals["windows_sysprep_text"] = cty.StringVal(p.WindowsSysprepText)
}

func EncodeVirtualMachine_Clone_Customize_DnsServerList(p Customize, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.DnsServerList {
		colVals = append(colVals, cty.StringVal(value))
	}
	if len(colVals) == 0 {
		vals["dns_server_list"] = cty.ListValEmpty(cty.String)
	} else {
		vals["dns_server_list"] = cty.ListVal(colVals)
	}
}

func EncodeVirtualMachine_Clone_Customize_LinuxOptions(p LinuxOptions, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeVirtualMachine_Clone_Customize_LinuxOptions_TimeZone(p, ctyVal)
	EncodeVirtualMachine_Clone_Customize_LinuxOptions_Domain(p, ctyVal)
	EncodeVirtualMachine_Clone_Customize_LinuxOptions_HostName(p, ctyVal)
	EncodeVirtualMachine_Clone_Customize_LinuxOptions_HwClockUtc(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	if len(valsForCollection) == 0 {
		vals["linux_options"] = cty.ListValEmpty(cty.EmptyObject)
	} else {
		vals["linux_options"] = cty.ListVal(valsForCollection)
	}
}

func EncodeVirtualMachine_Clone_Customize_LinuxOptions_TimeZone(p LinuxOptions, vals map[string]cty.Value) {
	vals["time_zone"] = cty.StringVal(p.TimeZone)
}

func EncodeVirtualMachine_Clone_Customize_LinuxOptions_Domain(p LinuxOptions, vals map[string]cty.Value) {
	vals["domain"] = cty.StringVal(p.Domain)
}

func EncodeVirtualMachine_Clone_Customize_LinuxOptions_HostName(p LinuxOptions, vals map[string]cty.Value) {
	vals["host_name"] = cty.StringVal(p.HostName)
}

func EncodeVirtualMachine_Clone_Customize_LinuxOptions_HwClockUtc(p LinuxOptions, vals map[string]cty.Value) {
	vals["hw_clock_utc"] = cty.BoolVal(p.HwClockUtc)
}

func EncodeVirtualMachine_Clone_Customize_NetworkInterface(p NetworkInterface, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeVirtualMachine_Clone_Customize_NetworkInterface_DnsDomain(p, ctyVal)
	EncodeVirtualMachine_Clone_Customize_NetworkInterface_DnsServerList(p, ctyVal)
	EncodeVirtualMachine_Clone_Customize_NetworkInterface_Ipv4Address(p, ctyVal)
	EncodeVirtualMachine_Clone_Customize_NetworkInterface_Ipv4Netmask(p, ctyVal)
	EncodeVirtualMachine_Clone_Customize_NetworkInterface_Ipv6Address(p, ctyVal)
	EncodeVirtualMachine_Clone_Customize_NetworkInterface_Ipv6Netmask(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	if len(valsForCollection) == 0 {
		vals["network_interface"] = cty.ListValEmpty(cty.EmptyObject)
	} else {
		vals["network_interface"] = cty.ListVal(valsForCollection)
	}
}

func EncodeVirtualMachine_Clone_Customize_NetworkInterface_DnsDomain(p NetworkInterface, vals map[string]cty.Value) {
	vals["dns_domain"] = cty.StringVal(p.DnsDomain)
}

func EncodeVirtualMachine_Clone_Customize_NetworkInterface_DnsServerList(p NetworkInterface, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.DnsServerList {
		colVals = append(colVals, cty.StringVal(value))
	}
	if len(colVals) == 0 {
		vals["dns_server_list"] = cty.ListValEmpty(cty.String)
	} else {
		vals["dns_server_list"] = cty.ListVal(colVals)
	}
}

func EncodeVirtualMachine_Clone_Customize_NetworkInterface_Ipv4Address(p NetworkInterface, vals map[string]cty.Value) {
	vals["ipv4_address"] = cty.StringVal(p.Ipv4Address)
}

func EncodeVirtualMachine_Clone_Customize_NetworkInterface_Ipv4Netmask(p NetworkInterface, vals map[string]cty.Value) {
	vals["ipv4_netmask"] = cty.NumberIntVal(p.Ipv4Netmask)
}

func EncodeVirtualMachine_Clone_Customize_NetworkInterface_Ipv6Address(p NetworkInterface, vals map[string]cty.Value) {
	vals["ipv6_address"] = cty.StringVal(p.Ipv6Address)
}

func EncodeVirtualMachine_Clone_Customize_NetworkInterface_Ipv6Netmask(p NetworkInterface, vals map[string]cty.Value) {
	vals["ipv6_netmask"] = cty.NumberIntVal(p.Ipv6Netmask)
}

func EncodeVirtualMachine_Clone_Customize_WindowsOptions(p WindowsOptions, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeVirtualMachine_Clone_Customize_WindowsOptions_AutoLogonCount(p, ctyVal)
	EncodeVirtualMachine_Clone_Customize_WindowsOptions_ComputerName(p, ctyVal)
	EncodeVirtualMachine_Clone_Customize_WindowsOptions_DomainAdminUser(p, ctyVal)
	EncodeVirtualMachine_Clone_Customize_WindowsOptions_FullName(p, ctyVal)
	EncodeVirtualMachine_Clone_Customize_WindowsOptions_OrganizationName(p, ctyVal)
	EncodeVirtualMachine_Clone_Customize_WindowsOptions_ProductKey(p, ctyVal)
	EncodeVirtualMachine_Clone_Customize_WindowsOptions_Workgroup(p, ctyVal)
	EncodeVirtualMachine_Clone_Customize_WindowsOptions_AdminPassword(p, ctyVal)
	EncodeVirtualMachine_Clone_Customize_WindowsOptions_AutoLogon(p, ctyVal)
	EncodeVirtualMachine_Clone_Customize_WindowsOptions_DomainAdminPassword(p, ctyVal)
	EncodeVirtualMachine_Clone_Customize_WindowsOptions_JoinDomain(p, ctyVal)
	EncodeVirtualMachine_Clone_Customize_WindowsOptions_RunOnceCommandList(p, ctyVal)
	EncodeVirtualMachine_Clone_Customize_WindowsOptions_TimeZone(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	if len(valsForCollection) == 0 {
		vals["windows_options"] = cty.ListValEmpty(cty.EmptyObject)
	} else {
		vals["windows_options"] = cty.ListVal(valsForCollection)
	}
}

func EncodeVirtualMachine_Clone_Customize_WindowsOptions_AutoLogonCount(p WindowsOptions, vals map[string]cty.Value) {
	vals["auto_logon_count"] = cty.NumberIntVal(p.AutoLogonCount)
}

func EncodeVirtualMachine_Clone_Customize_WindowsOptions_ComputerName(p WindowsOptions, vals map[string]cty.Value) {
	vals["computer_name"] = cty.StringVal(p.ComputerName)
}

func EncodeVirtualMachine_Clone_Customize_WindowsOptions_DomainAdminUser(p WindowsOptions, vals map[string]cty.Value) {
	vals["domain_admin_user"] = cty.StringVal(p.DomainAdminUser)
}

func EncodeVirtualMachine_Clone_Customize_WindowsOptions_FullName(p WindowsOptions, vals map[string]cty.Value) {
	vals["full_name"] = cty.StringVal(p.FullName)
}

func EncodeVirtualMachine_Clone_Customize_WindowsOptions_OrganizationName(p WindowsOptions, vals map[string]cty.Value) {
	vals["organization_name"] = cty.StringVal(p.OrganizationName)
}

func EncodeVirtualMachine_Clone_Customize_WindowsOptions_ProductKey(p WindowsOptions, vals map[string]cty.Value) {
	vals["product_key"] = cty.StringVal(p.ProductKey)
}

func EncodeVirtualMachine_Clone_Customize_WindowsOptions_Workgroup(p WindowsOptions, vals map[string]cty.Value) {
	vals["workgroup"] = cty.StringVal(p.Workgroup)
}

func EncodeVirtualMachine_Clone_Customize_WindowsOptions_AdminPassword(p WindowsOptions, vals map[string]cty.Value) {
	vals["admin_password"] = cty.StringVal(p.AdminPassword)
}

func EncodeVirtualMachine_Clone_Customize_WindowsOptions_AutoLogon(p WindowsOptions, vals map[string]cty.Value) {
	vals["auto_logon"] = cty.BoolVal(p.AutoLogon)
}

func EncodeVirtualMachine_Clone_Customize_WindowsOptions_DomainAdminPassword(p WindowsOptions, vals map[string]cty.Value) {
	vals["domain_admin_password"] = cty.StringVal(p.DomainAdminPassword)
}

func EncodeVirtualMachine_Clone_Customize_WindowsOptions_JoinDomain(p WindowsOptions, vals map[string]cty.Value) {
	vals["join_domain"] = cty.StringVal(p.JoinDomain)
}

func EncodeVirtualMachine_Clone_Customize_WindowsOptions_RunOnceCommandList(p WindowsOptions, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.RunOnceCommandList {
		colVals = append(colVals, cty.StringVal(value))
	}
	if len(colVals) == 0 {
		vals["run_once_command_list"] = cty.ListValEmpty(cty.String)
	} else {
		vals["run_once_command_list"] = cty.ListVal(colVals)
	}
}

func EncodeVirtualMachine_Clone_Customize_WindowsOptions_TimeZone(p WindowsOptions, vals map[string]cty.Value) {
	vals["time_zone"] = cty.NumberIntVal(p.TimeZone)
}

func EncodeVirtualMachine_Disk(p []Disk, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeVirtualMachine_Disk_DiskMode(v, ctyVal)
		EncodeVirtualMachine_Disk_KeepOnRemove(v, ctyVal)
		EncodeVirtualMachine_Disk_Path(v, ctyVal)
		EncodeVirtualMachine_Disk_Size(v, ctyVal)
		EncodeVirtualMachine_Disk_StoragePolicyId(v, ctyVal)
		EncodeVirtualMachine_Disk_UnitNumber(v, ctyVal)
		EncodeVirtualMachine_Disk_ControllerType(v, ctyVal)
		EncodeVirtualMachine_Disk_DeviceAddress(v, ctyVal)
		EncodeVirtualMachine_Disk_Uuid(v, ctyVal)
		EncodeVirtualMachine_Disk_IoShareCount(v, ctyVal)
		EncodeVirtualMachine_Disk_Name(v, ctyVal)
		EncodeVirtualMachine_Disk_ThinProvisioned(v, ctyVal)
		EncodeVirtualMachine_Disk_Attach(v, ctyVal)
		EncodeVirtualMachine_Disk_IoLimit(v, ctyVal)
		EncodeVirtualMachine_Disk_Key(v, ctyVal)
		EncodeVirtualMachine_Disk_Label(v, ctyVal)
		EncodeVirtualMachine_Disk_WriteThrough(v, ctyVal)
		EncodeVirtualMachine_Disk_DatastoreId(v, ctyVal)
		EncodeVirtualMachine_Disk_EagerlyScrub(v, ctyVal)
		EncodeVirtualMachine_Disk_IoShareLevel(v, ctyVal)
		EncodeVirtualMachine_Disk_DiskSharing(v, ctyVal)
		EncodeVirtualMachine_Disk_IoReservation(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	if len(valsForCollection) == 0 {
		vals["disk"] = cty.ListValEmpty(cty.EmptyObject)
	} else {
		vals["disk"] = cty.ListVal(valsForCollection)
	}
}

func EncodeVirtualMachine_Disk_DiskMode(p Disk, vals map[string]cty.Value) {
	vals["disk_mode"] = cty.StringVal(p.DiskMode)
}

func EncodeVirtualMachine_Disk_KeepOnRemove(p Disk, vals map[string]cty.Value) {
	vals["keep_on_remove"] = cty.BoolVal(p.KeepOnRemove)
}

func EncodeVirtualMachine_Disk_Path(p Disk, vals map[string]cty.Value) {
	vals["path"] = cty.StringVal(p.Path)
}

func EncodeVirtualMachine_Disk_Size(p Disk, vals map[string]cty.Value) {
	vals["size"] = cty.NumberIntVal(p.Size)
}

func EncodeVirtualMachine_Disk_StoragePolicyId(p Disk, vals map[string]cty.Value) {
	vals["storage_policy_id"] = cty.StringVal(p.StoragePolicyId)
}

func EncodeVirtualMachine_Disk_UnitNumber(p Disk, vals map[string]cty.Value) {
	vals["unit_number"] = cty.NumberIntVal(p.UnitNumber)
}

func EncodeVirtualMachine_Disk_ControllerType(p Disk, vals map[string]cty.Value) {
	vals["controller_type"] = cty.StringVal(p.ControllerType)
}

func EncodeVirtualMachine_Disk_DeviceAddress(p Disk, vals map[string]cty.Value) {
	vals["device_address"] = cty.StringVal(p.DeviceAddress)
}

func EncodeVirtualMachine_Disk_Uuid(p Disk, vals map[string]cty.Value) {
	vals["uuid"] = cty.StringVal(p.Uuid)
}

func EncodeVirtualMachine_Disk_IoShareCount(p Disk, vals map[string]cty.Value) {
	vals["io_share_count"] = cty.NumberIntVal(p.IoShareCount)
}

func EncodeVirtualMachine_Disk_Name(p Disk, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeVirtualMachine_Disk_ThinProvisioned(p Disk, vals map[string]cty.Value) {
	vals["thin_provisioned"] = cty.BoolVal(p.ThinProvisioned)
}

func EncodeVirtualMachine_Disk_Attach(p Disk, vals map[string]cty.Value) {
	vals["attach"] = cty.BoolVal(p.Attach)
}

func EncodeVirtualMachine_Disk_IoLimit(p Disk, vals map[string]cty.Value) {
	vals["io_limit"] = cty.NumberIntVal(p.IoLimit)
}

func EncodeVirtualMachine_Disk_Key(p Disk, vals map[string]cty.Value) {
	vals["key"] = cty.NumberIntVal(p.Key)
}

func EncodeVirtualMachine_Disk_Label(p Disk, vals map[string]cty.Value) {
	vals["label"] = cty.StringVal(p.Label)
}

func EncodeVirtualMachine_Disk_WriteThrough(p Disk, vals map[string]cty.Value) {
	vals["write_through"] = cty.BoolVal(p.WriteThrough)
}

func EncodeVirtualMachine_Disk_DatastoreId(p Disk, vals map[string]cty.Value) {
	vals["datastore_id"] = cty.StringVal(p.DatastoreId)
}

func EncodeVirtualMachine_Disk_EagerlyScrub(p Disk, vals map[string]cty.Value) {
	vals["eagerly_scrub"] = cty.BoolVal(p.EagerlyScrub)
}

func EncodeVirtualMachine_Disk_IoShareLevel(p Disk, vals map[string]cty.Value) {
	vals["io_share_level"] = cty.StringVal(p.IoShareLevel)
}

func EncodeVirtualMachine_Disk_DiskSharing(p Disk, vals map[string]cty.Value) {
	vals["disk_sharing"] = cty.StringVal(p.DiskSharing)
}

func EncodeVirtualMachine_Disk_IoReservation(p Disk, vals map[string]cty.Value) {
	vals["io_reservation"] = cty.NumberIntVal(p.IoReservation)
}

func EncodeVirtualMachine_NetworkInterface(p []NetworkInterface0, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeVirtualMachine_NetworkInterface_DeviceAddress(v, ctyVal)
		EncodeVirtualMachine_NetworkInterface_Key(v, ctyVal)
		EncodeVirtualMachine_NetworkInterface_MacAddress(v, ctyVal)
		EncodeVirtualMachine_NetworkInterface_NetworkId(v, ctyVal)
		EncodeVirtualMachine_NetworkInterface_OvfMapping(v, ctyVal)
		EncodeVirtualMachine_NetworkInterface_BandwidthLimit(v, ctyVal)
		EncodeVirtualMachine_NetworkInterface_BandwidthReservation(v, ctyVal)
		EncodeVirtualMachine_NetworkInterface_BandwidthShareLevel(v, ctyVal)
		EncodeVirtualMachine_NetworkInterface_UseStaticMac(v, ctyVal)
		EncodeVirtualMachine_NetworkInterface_AdapterType(v, ctyVal)
		EncodeVirtualMachine_NetworkInterface_BandwidthShareCount(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	if len(valsForCollection) == 0 {
		vals["network_interface"] = cty.ListValEmpty(cty.EmptyObject)
	} else {
		vals["network_interface"] = cty.ListVal(valsForCollection)
	}
}

func EncodeVirtualMachine_NetworkInterface_DeviceAddress(p NetworkInterface0, vals map[string]cty.Value) {
	vals["device_address"] = cty.StringVal(p.DeviceAddress)
}

func EncodeVirtualMachine_NetworkInterface_Key(p NetworkInterface0, vals map[string]cty.Value) {
	vals["key"] = cty.NumberIntVal(p.Key)
}

func EncodeVirtualMachine_NetworkInterface_MacAddress(p NetworkInterface0, vals map[string]cty.Value) {
	vals["mac_address"] = cty.StringVal(p.MacAddress)
}

func EncodeVirtualMachine_NetworkInterface_NetworkId(p NetworkInterface0, vals map[string]cty.Value) {
	vals["network_id"] = cty.StringVal(p.NetworkId)
}

func EncodeVirtualMachine_NetworkInterface_OvfMapping(p NetworkInterface0, vals map[string]cty.Value) {
	vals["ovf_mapping"] = cty.StringVal(p.OvfMapping)
}

func EncodeVirtualMachine_NetworkInterface_BandwidthLimit(p NetworkInterface0, vals map[string]cty.Value) {
	vals["bandwidth_limit"] = cty.NumberIntVal(p.BandwidthLimit)
}

func EncodeVirtualMachine_NetworkInterface_BandwidthReservation(p NetworkInterface0, vals map[string]cty.Value) {
	vals["bandwidth_reservation"] = cty.NumberIntVal(p.BandwidthReservation)
}

func EncodeVirtualMachine_NetworkInterface_BandwidthShareLevel(p NetworkInterface0, vals map[string]cty.Value) {
	vals["bandwidth_share_level"] = cty.StringVal(p.BandwidthShareLevel)
}

func EncodeVirtualMachine_NetworkInterface_UseStaticMac(p NetworkInterface0, vals map[string]cty.Value) {
	vals["use_static_mac"] = cty.BoolVal(p.UseStaticMac)
}

func EncodeVirtualMachine_NetworkInterface_AdapterType(p NetworkInterface0, vals map[string]cty.Value) {
	vals["adapter_type"] = cty.StringVal(p.AdapterType)
}

func EncodeVirtualMachine_NetworkInterface_BandwidthShareCount(p NetworkInterface0, vals map[string]cty.Value) {
	vals["bandwidth_share_count"] = cty.NumberIntVal(p.BandwidthShareCount)
}

func EncodeVirtualMachine_OvfDeploy(p OvfDeploy, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeVirtualMachine_OvfDeploy_IpProtocol(p, ctyVal)
	EncodeVirtualMachine_OvfDeploy_LocalOvfPath(p, ctyVal)
	EncodeVirtualMachine_OvfDeploy_OvfNetworkMap(p, ctyVal)
	EncodeVirtualMachine_OvfDeploy_RemoteOvfUrl(p, ctyVal)
	EncodeVirtualMachine_OvfDeploy_AllowUnverifiedSslCert(p, ctyVal)
	EncodeVirtualMachine_OvfDeploy_DiskProvisioning(p, ctyVal)
	EncodeVirtualMachine_OvfDeploy_IpAllocationPolicy(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	if len(valsForCollection) == 0 {
		vals["ovf_deploy"] = cty.ListValEmpty(cty.EmptyObject)
	} else {
		vals["ovf_deploy"] = cty.ListVal(valsForCollection)
	}
}

func EncodeVirtualMachine_OvfDeploy_IpProtocol(p OvfDeploy, vals map[string]cty.Value) {
	vals["ip_protocol"] = cty.StringVal(p.IpProtocol)
}

func EncodeVirtualMachine_OvfDeploy_LocalOvfPath(p OvfDeploy, vals map[string]cty.Value) {
	vals["local_ovf_path"] = cty.StringVal(p.LocalOvfPath)
}

func EncodeVirtualMachine_OvfDeploy_OvfNetworkMap(p OvfDeploy, vals map[string]cty.Value) {
	if len(p.OvfNetworkMap) == 0 {
		vals["ovf_network_map"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.OvfNetworkMap {
		mVals[key] = cty.StringVal(value)
	}
	vals["ovf_network_map"] = cty.MapVal(mVals)
}

func EncodeVirtualMachine_OvfDeploy_RemoteOvfUrl(p OvfDeploy, vals map[string]cty.Value) {
	vals["remote_ovf_url"] = cty.StringVal(p.RemoteOvfUrl)
}

func EncodeVirtualMachine_OvfDeploy_AllowUnverifiedSslCert(p OvfDeploy, vals map[string]cty.Value) {
	vals["allow_unverified_ssl_cert"] = cty.BoolVal(p.AllowUnverifiedSslCert)
}

func EncodeVirtualMachine_OvfDeploy_DiskProvisioning(p OvfDeploy, vals map[string]cty.Value) {
	vals["disk_provisioning"] = cty.StringVal(p.DiskProvisioning)
}

func EncodeVirtualMachine_OvfDeploy_IpAllocationPolicy(p OvfDeploy, vals map[string]cty.Value) {
	vals["ip_allocation_policy"] = cty.StringVal(p.IpAllocationPolicy)
}

func EncodeVirtualMachine_Vapp(p Vapp, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeVirtualMachine_Vapp_Properties(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	if len(valsForCollection) == 0 {
		vals["vapp"] = cty.ListValEmpty(cty.EmptyObject)
	} else {
		vals["vapp"] = cty.ListVal(valsForCollection)
	}
}

func EncodeVirtualMachine_Vapp_Properties(p Vapp, vals map[string]cty.Value) {
	if len(p.Properties) == 0 {
		vals["properties"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.Properties {
		mVals[key] = cty.StringVal(value)
	}
	vals["properties"] = cty.MapVal(mVals)
}

func EncodeVirtualMachine_ChangeVersion(p VirtualMachineObservation, vals map[string]cty.Value) {
	vals["change_version"] = cty.StringVal(p.ChangeVersion)
}

func EncodeVirtualMachine_DefaultIpAddress(p VirtualMachineObservation, vals map[string]cty.Value) {
	vals["default_ip_address"] = cty.StringVal(p.DefaultIpAddress)
}

func EncodeVirtualMachine_GuestIpAddresses(p VirtualMachineObservation, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.GuestIpAddresses {
		colVals = append(colVals, cty.StringVal(value))
	}
	if len(colVals) == 0 {
		vals["guest_ip_addresses"] = cty.ListValEmpty(cty.String)
	} else {
		vals["guest_ip_addresses"] = cty.ListVal(colVals)
	}
}

func EncodeVirtualMachine_Imported(p VirtualMachineObservation, vals map[string]cty.Value) {
	vals["imported"] = cty.BoolVal(p.Imported)
}

func EncodeVirtualMachine_Moid(p VirtualMachineObservation, vals map[string]cty.Value) {
	vals["moid"] = cty.StringVal(p.Moid)
}

func EncodeVirtualMachine_RebootRequired(p VirtualMachineObservation, vals map[string]cty.Value) {
	vals["reboot_required"] = cty.BoolVal(p.RebootRequired)
}

func EncodeVirtualMachine_Uuid(p VirtualMachineObservation, vals map[string]cty.Value) {
	vals["uuid"] = cty.StringVal(p.Uuid)
}

func EncodeVirtualMachine_VappTransport(p VirtualMachineObservation, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.VappTransport {
		colVals = append(colVals, cty.StringVal(value))
	}
	if len(colVals) == 0 {
		vals["vapp_transport"] = cty.ListValEmpty(cty.String)
	} else {
		vals["vapp_transport"] = cty.ListVal(colVals)
	}
}

func EncodeVirtualMachine_VmwareToolsStatus(p VirtualMachineObservation, vals map[string]cty.Value) {
	vals["vmware_tools_status"] = cty.StringVal(p.VmwareToolsStatus)
}

func EncodeVirtualMachine_VmxPath(p VirtualMachineObservation, vals map[string]cty.Value) {
	vals["vmx_path"] = cty.StringVal(p.VmxPath)
}

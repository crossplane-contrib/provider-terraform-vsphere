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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// +kubebuilder:object:root=true

// VirtualMachine is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type VirtualMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualMachineSpec   `json:"spec"`
	Status VirtualMachineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualMachine contains a list of VirtualMachineList
type VirtualMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualMachine `json:"items"`
}

// A VirtualMachineSpec defines the desired state of a VirtualMachine
type VirtualMachineSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       VirtualMachineParameters `json:"forProvider"`
}

// A VirtualMachineParameters defines the desired state of a VirtualMachine
type VirtualMachineParameters struct {
	AlternateGuestName                 string              `json:"alternate_guest_name"`
	Annotation                         string              `json:"annotation"`
	BootDelay                          int64               `json:"boot_delay"`
	BootRetryDelay                     int64               `json:"boot_retry_delay"`
	BootRetryEnabled                   bool                `json:"boot_retry_enabled"`
	CpuHotAddEnabled                   bool                `json:"cpu_hot_add_enabled"`
	CpuHotRemoveEnabled                bool                `json:"cpu_hot_remove_enabled"`
	CpuLimit                           int64               `json:"cpu_limit"`
	CpuPerformanceCountersEnabled      bool                `json:"cpu_performance_counters_enabled"`
	CpuReservation                     int64               `json:"cpu_reservation"`
	CpuShareCount                      int64               `json:"cpu_share_count"`
	CpuShareLevel                      string              `json:"cpu_share_level"`
	CustomAttributes                   map[string]string   `json:"custom_attributes,omitempty"`
	DatacenterId                       string              `json:"datacenter_id"`
	DatastoreClusterId                 string              `json:"datastore_cluster_id"`
	DatastoreId                        string              `json:"datastore_id"`
	EfiSecureBootEnabled               bool                `json:"efi_secure_boot_enabled"`
	EnableDiskUuid                     bool                `json:"enable_disk_uuid"`
	EnableLogging                      bool                `json:"enable_logging"`
	EptRviMode                         string              `json:"ept_rvi_mode"`
	ExtraConfig                        map[string]string   `json:"extra_config,omitempty"`
	Firmware                           string              `json:"firmware"`
	Folder                             string              `json:"folder"`
	ForcePowerOff                      bool                `json:"force_power_off"`
	GuestId                            string              `json:"guest_id"`
	HardwareVersion                    int64               `json:"hardware_version"`
	HostSystemId                       string              `json:"host_system_id"`
	HvMode                             string              `json:"hv_mode"`
	IdeControllerCount                 int64               `json:"ide_controller_count"`
	IgnoredGuestIps                    []string            `json:"ignored_guest_ips,omitempty"`
	LatencySensitivity                 string              `json:"latency_sensitivity"`
	Memory                             int64               `json:"memory"`
	MemoryHotAddEnabled                bool                `json:"memory_hot_add_enabled"`
	MemoryLimit                        int64               `json:"memory_limit"`
	MemoryReservation                  int64               `json:"memory_reservation"`
	MemoryShareCount                   int64               `json:"memory_share_count"`
	MemoryShareLevel                   string              `json:"memory_share_level"`
	MigrateWaitTimeout                 int64               `json:"migrate_wait_timeout"`
	Name                               string              `json:"name"`
	NestedHvEnabled                    bool                `json:"nested_hv_enabled"`
	NumCoresPerSocket                  int64               `json:"num_cores_per_socket"`
	NumCpus                            int64               `json:"num_cpus"`
	PciDeviceId                        []string            `json:"pci_device_id,omitempty"`
	PoweronTimeout                     int64               `json:"poweron_timeout"`
	ResourcePoolId                     string              `json:"resource_pool_id"`
	RunToolsScriptsAfterPowerOn        bool                `json:"run_tools_scripts_after_power_on"`
	RunToolsScriptsAfterResume         bool                `json:"run_tools_scripts_after_resume"`
	RunToolsScriptsBeforeGuestReboot   bool                `json:"run_tools_scripts_before_guest_reboot"`
	RunToolsScriptsBeforeGuestShutdown bool                `json:"run_tools_scripts_before_guest_shutdown"`
	RunToolsScriptsBeforeGuestStandby  bool                `json:"run_tools_scripts_before_guest_standby"`
	SataControllerCount                int64               `json:"sata_controller_count"`
	ScsiBusSharing                     string              `json:"scsi_bus_sharing"`
	ScsiControllerCount                int64               `json:"scsi_controller_count"`
	ScsiType                           string              `json:"scsi_type"`
	ShutdownWaitTimeout                int64               `json:"shutdown_wait_timeout"`
	StoragePolicyId                    string              `json:"storage_policy_id"`
	SwapPlacementPolicy                string              `json:"swap_placement_policy"`
	SyncTimeWithHost                   bool                `json:"sync_time_with_host"`
	Tags                               []string            `json:"tags,omitempty"`
	WaitForGuestIpTimeout              int64               `json:"wait_for_guest_ip_timeout"`
	WaitForGuestNetRoutable            bool                `json:"wait_for_guest_net_routable"`
	WaitForGuestNetTimeout             int64               `json:"wait_for_guest_net_timeout"`
	Cdrom                              Cdrom               `json:"cdrom"`
	Clone                              Clone               `json:"clone"`
	Disk                               []Disk              `json:"disk"`
	NetworkInterface                   []NetworkInterface0 `json:"network_interface"`
	OvfDeploy                          OvfDeploy           `json:"ovf_deploy"`
	Vapp                               Vapp                `json:"vapp"`
}

type Cdrom struct {
	ClientDevice  bool   `json:"client_device"`
	DatastoreId   string `json:"datastore_id"`
	DeviceAddress string `json:"device_address"`
	Key           int64  `json:"key"`
	Path          string `json:"path"`
}

type Clone struct {
	LinkedClone   bool              `json:"linked_clone"`
	OvfNetworkMap map[string]string `json:"ovf_network_map,omitempty"`
	OvfStorageMap map[string]string `json:"ovf_storage_map,omitempty"`
	TemplateUuid  string            `json:"template_uuid"`
	Timeout       int64             `json:"timeout"`
	Customize     Customize         `json:"customize"`
}

type Customize struct {
	DnsServerList      []string         `json:"dns_server_list,omitempty"`
	DnsSuffixList      []string         `json:"dns_suffix_list,omitempty"`
	Ipv4Gateway        string           `json:"ipv4_gateway"`
	Ipv6Gateway        string           `json:"ipv6_gateway"`
	Timeout            int64            `json:"timeout"`
	WindowsSysprepText string           `json:"windows_sysprep_text"`
	LinuxOptions       LinuxOptions     `json:"linux_options"`
	NetworkInterface   NetworkInterface `json:"network_interface"`
	WindowsOptions     WindowsOptions   `json:"windows_options"`
}

type LinuxOptions struct {
	Domain     string `json:"domain"`
	HostName   string `json:"host_name"`
	HwClockUtc bool   `json:"hw_clock_utc"`
	TimeZone   string `json:"time_zone"`
}

type NetworkInterface struct {
	DnsDomain     string   `json:"dns_domain"`
	DnsServerList []string `json:"dns_server_list,omitempty"`
	Ipv4Address   string   `json:"ipv4_address"`
	Ipv4Netmask   int64    `json:"ipv4_netmask"`
	Ipv6Address   string   `json:"ipv6_address"`
	Ipv6Netmask   int64    `json:"ipv6_netmask"`
}

type WindowsOptions struct {
	DomainAdminPassword string   `json:"domain_admin_password"`
	DomainAdminUser     string   `json:"domain_admin_user"`
	JoinDomain          string   `json:"join_domain"`
	ProductKey          string   `json:"product_key"`
	AdminPassword       string   `json:"admin_password"`
	AutoLogon           bool     `json:"auto_logon"`
	AutoLogonCount      int64    `json:"auto_logon_count"`
	ComputerName        string   `json:"computer_name"`
	RunOnceCommandList  []string `json:"run_once_command_list,omitempty"`
	Workgroup           string   `json:"workgroup"`
	FullName            string   `json:"full_name"`
	OrganizationName    string   `json:"organization_name"`
	TimeZone            int64    `json:"time_zone"`
}

type Disk struct {
	WriteThrough    bool   `json:"write_through"`
	IoShareLevel    string `json:"io_share_level"`
	Label           string `json:"label"`
	Path            string `json:"path"`
	Key             int64  `json:"key"`
	Name            string `json:"name"`
	ThinProvisioned bool   `json:"thin_provisioned"`
	Uuid            string `json:"uuid"`
	DatastoreId     string `json:"datastore_id"`
	DiskMode        string `json:"disk_mode"`
	EagerlyScrub    bool   `json:"eagerly_scrub"`
	KeepOnRemove    bool   `json:"keep_on_remove"`
	Attach          bool   `json:"attach"`
	ControllerType  string `json:"controller_type"`
	IoShareCount    int64  `json:"io_share_count"`
	IoReservation   int64  `json:"io_reservation"`
	Size            int64  `json:"size"`
	StoragePolicyId string `json:"storage_policy_id"`
	UnitNumber      int64  `json:"unit_number"`
	DeviceAddress   string `json:"device_address"`
	DiskSharing     string `json:"disk_sharing"`
	IoLimit         int64  `json:"io_limit"`
}

type NetworkInterface0 struct {
	Key                  int64  `json:"key"`
	AdapterType          string `json:"adapter_type"`
	BandwidthLimit       int64  `json:"bandwidth_limit"`
	DeviceAddress        string `json:"device_address"`
	MacAddress           string `json:"mac_address"`
	NetworkId            string `json:"network_id"`
	OvfMapping           string `json:"ovf_mapping"`
	UseStaticMac         bool   `json:"use_static_mac"`
	BandwidthReservation int64  `json:"bandwidth_reservation"`
	BandwidthShareCount  int64  `json:"bandwidth_share_count"`
	BandwidthShareLevel  string `json:"bandwidth_share_level"`
}

type OvfDeploy struct {
	RemoteOvfUrl           string            `json:"remote_ovf_url"`
	AllowUnverifiedSslCert bool              `json:"allow_unverified_ssl_cert"`
	DiskProvisioning       string            `json:"disk_provisioning"`
	IpAllocationPolicy     string            `json:"ip_allocation_policy"`
	IpProtocol             string            `json:"ip_protocol"`
	LocalOvfPath           string            `json:"local_ovf_path"`
	OvfNetworkMap          map[string]string `json:"ovf_network_map,omitempty"`
}

type Vapp struct {
	Properties map[string]string `json:"properties,omitempty"`
}

// A VirtualMachineStatus defines the observed state of a VirtualMachine
type VirtualMachineStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          VirtualMachineObservation `json:"atProvider"`
}

// A VirtualMachineObservation records the observed state of a VirtualMachine
type VirtualMachineObservation struct {
	ChangeVersion     string   `json:"change_version"`
	DefaultIpAddress  string   `json:"default_ip_address"`
	GuestIpAddresses  []string `json:"guest_ip_addresses,omitempty"`
	Imported          bool     `json:"imported"`
	Moid              string   `json:"moid"`
	RebootRequired    bool     `json:"reboot_required"`
	Uuid              string   `json:"uuid"`
	VappTransport     []string `json:"vapp_transport,omitempty"`
	VmwareToolsStatus string   `json:"vmware_tools_status"`
	VmxPath           string   `json:"vmx_path"`
}
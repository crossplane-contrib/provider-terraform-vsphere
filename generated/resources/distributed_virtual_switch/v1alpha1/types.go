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

// DistributedVirtualSwitch is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DistributedVirtualSwitch struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DistributedVirtualSwitchSpec   `json:"spec"`
	Status DistributedVirtualSwitchStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DistributedVirtualSwitch contains a list of DistributedVirtualSwitchList
type DistributedVirtualSwitchList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DistributedVirtualSwitch `json:"items"`
}

// A DistributedVirtualSwitchSpec defines the desired state of a DistributedVirtualSwitch
type DistributedVirtualSwitchSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DistributedVirtualSwitchParameters `json:"forProvider"`
}

// A DistributedVirtualSwitchParameters defines the desired state of a DistributedVirtualSwitch
type DistributedVirtualSwitchParameters struct {
	ActiveUplinks                  []string          `json:"active_uplinks,omitempty"`
	AllowForgedTransmits           bool              `json:"allow_forged_transmits"`
	AllowMacChanges                bool              `json:"allow_mac_changes"`
	AllowPromiscuous               bool              `json:"allow_promiscuous"`
	BlockAllPorts                  bool              `json:"block_all_ports"`
	CheckBeacon                    bool              `json:"check_beacon"`
	ContactDetail                  string            `json:"contact_detail"`
	ContactName                    string            `json:"contact_name"`
	CustomAttributes               map[string]string `json:"custom_attributes,omitempty"`
	DatacenterId                   string            `json:"datacenter_id"`
	Description                    string            `json:"description"`
	DirectpathGen2Allowed          bool              `json:"directpath_gen2_allowed"`
	EgressShapingAverageBandwidth  int64             `json:"egress_shaping_average_bandwidth"`
	EgressShapingBurstSize         int64             `json:"egress_shaping_burst_size"`
	EgressShapingEnabled           bool              `json:"egress_shaping_enabled"`
	EgressShapingPeakBandwidth     int64             `json:"egress_shaping_peak_bandwidth"`
	Failback                       bool              `json:"failback"`
	FaulttoleranceMaximumMbit      int64             `json:"faulttolerance_maximum_mbit"`
	FaulttoleranceReservationMbit  int64             `json:"faulttolerance_reservation_mbit"`
	FaulttoleranceShareCount       int64             `json:"faulttolerance_share_count"`
	FaulttoleranceShareLevel       string            `json:"faulttolerance_share_level"`
	Folder                         string            `json:"folder"`
	HbrMaximumMbit                 int64             `json:"hbr_maximum_mbit"`
	HbrReservationMbit             int64             `json:"hbr_reservation_mbit"`
	HbrShareCount                  int64             `json:"hbr_share_count"`
	HbrShareLevel                  string            `json:"hbr_share_level"`
	IgnoreOtherPvlanMappings       bool              `json:"ignore_other_pvlan_mappings"`
	IngressShapingAverageBandwidth int64             `json:"ingress_shaping_average_bandwidth"`
	IngressShapingBurstSize        int64             `json:"ingress_shaping_burst_size"`
	IngressShapingEnabled          bool              `json:"ingress_shaping_enabled"`
	IngressShapingPeakBandwidth    int64             `json:"ingress_shaping_peak_bandwidth"`
	Ipv4Address                    string            `json:"ipv4_address"`
	IscsiMaximumMbit               int64             `json:"iscsi_maximum_mbit"`
	IscsiReservationMbit           int64             `json:"iscsi_reservation_mbit"`
	IscsiShareCount                int64             `json:"iscsi_share_count"`
	IscsiShareLevel                string            `json:"iscsi_share_level"`
	LacpApiVersion                 string            `json:"lacp_api_version"`
	LacpEnabled                    bool              `json:"lacp_enabled"`
	LacpMode                       string            `json:"lacp_mode"`
	LinkDiscoveryOperation         string            `json:"link_discovery_operation"`
	LinkDiscoveryProtocol          string            `json:"link_discovery_protocol"`
	ManagementMaximumMbit          int64             `json:"management_maximum_mbit"`
	ManagementReservationMbit      int64             `json:"management_reservation_mbit"`
	ManagementShareCount           int64             `json:"management_share_count"`
	ManagementShareLevel           string            `json:"management_share_level"`
	MaxMtu                         int64             `json:"max_mtu"`
	MulticastFilteringMode         string            `json:"multicast_filtering_mode"`
	Name                           string            `json:"name"`
	NetflowActiveFlowTimeout       int64             `json:"netflow_active_flow_timeout"`
	NetflowCollectorIpAddress      string            `json:"netflow_collector_ip_address"`
	NetflowCollectorPort           int64             `json:"netflow_collector_port"`
	NetflowEnabled                 bool              `json:"netflow_enabled"`
	NetflowIdleFlowTimeout         int64             `json:"netflow_idle_flow_timeout"`
	NetflowInternalFlowsOnly       bool              `json:"netflow_internal_flows_only"`
	NetflowObservationDomainId     int64             `json:"netflow_observation_domain_id"`
	NetflowSamplingRate            int64             `json:"netflow_sampling_rate"`
	NetworkResourceControlEnabled  bool              `json:"network_resource_control_enabled"`
	NetworkResourceControlVersion  string            `json:"network_resource_control_version"`
	NfsMaximumMbit                 int64             `json:"nfs_maximum_mbit"`
	NfsReservationMbit             int64             `json:"nfs_reservation_mbit"`
	NfsShareCount                  int64             `json:"nfs_share_count"`
	NfsShareLevel                  string            `json:"nfs_share_level"`
	NotifySwitches                 bool              `json:"notify_switches"`
	PortPrivateSecondaryVlanId     int64             `json:"port_private_secondary_vlan_id"`
	StandbyUplinks                 []string          `json:"standby_uplinks,omitempty"`
	Tags                           []string          `json:"tags,omitempty"`
	TeamingPolicy                  string            `json:"teaming_policy"`
	TxUplink                       bool              `json:"tx_uplink"`
	Uplinks                        []string          `json:"uplinks,omitempty"`
	VdpMaximumMbit                 int64             `json:"vdp_maximum_mbit"`
	VdpReservationMbit             int64             `json:"vdp_reservation_mbit"`
	VdpShareCount                  int64             `json:"vdp_share_count"`
	VdpShareLevel                  string            `json:"vdp_share_level"`
	Version                        string            `json:"version"`
	VirtualmachineMaximumMbit      int64             `json:"virtualmachine_maximum_mbit"`
	VirtualmachineReservationMbit  int64             `json:"virtualmachine_reservation_mbit"`
	VirtualmachineShareCount       int64             `json:"virtualmachine_share_count"`
	VirtualmachineShareLevel       string            `json:"virtualmachine_share_level"`
	VlanId                         int64             `json:"vlan_id"`
	VmotionMaximumMbit             int64             `json:"vmotion_maximum_mbit"`
	VmotionReservationMbit         int64             `json:"vmotion_reservation_mbit"`
	VmotionShareCount              int64             `json:"vmotion_share_count"`
	VmotionShareLevel              string            `json:"vmotion_share_level"`
	VsanMaximumMbit                int64             `json:"vsan_maximum_mbit"`
	VsanReservationMbit            int64             `json:"vsan_reservation_mbit"`
	VsanShareCount                 int64             `json:"vsan_share_count"`
	VsanShareLevel                 string            `json:"vsan_share_level"`
	Host                           Host              `json:"host"`
	PvlanMapping                   PvlanMapping      `json:"pvlan_mapping"`
	VlanRange                      VlanRange         `json:"vlan_range"`
}

type Host struct {
	Devices      []string `json:"devices,omitempty"`
	HostSystemId string   `json:"host_system_id"`
}

type PvlanMapping struct {
	PvlanType       string `json:"pvlan_type"`
	SecondaryVlanId int64  `json:"secondary_vlan_id"`
	PrimaryVlanId   int64  `json:"primary_vlan_id"`
}

type VlanRange struct {
	MinVlan int64 `json:"min_vlan"`
	MaxVlan int64 `json:"max_vlan"`
}

// A DistributedVirtualSwitchStatus defines the observed state of a DistributedVirtualSwitch
type DistributedVirtualSwitchStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DistributedVirtualSwitchObservation `json:"atProvider"`
}

// A DistributedVirtualSwitchObservation records the observed state of a DistributedVirtualSwitch
type DistributedVirtualSwitchObservation struct {
	ConfigVersion string `json:"config_version"`
}

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

// DistributedPortGroup is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DistributedPortGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DistributedPortGroupSpec   `json:"spec"`
	Status DistributedPortGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DistributedPortGroup contains a list of DistributedPortGroupList
type DistributedPortGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DistributedPortGroup `json:"items"`
}

// A DistributedPortGroupSpec defines the desired state of a DistributedPortGroup
type DistributedPortGroupSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DistributedPortGroupParameters `json:"forProvider"`
}

// A DistributedPortGroupParameters defines the desired state of a DistributedPortGroup
type DistributedPortGroupParameters struct {
	ActiveUplinks                      []string          `json:"active_uplinks,omitempty"`
	AllowForgedTransmits               bool              `json:"allow_forged_transmits"`
	AllowMacChanges                    bool              `json:"allow_mac_changes"`
	AllowPromiscuous                   bool              `json:"allow_promiscuous"`
	AutoExpand                         bool              `json:"auto_expand"`
	BlockAllPorts                      bool              `json:"block_all_ports"`
	BlockOverrideAllowed               bool              `json:"block_override_allowed"`
	CheckBeacon                        bool              `json:"check_beacon"`
	CustomAttributes                   map[string]string `json:"custom_attributes,omitempty"`
	Description                        string            `json:"description"`
	DirectpathGen2Allowed              bool              `json:"directpath_gen2_allowed"`
	DistributedVirtualSwitchUuid       string            `json:"distributed_virtual_switch_uuid"`
	EgressShapingAverageBandwidth      int64             `json:"egress_shaping_average_bandwidth"`
	EgressShapingBurstSize             int64             `json:"egress_shaping_burst_size"`
	EgressShapingEnabled               bool              `json:"egress_shaping_enabled"`
	EgressShapingPeakBandwidth         int64             `json:"egress_shaping_peak_bandwidth"`
	Failback                           bool              `json:"failback"`
	IngressShapingAverageBandwidth     int64             `json:"ingress_shaping_average_bandwidth"`
	IngressShapingBurstSize            int64             `json:"ingress_shaping_burst_size"`
	IngressShapingEnabled              bool              `json:"ingress_shaping_enabled"`
	IngressShapingPeakBandwidth        int64             `json:"ingress_shaping_peak_bandwidth"`
	LacpEnabled                        bool              `json:"lacp_enabled"`
	LacpMode                           string            `json:"lacp_mode"`
	LivePortMovingAllowed              bool              `json:"live_port_moving_allowed"`
	Name                               string            `json:"name"`
	NetflowEnabled                     bool              `json:"netflow_enabled"`
	NetflowOverrideAllowed             bool              `json:"netflow_override_allowed"`
	NetworkResourcePoolKey             string            `json:"network_resource_pool_key"`
	NetworkResourcePoolOverrideAllowed bool              `json:"network_resource_pool_override_allowed"`
	NotifySwitches                     bool              `json:"notify_switches"`
	NumberOfPorts                      int64             `json:"number_of_ports"`
	PortConfigResetAtDisconnect        bool              `json:"port_config_reset_at_disconnect"`
	PortNameFormat                     string            `json:"port_name_format"`
	PortPrivateSecondaryVlanId         int64             `json:"port_private_secondary_vlan_id"`
	SecurityPolicyOverrideAllowed      bool              `json:"security_policy_override_allowed"`
	ShapingOverrideAllowed             bool              `json:"shaping_override_allowed"`
	StandbyUplinks                     []string          `json:"standby_uplinks,omitempty"`
	Tags                               []string          `json:"tags,omitempty"`
	TeamingPolicy                      string            `json:"teaming_policy"`
	TrafficFilterOverrideAllowed       bool              `json:"traffic_filter_override_allowed"`
	TxUplink                           bool              `json:"tx_uplink"`
	Type                               string            `json:"type"`
	UplinkTeamingOverrideAllowed       bool              `json:"uplink_teaming_override_allowed"`
	VlanId                             int64             `json:"vlan_id"`
	VlanOverrideAllowed                bool              `json:"vlan_override_allowed"`
	VlanRange                          VlanRange         `json:"vlan_range"`
}

type VlanRange struct {
	MaxVlan int64 `json:"max_vlan"`
	MinVlan int64 `json:"min_vlan"`
}

// A DistributedPortGroupStatus defines the observed state of a DistributedPortGroup
type DistributedPortGroupStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DistributedPortGroupObservation `json:"atProvider"`
}

// A DistributedPortGroupObservation records the observed state of a DistributedPortGroup
type DistributedPortGroupObservation struct {
	ConfigVersion string `json:"config_version"`
	Key           string `json:"key"`
}
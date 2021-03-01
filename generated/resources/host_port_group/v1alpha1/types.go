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

// HostPortGroup is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type HostPortGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HostPortGroupSpec   `json:"spec"`
	Status HostPortGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HostPortGroup contains a list of HostPortGroupList
type HostPortGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HostPortGroup `json:"items"`
}

// A HostPortGroupSpec defines the desired state of a HostPortGroup
type HostPortGroupSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       HostPortGroupParameters `json:"forProvider"`
}

// A HostPortGroupParameters defines the desired state of a HostPortGroup
type HostPortGroupParameters struct {
	ActiveNics              []string `json:"active_nics,omitempty"`
	AllowForgedTransmits    bool     `json:"allow_forged_transmits"`
	AllowMacChanges         bool     `json:"allow_mac_changes"`
	AllowPromiscuous        bool     `json:"allow_promiscuous"`
	CheckBeacon             bool     `json:"check_beacon"`
	Failback                bool     `json:"failback"`
	HostSystemId            string   `json:"host_system_id"`
	Name                    string   `json:"name"`
	NotifySwitches          bool     `json:"notify_switches"`
	ShapingAverageBandwidth int64    `json:"shaping_average_bandwidth"`
	ShapingBurstSize        int64    `json:"shaping_burst_size"`
	ShapingEnabled          bool     `json:"shaping_enabled"`
	ShapingPeakBandwidth    int64    `json:"shaping_peak_bandwidth"`
	StandbyNics             []string `json:"standby_nics,omitempty"`
	TeamingPolicy           string   `json:"teaming_policy"`
	VirtualSwitchName       string   `json:"virtual_switch_name"`
	VlanId                  int64    `json:"vlan_id"`
}

// A HostPortGroupStatus defines the observed state of a HostPortGroup
type HostPortGroupStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          HostPortGroupObservation `json:"atProvider"`
}

// A HostPortGroupObservation records the observed state of a HostPortGroup
type HostPortGroupObservation struct {
	ComputedPolicy map[string]string `json:"computed_policy,omitempty"`
	Key            string            `json:"key"`
	Ports          []Ports           `json:"ports,omitempty"`
}

type Ports struct {
	Type         string   `json:"type"`
	Key          string   `json:"key"`
	MacAddresses []string `json:"mac_addresses,omitempty"`
}

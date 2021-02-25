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

// Vnic is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Vnic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VnicSpec   `json:"spec"`
	Status VnicStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Vnic contains a list of VnicList
type VnicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Vnic `json:"items"`
}

// A VnicSpec defines the desired state of a Vnic
type VnicSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       VnicParameters `json:"forProvider"`
}

// A VnicParameters defines the desired state of a Vnic
type VnicParameters struct {
	DistributedPortGroup  string `json:"distributed_port_group"`
	DistributedSwitchPort string `json:"distributed_switch_port"`
	Host                  string `json:"host"`
	Mac                   string `json:"mac"`
	Mtu                   int64  `json:"mtu"`
	Netstack              string `json:"netstack"`
	Portgroup             string `json:"portgroup"`
	Ipv4                  Ipv4   `json:"ipv4"`
	Ipv6                  Ipv6   `json:"ipv6"`
}

type Ipv4 struct {
	Dhcp    bool   `json:"dhcp"`
	Gw      string `json:"gw"`
	Ip      string `json:"ip"`
	Netmask string `json:"netmask"`
}

type Ipv6 struct {
	Autoconfig bool     `json:"autoconfig"`
	Dhcp       bool     `json:"dhcp"`
	Gw         string   `json:"gw"`
	Addresses  []string `json:"addresses,omitempty"`
}

// A VnicStatus defines the observed state of a Vnic
type VnicStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          VnicObservation `json:"atProvider"`
}

// A VnicObservation records the observed state of a Vnic
type VnicObservation struct{}

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

// DpmHostOverride is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DpmHostOverride struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DpmHostOverrideSpec   `json:"spec"`
	Status DpmHostOverrideStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DpmHostOverride contains a list of DpmHostOverrideList
type DpmHostOverrideList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DpmHostOverride `json:"items"`
}

// A DpmHostOverrideSpec defines the desired state of a DpmHostOverride
type DpmHostOverrideSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DpmHostOverrideParameters `json:"forProvider"`
}

// A DpmHostOverrideParameters defines the desired state of a DpmHostOverride
type DpmHostOverrideParameters struct {
	ComputeClusterId   string `json:"compute_cluster_id"`
	DpmAutomationLevel string `json:"dpm_automation_level"`
	DpmEnabled         bool   `json:"dpm_enabled"`
	HostSystemId       string `json:"host_system_id"`
}

// A DpmHostOverrideStatus defines the observed state of a DpmHostOverride
type DpmHostOverrideStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DpmHostOverrideObservation `json:"atProvider"`
}

// A DpmHostOverrideObservation records the observed state of a DpmHostOverride
type DpmHostOverrideObservation struct{}
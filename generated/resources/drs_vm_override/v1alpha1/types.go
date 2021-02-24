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

// DrsVmOverride is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DrsVmOverride struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DrsVmOverrideSpec   `json:"spec"`
	Status DrsVmOverrideStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DrsVmOverride contains a list of DrsVmOverrideList
type DrsVmOverrideList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DrsVmOverride `json:"items"`
}

// A DrsVmOverrideSpec defines the desired state of a DrsVmOverride
type DrsVmOverrideSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DrsVmOverrideParameters `json:"forProvider"`
}

// A DrsVmOverrideParameters defines the desired state of a DrsVmOverride
type DrsVmOverrideParameters struct {
	ComputeClusterId   string `json:"compute_cluster_id"`
	DrsAutomationLevel string `json:"drs_automation_level"`
	DrsEnabled         bool   `json:"drs_enabled"`
	VirtualMachineId   string `json:"virtual_machine_id"`
}

// A DrsVmOverrideStatus defines the observed state of a DrsVmOverride
type DrsVmOverrideStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DrsVmOverrideObservation `json:"atProvider"`
}

// A DrsVmOverrideObservation records the observed state of a DrsVmOverride
type DrsVmOverrideObservation struct{}
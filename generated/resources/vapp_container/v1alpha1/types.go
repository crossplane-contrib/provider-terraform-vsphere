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

	runtimev1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
)

// +kubebuilder:object:root=true

// VappContainer is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type VappContainer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VappContainerSpec   `json:"spec"`
	Status VappContainerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VappContainer contains a list of VappContainerList
type VappContainerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VappContainer `json:"items"`
}

// A VappContainerSpec defines the desired state of a VappContainer
type VappContainerSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  VappContainerParameters `json:"forProvider"`
}

// A VappContainerParameters defines the desired state of a VappContainer
type VappContainerParameters struct {
	CustomAttributes     map[string]string `json:"custom_attributes,omitempty"`
	MemoryLimit          int64             `json:"memory_limit"`
	MemoryReservation    int64             `json:"memory_reservation"`
	Name                 string            `json:"name"`
	Tags                 []string          `json:"tags,omitempty"`
	ParentResourcePoolId string            `json:"parent_resource_pool_id"`
	CpuLimit             int64             `json:"cpu_limit"`
	CpuReservation       int64             `json:"cpu_reservation"`
	CpuShareLevel        string            `json:"cpu_share_level"`
	CpuShares            int64             `json:"cpu_shares"`
	ParentFolderId       string            `json:"parent_folder_id"`
	CpuExpandable        bool              `json:"cpu_expandable"`
	MemoryExpandable     bool              `json:"memory_expandable"`
	MemoryShareLevel     string            `json:"memory_share_level"`
	MemoryShares         int64             `json:"memory_shares"`
}

// A VappContainerStatus defines the observed state of a VappContainer
type VappContainerStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     VappContainerObservation `json:"atProvider"`
}

// A VappContainerObservation records the observed state of a VappContainer
type VappContainerObservation struct{}
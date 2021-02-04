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

// ResourcePool is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ResourcePool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ResourcePoolSpec   `json:"spec"`
	Status ResourcePoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResourcePool contains a list of ResourcePoolList
type ResourcePoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourcePool `json:"items"`
}

// A ResourcePoolSpec defines the desired state of a ResourcePool
type ResourcePoolSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ResourcePoolParameters `json:"forProvider"`
}

// A ResourcePoolParameters defines the desired state of a ResourcePool
type ResourcePoolParameters struct {
	MemoryShareLevel     string            `json:"memory_share_level"`
	ParentResourcePoolId string            `json:"parent_resource_pool_id"`
	Tags                 []string          `json:"tags,omitempty"`
	CpuReservation       int64             `json:"cpu_reservation"`
	Name                 string            `json:"name"`
	CpuExpandable        bool              `json:"cpu_expandable"`
	CpuLimit             int64             `json:"cpu_limit"`
	CpuShares            int64             `json:"cpu_shares"`
	MemoryExpandable     bool              `json:"memory_expandable"`
	MemoryLimit          int64             `json:"memory_limit"`
	MemoryShares         int64             `json:"memory_shares"`
	CpuShareLevel        string            `json:"cpu_share_level"`
	CustomAttributes     map[string]string `json:"custom_attributes,omitempty"`
	MemoryReservation    int64             `json:"memory_reservation"`
}

// A ResourcePoolStatus defines the observed state of a ResourcePool
type ResourcePoolStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ResourcePoolObservation `json:"atProvider"`
}

// A ResourcePoolObservation records the observed state of a ResourcePool
type ResourcePoolObservation struct{}
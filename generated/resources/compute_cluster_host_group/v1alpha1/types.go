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

// ComputeClusterHostGroup is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ComputeClusterHostGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeClusterHostGroupSpec   `json:"spec"`
	Status ComputeClusterHostGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ComputeClusterHostGroup contains a list of ComputeClusterHostGroupList
type ComputeClusterHostGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeClusterHostGroup `json:"items"`
}

// A ComputeClusterHostGroupSpec defines the desired state of a ComputeClusterHostGroup
type ComputeClusterHostGroupSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ComputeClusterHostGroupParameters `json:"forProvider"`
}

// A ComputeClusterHostGroupParameters defines the desired state of a ComputeClusterHostGroup
type ComputeClusterHostGroupParameters struct {
	ComputeClusterId string   `json:"compute_cluster_id"`
	HostSystemIds    []string `json:"host_system_ids,omitempty"`
	Name             string   `json:"name"`
}

// A ComputeClusterHostGroupStatus defines the observed state of a ComputeClusterHostGroup
type ComputeClusterHostGroupStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ComputeClusterHostGroupObservation `json:"atProvider"`
}

// A ComputeClusterHostGroupObservation records the observed state of a ComputeClusterHostGroup
type ComputeClusterHostGroupObservation struct{}
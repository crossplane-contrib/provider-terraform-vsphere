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

// ComputeClusterVmGroup is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ComputeClusterVmGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeClusterVmGroupSpec   `json:"spec"`
	Status ComputeClusterVmGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ComputeClusterVmGroup contains a list of ComputeClusterVmGroupList
type ComputeClusterVmGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeClusterVmGroup `json:"items"`
}

// A ComputeClusterVmGroupSpec defines the desired state of a ComputeClusterVmGroup
type ComputeClusterVmGroupSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ComputeClusterVmGroupParameters `json:"forProvider"`
}

// A ComputeClusterVmGroupParameters defines the desired state of a ComputeClusterVmGroup
type ComputeClusterVmGroupParameters struct {
	ComputeClusterId  string   `json:"compute_cluster_id"`
	Name              string   `json:"name"`
	VirtualMachineIds []string `json:"virtual_machine_ids,omitempty"`
}

// A ComputeClusterVmGroupStatus defines the observed state of a ComputeClusterVmGroup
type ComputeClusterVmGroupStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ComputeClusterVmGroupObservation `json:"atProvider"`
}

// A ComputeClusterVmGroupObservation records the observed state of a ComputeClusterVmGroup
type ComputeClusterVmGroupObservation struct{}

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

// ComputeClusterVmAntiAffinityRule is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ComputeClusterVmAntiAffinityRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeClusterVmAntiAffinityRuleSpec   `json:"spec"`
	Status ComputeClusterVmAntiAffinityRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ComputeClusterVmAntiAffinityRule contains a list of ComputeClusterVmAntiAffinityRuleList
type ComputeClusterVmAntiAffinityRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeClusterVmAntiAffinityRule `json:"items"`
}

// A ComputeClusterVmAntiAffinityRuleSpec defines the desired state of a ComputeClusterVmAntiAffinityRule
type ComputeClusterVmAntiAffinityRuleSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ComputeClusterVmAntiAffinityRuleParameters `json:"forProvider"`
}

// A ComputeClusterVmAntiAffinityRuleParameters defines the desired state of a ComputeClusterVmAntiAffinityRule
type ComputeClusterVmAntiAffinityRuleParameters struct {
	ComputeClusterId  string   `json:"compute_cluster_id"`
	Enabled           bool     `json:"enabled"`
	Mandatory         bool     `json:"mandatory"`
	Name              string   `json:"name"`
	VirtualMachineIds []string `json:"virtual_machine_ids,omitempty"`
}

// A ComputeClusterVmAntiAffinityRuleStatus defines the observed state of a ComputeClusterVmAntiAffinityRule
type ComputeClusterVmAntiAffinityRuleStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ComputeClusterVmAntiAffinityRuleObservation `json:"atProvider"`
}

// A ComputeClusterVmAntiAffinityRuleObservation records the observed state of a ComputeClusterVmAntiAffinityRule
type ComputeClusterVmAntiAffinityRuleObservation struct{}
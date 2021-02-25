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

// ComputeClusterVmAffinityRule is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ComputeClusterVmAffinityRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeClusterVmAffinityRuleSpec   `json:"spec"`
	Status ComputeClusterVmAffinityRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ComputeClusterVmAffinityRule contains a list of ComputeClusterVmAffinityRuleList
type ComputeClusterVmAffinityRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeClusterVmAffinityRule `json:"items"`
}

// A ComputeClusterVmAffinityRuleSpec defines the desired state of a ComputeClusterVmAffinityRule
type ComputeClusterVmAffinityRuleSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ComputeClusterVmAffinityRuleParameters `json:"forProvider"`
}

// A ComputeClusterVmAffinityRuleParameters defines the desired state of a ComputeClusterVmAffinityRule
type ComputeClusterVmAffinityRuleParameters struct {
	ComputeClusterId  string   `json:"compute_cluster_id"`
	Enabled           bool     `json:"enabled"`
	Mandatory         bool     `json:"mandatory"`
	Name              string   `json:"name"`
	VirtualMachineIds []string `json:"virtual_machine_ids,omitempty"`
}

// A ComputeClusterVmAffinityRuleStatus defines the observed state of a ComputeClusterVmAffinityRule
type ComputeClusterVmAffinityRuleStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ComputeClusterVmAffinityRuleObservation `json:"atProvider"`
}

// A ComputeClusterVmAffinityRuleObservation records the observed state of a ComputeClusterVmAffinityRule
type ComputeClusterVmAffinityRuleObservation struct{}

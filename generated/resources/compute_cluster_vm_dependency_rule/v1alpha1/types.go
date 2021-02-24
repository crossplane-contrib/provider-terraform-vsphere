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

// ComputeClusterVmDependencyRule is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ComputeClusterVmDependencyRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeClusterVmDependencyRuleSpec   `json:"spec"`
	Status ComputeClusterVmDependencyRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ComputeClusterVmDependencyRule contains a list of ComputeClusterVmDependencyRuleList
type ComputeClusterVmDependencyRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeClusterVmDependencyRule `json:"items"`
}

// A ComputeClusterVmDependencyRuleSpec defines the desired state of a ComputeClusterVmDependencyRule
type ComputeClusterVmDependencyRuleSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ComputeClusterVmDependencyRuleParameters `json:"forProvider"`
}

// A ComputeClusterVmDependencyRuleParameters defines the desired state of a ComputeClusterVmDependencyRule
type ComputeClusterVmDependencyRuleParameters struct {
	ComputeClusterId      string `json:"compute_cluster_id"`
	DependencyVmGroupName string `json:"dependency_vm_group_name"`
	Enabled               bool   `json:"enabled"`
	Mandatory             bool   `json:"mandatory"`
	Name                  string `json:"name"`
	VmGroupName           string `json:"vm_group_name"`
}

// A ComputeClusterVmDependencyRuleStatus defines the observed state of a ComputeClusterVmDependencyRule
type ComputeClusterVmDependencyRuleStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ComputeClusterVmDependencyRuleObservation `json:"atProvider"`
}

// A ComputeClusterVmDependencyRuleObservation records the observed state of a ComputeClusterVmDependencyRule
type ComputeClusterVmDependencyRuleObservation struct{}
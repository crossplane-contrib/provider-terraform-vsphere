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

// ComputeClusterVmHostRule is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ComputeClusterVmHostRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeClusterVmHostRuleSpec   `json:"spec"`
	Status ComputeClusterVmHostRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ComputeClusterVmHostRule contains a list of ComputeClusterVmHostRuleList
type ComputeClusterVmHostRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeClusterVmHostRule `json:"items"`
}

// A ComputeClusterVmHostRuleSpec defines the desired state of a ComputeClusterVmHostRule
type ComputeClusterVmHostRuleSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ComputeClusterVmHostRuleParameters `json:"forProvider"`
}

// A ComputeClusterVmHostRuleParameters defines the desired state of a ComputeClusterVmHostRule
type ComputeClusterVmHostRuleParameters struct {
	AntiAffinityHostGroupName string `json:"anti_affinity_host_group_name"`
	ComputeClusterId          string `json:"compute_cluster_id"`
	Enabled                   bool   `json:"enabled"`
	Mandatory                 bool   `json:"mandatory"`
	Name                      string `json:"name"`
	VmGroupName               string `json:"vm_group_name"`
	AffinityHostGroupName     string `json:"affinity_host_group_name"`
}

// A ComputeClusterVmHostRuleStatus defines the observed state of a ComputeClusterVmHostRule
type ComputeClusterVmHostRuleStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ComputeClusterVmHostRuleObservation `json:"atProvider"`
}

// A ComputeClusterVmHostRuleObservation records the observed state of a ComputeClusterVmHostRule
type ComputeClusterVmHostRuleObservation struct{}
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

// VmStoragePolicy is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type VmStoragePolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VmStoragePolicySpec   `json:"spec"`
	Status VmStoragePolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VmStoragePolicy contains a list of VmStoragePolicyList
type VmStoragePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VmStoragePolicy `json:"items"`
}

// A VmStoragePolicySpec defines the desired state of a VmStoragePolicy
type VmStoragePolicySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       VmStoragePolicyParameters `json:"forProvider"`
}

// A VmStoragePolicyParameters defines the desired state of a VmStoragePolicy
type VmStoragePolicyParameters struct {
	Description string     `json:"description"`
	Name        string     `json:"name"`
	TagRules    []TagRules `json:"tag_rules"`
}

type TagRules struct {
	IncludeDatastoresWithTags bool     `json:"include_datastores_with_tags"`
	TagCategory               string   `json:"tag_category"`
	Tags                      []string `json:"tags,omitempty"`
}

// A VmStoragePolicyStatus defines the observed state of a VmStoragePolicy
type VmStoragePolicyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          VmStoragePolicyObservation `json:"atProvider"`
}

// A VmStoragePolicyObservation records the observed state of a VmStoragePolicy
type VmStoragePolicyObservation struct{}
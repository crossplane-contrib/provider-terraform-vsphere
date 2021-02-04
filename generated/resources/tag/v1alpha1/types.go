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

// Tag is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Tag struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TagSpec   `json:"spec"`
	Status TagStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Tag contains a list of TagList
type TagList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Tag `json:"items"`
}

// A TagSpec defines the desired state of a Tag
type TagSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  TagParameters `json:"forProvider"`
}

// A TagParameters defines the desired state of a Tag
type TagParameters struct {
	CategoryId  string `json:"category_id"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

// A TagStatus defines the observed state of a Tag
type TagStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     TagObservation `json:"atProvider"`
}

// A TagObservation records the observed state of a Tag
type TagObservation struct{}
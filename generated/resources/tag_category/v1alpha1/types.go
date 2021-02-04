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

// TagCategory is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type TagCategory struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TagCategorySpec   `json:"spec"`
	Status TagCategoryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TagCategory contains a list of TagCategoryList
type TagCategoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TagCategory `json:"items"`
}

// A TagCategorySpec defines the desired state of a TagCategory
type TagCategorySpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  TagCategoryParameters `json:"forProvider"`
}

// A TagCategoryParameters defines the desired state of a TagCategory
type TagCategoryParameters struct {
	Cardinality     string   `json:"cardinality"`
	Description     string   `json:"description"`
	Name            string   `json:"name"`
	AssociableTypes []string `json:"associable_types,omitempty"`
}

// A TagCategoryStatus defines the observed state of a TagCategory
type TagCategoryStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     TagCategoryObservation `json:"atProvider"`
}

// A TagCategoryObservation records the observed state of a TagCategory
type TagCategoryObservation struct{}
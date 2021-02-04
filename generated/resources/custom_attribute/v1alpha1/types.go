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

// CustomAttribute is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type CustomAttribute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CustomAttributeSpec   `json:"spec"`
	Status CustomAttributeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CustomAttribute contains a list of CustomAttributeList
type CustomAttributeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CustomAttribute `json:"items"`
}

// A CustomAttributeSpec defines the desired state of a CustomAttribute
type CustomAttributeSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  CustomAttributeParameters `json:"forProvider"`
}

// A CustomAttributeParameters defines the desired state of a CustomAttribute
type CustomAttributeParameters struct {
	ManagedObjectType string `json:"managed_object_type"`
	Name              string `json:"name"`
}

// A CustomAttributeStatus defines the observed state of a CustomAttribute
type CustomAttributeStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     CustomAttributeObservation `json:"atProvider"`
}

// A CustomAttributeObservation records the observed state of a CustomAttribute
type CustomAttributeObservation struct{}
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

// ContentLibraryItem is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ContentLibraryItem struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ContentLibraryItemSpec   `json:"spec"`
	Status ContentLibraryItemStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ContentLibraryItem contains a list of ContentLibraryItemList
type ContentLibraryItemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContentLibraryItem `json:"items"`
}

// A ContentLibraryItemSpec defines the desired state of a ContentLibraryItem
type ContentLibraryItemSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ContentLibraryItemParameters `json:"forProvider"`
}

// A ContentLibraryItemParameters defines the desired state of a ContentLibraryItem
type ContentLibraryItemParameters struct {
	Description string `json:"description"`
	FileUrl     string `json:"file_url"`
	LibraryId   string `json:"library_id"`
	Name        string `json:"name"`
	SourceUuid  string `json:"source_uuid"`
	Type        string `json:"type"`
}

// A ContentLibraryItemStatus defines the observed state of a ContentLibraryItem
type ContentLibraryItemStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ContentLibraryItemObservation `json:"atProvider"`
}

// A ContentLibraryItemObservation records the observed state of a ContentLibraryItem
type ContentLibraryItemObservation struct{}

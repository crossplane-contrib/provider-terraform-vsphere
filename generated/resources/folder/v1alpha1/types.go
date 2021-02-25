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

// Folder is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Folder struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FolderSpec   `json:"spec"`
	Status FolderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Folder contains a list of FolderList
type FolderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Folder `json:"items"`
}

// A FolderSpec defines the desired state of a Folder
type FolderSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       FolderParameters `json:"forProvider"`
}

// A FolderParameters defines the desired state of a Folder
type FolderParameters struct {
	CustomAttributes map[string]string `json:"custom_attributes,omitempty"`
	DatacenterId     string            `json:"datacenter_id"`
	Path             string            `json:"path"`
	Tags             []string          `json:"tags,omitempty"`
	Type             string            `json:"type"`
}

// A FolderStatus defines the observed state of a Folder
type FolderStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          FolderObservation `json:"atProvider"`
}

// A FolderObservation records the observed state of a Folder
type FolderObservation struct{}

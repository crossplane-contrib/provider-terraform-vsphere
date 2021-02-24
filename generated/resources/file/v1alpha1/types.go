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

// File is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type File struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FileSpec   `json:"spec"`
	Status FileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// File contains a list of FileList
type FileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []File `json:"items"`
}

// A FileSpec defines the desired state of a File
type FileSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       FileParameters `json:"forProvider"`
}

// A FileParameters defines the desired state of a File
type FileParameters struct {
	CreateDirectories bool   `json:"create_directories"`
	Datacenter        string `json:"datacenter"`
	Datastore         string `json:"datastore"`
	DestinationFile   string `json:"destination_file"`
	SourceDatacenter  string `json:"source_datacenter"`
	SourceDatastore   string `json:"source_datastore"`
	SourceFile        string `json:"source_file"`
}

// A FileStatus defines the observed state of a File
type FileStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          FileObservation `json:"atProvider"`
}

// A FileObservation records the observed state of a File
type FileObservation struct{}
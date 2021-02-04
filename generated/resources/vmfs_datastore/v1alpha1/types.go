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

// VmfsDatastore is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type VmfsDatastore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VmfsDatastoreSpec   `json:"spec"`
	Status VmfsDatastoreStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VmfsDatastore contains a list of VmfsDatastoreList
type VmfsDatastoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VmfsDatastore `json:"items"`
}

// A VmfsDatastoreSpec defines the desired state of a VmfsDatastore
type VmfsDatastoreSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  VmfsDatastoreParameters `json:"forProvider"`
}

// A VmfsDatastoreParameters defines the desired state of a VmfsDatastore
type VmfsDatastoreParameters struct {
	CustomAttributes   map[string]string `json:"custom_attributes,omitempty"`
	Folder             string            `json:"folder"`
	Tags               []string          `json:"tags,omitempty"`
	DatastoreClusterId string            `json:"datastore_cluster_id"`
	Disks              []string          `json:"disks,omitempty"`
	HostSystemId       string            `json:"host_system_id"`
	Name               string            `json:"name"`
}

// A VmfsDatastoreStatus defines the observed state of a VmfsDatastore
type VmfsDatastoreStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     VmfsDatastoreObservation `json:"atProvider"`
}

// A VmfsDatastoreObservation records the observed state of a VmfsDatastore
type VmfsDatastoreObservation struct {
	Capacity           int64  `json:"capacity"`
	FreeSpace          int64  `json:"free_space"`
	UncommittedSpace   int64  `json:"uncommitted_space"`
	Accessible         bool   `json:"accessible"`
	MaintenanceMode    string `json:"maintenance_mode"`
	MultipleHostAccess bool   `json:"multiple_host_access"`
	Url                string `json:"url"`
}
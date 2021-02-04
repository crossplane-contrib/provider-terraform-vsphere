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

// NasDatastore is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type NasDatastore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NasDatastoreSpec   `json:"spec"`
	Status NasDatastoreStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NasDatastore contains a list of NasDatastoreList
type NasDatastoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NasDatastore `json:"items"`
}

// A NasDatastoreSpec defines the desired state of a NasDatastore
type NasDatastoreSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  NasDatastoreParameters `json:"forProvider"`
}

// A NasDatastoreParameters defines the desired state of a NasDatastore
type NasDatastoreParameters struct {
	AccessMode         string            `json:"access_mode"`
	CustomAttributes   map[string]string `json:"custom_attributes,omitempty"`
	SecurityType       string            `json:"security_type"`
	Tags               []string          `json:"tags,omitempty"`
	Folder             string            `json:"folder"`
	RemoteHosts        []string          `json:"remote_hosts,omitempty"`
	RemotePath         string            `json:"remote_path"`
	DatastoreClusterId string            `json:"datastore_cluster_id"`
	HostSystemIds      []string          `json:"host_system_ids,omitempty"`
	Name               string            `json:"name"`
	Type               string            `json:"type"`
}

// A NasDatastoreStatus defines the observed state of a NasDatastore
type NasDatastoreStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     NasDatastoreObservation `json:"atProvider"`
}

// A NasDatastoreObservation records the observed state of a NasDatastore
type NasDatastoreObservation struct {
	Capacity           int64  `json:"capacity"`
	ProtocolEndpoint   string `json:"protocol_endpoint"`
	Accessible         bool   `json:"accessible"`
	FreeSpace          int64  `json:"free_space"`
	Url                string `json:"url"`
	MaintenanceMode    string `json:"maintenance_mode"`
	MultipleHostAccess bool   `json:"multiple_host_access"`
	UncommittedSpace   int64  `json:"uncommitted_space"`
}
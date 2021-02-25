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

// Datacenter is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Datacenter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DatacenterSpec   `json:"spec"`
	Status DatacenterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Datacenter contains a list of DatacenterList
type DatacenterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Datacenter `json:"items"`
}

// A DatacenterSpec defines the desired state of a Datacenter
type DatacenterSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DatacenterParameters `json:"forProvider"`
}

// A DatacenterParameters defines the desired state of a Datacenter
type DatacenterParameters struct {
	CustomAttributes map[string]string `json:"custom_attributes,omitempty"`
	Folder           string            `json:"folder"`
	Name             string            `json:"name"`
	Tags             []string          `json:"tags,omitempty"`
}

// A DatacenterStatus defines the observed state of a Datacenter
type DatacenterStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DatacenterObservation `json:"atProvider"`
}

// A DatacenterObservation records the observed state of a Datacenter
type DatacenterObservation struct {
	Moid string `json:"moid"`
}

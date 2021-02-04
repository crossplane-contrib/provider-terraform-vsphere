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

// Role is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Role struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RoleSpec   `json:"spec"`
	Status RoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Role contains a list of RoleList
type RoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Role `json:"items"`
}

// A RoleSpec defines the desired state of a Role
type RoleSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  RoleParameters `json:"forProvider"`
}

// A RoleParameters defines the desired state of a Role
type RoleParameters struct {
	Name           string   `json:"name"`
	RolePrivileges []string `json:"role_privileges,omitempty"`
}

// A RoleStatus defines the observed state of a Role
type RoleStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     RoleObservation `json:"atProvider"`
}

// A RoleObservation records the observed state of a Role
type RoleObservation struct {
	Label string `json:"label"`
}
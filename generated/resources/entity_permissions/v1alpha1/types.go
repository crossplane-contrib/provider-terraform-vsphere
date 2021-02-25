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

// EntityPermissions is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type EntityPermissions struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EntityPermissionsSpec   `json:"spec"`
	Status EntityPermissionsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EntityPermissions contains a list of EntityPermissionsList
type EntityPermissionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EntityPermissions `json:"items"`
}

// A EntityPermissionsSpec defines the desired state of a EntityPermissions
type EntityPermissionsSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       EntityPermissionsParameters `json:"forProvider"`
}

// A EntityPermissionsParameters defines the desired state of a EntityPermissions
type EntityPermissionsParameters struct {
	EntityId    string        `json:"entity_id"`
	EntityType  string        `json:"entity_type"`
	Permissions []Permissions `json:"permissions"`
}

type Permissions struct {
	IsGroup     bool   `json:"is_group"`
	Propagate   bool   `json:"propagate"`
	RoleId      string `json:"role_id"`
	UserOrGroup string `json:"user_or_group"`
}

// A EntityPermissionsStatus defines the observed state of a EntityPermissions
type EntityPermissionsStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          EntityPermissionsObservation `json:"atProvider"`
}

// A EntityPermissionsObservation records the observed state of a EntityPermissions
type EntityPermissionsObservation struct{}

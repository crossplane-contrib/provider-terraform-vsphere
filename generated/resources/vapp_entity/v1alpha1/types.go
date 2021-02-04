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

// VappEntity is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type VappEntity struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VappEntitySpec   `json:"spec"`
	Status VappEntityStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VappEntity contains a list of VappEntityList
type VappEntityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VappEntity `json:"items"`
}

// A VappEntitySpec defines the desired state of a VappEntity
type VappEntitySpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  VappEntityParameters `json:"forProvider"`
}

// A VappEntityParameters defines the desired state of a VappEntity
type VappEntityParameters struct {
	StopAction       string            `json:"stop_action"`
	WaitForGuest     bool              `json:"wait_for_guest"`
	ContainerId      string            `json:"container_id"`
	CustomAttributes map[string]string `json:"custom_attributes,omitempty"`
	StartAction      string            `json:"start_action"`
	TargetId         string            `json:"target_id"`
	StartDelay       int64             `json:"start_delay"`
	StartOrder       int64             `json:"start_order"`
	StopDelay        int64             `json:"stop_delay"`
	Tags             []string          `json:"tags,omitempty"`
}

// A VappEntityStatus defines the observed state of a VappEntity
type VappEntityStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     VappEntityObservation `json:"atProvider"`
}

// A VappEntityObservation records the observed state of a VappEntity
type VappEntityObservation struct{}
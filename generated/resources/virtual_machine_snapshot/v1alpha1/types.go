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

// VirtualMachineSnapshot is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type VirtualMachineSnapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualMachineSnapshotSpec   `json:"spec"`
	Status VirtualMachineSnapshotStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualMachineSnapshot contains a list of VirtualMachineSnapshotList
type VirtualMachineSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualMachineSnapshot `json:"items"`
}

// A VirtualMachineSnapshotSpec defines the desired state of a VirtualMachineSnapshot
type VirtualMachineSnapshotSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       VirtualMachineSnapshotParameters `json:"forProvider"`
}

// A VirtualMachineSnapshotParameters defines the desired state of a VirtualMachineSnapshot
type VirtualMachineSnapshotParameters struct {
	Consolidate        bool   `json:"consolidate"`
	Description        string `json:"description"`
	Memory             bool   `json:"memory"`
	Quiesce            bool   `json:"quiesce"`
	RemoveChildren     bool   `json:"remove_children"`
	SnapshotName       string `json:"snapshot_name"`
	VirtualMachineUuid string `json:"virtual_machine_uuid"`
}

// A VirtualMachineSnapshotStatus defines the observed state of a VirtualMachineSnapshot
type VirtualMachineSnapshotStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          VirtualMachineSnapshotObservation `json:"atProvider"`
}

// A VirtualMachineSnapshotObservation records the observed state of a VirtualMachineSnapshot
type VirtualMachineSnapshotObservation struct{}

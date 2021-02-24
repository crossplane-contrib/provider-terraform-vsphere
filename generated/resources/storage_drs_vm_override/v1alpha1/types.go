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

// StorageDrsVmOverride is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type StorageDrsVmOverride struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StorageDrsVmOverrideSpec   `json:"spec"`
	Status StorageDrsVmOverrideStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StorageDrsVmOverride contains a list of StorageDrsVmOverrideList
type StorageDrsVmOverrideList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageDrsVmOverride `json:"items"`
}

// A StorageDrsVmOverrideSpec defines the desired state of a StorageDrsVmOverride
type StorageDrsVmOverrideSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       StorageDrsVmOverrideParameters `json:"forProvider"`
}

// A StorageDrsVmOverrideParameters defines the desired state of a StorageDrsVmOverride
type StorageDrsVmOverrideParameters struct {
	DatastoreClusterId  string `json:"datastore_cluster_id"`
	SdrsAutomationLevel string `json:"sdrs_automation_level"`
	SdrsEnabled         string `json:"sdrs_enabled"`
	SdrsIntraVmAffinity string `json:"sdrs_intra_vm_affinity"`
	VirtualMachineId    string `json:"virtual_machine_id"`
}

// A StorageDrsVmOverrideStatus defines the observed state of a StorageDrsVmOverride
type StorageDrsVmOverrideStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          StorageDrsVmOverrideObservation `json:"atProvider"`
}

// A StorageDrsVmOverrideObservation records the observed state of a StorageDrsVmOverride
type StorageDrsVmOverrideObservation struct{}
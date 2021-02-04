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

// License is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type License struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LicenseSpec   `json:"spec"`
	Status LicenseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// License contains a list of LicenseList
type LicenseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []License `json:"items"`
}

// A LicenseSpec defines the desired state of a License
type LicenseSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  LicenseParameters `json:"forProvider"`
}

// A LicenseParameters defines the desired state of a License
type LicenseParameters struct {
	Labels     map[string]string `json:"labels,omitempty"`
	LicenseKey string            `json:"license_key"`
}

// A LicenseStatus defines the observed state of a License
type LicenseStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     LicenseObservation `json:"atProvider"`
}

// A LicenseObservation records the observed state of a License
type LicenseObservation struct {
	Used       int64  `json:"used"`
	EditionKey string `json:"edition_key"`
	Name       string `json:"name"`
	Total      int64  `json:"total"`
}
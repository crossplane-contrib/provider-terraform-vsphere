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

// ContentLibrary is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ContentLibrary struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ContentLibrarySpec   `json:"spec"`
	Status ContentLibraryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ContentLibrary contains a list of ContentLibraryList
type ContentLibraryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContentLibrary `json:"items"`
}

// A ContentLibrarySpec defines the desired state of a ContentLibrary
type ContentLibrarySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ContentLibraryParameters `json:"forProvider"`
}

// A ContentLibraryParameters defines the desired state of a ContentLibrary
type ContentLibraryParameters struct {
	Description    string       `json:"description"`
	Name           string       `json:"name"`
	StorageBacking []string     `json:"storage_backing,omitempty"`
	Publication    Publication  `json:"publication"`
	Subscription   Subscription `json:"subscription"`
}

type Publication struct {
	AuthenticationMethod string `json:"authentication_method"`
	Password             string `json:"password"`
	PublishUrl           string `json:"publish_url"`
	Published            bool   `json:"published"`
	Username             string `json:"username"`
}

type Subscription struct {
	AuthenticationMethod string `json:"authentication_method"`
	AutomaticSync        bool   `json:"automatic_sync"`
	OnDemand             bool   `json:"on_demand"`
	Password             string `json:"password"`
	SubscriptionUrl      string `json:"subscription_url"`
	Username             string `json:"username"`
}

// A ContentLibraryStatus defines the observed state of a ContentLibrary
type ContentLibraryStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ContentLibraryObservation `json:"atProvider"`
}

// A ContentLibraryObservation records the observed state of a ContentLibrary
type ContentLibraryObservation struct{}

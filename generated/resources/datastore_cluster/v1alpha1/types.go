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

// DatastoreCluster is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DatastoreCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DatastoreClusterSpec   `json:"spec"`
	Status DatastoreClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatastoreCluster contains a list of DatastoreClusterList
type DatastoreClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatastoreCluster `json:"items"`
}

// A DatastoreClusterSpec defines the desired state of a DatastoreCluster
type DatastoreClusterSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DatastoreClusterParameters `json:"forProvider"`
}

// A DatastoreClusterParameters defines the desired state of a DatastoreCluster
type DatastoreClusterParameters struct {
	SdrsDefaultIntraVmAffinity           bool              `json:"sdrs_default_intra_vm_affinity"`
	SdrsIoLoadImbalanceThreshold         int64             `json:"sdrs_io_load_imbalance_threshold"`
	Tags                                 []string          `json:"tags,omitempty"`
	SdrsIoReservablePercentThreshold     int64             `json:"sdrs_io_reservable_percent_threshold"`
	SdrsIoReservableThresholdMode        string            `json:"sdrs_io_reservable_threshold_mode"`
	SdrsRuleEnforcementAutomationLevel   string            `json:"sdrs_rule_enforcement_automation_level"`
	CustomAttributes                     map[string]string `json:"custom_attributes,omitempty"`
	DatacenterId                         string            `json:"datacenter_id"`
	SdrsFreeSpaceUtilizationDifference   int64             `json:"sdrs_free_space_utilization_difference"`
	SdrsIoLatencyThreshold               int64             `json:"sdrs_io_latency_threshold"`
	SdrsSpaceBalanceAutomationLevel      string            `json:"sdrs_space_balance_automation_level"`
	SdrsSpaceUtilizationThreshold        int64             `json:"sdrs_space_utilization_threshold"`
	SdrsVmEvacuationAutomationLevel      string            `json:"sdrs_vm_evacuation_automation_level"`
	SdrsPolicyEnforcementAutomationLevel string            `json:"sdrs_policy_enforcement_automation_level"`
	Folder                               string            `json:"folder"`
	SdrsAutomationLevel                  string            `json:"sdrs_automation_level"`
	SdrsFreeSpaceThreshold               int64             `json:"sdrs_free_space_threshold"`
	SdrsIoBalanceAutomationLevel         string            `json:"sdrs_io_balance_automation_level"`
	SdrsLoadBalanceInterval              int64             `json:"sdrs_load_balance_interval"`
	SdrsIoReservableIopsThreshold        int64             `json:"sdrs_io_reservable_iops_threshold"`
	Name                                 string            `json:"name"`
	SdrsAdvancedOptions                  map[string]string `json:"sdrs_advanced_options,omitempty"`
	SdrsEnabled                          bool              `json:"sdrs_enabled"`
	SdrsFreeSpaceThresholdMode           string            `json:"sdrs_free_space_threshold_mode"`
	SdrsIoLoadBalanceEnabled             bool              `json:"sdrs_io_load_balance_enabled"`
}

// A DatastoreClusterStatus defines the observed state of a DatastoreCluster
type DatastoreClusterStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DatastoreClusterObservation `json:"atProvider"`
}

// A DatastoreClusterObservation records the observed state of a DatastoreCluster
type DatastoreClusterObservation struct{}
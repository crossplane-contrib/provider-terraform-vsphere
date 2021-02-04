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

// HaVmOverride is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type HaVmOverride struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HaVmOverrideSpec   `json:"spec"`
	Status HaVmOverrideStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HaVmOverride contains a list of HaVmOverrideList
type HaVmOverrideList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HaVmOverride `json:"items"`
}

// A HaVmOverrideSpec defines the desired state of a HaVmOverride
type HaVmOverrideSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  HaVmOverrideParameters `json:"forProvider"`
}

// A HaVmOverrideParameters defines the desired state of a HaVmOverride
type HaVmOverrideParameters struct {
	HaVmMaximumResets                int64  `json:"ha_vm_maximum_resets"`
	ComputeClusterId                 string `json:"compute_cluster_id"`
	HaDatastoreApdResponseDelay      int64  `json:"ha_datastore_apd_response_delay"`
	HaDatastoreApdRecoveryAction     string `json:"ha_datastore_apd_recovery_action"`
	HaVmMaximumFailureWindow         int64  `json:"ha_vm_maximum_failure_window"`
	HaVmMonitoringUseClusterDefaults bool   `json:"ha_vm_monitoring_use_cluster_defaults"`
	HaVmRestartTimeout               int64  `json:"ha_vm_restart_timeout"`
	VirtualMachineId                 string `json:"virtual_machine_id"`
	HaDatastorePdlResponse           string `json:"ha_datastore_pdl_response"`
	HaHostIsolationResponse          string `json:"ha_host_isolation_response"`
	HaVmFailureInterval              int64  `json:"ha_vm_failure_interval"`
	HaVmMinimumUptime                int64  `json:"ha_vm_minimum_uptime"`
	HaVmMonitoring                   string `json:"ha_vm_monitoring"`
	HaVmRestartPriority              string `json:"ha_vm_restart_priority"`
	HaDatastoreApdResponse           string `json:"ha_datastore_apd_response"`
}

// A HaVmOverrideStatus defines the observed state of a HaVmOverride
type HaVmOverrideStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     HaVmOverrideObservation `json:"atProvider"`
}

// A HaVmOverrideObservation records the observed state of a HaVmOverride
type HaVmOverrideObservation struct{}
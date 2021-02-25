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
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       HaVmOverrideParameters `json:"forProvider"`
}

// A HaVmOverrideParameters defines the desired state of a HaVmOverride
type HaVmOverrideParameters struct {
	ComputeClusterId                 string `json:"compute_cluster_id"`
	HaDatastoreApdRecoveryAction     string `json:"ha_datastore_apd_recovery_action"`
	HaDatastoreApdResponse           string `json:"ha_datastore_apd_response"`
	HaDatastoreApdResponseDelay      int64  `json:"ha_datastore_apd_response_delay"`
	HaDatastorePdlResponse           string `json:"ha_datastore_pdl_response"`
	HaHostIsolationResponse          string `json:"ha_host_isolation_response"`
	HaVmFailureInterval              int64  `json:"ha_vm_failure_interval"`
	HaVmMaximumFailureWindow         int64  `json:"ha_vm_maximum_failure_window"`
	HaVmMaximumResets                int64  `json:"ha_vm_maximum_resets"`
	HaVmMinimumUptime                int64  `json:"ha_vm_minimum_uptime"`
	HaVmMonitoring                   string `json:"ha_vm_monitoring"`
	HaVmMonitoringUseClusterDefaults bool   `json:"ha_vm_monitoring_use_cluster_defaults"`
	HaVmRestartPriority              string `json:"ha_vm_restart_priority"`
	HaVmRestartTimeout               int64  `json:"ha_vm_restart_timeout"`
	VirtualMachineId                 string `json:"virtual_machine_id"`
}

// A HaVmOverrideStatus defines the observed state of a HaVmOverride
type HaVmOverrideStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          HaVmOverrideObservation `json:"atProvider"`
}

// A HaVmOverrideObservation records the observed state of a HaVmOverride
type HaVmOverrideObservation struct{}

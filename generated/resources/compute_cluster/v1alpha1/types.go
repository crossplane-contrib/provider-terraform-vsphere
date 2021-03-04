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

// ComputeCluster is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ComputeCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeClusterSpec   `json:"spec"`
	Status ComputeClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ComputeCluster contains a list of ComputeClusterList
type ComputeClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeCluster `json:"items"`
}

// A ComputeClusterSpec defines the desired state of a ComputeCluster
type ComputeClusterSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ComputeClusterParameters `json:"forProvider"`
}

// A ComputeClusterParameters defines the desired state of a ComputeCluster
type ComputeClusterParameters struct {
	CustomAttributes                                map[string]string `json:"custom_attributes,omitempty"`
	DatacenterId                                    string            `json:"datacenter_id"`
	DpmAutomationLevel                              string            `json:"dpm_automation_level"`
	DpmEnabled                                      bool              `json:"dpm_enabled"`
	DpmThreshold                                    int64             `json:"dpm_threshold"`
	DrsAdvancedOptions                              map[string]string `json:"drs_advanced_options,omitempty"`
	DrsAutomationLevel                              string            `json:"drs_automation_level"`
	DrsEnablePredictiveDrs                          bool              `json:"drs_enable_predictive_drs"`
	DrsEnableVmOverrides                            bool              `json:"drs_enable_vm_overrides"`
	DrsEnabled                                      bool              `json:"drs_enabled"`
	DrsMigrationThreshold                           int64             `json:"drs_migration_threshold"`
	Folder                                          string            `json:"folder"`
	ForceEvacuateOnDestroy                          bool              `json:"force_evacuate_on_destroy"`
	HaAdmissionControlFailoverHostSystemIds         []string          `json:"ha_admission_control_failover_host_system_ids,omitempty"`
	HaAdmissionControlHostFailureTolerance          int64             `json:"ha_admission_control_host_failure_tolerance"`
	HaAdmissionControlPerformanceTolerance          int64             `json:"ha_admission_control_performance_tolerance"`
	HaAdmissionControlPolicy                        string            `json:"ha_admission_control_policy"`
	HaAdmissionControlResourcePercentageAutoCompute bool              `json:"ha_admission_control_resource_percentage_auto_compute"`
	HaAdmissionControlResourcePercentageCpu         int64             `json:"ha_admission_control_resource_percentage_cpu"`
	HaAdmissionControlResourcePercentageMemory      int64             `json:"ha_admission_control_resource_percentage_memory"`
	HaAdmissionControlSlotPolicyExplicitCpu         int64             `json:"ha_admission_control_slot_policy_explicit_cpu"`
	HaAdmissionControlSlotPolicyExplicitMemory      int64             `json:"ha_admission_control_slot_policy_explicit_memory"`
	HaAdmissionControlSlotPolicyUseExplicitSize     bool              `json:"ha_admission_control_slot_policy_use_explicit_size"`
	HaAdvancedOptions                               map[string]string `json:"ha_advanced_options,omitempty"`
	HaDatastoreApdRecoveryAction                    string            `json:"ha_datastore_apd_recovery_action"`
	HaDatastoreApdResponse                          string            `json:"ha_datastore_apd_response"`
	HaDatastoreApdResponseDelay                     int64             `json:"ha_datastore_apd_response_delay"`
	HaDatastorePdlResponse                          string            `json:"ha_datastore_pdl_response"`
	HaEnabled                                       bool              `json:"ha_enabled"`
	HaHeartbeatDatastoreIds                         []string          `json:"ha_heartbeat_datastore_ids,omitempty"`
	HaHeartbeatDatastorePolicy                      string            `json:"ha_heartbeat_datastore_policy"`
	HaHostIsolationResponse                         string            `json:"ha_host_isolation_response"`
	HaHostMonitoring                                string            `json:"ha_host_monitoring"`
	HaVmComponentProtection                         string            `json:"ha_vm_component_protection"`
	HaVmDependencyRestartCondition                  string            `json:"ha_vm_dependency_restart_condition"`
	HaVmFailureInterval                             int64             `json:"ha_vm_failure_interval"`
	HaVmMaximumFailureWindow                        int64             `json:"ha_vm_maximum_failure_window"`
	HaVmMaximumResets                               int64             `json:"ha_vm_maximum_resets"`
	HaVmMinimumUptime                               int64             `json:"ha_vm_minimum_uptime"`
	HaVmMonitoring                                  string            `json:"ha_vm_monitoring"`
	HaVmRestartAdditionalDelay                      int64             `json:"ha_vm_restart_additional_delay"`
	HaVmRestartPriority                             string            `json:"ha_vm_restart_priority"`
	HaVmRestartTimeout                              int64             `json:"ha_vm_restart_timeout"`
	HostClusterExitTimeout                          int64             `json:"host_cluster_exit_timeout"`
	HostManaged                                     bool              `json:"host_managed"`
	HostSystemIds                                   []string          `json:"host_system_ids,omitempty"`
	Name                                            string            `json:"name"`
	ProactiveHaAutomationLevel                      string            `json:"proactive_ha_automation_level"`
	ProactiveHaEnabled                              bool              `json:"proactive_ha_enabled"`
	ProactiveHaModerateRemediation                  string            `json:"proactive_ha_moderate_remediation"`
	ProactiveHaProviderIds                          []string          `json:"proactive_ha_provider_ids,omitempty"`
	ProactiveHaSevereRemediation                    string            `json:"proactive_ha_severe_remediation"`
	Tags                                            []string          `json:"tags,omitempty"`
	VsanEnabled                                     bool              `json:"vsan_enabled"`
	VsanDiskGroup                                   VsanDiskGroup     `json:"vsan_disk_group"`
}

type VsanDiskGroup struct {
	Storage []string `json:"storage,omitempty"`
	Cache   string   `json:"cache"`
}

// A ComputeClusterStatus defines the observed state of a ComputeCluster
type ComputeClusterStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ComputeClusterObservation `json:"atProvider"`
}

// A ComputeClusterObservation records the observed state of a ComputeCluster
type ComputeClusterObservation struct {
	ResourcePoolId string `json:"resource_pool_id"`
}

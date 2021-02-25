package generated

import (
	compute_cluster "github.com/crossplane-contrib/provider-terraform-vsphere/generated/resources/compute_cluster/v1alpha1"
	compute_cluster_host_group "github.com/crossplane-contrib/provider-terraform-vsphere/generated/resources/compute_cluster_host_group/v1alpha1"
	compute_cluster_vm_affinity_rule "github.com/crossplane-contrib/provider-terraform-vsphere/generated/resources/compute_cluster_vm_affinity_rule/v1alpha1"
	compute_cluster_vm_anti_affinity_rule "github.com/crossplane-contrib/provider-terraform-vsphere/generated/resources/compute_cluster_vm_anti_affinity_rule/v1alpha1"
	compute_cluster_vm_dependency_rule "github.com/crossplane-contrib/provider-terraform-vsphere/generated/resources/compute_cluster_vm_dependency_rule/v1alpha1"
	compute_cluster_vm_group "github.com/crossplane-contrib/provider-terraform-vsphere/generated/resources/compute_cluster_vm_group/v1alpha1"
	compute_cluster_vm_host_rule "github.com/crossplane-contrib/provider-terraform-vsphere/generated/resources/compute_cluster_vm_host_rule/v1alpha1"
	content_library "github.com/crossplane-contrib/provider-terraform-vsphere/generated/resources/content_library/v1alpha1"
	content_library_item "github.com/crossplane-contrib/provider-terraform-vsphere/generated/resources/content_library_item/v1alpha1"
	custom_attribute "github.com/crossplane-contrib/provider-terraform-vsphere/generated/resources/custom_attribute/v1alpha1"
	datacenter "github.com/crossplane-contrib/provider-terraform-vsphere/generated/resources/datacenter/v1alpha1"
	datastore_cluster "github.com/crossplane-contrib/provider-terraform-vsphere/generated/resources/datastore_cluster/v1alpha1"
	datastore_cluster_vm_anti_affinity_rule "github.com/crossplane-contrib/provider-terraform-vsphere/generated/resources/datastore_cluster_vm_anti_affinity_rule/v1alpha1"
	distributed_port_group "github.com/crossplane-contrib/provider-terraform-vsphere/generated/resources/distributed_port_group/v1alpha1"
	distributed_virtual_switch "github.com/crossplane-contrib/provider-terraform-vsphere/generated/resources/distributed_virtual_switch/v1alpha1"
	dpm_host_override "github.com/crossplane-contrib/provider-terraform-vsphere/generated/resources/dpm_host_override/v1alpha1"
	drs_vm_override "github.com/crossplane-contrib/provider-terraform-vsphere/generated/resources/drs_vm_override/v1alpha1"
	entity_permissions "github.com/crossplane-contrib/provider-terraform-vsphere/generated/resources/entity_permissions/v1alpha1"
	file "github.com/crossplane-contrib/provider-terraform-vsphere/generated/resources/file/v1alpha1"
	folder "github.com/crossplane-contrib/provider-terraform-vsphere/generated/resources/folder/v1alpha1"
	ha_vm_override "github.com/crossplane-contrib/provider-terraform-vsphere/generated/resources/ha_vm_override/v1alpha1"
	host "github.com/crossplane-contrib/provider-terraform-vsphere/generated/resources/host/v1alpha1"
	host_port_group "github.com/crossplane-contrib/provider-terraform-vsphere/generated/resources/host_port_group/v1alpha1"
	host_virtual_switch "github.com/crossplane-contrib/provider-terraform-vsphere/generated/resources/host_virtual_switch/v1alpha1"
	license "github.com/crossplane-contrib/provider-terraform-vsphere/generated/resources/license/v1alpha1"
	nas_datastore "github.com/crossplane-contrib/provider-terraform-vsphere/generated/resources/nas_datastore/v1alpha1"
	resource_pool "github.com/crossplane-contrib/provider-terraform-vsphere/generated/resources/resource_pool/v1alpha1"
	role "github.com/crossplane-contrib/provider-terraform-vsphere/generated/resources/role/v1alpha1"
	storage_drs_vm_override "github.com/crossplane-contrib/provider-terraform-vsphere/generated/resources/storage_drs_vm_override/v1alpha1"
	tag "github.com/crossplane-contrib/provider-terraform-vsphere/generated/resources/tag/v1alpha1"
	tag_category "github.com/crossplane-contrib/provider-terraform-vsphere/generated/resources/tag_category/v1alpha1"
	vapp_container "github.com/crossplane-contrib/provider-terraform-vsphere/generated/resources/vapp_container/v1alpha1"
	vapp_entity "github.com/crossplane-contrib/provider-terraform-vsphere/generated/resources/vapp_entity/v1alpha1"
	virtual_disk "github.com/crossplane-contrib/provider-terraform-vsphere/generated/resources/virtual_disk/v1alpha1"
	virtual_machine "github.com/crossplane-contrib/provider-terraform-vsphere/generated/resources/virtual_machine/v1alpha1"
	virtual_machine_snapshot "github.com/crossplane-contrib/provider-terraform-vsphere/generated/resources/virtual_machine_snapshot/v1alpha1"
	vm_storage_policy "github.com/crossplane-contrib/provider-terraform-vsphere/generated/resources/vm_storage_policy/v1alpha1"
	vmfs_datastore "github.com/crossplane-contrib/provider-terraform-vsphere/generated/resources/vmfs_datastore/v1alpha1"
	vnic "github.com/crossplane-contrib/provider-terraform-vsphere/generated/resources/vnic/v1alpha1"

	"github.com/crossplane-contrib/terraform-runtime/pkg/plugin"
)

var ResourceImplementations = []*plugin.Implementation{
	license.Implementation(),
	compute_cluster_vm_dependency_rule.Implementation(),
	file.Implementation(),
	content_library_item.Implementation(),
	vmfs_datastore.Implementation(),
	compute_cluster_vm_group.Implementation(),
	distributed_port_group.Implementation(),
	compute_cluster_vm_host_rule.Implementation(),
	tag.Implementation(),
	vnic.Implementation(),
	compute_cluster.Implementation(),
	nas_datastore.Implementation(),
	host_virtual_switch.Implementation(),
	entity_permissions.Implementation(),
	role.Implementation(),
	virtual_machine_snapshot.Implementation(),
	datacenter.Implementation(),
	custom_attribute.Implementation(),
	dpm_host_override.Implementation(),
	content_library.Implementation(),
	drs_vm_override.Implementation(),
	tag_category.Implementation(),
	distributed_virtual_switch.Implementation(),
	vapp_entity.Implementation(),
	ha_vm_override.Implementation(),
	datastore_cluster.Implementation(),
	host.Implementation(),
	virtual_disk.Implementation(),
	compute_cluster_host_group.Implementation(),
	vapp_container.Implementation(),
	datastore_cluster_vm_anti_affinity_rule.Implementation(),
	folder.Implementation(),
	host_port_group.Implementation(),
	compute_cluster_vm_affinity_rule.Implementation(),
	vm_storage_policy.Implementation(),
	resource_pool.Implementation(),
	compute_cluster_vm_anti_affinity_rule.Implementation(),
	virtual_machine.Implementation(),
	storage_drs_vm_override.Implementation(),
}

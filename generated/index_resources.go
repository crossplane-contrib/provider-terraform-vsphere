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

var generatedImplementations = []*plugin.Implementation{
	compute_cluster.Implementation(),
	compute_cluster_host_group.Implementation(),
	compute_cluster_vm_affinity_rule.Implementation(),
	compute_cluster_vm_anti_affinity_rule.Implementation(),
	compute_cluster_vm_dependency_rule.Implementation(),
	compute_cluster_vm_group.Implementation(),
	compute_cluster_vm_host_rule.Implementation(),
	content_library.Implementation(),
	content_library_item.Implementation(),
	custom_attribute.Implementation(),
	datacenter.Implementation(),
	datastore_cluster.Implementation(),
	datastore_cluster_vm_anti_affinity_rule.Implementation(),
	distributed_port_group.Implementation(),
	distributed_virtual_switch.Implementation(),
	dpm_host_override.Implementation(),
	drs_vm_override.Implementation(),
	entity_permissions.Implementation(),
	file.Implementation(),
	folder.Implementation(),
	ha_vm_override.Implementation(),
	host.Implementation(),
	host_port_group.Implementation(),
	host_virtual_switch.Implementation(),
	license.Implementation(),
	nas_datastore.Implementation(),
	resource_pool.Implementation(),
	role.Implementation(),
	storage_drs_vm_override.Implementation(),
	tag.Implementation(),
	tag_category.Implementation(),
	vapp_container.Implementation(),
	vapp_entity.Implementation(),
	virtual_disk.Implementation(),
	virtual_machine.Implementation(),
	virtual_machine_snapshot.Implementation(),
	vm_storage_policy.Implementation(),
	vmfs_datastore.Implementation(),
	vnic.Implementation(),
}

// this is deferred until init time to simplify the codegen workflow.
// index.go can be a simple templated, satisfying the needs of main.go so that
// the provider can be compiled (albeit in a non-functional state) enabling angryjet
// and controller-gen to run against the generated types.go before the a subsequent pass
// of terraform-provider-gen adds the compare/encode/decode methods.
func init() {
	for _, impl := range generatedImplementations {
		resourceImplementations = append(resourceImplementations, impl)
	}
}

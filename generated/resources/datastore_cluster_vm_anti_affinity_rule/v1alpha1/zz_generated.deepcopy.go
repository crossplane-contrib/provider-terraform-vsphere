// +build !ignore_autogenerated

/*
Copyright 2020 The Crossplane Authors.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatastoreClusterVmAntiAffinityRule) DeepCopyInto(out *DatastoreClusterVmAntiAffinityRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatastoreClusterVmAntiAffinityRule.
func (in *DatastoreClusterVmAntiAffinityRule) DeepCopy() *DatastoreClusterVmAntiAffinityRule {
	if in == nil {
		return nil
	}
	out := new(DatastoreClusterVmAntiAffinityRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatastoreClusterVmAntiAffinityRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatastoreClusterVmAntiAffinityRuleList) DeepCopyInto(out *DatastoreClusterVmAntiAffinityRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DatastoreClusterVmAntiAffinityRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatastoreClusterVmAntiAffinityRuleList.
func (in *DatastoreClusterVmAntiAffinityRuleList) DeepCopy() *DatastoreClusterVmAntiAffinityRuleList {
	if in == nil {
		return nil
	}
	out := new(DatastoreClusterVmAntiAffinityRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatastoreClusterVmAntiAffinityRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatastoreClusterVmAntiAffinityRuleObservation) DeepCopyInto(out *DatastoreClusterVmAntiAffinityRuleObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatastoreClusterVmAntiAffinityRuleObservation.
func (in *DatastoreClusterVmAntiAffinityRuleObservation) DeepCopy() *DatastoreClusterVmAntiAffinityRuleObservation {
	if in == nil {
		return nil
	}
	out := new(DatastoreClusterVmAntiAffinityRuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatastoreClusterVmAntiAffinityRuleParameters) DeepCopyInto(out *DatastoreClusterVmAntiAffinityRuleParameters) {
	*out = *in
	if in.VirtualMachineIds != nil {
		in, out := &in.VirtualMachineIds, &out.VirtualMachineIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatastoreClusterVmAntiAffinityRuleParameters.
func (in *DatastoreClusterVmAntiAffinityRuleParameters) DeepCopy() *DatastoreClusterVmAntiAffinityRuleParameters {
	if in == nil {
		return nil
	}
	out := new(DatastoreClusterVmAntiAffinityRuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatastoreClusterVmAntiAffinityRuleSpec) DeepCopyInto(out *DatastoreClusterVmAntiAffinityRuleSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatastoreClusterVmAntiAffinityRuleSpec.
func (in *DatastoreClusterVmAntiAffinityRuleSpec) DeepCopy() *DatastoreClusterVmAntiAffinityRuleSpec {
	if in == nil {
		return nil
	}
	out := new(DatastoreClusterVmAntiAffinityRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatastoreClusterVmAntiAffinityRuleStatus) DeepCopyInto(out *DatastoreClusterVmAntiAffinityRuleStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatastoreClusterVmAntiAffinityRuleStatus.
func (in *DatastoreClusterVmAntiAffinityRuleStatus) DeepCopy() *DatastoreClusterVmAntiAffinityRuleStatus {
	if in == nil {
		return nil
	}
	out := new(DatastoreClusterVmAntiAffinityRuleStatus)
	in.DeepCopyInto(out)
	return out
}

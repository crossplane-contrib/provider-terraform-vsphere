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
func (in *NasDatastore) DeepCopyInto(out *NasDatastore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NasDatastore.
func (in *NasDatastore) DeepCopy() *NasDatastore {
	if in == nil {
		return nil
	}
	out := new(NasDatastore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NasDatastore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NasDatastoreList) DeepCopyInto(out *NasDatastoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NasDatastore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NasDatastoreList.
func (in *NasDatastoreList) DeepCopy() *NasDatastoreList {
	if in == nil {
		return nil
	}
	out := new(NasDatastoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NasDatastoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NasDatastoreObservation) DeepCopyInto(out *NasDatastoreObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NasDatastoreObservation.
func (in *NasDatastoreObservation) DeepCopy() *NasDatastoreObservation {
	if in == nil {
		return nil
	}
	out := new(NasDatastoreObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NasDatastoreParameters) DeepCopyInto(out *NasDatastoreParameters) {
	*out = *in
	if in.CustomAttributes != nil {
		in, out := &in.CustomAttributes, &out.CustomAttributes
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.HostSystemIds != nil {
		in, out := &in.HostSystemIds, &out.HostSystemIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RemoteHosts != nil {
		in, out := &in.RemoteHosts, &out.RemoteHosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NasDatastoreParameters.
func (in *NasDatastoreParameters) DeepCopy() *NasDatastoreParameters {
	if in == nil {
		return nil
	}
	out := new(NasDatastoreParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NasDatastoreSpec) DeepCopyInto(out *NasDatastoreSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NasDatastoreSpec.
func (in *NasDatastoreSpec) DeepCopy() *NasDatastoreSpec {
	if in == nil {
		return nil
	}
	out := new(NasDatastoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NasDatastoreStatus) DeepCopyInto(out *NasDatastoreStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NasDatastoreStatus.
func (in *NasDatastoreStatus) DeepCopy() *NasDatastoreStatus {
	if in == nil {
		return nil
	}
	out := new(NasDatastoreStatus)
	in.DeepCopyInto(out)
	return out
}

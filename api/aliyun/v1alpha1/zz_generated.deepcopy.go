//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 The kubeall.com Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Addon) DeepCopyInto(out *Addon) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Addon.
func (in *Addon) DeepCopy() *Addon {
	if in == nil {
		return nil
	}
	out := new(Addon)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AliyunACK) DeepCopyInto(out *AliyunACK) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AliyunACK.
func (in *AliyunACK) DeepCopy() *AliyunACK {
	if in == nil {
		return nil
	}
	out := new(AliyunACK)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AliyunACK) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AliyunACKList) DeepCopyInto(out *AliyunACKList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AliyunACK, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AliyunACKList.
func (in *AliyunACKList) DeepCopy() *AliyunACKList {
	if in == nil {
		return nil
	}
	out := new(AliyunACKList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AliyunACKList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AliyunACKSpec) DeepCopyInto(out *AliyunACKSpec) {
	*out = *in
	in.Ask.DeepCopyInto(&out.Ask)
	out.Kubernetes = in.Kubernetes
	out.ManagedKubernetes = in.ManagedKubernetes
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AliyunACKSpec.
func (in *AliyunACKSpec) DeepCopy() *AliyunACKSpec {
	if in == nil {
		return nil
	}
	out := new(AliyunACKSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AliyunACKStatus) DeepCopyInto(out *AliyunACKStatus) {
	*out = *in
	in.Expiration.DeepCopyInto(&out.Expiration)
	if in.Tasks != nil {
		in, out := &in.Tasks, &out.Tasks
		*out = make([]AliyunTask, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AliyunACKStatus.
func (in *AliyunACKStatus) DeepCopy() *AliyunACKStatus {
	if in == nil {
		return nil
	}
	out := new(AliyunACKStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AliyunTask) DeepCopyInto(out *AliyunTask) {
	*out = *in
	in.Time.DeepCopyInto(&out.Time)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AliyunTask.
func (in *AliyunTask) DeepCopy() *AliyunTask {
	if in == nil {
		return nil
	}
	out := new(AliyunTask)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ask) DeepCopyInto(out *Ask) {
	*out = *in
	if in.ServiceDiscoveryTypes != nil {
		in, out := &in.ServiceDiscoveryTypes, &out.ServiceDiscoveryTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]Tag, len(*in))
		copy(*out, *in)
	}
	if in.Addons != nil {
		in, out := &in.Addons, &out.Addons
		*out = make([]Addon, len(*in))
		copy(*out, *in)
	}
	if in.VswitchIds != nil {
		in, out := &in.VswitchIds, &out.VswitchIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ask.
func (in *Ask) DeepCopy() *Ask {
	if in == nil {
		return nil
	}
	out := new(Ask)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Kubernetes) DeepCopyInto(out *Kubernetes) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Kubernetes.
func (in *Kubernetes) DeepCopy() *Kubernetes {
	if in == nil {
		return nil
	}
	out := new(Kubernetes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedKubernetes) DeepCopyInto(out *ManagedKubernetes) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedKubernetes.
func (in *ManagedKubernetes) DeepCopy() *ManagedKubernetes {
	if in == nil {
		return nil
	}
	out := new(ManagedKubernetes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tag) DeepCopyInto(out *Tag) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tag.
func (in *Tag) DeepCopy() *Tag {
	if in == nil {
		return nil
	}
	out := new(Tag)
	in.DeepCopyInto(out)
	return out
}

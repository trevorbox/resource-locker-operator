// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	apis "github.com/redhat-cop/operator-utils/pkg/util/apis"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Resource) DeepCopyInto(out *Resource) {
	*out = *in
	in.Object.DeepCopyInto(&out.Object)
	if in.ExcludedPaths != nil {
		in, out := &in.ExcludedPaths, &out.ExcludedPaths
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resource.
func (in *Resource) DeepCopy() *Resource {
	if in == nil {
		return nil
	}
	out := new(Resource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceLocker) DeepCopyInto(out *ResourceLocker) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceLocker.
func (in *ResourceLocker) DeepCopy() *ResourceLocker {
	if in == nil {
		return nil
	}
	out := new(ResourceLocker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceLocker) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceLockerList) DeepCopyInto(out *ResourceLockerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ResourceLocker, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceLockerList.
func (in *ResourceLockerList) DeepCopy() *ResourceLockerList {
	if in == nil {
		return nil
	}
	out := new(ResourceLockerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceLockerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceLockerSpec) DeepCopyInto(out *ResourceLockerSpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]Resource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Patches != nil {
		in, out := &in.Patches, &out.Patches
		*out = make([]apis.Patch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.ServiceAccountRef = in.ServiceAccountRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceLockerSpec.
func (in *ResourceLockerSpec) DeepCopy() *ResourceLockerSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceLockerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceLockerStatus) DeepCopyInto(out *ResourceLockerStatus) {
	*out = *in
	in.EnforcingReconcileStatus.DeepCopyInto(&out.EnforcingReconcileStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceLockerStatus.
func (in *ResourceLockerStatus) DeepCopy() *ResourceLockerStatus {
	if in == nil {
		return nil
	}
	out := new(ResourceLockerStatus)
	in.DeepCopyInto(out)
	return out
}

// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterCodeTaskRef) DeepCopyInto(out *ClusterCodeTaskRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterCodeTaskRef.
func (in *ClusterCodeTaskRef) DeepCopy() *ClusterCodeTaskRef {
	if in == nil {
		return nil
	}
	out := new(ClusterCodeTaskRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClustercodePlan) DeepCopyInto(out *ClustercodePlan) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClustercodePlan.
func (in *ClustercodePlan) DeepCopy() *ClustercodePlan {
	if in == nil {
		return nil
	}
	out := new(ClustercodePlan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClustercodePlan) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClustercodePlanList) DeepCopyInto(out *ClustercodePlanList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClustercodePlan, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClustercodePlanList.
func (in *ClustercodePlanList) DeepCopy() *ClustercodePlanList {
	if in == nil {
		return nil
	}
	out := new(ClustercodePlanList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClustercodePlanList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClustercodePlanSpec) DeepCopyInto(out *ClustercodePlanSpec) {
	*out = *in
	in.ScanSpec.DeepCopyInto(&out.ScanSpec)
	in.EncodeSpec.DeepCopyInto(&out.EncodeSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClustercodePlanSpec.
func (in *ClustercodePlanSpec) DeepCopy() *ClustercodePlanSpec {
	if in == nil {
		return nil
	}
	out := new(ClustercodePlanSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClustercodePlanStatus) DeepCopyInto(out *ClustercodePlanStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CurrentTasks != nil {
		in, out := &in.CurrentTasks, &out.CurrentTasks
		*out = make([]ClusterCodeTaskRef, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClustercodePlanStatus.
func (in *ClustercodePlanStatus) DeepCopy() *ClustercodePlanStatus {
	if in == nil {
		return nil
	}
	out := new(ClustercodePlanStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClustercodeTask) DeepCopyInto(out *ClustercodeTask) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClustercodeTask.
func (in *ClustercodeTask) DeepCopy() *ClustercodeTask {
	if in == nil {
		return nil
	}
	out := new(ClustercodeTask)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClustercodeTask) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClustercodeTaskList) DeepCopyInto(out *ClustercodeTaskList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClustercodeTask, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClustercodeTaskList.
func (in *ClustercodeTaskList) DeepCopy() *ClustercodeTaskList {
	if in == nil {
		return nil
	}
	out := new(ClustercodeTaskList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClustercodeTaskList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClustercodeTaskSpec) DeepCopyInto(out *ClustercodeTaskSpec) {
	*out = *in
	in.EncodeSpec.DeepCopyInto(&out.EncodeSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClustercodeTaskSpec.
func (in *ClustercodeTaskSpec) DeepCopy() *ClustercodeTaskSpec {
	if in == nil {
		return nil
	}
	out := new(ClustercodeTaskSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClustercodeTaskStatus) DeepCopyInto(out *ClustercodeTaskStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClustercodeTaskStatus.
func (in *ClustercodeTaskStatus) DeepCopy() *ClustercodeTaskStatus {
	if in == nil {
		return nil
	}
	out := new(ClustercodeTaskStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncodeSpec) DeepCopyInto(out *EncodeSpec) {
	*out = *in
	if in.DefaultCommandArgs != nil {
		in, out := &in.DefaultCommandArgs, &out.DefaultCommandArgs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SplitCommandArgs != nil {
		in, out := &in.SplitCommandArgs, &out.SplitCommandArgs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TranscodeCommandArgs != nil {
		in, out := &in.TranscodeCommandArgs, &out.TranscodeCommandArgs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MergeCommandArgs != nil {
		in, out := &in.MergeCommandArgs, &out.MergeCommandArgs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncodeSpec.
func (in *EncodeSpec) DeepCopy() *EncodeSpec {
	if in == nil {
		return nil
	}
	out := new(EncodeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScanSpec) DeepCopyInto(out *ScanSpec) {
	*out = *in
	if in.MediaFileExtensions != nil {
		in, out := &in.MediaFileExtensions, &out.MediaFileExtensions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScanSpec.
func (in *ScanSpec) DeepCopy() *ScanSpec {
	if in == nil {
		return nil
	}
	out := new(ScanSpec)
	in.DeepCopyInto(out)
	return out
}

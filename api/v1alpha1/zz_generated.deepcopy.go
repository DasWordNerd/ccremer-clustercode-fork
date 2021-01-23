// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Blueprint) DeepCopyInto(out *Blueprint) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Blueprint.
func (in *Blueprint) DeepCopy() *Blueprint {
	if in == nil {
		return nil
	}
	out := new(Blueprint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Blueprint) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlueprintList) DeepCopyInto(out *BlueprintList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Blueprint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlueprintList.
func (in *BlueprintList) DeepCopy() *BlueprintList {
	if in == nil {
		return nil
	}
	out := new(BlueprintList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BlueprintList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlueprintSpec) DeepCopyInto(out *BlueprintSpec) {
	*out = *in
	out.Storage = in.Storage
	in.TaskConcurrencyStrategy.DeepCopyInto(&out.TaskConcurrencyStrategy)
	in.ScanSpec.DeepCopyInto(&out.ScanSpec)
	in.EncodeSpec.DeepCopyInto(&out.EncodeSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlueprintSpec.
func (in *BlueprintSpec) DeepCopy() *BlueprintSpec {
	if in == nil {
		return nil
	}
	out := new(BlueprintSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlueprintStatus) DeepCopyInto(out *BlueprintStatus) {
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
		*out = make([]TaskRef, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlueprintStatus.
func (in *BlueprintStatus) DeepCopy() *BlueprintStatus {
	if in == nil {
		return nil
	}
	out := new(BlueprintStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterCodeVolumeRef) DeepCopyInto(out *ClusterCodeVolumeRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterCodeVolumeRef.
func (in *ClusterCodeVolumeRef) DeepCopy() *ClusterCodeVolumeRef {
	if in == nil {
		return nil
	}
	out := new(ClusterCodeVolumeRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClustercodeCountStrategy) DeepCopyInto(out *ClustercodeCountStrategy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClustercodeCountStrategy.
func (in *ClustercodeCountStrategy) DeepCopy() *ClustercodeCountStrategy {
	if in == nil {
		return nil
	}
	out := new(ClustercodeCountStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClustercodeSliceRef) DeepCopyInto(out *ClustercodeSliceRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClustercodeSliceRef.
func (in *ClustercodeSliceRef) DeepCopy() *ClustercodeSliceRef {
	if in == nil {
		return nil
	}
	out := new(ClustercodeSliceRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClustercodeStrategy) DeepCopyInto(out *ClustercodeStrategy) {
	*out = *in
	if in.ConcurrentCountStrategy != nil {
		in, out := &in.ConcurrentCountStrategy, &out.ConcurrentCountStrategy
		*out = new(ClustercodeCountStrategy)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClustercodeStrategy.
func (in *ClustercodeStrategy) DeepCopy() *ClustercodeStrategy {
	if in == nil {
		return nil
	}
	out := new(ClustercodeStrategy)
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

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageSpec) DeepCopyInto(out *StorageSpec) {
	*out = *in
	out.SourcePvc = in.SourcePvc
	out.IntermediatePvc = in.IntermediatePvc
	out.TargetPvc = in.TargetPvc
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageSpec.
func (in *StorageSpec) DeepCopy() *StorageSpec {
	if in == nil {
		return nil
	}
	out := new(StorageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Task) DeepCopyInto(out *Task) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Task.
func (in *Task) DeepCopy() *Task {
	if in == nil {
		return nil
	}
	out := new(Task)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Task) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskList) DeepCopyInto(out *TaskList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Task, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskList.
func (in *TaskList) DeepCopy() *TaskList {
	if in == nil {
		return nil
	}
	out := new(TaskList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TaskList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskRef) DeepCopyInto(out *TaskRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskRef.
func (in *TaskRef) DeepCopy() *TaskRef {
	if in == nil {
		return nil
	}
	out := new(TaskRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskSpec) DeepCopyInto(out *TaskSpec) {
	*out = *in
	out.Storage = in.Storage
	in.EncodeSpec.DeepCopyInto(&out.EncodeSpec)
	in.ConcurrencyStrategy.DeepCopyInto(&out.ConcurrencyStrategy)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskSpec.
func (in *TaskSpec) DeepCopy() *TaskSpec {
	if in == nil {
		return nil
	}
	out := new(TaskSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskStatus) DeepCopyInto(out *TaskStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SlicesScheduled != nil {
		in, out := &in.SlicesScheduled, &out.SlicesScheduled
		*out = make([]ClustercodeSliceRef, len(*in))
		copy(*out, *in)
	}
	if in.SlicesFinished != nil {
		in, out := &in.SlicesFinished, &out.SlicesFinished
		*out = make([]ClustercodeSliceRef, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskStatus.
func (in *TaskStatus) DeepCopy() *TaskStatus {
	if in == nil {
		return nil
	}
	out := new(TaskStatus)
	in.DeepCopyInto(out)
	return out
}

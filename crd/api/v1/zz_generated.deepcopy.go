//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*


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

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComparableProbeOutputItem) DeepCopyInto(out *ComparableProbeOutputItem) {
	*out = *in
	out.Source = in.Source
	out.Destination = in.Destination
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComparableProbeOutputItem.
func (in *ComparableProbeOutputItem) DeepCopy() *ComparableProbeOutputItem {
	if in == nil {
		return nil
	}
	out := new(ComparableProbeOutputItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Kubesonde) DeepCopyInto(out *Kubesonde) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Kubesonde.
func (in *Kubesonde) DeepCopy() *Kubesonde {
	if in == nil {
		return nil
	}
	out := new(Kubesonde)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Kubesonde) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubesondeList) DeepCopyInto(out *KubesondeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Kubesonde, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubesondeList.
func (in *KubesondeList) DeepCopy() *KubesondeList {
	if in == nil {
		return nil
	}
	out := new(KubesondeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubesondeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubesondeSpec) DeepCopyInto(out *KubesondeSpec) {
	*out = *in
	if in.Actions != nil {
		in, out := &in.Actions, &out.Actions
		*out = make([]ProbingAction, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubesondeSpec.
func (in *KubesondeSpec) DeepCopy() *KubesondeSpec {
	if in == nil {
		return nil
	}
	out := new(KubesondeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubesondeStatus) DeepCopyInto(out *KubesondeStatus) {
	*out = *in
	if in.LastProbeTime != nil {
		in, out := &in.LastProbeTime, &out.LastProbeTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubesondeStatus.
func (in *KubesondeStatus) DeepCopy() *KubesondeStatus {
	if in == nil {
		return nil
	}
	out := new(KubesondeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodNetworkingInfo) DeepCopyInto(out *PodNetworkingInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodNetworkingInfo.
func (in *PodNetworkingInfo) DeepCopy() *PodNetworkingInfo {
	if in == nil {
		return nil
	}
	out := new(PodNetworkingInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in PodNetworkingInfoV2) DeepCopyInto(out *PodNetworkingInfoV2) {
	{
		in := &in
		*out = make(PodNetworkingInfoV2, len(*in))
		for key, val := range *in {
			var outVal []PodNetworkingItem
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]PodNetworkingItem, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodNetworkingInfoV2.
func (in PodNetworkingInfoV2) DeepCopy() PodNetworkingInfoV2 {
	if in == nil {
		return nil
	}
	out := new(PodNetworkingInfoV2)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodNetworkingItem) DeepCopyInto(out *PodNetworkingItem) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodNetworkingItem.
func (in *PodNetworkingItem) DeepCopy() *PodNetworkingItem {
	if in == nil {
		return nil
	}
	out := new(PodNetworkingItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProbeEndpointInfo) DeepCopyInto(out *ProbeEndpointInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProbeEndpointInfo.
func (in *ProbeEndpointInfo) DeepCopy() *ProbeEndpointInfo {
	if in == nil {
		return nil
	}
	out := new(ProbeEndpointInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProbeOutput) DeepCopyInto(out *ProbeOutput) {
	*out = *in
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ProbeOutputItem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Errors != nil {
		in, out := &in.Errors, &out.Errors
		*out = make([]ProbeOutputError, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PodNetworking != nil {
		in, out := &in.PodNetworking, &out.PodNetworking
		*out = make([]PodNetworkingInfo, len(*in))
		copy(*out, *in)
	}
	if in.PodNetworkingV2 != nil {
		in, out := &in.PodNetworkingV2, &out.PodNetworkingV2
		*out = make(PodNetworkingInfoV2, len(*in))
		for key, val := range *in {
			var outVal []PodNetworkingItem
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]PodNetworkingItem, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	if in.PodConfigurationNetworking != nil {
		in, out := &in.PodConfigurationNetworking, &out.PodConfigurationNetworking
		*out = make(PodNetworkingInfoV2, len(*in))
		for key, val := range *in {
			var outVal []PodNetworkingItem
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]PodNetworkingItem, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProbeOutput.
func (in *ProbeOutput) DeepCopy() *ProbeOutput {
	if in == nil {
		return nil
	}
	out := new(ProbeOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProbeOutputError) DeepCopyInto(out *ProbeOutputError) {
	*out = *in
	in.Value.DeepCopyInto(&out.Value)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProbeOutputError.
func (in *ProbeOutputError) DeepCopy() *ProbeOutputError {
	if in == nil {
		return nil
	}
	out := new(ProbeOutputError)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProbeOutputItem) DeepCopyInto(out *ProbeOutputItem) {
	*out = *in
	out.Source = in.Source
	out.Destination = in.Destination
	if in.DestinationHostnames != nil {
		in, out := &in.DestinationHostnames, &out.DestinationHostnames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProbeOutputItem.
func (in *ProbeOutputItem) DeepCopy() *ProbeOutputItem {
	if in == nil {
		return nil
	}
	out := new(ProbeOutputItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProbingAction) DeepCopyInto(out *ProbingAction) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProbingAction.
func (in *ProbingAction) DeepCopy() *ProbingAction {
	if in == nil {
		return nil
	}
	out := new(ProbingAction)
	in.DeepCopyInto(out)
	return out
}
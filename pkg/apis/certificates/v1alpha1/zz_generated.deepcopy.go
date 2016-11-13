// +build !ignore_autogenerated

/*
Copyright 2016 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1alpha1

import (
	unversioned "k8s.io/kubernetes/pkg/api/unversioned"
)

// DeepCopyInto will perform a deep copy of the receiver, writing to out. in must be non-nil.
func (in *CertificateSigningRequest) DeepCopyInto(out *CertificateSigningRequest) {
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy will perform a deep copy of the receiver, creating a new CertificateSigningRequest.
func (x *CertificateSigningRequest) DeepCopy() *CertificateSigningRequest {
	if x == nil {
		return nil
	}
	out := new(CertificateSigningRequest)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject will perform a deep copy of the receiver, creating a new object.
func (x *CertificateSigningRequest) DeepCopyObject() unversioned.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto will perform a deep copy of the receiver, writing to out. in must be non-nil.
func (in *CertificateSigningRequestCondition) DeepCopyInto(out *CertificateSigningRequestCondition) {
	out.Type = in.Type
	out.Reason = in.Reason
	out.Message = in.Message
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	return
}

// DeepCopy will perform a deep copy of the receiver, creating a new CertificateSigningRequestCondition.
func (x *CertificateSigningRequestCondition) DeepCopy() *CertificateSigningRequestCondition {
	if x == nil {
		return nil
	}
	out := new(CertificateSigningRequestCondition)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto will perform a deep copy of the receiver, writing to out. in must be non-nil.
func (in *CertificateSigningRequestList) DeepCopyInto(out *CertificateSigningRequestList) {
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CertificateSigningRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	} else {
		out.Items = nil
	}
	return
}

// DeepCopy will perform a deep copy of the receiver, creating a new CertificateSigningRequestList.
func (x *CertificateSigningRequestList) DeepCopy() *CertificateSigningRequestList {
	if x == nil {
		return nil
	}
	out := new(CertificateSigningRequestList)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject will perform a deep copy of the receiver, creating a new object.
func (x *CertificateSigningRequestList) DeepCopyObject() unversioned.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto will perform a deep copy of the receiver, writing to out. in must be non-nil.
func (in *CertificateSigningRequestSpec) DeepCopyInto(out *CertificateSigningRequestSpec) {
	if in.Request != nil {
		in, out := &in.Request, &out.Request
		*out = make([]byte, len(*in))
		copy(*out, *in)
	} else {
		out.Request = nil
	}
	out.Username = in.Username
	out.UID = in.UID
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]string, len(*in))
		copy(*out, *in)
	} else {
		out.Groups = nil
	}
	return
}

// DeepCopy will perform a deep copy of the receiver, creating a new CertificateSigningRequestSpec.
func (x *CertificateSigningRequestSpec) DeepCopy() *CertificateSigningRequestSpec {
	if x == nil {
		return nil
	}
	out := new(CertificateSigningRequestSpec)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto will perform a deep copy of the receiver, writing to out. in must be non-nil.
func (in *CertificateSigningRequestStatus) DeepCopyInto(out *CertificateSigningRequestStatus) {
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]CertificateSigningRequestCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	} else {
		out.Conditions = nil
	}
	if in.Certificate != nil {
		in, out := &in.Certificate, &out.Certificate
		*out = make([]byte, len(*in))
		copy(*out, *in)
	} else {
		out.Certificate = nil
	}
	return
}

// DeepCopy will perform a deep copy of the receiver, creating a new CertificateSigningRequestStatus.
func (x *CertificateSigningRequestStatus) DeepCopy() *CertificateSigningRequestStatus {
	if x == nil {
		return nil
	}
	out := new(CertificateSigningRequestStatus)
	x.DeepCopyInto(out)
	return out
}

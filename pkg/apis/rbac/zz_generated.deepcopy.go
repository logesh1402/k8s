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

package rbac

import (
	unversioned "k8s.io/kubernetes/pkg/api/unversioned"
)

// DeepCopyInto will perform a deep copy of the receiver, writing to out. in must be non-nil.
func (in *ClusterRole) DeepCopyInto(out *ClusterRole) {
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]PolicyRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	} else {
		out.Rules = nil
	}
	return
}

// DeepCopy will perform a deep copy of the receiver, creating a new ClusterRole.
func (x *ClusterRole) DeepCopy() *ClusterRole {
	if x == nil {
		return nil
	}
	out := new(ClusterRole)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject will perform a deep copy of the receiver, creating a new object.
func (x *ClusterRole) DeepCopyObject() unversioned.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto will perform a deep copy of the receiver, writing to out. in must be non-nil.
func (in *ClusterRoleBinding) DeepCopyInto(out *ClusterRoleBinding) {
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Subjects != nil {
		in, out := &in.Subjects, &out.Subjects
		*out = make([]Subject, len(*in))
		for i := range *in {
			(*out)[i] = (*in)[i]
		}
	} else {
		out.Subjects = nil
	}
	out.RoleRef = in.RoleRef
	return
}

// DeepCopy will perform a deep copy of the receiver, creating a new ClusterRoleBinding.
func (x *ClusterRoleBinding) DeepCopy() *ClusterRoleBinding {
	if x == nil {
		return nil
	}
	out := new(ClusterRoleBinding)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject will perform a deep copy of the receiver, creating a new object.
func (x *ClusterRoleBinding) DeepCopyObject() unversioned.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto will perform a deep copy of the receiver, writing to out. in must be non-nil.
func (in *ClusterRoleBindingList) DeepCopyInto(out *ClusterRoleBindingList) {
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterRoleBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	} else {
		out.Items = nil
	}
	return
}

// DeepCopy will perform a deep copy of the receiver, creating a new ClusterRoleBindingList.
func (x *ClusterRoleBindingList) DeepCopy() *ClusterRoleBindingList {
	if x == nil {
		return nil
	}
	out := new(ClusterRoleBindingList)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject will perform a deep copy of the receiver, creating a new object.
func (x *ClusterRoleBindingList) DeepCopyObject() unversioned.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto will perform a deep copy of the receiver, writing to out. in must be non-nil.
func (in *ClusterRoleList) DeepCopyInto(out *ClusterRoleList) {
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	} else {
		out.Items = nil
	}
	return
}

// DeepCopy will perform a deep copy of the receiver, creating a new ClusterRoleList.
func (x *ClusterRoleList) DeepCopy() *ClusterRoleList {
	if x == nil {
		return nil
	}
	out := new(ClusterRoleList)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject will perform a deep copy of the receiver, creating a new object.
func (x *ClusterRoleList) DeepCopyObject() unversioned.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto will perform a deep copy of the receiver, writing to out. in must be non-nil.
func (in *PolicyRule) DeepCopyInto(out *PolicyRule) {
	if in.Verbs != nil {
		in, out := &in.Verbs, &out.Verbs
		*out = make([]string, len(*in))
		copy(*out, *in)
	} else {
		out.Verbs = nil
	}
	out.AttributeRestrictions = in.AttributeRestrictions.DeepCopyObject()
	if in.APIGroups != nil {
		in, out := &in.APIGroups, &out.APIGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	} else {
		out.APIGroups = nil
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]string, len(*in))
		copy(*out, *in)
	} else {
		out.Resources = nil
	}
	if in.ResourceNames != nil {
		in, out := &in.ResourceNames, &out.ResourceNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	} else {
		out.ResourceNames = nil
	}
	if in.NonResourceURLs != nil {
		in, out := &in.NonResourceURLs, &out.NonResourceURLs
		*out = make([]string, len(*in))
		copy(*out, *in)
	} else {
		out.NonResourceURLs = nil
	}
	return
}

// DeepCopy will perform a deep copy of the receiver, creating a new PolicyRule.
func (x *PolicyRule) DeepCopy() *PolicyRule {
	if x == nil {
		return nil
	}
	out := new(PolicyRule)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto will perform a deep copy of the receiver, writing to out. in must be non-nil.
func (in *Role) DeepCopyInto(out *Role) {
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]PolicyRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	} else {
		out.Rules = nil
	}
	return
}

// DeepCopy will perform a deep copy of the receiver, creating a new Role.
func (x *Role) DeepCopy() *Role {
	if x == nil {
		return nil
	}
	out := new(Role)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject will perform a deep copy of the receiver, creating a new object.
func (x *Role) DeepCopyObject() unversioned.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto will perform a deep copy of the receiver, writing to out. in must be non-nil.
func (in *RoleBinding) DeepCopyInto(out *RoleBinding) {
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Subjects != nil {
		in, out := &in.Subjects, &out.Subjects
		*out = make([]Subject, len(*in))
		for i := range *in {
			(*out)[i] = (*in)[i]
		}
	} else {
		out.Subjects = nil
	}
	out.RoleRef = in.RoleRef
	return
}

// DeepCopy will perform a deep copy of the receiver, creating a new RoleBinding.
func (x *RoleBinding) DeepCopy() *RoleBinding {
	if x == nil {
		return nil
	}
	out := new(RoleBinding)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject will perform a deep copy of the receiver, creating a new object.
func (x *RoleBinding) DeepCopyObject() unversioned.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto will perform a deep copy of the receiver, writing to out. in must be non-nil.
func (in *RoleBindingList) DeepCopyInto(out *RoleBindingList) {
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RoleBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	} else {
		out.Items = nil
	}
	return
}

// DeepCopy will perform a deep copy of the receiver, creating a new RoleBindingList.
func (x *RoleBindingList) DeepCopy() *RoleBindingList {
	if x == nil {
		return nil
	}
	out := new(RoleBindingList)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject will perform a deep copy of the receiver, creating a new object.
func (x *RoleBindingList) DeepCopyObject() unversioned.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto will perform a deep copy of the receiver, writing to out. in must be non-nil.
func (in *RoleList) DeepCopyInto(out *RoleList) {
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Role, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	} else {
		out.Items = nil
	}
	return
}

// DeepCopy will perform a deep copy of the receiver, creating a new RoleList.
func (x *RoleList) DeepCopy() *RoleList {
	if x == nil {
		return nil
	}
	out := new(RoleList)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject will perform a deep copy of the receiver, creating a new object.
func (x *RoleList) DeepCopyObject() unversioned.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto will perform a deep copy of the receiver, writing to out. in must be non-nil.
func (in *RoleRef) DeepCopyInto(out *RoleRef) {
	out.APIGroup = in.APIGroup
	out.Kind = in.Kind
	out.Name = in.Name
	return
}

// DeepCopy will perform a deep copy of the receiver, creating a new RoleRef.
func (x *RoleRef) DeepCopy() *RoleRef {
	if x == nil {
		return nil
	}
	out := new(RoleRef)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto will perform a deep copy of the receiver, writing to out. in must be non-nil.
func (in *Subject) DeepCopyInto(out *Subject) {
	out.Kind = in.Kind
	out.APIVersion = in.APIVersion
	out.Name = in.Name
	out.Namespace = in.Namespace
	return
}

// DeepCopy will perform a deep copy of the receiver, creating a new Subject.
func (x *Subject) DeepCopy() *Subject {
	if x == nil {
		return nil
	}
	out := new(Subject)
	x.DeepCopyInto(out)
	return out
}

// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

package apiextensions

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_apiextensions_CustomResource, InType: reflect.TypeOf(&CustomResource{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_apiextensions_CustomResourceCondition, InType: reflect.TypeOf(&CustomResourceCondition{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_apiextensions_CustomResourceList, InType: reflect.TypeOf(&CustomResourceList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_apiextensions_CustomResourceNames, InType: reflect.TypeOf(&CustomResourceNames{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_apiextensions_CustomResourceSpec, InType: reflect.TypeOf(&CustomResourceSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_apiextensions_CustomResourceStatus, InType: reflect.TypeOf(&CustomResourceStatus{})},
	)
}

// DeepCopy_apiextensions_CustomResource is an autogenerated deepcopy function.
func DeepCopy_apiextensions_CustomResource(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CustomResource)
		out := out.(*CustomResource)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if newVal, err := c.DeepCopy(&in.Spec); err != nil {
			return err
		} else {
			out.Spec = *newVal.(*CustomResourceSpec)
		}
		if newVal, err := c.DeepCopy(&in.Status); err != nil {
			return err
		} else {
			out.Status = *newVal.(*CustomResourceStatus)
		}
		return nil
	}
}

// DeepCopy_apiextensions_CustomResourceCondition is an autogenerated deepcopy function.
func DeepCopy_apiextensions_CustomResourceCondition(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CustomResourceCondition)
		out := out.(*CustomResourceCondition)
		*out = *in
		out.LastTransitionTime = in.LastTransitionTime.DeepCopy()
		return nil
	}
}

// DeepCopy_apiextensions_CustomResourceList is an autogenerated deepcopy function.
func DeepCopy_apiextensions_CustomResourceList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CustomResourceList)
		out := out.(*CustomResourceList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]CustomResource, len(*in))
			for i := range *in {
				if newVal, err := c.DeepCopy(&(*in)[i]); err != nil {
					return err
				} else {
					(*out)[i] = *newVal.(*CustomResource)
				}
			}
		}
		return nil
	}
}

// DeepCopy_apiextensions_CustomResourceNames is an autogenerated deepcopy function.
func DeepCopy_apiextensions_CustomResourceNames(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CustomResourceNames)
		out := out.(*CustomResourceNames)
		*out = *in
		if in.ShortNames != nil {
			in, out := &in.ShortNames, &out.ShortNames
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopy_apiextensions_CustomResourceSpec is an autogenerated deepcopy function.
func DeepCopy_apiextensions_CustomResourceSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CustomResourceSpec)
		out := out.(*CustomResourceSpec)
		*out = *in
		if newVal, err := c.DeepCopy(&in.Names); err != nil {
			return err
		} else {
			out.Names = *newVal.(*CustomResourceNames)
		}
		return nil
	}
}

// DeepCopy_apiextensions_CustomResourceStatus is an autogenerated deepcopy function.
func DeepCopy_apiextensions_CustomResourceStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CustomResourceStatus)
		out := out.(*CustomResourceStatus)
		*out = *in
		if in.Conditions != nil {
			in, out := &in.Conditions, &out.Conditions
			*out = make([]CustomResourceCondition, len(*in))
			for i := range *in {
				if newVal, err := c.DeepCopy(&(*in)[i]); err != nil {
					return err
				} else {
					(*out)[i] = *newVal.(*CustomResourceCondition)
				}
			}
		}
		if newVal, err := c.DeepCopy(&in.AcceptedNames); err != nil {
			return err
		} else {
			out.AcceptedNames = *newVal.(*CustomResourceNames)
		}
		return nil
	}
}

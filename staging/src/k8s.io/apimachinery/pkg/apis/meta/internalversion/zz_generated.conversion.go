// +build !ignore_autogenerated

/*
Copyright 2018 The Kubernetes Authors.

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

// Code generated by conversion-gen. DO NOT EDIT.

// This file was autogenerated by conversion-gen. Do not edit it manually!

package internalversion

import (
	unsafe "unsafe"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_internalversion_List_To_v1_List,
		Convert_v1_List_To_internalversion_List,
		Convert_internalversion_ListOptions_To_v1_ListOptions,
		Convert_v1_ListOptions_To_internalversion_ListOptions,
	)
}

func autoConvert_internalversion_List_To_v1_List(in *List, out *v1.List, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]runtime.RawExtension, len(*in))
		for i := range *in {
			if err := runtime.Convert_runtime_Object_To_runtime_RawExtension(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_internalversion_List_To_v1_List is an autogenerated conversion function.
func Convert_internalversion_List_To_v1_List(in *List, out *v1.List, s conversion.Scope) error {
	return autoConvert_internalversion_List_To_v1_List(in, out, s)
}

func autoConvert_v1_List_To_internalversion_List(in *v1.List, out *List, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]runtime.Object, len(*in))
		for i := range *in {
			if err := runtime.Convert_runtime_RawExtension_To_runtime_Object(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1_List_To_internalversion_List is an autogenerated conversion function.
func Convert_v1_List_To_internalversion_List(in *v1.List, out *List, s conversion.Scope) error {
	return autoConvert_v1_List_To_internalversion_List(in, out, s)
}

func autoConvert_internalversion_ListOptions_To_v1_ListOptions(in *ListOptions, out *v1.ListOptions, s conversion.Scope) error {
	if err := v1.Convert_labels_Selector_To_string(&in.LabelSelector, &out.LabelSelector, s); err != nil {
		return err
	}
	if err := v1.Convert_fields_Selector_To_string(&in.FieldSelector, &out.FieldSelector, s); err != nil {
		return err
	}
	out.IncludeUninitialized = in.IncludeUninitialized
	out.Watch = in.Watch
	out.ResourceVersion = in.ResourceVersion
	out.TimeoutSeconds = (*int64)(unsafe.Pointer(in.TimeoutSeconds))
	out.Limit = in.Limit
	out.Continue = in.Continue
	return nil
}

func autoConvert_v1_ListOptions_To_internalversion_ListOptions(in *v1.ListOptions, out *ListOptions, s conversion.Scope) error {
	if err := v1.Convert_string_To_labels_Selector(&in.LabelSelector, &out.LabelSelector, s); err != nil {
		return err
	}
	if err := v1.Convert_string_To_fields_Selector(&in.FieldSelector, &out.FieldSelector, s); err != nil {
		return err
	}
	out.IncludeUninitialized = in.IncludeUninitialized
	out.Watch = in.Watch
	out.ResourceVersion = in.ResourceVersion
	out.TimeoutSeconds = (*int64)(unsafe.Pointer(in.TimeoutSeconds))
	out.Limit = in.Limit
	out.Continue = in.Continue
	return nil
}

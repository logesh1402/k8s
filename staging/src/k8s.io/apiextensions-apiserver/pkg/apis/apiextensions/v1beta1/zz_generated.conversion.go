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

package v1beta1

import (
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	unsafe "unsafe"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1beta1_CustomResourceDefinition_To_apiextensions_CustomResourceDefinition,
		Convert_apiextensions_CustomResourceDefinition_To_v1beta1_CustomResourceDefinition,
		Convert_v1beta1_CustomResourceDefinitionCondition_To_apiextensions_CustomResourceDefinitionCondition,
		Convert_apiextensions_CustomResourceDefinitionCondition_To_v1beta1_CustomResourceDefinitionCondition,
		Convert_v1beta1_CustomResourceDefinitionList_To_apiextensions_CustomResourceDefinitionList,
		Convert_apiextensions_CustomResourceDefinitionList_To_v1beta1_CustomResourceDefinitionList,
		Convert_v1beta1_CustomResourceDefinitionNames_To_apiextensions_CustomResourceDefinitionNames,
		Convert_apiextensions_CustomResourceDefinitionNames_To_v1beta1_CustomResourceDefinitionNames,
		Convert_v1beta1_CustomResourceDefinitionSpec_To_apiextensions_CustomResourceDefinitionSpec,
		Convert_apiextensions_CustomResourceDefinitionSpec_To_v1beta1_CustomResourceDefinitionSpec,
		Convert_v1beta1_CustomResourceDefinitionStatus_To_apiextensions_CustomResourceDefinitionStatus,
		Convert_apiextensions_CustomResourceDefinitionStatus_To_v1beta1_CustomResourceDefinitionStatus,
		Convert_v1beta1_CustomResourceValidation_To_apiextensions_CustomResourceValidation,
		Convert_apiextensions_CustomResourceValidation_To_v1beta1_CustomResourceValidation,
		Convert_v1beta1_ExternalDocumentation_To_apiextensions_ExternalDocumentation,
		Convert_apiextensions_ExternalDocumentation_To_v1beta1_ExternalDocumentation,
		Convert_v1beta1_JSON_To_apiextensions_JSON,
		Convert_apiextensions_JSON_To_v1beta1_JSON,
		Convert_v1beta1_JSONSchemaProps_To_apiextensions_JSONSchemaProps,
		Convert_apiextensions_JSONSchemaProps_To_v1beta1_JSONSchemaProps,
		Convert_v1beta1_JSONSchemaPropsOrArray_To_apiextensions_JSONSchemaPropsOrArray,
		Convert_apiextensions_JSONSchemaPropsOrArray_To_v1beta1_JSONSchemaPropsOrArray,
		Convert_v1beta1_JSONSchemaPropsOrBool_To_apiextensions_JSONSchemaPropsOrBool,
		Convert_apiextensions_JSONSchemaPropsOrBool_To_v1beta1_JSONSchemaPropsOrBool,
		Convert_v1beta1_JSONSchemaPropsOrStringArray_To_apiextensions_JSONSchemaPropsOrStringArray,
		Convert_apiextensions_JSONSchemaPropsOrStringArray_To_v1beta1_JSONSchemaPropsOrStringArray,
	)
}

func autoConvert_v1beta1_CustomResourceDefinition_To_apiextensions_CustomResourceDefinition(in *CustomResourceDefinition, out *apiextensions.CustomResourceDefinition, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_CustomResourceDefinitionSpec_To_apiextensions_CustomResourceDefinitionSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_CustomResourceDefinitionStatus_To_apiextensions_CustomResourceDefinitionStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_CustomResourceDefinition_To_apiextensions_CustomResourceDefinition is an autogenerated conversion function.
func Convert_v1beta1_CustomResourceDefinition_To_apiextensions_CustomResourceDefinition(in *CustomResourceDefinition, out *apiextensions.CustomResourceDefinition, s conversion.Scope) error {
	return autoConvert_v1beta1_CustomResourceDefinition_To_apiextensions_CustomResourceDefinition(in, out, s)
}

func autoConvert_apiextensions_CustomResourceDefinition_To_v1beta1_CustomResourceDefinition(in *apiextensions.CustomResourceDefinition, out *CustomResourceDefinition, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_apiextensions_CustomResourceDefinitionSpec_To_v1beta1_CustomResourceDefinitionSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_apiextensions_CustomResourceDefinitionStatus_To_v1beta1_CustomResourceDefinitionStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_apiextensions_CustomResourceDefinition_To_v1beta1_CustomResourceDefinition is an autogenerated conversion function.
func Convert_apiextensions_CustomResourceDefinition_To_v1beta1_CustomResourceDefinition(in *apiextensions.CustomResourceDefinition, out *CustomResourceDefinition, s conversion.Scope) error {
	return autoConvert_apiextensions_CustomResourceDefinition_To_v1beta1_CustomResourceDefinition(in, out, s)
}

func autoConvert_v1beta1_CustomResourceDefinitionCondition_To_apiextensions_CustomResourceDefinitionCondition(in *CustomResourceDefinitionCondition, out *apiextensions.CustomResourceDefinitionCondition, s conversion.Scope) error {
	out.Type = apiextensions.CustomResourceDefinitionConditionType(in.Type)
	out.Status = apiextensions.ConditionStatus(in.Status)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_v1beta1_CustomResourceDefinitionCondition_To_apiextensions_CustomResourceDefinitionCondition is an autogenerated conversion function.
func Convert_v1beta1_CustomResourceDefinitionCondition_To_apiextensions_CustomResourceDefinitionCondition(in *CustomResourceDefinitionCondition, out *apiextensions.CustomResourceDefinitionCondition, s conversion.Scope) error {
	return autoConvert_v1beta1_CustomResourceDefinitionCondition_To_apiextensions_CustomResourceDefinitionCondition(in, out, s)
}

func autoConvert_apiextensions_CustomResourceDefinitionCondition_To_v1beta1_CustomResourceDefinitionCondition(in *apiextensions.CustomResourceDefinitionCondition, out *CustomResourceDefinitionCondition, s conversion.Scope) error {
	out.Type = CustomResourceDefinitionConditionType(in.Type)
	out.Status = ConditionStatus(in.Status)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_apiextensions_CustomResourceDefinitionCondition_To_v1beta1_CustomResourceDefinitionCondition is an autogenerated conversion function.
func Convert_apiextensions_CustomResourceDefinitionCondition_To_v1beta1_CustomResourceDefinitionCondition(in *apiextensions.CustomResourceDefinitionCondition, out *CustomResourceDefinitionCondition, s conversion.Scope) error {
	return autoConvert_apiextensions_CustomResourceDefinitionCondition_To_v1beta1_CustomResourceDefinitionCondition(in, out, s)
}

func autoConvert_v1beta1_CustomResourceDefinitionList_To_apiextensions_CustomResourceDefinitionList(in *CustomResourceDefinitionList, out *apiextensions.CustomResourceDefinitionList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]apiextensions.CustomResourceDefinition, len(*in))
		for i := range *in {
			if err := Convert_v1beta1_CustomResourceDefinition_To_apiextensions_CustomResourceDefinition(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1beta1_CustomResourceDefinitionList_To_apiextensions_CustomResourceDefinitionList is an autogenerated conversion function.
func Convert_v1beta1_CustomResourceDefinitionList_To_apiextensions_CustomResourceDefinitionList(in *CustomResourceDefinitionList, out *apiextensions.CustomResourceDefinitionList, s conversion.Scope) error {
	return autoConvert_v1beta1_CustomResourceDefinitionList_To_apiextensions_CustomResourceDefinitionList(in, out, s)
}

func autoConvert_apiextensions_CustomResourceDefinitionList_To_v1beta1_CustomResourceDefinitionList(in *apiextensions.CustomResourceDefinitionList, out *CustomResourceDefinitionList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CustomResourceDefinition, len(*in))
		for i := range *in {
			if err := Convert_apiextensions_CustomResourceDefinition_To_v1beta1_CustomResourceDefinition(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_apiextensions_CustomResourceDefinitionList_To_v1beta1_CustomResourceDefinitionList is an autogenerated conversion function.
func Convert_apiextensions_CustomResourceDefinitionList_To_v1beta1_CustomResourceDefinitionList(in *apiextensions.CustomResourceDefinitionList, out *CustomResourceDefinitionList, s conversion.Scope) error {
	return autoConvert_apiextensions_CustomResourceDefinitionList_To_v1beta1_CustomResourceDefinitionList(in, out, s)
}

func autoConvert_v1beta1_CustomResourceDefinitionNames_To_apiextensions_CustomResourceDefinitionNames(in *CustomResourceDefinitionNames, out *apiextensions.CustomResourceDefinitionNames, s conversion.Scope) error {
	out.Plural = in.Plural
	out.Singular = in.Singular
	out.ShortNames = *(*[]string)(unsafe.Pointer(&in.ShortNames))
	out.Kind = in.Kind
	out.ListKind = in.ListKind
	return nil
}

// Convert_v1beta1_CustomResourceDefinitionNames_To_apiextensions_CustomResourceDefinitionNames is an autogenerated conversion function.
func Convert_v1beta1_CustomResourceDefinitionNames_To_apiextensions_CustomResourceDefinitionNames(in *CustomResourceDefinitionNames, out *apiextensions.CustomResourceDefinitionNames, s conversion.Scope) error {
	return autoConvert_v1beta1_CustomResourceDefinitionNames_To_apiextensions_CustomResourceDefinitionNames(in, out, s)
}

func autoConvert_apiextensions_CustomResourceDefinitionNames_To_v1beta1_CustomResourceDefinitionNames(in *apiextensions.CustomResourceDefinitionNames, out *CustomResourceDefinitionNames, s conversion.Scope) error {
	out.Plural = in.Plural
	out.Singular = in.Singular
	out.ShortNames = *(*[]string)(unsafe.Pointer(&in.ShortNames))
	out.Kind = in.Kind
	out.ListKind = in.ListKind
	return nil
}

// Convert_apiextensions_CustomResourceDefinitionNames_To_v1beta1_CustomResourceDefinitionNames is an autogenerated conversion function.
func Convert_apiextensions_CustomResourceDefinitionNames_To_v1beta1_CustomResourceDefinitionNames(in *apiextensions.CustomResourceDefinitionNames, out *CustomResourceDefinitionNames, s conversion.Scope) error {
	return autoConvert_apiextensions_CustomResourceDefinitionNames_To_v1beta1_CustomResourceDefinitionNames(in, out, s)
}

func autoConvert_v1beta1_CustomResourceDefinitionSpec_To_apiextensions_CustomResourceDefinitionSpec(in *CustomResourceDefinitionSpec, out *apiextensions.CustomResourceDefinitionSpec, s conversion.Scope) error {
	out.Group = in.Group
	out.Version = in.Version
	if err := Convert_v1beta1_CustomResourceDefinitionNames_To_apiextensions_CustomResourceDefinitionNames(&in.Names, &out.Names, s); err != nil {
		return err
	}
	out.Scope = apiextensions.ResourceScope(in.Scope)
	if in.Validation != nil {
		in, out := &in.Validation, &out.Validation
		*out = new(apiextensions.CustomResourceValidation)
		if err := Convert_v1beta1_CustomResourceValidation_To_apiextensions_CustomResourceValidation(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Validation = nil
	}
	return nil
}

// Convert_v1beta1_CustomResourceDefinitionSpec_To_apiextensions_CustomResourceDefinitionSpec is an autogenerated conversion function.
func Convert_v1beta1_CustomResourceDefinitionSpec_To_apiextensions_CustomResourceDefinitionSpec(in *CustomResourceDefinitionSpec, out *apiextensions.CustomResourceDefinitionSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_CustomResourceDefinitionSpec_To_apiextensions_CustomResourceDefinitionSpec(in, out, s)
}

func autoConvert_apiextensions_CustomResourceDefinitionSpec_To_v1beta1_CustomResourceDefinitionSpec(in *apiextensions.CustomResourceDefinitionSpec, out *CustomResourceDefinitionSpec, s conversion.Scope) error {
	out.Group = in.Group
	out.Version = in.Version
	if err := Convert_apiextensions_CustomResourceDefinitionNames_To_v1beta1_CustomResourceDefinitionNames(&in.Names, &out.Names, s); err != nil {
		return err
	}
	out.Scope = ResourceScope(in.Scope)
	if in.Validation != nil {
		in, out := &in.Validation, &out.Validation
		*out = new(CustomResourceValidation)
		if err := Convert_apiextensions_CustomResourceValidation_To_v1beta1_CustomResourceValidation(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Validation = nil
	}
	return nil
}

// Convert_apiextensions_CustomResourceDefinitionSpec_To_v1beta1_CustomResourceDefinitionSpec is an autogenerated conversion function.
func Convert_apiextensions_CustomResourceDefinitionSpec_To_v1beta1_CustomResourceDefinitionSpec(in *apiextensions.CustomResourceDefinitionSpec, out *CustomResourceDefinitionSpec, s conversion.Scope) error {
	return autoConvert_apiextensions_CustomResourceDefinitionSpec_To_v1beta1_CustomResourceDefinitionSpec(in, out, s)
}

func autoConvert_v1beta1_CustomResourceDefinitionStatus_To_apiextensions_CustomResourceDefinitionStatus(in *CustomResourceDefinitionStatus, out *apiextensions.CustomResourceDefinitionStatus, s conversion.Scope) error {
	out.Conditions = *(*[]apiextensions.CustomResourceDefinitionCondition)(unsafe.Pointer(&in.Conditions))
	if err := Convert_v1beta1_CustomResourceDefinitionNames_To_apiextensions_CustomResourceDefinitionNames(&in.AcceptedNames, &out.AcceptedNames, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_CustomResourceDefinitionStatus_To_apiextensions_CustomResourceDefinitionStatus is an autogenerated conversion function.
func Convert_v1beta1_CustomResourceDefinitionStatus_To_apiextensions_CustomResourceDefinitionStatus(in *CustomResourceDefinitionStatus, out *apiextensions.CustomResourceDefinitionStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_CustomResourceDefinitionStatus_To_apiextensions_CustomResourceDefinitionStatus(in, out, s)
}

func autoConvert_apiextensions_CustomResourceDefinitionStatus_To_v1beta1_CustomResourceDefinitionStatus(in *apiextensions.CustomResourceDefinitionStatus, out *CustomResourceDefinitionStatus, s conversion.Scope) error {
	out.Conditions = *(*[]CustomResourceDefinitionCondition)(unsafe.Pointer(&in.Conditions))
	if err := Convert_apiextensions_CustomResourceDefinitionNames_To_v1beta1_CustomResourceDefinitionNames(&in.AcceptedNames, &out.AcceptedNames, s); err != nil {
		return err
	}
	return nil
}

// Convert_apiextensions_CustomResourceDefinitionStatus_To_v1beta1_CustomResourceDefinitionStatus is an autogenerated conversion function.
func Convert_apiextensions_CustomResourceDefinitionStatus_To_v1beta1_CustomResourceDefinitionStatus(in *apiextensions.CustomResourceDefinitionStatus, out *CustomResourceDefinitionStatus, s conversion.Scope) error {
	return autoConvert_apiextensions_CustomResourceDefinitionStatus_To_v1beta1_CustomResourceDefinitionStatus(in, out, s)
}

func autoConvert_v1beta1_CustomResourceValidation_To_apiextensions_CustomResourceValidation(in *CustomResourceValidation, out *apiextensions.CustomResourceValidation, s conversion.Scope) error {
	if in.OpenAPIV3Schema != nil {
		in, out := &in.OpenAPIV3Schema, &out.OpenAPIV3Schema
		*out = new(apiextensions.JSONSchemaProps)
		if err := Convert_v1beta1_JSONSchemaProps_To_apiextensions_JSONSchemaProps(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.OpenAPIV3Schema = nil
	}
	return nil
}

// Convert_v1beta1_CustomResourceValidation_To_apiextensions_CustomResourceValidation is an autogenerated conversion function.
func Convert_v1beta1_CustomResourceValidation_To_apiextensions_CustomResourceValidation(in *CustomResourceValidation, out *apiextensions.CustomResourceValidation, s conversion.Scope) error {
	return autoConvert_v1beta1_CustomResourceValidation_To_apiextensions_CustomResourceValidation(in, out, s)
}

func autoConvert_apiextensions_CustomResourceValidation_To_v1beta1_CustomResourceValidation(in *apiextensions.CustomResourceValidation, out *CustomResourceValidation, s conversion.Scope) error {
	if in.OpenAPIV3Schema != nil {
		in, out := &in.OpenAPIV3Schema, &out.OpenAPIV3Schema
		*out = new(JSONSchemaProps)
		if err := Convert_apiextensions_JSONSchemaProps_To_v1beta1_JSONSchemaProps(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.OpenAPIV3Schema = nil
	}
	return nil
}

// Convert_apiextensions_CustomResourceValidation_To_v1beta1_CustomResourceValidation is an autogenerated conversion function.
func Convert_apiextensions_CustomResourceValidation_To_v1beta1_CustomResourceValidation(in *apiextensions.CustomResourceValidation, out *CustomResourceValidation, s conversion.Scope) error {
	return autoConvert_apiextensions_CustomResourceValidation_To_v1beta1_CustomResourceValidation(in, out, s)
}

func autoConvert_v1beta1_ExternalDocumentation_To_apiextensions_ExternalDocumentation(in *ExternalDocumentation, out *apiextensions.ExternalDocumentation, s conversion.Scope) error {
	out.Description = in.Description
	out.URL = in.URL
	return nil
}

// Convert_v1beta1_ExternalDocumentation_To_apiextensions_ExternalDocumentation is an autogenerated conversion function.
func Convert_v1beta1_ExternalDocumentation_To_apiextensions_ExternalDocumentation(in *ExternalDocumentation, out *apiextensions.ExternalDocumentation, s conversion.Scope) error {
	return autoConvert_v1beta1_ExternalDocumentation_To_apiextensions_ExternalDocumentation(in, out, s)
}

func autoConvert_apiextensions_ExternalDocumentation_To_v1beta1_ExternalDocumentation(in *apiextensions.ExternalDocumentation, out *ExternalDocumentation, s conversion.Scope) error {
	out.Description = in.Description
	out.URL = in.URL
	return nil
}

// Convert_apiextensions_ExternalDocumentation_To_v1beta1_ExternalDocumentation is an autogenerated conversion function.
func Convert_apiextensions_ExternalDocumentation_To_v1beta1_ExternalDocumentation(in *apiextensions.ExternalDocumentation, out *ExternalDocumentation, s conversion.Scope) error {
	return autoConvert_apiextensions_ExternalDocumentation_To_v1beta1_ExternalDocumentation(in, out, s)
}

func autoConvert_v1beta1_JSON_To_apiextensions_JSON(in *JSON, out *apiextensions.JSON, s conversion.Scope) error {
	// WARNING: in.Raw requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_apiextensions_JSON_To_v1beta1_JSON(in *apiextensions.JSON, out *JSON, s conversion.Scope) error {
	// FIXME: Type apiextensions.JSON is unsupported.
	return nil
}

func autoConvert_v1beta1_JSONSchemaProps_To_apiextensions_JSONSchemaProps(in *JSONSchemaProps, out *apiextensions.JSONSchemaProps, s conversion.Scope) error {
	out.ID = in.ID
	out.Schema = apiextensions.JSONSchemaURL(in.Schema)
	out.Ref = (*string)(unsafe.Pointer(in.Ref))
	out.Description = in.Description
	out.Type = in.Type
	out.Format = in.Format
	out.Title = in.Title
	if in.Default != nil {
		in, out := &in.Default, &out.Default
		*out = new(apiextensions.JSON)
		if err := Convert_v1beta1_JSON_To_apiextensions_JSON(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Default = nil
	}
	out.Maximum = (*float64)(unsafe.Pointer(in.Maximum))
	out.ExclusiveMaximum = in.ExclusiveMaximum
	out.Minimum = (*float64)(unsafe.Pointer(in.Minimum))
	out.ExclusiveMinimum = in.ExclusiveMinimum
	out.MaxLength = (*int64)(unsafe.Pointer(in.MaxLength))
	out.MinLength = (*int64)(unsafe.Pointer(in.MinLength))
	out.Pattern = in.Pattern
	out.MaxItems = (*int64)(unsafe.Pointer(in.MaxItems))
	out.MinItems = (*int64)(unsafe.Pointer(in.MinItems))
	out.UniqueItems = in.UniqueItems
	out.MultipleOf = (*float64)(unsafe.Pointer(in.MultipleOf))
	if in.Enum != nil {
		in, out := &in.Enum, &out.Enum
		*out = make([]apiextensions.JSON, len(*in))
		for i := range *in {
			if err := Convert_v1beta1_JSON_To_apiextensions_JSON(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Enum = nil
	}
	out.MaxProperties = (*int64)(unsafe.Pointer(in.MaxProperties))
	out.MinProperties = (*int64)(unsafe.Pointer(in.MinProperties))
	out.Required = *(*[]string)(unsafe.Pointer(&in.Required))
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = new(apiextensions.JSONSchemaPropsOrArray)
		if err := Convert_v1beta1_JSONSchemaPropsOrArray_To_apiextensions_JSONSchemaPropsOrArray(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Items = nil
	}
	if in.AllOf != nil {
		in, out := &in.AllOf, &out.AllOf
		*out = make([]apiextensions.JSONSchemaProps, len(*in))
		for i := range *in {
			if err := Convert_v1beta1_JSONSchemaProps_To_apiextensions_JSONSchemaProps(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.AllOf = nil
	}
	if in.OneOf != nil {
		in, out := &in.OneOf, &out.OneOf
		*out = make([]apiextensions.JSONSchemaProps, len(*in))
		for i := range *in {
			if err := Convert_v1beta1_JSONSchemaProps_To_apiextensions_JSONSchemaProps(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.OneOf = nil
	}
	if in.AnyOf != nil {
		in, out := &in.AnyOf, &out.AnyOf
		*out = make([]apiextensions.JSONSchemaProps, len(*in))
		for i := range *in {
			if err := Convert_v1beta1_JSONSchemaProps_To_apiextensions_JSONSchemaProps(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.AnyOf = nil
	}
	if in.Not != nil {
		in, out := &in.Not, &out.Not
		*out = new(apiextensions.JSONSchemaProps)
		if err := Convert_v1beta1_JSONSchemaProps_To_apiextensions_JSONSchemaProps(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Not = nil
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make(map[string]apiextensions.JSONSchemaProps, len(*in))
		for key, val := range *in {
			newVal := new(apiextensions.JSONSchemaProps)
			if err := Convert_v1beta1_JSONSchemaProps_To_apiextensions_JSONSchemaProps(&val, newVal, s); err != nil {
				return err
			}
			(*out)[key] = *newVal
		}
	} else {
		out.Properties = nil
	}
	if in.AdditionalProperties != nil {
		in, out := &in.AdditionalProperties, &out.AdditionalProperties
		*out = new(apiextensions.JSONSchemaPropsOrBool)
		if err := Convert_v1beta1_JSONSchemaPropsOrBool_To_apiextensions_JSONSchemaPropsOrBool(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.AdditionalProperties = nil
	}
	if in.PatternProperties != nil {
		in, out := &in.PatternProperties, &out.PatternProperties
		*out = make(map[string]apiextensions.JSONSchemaProps, len(*in))
		for key, val := range *in {
			newVal := new(apiextensions.JSONSchemaProps)
			if err := Convert_v1beta1_JSONSchemaProps_To_apiextensions_JSONSchemaProps(&val, newVal, s); err != nil {
				return err
			}
			(*out)[key] = *newVal
		}
	} else {
		out.PatternProperties = nil
	}
	if in.Dependencies != nil {
		in, out := &in.Dependencies, &out.Dependencies
		*out = make(apiextensions.JSONSchemaDependencies, len(*in))
		for key, val := range *in {
			newVal := new(apiextensions.JSONSchemaPropsOrStringArray)
			if err := Convert_v1beta1_JSONSchemaPropsOrStringArray_To_apiextensions_JSONSchemaPropsOrStringArray(&val, newVal, s); err != nil {
				return err
			}
			(*out)[key] = *newVal
		}
	} else {
		out.Dependencies = nil
	}
	if in.AdditionalItems != nil {
		in, out := &in.AdditionalItems, &out.AdditionalItems
		*out = new(apiextensions.JSONSchemaPropsOrBool)
		if err := Convert_v1beta1_JSONSchemaPropsOrBool_To_apiextensions_JSONSchemaPropsOrBool(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.AdditionalItems = nil
	}
	if in.Definitions != nil {
		in, out := &in.Definitions, &out.Definitions
		*out = make(apiextensions.JSONSchemaDefinitions, len(*in))
		for key, val := range *in {
			newVal := new(apiextensions.JSONSchemaProps)
			if err := Convert_v1beta1_JSONSchemaProps_To_apiextensions_JSONSchemaProps(&val, newVal, s); err != nil {
				return err
			}
			(*out)[key] = *newVal
		}
	} else {
		out.Definitions = nil
	}
	out.ExternalDocs = (*apiextensions.ExternalDocumentation)(unsafe.Pointer(in.ExternalDocs))
	if in.Example != nil {
		in, out := &in.Example, &out.Example
		*out = new(apiextensions.JSON)
		if err := Convert_v1beta1_JSON_To_apiextensions_JSON(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Example = nil
	}
	return nil
}

// Convert_v1beta1_JSONSchemaProps_To_apiextensions_JSONSchemaProps is an autogenerated conversion function.
func Convert_v1beta1_JSONSchemaProps_To_apiextensions_JSONSchemaProps(in *JSONSchemaProps, out *apiextensions.JSONSchemaProps, s conversion.Scope) error {
	return autoConvert_v1beta1_JSONSchemaProps_To_apiextensions_JSONSchemaProps(in, out, s)
}

func autoConvert_apiextensions_JSONSchemaProps_To_v1beta1_JSONSchemaProps(in *apiextensions.JSONSchemaProps, out *JSONSchemaProps, s conversion.Scope) error {
	out.ID = in.ID
	out.Schema = JSONSchemaURL(in.Schema)
	out.Ref = (*string)(unsafe.Pointer(in.Ref))
	out.Description = in.Description
	out.Type = in.Type
	out.Format = in.Format
	out.Title = in.Title
	if in.Default != nil {
		in, out := &in.Default, &out.Default
		*out = new(JSON)
		if err := Convert_apiextensions_JSON_To_v1beta1_JSON(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Default = nil
	}
	out.Maximum = (*float64)(unsafe.Pointer(in.Maximum))
	out.ExclusiveMaximum = in.ExclusiveMaximum
	out.Minimum = (*float64)(unsafe.Pointer(in.Minimum))
	out.ExclusiveMinimum = in.ExclusiveMinimum
	out.MaxLength = (*int64)(unsafe.Pointer(in.MaxLength))
	out.MinLength = (*int64)(unsafe.Pointer(in.MinLength))
	out.Pattern = in.Pattern
	out.MaxItems = (*int64)(unsafe.Pointer(in.MaxItems))
	out.MinItems = (*int64)(unsafe.Pointer(in.MinItems))
	out.UniqueItems = in.UniqueItems
	out.MultipleOf = (*float64)(unsafe.Pointer(in.MultipleOf))
	if in.Enum != nil {
		in, out := &in.Enum, &out.Enum
		*out = make([]JSON, len(*in))
		for i := range *in {
			if err := Convert_apiextensions_JSON_To_v1beta1_JSON(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Enum = nil
	}
	out.MaxProperties = (*int64)(unsafe.Pointer(in.MaxProperties))
	out.MinProperties = (*int64)(unsafe.Pointer(in.MinProperties))
	out.Required = *(*[]string)(unsafe.Pointer(&in.Required))
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = new(JSONSchemaPropsOrArray)
		if err := Convert_apiextensions_JSONSchemaPropsOrArray_To_v1beta1_JSONSchemaPropsOrArray(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Items = nil
	}
	if in.AllOf != nil {
		in, out := &in.AllOf, &out.AllOf
		*out = make([]JSONSchemaProps, len(*in))
		for i := range *in {
			if err := Convert_apiextensions_JSONSchemaProps_To_v1beta1_JSONSchemaProps(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.AllOf = nil
	}
	if in.OneOf != nil {
		in, out := &in.OneOf, &out.OneOf
		*out = make([]JSONSchemaProps, len(*in))
		for i := range *in {
			if err := Convert_apiextensions_JSONSchemaProps_To_v1beta1_JSONSchemaProps(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.OneOf = nil
	}
	if in.AnyOf != nil {
		in, out := &in.AnyOf, &out.AnyOf
		*out = make([]JSONSchemaProps, len(*in))
		for i := range *in {
			if err := Convert_apiextensions_JSONSchemaProps_To_v1beta1_JSONSchemaProps(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.AnyOf = nil
	}
	if in.Not != nil {
		in, out := &in.Not, &out.Not
		*out = new(JSONSchemaProps)
		if err := Convert_apiextensions_JSONSchemaProps_To_v1beta1_JSONSchemaProps(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Not = nil
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make(map[string]JSONSchemaProps, len(*in))
		for key, val := range *in {
			newVal := new(JSONSchemaProps)
			if err := Convert_apiextensions_JSONSchemaProps_To_v1beta1_JSONSchemaProps(&val, newVal, s); err != nil {
				return err
			}
			(*out)[key] = *newVal
		}
	} else {
		out.Properties = nil
	}
	if in.AdditionalProperties != nil {
		in, out := &in.AdditionalProperties, &out.AdditionalProperties
		*out = new(JSONSchemaPropsOrBool)
		if err := Convert_apiextensions_JSONSchemaPropsOrBool_To_v1beta1_JSONSchemaPropsOrBool(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.AdditionalProperties = nil
	}
	if in.PatternProperties != nil {
		in, out := &in.PatternProperties, &out.PatternProperties
		*out = make(map[string]JSONSchemaProps, len(*in))
		for key, val := range *in {
			newVal := new(JSONSchemaProps)
			if err := Convert_apiextensions_JSONSchemaProps_To_v1beta1_JSONSchemaProps(&val, newVal, s); err != nil {
				return err
			}
			(*out)[key] = *newVal
		}
	} else {
		out.PatternProperties = nil
	}
	if in.Dependencies != nil {
		in, out := &in.Dependencies, &out.Dependencies
		*out = make(JSONSchemaDependencies, len(*in))
		for key, val := range *in {
			newVal := new(JSONSchemaPropsOrStringArray)
			if err := Convert_apiextensions_JSONSchemaPropsOrStringArray_To_v1beta1_JSONSchemaPropsOrStringArray(&val, newVal, s); err != nil {
				return err
			}
			(*out)[key] = *newVal
		}
	} else {
		out.Dependencies = nil
	}
	if in.AdditionalItems != nil {
		in, out := &in.AdditionalItems, &out.AdditionalItems
		*out = new(JSONSchemaPropsOrBool)
		if err := Convert_apiextensions_JSONSchemaPropsOrBool_To_v1beta1_JSONSchemaPropsOrBool(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.AdditionalItems = nil
	}
	if in.Definitions != nil {
		in, out := &in.Definitions, &out.Definitions
		*out = make(JSONSchemaDefinitions, len(*in))
		for key, val := range *in {
			newVal := new(JSONSchemaProps)
			if err := Convert_apiextensions_JSONSchemaProps_To_v1beta1_JSONSchemaProps(&val, newVal, s); err != nil {
				return err
			}
			(*out)[key] = *newVal
		}
	} else {
		out.Definitions = nil
	}
	out.ExternalDocs = (*ExternalDocumentation)(unsafe.Pointer(in.ExternalDocs))
	if in.Example != nil {
		in, out := &in.Example, &out.Example
		*out = new(JSON)
		if err := Convert_apiextensions_JSON_To_v1beta1_JSON(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Example = nil
	}
	return nil
}

func autoConvert_v1beta1_JSONSchemaPropsOrArray_To_apiextensions_JSONSchemaPropsOrArray(in *JSONSchemaPropsOrArray, out *apiextensions.JSONSchemaPropsOrArray, s conversion.Scope) error {
	if in.Schema != nil {
		in, out := &in.Schema, &out.Schema
		*out = new(apiextensions.JSONSchemaProps)
		if err := Convert_v1beta1_JSONSchemaProps_To_apiextensions_JSONSchemaProps(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Schema = nil
	}
	if in.JSONSchemas != nil {
		in, out := &in.JSONSchemas, &out.JSONSchemas
		*out = make([]apiextensions.JSONSchemaProps, len(*in))
		for i := range *in {
			if err := Convert_v1beta1_JSONSchemaProps_To_apiextensions_JSONSchemaProps(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.JSONSchemas = nil
	}
	return nil
}

// Convert_v1beta1_JSONSchemaPropsOrArray_To_apiextensions_JSONSchemaPropsOrArray is an autogenerated conversion function.
func Convert_v1beta1_JSONSchemaPropsOrArray_To_apiextensions_JSONSchemaPropsOrArray(in *JSONSchemaPropsOrArray, out *apiextensions.JSONSchemaPropsOrArray, s conversion.Scope) error {
	return autoConvert_v1beta1_JSONSchemaPropsOrArray_To_apiextensions_JSONSchemaPropsOrArray(in, out, s)
}

func autoConvert_apiextensions_JSONSchemaPropsOrArray_To_v1beta1_JSONSchemaPropsOrArray(in *apiextensions.JSONSchemaPropsOrArray, out *JSONSchemaPropsOrArray, s conversion.Scope) error {
	if in.Schema != nil {
		in, out := &in.Schema, &out.Schema
		*out = new(JSONSchemaProps)
		if err := Convert_apiextensions_JSONSchemaProps_To_v1beta1_JSONSchemaProps(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Schema = nil
	}
	if in.JSONSchemas != nil {
		in, out := &in.JSONSchemas, &out.JSONSchemas
		*out = make([]JSONSchemaProps, len(*in))
		for i := range *in {
			if err := Convert_apiextensions_JSONSchemaProps_To_v1beta1_JSONSchemaProps(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.JSONSchemas = nil
	}
	return nil
}

// Convert_apiextensions_JSONSchemaPropsOrArray_To_v1beta1_JSONSchemaPropsOrArray is an autogenerated conversion function.
func Convert_apiextensions_JSONSchemaPropsOrArray_To_v1beta1_JSONSchemaPropsOrArray(in *apiextensions.JSONSchemaPropsOrArray, out *JSONSchemaPropsOrArray, s conversion.Scope) error {
	return autoConvert_apiextensions_JSONSchemaPropsOrArray_To_v1beta1_JSONSchemaPropsOrArray(in, out, s)
}

func autoConvert_v1beta1_JSONSchemaPropsOrBool_To_apiextensions_JSONSchemaPropsOrBool(in *JSONSchemaPropsOrBool, out *apiextensions.JSONSchemaPropsOrBool, s conversion.Scope) error {
	out.Allows = in.Allows
	if in.Schema != nil {
		in, out := &in.Schema, &out.Schema
		*out = new(apiextensions.JSONSchemaProps)
		if err := Convert_v1beta1_JSONSchemaProps_To_apiextensions_JSONSchemaProps(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Schema = nil
	}
	return nil
}

// Convert_v1beta1_JSONSchemaPropsOrBool_To_apiextensions_JSONSchemaPropsOrBool is an autogenerated conversion function.
func Convert_v1beta1_JSONSchemaPropsOrBool_To_apiextensions_JSONSchemaPropsOrBool(in *JSONSchemaPropsOrBool, out *apiextensions.JSONSchemaPropsOrBool, s conversion.Scope) error {
	return autoConvert_v1beta1_JSONSchemaPropsOrBool_To_apiextensions_JSONSchemaPropsOrBool(in, out, s)
}

func autoConvert_apiextensions_JSONSchemaPropsOrBool_To_v1beta1_JSONSchemaPropsOrBool(in *apiextensions.JSONSchemaPropsOrBool, out *JSONSchemaPropsOrBool, s conversion.Scope) error {
	out.Allows = in.Allows
	if in.Schema != nil {
		in, out := &in.Schema, &out.Schema
		*out = new(JSONSchemaProps)
		if err := Convert_apiextensions_JSONSchemaProps_To_v1beta1_JSONSchemaProps(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Schema = nil
	}
	return nil
}

// Convert_apiextensions_JSONSchemaPropsOrBool_To_v1beta1_JSONSchemaPropsOrBool is an autogenerated conversion function.
func Convert_apiextensions_JSONSchemaPropsOrBool_To_v1beta1_JSONSchemaPropsOrBool(in *apiextensions.JSONSchemaPropsOrBool, out *JSONSchemaPropsOrBool, s conversion.Scope) error {
	return autoConvert_apiextensions_JSONSchemaPropsOrBool_To_v1beta1_JSONSchemaPropsOrBool(in, out, s)
}

func autoConvert_v1beta1_JSONSchemaPropsOrStringArray_To_apiextensions_JSONSchemaPropsOrStringArray(in *JSONSchemaPropsOrStringArray, out *apiextensions.JSONSchemaPropsOrStringArray, s conversion.Scope) error {
	if in.Schema != nil {
		in, out := &in.Schema, &out.Schema
		*out = new(apiextensions.JSONSchemaProps)
		if err := Convert_v1beta1_JSONSchemaProps_To_apiextensions_JSONSchemaProps(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Schema = nil
	}
	out.Property = *(*[]string)(unsafe.Pointer(&in.Property))
	return nil
}

// Convert_v1beta1_JSONSchemaPropsOrStringArray_To_apiextensions_JSONSchemaPropsOrStringArray is an autogenerated conversion function.
func Convert_v1beta1_JSONSchemaPropsOrStringArray_To_apiextensions_JSONSchemaPropsOrStringArray(in *JSONSchemaPropsOrStringArray, out *apiextensions.JSONSchemaPropsOrStringArray, s conversion.Scope) error {
	return autoConvert_v1beta1_JSONSchemaPropsOrStringArray_To_apiextensions_JSONSchemaPropsOrStringArray(in, out, s)
}

func autoConvert_apiextensions_JSONSchemaPropsOrStringArray_To_v1beta1_JSONSchemaPropsOrStringArray(in *apiextensions.JSONSchemaPropsOrStringArray, out *JSONSchemaPropsOrStringArray, s conversion.Scope) error {
	if in.Schema != nil {
		in, out := &in.Schema, &out.Schema
		*out = new(JSONSchemaProps)
		if err := Convert_apiextensions_JSONSchemaProps_To_v1beta1_JSONSchemaProps(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Schema = nil
	}
	out.Property = *(*[]string)(unsafe.Pointer(&in.Property))
	return nil
}

// Convert_apiextensions_JSONSchemaPropsOrStringArray_To_v1beta1_JSONSchemaPropsOrStringArray is an autogenerated conversion function.
func Convert_apiextensions_JSONSchemaPropsOrStringArray_To_v1beta1_JSONSchemaPropsOrStringArray(in *apiextensions.JSONSchemaPropsOrStringArray, out *JSONSchemaPropsOrStringArray, s conversion.Scope) error {
	return autoConvert_apiextensions_JSONSchemaPropsOrStringArray_To_v1beta1_JSONSchemaPropsOrStringArray(in, out, s)
}

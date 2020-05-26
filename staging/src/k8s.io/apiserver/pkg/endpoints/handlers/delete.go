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

package handlers

import (
	"context"
	"fmt"
	"k8s.io/klog/v2"
	"net/http"
	"time"

	"k8s.io/apimachinery/pkg/api/errors"
	metainternalversion "k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metainternalversionscheme "k8s.io/apimachinery/pkg/apis/meta/internalversion/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/validation"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/admission"
	"k8s.io/apiserver/pkg/audit"
	"k8s.io/apiserver/pkg/endpoints/handlers/negotiation"
	"k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/features"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/apiserver/pkg/util/dryrun"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	utiltrace "k8s.io/utils/trace"
)

// DeleteResource returns a function that will handle a resource deletion
// TODO admission here becomes solely validating admission
func DeleteResource(r rest.GracefulDeleter, allowsOptions bool, scope *RequestScope, admit admission.Interface) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// For performance tracking purposes.
		trace := utiltrace.New("Delete", utiltrace.Field{Key: "url", Value: req.URL.Path}, utiltrace.Field{Key: "user-agent", Value: &lazyTruncatedUserAgent{req}}, utiltrace.Field{Key: "client", Value: &lazyClientIP{req}})
		defer trace.LogIfLong(500 * time.Millisecond)

		if isDryRun(req.URL) && !utilfeature.DefaultFeatureGate.Enabled(features.DryRun) {
			scope.err(errors.NewBadRequest("the dryRun feature is disabled"), w, req)
			return
		}

		//// TODO: we either want to remove timeout or document it (if we document, move timeout out of this function and declare it in api_installer)
		//timeout := parseTimeout(req.URL.Query().Get("timeout"))
		timeout := 90 * time.Second

		namespace, name, err := scope.Namer.Name(req)
		if err != nil {
			scope.err(err, w, req)
			return
		}
		ctx, cancel := context.WithTimeout(req.Context(), timeout)
		defer cancel()
		ctx = request.WithNamespace(ctx, namespace)
		ae := request.AuditEventFrom(ctx)
		admit = admission.WithAudit(admit, ae)

		outputMediaType, _, err := negotiation.NegotiateOutputMediaType(req, scope.Serializer, scope)
		if err != nil {
			scope.err(err, w, req)
			return
		}

		options := &metav1.DeleteOptions{}
		if allowsOptions {
			body, err := limitedReadBody(req, scope.MaxRequestBodyBytes)
			if err != nil {
				scope.err(err, w, req)
				return
			}
			if len(body) > 0 {
				s, err := negotiation.NegotiateInputSerializer(req, false, metainternalversionscheme.Codecs)
				if err != nil {
					scope.err(err, w, req)
					return
				}
				// For backwards compatibility, we need to allow existing clients to submit per group DeleteOptions
				// It is also allowed to pass a body with meta.k8s.io/v1.DeleteOptions
				defaultGVK := scope.MetaGroupVersion.WithKind("DeleteOptions")
				obj, _, err := metainternalversionscheme.Codecs.DecoderToVersion(s.Serializer, defaultGVK.GroupVersion()).Decode(body, &defaultGVK, options)
				if err != nil {
					scope.err(err, w, req)
					return
				}
				if obj != options {
					scope.err(fmt.Errorf("decoded object cannot be converted to DeleteOptions"), w, req)
					return
				}
				trace.Step("Decoded delete options")

				ae := request.AuditEventFrom(ctx)
				audit.LogRequestObject(ae, obj, scope.Resource, scope.Subresource, scope.Serializer)
				trace.Step("Recorded the audit event")
			} else {
				if err := metainternalversionscheme.ParameterCodec.DecodeParameters(req.URL.Query(), scope.MetaGroupVersion, options); err != nil {
					err = errors.NewBadRequest(err.Error())
					scope.err(err, w, req)
					return
				}
			}
		}
		if errs := validation.ValidateDeleteOptions(options); len(errs) > 0 {
			err := errors.NewInvalid(schema.GroupKind{Group: metav1.GroupName, Kind: "DeleteOptions"}, "", errs)
			scope.err(err, w, req)
			return
		}
		options.TypeMeta.SetGroupVersionKind(metav1.SchemeGroupVersion.WithKind("DeleteOptions"))

		trace.Step("About to delete object from database")
		wasDeleted := true
		userInfo, _ := request.UserFrom(ctx)
		staticAdmissionAttrs := admission.NewAttributesRecord(nil, nil, scope.Kind, namespace, name, scope.Resource, scope.Subresource, admission.Delete, options, dryrun.IsDryRun(options.DryRun), userInfo)
		result, err := finishRequest(timeout, func() (runtime.Object, error) {
			obj, deleted, err := r.Delete(ctx, name, rest.AdmissionToValidateObjectDeleteFunc(admit, staticAdmissionAttrs, scope), options)
			wasDeleted = deleted
			return obj, err
		})
		if err != nil {
			scope.err(err, w, req)
			return
		}
		trace.Step("Object deleted from database")

		status := http.StatusOK
		// Return http.StatusAccepted if the resource was not deleted immediately and
		// user requested cascading deletion by setting OrphanDependents=false.
		// Note: We want to do this always if resource was not deleted immediately, but
		// that will break existing clients.
		// Other cases where resource is not instantly deleted are: namespace deletion
		// and pod graceful deletion.
		if !wasDeleted && options.OrphanDependents != nil && *options.OrphanDependents == false {
			status = http.StatusAccepted
		}
		// if the rest.Deleter returns a nil object, fill out a status. Callers may return a valid
		// object with the response.
		if result == nil {
			result = &metav1.Status{
				Status: metav1.StatusSuccess,
				Code:   int32(status),
				Details: &metav1.StatusDetails{
					Name: name,
					Kind: scope.Kind.Kind,
				},
			}
		}

		transformResponseObject(ctx, scope, trace, req, w, status, outputMediaType, result)
	}
}

// DeleteCollection returns a function that will handle a collection deletion
func DeleteCollection(r rest.CollectionDeleter, checkBody bool, scope *RequestScope, admit admission.Interface) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		startTime := time.Now()
		trace := utiltrace.New("Delete", utiltrace.Field{"url", req.URL.Path})
		defer trace.LogIfLong(500 * time.Millisecond)
		klog.V(1).Infof("***BP*** Entering DeleteCollection for %s", req.URL.Path)
		defer func() {
			klog.V(1).Infof("***BP*** Leaving DeleteCollection for %s (Duration %v)", req.URL.Path, time.Since(startTime))
		}()
		if isDryRun(req.URL) && !utilfeature.DefaultFeatureGate.Enabled(features.DryRun) {
			scope.err(errors.NewBadRequest("the dryRun feature is disabled"), w, req)
			return
		}

		// TODO: we either want to remove timeout or document it (if we document, move timeout out of this function and declare it in api_installer)
		timeout := parseTimeout(req.URL.Query().Get("timeout"))

		namespace, err := scope.Namer.Namespace(req)
		if err != nil {
			klog.V(1).Infof("***BP*** DeleteCollection failed to get namespace for %s: %v", req.URL.Path, err)
			scope.err(err, w, req)
			return
		}
		klog.V(1).Infof("***BP*** DeleteCollection namespace for %s: %v", req.URL.Path, namespace)

		ctx, cancel := context.WithTimeout(req.Context(), timeout)
		defer cancel()
		ctx = request.WithNamespace(ctx, namespace)
		ae := request.AuditEventFrom(ctx)

		outputMediaType, _, err := negotiation.NegotiateOutputMediaType(req, scope.Serializer, scope)
		if err != nil {
			klog.V(1).Infof("***BP*** DeleteCollection failed to NegotiateOutputMediaType for %s: %v", req.URL.Path, err)
			scope.err(err, w, req)
			return
		}
		klog.V(1).Infof("***BP*** DeleteCollection outputMediaType for %s: %v", req.URL.Path, outputMediaType)

		listOptions := metainternalversion.ListOptions{}
		if err := metainternalversionscheme.ParameterCodec.DecodeParameters(req.URL.Query(), scope.MetaGroupVersion, &listOptions); err != nil {
			klog.V(1).Infof("***BP*** DeleteCollection failed to DecodeParameters for %s: %v", req.URL.Path, err)
			err = errors.NewBadRequest(err.Error())
			scope.err(err, w, req)
			return
		}

		// transform fields
		// TODO: DecodeParametersInto should do this.
		if listOptions.FieldSelector != nil {
			fn := func(label, value string) (newLabel, newValue string, err error) {
				return scope.Convertor.ConvertFieldLabel(scope.Kind, label, value)
			}
			if listOptions.FieldSelector, err = listOptions.FieldSelector.Transform(fn); err != nil {
				klog.V(1).Infof("***BP*** DeleteCollection failed to transform field selector for %s: %v", req.URL.Path, err)
				// TODO: allow bad request to set field causes based on query parameters
				err = errors.NewBadRequest(err.Error())
				scope.err(err, w, req)
				return
			}
		}

		options := &metav1.DeleteOptions{}
		if checkBody {
			body, err := limitedReadBody(req, scope.MaxRequestBodyBytes)
			if err != nil {
				klog.V(1).Infof("***BP*** DeleteCollection failed to read body for %s: %v", req.URL.Path, err)
				scope.err(err, w, req)
				return
			}
			if len(body) > 0 {
				s, err := negotiation.NegotiateInputSerializer(req, false, scope.Serializer)
				if err != nil {
					klog.V(1).Infof("***BP*** DeleteCollection failed to negotiate input serializer for %s: %v", req.URL.Path, err)
					scope.err(err, w, req)
					return
				}
				// For backwards compatibility, we need to allow existing clients to submit per group DeleteOptions
				// It is also allowed to pass a body with meta.k8s.io/v1.DeleteOptions
				defaultGVK := scope.Kind.GroupVersion().WithKind("DeleteOptions")
				obj, _, err := scope.Serializer.DecoderToVersion(s.Serializer, defaultGVK.GroupVersion()).Decode(body, &defaultGVK, options)
				if err != nil {
					klog.V(1).Infof("***BP*** DeleteCollection failed to DecoderToVersion for %s: %v", req.URL.Path, err)
					scope.err(err, w, req)
					return
				}
				if obj != options {
					klog.V(1).Infof("***BP*** decoded object cannot be converted to DeleteOptions for %s: %v", req.URL.Path, err)
					scope.err(fmt.Errorf("decoded object cannot be converted to DeleteOptions"), w, req)
					return
				}

				ae := request.AuditEventFrom(ctx)
				audit.LogRequestObject(ae, obj, scope.Resource, scope.Subresource, scope.Serializer)
			} else {
				if err := metainternalversionscheme.ParameterCodec.DecodeParameters(req.URL.Query(), scope.MetaGroupVersion, options); err != nil {
					klog.V(1).Infof("***BP*** DeleteCollection failed to Decode parameters for %s: %v", req.URL.Path, err)
					err = errors.NewBadRequest(err.Error())
					scope.err(err, w, req)
					return
				}
			}
		}
		if errs := validation.ValidateDeleteOptions(options); len(errs) > 0 {
			klog.V(1).Infof("***BP*** DeleteCollection failed to validate delete options for %s: %v", req.URL.Path, err)
			err := errors.NewInvalid(schema.GroupKind{Group: metav1.GroupName, Kind: "DeleteOptions"}, "", errs)
			scope.err(err, w, req)
			return
		}
		options.TypeMeta.SetGroupVersionKind(metav1.SchemeGroupVersion.WithKind("DeleteOptions"))

		admit = admission.WithAudit(admit, ae)
		userInfo, _ := request.UserFrom(ctx)
		staticAdmissionAttrs := admission.NewAttributesRecord(nil, nil, scope.Kind, namespace, "", scope.Resource, scope.Subresource, admission.Delete, options, dryrun.IsDryRun(options.DryRun), userInfo)
		result, err := finishRequest(timeout, func() (runtime.Object, error) {
			klog.V(1).Info("***BP*** Before r.DeleteCollection")
			startTime := time.Now()
			obj, err := r.DeleteCollection(ctx, rest.AdmissionToValidateObjectDeleteFunc(admit, staticAdmissionAttrs, scope), options, &listOptions)
			klog.V(1).Infof("***BP*** After r.DeleteCollection, err = %v, duration = %v", err, time.Since(startTime))
			return obj, err
		})
		if err != nil {
			klog.V(1).Infof("***BP*** DeleteCollection failed to finish request for %s: %v", req.URL.Path, err)
			scope.err(err, w, req)
			return
		}

		// if the rest.Deleter returns a nil object, fill out a status. Callers may return a valid
		// object with the response.
		if result == nil {
			result = &metav1.Status{
				Status: metav1.StatusSuccess,
				Code:   http.StatusOK,
				Details: &metav1.StatusDetails{
					Kind: scope.Kind.Kind,
				},
			}
		}

		transformResponseObject(ctx, scope, trace, req, w, http.StatusOK, outputMediaType, result)
		klog.V(1).Infof("***BP*** DeleteCollection reached the end for %s", req.URL.Path)
	}
}

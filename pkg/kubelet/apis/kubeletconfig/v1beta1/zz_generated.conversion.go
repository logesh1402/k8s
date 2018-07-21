// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

package v1beta1

import (
	unsafe "unsafe"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	kubeletconfig "k8s.io/kubernetes/pkg/kubelet/apis/kubeletconfig"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*KubeletAnonymousAuthentication)(nil), (*kubeletconfig.KubeletAnonymousAuthentication)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_KubeletAnonymousAuthentication_To_kubeletconfig_KubeletAnonymousAuthentication(a.(*KubeletAnonymousAuthentication), b.(*kubeletconfig.KubeletAnonymousAuthentication), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*kubeletconfig.KubeletAnonymousAuthentication)(nil), (*KubeletAnonymousAuthentication)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_kubeletconfig_KubeletAnonymousAuthentication_To_v1beta1_KubeletAnonymousAuthentication(a.(*kubeletconfig.KubeletAnonymousAuthentication), b.(*KubeletAnonymousAuthentication), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*KubeletAuthentication)(nil), (*kubeletconfig.KubeletAuthentication)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_KubeletAuthentication_To_kubeletconfig_KubeletAuthentication(a.(*KubeletAuthentication), b.(*kubeletconfig.KubeletAuthentication), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*kubeletconfig.KubeletAuthentication)(nil), (*KubeletAuthentication)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_kubeletconfig_KubeletAuthentication_To_v1beta1_KubeletAuthentication(a.(*kubeletconfig.KubeletAuthentication), b.(*KubeletAuthentication), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*KubeletAuthorization)(nil), (*kubeletconfig.KubeletAuthorization)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_KubeletAuthorization_To_kubeletconfig_KubeletAuthorization(a.(*KubeletAuthorization), b.(*kubeletconfig.KubeletAuthorization), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*kubeletconfig.KubeletAuthorization)(nil), (*KubeletAuthorization)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_kubeletconfig_KubeletAuthorization_To_v1beta1_KubeletAuthorization(a.(*kubeletconfig.KubeletAuthorization), b.(*KubeletAuthorization), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*KubeletConfiguration)(nil), (*kubeletconfig.KubeletConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_KubeletConfiguration_To_kubeletconfig_KubeletConfiguration(a.(*KubeletConfiguration), b.(*kubeletconfig.KubeletConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*kubeletconfig.KubeletConfiguration)(nil), (*KubeletConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_kubeletconfig_KubeletConfiguration_To_v1beta1_KubeletConfiguration(a.(*kubeletconfig.KubeletConfiguration), b.(*KubeletConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*KubeletWebhookAuthentication)(nil), (*kubeletconfig.KubeletWebhookAuthentication)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_KubeletWebhookAuthentication_To_kubeletconfig_KubeletWebhookAuthentication(a.(*KubeletWebhookAuthentication), b.(*kubeletconfig.KubeletWebhookAuthentication), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*kubeletconfig.KubeletWebhookAuthentication)(nil), (*KubeletWebhookAuthentication)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_kubeletconfig_KubeletWebhookAuthentication_To_v1beta1_KubeletWebhookAuthentication(a.(*kubeletconfig.KubeletWebhookAuthentication), b.(*KubeletWebhookAuthentication), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*KubeletWebhookAuthorization)(nil), (*kubeletconfig.KubeletWebhookAuthorization)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_KubeletWebhookAuthorization_To_kubeletconfig_KubeletWebhookAuthorization(a.(*KubeletWebhookAuthorization), b.(*kubeletconfig.KubeletWebhookAuthorization), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*kubeletconfig.KubeletWebhookAuthorization)(nil), (*KubeletWebhookAuthorization)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_kubeletconfig_KubeletWebhookAuthorization_To_v1beta1_KubeletWebhookAuthorization(a.(*kubeletconfig.KubeletWebhookAuthorization), b.(*KubeletWebhookAuthorization), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*KubeletX509Authentication)(nil), (*kubeletconfig.KubeletX509Authentication)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_KubeletX509Authentication_To_kubeletconfig_KubeletX509Authentication(a.(*KubeletX509Authentication), b.(*kubeletconfig.KubeletX509Authentication), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*kubeletconfig.KubeletX509Authentication)(nil), (*KubeletX509Authentication)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_kubeletconfig_KubeletX509Authentication_To_v1beta1_KubeletX509Authentication(a.(*kubeletconfig.KubeletX509Authentication), b.(*KubeletX509Authentication), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*SerializedNodeConfigSource)(nil), (*kubeletconfig.SerializedNodeConfigSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_SerializedNodeConfigSource_To_kubeletconfig_SerializedNodeConfigSource(a.(*SerializedNodeConfigSource), b.(*kubeletconfig.SerializedNodeConfigSource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*kubeletconfig.SerializedNodeConfigSource)(nil), (*SerializedNodeConfigSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_kubeletconfig_SerializedNodeConfigSource_To_v1beta1_SerializedNodeConfigSource(a.(*kubeletconfig.SerializedNodeConfigSource), b.(*SerializedNodeConfigSource), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1beta1_KubeletAnonymousAuthentication_To_kubeletconfig_KubeletAnonymousAuthentication(in *KubeletAnonymousAuthentication, out *kubeletconfig.KubeletAnonymousAuthentication, s conversion.Scope) error {
	if err := v1.Convert_Pointer_bool_To_bool(&in.Enabled, &out.Enabled, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_KubeletAnonymousAuthentication_To_kubeletconfig_KubeletAnonymousAuthentication is an autogenerated conversion function.
func Convert_v1beta1_KubeletAnonymousAuthentication_To_kubeletconfig_KubeletAnonymousAuthentication(in *KubeletAnonymousAuthentication, out *kubeletconfig.KubeletAnonymousAuthentication, s conversion.Scope) error {
	return autoConvert_v1beta1_KubeletAnonymousAuthentication_To_kubeletconfig_KubeletAnonymousAuthentication(in, out, s)
}

func autoConvert_kubeletconfig_KubeletAnonymousAuthentication_To_v1beta1_KubeletAnonymousAuthentication(in *kubeletconfig.KubeletAnonymousAuthentication, out *KubeletAnonymousAuthentication, s conversion.Scope) error {
	if err := v1.Convert_bool_To_Pointer_bool(&in.Enabled, &out.Enabled, s); err != nil {
		return err
	}
	return nil
}

// Convert_kubeletconfig_KubeletAnonymousAuthentication_To_v1beta1_KubeletAnonymousAuthentication is an autogenerated conversion function.
func Convert_kubeletconfig_KubeletAnonymousAuthentication_To_v1beta1_KubeletAnonymousAuthentication(in *kubeletconfig.KubeletAnonymousAuthentication, out *KubeletAnonymousAuthentication, s conversion.Scope) error {
	return autoConvert_kubeletconfig_KubeletAnonymousAuthentication_To_v1beta1_KubeletAnonymousAuthentication(in, out, s)
}

func autoConvert_v1beta1_KubeletAuthentication_To_kubeletconfig_KubeletAuthentication(in *KubeletAuthentication, out *kubeletconfig.KubeletAuthentication, s conversion.Scope) error {
	if err := Convert_v1beta1_KubeletX509Authentication_To_kubeletconfig_KubeletX509Authentication(&in.X509, &out.X509, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_KubeletWebhookAuthentication_To_kubeletconfig_KubeletWebhookAuthentication(&in.Webhook, &out.Webhook, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_KubeletAnonymousAuthentication_To_kubeletconfig_KubeletAnonymousAuthentication(&in.Anonymous, &out.Anonymous, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_KubeletAuthentication_To_kubeletconfig_KubeletAuthentication is an autogenerated conversion function.
func Convert_v1beta1_KubeletAuthentication_To_kubeletconfig_KubeletAuthentication(in *KubeletAuthentication, out *kubeletconfig.KubeletAuthentication, s conversion.Scope) error {
	return autoConvert_v1beta1_KubeletAuthentication_To_kubeletconfig_KubeletAuthentication(in, out, s)
}

func autoConvert_kubeletconfig_KubeletAuthentication_To_v1beta1_KubeletAuthentication(in *kubeletconfig.KubeletAuthentication, out *KubeletAuthentication, s conversion.Scope) error {
	if err := Convert_kubeletconfig_KubeletX509Authentication_To_v1beta1_KubeletX509Authentication(&in.X509, &out.X509, s); err != nil {
		return err
	}
	if err := Convert_kubeletconfig_KubeletWebhookAuthentication_To_v1beta1_KubeletWebhookAuthentication(&in.Webhook, &out.Webhook, s); err != nil {
		return err
	}
	if err := Convert_kubeletconfig_KubeletAnonymousAuthentication_To_v1beta1_KubeletAnonymousAuthentication(&in.Anonymous, &out.Anonymous, s); err != nil {
		return err
	}
	return nil
}

// Convert_kubeletconfig_KubeletAuthentication_To_v1beta1_KubeletAuthentication is an autogenerated conversion function.
func Convert_kubeletconfig_KubeletAuthentication_To_v1beta1_KubeletAuthentication(in *kubeletconfig.KubeletAuthentication, out *KubeletAuthentication, s conversion.Scope) error {
	return autoConvert_kubeletconfig_KubeletAuthentication_To_v1beta1_KubeletAuthentication(in, out, s)
}

func autoConvert_v1beta1_KubeletAuthorization_To_kubeletconfig_KubeletAuthorization(in *KubeletAuthorization, out *kubeletconfig.KubeletAuthorization, s conversion.Scope) error {
	out.Mode = kubeletconfig.KubeletAuthorizationMode(in.Mode)
	if err := Convert_v1beta1_KubeletWebhookAuthorization_To_kubeletconfig_KubeletWebhookAuthorization(&in.Webhook, &out.Webhook, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_KubeletAuthorization_To_kubeletconfig_KubeletAuthorization is an autogenerated conversion function.
func Convert_v1beta1_KubeletAuthorization_To_kubeletconfig_KubeletAuthorization(in *KubeletAuthorization, out *kubeletconfig.KubeletAuthorization, s conversion.Scope) error {
	return autoConvert_v1beta1_KubeletAuthorization_To_kubeletconfig_KubeletAuthorization(in, out, s)
}

func autoConvert_kubeletconfig_KubeletAuthorization_To_v1beta1_KubeletAuthorization(in *kubeletconfig.KubeletAuthorization, out *KubeletAuthorization, s conversion.Scope) error {
	out.Mode = KubeletAuthorizationMode(in.Mode)
	if err := Convert_kubeletconfig_KubeletWebhookAuthorization_To_v1beta1_KubeletWebhookAuthorization(&in.Webhook, &out.Webhook, s); err != nil {
		return err
	}
	return nil
}

// Convert_kubeletconfig_KubeletAuthorization_To_v1beta1_KubeletAuthorization is an autogenerated conversion function.
func Convert_kubeletconfig_KubeletAuthorization_To_v1beta1_KubeletAuthorization(in *kubeletconfig.KubeletAuthorization, out *KubeletAuthorization, s conversion.Scope) error {
	return autoConvert_kubeletconfig_KubeletAuthorization_To_v1beta1_KubeletAuthorization(in, out, s)
}

func autoConvert_v1beta1_KubeletConfiguration_To_kubeletconfig_KubeletConfiguration(in *KubeletConfiguration, out *kubeletconfig.KubeletConfiguration, s conversion.Scope) error {
	out.EnableServer = in.EnableServer
	out.StaticPodPath = in.StaticPodPath
	out.SyncFrequency = in.SyncFrequency
	out.FileCheckFrequency = in.FileCheckFrequency
	out.HTTPCheckFrequency = in.HTTPCheckFrequency
	out.StaticPodURL = in.StaticPodURL
	out.StaticPodURLHeader = *(*map[string][]string)(unsafe.Pointer(&in.StaticPodURLHeader))
	out.Address = in.Address
	out.Port = in.Port
	out.ReadOnlyPort = in.ReadOnlyPort
	out.TLSCertFile = in.TLSCertFile
	out.TLSPrivateKeyFile = in.TLSPrivateKeyFile
	out.TLSCipherSuites = *(*[]string)(unsafe.Pointer(&in.TLSCipherSuites))
	out.TLSMinVersion = in.TLSMinVersion
	out.RotateCertificates = in.RotateCertificates
	out.ServerTLSBootstrap = in.ServerTLSBootstrap
	if err := Convert_v1beta1_KubeletAuthentication_To_kubeletconfig_KubeletAuthentication(&in.Authentication, &out.Authentication, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_KubeletAuthorization_To_kubeletconfig_KubeletAuthorization(&in.Authorization, &out.Authorization, s); err != nil {
		return err
	}
	if err := v1.Convert_Pointer_int32_To_int32(&in.RegistryPullQPS, &out.RegistryPullQPS, s); err != nil {
		return err
	}
	out.RegistryBurst = in.RegistryBurst
	if err := v1.Convert_Pointer_int32_To_int32(&in.EventRecordQPS, &out.EventRecordQPS, s); err != nil {
		return err
	}
	out.EventBurst = in.EventBurst
	if err := v1.Convert_Pointer_bool_To_bool(&in.EnableDebuggingHandlers, &out.EnableDebuggingHandlers, s); err != nil {
		return err
	}
	out.EnableContentionProfiling = in.EnableContentionProfiling
	if err := v1.Convert_Pointer_int32_To_int32(&in.HealthzPort, &out.HealthzPort, s); err != nil {
		return err
	}
	out.HealthzBindAddress = in.HealthzBindAddress
	if err := v1.Convert_Pointer_int32_To_int32(&in.OOMScoreAdj, &out.OOMScoreAdj, s); err != nil {
		return err
	}
	out.ClusterDomain = in.ClusterDomain
	out.ClusterDNS = *(*[]string)(unsafe.Pointer(&in.ClusterDNS))
	out.StreamingConnectionIdleTimeout = in.StreamingConnectionIdleTimeout
	out.NodeStatusUpdateFrequency = in.NodeStatusUpdateFrequency
	out.ImageMinimumGCAge = in.ImageMinimumGCAge
	if err := v1.Convert_Pointer_int32_To_int32(&in.ImageGCHighThresholdPercent, &out.ImageGCHighThresholdPercent, s); err != nil {
		return err
	}
	if err := v1.Convert_Pointer_int32_To_int32(&in.ImageGCLowThresholdPercent, &out.ImageGCLowThresholdPercent, s); err != nil {
		return err
	}
	out.VolumeStatsAggPeriod = in.VolumeStatsAggPeriod
	out.KubeletCgroups = in.KubeletCgroups
	out.SystemCgroups = in.SystemCgroups
	out.CgroupRoot = in.CgroupRoot
	if err := v1.Convert_Pointer_bool_To_bool(&in.CgroupsPerQOS, &out.CgroupsPerQOS, s); err != nil {
		return err
	}
	out.CgroupDriver = in.CgroupDriver
	out.CPUManagerPolicy = in.CPUManagerPolicy
	out.CPUManagerReconcilePeriod = in.CPUManagerReconcilePeriod
	out.QOSReserved = *(*map[string]string)(unsafe.Pointer(&in.QOSReserved))
	out.RuntimeRequestTimeout = in.RuntimeRequestTimeout
	out.HairpinMode = in.HairpinMode
	out.MaxPods = in.MaxPods
	out.PodCIDR = in.PodCIDR
	if err := v1.Convert_Pointer_int64_To_int64(&in.PodPidsLimit, &out.PodPidsLimit, s); err != nil {
		return err
	}
	out.ResolverConfig = in.ResolverConfig
	if err := v1.Convert_Pointer_bool_To_bool(&in.CPUCFSQuota, &out.CPUCFSQuota, s); err != nil {
		return err
	}
	out.MaxOpenFiles = in.MaxOpenFiles
	out.ContentType = in.ContentType
	if err := v1.Convert_Pointer_int32_To_int32(&in.KubeAPIQPS, &out.KubeAPIQPS, s); err != nil {
		return err
	}
	out.KubeAPIBurst = in.KubeAPIBurst
	if err := v1.Convert_Pointer_bool_To_bool(&in.SerializeImagePulls, &out.SerializeImagePulls, s); err != nil {
		return err
	}
	out.EvictionHard = *(*map[string]string)(unsafe.Pointer(&in.EvictionHard))
	out.EvictionSoft = *(*map[string]string)(unsafe.Pointer(&in.EvictionSoft))
	out.EvictionSoftGracePeriod = *(*map[string]string)(unsafe.Pointer(&in.EvictionSoftGracePeriod))
	out.EvictionPressureTransitionPeriod = in.EvictionPressureTransitionPeriod
	out.EvictionMaxPodGracePeriod = in.EvictionMaxPodGracePeriod
	out.EvictionMinimumReclaim = *(*map[string]string)(unsafe.Pointer(&in.EvictionMinimumReclaim))
	out.PodsPerCore = in.PodsPerCore
	if err := v1.Convert_Pointer_bool_To_bool(&in.EnableControllerAttachDetach, &out.EnableControllerAttachDetach, s); err != nil {
		return err
	}
	out.ProtectKernelDefaults = in.ProtectKernelDefaults
	if err := v1.Convert_Pointer_bool_To_bool(&in.MakeIPTablesUtilChains, &out.MakeIPTablesUtilChains, s); err != nil {
		return err
	}
	if err := v1.Convert_Pointer_int32_To_int32(&in.IPTablesMasqueradeBit, &out.IPTablesMasqueradeBit, s); err != nil {
		return err
	}
	if err := v1.Convert_Pointer_int32_To_int32(&in.IPTablesDropBit, &out.IPTablesDropBit, s); err != nil {
		return err
	}
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	if err := v1.Convert_Pointer_bool_To_bool(&in.FailSwapOn, &out.FailSwapOn, s); err != nil {
		return err
	}
	out.ContainerLogMaxSize = in.ContainerLogMaxSize
	if err := v1.Convert_Pointer_int32_To_int32(&in.ContainerLogMaxFiles, &out.ContainerLogMaxFiles, s); err != nil {
		return err
	}
	out.ConfigMapAndSecretChangeDetectionStrategy = kubeletconfig.ResourceChangeDetectionStrategy(in.ConfigMapAndSecretChangeDetectionStrategy)
	out.SystemReserved = *(*map[string]string)(unsafe.Pointer(&in.SystemReserved))
	out.KubeReserved = *(*map[string]string)(unsafe.Pointer(&in.KubeReserved))
	out.SystemReservedCgroup = in.SystemReservedCgroup
	out.KubeReservedCgroup = in.KubeReservedCgroup
	out.EnforceNodeAllocatable = *(*[]string)(unsafe.Pointer(&in.EnforceNodeAllocatable))
	return nil
}

// Convert_v1beta1_KubeletConfiguration_To_kubeletconfig_KubeletConfiguration is an autogenerated conversion function.
func Convert_v1beta1_KubeletConfiguration_To_kubeletconfig_KubeletConfiguration(in *KubeletConfiguration, out *kubeletconfig.KubeletConfiguration, s conversion.Scope) error {
	return autoConvert_v1beta1_KubeletConfiguration_To_kubeletconfig_KubeletConfiguration(in, out, s)
}

func autoConvert_kubeletconfig_KubeletConfiguration_To_v1beta1_KubeletConfiguration(in *kubeletconfig.KubeletConfiguration, out *KubeletConfiguration, s conversion.Scope) error {
	out.EnableServer = in.EnableServer
	out.StaticPodPath = in.StaticPodPath
	out.SyncFrequency = in.SyncFrequency
	out.FileCheckFrequency = in.FileCheckFrequency
	out.HTTPCheckFrequency = in.HTTPCheckFrequency
	out.StaticPodURL = in.StaticPodURL
	out.StaticPodURLHeader = *(*map[string][]string)(unsafe.Pointer(&in.StaticPodURLHeader))
	out.Address = in.Address
	out.Port = in.Port
	out.ReadOnlyPort = in.ReadOnlyPort
	out.TLSCertFile = in.TLSCertFile
	out.TLSPrivateKeyFile = in.TLSPrivateKeyFile
	out.TLSCipherSuites = *(*[]string)(unsafe.Pointer(&in.TLSCipherSuites))
	out.TLSMinVersion = in.TLSMinVersion
	out.RotateCertificates = in.RotateCertificates
	out.ServerTLSBootstrap = in.ServerTLSBootstrap
	if err := Convert_kubeletconfig_KubeletAuthentication_To_v1beta1_KubeletAuthentication(&in.Authentication, &out.Authentication, s); err != nil {
		return err
	}
	if err := Convert_kubeletconfig_KubeletAuthorization_To_v1beta1_KubeletAuthorization(&in.Authorization, &out.Authorization, s); err != nil {
		return err
	}
	if err := v1.Convert_int32_To_Pointer_int32(&in.RegistryPullQPS, &out.RegistryPullQPS, s); err != nil {
		return err
	}
	out.RegistryBurst = in.RegistryBurst
	if err := v1.Convert_int32_To_Pointer_int32(&in.EventRecordQPS, &out.EventRecordQPS, s); err != nil {
		return err
	}
	out.EventBurst = in.EventBurst
	if err := v1.Convert_bool_To_Pointer_bool(&in.EnableDebuggingHandlers, &out.EnableDebuggingHandlers, s); err != nil {
		return err
	}
	out.EnableContentionProfiling = in.EnableContentionProfiling
	if err := v1.Convert_int32_To_Pointer_int32(&in.HealthzPort, &out.HealthzPort, s); err != nil {
		return err
	}
	out.HealthzBindAddress = in.HealthzBindAddress
	if err := v1.Convert_int32_To_Pointer_int32(&in.OOMScoreAdj, &out.OOMScoreAdj, s); err != nil {
		return err
	}
	out.ClusterDomain = in.ClusterDomain
	out.ClusterDNS = *(*[]string)(unsafe.Pointer(&in.ClusterDNS))
	out.StreamingConnectionIdleTimeout = in.StreamingConnectionIdleTimeout
	out.NodeStatusUpdateFrequency = in.NodeStatusUpdateFrequency
	out.ImageMinimumGCAge = in.ImageMinimumGCAge
	if err := v1.Convert_int32_To_Pointer_int32(&in.ImageGCHighThresholdPercent, &out.ImageGCHighThresholdPercent, s); err != nil {
		return err
	}
	if err := v1.Convert_int32_To_Pointer_int32(&in.ImageGCLowThresholdPercent, &out.ImageGCLowThresholdPercent, s); err != nil {
		return err
	}
	out.VolumeStatsAggPeriod = in.VolumeStatsAggPeriod
	out.KubeletCgroups = in.KubeletCgroups
	out.SystemCgroups = in.SystemCgroups
	out.CgroupRoot = in.CgroupRoot
	if err := v1.Convert_bool_To_Pointer_bool(&in.CgroupsPerQOS, &out.CgroupsPerQOS, s); err != nil {
		return err
	}
	out.CgroupDriver = in.CgroupDriver
	out.CPUManagerPolicy = in.CPUManagerPolicy
	out.CPUManagerReconcilePeriod = in.CPUManagerReconcilePeriod
	out.QOSReserved = *(*map[string]string)(unsafe.Pointer(&in.QOSReserved))
	out.RuntimeRequestTimeout = in.RuntimeRequestTimeout
	out.HairpinMode = in.HairpinMode
	out.MaxPods = in.MaxPods
	out.PodCIDR = in.PodCIDR
	if err := v1.Convert_int64_To_Pointer_int64(&in.PodPidsLimit, &out.PodPidsLimit, s); err != nil {
		return err
	}
	out.ResolverConfig = in.ResolverConfig
	if err := v1.Convert_bool_To_Pointer_bool(&in.CPUCFSQuota, &out.CPUCFSQuota, s); err != nil {
		return err
	}
	out.MaxOpenFiles = in.MaxOpenFiles
	out.ContentType = in.ContentType
	if err := v1.Convert_int32_To_Pointer_int32(&in.KubeAPIQPS, &out.KubeAPIQPS, s); err != nil {
		return err
	}
	out.KubeAPIBurst = in.KubeAPIBurst
	if err := v1.Convert_bool_To_Pointer_bool(&in.SerializeImagePulls, &out.SerializeImagePulls, s); err != nil {
		return err
	}
	out.EvictionHard = *(*map[string]string)(unsafe.Pointer(&in.EvictionHard))
	out.EvictionSoft = *(*map[string]string)(unsafe.Pointer(&in.EvictionSoft))
	out.EvictionSoftGracePeriod = *(*map[string]string)(unsafe.Pointer(&in.EvictionSoftGracePeriod))
	out.EvictionPressureTransitionPeriod = in.EvictionPressureTransitionPeriod
	out.EvictionMaxPodGracePeriod = in.EvictionMaxPodGracePeriod
	out.EvictionMinimumReclaim = *(*map[string]string)(unsafe.Pointer(&in.EvictionMinimumReclaim))
	out.PodsPerCore = in.PodsPerCore
	if err := v1.Convert_bool_To_Pointer_bool(&in.EnableControllerAttachDetach, &out.EnableControllerAttachDetach, s); err != nil {
		return err
	}
	out.ProtectKernelDefaults = in.ProtectKernelDefaults
	if err := v1.Convert_bool_To_Pointer_bool(&in.MakeIPTablesUtilChains, &out.MakeIPTablesUtilChains, s); err != nil {
		return err
	}
	if err := v1.Convert_int32_To_Pointer_int32(&in.IPTablesMasqueradeBit, &out.IPTablesMasqueradeBit, s); err != nil {
		return err
	}
	if err := v1.Convert_int32_To_Pointer_int32(&in.IPTablesDropBit, &out.IPTablesDropBit, s); err != nil {
		return err
	}
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	if err := v1.Convert_bool_To_Pointer_bool(&in.FailSwapOn, &out.FailSwapOn, s); err != nil {
		return err
	}
	out.ContainerLogMaxSize = in.ContainerLogMaxSize
	if err := v1.Convert_int32_To_Pointer_int32(&in.ContainerLogMaxFiles, &out.ContainerLogMaxFiles, s); err != nil {
		return err
	}
	out.ConfigMapAndSecretChangeDetectionStrategy = ResourceChangeDetectionStrategy(in.ConfigMapAndSecretChangeDetectionStrategy)
	out.SystemReserved = *(*map[string]string)(unsafe.Pointer(&in.SystemReserved))
	out.KubeReserved = *(*map[string]string)(unsafe.Pointer(&in.KubeReserved))
	out.SystemReservedCgroup = in.SystemReservedCgroup
	out.KubeReservedCgroup = in.KubeReservedCgroup
	out.EnforceNodeAllocatable = *(*[]string)(unsafe.Pointer(&in.EnforceNodeAllocatable))
	return nil
}

// Convert_kubeletconfig_KubeletConfiguration_To_v1beta1_KubeletConfiguration is an autogenerated conversion function.
func Convert_kubeletconfig_KubeletConfiguration_To_v1beta1_KubeletConfiguration(in *kubeletconfig.KubeletConfiguration, out *KubeletConfiguration, s conversion.Scope) error {
	return autoConvert_kubeletconfig_KubeletConfiguration_To_v1beta1_KubeletConfiguration(in, out, s)
}

func autoConvert_v1beta1_KubeletWebhookAuthentication_To_kubeletconfig_KubeletWebhookAuthentication(in *KubeletWebhookAuthentication, out *kubeletconfig.KubeletWebhookAuthentication, s conversion.Scope) error {
	if err := v1.Convert_Pointer_bool_To_bool(&in.Enabled, &out.Enabled, s); err != nil {
		return err
	}
	out.CacheTTL = in.CacheTTL
	return nil
}

// Convert_v1beta1_KubeletWebhookAuthentication_To_kubeletconfig_KubeletWebhookAuthentication is an autogenerated conversion function.
func Convert_v1beta1_KubeletWebhookAuthentication_To_kubeletconfig_KubeletWebhookAuthentication(in *KubeletWebhookAuthentication, out *kubeletconfig.KubeletWebhookAuthentication, s conversion.Scope) error {
	return autoConvert_v1beta1_KubeletWebhookAuthentication_To_kubeletconfig_KubeletWebhookAuthentication(in, out, s)
}

func autoConvert_kubeletconfig_KubeletWebhookAuthentication_To_v1beta1_KubeletWebhookAuthentication(in *kubeletconfig.KubeletWebhookAuthentication, out *KubeletWebhookAuthentication, s conversion.Scope) error {
	if err := v1.Convert_bool_To_Pointer_bool(&in.Enabled, &out.Enabled, s); err != nil {
		return err
	}
	out.CacheTTL = in.CacheTTL
	return nil
}

// Convert_kubeletconfig_KubeletWebhookAuthentication_To_v1beta1_KubeletWebhookAuthentication is an autogenerated conversion function.
func Convert_kubeletconfig_KubeletWebhookAuthentication_To_v1beta1_KubeletWebhookAuthentication(in *kubeletconfig.KubeletWebhookAuthentication, out *KubeletWebhookAuthentication, s conversion.Scope) error {
	return autoConvert_kubeletconfig_KubeletWebhookAuthentication_To_v1beta1_KubeletWebhookAuthentication(in, out, s)
}

func autoConvert_v1beta1_KubeletWebhookAuthorization_To_kubeletconfig_KubeletWebhookAuthorization(in *KubeletWebhookAuthorization, out *kubeletconfig.KubeletWebhookAuthorization, s conversion.Scope) error {
	out.CacheAuthorizedTTL = in.CacheAuthorizedTTL
	out.CacheUnauthorizedTTL = in.CacheUnauthorizedTTL
	return nil
}

// Convert_v1beta1_KubeletWebhookAuthorization_To_kubeletconfig_KubeletWebhookAuthorization is an autogenerated conversion function.
func Convert_v1beta1_KubeletWebhookAuthorization_To_kubeletconfig_KubeletWebhookAuthorization(in *KubeletWebhookAuthorization, out *kubeletconfig.KubeletWebhookAuthorization, s conversion.Scope) error {
	return autoConvert_v1beta1_KubeletWebhookAuthorization_To_kubeletconfig_KubeletWebhookAuthorization(in, out, s)
}

func autoConvert_kubeletconfig_KubeletWebhookAuthorization_To_v1beta1_KubeletWebhookAuthorization(in *kubeletconfig.KubeletWebhookAuthorization, out *KubeletWebhookAuthorization, s conversion.Scope) error {
	out.CacheAuthorizedTTL = in.CacheAuthorizedTTL
	out.CacheUnauthorizedTTL = in.CacheUnauthorizedTTL
	return nil
}

// Convert_kubeletconfig_KubeletWebhookAuthorization_To_v1beta1_KubeletWebhookAuthorization is an autogenerated conversion function.
func Convert_kubeletconfig_KubeletWebhookAuthorization_To_v1beta1_KubeletWebhookAuthorization(in *kubeletconfig.KubeletWebhookAuthorization, out *KubeletWebhookAuthorization, s conversion.Scope) error {
	return autoConvert_kubeletconfig_KubeletWebhookAuthorization_To_v1beta1_KubeletWebhookAuthorization(in, out, s)
}

func autoConvert_v1beta1_KubeletX509Authentication_To_kubeletconfig_KubeletX509Authentication(in *KubeletX509Authentication, out *kubeletconfig.KubeletX509Authentication, s conversion.Scope) error {
	out.ClientCAFile = in.ClientCAFile
	return nil
}

// Convert_v1beta1_KubeletX509Authentication_To_kubeletconfig_KubeletX509Authentication is an autogenerated conversion function.
func Convert_v1beta1_KubeletX509Authentication_To_kubeletconfig_KubeletX509Authentication(in *KubeletX509Authentication, out *kubeletconfig.KubeletX509Authentication, s conversion.Scope) error {
	return autoConvert_v1beta1_KubeletX509Authentication_To_kubeletconfig_KubeletX509Authentication(in, out, s)
}

func autoConvert_kubeletconfig_KubeletX509Authentication_To_v1beta1_KubeletX509Authentication(in *kubeletconfig.KubeletX509Authentication, out *KubeletX509Authentication, s conversion.Scope) error {
	out.ClientCAFile = in.ClientCAFile
	return nil
}

// Convert_kubeletconfig_KubeletX509Authentication_To_v1beta1_KubeletX509Authentication is an autogenerated conversion function.
func Convert_kubeletconfig_KubeletX509Authentication_To_v1beta1_KubeletX509Authentication(in *kubeletconfig.KubeletX509Authentication, out *KubeletX509Authentication, s conversion.Scope) error {
	return autoConvert_kubeletconfig_KubeletX509Authentication_To_v1beta1_KubeletX509Authentication(in, out, s)
}

func autoConvert_v1beta1_SerializedNodeConfigSource_To_kubeletconfig_SerializedNodeConfigSource(in *SerializedNodeConfigSource, out *kubeletconfig.SerializedNodeConfigSource, s conversion.Scope) error {
	out.Source = in.Source
	return nil
}

// Convert_v1beta1_SerializedNodeConfigSource_To_kubeletconfig_SerializedNodeConfigSource is an autogenerated conversion function.
func Convert_v1beta1_SerializedNodeConfigSource_To_kubeletconfig_SerializedNodeConfigSource(in *SerializedNodeConfigSource, out *kubeletconfig.SerializedNodeConfigSource, s conversion.Scope) error {
	return autoConvert_v1beta1_SerializedNodeConfigSource_To_kubeletconfig_SerializedNodeConfigSource(in, out, s)
}

func autoConvert_kubeletconfig_SerializedNodeConfigSource_To_v1beta1_SerializedNodeConfigSource(in *kubeletconfig.SerializedNodeConfigSource, out *SerializedNodeConfigSource, s conversion.Scope) error {
	out.Source = in.Source
	return nil
}

// Convert_kubeletconfig_SerializedNodeConfigSource_To_v1beta1_SerializedNodeConfigSource is an autogenerated conversion function.
func Convert_kubeletconfig_SerializedNodeConfigSource_To_v1beta1_SerializedNodeConfigSource(in *kubeletconfig.SerializedNodeConfigSource, out *SerializedNodeConfigSource, s conversion.Scope) error {
	return autoConvert_kubeletconfig_SerializedNodeConfigSource_To_v1beta1_SerializedNodeConfigSource(in, out, s)
}

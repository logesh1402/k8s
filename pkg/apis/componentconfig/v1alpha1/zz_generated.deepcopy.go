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
func (in *KubeProxyConfiguration) DeepCopyInto(out *KubeProxyConfiguration) {
	out.TypeMeta = in.TypeMeta
	out.BindAddress = in.BindAddress
	out.ClusterCIDR = in.ClusterCIDR
	out.HealthzBindAddress = in.HealthzBindAddress
	out.HealthzPort = in.HealthzPort
	out.HostnameOverride = in.HostnameOverride
	if in.IPTablesMasqueradeBit != nil {
		in, out := &in.IPTablesMasqueradeBit, &out.IPTablesMasqueradeBit
		*out = new(int32)
		**out = **in
	} else {
		out.IPTablesMasqueradeBit = nil
	}
	out.IPTablesSyncPeriod = in.IPTablesSyncPeriod
	out.KubeconfigPath = in.KubeconfigPath
	out.MasqueradeAll = in.MasqueradeAll
	out.Master = in.Master
	if in.OOMScoreAdj != nil {
		in, out := &in.OOMScoreAdj, &out.OOMScoreAdj
		*out = new(int32)
		**out = **in
	} else {
		out.OOMScoreAdj = nil
	}
	out.Mode = in.Mode
	out.PortRange = in.PortRange
	out.ResourceContainer = in.ResourceContainer
	out.UDPIdleTimeout = in.UDPIdleTimeout
	out.ConntrackMax = in.ConntrackMax
	out.ConntrackMaxPerCore = in.ConntrackMaxPerCore
	out.ConntrackMin = in.ConntrackMin
	out.ConntrackTCPEstablishedTimeout = in.ConntrackTCPEstablishedTimeout
	return
}

// DeepCopy will perform a deep copy of the receiver, creating a new KubeProxyConfiguration.
func (x *KubeProxyConfiguration) DeepCopy() *KubeProxyConfiguration {
	if x == nil {
		return nil
	}
	out := new(KubeProxyConfiguration)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject will perform a deep copy of the receiver, creating a new object.
func (x *KubeProxyConfiguration) DeepCopyObject() unversioned.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto will perform a deep copy of the receiver, writing to out. in must be non-nil.
func (in *KubeSchedulerConfiguration) DeepCopyInto(out *KubeSchedulerConfiguration) {
	out.TypeMeta = in.TypeMeta
	out.Port = in.Port
	out.Address = in.Address
	out.AlgorithmProvider = in.AlgorithmProvider
	out.PolicyConfigFile = in.PolicyConfigFile
	if in.EnableProfiling != nil {
		in, out := &in.EnableProfiling, &out.EnableProfiling
		*out = new(bool)
		**out = **in
	} else {
		out.EnableProfiling = nil
	}
	out.ContentType = in.ContentType
	out.KubeAPIQPS = in.KubeAPIQPS
	out.KubeAPIBurst = in.KubeAPIBurst
	out.SchedulerName = in.SchedulerName
	out.HardPodAffinitySymmetricWeight = in.HardPodAffinitySymmetricWeight
	out.FailureDomains = in.FailureDomains
	in.LeaderElection.DeepCopyInto(&out.LeaderElection)
	return
}

// DeepCopy will perform a deep copy of the receiver, creating a new KubeSchedulerConfiguration.
func (x *KubeSchedulerConfiguration) DeepCopy() *KubeSchedulerConfiguration {
	if x == nil {
		return nil
	}
	out := new(KubeSchedulerConfiguration)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject will perform a deep copy of the receiver, creating a new object.
func (x *KubeSchedulerConfiguration) DeepCopyObject() unversioned.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto will perform a deep copy of the receiver, writing to out. in must be non-nil.
func (in *KubeletAnonymousAuthentication) DeepCopyInto(out *KubeletAnonymousAuthentication) {
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	} else {
		out.Enabled = nil
	}
	return
}

// DeepCopy will perform a deep copy of the receiver, creating a new KubeletAnonymousAuthentication.
func (x *KubeletAnonymousAuthentication) DeepCopy() *KubeletAnonymousAuthentication {
	if x == nil {
		return nil
	}
	out := new(KubeletAnonymousAuthentication)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto will perform a deep copy of the receiver, writing to out. in must be non-nil.
func (in *KubeletAuthentication) DeepCopyInto(out *KubeletAuthentication) {
	out.X509 = in.X509
	in.Webhook.DeepCopyInto(&out.Webhook)
	in.Anonymous.DeepCopyInto(&out.Anonymous)
	return
}

// DeepCopy will perform a deep copy of the receiver, creating a new KubeletAuthentication.
func (x *KubeletAuthentication) DeepCopy() *KubeletAuthentication {
	if x == nil {
		return nil
	}
	out := new(KubeletAuthentication)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto will perform a deep copy of the receiver, writing to out. in must be non-nil.
func (in *KubeletAuthorization) DeepCopyInto(out *KubeletAuthorization) {
	out.Mode = in.Mode
	out.Webhook = in.Webhook
	return
}

// DeepCopy will perform a deep copy of the receiver, creating a new KubeletAuthorization.
func (x *KubeletAuthorization) DeepCopy() *KubeletAuthorization {
	if x == nil {
		return nil
	}
	out := new(KubeletAuthorization)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto will perform a deep copy of the receiver, writing to out. in must be non-nil.
func (in *KubeletConfiguration) DeepCopyInto(out *KubeletConfiguration) {
	out.TypeMeta = in.TypeMeta
	out.PodManifestPath = in.PodManifestPath
	out.SyncFrequency = in.SyncFrequency
	out.FileCheckFrequency = in.FileCheckFrequency
	out.HTTPCheckFrequency = in.HTTPCheckFrequency
	out.ManifestURL = in.ManifestURL
	out.ManifestURLHeader = in.ManifestURLHeader
	if in.EnableServer != nil {
		in, out := &in.EnableServer, &out.EnableServer
		*out = new(bool)
		**out = **in
	} else {
		out.EnableServer = nil
	}
	out.Address = in.Address
	out.Port = in.Port
	out.ReadOnlyPort = in.ReadOnlyPort
	out.TLSCertFile = in.TLSCertFile
	out.TLSPrivateKeyFile = in.TLSPrivateKeyFile
	out.CertDirectory = in.CertDirectory
	in.Authentication.DeepCopyInto(&out.Authentication)
	out.Authorization = in.Authorization
	out.HostnameOverride = in.HostnameOverride
	out.PodInfraContainerImage = in.PodInfraContainerImage
	out.DockerEndpoint = in.DockerEndpoint
	out.RootDirectory = in.RootDirectory
	out.SeccompProfileRoot = in.SeccompProfileRoot
	if in.AllowPrivileged != nil {
		in, out := &in.AllowPrivileged, &out.AllowPrivileged
		*out = new(bool)
		**out = **in
	} else {
		out.AllowPrivileged = nil
	}
	if in.HostNetworkSources != nil {
		in, out := &in.HostNetworkSources, &out.HostNetworkSources
		*out = make([]string, len(*in))
		copy(*out, *in)
	} else {
		out.HostNetworkSources = nil
	}
	if in.HostPIDSources != nil {
		in, out := &in.HostPIDSources, &out.HostPIDSources
		*out = make([]string, len(*in))
		copy(*out, *in)
	} else {
		out.HostPIDSources = nil
	}
	if in.HostIPCSources != nil {
		in, out := &in.HostIPCSources, &out.HostIPCSources
		*out = make([]string, len(*in))
		copy(*out, *in)
	} else {
		out.HostIPCSources = nil
	}
	if in.RegistryPullQPS != nil {
		in, out := &in.RegistryPullQPS, &out.RegistryPullQPS
		*out = new(int32)
		**out = **in
	} else {
		out.RegistryPullQPS = nil
	}
	out.RegistryBurst = in.RegistryBurst
	if in.EventRecordQPS != nil {
		in, out := &in.EventRecordQPS, &out.EventRecordQPS
		*out = new(int32)
		**out = **in
	} else {
		out.EventRecordQPS = nil
	}
	out.EventBurst = in.EventBurst
	if in.EnableDebuggingHandlers != nil {
		in, out := &in.EnableDebuggingHandlers, &out.EnableDebuggingHandlers
		*out = new(bool)
		**out = **in
	} else {
		out.EnableDebuggingHandlers = nil
	}
	out.MinimumGCAge = in.MinimumGCAge
	out.MaxPerPodContainerCount = in.MaxPerPodContainerCount
	if in.MaxContainerCount != nil {
		in, out := &in.MaxContainerCount, &out.MaxContainerCount
		*out = new(int32)
		**out = **in
	} else {
		out.MaxContainerCount = nil
	}
	out.CAdvisorPort = in.CAdvisorPort
	out.HealthzPort = in.HealthzPort
	out.HealthzBindAddress = in.HealthzBindAddress
	if in.OOMScoreAdj != nil {
		in, out := &in.OOMScoreAdj, &out.OOMScoreAdj
		*out = new(int32)
		**out = **in
	} else {
		out.OOMScoreAdj = nil
	}
	if in.RegisterNode != nil {
		in, out := &in.RegisterNode, &out.RegisterNode
		*out = new(bool)
		**out = **in
	} else {
		out.RegisterNode = nil
	}
	out.ClusterDomain = in.ClusterDomain
	out.MasterServiceNamespace = in.MasterServiceNamespace
	out.ClusterDNS = in.ClusterDNS
	out.StreamingConnectionIdleTimeout = in.StreamingConnectionIdleTimeout
	out.NodeStatusUpdateFrequency = in.NodeStatusUpdateFrequency
	out.ImageMinimumGCAge = in.ImageMinimumGCAge
	if in.ImageGCHighThresholdPercent != nil {
		in, out := &in.ImageGCHighThresholdPercent, &out.ImageGCHighThresholdPercent
		*out = new(int32)
		**out = **in
	} else {
		out.ImageGCHighThresholdPercent = nil
	}
	if in.ImageGCLowThresholdPercent != nil {
		in, out := &in.ImageGCLowThresholdPercent, &out.ImageGCLowThresholdPercent
		*out = new(int32)
		**out = **in
	} else {
		out.ImageGCLowThresholdPercent = nil
	}
	out.LowDiskSpaceThresholdMB = in.LowDiskSpaceThresholdMB
	out.VolumeStatsAggPeriod = in.VolumeStatsAggPeriod
	out.NetworkPluginName = in.NetworkPluginName
	out.NetworkPluginDir = in.NetworkPluginDir
	out.CNIConfDir = in.CNIConfDir
	out.CNIBinDir = in.CNIBinDir
	out.NetworkPluginMTU = in.NetworkPluginMTU
	out.VolumePluginDir = in.VolumePluginDir
	out.CloudProvider = in.CloudProvider
	out.CloudConfigFile = in.CloudConfigFile
	out.KubeletCgroups = in.KubeletCgroups
	out.RuntimeCgroups = in.RuntimeCgroups
	out.SystemCgroups = in.SystemCgroups
	out.CgroupRoot = in.CgroupRoot
	if in.CgroupsPerQOS != nil {
		in, out := &in.CgroupsPerQOS, &out.CgroupsPerQOS
		*out = new(bool)
		**out = **in
	} else {
		out.CgroupsPerQOS = nil
	}
	out.CgroupDriver = in.CgroupDriver
	out.ContainerRuntime = in.ContainerRuntime
	out.RemoteRuntimeEndpoint = in.RemoteRuntimeEndpoint
	out.RemoteImageEndpoint = in.RemoteImageEndpoint
	out.RuntimeRequestTimeout = in.RuntimeRequestTimeout
	out.RktPath = in.RktPath
	out.ExperimentalMounterPath = in.ExperimentalMounterPath
	out.ExperimentalMounterRootfsPath = in.ExperimentalMounterRootfsPath
	out.RktAPIEndpoint = in.RktAPIEndpoint
	out.RktStage1Image = in.RktStage1Image
	if in.LockFilePath != nil {
		in, out := &in.LockFilePath, &out.LockFilePath
		*out = new(string)
		**out = **in
	} else {
		out.LockFilePath = nil
	}
	out.ExitOnLockContention = in.ExitOnLockContention
	out.HairpinMode = in.HairpinMode
	out.BabysitDaemons = in.BabysitDaemons
	out.MaxPods = in.MaxPods
	out.NvidiaGPUs = in.NvidiaGPUs
	out.DockerExecHandlerName = in.DockerExecHandlerName
	out.PodCIDR = in.PodCIDR
	out.ResolverConfig = in.ResolverConfig
	if in.CPUCFSQuota != nil {
		in, out := &in.CPUCFSQuota, &out.CPUCFSQuota
		*out = new(bool)
		**out = **in
	} else {
		out.CPUCFSQuota = nil
	}
	if in.Containerized != nil {
		in, out := &in.Containerized, &out.Containerized
		*out = new(bool)
		**out = **in
	} else {
		out.Containerized = nil
	}
	out.MaxOpenFiles = in.MaxOpenFiles
	if in.ReconcileCIDR != nil {
		in, out := &in.ReconcileCIDR, &out.ReconcileCIDR
		*out = new(bool)
		**out = **in
	} else {
		out.ReconcileCIDR = nil
	}
	if in.RegisterSchedulable != nil {
		in, out := &in.RegisterSchedulable, &out.RegisterSchedulable
		*out = new(bool)
		**out = **in
	} else {
		out.RegisterSchedulable = nil
	}
	out.ContentType = in.ContentType
	if in.KubeAPIQPS != nil {
		in, out := &in.KubeAPIQPS, &out.KubeAPIQPS
		*out = new(int32)
		**out = **in
	} else {
		out.KubeAPIQPS = nil
	}
	out.KubeAPIBurst = in.KubeAPIBurst
	if in.SerializeImagePulls != nil {
		in, out := &in.SerializeImagePulls, &out.SerializeImagePulls
		*out = new(bool)
		**out = **in
	} else {
		out.SerializeImagePulls = nil
	}
	out.OutOfDiskTransitionFrequency = in.OutOfDiskTransitionFrequency
	out.NodeIP = in.NodeIP
	if in.NodeLabels != nil {
		in, out := &in.NodeLabels, &out.NodeLabels
		*out = make(map[string]string)
		for key, val := range *in {
			(*out)[key] = val
		}
	} else {
		out.NodeLabels = nil
	}
	out.NonMasqueradeCIDR = in.NonMasqueradeCIDR
	out.EnableCustomMetrics = in.EnableCustomMetrics
	if in.EvictionHard != nil {
		in, out := &in.EvictionHard, &out.EvictionHard
		*out = new(string)
		**out = **in
	} else {
		out.EvictionHard = nil
	}
	out.EvictionSoft = in.EvictionSoft
	out.EvictionSoftGracePeriod = in.EvictionSoftGracePeriod
	out.EvictionPressureTransitionPeriod = in.EvictionPressureTransitionPeriod
	out.EvictionMaxPodGracePeriod = in.EvictionMaxPodGracePeriod
	out.EvictionMinimumReclaim = in.EvictionMinimumReclaim
	out.PodsPerCore = in.PodsPerCore
	if in.EnableControllerAttachDetach != nil {
		in, out := &in.EnableControllerAttachDetach, &out.EnableControllerAttachDetach
		*out = new(bool)
		**out = **in
	} else {
		out.EnableControllerAttachDetach = nil
	}
	if in.SystemReserved != nil {
		in, out := &in.SystemReserved, &out.SystemReserved
		*out = make(map[string]string)
		for key, val := range *in {
			(*out)[key] = val
		}
	} else {
		out.SystemReserved = nil
	}
	if in.KubeReserved != nil {
		in, out := &in.KubeReserved, &out.KubeReserved
		*out = make(map[string]string)
		for key, val := range *in {
			(*out)[key] = val
		}
	} else {
		out.KubeReserved = nil
	}
	out.ProtectKernelDefaults = in.ProtectKernelDefaults
	if in.MakeIPTablesUtilChains != nil {
		in, out := &in.MakeIPTablesUtilChains, &out.MakeIPTablesUtilChains
		*out = new(bool)
		**out = **in
	} else {
		out.MakeIPTablesUtilChains = nil
	}
	if in.IPTablesMasqueradeBit != nil {
		in, out := &in.IPTablesMasqueradeBit, &out.IPTablesMasqueradeBit
		*out = new(int32)
		**out = **in
	} else {
		out.IPTablesMasqueradeBit = nil
	}
	if in.IPTablesDropBit != nil {
		in, out := &in.IPTablesDropBit, &out.IPTablesDropBit
		*out = new(int32)
		**out = **in
	} else {
		out.IPTablesDropBit = nil
	}
	if in.AllowedUnsafeSysctls != nil {
		in, out := &in.AllowedUnsafeSysctls, &out.AllowedUnsafeSysctls
		*out = make([]string, len(*in))
		copy(*out, *in)
	} else {
		out.AllowedUnsafeSysctls = nil
	}
	out.ExperimentalRuntimeIntegrationType = in.ExperimentalRuntimeIntegrationType
	return
}

// DeepCopy will perform a deep copy of the receiver, creating a new KubeletConfiguration.
func (x *KubeletConfiguration) DeepCopy() *KubeletConfiguration {
	if x == nil {
		return nil
	}
	out := new(KubeletConfiguration)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject will perform a deep copy of the receiver, creating a new object.
func (x *KubeletConfiguration) DeepCopyObject() unversioned.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto will perform a deep copy of the receiver, writing to out. in must be non-nil.
func (in *KubeletWebhookAuthentication) DeepCopyInto(out *KubeletWebhookAuthentication) {
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	} else {
		out.Enabled = nil
	}
	out.CacheTTL = in.CacheTTL
	return
}

// DeepCopy will perform a deep copy of the receiver, creating a new KubeletWebhookAuthentication.
func (x *KubeletWebhookAuthentication) DeepCopy() *KubeletWebhookAuthentication {
	if x == nil {
		return nil
	}
	out := new(KubeletWebhookAuthentication)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto will perform a deep copy of the receiver, writing to out. in must be non-nil.
func (in *KubeletWebhookAuthorization) DeepCopyInto(out *KubeletWebhookAuthorization) {
	out.CacheAuthorizedTTL = in.CacheAuthorizedTTL
	out.CacheUnauthorizedTTL = in.CacheUnauthorizedTTL
	return
}

// DeepCopy will perform a deep copy of the receiver, creating a new KubeletWebhookAuthorization.
func (x *KubeletWebhookAuthorization) DeepCopy() *KubeletWebhookAuthorization {
	if x == nil {
		return nil
	}
	out := new(KubeletWebhookAuthorization)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto will perform a deep copy of the receiver, writing to out. in must be non-nil.
func (in *KubeletX509Authentication) DeepCopyInto(out *KubeletX509Authentication) {
	out.ClientCAFile = in.ClientCAFile
	return
}

// DeepCopy will perform a deep copy of the receiver, creating a new KubeletX509Authentication.
func (x *KubeletX509Authentication) DeepCopy() *KubeletX509Authentication {
	if x == nil {
		return nil
	}
	out := new(KubeletX509Authentication)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto will perform a deep copy of the receiver, writing to out. in must be non-nil.
func (in *LeaderElectionConfiguration) DeepCopyInto(out *LeaderElectionConfiguration) {
	if in.LeaderElect != nil {
		in, out := &in.LeaderElect, &out.LeaderElect
		*out = new(bool)
		**out = **in
	} else {
		out.LeaderElect = nil
	}
	out.LeaseDuration = in.LeaseDuration
	out.RenewDeadline = in.RenewDeadline
	out.RetryPeriod = in.RetryPeriod
	return
}

// DeepCopy will perform a deep copy of the receiver, creating a new LeaderElectionConfiguration.
func (x *LeaderElectionConfiguration) DeepCopy() *LeaderElectionConfiguration {
	if x == nil {
		return nil
	}
	out := new(LeaderElectionConfiguration)
	x.DeepCopyInto(out)
	return out
}

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

package e2enode

import (
	"context"
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/util/sets"
	kubeletdevicepluginv1beta1 "k8s.io/kubelet/pkg/apis/deviceplugin/v1beta1"
	admissionapi "k8s.io/pod-security-admission/api"

	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/uuid"
	kubeletpodresourcesv1 "k8s.io/kubelet/pkg/apis/podresources/v1"
	kubeletpodresourcesv1alpha1 "k8s.io/kubelet/pkg/apis/podresources/v1alpha1"
	"k8s.io/kubernetes/test/e2e/framework"
	e2enode "k8s.io/kubernetes/test/e2e/framework/node"
	e2epod "k8s.io/kubernetes/test/e2e/framework/pod"
)

var (
	appsScheme = runtime.NewScheme()
	appsCodecs = serializer.NewCodecFactory(appsScheme)
)

// Serial because the test restarts Kubelet
var _ = SIGDescribe("Device Plugin [Feature:DevicePluginProbe][NodeFeature:DevicePluginProbe][Serial]", func() {
	f := framework.NewDefaultFramework("device-plugin-errors")
	f.NamespacePodSecurityEnforceLevel = admissionapi.LevelPrivileged
	testDevicePlugin(f, kubeletdevicepluginv1beta1.DevicePluginPath)
})

// readDaemonSetV1OrDie reads daemonset object from bytes. Panics on error.
func readDaemonSetV1OrDie(objBytes []byte) *appsv1.DaemonSet {
	appsv1.AddToScheme(appsScheme)
	requiredObj, err := runtime.Decode(appsCodecs.UniversalDecoder(appsv1.SchemeGroupVersion), objBytes)
	if err != nil {
		panic(err)
	}
	return requiredObj.(*appsv1.DaemonSet)
}

const (
	// TODO(vikasc): Instead of hard-coding number of devices, provide number of devices in the sample-device-plugin using configmap
	// and then use the same here
	expectedSampleDevsAmount int64 = 2
)

func testDevicePlugin(f *framework.Framework, pluginSockDir string) {
	pluginSockDir = filepath.Join(pluginSockDir) + "/"
	ginkgo.Context("DevicePlugin [Serial] [Disruptive]", func() {
		var devicePluginPod, dptemplate *v1.Pod
		var v1alphaPodResources *kubeletpodresourcesv1alpha1.ListPodResourcesResponse
		var v1PodResources *kubeletpodresourcesv1.ListPodResourcesResponse

		ginkgo.BeforeEach(func() {
			ginkgo.By("Wait for node to be ready")
			gomega.Eventually(func() bool {
				nodes, err := e2enode.TotalReady(f.ClientSet)
				framework.ExpectNoError(err)
				return nodes == 1
			}, time.Minute, time.Second).Should(gomega.BeTrue())

			ginkgo.By("Scheduling a sample device plugin pod")
			dp := getSampleDevicePluginPod(pluginSockDir)
			dptemplate = dp.DeepCopy()
			devicePluginPod = e2epod.NewPodClient(f).CreateSync(dp)

			ginkgo.By("Waiting for devices to become available on the local node")
			gomega.Eventually(func() bool {
				node, ready := getLocalTestNode(f)
				return ready && CountSampleDeviceCapacity(node) > 0
			}, 5*time.Minute, framework.Poll).Should(gomega.BeTrue())
			framework.Logf("Successfully created device plugin pod")

			ginkgo.By(fmt.Sprintf("Waiting for the resource exported by the sample device plugin to become available on the local node (instances: %d)", expectedSampleDevsAmount))
			gomega.Eventually(func() bool {
				node, ready := getLocalTestNode(f)
				return ready &&
					CountSampleDeviceCapacity(node) == expectedSampleDevsAmount &&
					CountSampleDeviceAllocatable(node) == expectedSampleDevsAmount
			}, 30*time.Second, framework.Poll).Should(gomega.BeTrue())
		})

		ginkgo.AfterEach(func() {
			ginkgo.By("Deleting the device plugin pod")
			e2epod.NewPodClient(f).DeleteSync(devicePluginPod.Name, metav1.DeleteOptions{}, time.Minute)

			ginkgo.By("Deleting any Pods created by the test")
			l, err := e2epod.NewPodClient(f).List(context.TODO(), metav1.ListOptions{})
			framework.ExpectNoError(err)
			for _, p := range l.Items {
				if p.Namespace != f.Namespace.Name {
					continue
				}

				framework.Logf("Deleting pod: %s", p.Name)
				e2epod.NewPodClient(f).DeleteSync(p.Name, metav1.DeleteOptions{}, 2*time.Minute)
			}

			restartKubelet(true)

			ginkgo.By("Waiting for devices to become unavailable on the local node")
			gomega.Eventually(func() bool {
				node, ready := getLocalTestNode(f)
				return ready && CountSampleDeviceCapacity(node) <= 0
			}, 5*time.Minute, framework.Poll).Should(gomega.BeTrue())

			ginkgo.By("devices now unavailable on the local node")
		})

		ginkgo.It("Can schedule a pod that requires a device", func() {
			podRECMD := "devs=$(ls /tmp/ | egrep '^Dev-[0-9]+$') && echo stub devices: $devs && sleep 60"
			pod1 := e2epod.NewPodClient(f).CreateSync(makeBusyboxPod(SampleDeviceResourceName, podRECMD))
			deviceIDRE := "stub devices: (Dev-[0-9]+)"
			devID1, err := parseLog(f, pod1.Name, pod1.Name, deviceIDRE)
			framework.ExpectNoError(err, "getting logs for pod %q", pod1.Name)
			gomega.Expect(devID1).To(gomega.Not(gomega.Equal("")), "pod1 requested a device but started successfully without")

			v1alphaPodResources, err = getV1alpha1NodeDevices()
			framework.ExpectNoError(err)

			v1PodResources, err := getV1NodeDevices()
			framework.ExpectNoError(err)

			framework.Logf("v1alphaPodResources.PodResources:%+v\n", v1alphaPodResources.PodResources)
			framework.Logf("v1PodResources.PodResources:%+v\n", v1PodResources.PodResources)
			framework.Logf("len(v1alphaPodResources.PodResources):%+v", len(v1alphaPodResources.PodResources))
			framework.Logf("len(v1PodResources.PodResources):%+v", len(v1PodResources.PodResources))

			framework.ExpectEqual(len(v1alphaPodResources.PodResources), 2)
			framework.ExpectEqual(len(v1PodResources.PodResources), 2)

			var v1alphaResourcesForOurPod *kubeletpodresourcesv1alpha1.PodResources
			for _, res := range v1alphaPodResources.GetPodResources() {
				if res.Name == pod1.Name {
					v1alphaResourcesForOurPod = res
				}
			}

			var v1ResourcesForOurPod *kubeletpodresourcesv1.PodResources
			for _, res := range v1PodResources.GetPodResources() {
				if res.Name == pod1.Name {
					v1ResourcesForOurPod = res
				}
			}

			gomega.Expect(v1alphaResourcesForOurPod).NotTo(gomega.BeNil())
			gomega.Expect(v1ResourcesForOurPod).NotTo(gomega.BeNil())

			framework.ExpectEqual(v1alphaResourcesForOurPod.Name, pod1.Name)
			framework.ExpectEqual(v1ResourcesForOurPod.Name, pod1.Name)

			framework.ExpectEqual(v1alphaResourcesForOurPod.Namespace, pod1.Namespace)
			framework.ExpectEqual(v1ResourcesForOurPod.Namespace, pod1.Namespace)

			framework.ExpectEqual(len(v1alphaResourcesForOurPod.Containers), 1)
			framework.ExpectEqual(len(v1ResourcesForOurPod.Containers), 1)

			framework.ExpectEqual(v1alphaResourcesForOurPod.Containers[0].Name, pod1.Spec.Containers[0].Name)
			framework.ExpectEqual(v1ResourcesForOurPod.Containers[0].Name, pod1.Spec.Containers[0].Name)

			framework.ExpectEqual(len(v1alphaResourcesForOurPod.Containers[0].Devices), 1)
			framework.ExpectEqual(len(v1ResourcesForOurPod.Containers[0].Devices), 1)

			framework.ExpectEqual(v1alphaResourcesForOurPod.Containers[0].Devices[0].ResourceName, SampleDeviceResourceName)
			framework.ExpectEqual(v1ResourcesForOurPod.Containers[0].Devices[0].ResourceName, SampleDeviceResourceName)

			framework.ExpectEqual(len(v1alphaResourcesForOurPod.Containers[0].Devices[0].DeviceIds), 1)
			framework.ExpectEqual(len(v1ResourcesForOurPod.Containers[0].Devices[0].DeviceIds), 1)
		})

		// simulate container restart, while all other involved components (kubelet, device plugin) stay stable. To do so, in the container
		// entry point we sleep for a limited and short period of time. The device assignment should be kept and be stable across the container
		// restarts. For the sake of brevity we however check just the fist restart.
		ginkgo.It("Keeps device plugin assignments across pod restarts (no kubelet restart, device plugin re-registration)", func(ctx context.Context) {
			podRECMD := "devs=$(ls /tmp/ | egrep '^Dev-[0-9]+$') && echo stub devices: $devs && sleep 60"
			pod1 := e2epod.NewPodClient(f).CreateSync(makeBusyboxPod(SampleDeviceResourceName, podRECMD))
			deviceIDRE := "stub devices: (Dev-[0-9]+)"
			devID1, err := parseLog(f, pod1.Name, pod1.Name, deviceIDRE)
			framework.ExpectNoError(err, "getting logs for pod %q", pod1.Name)
			gomega.Expect(devID1).To(gomega.Not(gomega.Equal("")), "pod1 requested a device but started successfully without")

			pod1, err = e2epod.NewPodClient(f).Get(ctx, pod1.Name, metav1.GetOptions{})
			framework.ExpectNoError(err)

			ginkgo.By("Waiting for container to restart")
			ensurePodContainerRestart(f, pod1.Name, pod1.Name)

			// check from the device assignment is preserved and stable from perspective of the container
			ginkgo.By("Confirming that after a container restart, fake-device assignment is kept")
			devIDRestart1, err := parseLog(f, pod1.Name, pod1.Name, deviceIDRE)
			framework.ExpectNoError(err, "getting logs for pod %q", pod1.Name)
			framework.ExpectEqual(devIDRestart1, devID1)

			// crosscheck from the device assignment is preserved and stable from perspective of the kubelet.
			// needs to match the container perspective.
			ginkgo.By("Verifying the device assignment after container restart using podresources API")
			v1PodResources, err = getV1NodeDevices()
			if err != nil {
				framework.ExpectNoError(err, "getting pod resources assignment after pod restart")
			}
			err = checkPodResourcesAssignment(v1PodResources, pod1.Namespace, pod1.Name, pod1.Spec.Containers[0].Name, SampleDeviceResourceName, []string{devID1})
			framework.ExpectNoError(err, "inconsistent device assignment after pod restart")

			ginkgo.By("Creating another pod")
			pod2 := e2epod.NewPodClient(f).CreateSync(makeBusyboxPod(SampleDeviceResourceName, podRECMD))
			err = e2epod.WaitTimeoutForPodRunningInNamespace(f.ClientSet, pod2.Name, f.Namespace.Name, 1*time.Minute)
			framework.ExpectNoError(err)

			ginkgo.By("Checking that pod got a fake device")
			devID2, err := parseLog(f, pod2.Name, pod2.Name, deviceIDRE)
			framework.ExpectNoError(err, "getting logs for pod %q", pod2.Name)

			gomega.Expect(devID2).To(gomega.Not(gomega.Equal("")), "pod2 requested a device but started successfully without")

			ginkgo.By("Verifying the device assignment after extra container start using podresources API")
			v1PodResources, err = getV1NodeDevices()
			if err != nil {
				framework.ExpectNoError(err, "getting pod resources assignment after pod restart")
			}
			err = checkPodResourcesAssignment(v1PodResources, pod1.Namespace, pod1.Name, pod1.Spec.Containers[0].Name, SampleDeviceResourceName, []string{devID1})
			framework.ExpectNoError(err, "inconsistent device assignment after extra container restart - pod1")
			err = checkPodResourcesAssignment(v1PodResources, pod2.Namespace, pod2.Name, pod2.Spec.Containers[0].Name, SampleDeviceResourceName, []string{devID2})
			framework.ExpectNoError(err, "inconsistent device assignment after extra container restart - pod2")
		})

		ginkgo.It("Keeps device plugin assignments after the device plugin has been re-registered", func(ctx context.Context) {
			podRECMD := "devs=$(ls /tmp/ | egrep '^Dev-[0-9]+$') && echo stub devices: $devs && sleep 60"
			pod1 := e2epod.NewPodClient(f).CreateSync(makeBusyboxPod(SampleDeviceResourceName, podRECMD))
			deviceIDRE := "stub devices: (Dev-[0-9]+)"
			devID1, err := parseLog(f, pod1.Name, pod1.Name, deviceIDRE)
			framework.ExpectNoError(err, "getting logs for pod %q", pod1.Name)

			gomega.Expect(devID1).To(gomega.Not(gomega.Equal("")), "pod1 requested a device but started successfully without")

			pod1, err = e2epod.NewPodClient(f).Get(ctx, pod1.Name, metav1.GetOptions{})
			framework.ExpectNoError(err)

			ginkgo.By("Restarting Kubelet")
			restartKubelet(true)

			ginkgo.By("Wait for node to be ready again")
			e2enode.WaitForAllNodesSchedulable(f.ClientSet, 5*time.Minute)

			ginkgo.By("Re-Register resources and delete the plugin pod")
			gp := int64(0)
			deleteOptions := metav1.DeleteOptions{
				GracePeriodSeconds: &gp,
			}
			e2epod.NewPodClient(f).DeleteSync(devicePluginPod.Name, deleteOptions, time.Minute)
			waitForContainerRemoval(devicePluginPod.Spec.Containers[0].Name, devicePluginPod.Name, devicePluginPod.Namespace)

			ginkgo.By("Recreating the plugin pod")
			devicePluginPod = e2epod.NewPodClient(f).CreateSync(dptemplate)

			ginkgo.By("Confirming that after a kubelet and pod restart, fake-device assignment is kept")
			ensurePodContainerRestart(f, pod1.Name, pod1.Name)
			devIDRestart1, err := parseLog(f, pod1.Name, pod1.Name, deviceIDRE)
			framework.ExpectNoError(err, "getting logs for pod %q", pod1.Name)
			framework.ExpectEqual(devIDRestart1, devID1)

			ginkgo.By("Waiting for resource to become available on the local node after re-registration")
			gomega.Eventually(func() bool {
				node, ready := getLocalTestNode(f)
				return ready &&
					CountSampleDeviceCapacity(node) == expectedSampleDevsAmount &&
					CountSampleDeviceAllocatable(node) == expectedSampleDevsAmount
			}, 30*time.Second, framework.Poll).Should(gomega.BeTrue())

			ginkgo.By("Creating another pod")
			pod2 := e2epod.NewPodClient(f).CreateSync(makeBusyboxPod(SampleDeviceResourceName, podRECMD))

			ginkgo.By("Checking that pod got a different fake device")
			devID2, err := parseLog(f, pod2.Name, pod2.Name, deviceIDRE)
			framework.ExpectNoError(err, "getting logs for pod %q", pod2.Name)

			gomega.Expect(devID1).To(gomega.Not(gomega.Equal(devID2)), "pod2 requested a device but started successfully without")
		})
	})
}

// makeBusyboxPod returns a simple Pod spec with a busybox container
// that requests SampleDeviceResourceName and runs the specified command.
func makeBusyboxPod(SampleDeviceResourceName, cmd string) *v1.Pod {
	podName := "device-plugin-test-" + string(uuid.NewUUID())
	rl := v1.ResourceList{v1.ResourceName(SampleDeviceResourceName): *resource.NewQuantity(1, resource.DecimalSI)}

	return &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{Name: podName},
		Spec: v1.PodSpec{
			RestartPolicy: v1.RestartPolicyAlways,
			Containers: []v1.Container{{
				Image: busyboxImage,
				Name:  podName,
				// Runs the specified command in the test pod.
				Command: []string{"sh", "-c", cmd},
				Resources: v1.ResourceRequirements{
					Limits:   rl,
					Requests: rl,
				},
			}},
		},
	}
}

// ensurePodContainerRestart confirms that pod container has restarted at least once
func ensurePodContainerRestart(f *framework.Framework, podName string, contName string) {
	var initialCount int32
	var currentCount int32
	p, err := e2epod.NewPodClient(f).Get(context.TODO(), podName, metav1.GetOptions{})
	if err != nil || len(p.Status.ContainerStatuses) < 1 {
		framework.Failf("ensurePodContainerRestart failed for pod %q: %v", podName, err)
	}
	initialCount = p.Status.ContainerStatuses[0].RestartCount
	gomega.Eventually(func() bool {
		p, err = e2epod.NewPodClient(f).Get(context.TODO(), podName, metav1.GetOptions{})
		if err != nil || len(p.Status.ContainerStatuses) < 1 {
			return false
		}
		currentCount = p.Status.ContainerStatuses[0].RestartCount
		framework.Logf("initial %v, current %v", initialCount, currentCount)
		return currentCount > initialCount
	}, 5*time.Minute, framework.Poll).Should(gomega.BeTrue())
}

// parseLog returns the matching string for the specified regular expression parsed from the container logs.
func parseLog(f *framework.Framework, podName string, contName string, re string) (string, error) {
	logs, err := e2epod.GetPodLogs(f.ClientSet, f.Namespace.Name, podName, contName)
	if err != nil {
		return "", err
	}

	framework.Logf("got pod logs: %v", logs)
	regex := regexp.MustCompile(re)
	matches := regex.FindStringSubmatch(logs)
	if len(matches) < 2 {
		return "", fmt.Errorf("unexpected match in logs: %q", logs)
	}

	return matches[1], nil
}

func checkPodResourcesAssignment(v1PodRes *kubeletpodresourcesv1.ListPodResourcesResponse, podNamespace, podName, containerName, resourceName string, devs []string) error {
	for _, podRes := range v1PodRes.PodResources {
		if podRes.Namespace != podNamespace || podRes.Name != podName {
			continue
		}
		for _, contRes := range podRes.Containers {
			if contRes.Name != containerName {
				continue
			}
			return matchContainerDevices(podNamespace+"/"+podName+"/"+containerName, contRes.Devices, resourceName, devs)
		}
	}
	return fmt.Errorf("no resources found for %s/%s/%s", podNamespace, podName, containerName)
}

func matchContainerDevices(ident string, contDevs []*kubeletpodresourcesv1.ContainerDevices, resourceName string, devs []string) error {
	expected := sets.NewString(devs...)
	assigned := sets.NewString()
	for _, contDev := range contDevs {
		if contDev.ResourceName != resourceName {
			continue
		}
		assigned = assigned.Insert(contDev.DeviceIds...)
	}
	expectedStr := strings.Join(expected.UnsortedList(), ",")
	assignedStr := strings.Join(assigned.UnsortedList(), ",")
	framework.Logf("%s: devices expected %q assigned %q", ident, expectedStr, assignedStr)
	if !assigned.Equal(expected) {
		return fmt.Errorf("device allocation mismatch for %s expected %s assigned %s", ident, expectedStr, assignedStr)
	}
	return nil
}

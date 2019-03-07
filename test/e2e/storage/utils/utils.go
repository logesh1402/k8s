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

package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/rand"
	"path/filepath"
	"strings"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/kubernetes/test/e2e/framework"
	imageutils "k8s.io/kubernetes/test/utils/image"
	uexec "k8s.io/utils/exec"
)

type KubeletOpt string

const (
	NodeStateTimeout            = 1 * time.Minute
	KStart           KubeletOpt = "start"
	KStop            KubeletOpt = "stop"
	KRestart         KubeletOpt = "restart"
)

const (
	// ClusterRole name for e2e test Priveledged Pod Security Policy User
	podSecurityPolicyPrivilegedClusterRoleName = "e2e-test-privileged-psp"
)

// PodExec wraps RunKubectl to execute a bash cmd in target pod
func PodExec(pod *v1.Pod, bashExec string) (string, error) {
	return framework.RunKubectl("exec", fmt.Sprintf("--namespace=%s", pod.Namespace), pod.Name, "--", "/bin/sh", "-c", bashExec)
}

// VerifyExecInPodSucceed verifies bash cmd in target pod succeed
func VerifyExecInPodSucceed(pod *v1.Pod, bashExec string) {
	_, err := PodExec(pod, bashExec)
	if err != nil {
		if err, ok := err.(uexec.CodeExitError); ok {
			exitCode := err.ExitStatus()
			framework.ExpectNoError(err,
				"%q should succeed, but failed with exit code %d and error message %q",
				bashExec, exitCode, err)
		} else {
			framework.ExpectNoError(err,
				"%q should succeed, but failed with error message %q",
				bashExec, err)
		}
	}
}

// VerifyExecInPodFail verifies bash cmd in target pod fail with certain exit code
func VerifyExecInPodFail(pod *v1.Pod, bashExec string, exitCode int) {
	_, err := PodExec(pod, bashExec)
	if err != nil {
		if err, ok := err.(uexec.CodeExitError); ok {
			actualExitCode := err.ExitStatus()
			Expect(actualExitCode).To(Equal(exitCode),
				"%q should fail with exit code %d, but failed with exit code %d and error message %q",
				bashExec, exitCode, actualExitCode, err)
		} else {
			framework.ExpectNoError(err,
				"%q should fail with exit code %d, but failed with error message %q",
				bashExec, exitCode, err)
		}
	}
	Expect(err).To(HaveOccurred(), "%q should fail with exit code %d, but exit without error", bashExec, exitCode)
}

// KubeletCommand performs `start`, `restart`, or `stop` on the kubelet running on the node of the target pod and waits
// for the desired statues..
// - First issues the command via `systemctl`
// - If `systemctl` returns stderr "command not found, issues the command via `service`
// - If `service` also returns stderr "command not found", the test is aborted.
// Allowed kubeletOps are `KStart`, `KStop`, and `KRestart`
func KubeletCommand(kOp KubeletOpt, c clientset.Interface, pod *v1.Pod) {
	command := ""
	sudoPresent := false
	systemctlPresent := false
	kubeletPid := ""

	nodeIP, err := framework.GetHostExternalAddress(c, pod)
	framework.ExpectNoError(err)
	nodeIP = nodeIP + ":22"

	framework.Logf("Checking if sudo command is present")
	sshResult, err := framework.SSH("sudo --version", nodeIP, framework.TestContext.Provider)
	framework.ExpectNoError(err, fmt.Sprintf("SSH to Node %q errored.", pod.Spec.NodeName))
	if !strings.Contains(sshResult.Stderr, "command not found") {
		sudoPresent = true
	}

	framework.Logf("Checking if systemctl command is present")
	sshResult, err = framework.SSH("systemctl --version", nodeIP, framework.TestContext.Provider)
	framework.ExpectNoError(err, fmt.Sprintf("SSH to Node %q errored.", pod.Spec.NodeName))
	if !strings.Contains(sshResult.Stderr, "command not found") {
		command = fmt.Sprintf("systemctl %s kubelet", string(kOp))
		systemctlPresent = true
	} else {
		command = fmt.Sprintf("service kubelet %s", string(kOp))
	}
	if sudoPresent {
		command = fmt.Sprintf("sudo %s", command)
	}

	if kOp == KRestart {
		kubeletPid = getKubeletMainPid(nodeIP, sudoPresent, systemctlPresent)
	}

	framework.Logf("Attempting `%s`", command)
	sshResult, err = framework.SSH(command, nodeIP, framework.TestContext.Provider)
	framework.ExpectNoError(err, fmt.Sprintf("SSH to Node %q errored.", pod.Spec.NodeName))
	framework.LogSSHResult(sshResult)
	Expect(sshResult.Code).To(BeZero(), "Failed to [%s] kubelet:\n%#v", string(kOp), sshResult)

	if kOp == KStop {
		if ok := framework.WaitForNodeToBeNotReady(c, pod.Spec.NodeName, NodeStateTimeout); !ok {
			framework.Failf("Node %s failed to enter NotReady state", pod.Spec.NodeName)
		}
	}
	if kOp == KRestart {
		// Wait for a minute to check if kubelet Pid is getting changed
		isPidChanged := false
		for start := time.Now(); time.Since(start) < 1*time.Minute; time.Sleep(2 * time.Second) {
			kubeletPidAfterRestart := getKubeletMainPid(nodeIP, sudoPresent, systemctlPresent)
			if kubeletPid != kubeletPidAfterRestart {
				isPidChanged = true
				break
			}
		}
		Expect(isPidChanged).To(BeTrue(), "Kubelet PID remained unchanged after restarting Kubelet")
		framework.Logf("Noticed that kubelet PID is changed. Waiting for 30 Seconds for Kubelet to come back")
		time.Sleep(30 * time.Second)
	}
	if kOp == KStart || kOp == KRestart {
		// For kubelet start and restart operations, Wait until Node becomes Ready
		if ok := framework.WaitForNodeToBeReady(c, pod.Spec.NodeName, NodeStateTimeout); !ok {
			framework.Failf("Node %s failed to enter Ready state", pod.Spec.NodeName)
		}
	}
}

// getKubeletMainPid return the Main PID of the Kubelet Process
func getKubeletMainPid(nodeIP string, sudoPresent bool, systemctlPresent bool) string {
	command := ""
	if systemctlPresent {
		command = "systemctl status kubelet | grep 'Main PID'"
	} else {
		command = "service kubelet status | grep 'Main PID'"
	}
	if sudoPresent {
		command = fmt.Sprintf("sudo %s", command)
	}
	framework.Logf("Attempting `%s`", command)
	sshResult, err := framework.SSH(command, nodeIP, framework.TestContext.Provider)
	framework.ExpectNoError(err, fmt.Sprintf("SSH to Node %q errored.", nodeIP))
	framework.LogSSHResult(sshResult)
	Expect(sshResult.Code).To(BeZero(), "Failed to get kubelet PID")
	Expect(sshResult.Stdout).NotTo(BeEmpty(), "Kubelet Main PID should not be Empty")
	return sshResult.Stdout
}

// TestKubeletRestartsAndRestoresMount tests that a volume mounted to a pod remains mounted after a kubelet restarts
func TestKubeletRestartsAndRestoresMount(c clientset.Interface, f *framework.Framework, clientPod *v1.Pod) {
	By("Writing to the volume.")
	file := "/mnt/_SUCCESS"
	out, err := PodExec(clientPod, fmt.Sprintf("touch %s", file))
	framework.Logf(out)
	framework.ExpectNoError(err)

	By("Restarting kubelet")
	KubeletCommand(KRestart, c, clientPod)

	By("Testing that written file is accessible.")
	out, err = PodExec(clientPod, fmt.Sprintf("cat %s", file))
	framework.Logf(out)
	framework.ExpectNoError(err)
	framework.Logf("Volume mount detected on pod %s and written file %s is readable post-restart.", clientPod.Name, file)
}

// TestVolumeUnmountsFromDeletedPod tests that a volume unmounts if the client pod was deleted while the kubelet was down.
// forceDelete is true indicating whether the pod is forcefully deleted.
func TestVolumeUnmountsFromDeletedPodWithForceOption(c clientset.Interface, f *framework.Framework, clientPod *v1.Pod, forceDelete bool, checkSubpath bool) {
	nodeIP, err := framework.GetHostExternalAddress(c, clientPod)
	framework.ExpectNoError(err)
	nodeIP = nodeIP + ":22"

	By("Expecting the volume mount to be found.")
	result, err := framework.SSH(fmt.Sprintf("mount | grep %s | grep -v volume-subpaths", clientPod.UID), nodeIP, framework.TestContext.Provider)
	framework.LogSSHResult(result)
	framework.ExpectNoError(err, "Encountered SSH error.")
	Expect(result.Code).To(BeZero(), fmt.Sprintf("Expected grep exit code of 0, got %d", result.Code))

	if checkSubpath {
		By("Expecting the volume subpath mount to be found.")
		result, err := framework.SSH(fmt.Sprintf("cat /proc/self/mountinfo | grep %s | grep volume-subpaths", clientPod.UID), nodeIP, framework.TestContext.Provider)
		framework.LogSSHResult(result)
		framework.ExpectNoError(err, "Encountered SSH error.")
		Expect(result.Code).To(BeZero(), fmt.Sprintf("Expected grep exit code of 0, got %d", result.Code))
	}

	// This command is to make sure kubelet is started after test finishes no matter it fails or not.
	defer func() {
		KubeletCommand(KStart, c, clientPod)
	}()
	By("Stopping the kubelet.")
	KubeletCommand(KStop, c, clientPod)

	By(fmt.Sprintf("Deleting Pod %q", clientPod.Name))
	if forceDelete {
		err = c.CoreV1().Pods(clientPod.Namespace).Delete(clientPod.Name, metav1.NewDeleteOptions(0))
	} else {
		err = c.CoreV1().Pods(clientPod.Namespace).Delete(clientPod.Name, &metav1.DeleteOptions{})
	}
	framework.ExpectNoError(err)

	By("Starting the kubelet and waiting for pod to delete.")
	KubeletCommand(KStart, c, clientPod)
	err = f.WaitForPodNotFound(clientPod.Name, framework.PodDeleteTimeout)
	if err != nil {
		framework.ExpectNoError(err, "Expected pod to be not found.")
	}

	if forceDelete {
		// With forceDelete, since pods are immediately deleted from API server, there is no way to be sure when volumes are torn down
		// so wait some time to finish
		time.Sleep(30 * time.Second)
	}

	By("Expecting the volume mount not to be found.")
	result, err = framework.SSH(fmt.Sprintf("mount | grep %s | grep -v volume-subpaths", clientPod.UID), nodeIP, framework.TestContext.Provider)
	framework.LogSSHResult(result)
	framework.ExpectNoError(err, "Encountered SSH error.")
	Expect(result.Stdout).To(BeEmpty(), "Expected grep stdout to be empty (i.e. no mount found).")
	framework.Logf("Volume unmounted on node %s", clientPod.Spec.NodeName)

	if checkSubpath {
		By("Expecting the volume subpath mount not to be found.")
		result, err = framework.SSH(fmt.Sprintf("cat /proc/self/mountinfo | grep %s | grep volume-subpaths", clientPod.UID), nodeIP, framework.TestContext.Provider)
		framework.LogSSHResult(result)
		framework.ExpectNoError(err, "Encountered SSH error.")
		Expect(result.Stdout).To(BeEmpty(), "Expected grep stdout to be empty (i.e. no subpath mount found).")
		framework.Logf("Subpath volume unmounted on node %s", clientPod.Spec.NodeName)
	}
}

// TestVolumeUnmountsFromDeletedPod tests that a volume unmounts if the client pod was deleted while the kubelet was down.
func TestVolumeUnmountsFromDeletedPod(c clientset.Interface, f *framework.Framework, clientPod *v1.Pod) {
	TestVolumeUnmountsFromDeletedPodWithForceOption(c, f, clientPod, false, false)
}

// TestVolumeUnmountsFromFoceDeletedPod tests that a volume unmounts if the client pod was forcefully deleted while the kubelet was down.
func TestVolumeUnmountsFromForceDeletedPod(c clientset.Interface, f *framework.Framework, clientPod *v1.Pod) {
	TestVolumeUnmountsFromDeletedPodWithForceOption(c, f, clientPod, true, false)
}

// RunInPodWithVolume runs a command in a pod with given claim mounted to /mnt directory.
func RunInPodWithVolume(c clientset.Interface, ns, claimName, command string) {
	pod := &v1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: "pvc-volume-tester-",
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:    "volume-tester",
					Image:   imageutils.GetE2EImage(imageutils.BusyBox),
					Command: []string{"/bin/sh"},
					Args:    []string{"-c", command},
					VolumeMounts: []v1.VolumeMount{
						{
							Name:      "my-volume",
							MountPath: "/mnt/test",
						},
					},
				},
			},
			RestartPolicy: v1.RestartPolicyNever,
			Volumes: []v1.Volume{
				{
					Name: "my-volume",
					VolumeSource: v1.VolumeSource{
						PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
							ClaimName: claimName,
							ReadOnly:  false,
						},
					},
				},
			},
		},
	}
	pod, err := c.CoreV1().Pods(ns).Create(pod)
	framework.ExpectNoError(err, "Failed to create pod: %v", err)
	defer func() {
		framework.DeletePodOrFail(c, ns, pod.Name)
	}()
	framework.ExpectNoError(framework.WaitForPodSuccessInNamespaceSlow(c, pod.Name, pod.Namespace))
}

func StartExternalProvisioner(c clientset.Interface, ns string, externalPluginName, serverIP string) *v1.Pod {
	podClient := c.CoreV1().Pods(ns)

	provisionerPod := &v1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: "external-provisioner-",
		},

		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:            "nfs-client-provisioner",
					Image:           "quay.io/external_storage/nfs-client-provisioner:v3.1.0-k8s1.11",
					ImagePullPolicy: v1.PullIfNotPresent,
					VolumeMounts: []v1.VolumeMount{
						{
							Name:      "nfs-client-root",
							MountPath: "/persistentvolumes",
						},
					},
					Env: []v1.EnvVar{
						{
							Name:  "PROVISIONER_NAME",
							Value: externalPluginName,
						},
						{
							Name:  "NFS_SERVER",
							Value: serverIP,
						},
						{
							Name:  "NFS_PATH",
							Value: "/",
						},
					},
				},
			},
			Volumes: []v1.Volume{
				{
					Name: "nfs-client-root",
					VolumeSource: v1.VolumeSource{
						NFS: &v1.NFSVolumeSource{
							Server: serverIP,
							Path:   "/",
						},
					},
				},
			},
		},
	}
	provisionerPod, err := podClient.Create(provisionerPod)
	framework.ExpectNoError(err, "Failed to create %s pod: %v", provisionerPod.Name, err)

	framework.ExpectNoError(framework.WaitForPodRunningInNamespace(c, provisionerPod))

	By("locating the provisioner pod")
	pod, err := podClient.Get(provisionerPod.Name, metav1.GetOptions{})
	framework.ExpectNoError(err, "Cannot locate the provisioner pod %v: %v", provisionerPod.Name, err)

	return pod
}

func PrivilegedTestPSPClusterRoleBinding(client clientset.Interface,
	namespace string,
	teardown bool,
	saNames []string) {
	bindingString := "Binding"
	if teardown {
		bindingString = "Unbinding"
	}
	roleBindingClient := client.RbacV1().RoleBindings(namespace)
	for _, saName := range saNames {
		By(fmt.Sprintf("%v priviledged Pod Security Policy to the service account %s", bindingString, saName))
		binding := &rbacv1.RoleBinding{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "psp-" + saName,
				Namespace: namespace,
			},
			Subjects: []rbacv1.Subject{
				{
					Kind:      rbacv1.ServiceAccountKind,
					Name:      saName,
					Namespace: namespace,
				},
			},
			RoleRef: rbacv1.RoleRef{
				Kind:     "ClusterRole",
				Name:     podSecurityPolicyPrivilegedClusterRoleName,
				APIGroup: "rbac.authorization.k8s.io",
			},
		}

		roleBindingClient.Delete(binding.GetName(), &metav1.DeleteOptions{})
		err := wait.Poll(2*time.Second, 2*time.Minute, func() (bool, error) {
			_, err := roleBindingClient.Get(binding.GetName(), metav1.GetOptions{})
			return apierrs.IsNotFound(err), nil
		})
		framework.ExpectNoError(err, "Timed out waiting for deletion: %v", err)

		if teardown {
			continue
		}

		_, err = roleBindingClient.Create(binding)
		framework.ExpectNoError(err, "Failed to create %s role binding: %v", binding.GetName(), err)

	}
}

func CheckVolumeModeOfPath(pod *v1.Pod, volMode v1.PersistentVolumeMode, path string) {
	if volMode == v1.PersistentVolumeBlock {
		// Check if block exists
		VerifyExecInPodSucceed(pod, fmt.Sprintf("test -b %s", path))

		// Double check that it's not directory
		VerifyExecInPodFail(pod, fmt.Sprintf("test -d %s", path), 1)
	} else {
		// Check if directory exists
		VerifyExecInPodSucceed(pod, fmt.Sprintf("test -d %s", path))

		// Double check that it's not block
		VerifyExecInPodFail(pod, fmt.Sprintf("test -b %s", path), 1)
	}
}

func CheckReadWriteToPath(pod *v1.Pod, volMode v1.PersistentVolumeMode, path string) {
	if volMode == v1.PersistentVolumeBlock {
		// random -> file1
		VerifyExecInPodSucceed(pod, "dd if=/dev/urandom of=/tmp/file1 bs=64 count=1")
		// file1 -> dev (write to dev)
		VerifyExecInPodSucceed(pod, fmt.Sprintf("dd if=/tmp/file1 of=%s bs=64 count=1", path))
		// dev -> file2 (read from dev)
		VerifyExecInPodSucceed(pod, fmt.Sprintf("dd if=%s of=/tmp/file2 bs=64 count=1", path))
		// file1 == file2 (check contents)
		VerifyExecInPodSucceed(pod, "diff /tmp/file1 /tmp/file2")
		// Clean up temp files
		VerifyExecInPodSucceed(pod, "rm -f /tmp/file1 /tmp/file2")

		// Check that writing file to block volume fails
		VerifyExecInPodFail(pod, fmt.Sprintf("echo 'Hello world.' > %s/file1.txt", path), 1)
	} else {
		// text -> file1 (write to file)
		VerifyExecInPodSucceed(pod, fmt.Sprintf("echo 'Hello world.' > %s/file1.txt", path))
		// grep file1 (read from file and check contents)
		VerifyExecInPodSucceed(pod, fmt.Sprintf("grep 'Hello world.' %s/file1.txt", path))

		// Check that writing to directory as block volume fails
		VerifyExecInPodFail(pod, fmt.Sprintf("dd if=/dev/urandom of=%s bs=64 count=1", path), 1)
	}
}

func genBinDataFromSeed(len int, seed int64) []byte {
	binData := make([]byte, len)
	rand.Seed(seed)

	len, err := rand.Read(binData)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	return binData
}

func CheckReadFromPath(pod *v1.Pod, volMode v1.PersistentVolumeMode, path string, len int, seed int64) {
	var pathForVolMode string
	if volMode == v1.PersistentVolumeBlock {
		pathForVolMode = path
	} else {
		pathForVolMode = filepath.Join(path, "file1.txt")
	}

	sum := sha256.Sum256(genBinDataFromSeed(len, seed))

	VerifyExecInPodSucceed(pod, fmt.Sprintf("dd if=%s bs=%d count=1 | sha256sum", pathForVolMode, len))
	VerifyExecInPodSucceed(pod, fmt.Sprintf("dd if=%s bs=%d count=1 | sha256sum | grep -Fq %x", pathForVolMode, len, sum))
}

func CheckWriteToPath(pod *v1.Pod, volMode v1.PersistentVolumeMode, path string, len int, seed int64) {
	var pathForVolMode string
	if volMode == v1.PersistentVolumeBlock {
		pathForVolMode = path
	} else {
		pathForVolMode = filepath.Join(path, "file1.txt")
	}

	encoded := base64.StdEncoding.EncodeToString(genBinDataFromSeed(len, seed))

	VerifyExecInPodSucceed(pod, fmt.Sprintf("echo %s | base64 -d | sha256sum", encoded))
	VerifyExecInPodSucceed(pod, fmt.Sprintf("echo %s | base64 -d | dd of=%s bs=%d count=1", encoded, pathForVolMode, len))
}

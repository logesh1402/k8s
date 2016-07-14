/*
Copyright 2014 The Kubernetes Authors All rights reserved.

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

package cmd

import (
	"fmt"
	"io"
	"net/url"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/client/restclient"
	client "k8s.io/kubernetes/pkg/client/unversioned"
	"k8s.io/kubernetes/pkg/client/unversioned/remotecommand"
	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	remotecommandserver "k8s.io/kubernetes/pkg/kubelet/server/remotecommand"
	"k8s.io/kubernetes/pkg/util/interrupt"
	"k8s.io/kubernetes/pkg/util/term"
)

const (
	exec_example = `# Get output from running 'date' from pod 123456-7890, using the first container by default
kubectl exec 123456-7890 date
	
# Get output from running 'date' in ruby-container from pod 123456-7890
kubectl exec 123456-7890 -c ruby-container date

# Switch to raw terminal mode, sends stdin to 'bash' in ruby-container from pod 123456-7890
# and sends stdout/stderr from 'bash' back to the client
kubectl exec 123456-7890 -c ruby-container -i -t -- bash -il`
)

func NewCmdExec(f *cmdutil.Factory, cmdIn io.Reader, cmdOut, cmdErr io.Writer) *cobra.Command {
	options := &ExecOptions{
		In:  cmdIn,
		Out: cmdOut,
		Err: cmdErr,

		Executor: &DefaultRemoteExecutor{},
	}
	cmd := &cobra.Command{
		Use:     "exec POD [-c CONTAINER] -- COMMAND [args...]",
		Short:   "Execute a command in a container.",
		Long:    "Execute a command in a container.",
		Example: exec_example,
		Run: func(cmd *cobra.Command, args []string) {
			argsLenAtDash := cmd.ArgsLenAtDash()
			cmdutil.CheckErr(options.Complete(f, cmd, args, argsLenAtDash))
			cmdutil.CheckErr(options.Validate())
			cmdutil.CheckErr(options.Run())
		},
	}
	cmd.Flags().StringVarP(&options.PodName, "pod", "p", "", "Pod name")
	// TODO support UID
	cmd.Flags().StringVarP(&options.ContainerName, "container", "c", "", "Container name. If omitted, the first container in the pod will be chosen")
	cmd.Flags().BoolVarP(&options.Stdin, "stdin", "i", false, "Pass stdin to the container")
	cmd.Flags().BoolVarP(&options.TTY, "tty", "t", false, "Stdin is a TTY")
	return cmd
}

// RemoteExecutor defines the interface accepted by the Exec command - provided for test stubbing
type RemoteExecutor interface {
	Execute(method string, url *url.URL, config *restclient.Config, stdin io.Reader, stdout, stderr io.Writer, tty bool, terminalSizeQueue term.TerminalSizeQueue) error
}

// DefaultRemoteExecutor is the standard implementation of remote command execution
type DefaultRemoteExecutor struct{}

func (*DefaultRemoteExecutor) Execute(method string, url *url.URL, config *restclient.Config, stdin io.Reader, stdout, stderr io.Writer, tty bool, terminalSizeQueue term.TerminalSizeQueue) error {
	exec, err := remotecommand.NewExecutor(config, method, url)
	if err != nil {
		return err
	}
	return exec.Stream(remotecommand.StreamOptions{
		SupportedProtocols: remotecommandserver.SupportedStreamingProtocols,
		Stdin:              stdin,
		Stdout:             stdout,
		Stderr:             stderr,
		Tty:                tty,
		TerminalSizeQueue:  terminalSizeQueue,
	})
}

// ExecOptions declare the arguments accepted by the Exec command
type ExecOptions struct {
	Namespace     string
	PodName       string
	ContainerName string
	Stdin         bool
	TTY           bool
	Command       []string

	// InterruptParent, if set, is used to handle interrupts while attached
	InterruptParent *interrupt.Handler

	In  io.Reader
	Out io.Writer
	Err io.Writer

	Executor RemoteExecutor
	Client   *client.Client
	Config   *restclient.Config
}

// Complete verifies command line arguments and loads data from the command environment
func (p *ExecOptions) Complete(f *cmdutil.Factory, cmd *cobra.Command, argsIn []string, argsLenAtDash int) error {
	// Let kubectl exec follow rules for `--`, see #13004 issue
	if len(p.PodName) == 0 && (len(argsIn) == 0 || argsLenAtDash == 0) {
		return cmdutil.UsageError(cmd, "POD is required for exec")
	}
	if len(p.PodName) != 0 {
		printDeprecationWarning("exec POD", "-p POD")
		if len(argsIn) < 1 {
			return cmdutil.UsageError(cmd, "COMMAND is required for exec")
		}
		p.Command = argsIn
	} else {
		p.PodName = argsIn[0]
		p.Command = argsIn[1:]
		if len(p.Command) < 1 {
			return cmdutil.UsageError(cmd, "COMMAND is required for exec")
		}
	}

	namespace, _, err := f.DefaultNamespace()
	if err != nil {
		return err
	}
	p.Namespace = namespace

	config, err := f.ClientConfig()
	if err != nil {
		return err
	}
	p.Config = config

	client, err := f.Client()
	if err != nil {
		return err
	}
	p.Client = client

	return nil
}

// Validate checks that the provided exec options are specified.
func (p *ExecOptions) Validate() error {
	if len(p.PodName) == 0 {
		return fmt.Errorf("pod name must be specified")
	}
	if len(p.Command) == 0 {
		return fmt.Errorf("you must specify at least one command for the container")
	}
	if p.Out == nil || p.Err == nil {
		return fmt.Errorf("both output and error output must be provided")
	}
	if p.Executor == nil || p.Client == nil || p.Config == nil {
		return fmt.Errorf("client, client config, and executor must be provided")
	}
	return nil
}

// Run executes a validated remote execution against a pod.
func (p *ExecOptions) Run() error {
	pod, err := p.Client.Pods(p.Namespace).Get(p.PodName)
	if err != nil {
		return err
	}

	containerName := p.ContainerName
	if len(containerName) == 0 {
		glog.V(4).Infof("defaulting container name to %s", pod.Spec.Containers[0].Name)
		containerName = pod.Spec.Containers[0].Name
	}

	// ensure we can recover the terminal while attached
	t := term.TTY{
		Parent: p.InterruptParent,
		Out:    p.Out,
	}

	// check for TTY
	tty := p.TTY
	if p.Stdin {
		t.In = p.In
		if tty && !t.IsTerminalIn() {
			tty = false
			fmt.Fprintln(p.Err, "Unable to use a TTY - input is not a terminal or the right kind of file")
		}
	} else {
		p.In = nil
	}
	t.Raw = tty

	var sizeQueue term.TerminalSizeQueue
	if tty {
		// this call spawns a goroutine to monitor/update the terminal size
		sizeQueue = t.MonitorSize(t.GetSize())

		// unset p.Err if it was previously set because both stdout and stderr go over p.Out when tty is
		// true
		p.Err = nil
	}

	fn := func() error {
		// TODO: consider abstracting into a client invocation or client helper
		req := p.Client.RESTClient.Post().
			Resource("pods").
			Name(pod.Name).
			Namespace(pod.Namespace).
			SubResource("exec").
			Param("container", containerName)
		req.VersionedParams(&api.PodExecOptions{
			Container: containerName,
			Command:   p.Command,
			Stdin:     p.Stdin,
			Stdout:    p.Out != nil,
			Stderr:    p.Err != nil,
			TTY:       tty,
		}, api.ParameterCodec)

		return p.Executor.Execute("POST", req.URL(), p.Config, p.In, p.Out, p.Err, tty, sizeQueue)
	}

	if err := t.Safe(fn); err != nil {
		return err
	}

	return nil
}

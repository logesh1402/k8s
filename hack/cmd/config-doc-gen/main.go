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

package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/golang/glog"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kubernetes/pkg/apis/componentconfig"
	"k8s.io/kubernetes/pkg/apis/componentconfig/v1alpha1"
	"k8s.io/kubernetes/pkg/util/yaml"
)

func main() {
	generateKubeProxy()
}

func generateKubeProxy() {
	scheme := runtime.NewScheme()

	if err := componentconfig.AddToScheme(scheme); err != nil {
		glog.Fatalf("error adding componentconfig internal verson to scheme: %v", err)
	}
	if err := v1alpha1.AddToScheme(scheme); err != nil {
		glog.Fatalf("error adding componentconfig v1alpha1 verson to scheme: %v", err)
	}

	config := &v1alpha1.KubeProxyConfiguration{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "componentconfig/v1alpha1",
			Kind:       "KubeProxyConfiguration",
		},
	}

	writeConfigYaml("cmd/kube-proxy/app", scheme, config)
}

func writeBoilerplate(f io.Writer) error {
	b, err := ioutil.ReadFile("hack/boilerplate/boilerplate.go.txt")
	if err != nil {
		return err
	}

	b = bytes.Replace(b, []byte("YEAR"), []byte(strconv.Itoa(time.Now().Year())), -1)
	if _, err := f.Write(b); err != nil {
		return err
	}

	return nil
}

func writeConfigYaml(destinationPackage string, scheme *runtime.Scheme, obj runtime.Object) error {
	scheme.Default(obj)

	file := filepath.Join(destinationPackage, "zz_generated.default_config_yaml.go")
	f, err := os.Create(file)
	if err != nil {
		return fmt.Errorf("error creating %q: %v", file, err)
	}
	defer f.Close()

	if err := writeBoilerplate(f); err != nil {
		return fmt.Errorf("error writing boilerplate: %v", err)
	}

	packageName := filepath.Base(destinationPackage)
	if _, err := fmt.Fprintf(f, `// THIS FILE IS AUTOGENERATED - DO NOT EDIT

package %s

const defaultConfigYaml = `+"`", packageName); err != nil {
		return fmt.Errorf("error writing to %q: %v", file, err)
	}

	if err := yaml.Marshal(obj, f); err != nil {
		return fmt.Errorf("error writing to %q: %v", file, err)
	}

	_, err = fmt.Fprintf(f, "`\n")
	if err != nil {
		return fmt.Errorf("error writing to %q: %v", file, err)
	}

	return nil
}

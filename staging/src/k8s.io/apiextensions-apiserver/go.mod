// This is a generated file. Do not edit directly.

module k8s.io/apiextensions-apiserver

go 1.13

require (
	github.com/emicklei/go-restful v2.9.5+incompatible
	github.com/go-openapi/errors v0.19.2
	github.com/go-openapi/spec v0.19.3
	github.com/go-openapi/strfmt v0.19.3
	github.com/go-openapi/validate v0.19.5
	github.com/gogo/protobuf v1.3.1
	github.com/google/go-cmp v0.4.0
	github.com/google/gofuzz v1.1.0
	github.com/google/uuid v1.1.1
	github.com/googleapis/gnostic v0.4.1
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.5.1
	go.etcd.io/etcd v0.5.0-alpha.5.0.20200401174654-e694b7bb0875
	google.golang.org/grpc v1.27.1
	gopkg.in/yaml.v2 v2.2.8
	k8s.io/api v0.0.0
	k8s.io/apimachinery v0.0.0
	k8s.io/apiserver v0.0.0
	k8s.io/client-go v0.0.0
	k8s.io/code-generator v0.0.0
	k8s.io/component-base v0.0.0
	k8s.io/klog v1.0.0
	k8s.io/kube-openapi v0.0.0-20200403204345-e1beb1bd0f35
	k8s.io/utils v0.0.0-20200324210504-a9aa75ae1b89
	sigs.k8s.io/yaml v1.2.0
)

replace (
	cloud.google.com/go => cloud.google.com/go v0.38.0
	github.com/NYTimes/gziphandler => github.com/NYTimes/gziphandler v0.0.0-20170623195520-56545f4a5d46
	github.com/kr/text => github.com/kr/text v0.1.0
	github.com/onsi/gomega => github.com/onsi/gomega v1.7.0
	github.com/stretchr/testify => github.com/stretchr/testify v1.4.0
	golang.org/x/oauth2 => golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	golang.org/x/sys => golang.org/x/sys v0.0.0-20190813064441-fde4db37ae7a // pinned to release-branch.go1.13
	golang.org/x/tools => golang.org/x/tools v0.0.0-20190821162956-65e3620a7ae7 // pinned to release-branch.go1.13
	google.golang.org/api => google.golang.org/api v0.15.1
	google.golang.org/appengine => google.golang.org/appengine v1.5.0
	google.golang.org/genproto => google.golang.org/genproto v0.0.0-20191230161307-f3c370f40bfb
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
	gopkg.in/check.v1 => gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15
	k8s.io/api => ../api
	k8s.io/apiextensions-apiserver => ../apiextensions-apiserver
	k8s.io/apimachinery => ../apimachinery
	k8s.io/apiserver => ../apiserver
	k8s.io/client-go => ../client-go
	k8s.io/code-generator => ../code-generator
	k8s.io/component-base => ../component-base
	sigs.k8s.io/structured-merge-diff/v3 => github.com/kwiesmueller/structured-merge-diff/v3 v3.0.0-20200325145223-0548f89a53da
)

module github.com/bhojpur/policy

go 1.16

require (
	github.com/Knetic/govaluate v3.0.0+incompatible
	github.com/bhojpur/dbm v0.0.3
	github.com/go-sql-driver/mysql v1.6.0
	github.com/golang/mock v1.6.0
	github.com/lib/pq v1.10.4
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.3.0
	google.golang.org/grpc v1.44.0
	google.golang.org/protobuf v1.27.1
	k8s.io/apimachinery v0.23.6
	k8s.io/client-go v1.5.2
)

require golang.org/x/tools v0.1.6-0.20210820212750-d4cc65f0b2ff // indirect

replace k8s.io/api => k8s.io/api v0.23.6

replace k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.23.6

replace k8s.io/apimachinery => k8s.io/apimachinery v0.23.7-rc.0

replace k8s.io/apiserver => k8s.io/apiserver v0.23.6

replace k8s.io/cli-runtime => k8s.io/cli-runtime v0.23.6

replace k8s.io/client-go => k8s.io/client-go v0.23.6

replace k8s.io/cloud-provider => k8s.io/cloud-provider v0.23.6

replace k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.23.6

replace k8s.io/code-generator => k8s.io/code-generator v0.23.7-rc.0

replace k8s.io/component-base => k8s.io/component-base v0.23.6

replace k8s.io/cri-api => k8s.io/cri-api v0.23.7-rc.0

replace k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.23.6

replace k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.23.6

replace k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.23.6

replace k8s.io/kube-proxy => k8s.io/kube-proxy v0.23.6

replace k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.23.6

replace k8s.io/kubelet => k8s.io/kubelet v0.23.6

replace k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.23.6

replace k8s.io/metrics => k8s.io/metrics v0.23.6

replace k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.23.6

replace k8s.io/component-helpers => k8s.io/component-helpers v0.23.6

replace k8s.io/controller-manager => k8s.io/controller-manager v0.23.6

replace k8s.io/kubectl => k8s.io/kubectl v0.23.6

replace k8s.io/mount-utils => k8s.io/mount-utils v0.23.7-rc.0

replace k8s.io/pod-security-admission => k8s.io/pod-security-admission v0.23.6

replace k8s.io/sample-cli-plugin => k8s.io/sample-cli-plugin v0.23.6

replace k8s.io/sample-controller => k8s.io/sample-controller v0.23.6

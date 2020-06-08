module github.com/kyma-project/kyma/components/application-broker

go 1.13

require (
	github.com/hashicorp/go-multierror v1.1.0 // indirect
	github.com/kubernetes-sigs/go-open-service-broker-client v0.0.0-20200527163240-4406bd2cb6b8 // indirect
	github.com/kubernetes-sigs/service-catalog v0.3.0 // indirect
	github.com/kyma-incubator/compass/components/connectivity-adapter v0.0.0-20200608084054-64f737ad7e1d // indirect
	github.com/kyma-incubator/compass/components/director v0.0.0-20200608084054-64f737ad7e1d // indirect
	github.com/kyma-project/kyma v0.5.1-0.20200608133559-bbd522ccf10b // indirect
	github.com/mcuadros/go-defaults v1.2.0 // indirect
	github.com/meatballhat/negroni-logrus v1.1.0 // indirect
	istio.io/client-go v0.0.0-20200606171330-0db30011a5d0 // indirect
	k8s.io/kubernetes v1.18.3 // indirect
	knative.dev/eventing v0.15.0 // indirect
)

replace (
	gopkg.in/yaml.v2 => gopkg.in/yaml.v2 v2.2.8

	// https://github.com/kubernetes/kubernetes/issues/79384
	k8s.io/api => k8s.io/api v0.18.3
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.18.3
	k8s.io/apimachinery => k8s.io/apimachinery v0.18.3
	k8s.io/apiserver => k8s.io/apiserver v0.18.3
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.18.3
	k8s.io/client-go => k8s.io/client-go v0.18.3
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.18.3
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.18.3
	k8s.io/code-generator => k8s.io/code-generator v0.18.3
	k8s.io/component-base => k8s.io/component-base v0.18.3
	k8s.io/cri-api => k8s.io/cri-api v0.18.3
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.18.3
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.18.3
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.18.3
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20200410145947-61e04a5be9a6
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.18.3
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.18.3
	k8s.io/kubectl => k8s.io/kubectl v0.18.3
	k8s.io/kubelet => k8s.io/kubelet v0.18.3
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.18.3
	k8s.io/metrics => k8s.io/metrics v0.18.3
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.18.3
	sigs.k8s.io/structured-merge-diff/v3 => sigs.k8s.io/structured-merge-diff/v3 v3.0.0
)

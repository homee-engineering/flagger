# Flagger end-to-end testing

The e2e testing infrastructure is powered by CircleCI and [Kubernetes Kind](https://github.com/kubernetes-sigs/kind).

CircleCI e2e workflow:

* install latest stable kubectl [e2e-kind.sh](e2e-kind.sh)
* build Kubernetes Kind from master [e2e-kind.sh](e2e-kind.sh)
* create local Kubernetes cluster with kind [e2e-kind.sh](e2e-kind.sh)
* install latest stable Helm CLI [e2e-istio.sh](e2e-istio.sh)
* deploy Tiller on the local cluster [e2e-istio.sh](e2e-istio.sh)
* install Istio CRDs with Helm [e2e-istio.sh](e2e-istio.sh)
* install Istio control plane and Prometheus with Helm [e2e-istio.sh](e2e-istio.sh)
* build Flagger container image [e2e-build.sh](e2e-build.sh)
* load Flagger image onto the local cluster [e2e-build.sh](e2e-build.sh)
* deploy Flagger in the istio-system namespace [e2e-build.sh](e2e-build.sh)
* create a test namespace with Istio injection enabled [e2e-tests.sh](e2e-tests.sh)
* deploy the load tester in the test namespace [e2e-tests.sh](e2e-tests.sh)
* deploy a demo workload (podinfo) in the test namespace [e2e-tests.sh](e2e-tests.sh)
* test the canary initialization [e2e-tests.sh](e2e-tests.sh)
* test the canary analysis and promotion [e2e-tests.sh](e2e-tests.sh)




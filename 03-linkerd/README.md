# LinkerD Service Mesh

**This installation is required only once per cluster.**

* [Download linkerd CLI](https://linkerd.io/2/getting-started/#step-1-install-the-cli)

Linux: `curl -sL run.linkerd.io/install | sh`
MacOS: `brew install linkerd`

* Check if everything is correct for LinkerD
* For Kubernetes >=1.17.0 the Clock Skew error can be ignored

```shell
linkerd check --pre
```

* Install Linkerd in Cluster and verify installation

```shell
linkerd install --crds | kubectl apply -f -
linkerd install | kubectl apply -f -
linkerd check
```

* Add Ingress for LinkerD dashboard or use `linkerd viz dashboard` without defining ingress and secret resource

* Add linkerd-viz extension

```shell
linkerd viz install | kubectl apply -f -
linkerd check
```

```shell
kubectl apply -f basic-auth.yaml
kubectl apply -f linkerd-ingress.yaml
```

* Inject LinkerD proxy into existing deployments

Add the "linkerd.io/inject: enabled" anntation to pods

```shell
kubectl edit deployment web-application

# in metadata.annotations, add a new line with "linkerd.io/inject: enabled"
# save and close your editor
```

* you can also inject linkerd to all deployments in a specific namespace

```shell
kubectl -n <NAMESPACE> get deployments -o yaml | linkerd inject - | kubectl apply -f -
```

* Verify if its pods are finally there:

`kubectl -n <NAMESPACE> get po -o jsonpath='{.items[0].spec.containers[*].name}'`

* Also mesh the ingress controller and check LinkerD dashboard

```shell
linkerd viz dashboard &
kubectl -n ingress-nginx get deployments -o yaml | linkerd inject - | kubectl apply -f -
```

* Verify mTLS traffic is working

`linkerd viz -n ${YOURNAME} edges pod`

* and in more detail

`linkerd viz -n simple-httpd tap pod`

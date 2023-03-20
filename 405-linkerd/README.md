# LinkerD Service Mesh

## Preparation

* [Download linkerd CLI](https://linkerd.io/2/getting-started/#step-1-install-the-cli)

Linux: `curl -sL run.linkerd.io/install | sh`

MacOS: `brew install linkerd`

### Verify client installation

Verify by this command: `linkerd version`

Expected output: `Client version: stable-2.12.4`

## Install LinkerD in the cluster

**This installation is required only once per cluster.**

* Check if everything is correct for LinkerD

```shell
linkerd check --pre
```

* Install Linkerd in Cluster and verify installation

```shell
linkerd install --crds | kubectl apply -f -
linkerd install | kubectl apply -f -
linkerd check
```

* Add linkerd-viz Dashboard extension

```shell
linkerd viz install | kubectl apply -f -
linkerd check
```

* Access `linkerd viz dashboard &` directly - this command automatically opens your desktop browser.

---

**Optional step**

* Add an Ingress for LinkerD dashboard

```shell
kubectl apply -f basic-auth.yaml
kubectl apply -f linkerd-ingress.yaml
```

---

Option 1:

* Inject LinkerD proxy into existing deployments

Add the "linkerd.io/inject: enabled" anntation to pods

  ```shell
    read -p "Please enter your name (without blanks e.g. johndoe): " YOURNAME
    export YOURNAME
  ```

  ```shell
  kubectl -n ${YOURNAME} edit deployment web-application
  
  # in metadata.annotations, add a new line with "linkerd.io/inject: enabled"
  # save and close your editor
  ```

---

Option 2:

* you can also inject linkerd to all deployments in a specific namespace

```shell
kubectl -n ${YOURNAME} get deployments -o yaml | linkerd inject - | kubectl apply -f -
```

* Result: your pods now consist of +1 additional container.

* Verify if its pods are finally there:

`kubectl -n ${YOURNAME} get po -o jsonpath='{.items[0].spec.containers[*].name}'`

Result:
* Observe the LinkerD dashboard and the connections to web-application

---

* Also mesh the ingress controller and check LinkerD dashboard

```shell
kubectl -n ingress-nginx get deployments -o yaml | linkerd inject - | kubectl apply -f -
```

* Enable "refresh" on web-application demo page

Result:

* Observe the LinkerD dashboard and the connections to web-application
* Now connections from the ingress-controller appear
* What are the connections from the "unmeshed node" source?
* Observe the web-application logs: `kubectl -n ${YOURNAME} logs -f deployments/web-application -c hello`
  * notice the probes
  * notice web traffic

---

* Verify mTLS traffic is working

`linkerd viz -n ${YOURNAME} edges pod`

* and in more detail

`linkerd viz -n ${YOURNAME} tap pod`

# kube-prometheus-stack

## Enable basic authentication

**This installation is required only once per cluster.**

* Create a secret which will be used as basic authentication credentials later

  ```shell
  kubectl -n monitoring apply -f basic-auth.yaml
  ```

* Take notice of the altered prom-values file where we enable prometheus ingress

* Update our prometheus deployment with the new settings:

  ```shell
  helm upgrade --install --namespace monitoring -f prom-values.yaml --version 45.6.0 prom prometheus-community/kube-prometheus-stack
  ```

* Verify basic authentication for prometheus web interface is enabled and enforced 

  ```shell
  kubectl -n monitoring port-forward svc/prom-kube-prometheus-stack-prometheus 9090:9090
  ```

**Test 1:**

* Visit http://localhost:9090/
* Result: Basic authentication is not enabled with direct access

**Test 2:**

* Now verify the ingress resource is working and prometheus web interface is now
also available over its public DNS name. Basic authentication

* Verify and obtain ingress address:

  `kubectl -n monitoring get ing`

* Visit https://prometheus.workshop.metakube.org

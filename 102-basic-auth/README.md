# kube-prometheus-stack

## Enable basic authentication

**This installation is required only once per cluster.**

* Create a secret which will be used as basic authentication credentials later

  ```shell
  kubectl -n monitoring apply -f basic-auth.yaml
  ```

* It contains these credentials:

  ```text
  User: foo
  Password: bar
  ```

* Take notice of the altered prom-values file where we enable prometheus ingress

* Update our prometheus deployment with the new settings:

  ```shell
  helm upgrade --install --namespace monitoring -f prom-values.yaml --version 45.6.0 prom prometheus-community/kube-prometheus-stack
  ```

* wait until the ingress resource is applied and shows a DNS name and public IP
  * `watch kubectl -n monitoring get ing`

* Verify basic authentication for prometheus web interface is enabled and enforced 

  ```shell
  kubectl -n monitoring port-forward svc/prom-kube-prometheus-stack-prometheus 9090:9090
  ```

**Test 1 - internal:**

* Visit http://localhost:9090/
  * Result: Basic authentication is enforced with direct access

**Test 2 - public:**

* Verify and obtain ingress address:

  `kubectl -n monitoring get ing`

* Visit https://prometheus.workshop.metakube.org
* Now verify the ingress resource is working:
  * Result: Prometheus web interface is now also available over its public DNS name.
  * Basic authentication is enforced, so enter the credentials from above.

### Conclusion

* basic authentication is enabled and prometheus web interface is publically available via ingress
* From now on you may access Prometheus web interface via its DNS name
* alternatively you may always use a port-forward with direct access

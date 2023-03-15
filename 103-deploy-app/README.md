# kube-prometheus-stack

### Deploy your own application

**This installation is required in every participants namespace.**

* Deploy the application

  ```shell
  read -p "Please enter your name (without blanks e.g. johndoe): " YOURNAME
  export YOURNAME
  helm -n ${YOURNAME} upgrade --install ${YOURNAME}-metrics-app metrics-app --set ingress.host=${YOURNAME}.workshop.metakube.org --set tls.cn=${YOURNAME}.workshop.metakube.org
  ```

* Verify it is installed properly and obtain your ingress endpoints

  ```shell
  helm -n ${YOURNAME} list
  ```

  ```shell
  kubectl -n ${YOURNAME} describe ingress metrics-app-ingress
  ```

  Result example:
  
  ```shell
  Rules:
    Host             Path  Backends
    ----             ----  --------
    johndoe.workshop.metakube.org  
                     /          metrics-app-service:8090 (172.25.1.106:8090)
                     /ping      metrics-app-service:8090 (172.25.1.106:8090)
                     /metrics   metrics-app-service:8090 (172.25.1.106:8090)
  ```

* Visit the particular metric endpoints of the application:
  * Visit https://YOURNAME.workshop.metakube.org/
    * Result: `hello world`
  * Visit https://YOURNAME.workshop.metakube.org/ping
    * Result: `pong`
  * Visit https://YOURNAME.workshop.metakube.org/metrics
    * Result: all metrics from go_lang prometheus client library
    * take notice of the metric: `helloworld_request_count`
    * take notice of the metric: `ping_request_count`

---

### Reconfigure Prometheus with a manual configuration

**This installation is required only once in the cluster.**

Now we need Prometheus to pick up metrics from our application

* Take notice of the altered `prom-values.yaml` file where we add `additionalScrapeConfigs`.

* Update our prometheus deployment with the new settings:

  ```shell
  helm upgrade --install --namespace monitoring -f prom-values.yaml --version 45.6.0 prom prometheus-community/kube-prometheus-stack
  ```

### View the results

* Visit the prometheus web interface https://prometheus.workshop.metakube.org

* optionally access it through a local port-forward on http://localhost:9090

  `kubectl -n monitoring port-forward svc/prom-kube-prometheus-stack-prometheus 9090:9090`

* under Status / Configuration view the configuration which contains the manual settings from helm values
  * look for `job_name: metrics-app`

* under Status / Targets view the targets and especially the new metrics-app target we configured by service discovery

* under Prometheus home page filter for these metrics:
  * `ping_request_count`
  * `helloworld_request_count`

---

### Clean up

Since we will learn a more straight-forward way how to pick up application metrics
we will now reset kube-prometheus-stack configuration.

**This installation is required only once in the cluster.**

* Revert our prometheus deployment with the old settings:

  ```shell
  helm upgrade --install --namespace monitoring -f prom-values-reset.yaml --version 45.6.0 prom prometheus-community/kube-prometheus-stack
  ```
**Result:** after a while the metrics-app should be gone from the targets on the prometheus web interface 
and from its configuration

---

## Conclusion

* We deployed our own application which provides custom metrics.

* Manual changes to the prometheus configuration are time-consuming, error-prone and
can only be rolled out by a prometheus admin person.

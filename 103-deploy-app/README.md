# kube-prometheus-stack

## Deploy your own application

**This installation is required in every participants namespace.**

* Deploy the application

  ```shell
  export YOURNAME=<YOURNAME>
  helm -n ${YOURNAME} upgrade --install ${YOURNAME}-metrics-app metrics-app --set ingress.host=${YOURNAME}.choerl.metakube.io #TODO
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
    metrics-app.YOURNAME.metakube.org  
                     /          metrics-app-service:8090 (172.25.1.106:8090)
                     /ping      metrics-app-service:8090 (172.25.1.106:8090)
                     /metrics   metrics-app-service:8090 (172.25.1.106:8090)
  ```

* Visit the particular metric endpoints of the application:
* Visit http://metrics-app.YOURNAME.metakube.org/
* Visit http://metrics-app.YOURNAME.metakube.org/ping
* Visit http://metrics-app.YOURNAME.metakube.org/metrics

## Reconfigure Prometheus with a manual configuration

**This installation is required only once in the cluster.**

* Now we need Prometheus to pick up metrics from our application

* Take notice of the altered prom-values.yaml file where we enable prometheus ingress

* Update our prometheus deployment with the new settings:

  ```shell
  helm upgrade --install --namespace monitoring -f prom-values.yaml --version 45.6.0 prom prometheus-community/kube-prometheus-stack
  ```

## View the results

* Visit the prometheus web interface http://prometheus.workshop.metakube.io

* View the configuration which contains the manual configuration from helm values

* View the targets and especially the new metrics-app target we configured by service discovery

## Conclusion

Manual changes to the prometheus configuration are time-consuming, error-prone and
can only be rolled out by a prometheus admin person.

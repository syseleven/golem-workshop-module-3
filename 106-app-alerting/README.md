# kube-prometheus-stack

## AlertManager Alerting Rules

### Scope

**This installation is required by every participant.**

Let's deploy an alerting rule for our application to instruct AlertManager
when to fire an alert event.

### Create an alerting rule

Alerting rules are defined as CRD resources of kind "PrometheusRule".

* First adjust your personal namespace in the rule manifest. Edit the file `rules/metrics-app-rule.yaml`

  ```yaml
  # metrics-app-rule.yaml
  [...]
    expr: sum(ping_request_count{namespace="<YOURNAME>"}) #<-- adjust your namespace
  [...] 
  ```

### Deploy alerting rule

* Deploy

  ```shell
  read -p "Please enter your name (without blanks e.g. johndoe): " YOURNAME
  export YOURNAME
  kubectl -n ${YOURNAME} apply -f rules/metrics-app-rule.yaml
  ``` 

* Verify

  ```shell
  kubectl -n ${YOURNAME} get prometheusrules.monitoring.coreos.com
  ```

### Trigger an alert

* to trigger AlertManager firing an alert call your application endpoint to exceed the threshold of the rule
  * Browse or curl the URL more than 5 times:
  * `curl https://${YOURNAME}.workshop.metakube.org/ping`

* View the result in the prometheus web interface
  * Status - Rules: new rule appears
  * Alerts: firing alerts

* Want to try again?
  * reset the metric counter to zero by this command:
  * `kubectl -n ${YOURNAME} rollout restart deployment metrics-app`

### Conclusion

* Metrics are reset to zero when the pod / deployment is deleted and re-installed.
* Our "static" alerting rule for Prometheus AlertManager is set up correctly.

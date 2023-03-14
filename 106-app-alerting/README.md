# kube-prometheus-stack

## Alerting Rules

### Scope

**This installation is required by every participant.**

Let's deploy an alerting rule for our application to instruct AlertManager
when to fire an alert event.

### Create an alerting rule

Alerting rules are defined as a CRD resource of kind PrometheusRule.

* Adjust the namespace in the rule manifest:
* edit rules/metrics-app-rule.yaml

  ```yaml
  # metrics-app-rule.yaml
  [...]
    expr: sum(ping_request_count{namespace="<YOURNAME>"}) #<-- adjust your namespace
  [...] 
  ```

### Deploy alerting rule

* deploy

  ```shell
  export YOURNAME=<YOURNAME>
  kubectl -n ${YOURNAME} apply -f rules/matrics-app-rule.yaml
  ``` 

### Trigger an alert

* to trigger AlertManager firing an alert call your application endpoint to exceed the threshold of the rule

  Browse or curl more than 5 times:
  
  `curl http://${YOURNAME}.workshop.metakube.org/ping`

* View the result in the prometheus web interface
  * Status - Rules: new rule appears
  * Alerts: firing alerts

### Conclusion

Metrics are reset to 0 when the pod / deployment is deleted and re-installed.

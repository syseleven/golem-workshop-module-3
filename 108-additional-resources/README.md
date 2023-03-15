# kube-prometheus-stack

## Additional resources

### Scope

**This installation is required only once in the cluster.**

We now deploy additional resources such as rules and serviceMonitors
for other components we would like to get metrics from.

Their metrics will make a lot more sense in Grafana later in this workshop.

### Inspect

Inspect the yaml files in the `rules` and `servicemonitors` folders.

### Deploy

* Deploy the resources

  ```shell
  kubectl -n monitoring apply -f rules/
  kubectl -n monitoring apply -f servicemonitors/
  ```

---

### Conclusion

We now have set up Prometheus completely with these components:
- password-protected web interface
- AlertManager

Further we deployed:
- our own application providing custom metrics
- Alerting rules
- serviceMonitors
- exporters (blackbox-exporter)
- we used the CRDs prometheus provides us with

Finally we are prepared to take the next step: Grafana

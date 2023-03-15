# kube-prometheus-stack

## Exporters - Blackbox-Exporter

### Scope

**This installation is required only once in the cluster.**

Blackbox-exporter is meant for calling generic targets over standardized methods such as
HTTP/S. For example you can use it to monitor a website and its certificates.

### Deploy the blackbox-exporter

  ```shell
  kubectl -n monitoring apply -f blackbox-exporter/
  ``` 
* View the result in the k8s cluster

  ```shell
  kubectl -n monitoring get po
  ```

* After a while view the result in prometheus web interface
  * a new target "blackbox-exporter" appears through its serviceMonitor

### Conclusion

In this exercise we deployed blackbox-exporter which already collects metrics
but will be of more interest when we view its dashboard in Grafana later in this workshop.

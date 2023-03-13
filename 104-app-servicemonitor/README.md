# kube-prometheus-stack

## Deploy a serviceMonitor for your own application

### Reset kube-prometheus-stack configuration

**This installation is required only once in the cluster.**

* Update our prometheus deployment with the old settings:

  ```shell
  helm upgrade --install --namespace monitoring -f prom-values.yaml --version 45.6.0 prom prometheus-community/kube-prometheus-stack
  ```

**Result:** the metrics-app should be gone from the targets on the prometheus web interface 
and from its configuration

### Prepare a serviceMonitor

**This installation is required in every participants namespace.**

* View the Custom Resource Definitions (CRDs) Prometheus installed

  ```shell
  kubectl get crd | grep monitoring
  ```

  Result:
  
  ```shell
  alertmanagerconfigs.monitoring.coreos.com             2023-01-19T15:26:12Z
  alertmanagers.monitoring.coreos.com                   2023-01-19T15:26:13Z
  podmonitors.monitoring.coreos.com                     2023-01-19T15:26:15Z
  probes.monitoring.coreos.com                          2023-01-19T15:26:16Z
  prometheuses.monitoring.coreos.com                    2023-01-19T15:26:18Z
  prometheusrules.monitoring.coreos.com                 2023-01-19T15:26:19Z
  servicemonitors.monitoring.coreos.com                 2023-01-19T15:26:21Z
  thanosrulers.monitoring.coreos.com                    2023-01-19T15:26:22Z
  ```

* We make use of the serviceMonitor CRD and deploy it for our application

* edit the file app-servicemonitor.yaml and adjust the name of your namespace

  ```yaml
    [...]
    namespaceSelector:
      matchNames:
        - <YOURNAME> #<-- please adjust your namespace here
    [...]
  ```

### Deploy the serviceMonitor

* then deploy the serviceMonitor to your namespace 

  ```shell
  export YOURNAME=<YOURNAME>
  kubectl -n ${YOURNAME} apply -f app-servicemonitor.yaml
  ```
* View the result in prometheus web interface
  * targets
  * metrics with detailled labels

* Recap how serviceMonitor connects to the service and how service connects to the pod

  ```shell
  kubectl -n <YOURNAME> get po,svc,servicemonitors.monitoring.coreos.com  --show-labels
  ```

### Conclusion

In this exercise we deployed a serviceMonitor for out application to leverage metrics discovery
by prometheus. As a result metrics contain even more pieces of information through labels.

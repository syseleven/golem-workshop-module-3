# Grafana

## Scope

**This installation is required only once per cluster.**

We already deployed Grafana before and we can already use it internally
but we would like to enable public (password-protected) access to it.

* This we achieve by enabling its ingress resource in the `prom-values.yaml`.

* First inspect the file

* Now update kube-prometheus-stack with these settings

  ```shell
  helm upgrade --install --namespace monitoring -f prom-values.yaml --version 45.6.0 prom prometheus-community/kube-prometheus-stack
  ```

* Verify it is installed

  ```shell
  helm -n monitoring list
  helm -n monitoring get values prom
  kubectl -n monitoring get ing
  ```

* get Grafana password and port-forward to UI

  ```shell
  kubectl -n monitoring get secret prom-grafana -o jsonpath="{.data.admin-password}" | base64 --decode
  kubectl -n monitoring port-forward service/prom-grafana 8080:80
  ```

Option 1:

* Visit http://localhost:8080/

Option 2:

* Visit public URL: https://grafana.workshop.metakube.org

* import [Node Exporter Dashboard](https://grafana.com/grafana/dashboards/13978?pg=dashboards&plcmt=featured-sub1)

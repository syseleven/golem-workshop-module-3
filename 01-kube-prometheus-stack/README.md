# Monitoring with prometheus

**This installation is required only once per cluster.**

### Prepare installation of kube-prometheus-stack

[ArtifactHub - kube-prometheus-stack](https://artifacthub.io/packages/helm/prometheus-community/kube-prometheus-stack)

```shell
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
```

```shell
kubectl create namespace monitoring
kubectl apply -f basic-auth.yaml
```

* Optional: Before installation of kube prometheus stack lets check what would be installed

  ```shell
  # This step is optional!
  
  # Install PlugIn for helm that gives diff info beforehand
  helm plugin install https://github.com/databus23/helm-diff
  
  # Run the diff before the real upgrade
  helm diff upgrade --disable-validation --allow-unreleased prom prometheus-community/kube-prometheus-stack --values values.yaml
  ```

---

### Install kube-prometheus-stack

* Install

  ```shell
  helm upgrade --install --namespace monitoring -f values.yaml --version 57.1.0 prom prometheus-community/kube-prometheus-stack
  
  kubectl apply -f dashboards/
  kubectl apply -f service-monitors/
  kubectl apply -f blackbox-exporter/
  kubectl apply -f rules/
  
  # Force reloading of grafana dashboards if needed
  kubectl -n monitoring rollout restart deployment prom-grafana
  ```

---

### Explore Promtheus Web UI

* port-forward to Prometheus UI

  ```shell
  kubectl -n monitoring port-forward service/prom-kube-prometheus-stack-prometheus 9090:9090
  ```

* Visit http://localhost:9090/
* Please notice: the Web UI is not protected by a password

---

### Explore Grafana

* get Grafana password and port-forward to UI

  ```shell
  kubectl -n monitoring get secret prom-grafana -o jsonpath="{.data.admin-password}" | base64 --decode
  kubectl -n monitoring port-forward service/prom-grafana 8080:80
  ```

---

* ONLY FOR JUMPHOST SETUP USERS

  ```shell
  ssh -L 8080:localhost:8080 <user>@<IP-of-Jumphost>
  ```

---

* Visit http://localhost:8080/

* import [Node Exporter Dashboard](https://grafana.com/grafana/dashboards/13978?pg=dashboards&plcmt=featured-sub1)

---

### Explore Alertmanager Web UI

* port-forward to Alertmanager UI

  ```shell
  kubectl -n monitoring port-forward svc/prom-kube-prometheus-stack-alertmanager 9093:9093
  ```

* Visit http://localhost:9093/
* Please notice: the Web UI is not protected by a password

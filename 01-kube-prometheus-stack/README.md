# Monitoring with prometheus

**This installation is required only once per cluster.**

[ArtifactHub - kube-prometheus-stack](https://artifacthub.io/packages/helm/prometheus-community/kube-prometheus-stack)

```shell
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
```

```shell
kubectl create namespace monitoring
kubectl label namespace monitoring golem-workshop=true
kubectl apply -f basic-auth.yaml
```

* Before installation of kube prometheus stack lets check what would be installed

```sh
helm diff upgrade --disable-validation --allow-unreleased prom prometheus-community/kube-prometheus-stack --values values.yaml
```

```shell
helm upgrade --install --namespace monitoring -f values.yaml --version 20.0.1 prom prometheus-community/kube-prometheus-stack

kubectl apply -f dashboards/
kubectl apply -f service-monitors/
kubectl apply -f blackbox-exporter/
kubectl apply -f rules/

# Force reloading of grafana dashboards if needed
kubectl rollout restart deployment prom-grafana
```

* get Grafana password and port-forward to UI

```shell
kubectl -n monitoring get secret prom-grafana -o jsonpath="{.data.admin-password}" | base64 --decode
kubectl -n monitoring port-forward service/prom-grafana 8080:80
```

* import [Node Exporter Dashboard](https://grafana.com/grafana/dashboards/13978?pg=dashboards&plcmt=featured-sub1)

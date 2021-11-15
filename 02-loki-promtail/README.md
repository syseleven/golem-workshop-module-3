# Logging with Loki

**This installation is required only once per cluster.**

## Install Loki and configure Grafana

[loki](https://artifacthub.io/packages/helm/grafana/loki)
[promtail](https://artifacthub.io/packages/helm/grafana/promtail)

```shell
helm repo add grafana https://grafana.github.io/helm-charts
helm repo update
helm upgrade --install loki grafana/loki --namespace monitoring -f values-loki.yaml --version 2.8.1
helm upgrade --install promtail grafana/promtail --namespace monitoring -f values-promtail.yaml --version 3.9.1
kubectl apply -f datasource.yaml
```

## Restart Grafana

```shell
kubectl rollout restart deployment -n monitoring prom-grafana
```

# kube-prometheus-stack

## Preparation

* Before you begin with the actual exercise please make sure to follow these steps to work in your own environment:

  ```shell
  kubectl create ns <YOURNAME>
  kubectl label namespace <YOURNAME> deepdive-observability=true
  kubectl config set-context --current --namespace=<YOURNAME>
  ```

* Clone this repository to your working station and change into the directory for the following exercises

## Deployment

**This installation is required only once per cluster.**

[ArtifactHub - kube-prometheus-stack](https://artifacthub.io/packages/helm/prometheus-community/kube-prometheus-stack)

* Add the Helm repository to you local helm client and update its inventory

```shell
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
```

```shell
kubectl create namespace monitoring
kubectl label namespace monitoring deepdive-observability=true
```

* Before installation of kube-prometheus-stack let's check what would be installed

```sh
# This step is optional!

# Install PlugIn for helm that gives diff info beforehand
helm plugin install https://github.com/databus23/helm-diff

# Run the diff before the real upgrade
helm diff upgrade --disable-validation --allow-unreleased prom prometheus-community/kube-prometheus-stack --values values.yaml
```

* Now deploy kube-prometheus-stack

```shell
helm upgrade --install --namespace monitoring -f values.yaml --version 45.6.0 prom prometheus-community/kube-prometheus-stack
```

* Verify it is installed

```shell
helm -n monitoring list
helm -n monitoring get values prom
```

* get Grafana password and port-forward to UI

```shell
kubectl -n monitoring get secret prom-grafana -o jsonpath="{.data.admin-password}" | base64 --decode
kubectl -n monitoring port-forward service/prom-grafana 8080:80
```
* ONLY FOR JUMPHOST SETUP USERS

```shell
ssh -L 8080:localhost:8080 <user>@<IP-of-Jumphost>
```

* Visit http://localhost:8080/

* import [Node Exporter Dashboard](https://grafana.com/grafana/dashboards/13978?pg=dashboards&plcmt=featured-sub1)

# kube-prometheus-stack

## Preparation

* Before you begin with the actual exercise please make sure to follow these steps to work in your own environment:

  ```shell
  read -p "Please enter your name (without blanks e.g. johndoe): " YOURNAME
  export YOURNAME
  kubectl create ns ${YOURNAME}
  kubectl label namespace ${YOURNAME} deepdive-observability=true
  kubectl config set-context --current --namespace=${YOURNAME}
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
  helm diff upgrade --namespace monitoring --disable-validation --allow-unreleased prom prometheus-community/kube-prometheus-stack --version 45.6.0 --values prom-values.yaml
  ```

* Inspect the values it will deployed with: `prom-values.yaml`
  * take notice of the `retention` information
  * take notice of the `persistence` settings

* Inspect the persistent volumes and volume claims in k8s
  * `kubectl -n monitoring get pvc,pv`

* Now deploy kube-prometheus-stack

  ```shell
  helm upgrade --install --namespace monitoring -f prom-values.yaml --version 45.6.0 prom prometheus-community/kube-prometheus-stack
  ```

* Verify it is installed

  ```shell
  helm -n monitoring list
  helm -n monitoring get values prom
  ```

* Verify prometheus service and create a port-forward to its web interface

  ```shell
  kubectl -n monitoring get svc
  kubectl -n monitoring port-forward svc/prom-kube-prometheus-stack-prometheus 9090:9090
  ```

* For the prometheus web interface visit http://localhost:9090/

## Conclusion

In this exercise you installed kube-prometheus-stack helm chart and got access to its web interface by a port-forward.

**Please note:** The web interface is available without any authorization by default, 
you might want to protect it by a password at least.

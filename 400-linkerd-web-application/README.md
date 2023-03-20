# Deployment with Ingress and TLS

**This installation is required only once per cluster.**

## Task

Let's deploy a demo web application to generate some traffic.

* Edit hostnames in `web-application/deployment/ingress.yaml`

* Deploy application

  ```shell
    read -p "Please enter your name (without blanks e.g. johndoe): " YOURNAME
    export YOURNAME
  ```

  ```shell
  kubectl -n ${YOURNAME} apply -f web-application/deployment/web-application.yaml
  ```

* Deploy Ingress

  ```shell
  kubectl -n ${YOURNAME} apply -f web-application/deployment/ingress.yaml
  ```

* Access your web application

  `kubectl -n ${YOURNAME} get ing`

  Visit: https://web-application-YOURNAME.workshop.metakube.org

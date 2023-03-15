# Introduction

## Prerequisites

You need to have these tools installed on your local machine:

- kubectl
- helm
- linkerd

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

## Landscape

You will share a single Kubernetes cluster with other workshop participants.
Your own examples can be executed and tested in you personal namespace.

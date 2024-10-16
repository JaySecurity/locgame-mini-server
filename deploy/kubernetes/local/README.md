# Introduction

These are instructions to set up a minikube cluster to run locg-service locally.
Everything is assumed to be run from locg-service root.

#Automatic compatible with windows and WSL
###Do not forget to place your secret files on Secrets' folder next to ../local directory. Although, the database is created and hardcoded locally
```bash
./deploy/kubernetes/local/start.sh
```
### There are 4 steps that you can run each depending on your need
```bash
./deploy/kubernetes/local/init_minikube.sh

./deploy/kubernetes/local/import_yaml_files.sh

./deploy/kubernetes/local/build_and_deploy_locg_service.sh

./deploy/kubernetes/local/expose_ports.sh

```
####You may need to change docker file builder images both on main directory and ./cmd/locgame-jobs-service/build/Dockerfile for better speed time
```dockerfile
FROM docker.io/library/locg-jobs-service:builder as builder
```
# Manual
# Install dependencies

```bash
brew install minikube
brew install docker
brew install helm
```

# Start cluster:

```bash
minikube start --driver=docker
```

# Set up minikube container registry

```bash
minikube addons enable registry
```

# Set up cluster

## Create namespace

Since we aim to be fully isolated, we'll create a new `namespace` and deploy into it

```bash
kubectl create ns locg-service
kubectl config set-context --current --namespace=locg-service
```

## Add NATS

```bash
helm repo add nats https://nats-io.github.io/k8s/helm/charts/
helm install locg-nats nats/nats --namespace locg-service

```

<!-- ## Add Mongo
```bash
helm repo add groundhog2k https://groundhog2k.github.io/helm-charts/
helm install locg-mongodb groundhog2k/mongodb --namespace locg-service
``` -->

## Add Redis

```bash
helm repo add bitnami https://charts.bitnami.com/bitnami
helm install locg-redis bitnami/redis --namespace locg-service
```

## Add locg-service and jobs-service

```bash
kubectl apply -f ./deploy/kubernetes/local/service.yaml --namespace locg-service
kubectl apply -f ./deploy/kubernetes/local/locg-service-deployment.yaml --namespace locg-service
kubectl apply -f ./deploy/kubernetes/local/jobs-service-deployment.yaml --namespace locg-service
kubectl apply -f ./deploy/kubernetes/secrets/secret-local.yaml --namespace locg-service
```

This will result in the service and deployments being created, but we still need to build and push the images

# Build LOCG Service

```bash
./build_and_deploy_locg_service.sh
```

# Expose them outside the cluster

```bash
kubectl port-forward --address 0.0.0.0 service/locg-service 8080:8080
```

# Clean Up

When you are done, just remove the namespace and everything should be deleted. This can be done by typing

```bash
kubectl delete ns locg-service
```

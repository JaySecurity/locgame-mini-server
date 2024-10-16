#! /usr/bin/bash

minikube start --driver=docker
minikube addons enable registry
kubectl create ns locg-service
kubectl config set-context --current --namespace=locg-service
helm install locg-nats nats/nats --namespace locg-service
helm install locg-redis bitnami/redis --namespace locg-service

#docker run -p 1358:1358 -d appbaseio/dejavu
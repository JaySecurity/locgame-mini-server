#! /usr/bin/bash

kubectl apply -f ./deploy/kubernetes/local/Mongo.yaml --namespace locg-service
kubectl apply -f ./deploy/kubernetes/local/locg-service-deployment.yaml --namespace locg-service
kubectl apply -f ./deploy/kubernetes/local/jobs-service-deployment.yaml --namespace locg-service
kubectl apply -f ./deploy/kubernetes/secrets/secret-local.yaml --namespace locg-service
kubectl apply -f ./deploy/kubernetes/local/service.yaml --namespace locg-service
kubectl apply -f ./deploy/kubernetes/local/elasticsearch-deployment.yaml  --namespace locg-service
#! /usr/bin/bash

kubectl port-forward --address 0.0.0.0 service/locg-service 8080:8080
kubectl port-forward --address 0.0.0.0 service/locg-service 8080:8080 &
kubectl port-forward --address 0.0.0.0 svc/exposed-service 27017:27017 & #add mongo

kubectl port-forward service/elasticsearch 9200:9200

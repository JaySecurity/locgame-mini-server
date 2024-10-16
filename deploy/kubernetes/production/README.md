# Introduction
These are instructions to run locg-service in Kubernetes.
Everything is assumed to be run from locg-service root.

# Set up cluster

## Create namespace
Since we aim to be fully isolated, we'll create a new `namespace` and deploy into it
```bash
kubectl create ns locg-prod
``` 

## Add NATS
```bash
helm repo add nats https://nats-io.github.io/k8s/helm/charts/    
helm install locg-nats nats/nats --namespace locg-prod
```

## Add Redis
```bash
helm repo add bitnami https://charts.bitnami.com/bitnami 
helm install locg-redis bitnami/redis --set master.persistence.enabled=false --set replica.persistence.enabled=false --namespace locg-prod
```

## Add Secrets
It is required to change the values of secrets in the file: `secrets.yaml`

After that:
```bash
kubectl apply -f secrets.yaml --namespace locg-prod
```

## Add Registry secrets
It is required to create Deploy Token for Docker and create this secret in Kubernetes

## Add locg-service and jobs-service
It is required to specify the correct address for MongoDB Atlas (`DATABASE_HOST`) and specify the correct address for Docker Image (`containers.image`).
```bash
kubectl apply -f locg-service-deployment.yaml --namespace locg-prod
kubectl apply -f jobs-service-deployment.yaml --namespace locg-prod
```
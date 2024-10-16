#! /usr/bin/bash
eval "$(minikube -p minikube docker-env)"

# create builder image
docker build --target builder -t locg-jobs-service:builder -f ./cmd/locgame-jobs-service/build/Dockerfile .

docker build -t localhost:5000/locg-service .
docker push localhost:5000/locg-service
kubectl rollout restart deployment/locg-service --namespace locg-service

docker build -t localhost:5000/locg-service/jobs-service -f ./cmd/locgame-jobs-service/build/Dockerfile .
docker push localhost:5000/locg-service/jobs-service
kubectl rollout restart deployment/locg-jobs-service --namespace locg-service

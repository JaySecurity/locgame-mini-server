#! /usr/bin/bash
#minikube mount /mnt/d/git-projects/locgame-service:/mnt/codes
minikube ssh -- cp -r /mnt/codes/internal /tmp/hostpath-provisioner/locg-service/my-code
kubectl exec -it locg-service-649959fdc4-plj6f -- rm -rf /build/internal
kubectl exec -it locg-service-649959fdc4-plj6f -- mv /build/codes/internal /build
kubectl exec -it locg-service-649959fdc4-plj6f -- go build -o /build/locg-service ./cmd/locgame-mini-server/service.go
kubectl exec -it locg-service-649959fdc4-plj6f -- killall locg-service
#kubectl exec -d locg-service-649959fdc4-plj6f -- /build/locg-service
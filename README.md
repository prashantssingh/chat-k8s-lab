# chat-k8s-lab

Hands-on microservices lab using **Go + Python**, deployed on **Kubernetes** with
multi-pod load balancing.  
Built for interview preparation and production-style learning.

---

## Tech Stack

- Go (client service)
- Python (FastAPI service)
- Docker
- Kubernetes (kind)
- kubectl

---

## Prerequisites

- Docker Desktop (running)
- Homebrew (macOS)
- kubectl
- kind

Install tools:
```bash
brew install kubernetes-cli kind

# verify 
kubectl version --client
kind version
docker info
```

---

## Local Kubernetes Setup (kind)
This project uses kind (Kubernetes in Docker) to create a local multi-node cluster.

### Create the Cluster
```bash
kind create cluster --config kind-cluster.yaml

# verify
kubectl cluster-info
kubectl get nodes -o wide
```

---

## Build and Load Docker Images
kind does not automatically see local Docker images.
Images must be loaded explicitly.

### Build python-chat image and Load into kind
```bash
docker build -t python-chat:0.1 ./services/python-chat

kind load docker-image python-chat:0.1 --name chat-lab
```

---

## Deploy python-chat to Kubernetes

Apply the Kubernetes manifests:

```bash
kubectl apply -f deploy/k8s/python-chat.yaml

# Watch the pods come up:
kubectl get pods -l app=python-chat -w

# verify the service 
kubectl get svc python-chat
```

---

## Validate Load Balancing

Run a temporary pod inside the cluster:

```bash
kubectl run tmp --rm -it --image=busybox -- sh
```

### Inside the pod, send multiple requests:
```bash
for i in 1 2 3 4 5 6 7 8 9 10; do
  wget -qO- http://python-chat/reply \
    --post-data='{"message":"hello","trace_id":"lb-test"}' \
    --header='Content-Type: application/json'
  echo
  sleep 2
done
```

### In another terminal, observe logs:
```bash
kubectl logs -l app=python-chat --tail=50
```
Each response includes the pod name, proving traffic is load-balanced across replicas.

---

## Cleanup
Delete the local cluster:
```bash
kind delete cluster --name chat-lab
```
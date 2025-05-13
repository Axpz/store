#!/bin/bash

set -e

# Define the kind cluster name
CLUSTER_NAME="axpz"
CONFIG_FILE="kind-config.yaml"

# Create the kind configuration file
cat <<EOF > "$CONFIG_FILE"
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
  extraPortMappings:
  - containerPort: 30080
    hostPort: 30080
  - containerPort: 30443
    hostPort: 30443
- role: worker
- role: worker
EOF

# Create the kind cluster
kind create cluster --config "$CONFIG_FILE" --name "$CLUSTER_NAME"

# - --kubelet-insecure-tls should be add to deployment
kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml

# Get cluster information
kubectl cluster-info --context "kind-$CLUSTER_NAME"

# Get the list of nodes
kubectl get nodes --context "kind-$CLUSTER_NAME"

kubectl top nodes && kubectl top pods -A

echo "Kind cluster '$CLUSTER_NAME' has been created and its information displayed."


helm repo add traefik https://helm.traefik.io/traefik --force-update
kubectl create namespace traefik --dry-run -o yaml | kubectl apply -f -
helm upgrade --install traefik-ingress traefik/traefik \
  --namespace traefik \
  --set api.insecure=true \
  --set api.dashboard=true \
  --set ui.enabled=true \
  --set service.type=NodePort \
  --set ports.web.nodePort=30080 \
  --set ports.websecure.nodePort=30443


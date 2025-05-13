#!/bin/bash

set -e

# Define the kind cluster name
CLUSTER_NAME="axpz"
CONFIG_FILE="kind-config.yaml"

if ! kind get clusters | grep -q "^${CLUSTER_NAME}$"; then
  echo "🛠 Creating Kind cluster: $CLUSTER_NAME"

  # Create Kind config file
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

  sleep 3

  # - --kubelet-insecure-tls should be add to deployment
  kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml

  # Get cluster information
  kubectl cluster-info --context "kind-$CLUSTER_NAME"

  # Get the list of nodes
  kubectl get nodes --context "kind-$CLUSTER_NAME"

  # Show cluster info
  kubectl cluster-info --context "kind-$CLUSTER_NAME"
  kubectl get nodes --context "kind-$CLUSTER_NAME"
  kubectl top nodes || true
  kubectl top pods -A || true
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
fi

# Load Docker images into the Kind cluster
echo "📦 Loading Docker images"
kind load docker-image store-frontend:latest --name "$CLUSTER_NAME"
kind load docker-image store:latest --name "$CLUSTER_NAME"

# Deploy the application
echo "🚀 Deploying application..."
kubectl create configmap zxenv --from-env-file="$HOME/.zxenv" || true
kubectl apply -f store-app.yaml

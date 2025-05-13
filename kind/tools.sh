#!/bin/bash

set -e

HELM_VERSION="v3.17.2"
KIND_VERSION="v0.27.0"
KUBECTL_VERSION="v1.32.2"

OS=$(uname -s)
ARCH=$(uname -m)

echo "Starting installation check for Helm, Kind, kubectl, and cloudflared..."

# --- Check and Install Helm ---
if ! command -v helm &> /dev/null; then
  echo "\n--- Helm not found. Installing Helm ${HELM_VERSION} ---"
  case "$ARCH" in
    x86_64) HELM_ARCH="linux-amd64" ;;
    arm64)  HELM_ARCH="linux-arm64" ;;
    aarch64) HELM_ARCH="linux-arm64" ;; # Some systems report aarch64
    Darwin)
      case "$ARCH" in
        x86_64) HELM_ARCH="darwin-amd64" ;;
        arm64)  HELM_ARCH="darwin-arm64" ;;
      esac
      ;;
    *)
      echo "Unsupported architecture: $ARCH for Helm"
      HELM_ARCH=""
      ;;
  esac

  if [ -n "$HELM_ARCH" ]; then
    HELM_URL="https://get.helm.sh/helm-${HELM_VERSION}-${HELM_ARCH}.tar.gz"
    echo "Downloading Helm from: $HELM_URL"
    curl -sSL "$HELM_URL" -o helm.tar.gz
    tar -zxvf helm.tar.gz
    sudo mv "./${HELM_ARCH}/helm" /usr/local/bin/helm
    rm -rf "./${HELM_ARCH}" helm.tar.gz
    echo "Helm ${HELM_VERSION} installed. Run 'helm version' to verify."
  else
    echo "Skipping Helm installation."
  fi
else
  echo "Helm found in PATH. Skipping installation."
fi

# --- Check and Install Kind ---
if ! command -v kind &> /dev/null; then
  echo "\n--- Kind not found. Installing Kind ${KIND_VERSION} ---"
  case "$ARCH" in
    x86_64) KIND_ARCH="amd64" ;;
    arm64)  KIND_ARCH="arm64" ;;
    aarch64) KIND_ARCH="arm64" ;; # Some systems report aarch64
    *)
      echo "Unsupported architecture: $ARCH for Kind"
      KIND_ARCH=""
      ;;
  esac

  if [ -n "$KIND_ARCH" ]; then
    KIND_URL="https://github.com/kubernetes-sigs/kind/releases/download/${KIND_VERSION}/kind-${OS}-${KIND_ARCH}"
    echo "Downloading Kind from: $KIND_URL"
    curl -sSL "$KIND_URL" -o kind
    chmod +x kind
    sudo mv kind /usr/local/bin/kind
    echo "Kind ${KIND_VERSION} installed. Run 'kind version' to verify."
  else
    echo "Skipping Kind installation."
  fi
else
  echo "Kind found in PATH. Skipping installation."
fi

# --- Check and Install kubectl ---
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
chmod +x kubectl
sudo mv kubectl /usr/local/bin/kubectl

# --- Check and Install cloudflared ---
if ! command -v cloudflared &> /dev/null; then
  echo "\n--- cloudflared not found. Installing cloudflared (latest stable) ---"
  case "$OS" in
    Linux)
      case "$ARCH" in
        x86_64) CLOUDFLARED_ARCH="amd64" CLOUDFLARED_URL="https://github.com/cloudflare/cloudflared/releases/latest/download/cloudflared-linux-$CLOUDFLARED_ARCH" ;;
        arm64)  CLOUDFLARED_ARCH="arm64" CLOUDFLARED_URL="https://github.com/cloudflare/cloudflared/releases/latest/download/cloudflared-linux-$CLOUDFLARED_ARCH" ;;
        aarch64) CLOUDFLARED_ARCH="arm64" CLOUDFLARED_URL="https://github.com/cloudflare/cloudflared/releases/latest/download/cloudflared-linux-$CLOUDFLARED_ARCH" ;;
        armv7l) CLOUDFLARED_ARCH="arm" CLOUDFLARED_URL="https://github.com/cloudflare/cloudflared/releases/latest/download/cloudflared-linux-$CLOUDFLARED_ARCH" ;;
        *)
          echo "Unsupported Linux architecture: $ARCH for cloudflared"
          CLOUDFLARED_URL=""
          ;;
      esac
      ;;
    Darwin)
      case "$ARCH" in
        x86_64) CLOUDFLARED_ARCH="amd64" CLOUDFLARED_URL="https://github.com/cloudflare/cloudflared/releases/latest/download/cloudflared-darwin-$CLOUDFLARED_ARCH" ;;
        arm64)  CLOUDFLARED_ARCH="arm64" CLOUDFLARED_URL="https://github.com/cloudflare/cloudflared/releases/latest/download/cloudflared-darwin-$CLOUDFLARED_ARCH" ;;
        *)
          echo "Unsupported macOS architecture: $ARCH for cloudflared"
          CLOUDFLARED_URL=""
          ;;
      esac
      ;;
    *)
      echo "Unsupported operating system: $OS for cloudflared"
      CLOUDFLARED_URL=""
      ;;
  esac

  if [ -n "$CLOUDFLARED_URL" ]; then
    echo "Downloading cloudflared from: $CLOUDFLARED_URL"
    curl -sSL "$CLOUDFLARED_URL" -o cloudflared
    chmod +x cloudflared
    sudo mv cloudflared /usr/local/bin/cloudflared
    echo "cloudflared installed. Run 'cloudflared --version' to verify."
  else
    echo "Skipping cloudflared installation due to unsupported OS or architecture."
  fi
else
  echo "cloudflared found in PATH. Skipping installation."
fi

echo "\nInstallation check and execution complete."
echo "Please verify each tool's installation if you haven't already."

#!/bin/bash
set -e

# Make scripts executable
chmod +x scripts/*.sh

# Install K3s
./scripts/install_k3s.sh

# Set KUBECONFIG for current session if not set
if [ -z "$KUBECONFIG" ]; then
    export KUBECONFIG=~/.kube/k3s-config
fi

# Wait for K3s to be ready
echo "Waiting for K3s to be ready..."
until kubectl get nodes &>/dev/null; do
    echo "Waiting for Kubernetes API..."
    sleep 5
done

# Install Ollama
./scripts/install_ollama.sh

# Install monitoring stack
./scripts/install_monitoring.sh

# Install lorvi
echo "Installing lorvi..."
make build
mkdir -p ~/bin
cp lorvi ~/bin/

# Add ~/bin to PATH for current session and permanently
export PATH="$HOME/bin:$PATH"

# Add to shell config if not already there
PATH_LINE='export PATH="$HOME/bin:$PATH"'
SHELL_CONFIG=""

if [ -f "$HOME/.zshrc" ]; then
    SHELL_CONFIG="$HOME/.zshrc"
elif [ -f "$HOME/.bashrc" ]; then
    SHELL_CONFIG="$HOME/.bashrc"
fi

if [ -n "$SHELL_CONFIG" ] && ! grep -q "$PATH_LINE" "$SHELL_CONFIG"; then
    echo "$PATH_LINE" >> "$SHELL_CONFIG"
fi

# Verify lorvi is in PATH
if ! which lorvi >/dev/null; then
    echo "Error: lorvi not found in PATH. Current PATH: $PATH"
    echo "lorvi binary location: $HOME/bin/lorvi"
    exit 1
fi

# Show lorvi version to confirm installation
echo "Lorvi installed successfully at: $(which lorvi)"
lorvi version

# Create/Update test namespace and workload
echo "Setting up test workload..."
kubectl create namespace lorvi-test 2>/dev/null || true

# Check if nginx deployment exists
if ! kubectl -n lorvi-test get deployment nginx &>/dev/null; then
    kubectl -n lorvi-test create deployment nginx --image=nginx
    kubectl -n lorvi-test expose deployment nginx --port=80
else
    echo "Nginx deployment already exists in lorvi-test namespace"
fi

# Update monitoring annotations
kubectl -n lorvi-test patch deployment nginx --patch '{
  "spec": {
    "template": {
      "metadata": {
        "annotations": {
          "prometheus.io/scrape": "true",
          "prometheus.io/port": "80"
        }
      }
    }
  }
}' || true

echo "Test environment setup complete!"
echo "Lorvi is installed at: $(which lorvi)"
echo "Version: $(lorvi version)"
echo ""
echo "Try these commands (they should work now without restarting your shell):"
echo "lorvi kubectl get pods -n lorvi-test"
echo "lorvi kubectl get svc -n lorvi-test"
echo "lorvi kubectl port-forward -n monitoring svc/monitoring-grafana 3000:80"

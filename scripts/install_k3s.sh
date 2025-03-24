#!/bin/bash
set -e

echo "Checking K3s status..."

# Check if K3s instance exists and is running
if multipass info k3s &> /dev/null; then
    echo "K3s instance already exists and running. Skipping installation."
    exit 0
fi

echo "Installing K3s on macOS..."

# Install multipass if not present
if ! command -v multipass &> /dev/null; then
    brew install --cask multipass
fi

# Launch K3s VM
multipass launch --name k3s --cpus 2 --memory 4G --disk 10G

# Install K3s inside VM
multipass exec k3s -- bash -c 'curl -sfL https://get.k3s.io | sh -'

# Get kubeconfig
mkdir -p ~/.kube
multipass exec k3s -- sudo cat /etc/rancher/k3s/k3s.yaml > ~/.kube/k3s-config
sed -i '' "s/127.0.0.1/$(multipass info k3s | grep IPv4 | awk '{print $2}')/g" ~/.kube/k3s-config

# Set KUBECONFIG in shell config
KUBECONFIG_LINE='export KUBECONFIG=~/.kube/k3s-config'
SHELL_CONFIG=""

if [ -f "$HOME/.zshrc" ]; then
    SHELL_CONFIG="$HOME/.zshrc"
elif [ -f "$HOME/.bashrc" ]; then
    SHELL_CONFIG="$HOME/.bashrc"
fi

if [ -n "$SHELL_CONFIG" ] && ! grep -q "$KUBECONFIG_LINE" "$SHELL_CONFIG"; then
    echo "$KUBECONFIG_LINE" >> "$SHELL_CONFIG"
fi

# Set KUBECONFIG for current session
export KUBECONFIG=~/.kube/k3s-config

echo "K3s installed successfully!"
echo "Kubeconfig has been configured at ~/.kube/k3s-config"
echo "KUBECONFIG environment variable set for current session"
echo "Please run: source ~/.zshrc or source ~/.bashrc"
echo "Test with: kubectl get nodes"

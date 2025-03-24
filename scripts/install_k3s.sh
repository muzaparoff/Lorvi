#!/bin/bash
set -e

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

# Set KUBECONFIG
echo 'export KUBECONFIG=~/.kube/k3s-config' >> ~/.zshrc

echo "K3s installed successfully!"
echo "Please run: source ~/.zshrc"
echo "Test with: kubectl get nodes"

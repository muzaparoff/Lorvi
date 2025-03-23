#!/bin/bash
set -e

# Make scripts executable
chmod +x scripts/*.sh

# Install K3s
./scripts/install_k3s.sh

# Install Ollama
./scripts/install_ollama.sh

# Install monitoring stack
./scripts/install_monitoring.sh

# Create test namespace
kubectl create namespace lorvi-test

# Deploy test workload
kubectl -n lorvi-test create deployment nginx --image=nginx
kubectl -n lorvi-test expose deployment nginx --port=80

# Add monitoring annotations
kubectl -n lorvi-test patch deployment nginx -p '{
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
}'

echo "Test environment setup complete!"
echo "Try: lorvi kubectl get pods -n lorvi-test"
echo "Try: lorvi kubectl get svc -n lorvi-test"
echo "Try: lorvi kubectl port-forward -n monitoring svc/monitoring-grafana 3000:80"

#!/bin/bash
set -e

echo "Checking monitoring stack status..."

# Check if monitoring namespace exists and has deployments
if kubectl get namespace monitoring &> /dev/null && \
   kubectl get deployments -n monitoring &> /dev/null; then
    echo "Monitoring stack already installed. Skipping installation."
    exit 0
fi

echo "Installing monitoring stack..."

# Install Helm if not present
if ! command -v helm &> /dev/null; then
    echo "Installing Helm..."
    brew install helm
fi

# Add Helm repositories
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo add grafana https://grafana.github.io/helm-charts
helm repo update

# Create monitoring namespace
kubectl create namespace monitoring || true

# Install kube-prometheus-stack (includes Prometheus and Grafana)
helm upgrade --install monitoring prometheus-community/kube-prometheus-stack \
  --namespace monitoring \
  --set grafana.adminPassword=admin123 \
  --set grafana.persistence.enabled=true \
  --set prometheus.prometheusSpec.retention=5d

# Install Loki and Promtail
helm upgrade --install loki grafana/loki-stack \
  --namespace monitoring \
  --set grafana.enabled=false \
  --set promtail.enabled=true \
  --set loki.persistence.enabled=true

echo "Monitoring stack installed successfully!"
echo "Access Grafana:"
echo "1. Run: kubectl port-forward -n monitoring svc/monitoring-grafana 3000:80"
echo "2. Open: http://localhost:3000"
echo "3. Login with admin/admin123"

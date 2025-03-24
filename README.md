# Lorvi

[![Build Status](https://github.com/muzaparoff/lorvi/actions/workflows/build.yml/badge.svg)](https://github.com/muzaparoff/lorvi/actions)
[![Security Scan](https://img.shields.io/badge/Security-passing-brightgreen)](#)
[![Latest Release](https://img.shields.io/github/v/release/muzaparoff/lorvi)](https://github.com/muzaparoff/lorvi/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/muzaparoff/lorvi)](https://goreportcard.com/report/github.com/muzaparoff/lorvi)

Lorvi is a terminal-based, AI-powered DevOps assistant built in Golang. It provides a unified CLI interface for common DevOps tools like **kubectl** and **terraform**, integrates with multiple AI backends (Ollama, OpenAI, Claude, Gemini), and supports both local and cloud API integrations (AWS, Azure, GCP).

## Features

- **CLI Integration with Cobra:**  
  Modular commands for various tools with dynamic flags.
  - **kubectl:** Execute Kubernetes commands with an environment/cluster context.
  - **terraform:** Run Terraform commands using environment-specific variable files.

- **AI Backend Integration:**  
  Easily switch between AI backends using the `--ai-backend` flag.
  - Options include: **ollama**, **openai**, **claude**, and **gemini**.

- **Cloud API Support:**  
  Validate and use cloud provider credentials for AWS, Azure, and GCP.
  - Use the `--cloud` flag to target cloud providers. Defaults to local mode if not specified.

- **CI/CD with GitHub Actions:**  
  - Build macOS binaries automatically.
  - Automated security scanning with `gosec` and `govulncheck`.
  - GitHub Releases with semantic versioning (starting from **v0.0.1**).  
  Semantic versioning follows:
    - **Patch**: Bug fixes.
    - **Minor**: New features.
    - **Major**: Breaking changes or API modifications.

## Getting Started

### Prerequisites

- **Go 1.20+**
- CLI tools installed: `kubectl`, `terraform`
- Cloud provider credentials set in environment variables (if using cloud features):
  - **AWS:** `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`
  - **Azure:** Appropriate Azure CLI authentication or service principal details.
  - **GCP:** `GOOGLE_APPLICATION_CREDENTIALS`

### Installation

Clone the repository and build Lorvi:

```bash
git clone https://github.com/muzaparoff/lorvi.git
cd lorvi
go build -o lorvi .
```

## Usage Examples
### Running kubectl commands:
```bash
# Execute a kubectl command targeting the "prod" environment.
./lorvi kubectl get pods -e prod
```

### Running terraform commands:
```bash
# Execute a terraform plan using the "staging.tfvars" file.
./lorvi terraform plan -e staging
```

### Using AI Features (stub implementation):
```bash
# Ask the configured AI backend for help.
./lorvi --ai-backend openai -- ask "How do I resolve a CrashLoopBackOff error?"
```

### Using Cloud Integration:
```bash
# Run a command with cloud provider context (e.g., AWS).
./lorvi kubectl get services -e prod --cloud aws
```

## Test Environment

Lorvi comes with a complete test environment setup including:
- Local K3s cluster (via Multipass)
- Ollama AI backend
- Full monitoring stack

### Prerequisites
- macOS (tested on M1/M2)
- Homebrew
- Go 1.20+

### Quick Start

1. Clone and build:
```bash
git clone https://github.com/muzaparoff/lorvi.git
cd lorvi
make build
```

2. Set up test environment:
```bash
chmod +x scripts/*.sh
make setup-test-env
```

This will:
- Install K3s via Multipass
- Install Ollama with CodeLlama model
- Deploy monitoring stack (Prometheus, Loki, Grafana)
- Create test workload

### Monitoring Stack

The test environment includes:
- Prometheus for metrics collection
- Loki for log aggregation
- Grafana for visualization

Access Grafana:
```bash
make grafana-dashboard
```
Then open http://localhost:3000 (admin/admin123)

### Test Commands

Try these commands:
```bash
# View test pods
lorvi kubectl get pods -n lorvi-test

# View monitoring stack
lorvi kubectl get pods -n monitoring

# Access Grafana
make grafana-dashboard
```

### Cleanup

Remove test environment:
```bash
make clean-test-env
make clean-monitoring
```

## CI/CD and Releases

Lorvi uses GitHub Actions for continuous integration:
- **Build & Test:**
  Builds a macOS binary, runs tests, and performs security scans.
- **Security Scanning:**
  Uses gosec and govulncheck to ensure code security.
- **Releases:**
  Semantic versioning is managed using Git tags (starting at v0.0.1).
  On release, GitHub Actions build binaries and attach them to the GitHub Release.

## Contributing

Contributions are welcome! Please see CONTRIBUTING.md for guidelines.

## License

This project is licensed under the MIT License – see the LICENSE file for details.

## Security Badge

---

### Final Notes

- **Extensibility:**  
  - To add new DevOps tools, create additional subcommands under `cmd/` and implement their logic in separate packages if needed.
  - To integrate a new AI backend, extend the `internal/ai/ai.go` file with a new client that implements the `AIClient` interface.

- **Versioning:**  
  - Use Git tags for versioning (e.g., `git tag v0.0.1`), and automate version bumps and changelog generation via GitHub Actions or tools like [release-please](https://github.com/googleapis/release-please).

- **Security:**  
  - Ensure that you update the stub functions to perform real cloud credential checks and AI API integrations as you extend the project.

You can now open this project in VSCode, customize it further, and push it to GitHub to start developing Lorvi!
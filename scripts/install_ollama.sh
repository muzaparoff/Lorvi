#!/bin/bash
set -e

echo "Installing Ollama on macOS..."

# Install Ollama
if ! command -v ollama &> /dev/null; then
    curl -fsSL https://ollama.com/install.sh | sh
fi

# Start Ollama service
ollama serve &

# Pull CodeLlama model for development
echo "Pulling CodeLlama model..."
ollama pull codellama

echo "Ollama installed successfully!"
echo "Test with: ollama run codellama 'Write a hello world in Go'"

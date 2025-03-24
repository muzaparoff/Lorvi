#!/bin/bash
set -e

echo "Checking Ollama status..."

# Check if Ollama is running
if pgrep -x "ollama" > /dev/null; then
    echo "Ollama is already running. Skipping installation."
    exit 0
fi

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

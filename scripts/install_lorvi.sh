#!/bin/bash
set -e

echo "Installing lorvi..."

# Build lorvi
make build

# Create bin directory if it doesn't exist
mkdir -p ~/bin

# Copy lorvi to bin directory
cp lorvi ~/bin/

# Add ~/bin to PATH if not already there
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

echo "Lorvi installed successfully!"
echo "Please run: source $SHELL_CONFIG"
echo "Then try: lorvi version"

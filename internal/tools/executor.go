package tools

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
)

var (
	// Allowlist of safe characters in command arguments
	safeArgPattern = regexp.MustCompile(`^[a-zA-Z0-9\-\._/=@:]+$`)
	// Allowlist of safe characters in command paths
	safePathPattern = regexp.MustCompile(`^[a-zA-Z0-9\-\._/]+$`)
	// Whitelist of allowed commands
	allowedCommands = map[string]bool{
		"kubectl":   true,
		"terraform": true,
	}
)

// SecureCommandExecutor handles safe command execution
type SecureCommandExecutor struct {
	allowedCommands map[string]bool
	commandPaths    map[string]string
}

// NewSecureCommandExecutor creates a new executor with allowed commands
func NewSecureCommandExecutor(commands []string) *SecureCommandExecutor {
	allowed := make(map[string]bool)
	paths := make(map[string]string)
	for _, cmd := range commands {
		if path, err := exec.LookPath(cmd); err == nil {
			cleanPath := filepath.Clean(path)
			if safePathPattern.MatchString(cleanPath) {
				allowed[cleanPath] = true
				paths[cmd] = cleanPath
			}
		}
	}
	return &SecureCommandExecutor{
		allowedCommands: allowed,
		commandPaths:    paths,
	}
}

// ValidateArgs checks if command arguments are safe
func (e *SecureCommandExecutor) ValidateArgs(args []string) error {
	for _, arg := range args {
		if !safeArgPattern.MatchString(arg) {
			return fmt.Errorf("invalid argument format: %s", arg)
		}
	}
	return nil
}

// Execute runs a command securely
func (e *SecureCommandExecutor) Execute(command string, args []string) ([]byte, error) {
	if !allowedCommands[command] {
		return nil, fmt.Errorf("command not in whitelist: %s", command)
	}

	// Get pre-validated command path
	path, ok := e.commandPaths[command]
	if !ok {
		return nil, fmt.Errorf("command not allowed: %s", command)
	}

	if err := e.ValidateArgs(args); err != nil {
		return nil, err
	}

	// #nosec G204 -- path and args are validated above
	cmd := exec.Command(path, args...)
	cmd.Env = os.Environ()
	return cmd.CombinedOutput()
}

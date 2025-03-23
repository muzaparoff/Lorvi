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
)

// SecureCommandExecutor handles safe command execution
type SecureCommandExecutor struct {
	allowedCommands map[string]bool
}

// NewSecureCommandExecutor creates a new executor with allowed commands
func NewSecureCommandExecutor(commands []string) *SecureCommandExecutor {
	allowed := make(map[string]bool)
	for _, cmd := range commands {
		if path, err := exec.LookPath(cmd); err == nil {
			allowed[filepath.Clean(path)] = true
		}
	}
	return &SecureCommandExecutor{allowedCommands: allowed}
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
	path, err := exec.LookPath(command)
	if err != nil {
		return nil, fmt.Errorf("command not found: %s", command)
	}

	if !e.allowedCommands[filepath.Clean(path)] {
		return nil, fmt.Errorf("command not allowed: %s", command)
	}

	if err := e.ValidateArgs(args); err != nil {
		return nil, err
	}

	cmd := exec.Command(path, args...)
	cmd.Env = os.Environ()
	return cmd.CombinedOutput()
}

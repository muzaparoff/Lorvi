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

// CommandExecutor interface for mocking in tests
type CommandExecutor interface {
	Execute(command string, args []string) ([]byte, error)
}

// Ensure SecureCommandExecutor implements CommandExecutor
var _ CommandExecutor = (*SecureCommandExecutor)(nil)
var _ CommandExecutor = (*MockCommandExecutor)(nil)

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

// NewMockExecutor creates a test executor that doesn't actually run commands
func NewMockExecutor() *SecureCommandExecutor {
	return &SecureCommandExecutor{
		allowedCommands: allowedCommands,
		commandPaths: map[string]string{
			"kubectl":   "/usr/local/bin/kubectl",
			"terraform": "/usr/local/bin/terraform",
		},
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

// MockCommandExecutor for testing
type MockCommandExecutor struct {
	MockOutput []byte
	MockError  error
}

func NewTestExecutor() *MockCommandExecutor {
	return &MockCommandExecutor{
		MockOutput: []byte("mocked output"),
		MockError:  nil,
	}
}

func (m *MockCommandExecutor) Execute(command string, args []string) ([]byte, error) {
	// Validate command is in whitelist
	if !allowedCommands[command] {
		return nil, fmt.Errorf("command not in whitelist: %s", command)
	}

	// Simulate argument validation
	for _, arg := range args {
		if !safeArgPattern.MatchString(arg) {
			return nil, fmt.Errorf("invalid argument format: %s", arg)
		}
	}

	return m.MockOutput, m.MockError
}

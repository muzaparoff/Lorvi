package cmd

import (
	"testing"

	"github.com/muzaparoff/lorvi/internal/tools"
)

func TestRunKubectl(t *testing.T) {
	// Create mock executor and save original
	mockExecutor := tools.NewTestExecutor()
	savedExecutor := executor
	// Replace global executor with mock
	executor = mockExecutor
	// Restore original executor after test
	defer func() { executor = savedExecutor }()

	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{"valid args", []string{"get", "pods"}, false},
		{"invalid args", []string{";dangerous"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RunKubectl(tt.args); (err != nil) != tt.wantErr {
				t.Errorf("RunKubectl() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

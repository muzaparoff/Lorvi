package cmd

import (
	"testing"

	"github.com/yourusername/lorvi/internal/tools"
)

func TestRunKubectl(t *testing.T) {
	mockExecutor := tools.NewTestExecutor()
	savedExecutor := executor
	executor = &tools.SecureCommandExecutor{CommandExecutor: mockExecutor}
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

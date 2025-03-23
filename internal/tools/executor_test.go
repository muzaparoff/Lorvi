package tools

import (
	"testing"
)

func TestSecureCommandExecutor(t *testing.T) {
	tests := []struct {
		name    string
		cmd     string
		args    []string
		wantErr bool
	}{
		{"valid kubectl command", "kubectl", []string{"get", "pods"}, false},
		{"valid terraform command", "terraform", []string{"plan"}, false},
		{"invalid command", "malicious", []string{}, true},
		{"invalid args", "kubectl", []string{";dangerous"}, true},
	}

	executor := NewSecureCommandExecutor([]string{"kubectl", "terraform"})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := executor.Execute(tt.cmd, tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

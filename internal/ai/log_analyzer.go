package ai

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"

	"github.com/muzaparoff/lorvi/internal/tools"
)

var (
	// Safe namespace pattern
	safeNamespacePattern = regexp.MustCompile(`^[a-z0-9][a-z0-9-]*[a-z0-9]$`)
	// Safe pod name pattern
	safePodNamePattern = regexp.MustCompile(`^[a-z0-9][a-z0-9-]*[a-z0-9]$`)
)

type LogAnalyzer struct {
	namespace string
	executor  tools.CommandExecutor
}

func NewLogAnalyzer(namespace string) *LogAnalyzer {
	return &LogAnalyzer{
		namespace: namespace,
		executor:  tools.NewSecureCommandExecutor([]string{"kubectl", "ollama"}),
	}
}

func (la *LogAnalyzer) validateNamespace() error {
	if !safeNamespacePattern.MatchString(la.namespace) {
		return fmt.Errorf("invalid namespace format: %s", la.namespace)
	}
	return nil
}

func (la *LogAnalyzer) validatePodName(pod string) error {
	if !safePodNamePattern.MatchString(pod) {
		return fmt.Errorf("invalid pod name format: %s", pod)
	}
	return nil
}

func (la *LogAnalyzer) AnalyzeLogs() (string, error) {
	// Get all pods in namespace
	pods, err := la.getPods()
	if err != nil {
		return "", err
	}

	if len(pods) == 0 {
		return "No pods found in namespace. This might indicate deployment issues or incorrect namespace.", nil
	}

	// Collect logs from all pods
	allLogs := make(map[string]string)
	hasLogs := false

	for _, pod := range pods {
		logs, _ := la.getPodLogs(pod)
		if logs != "" {
			hasLogs = true
			allLogs[pod] = logs
		}
	}

	if !hasLogs {
		return "No logs found in any pods. This might indicate that pods are not running properly or applications are not producing logs.", nil
	}

	// Prepare prompt for Ollama
	prompt := fmt.Sprintf(`Analyze these Kubernetes pod logs and provide:
1. Overall health assessment
2. Potential issues or warnings
3. Performance improvement suggestions
4. Best practices recommendations

Logs by pod:
%s`, la.formatLogsForPrompt(allLogs))

	// Get analysis from Ollama
	return la.getOllamaAnalysis(prompt)
}

func (la *LogAnalyzer) getPods() ([]string, error) {
	if err := la.validateNamespace(); err != nil {
		return nil, err
	}

	output, err := la.executor.Execute("kubectl", []string{
		"get", "pods",
		"-n", la.namespace,
		"-o", "jsonpath={.items[*].metadata.name}",
	})
	if err != nil {
		return nil, err
	}
	return strings.Fields(string(output)), nil
}

func (la *LogAnalyzer) getPodLogs(pod string) (string, error) {
	if err := la.validatePodName(pod); err != nil {
		return "", err
	}

	output, err := la.executor.Execute("kubectl", []string{
		"logs",
		"-n", la.namespace,
		pod,
		"--tail=100",
	})
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func (la *LogAnalyzer) formatLogsForPrompt(logs map[string]string) string {
	var buffer bytes.Buffer
	for pod, log := range logs {
		buffer.WriteString(fmt.Sprintf("\n=== Pod: %s ===\n%s\n", pod, log))
	}
	return buffer.String()
}

func (la *LogAnalyzer) getOllamaAnalysis(prompt string) (string, error) {
	output, err := la.executor.Execute("ollama", []string{"run", "codellama", prompt})
	if err != nil {
		return "", fmt.Errorf("ollama analysis failed: %v", err)
	}
	return string(output), nil
}

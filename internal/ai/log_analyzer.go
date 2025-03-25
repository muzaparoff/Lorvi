package ai

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

type LogAnalyzer struct {
	namespace string
}

func NewLogAnalyzer(namespace string) *LogAnalyzer {
	return &LogAnalyzer{namespace: namespace}
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
	cmd := exec.Command("kubectl", "get", "pods", "-n", la.namespace, "-o", "jsonpath={.items[*].metadata.name}")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	return strings.Fields(string(output)), nil
}

func (la *LogAnalyzer) getPodLogs(pod string) (string, error) {
	cmd := exec.Command("kubectl", "logs", "-n", la.namespace, pod, "--tail=100")
	output, err := cmd.CombinedOutput()
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
	cmd := exec.Command("ollama", "run", "codellama", prompt)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("ollama analysis failed: %v", err)
	}
	return string(output), nil
}

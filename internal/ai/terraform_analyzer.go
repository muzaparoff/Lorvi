package ai

import (
	"fmt"
	"os/exec"
)

type TerraformAnalyzer struct {
	workingDir string
}

func NewTerraformAnalyzer(workingDir string) *TerraformAnalyzer {
	return &TerraformAnalyzer{workingDir: workingDir}
}

func (ta *TerraformAnalyzer) AnalyzePlan() (string, error) {
	// Get terraform plan output
	planOutput, err := ta.getTerraformPlan()
	if err != nil {
		return "", err
	}

	if planOutput == "" {
		return "No changes. Infrastructure is up-to-date.", nil
	}

	// Prepare prompt for Ollama
	prompt := fmt.Sprintf(`Analyze this Terraform plan and provide:
1. Risk assessment of proposed changes
2. Security implications
3. Cost impact analysis
4. Best practices recommendations
5. Potential issues or warnings

Terraform Plan:
%s`, planOutput)

	return ta.getOllamaAnalysis(prompt)
}

func (ta *TerraformAnalyzer) AnalyzeState() (string, error) {
	// Get terraform state output
	stateOutput, err := ta.getTerraformState()
	if err != nil {
		return "", err
	}

	if stateOutput == "" {
		return "No state found. Infrastructure might not be initialized or may be empty.", nil
	}

	// Prepare prompt for Ollama
	prompt := fmt.Sprintf(`Analyze this Terraform state and provide:
1. Infrastructure health assessment
2. Security concerns
3. Cost optimization opportunities
4. Compliance checks
5. Architecture improvement suggestions

Terraform State:
%s`, stateOutput)

	return ta.getOllamaAnalysis(prompt)
}

func (ta *TerraformAnalyzer) getTerraformPlan() (string, error) {
	cmd := exec.Command("terraform", "plan", "-no-color")
	cmd.Dir = ta.workingDir
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("terraform plan failed: %v", err)
	}
	return string(output), nil
}

func (ta *TerraformAnalyzer) getTerraformState() (string, error) {
	cmd := exec.Command("terraform", "show", "-no-color")
	cmd.Dir = ta.workingDir
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("terraform show failed: %v", err)
	}
	return string(output), nil
}

func (ta *TerraformAnalyzer) getOllamaAnalysis(prompt string) (string, error) {
	cmd := exec.Command("ollama", "run", "codellama", prompt)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("ollama analysis failed: %v", err)
	}
	return string(output), nil
}

package cmd

import (
	"fmt"

	"github.com/muzaparoff/lorvi/internal/ai"
	"github.com/muzaparoff/lorvi/internal/tools"
	"github.com/spf13/cobra"
)

var executor tools.CommandExecutor = tools.NewSecureCommandExecutor([]string{"kubectl"})

func RunKubectl(cmdArgs []string) error {
	out, err := executor.Execute("kubectl", cmdArgs)
	if err != nil {
		return fmt.Errorf("error: %v\noutput: %s", err, out)
	}
	fmt.Printf("%s", out)
	return nil
}

var kubectlCmd = &cobra.Command{
	Use:                "kubectl [flags] [command]",
	Short:              "Run kubectl commands with environment context",
	Long:               `Execute kubectl commands for a specified environment or cluster.`,
	DisableFlagParsing: true, // This allows passing through all flags to kubectl
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide kubectl arguments, e.g., 'get pods'")
			return
		}

		// Check if this is a logs analysis request
		if args[0] == "analyze-logs" {
			namespace := "default"
			if len(args) > 1 {
				namespace = args[1]
			}
			analyzer := ai.NewLogAnalyzer(namespace)
			analysis, err := analyzer.AnalyzeLogs()
			if err != nil {
				fmt.Printf("Error analyzing logs: %v\n", err)
				return
			}
			fmt.Printf("\nAI Log Analysis for namespace %s:\n%s\n", namespace, analysis)
			return
		}

		// Build the command with environment context if provided
		cmdArgs := []string{}
		if env != "" {
			cmdArgs = append(cmdArgs, "--context", env)
		}
		cmdArgs = append(cmdArgs, args...)

		if err := RunKubectl(cmdArgs); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	},
}

package cmd

import (
	"fmt"

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

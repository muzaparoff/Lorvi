package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yourusername/lorvi/internal/tools"
)

var executor = tools.NewSecureCommandExecutor([]string{"kubectl"})

func RunKubectl(cmdArgs []string) error {
	out, err := executor.Execute("kubectl", cmdArgs)
	if err != nil {
		return fmt.Errorf("error: %v\noutput: %s", err, out)
	}
	fmt.Printf("Output: %s\n", out)
	return nil
}

var kubectlCmd = &cobra.Command{
	Use:   "kubectl",
	Short: "Run kubectl commands with environment context",
	Long:  `Execute kubectl commands for a specified environment or cluster.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide kubectl arguments, e.g., 'get pods'")
			return
		}

		// Build the command. This is a simple example.
		cmdArgs := args
		if env != "" {
			// Append --context flag to target the environment/cluster.
			cmdArgs = append([]string{"--context", env}, cmdArgs...)
		}
		if err := RunKubectl(cmdArgs); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	},
}

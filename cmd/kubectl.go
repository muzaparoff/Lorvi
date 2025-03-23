package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

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
		fmt.Printf("Running kubectl with args: %v\n", cmdArgs)
		out, err := exec.Command("kubectl", cmdArgs...).CombinedOutput()
		if err != nil {
			fmt.Printf("Error: %v\nOutput: %s\n", err, out)
		} else {
			fmt.Printf("Output: %s\n", out)
		}
	},
}

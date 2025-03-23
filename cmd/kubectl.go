package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

func validateKubectlArgs(args []string) error {
	for _, arg := range args {
		if strings.Contains(arg, ";") || strings.Contains(arg, "|") {
			return fmt.Errorf("invalid character in argument: %s", arg)
		}
	}
	return nil
}

func RunKubectl(cmdArgs []string) error {
	if err := validateKubectlArgs(cmdArgs); err != nil {
		return err
	}
	fmt.Printf("Running kubectl with args: %v\n", cmdArgs)
	cmd := exec.Command("kubectl", cmdArgs...)
	cmd.Env = os.Environ()
	out, err := cmd.CombinedOutput()
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

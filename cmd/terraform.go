package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var terraformCmd = &cobra.Command{
	Use:   "terraform",
	Short: "Run terraform commands with environment context",
	Long:  `Execute terraform commands for a specified environment or workspace.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide terraform arguments, e.g., 'plan'")
			return
		}

		// Build the command. Append environment flag if provided.
		cmdArgs := args
		if env != "" {
			// For example, pass env as a variable file.
			cmdArgs = append(cmdArgs, "-var-file="+env+".tfvars")
		}
		fmt.Printf("Running terraform with args: %v\n", cmdArgs)
		out, err := exec.Command("terraform", cmdArgs...).CombinedOutput()
		if err != nil {
			fmt.Printf("Error: %v\nOutput: %s\n", err, out)
		} else {
			fmt.Printf("Output: %s\n", out)
		}
	},
}

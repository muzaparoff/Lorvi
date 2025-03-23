package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

func validateTerraformArgs(args []string) error {
	for _, arg := range args {
		if strings.Contains(arg, ";") || strings.Contains(arg, "|") {
			return fmt.Errorf("invalid character in argument: %s", arg)
		}
	}
	return nil
}

func RunTerraform(cmdArgs []string) error {
	if err := validateTerraformArgs(cmdArgs); err != nil {
		return err
	}
	fmt.Printf("Running terraform with args: %v\n", cmdArgs)
	cmd := exec.Command("terraform", cmdArgs...)
	cmd.Env = os.Environ()
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error: %v\noutput: %s", err, out)
	}
	fmt.Printf("Output: %s\n", out)
	return nil
}

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
		if err := RunTerraform(cmdArgs); err != nil {
			fmt.Println(err)
		}
	},
}

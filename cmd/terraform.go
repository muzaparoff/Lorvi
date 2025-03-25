package cmd

import (
	"fmt"
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
	out, err := executor.Execute("terraform", cmdArgs)
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

		// Check if this is an analysis request
		if args[0] == "analyze" {
			analyzer := ai.NewTerraformAnalyzer(".")
			var analysis string
			var err error

			if len(args) > 1 && args[1] == "state" {
				analysis, err = analyzer.AnalyzeState()
			} else {
				analysis, err = analyzer.AnalyzePlan()
			}

			if err != nil {
				fmt.Printf("Error analyzing Terraform: %v\n", err)
				return
			}
			fmt.Printf("\nAI Terraform Analysis:\n%s\n", analysis)
			return
		}

		// Regular terraform command handling
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

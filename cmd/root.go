package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	env       string
	cloud     string
	aiBackend string
)

var rootCmd = &cobra.Command{
	Use:   "lorvi",
	Short: "Lorvi - AI-powered DevOps Assistant",
	Long: `Lorvi is a terminal-based DevOps assistant that integrates with 
various tools like kubectl and terraform, and supports AI backends 
such as Ollama, OpenAI, Claude, and Gemini.`,
	Run: func(cmd *cobra.Command, args []string) {
		// If no subcommand is provided, show help.
		cmd.Help()
	},
}

// Execute runs the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Global persistent flags
	rootCmd.PersistentFlags().StringVarP(&env, "environment", "e", "", "Environment or cluster name (e.g., dev, staging, prod)")
	rootCmd.PersistentFlags().StringVar(&cloud, "cloud", "", "Cloud provider (aws|azure|gcp)")
	rootCmd.PersistentFlags().StringVar(&aiBackend, "ai-backend", "ollama", "AI backend to use (ollama, openai, claude, gemini)")

	// Add subcommands
	rootCmd.AddCommand(kubectlCmd)
	rootCmd.AddCommand(terraformCmd)
}

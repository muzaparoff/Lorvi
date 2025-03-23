package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// Version is set during build
	Version = "dev"
	// Commit is set during build
	Commit = "none"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Lorvi",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Lorvi version %s (commit: %s)\n", Version, Commit)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

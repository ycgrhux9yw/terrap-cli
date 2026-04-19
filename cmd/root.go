package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// Version is set at build time via ldflags
	Version = "dev"
	// Commit is set at build time via ldflags
	Commit = "none"
)

// rootCmd is the base command for the terrap CLI.
// All subcommands are registered as children of this command.
var rootCmd = &cobra.Command{
	Use:   "terrap",
	Short: "Terrap - Terraform wrapper for multi-environment management",
	Long: `Terrap is a CLI tool that wraps Terraform to provide
simplified multi-environment and multi-region infrastructure management.

It helps teams manage Terraform workspaces, state, and deployments
across different environments with a consistent workflow.`,
	SilenceUsage:  true,
	SilenceErrors: true,
}

// versionCmd prints the current version of the CLI.
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of terrap",
	Aliases: []string{"v"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("terrap version %s (commit: %s)\n", Version, Commit)
	},
}

// Execute runs the root command and handles any top-level errors.
// This is called from main.go.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var (
	// Version variables injected at build time via -ldflags
	Version   = "1.0.1"
	Commit    = "none"
	BuildDate = "unknown"
)

// RootCmd represents the base command when called without any subcommands.
var RootCmd = &cobra.Command{
	Use:   "bender",
	Short: "Generic Go CLI, Self-Updater & MCP Server Engine",
	Long: fmt.Sprintf(`🤖 bender - Generic Go CLI, Self-Updater & MCP Server Engine (v%s)

Standardized CLI architecture with built-in GitHub Releases self-updater,
embedded Model Context Protocol (MCP) server, and progressive installation scaffolding.`, Version),
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

// ResetFlags recursively resets all flags on a command and its subcommands to their default values.
func ResetFlags(cmd *cobra.Command) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		_ = f.Value.Set(f.DefValue)
		f.Changed = false
	})
	for _, child := range cmd.Commands() {
		ResetFlags(child)
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() error {
	defer ResetFlags(RootCmd)
	return RootCmd.Execute()
}

func init() {
	RootCmd.Version = Version
	RootCmd.SetVersionTemplate(fmt.Sprintf("bender v{{.Version}} (commit: %s, built: %s)\n", Commit, BuildDate))
}

func getWorkingDir() string {
	cwd, err := os.Getwd()
	if err != nil {
		return "."
	}
	return cwd
}

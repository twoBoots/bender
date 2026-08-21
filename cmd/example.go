package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

// ExampleCmd represents a sample subcommand demonstrating CLI extensibility.
var ExampleCmd = &cobra.Command{
	Use:     "hello [name]",
	Aliases: []string{"greet"},
	Short:   "Sample greeting command demonstrating CLI subcommand extension",
	Long: `Demonstrates how projects consuming or scaffolding from Bender can add
new Cobra subcommands while leveraging built-in self-updater and MCP servers.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		name := "World"
		if len(args) > 0 && args[0] != "" {
			name = args[0]
		}
		return RunHello(cmd.OutOrStdout(), name)
	},
}

// RunHello prints a greeting message to out.
func RunHello(out io.Writer, name string) error {
	fmt.Fprintf(out, "🤖 Hello, %s! Welcome to Bender.\n", name)
	return nil
}

func init() {
	RootCmd.AddCommand(ExampleCmd)
}

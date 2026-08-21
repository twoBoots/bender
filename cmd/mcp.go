package cmd

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/spf13/cobra"
	"github.com/twoBoots/bender/pkg/mcp"
)

var (
	mcpTransport             string
	mcpInstallClients        string
	mcpInstallAll            bool
	mcpInstallNonInteractive bool
)

// MCPCmd represents the 'bender mcp' command.
var MCPCmd = &cobra.Command{
	Use:     "mcp",
	Aliases: []string{"serve"},
	Short:   "Start the Model Context Protocol (MCP) server over stdio",
	Long: `Start the Bender Model Context Protocol (MCP) server over stdio.

Exposes CLI orchestration tools, living context resources, and prompt templates to AI coding assistants (e.g. Antigravity, Claude Code, Cursor, Windsurf).`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return RunMCPServer(cmd.InOrStdin(), cmd.OutOrStdout(), getWorkingDir())
	},
}

// MCPInstallCmd represents the 'bender mcp install' command.
var MCPInstallCmd = &cobra.Command{
	Use:     "install",
	Aliases: []string{"setup", "configure"},
	Short:   "Configure Bender MCP server in AI coding assistants (Cursor, Antigravity, Claude, Windsurf, VS Code)",
	Long: `Automatically detect and configure Bender's MCP server in supported AI assistant configuration files.

Supported clients:
  * cursor          - Cursor IDE (.cursor/mcp.json)
  * antigravity     - Google Antigravity / agy (~/.gemini/config/mcp_config.json)
  * claude-desktop  - Anthropic Claude Desktop
  * claude-code     - Anthropic Claude Code (~/.claude.json)
  * windsurf        - Windsurf IDE (~/.codeium/windsurf/mcp_config.json)
  * vscode          - VS Code (.vscode/mcp.json)`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var clients []string
		if mcpInstallClients != "" {
			for _, c := range strings.Split(mcpInstallClients, ",") {
				if trimmed := strings.TrimSpace(c); trimmed != "" {
					clients = append(clients, trimmed)
				}
			}
		}
		return RunMCPInstall(cmd.OutOrStdout(), getWorkingDir(), "", clients, mcpInstallAll, mcpInstallNonInteractive)
	},
}

// RunMCPServer starts the stdio MCP server.
func RunMCPServer(in io.Reader, out io.Writer, cwd string) error {
	srv := mcp.NewServer("bender-mcp", Version, cwd)

	// Register default example tool
	srv.RegisterTool(mcp.Tool{
		Name:        "get_version",
		Description: "Returns the current Bender version and commit metadata",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
		},
	}, func(ctx context.Context, args map[string]interface{}) (mcp.CallToolResult, error) {
		return mcp.NewTextResult(fmt.Sprintf("Bender v%s (commit: %s, date: %s)", Version, Commit, BuildDate), false), nil
	})

	ctx := context.Background()
	return srv.Serve(ctx, in, out)
}

// RunMCPInstall configures client target configuration files.
func RunMCPInstall(out io.Writer, cwd string, homeDir string, clientIDs []string, all bool, nonInteractive bool) error {
	supported := mcp.GetSupportedClients(cwd, homeDir)
	var selectedIDs []string

	if all {
		for _, s := range supported {
			selectedIDs = append(selectedIDs, s.ID)
		}
	} else if len(clientIDs) > 0 {
		selectedIDs = clientIDs
	} else {
		// Non-interactive or default: install into detected clients or all if none detected
		for _, s := range supported {
			if s.Detected {
				selectedIDs = append(selectedIDs, s.ID)
			}
		}
		if len(selectedIDs) == 0 {
			for _, s := range supported {
				selectedIDs = append(selectedIDs, s.ID)
			}
		}
	}

	if len(selectedIDs) == 0 {
		fmt.Fprintln(out, "ℹ️ No AI clients selected. MCP configuration unchanged.")
		return nil
	}

	fmt.Fprintln(out, "🔌 Configuring Bender MCP Server...")

	opts := mcp.InstallerOptions{
		ServerName: "bender",
		Command:    "bender",
		Args:       []string{"mcp"},
		Cwd:        cwd,
		HomeDir:    homeDir,
		ClientIDs:  selectedIDs,
	}

	results, err := mcp.InstallClients(opts)
	if err != nil {
		return err
	}

	for _, res := range results {
		if res.Error != nil {
			fmt.Fprintf(out, "  [✗] Failed %s (%s): %v\n", res.DisplayName, res.ConfigPath, res.Error)
		} else if res.Created {
			fmt.Fprintf(out, "  [✓] Configured %s -> Created %s\n", res.DisplayName, res.ConfigPath)
		} else if res.Updated {
			fmt.Fprintf(out, "  [✓] Configured %s -> Updated %s\n", res.DisplayName, res.ConfigPath)
		}
	}

	fmt.Fprintln(out, "\n✨ MCP configuration completed!")
	return nil
}

func init() {
	MCPCmd.Flags().StringVarP(&mcpTransport, "transport", "t", "stdio", "Transport protocol to use (stdio)")

	MCPInstallCmd.Flags().StringVarP(&mcpInstallClients, "client", "c", "", "Comma-separated list of clients to configure (cursor, antigravity, claude-desktop, claude-code, windsurf, vscode)")
	MCPInstallCmd.Flags().BoolVarP(&mcpInstallAll, "all", "a", false, "Configure all supported AI clients")
	MCPInstallCmd.Flags().BoolVarP(&mcpInstallNonInteractive, "non-interactive", "y", false, "Run non-interactively configuring detected clients")

	MCPCmd.AddCommand(MCPInstallCmd)
	RootCmd.AddCommand(MCPCmd)
}

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/spf13/cobra"
	"github.com/thereisnotime/sdrangel-mcp/internal/sdrangel"
	"github.com/thereisnotime/sdrangel-mcp/internal/tools"
)

func newCLI(version string) *cobra.Command {
	var envFile string

	root := &cobra.Command{
		Use:     "sdrangel-mcp",
		Short:   "MCP server for SDRAngel",
		Version: version,
		PersistentPreRunE: func(_ *cobra.Command, _ []string) error {
			return loadDotenv(envFile)
		},
	}
	root.PersistentFlags().StringVar(&envFile, "env-file", "", "path to .env file (default: auto-discovered)")

	root.AddCommand(serveCmd(version))
	root.AddCommand(toolsCmd(version))
	root.AddCommand(callCmd(version))

	// Make serve the default when no subcommand is given.
	root.RunE = serveCmd(version).RunE

	return root
}

func buildClient() *sdrangel.Client {
	baseURL := envOr("SDRANGEL_BASE_URL", "http://localhost:8091")
	timeout := parseDuration(envOr("SDRANGEL_TIMEOUT", "10s"), 10*time.Second)
	username := os.Getenv("SDRANGEL_USERNAME")
	password := os.Getenv("SDRANGEL_PASSWORD")
	return sdrangel.New(sdrangel.Options{
		BaseURL:  baseURL,
		Timeout:  timeout,
		Username: username,
		Password: password,
	})
}

func serveCmd(version string) *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "Start the MCP server on stdio (default)",
		RunE: func(cmd *cobra.Command, _ []string) error {
			c := buildClient()
			srv := tools.New(c, tools.Options{Version: version})
			return srv.Run(cmd.Context(), &mcp.StdioTransport{})
		},
	}
}

func toolsCmd(version string) *cobra.Command {
	return &cobra.Command{
		Use:   "tools",
		Short: "List all available MCP tools by connecting in-memory and calling ListTools",
		RunE: func(cmd *cobra.Command, _ []string) error {
			c := buildClient()
			srv := tools.New(c, tools.Options{Version: version})

			ct, ss := mcp.NewInMemoryTransports()
			ctx := cmd.Context()

			srvDone := make(chan error, 1)
			go func() { srvDone <- srv.Run(ctx, ss) }()

			client := mcp.NewClient(&mcp.Implementation{Name: "sdrangel-mcp-cli", Version: version}, nil)
			session, err := client.Connect(ctx, ct, nil)
			if err != nil {
				return fmt.Errorf("connect: %w", err)
			}
			defer session.Close()

			result, err := session.ListTools(ctx, nil)
			if err != nil {
				return fmt.Errorf("list tools: %w", err)
			}
			for _, t := range result.Tools {
				fmt.Printf("%-40s %s\n", t.Name, t.Description)
			}
			return nil
		},
	}
}

func callCmd(version string) *cobra.Command {
	var jsonArg string

	cmd := &cobra.Command{
		Use:   "call <tool> [--json <args>]",
		Short: "Call a tool by name and print the result",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			toolName := args[0]

			var rawArgs any
			switch {
			case jsonArg == "-":
				b, err := io.ReadAll(os.Stdin)
				if err != nil {
					return fmt.Errorf("read stdin: %w", err)
				}
				if err := json.Unmarshal(b, &rawArgs); err != nil {
					return fmt.Errorf("parse json: %w", err)
				}
			case strings.HasPrefix(jsonArg, "@"):
				b, err := os.ReadFile(jsonArg[1:])
				if err != nil {
					return fmt.Errorf("read file: %w", err)
				}
				if err := json.Unmarshal(b, &rawArgs); err != nil {
					return fmt.Errorf("parse json: %w", err)
				}
			case jsonArg != "":
				if err := json.Unmarshal([]byte(jsonArg), &rawArgs); err != nil {
					return fmt.Errorf("parse json: %w", err)
				}
			default:
				rawArgs = map[string]any{}
			}

			c := buildClient()
			srv := tools.New(c, tools.Options{Version: version})

			ct, ss := mcp.NewInMemoryTransports()
			ctx := cmd.Context()

			go func() { _ = srv.Run(ctx, ss) }()

			client := mcp.NewClient(&mcp.Implementation{Name: "sdrangel-mcp-cli", Version: version}, nil)
			session, err := client.Connect(ctx, ct, nil)
			if err != nil {
				return fmt.Errorf("connect: %w", err)
			}
			defer session.Close()

			result, err := session.CallTool(ctx, &mcp.CallToolParams{
				Name:      toolName,
				Arguments: rawArgs,
			})
			if err != nil {
				return fmt.Errorf("call tool: %w", err)
			}

			for _, content := range result.Content {
				if tc, ok := content.(*mcp.TextContent); ok {
					fmt.Println(tc.Text)
				}
			}
			return nil
		},
	}
	cmd.Flags().StringVar(&jsonArg, "json", "", "JSON arguments (literal, - for stdin, @file for file)")
	return cmd
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func parseDuration(s string, fallback time.Duration) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		return fallback
	}
	return d
}

// suppress unused import warning for context package used via cmd.Context()
var _ = context.Background

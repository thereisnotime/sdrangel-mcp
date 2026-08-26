package tools

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/thereisnotime/sdrangel-mcp/internal/sdrangel"
)

func registerConfigurationTools(srv *mcp.Server, c *sdrangel.Client) {
	mcp.AddTool(srv, &mcp.Tool{
		Name:        "list_configurations",
		Description: "List all saved configurations grouped by group name. A configuration stores the complete SDRAngel workspace layout including all device sets, channels, and features.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, _ struct{}) (*mcp.CallToolResult, sdrangel.Configurations, error) {
		result, err := c.ListConfigurations(ctx)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "load_configuration",
		Description: "Load a saved configuration by groupName and name. This replaces the current workspace layout.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args sdrangel.ConfigurationKeys) (*mcp.CallToolResult, sdrangel.SuccessResponse, error) {
		result, err := c.LoadConfiguration(ctx, args)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "save_configuration",
		Description: "Save the current workspace layout into an existing configuration by groupName and name.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args sdrangel.ConfigurationKeys) (*mcp.CallToolResult, sdrangel.SuccessResponse, error) {
		result, err := c.SaveConfiguration(ctx, args)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "create_configuration",
		Description: "Create a new configuration from the current workspace layout with the given groupName and name.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args sdrangel.ConfigurationKeys) (*mcp.CallToolResult, sdrangel.SuccessResponse, error) {
		result, err := c.CreateConfiguration(ctx, args)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "delete_configuration",
		Description: "Delete a saved configuration by groupName and name.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args sdrangel.ConfigurationKeys) (*mcp.CallToolResult, sdrangel.ConfigurationKeys, error) {
		result, err := c.DeleteConfiguration(ctx, args)
		return nil, result, err
	})
}

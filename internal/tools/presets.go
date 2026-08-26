package tools

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/thereisnotime/sdrangel-mcp/internal/sdrangel"
)

func registerPresetTools(srv *mcp.Server, c *sdrangel.Client) {
	mcp.AddTool(srv, &mcp.Tool{
		Name:        "list_presets",
		Description: "List all saved presets grouped by group name. Each preset stores device and channel configuration for a specific frequency/mode.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, _ struct{}) (*mcp.CallToolResult, sdrangel.Presets, error) {
		result, err := c.ListPresets(ctx)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "load_preset",
		Description: "Load a preset into a device set. Specify deviceSetIndex and the preset identifier (groupName, name, type). type: R=Rx, T=Tx, M=MIMO.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args sdrangel.PresetTransfer) (*mcp.CallToolResult, sdrangel.SuccessResponse, error) {
		result, err := c.LoadPreset(ctx, args)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "save_preset",
		Description: "Save the current state of a device set into an existing preset. Specify deviceSetIndex and the preset identifier.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args sdrangel.PresetTransfer) (*mcp.CallToolResult, sdrangel.SuccessResponse, error) {
		result, err := c.SavePreset(ctx, args)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "create_preset",
		Description: "Create a new preset from the current state of a device set. Specify deviceSetIndex and the new preset identifier (groupName, name, type).",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args sdrangel.PresetTransfer) (*mcp.CallToolResult, sdrangel.SuccessResponse, error) {
		result, err := c.CreatePreset(ctx, args)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "delete_preset",
		Description: "Delete a saved preset by group name, name, and type.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args sdrangel.PresetKeys) (*mcp.CallToolResult, sdrangel.PresetIdentifier, error) {
		result, err := c.DeletePreset(ctx, args)
		return nil, result, err
	})
}

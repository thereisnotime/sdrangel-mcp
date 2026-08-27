package tools

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/thereisnotime/sdrangel-mcp/internal/sdrangel"
)

func registerFeaturePresetTools(srv *mcp.Server, c *sdrangel.Client) {
	mcp.AddTool(srv, &mcp.Tool{
		Name:        "list_feature_presets",
		Description: "List all saved feature presets grouped by group name. A feature preset stores the configuration of a set of features (as opposed to a device/channel preset). This is a separate catalog from the feature-set preset load/save/create tools.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, _ struct{}) (*mcp.CallToolResult, sdrangel.FeaturePresets, error) {
		result, err := c.ListFeaturePresets(ctx)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "delete_feature_preset",
		Description: "Delete a saved feature preset by group name and description.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args sdrangel.FeaturePresetIdentifier) (*mcp.CallToolResult, sdrangel.FeaturePresetIdentifier, error) {
		result, err := c.DeleteFeaturePreset(ctx, args)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "load_feature_set_preset",
		Description: "Load a preset into the current feature set (applies to the whole feature set, not a single device set). Specify the preset identifier (groupName, description).",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args sdrangel.FeaturePresetIdentifier) (*mcp.CallToolResult, sdrangel.FeaturePresetIdentifier, error) {
		result, err := c.LoadFeatureSetPreset(ctx, args)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "save_feature_set_preset",
		Description: "Save the current feature set state into an existing feature preset. Specify the preset identifier (groupName, description).",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args sdrangel.FeaturePresetIdentifier) (*mcp.CallToolResult, sdrangel.FeaturePresetIdentifier, error) {
		result, err := c.SaveFeatureSetPreset(ctx, args)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "create_feature_set_preset",
		Description: "Create a new feature preset from the current feature set state. Specify the new preset identifier (groupName, description).",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args sdrangel.FeaturePresetIdentifier) (*mcp.CallToolResult, sdrangel.FeaturePresetIdentifier, error) {
		result, err := c.CreateFeatureSetPreset(ctx, args)
		return nil, result, err
	})
}

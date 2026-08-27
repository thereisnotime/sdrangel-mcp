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
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args sdrangel.PresetTransfer) (*mcp.CallToolResult, sdrangel.PresetIdentifier, error) {
		result, err := c.LoadPreset(ctx, args)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "save_preset",
		Description: "Save the current state of a device set into an existing preset. Specify deviceSetIndex and the preset identifier.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args sdrangel.PresetTransfer) (*mcp.CallToolResult, sdrangel.PresetIdentifier, error) {
		result, err := c.SavePreset(ctx, args)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "create_preset",
		Description: "Create a new preset from the current state of a device set. Specify deviceSetIndex and the new preset identifier (groupName, name, type).",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args sdrangel.PresetTransfer) (*mcp.CallToolResult, sdrangel.PresetIdentifier, error) {
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

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "import_preset_from_file",
		Description: "Import a preset from a file path, creating a new preset. The file path is resolved on the server's filesystem (the machine running SDRAngel), not the MCP client's filesystem.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args sdrangel.FilePath) (*mcp.CallToolResult, sdrangel.PresetIdentifier, error) {
		result, err := c.ImportPresetFromFile(ctx, args)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "export_preset_to_file",
		Description: "Export an existing preset to a file path. The file path is resolved on the server's filesystem (the machine running SDRAngel), not the MCP client's filesystem.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args sdrangel.PresetExport) (*mcp.CallToolResult, sdrangel.PresetIdentifier, error) {
		result, err := c.ExportPresetToFile(ctx, args)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "import_preset_from_blob",
		Description: "Import a preset from a base64-encoded blob, creating a new preset.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args sdrangel.Base64Blob) (*mcp.CallToolResult, sdrangel.PresetIdentifier, error) {
		result, err := c.ImportPresetFromBlob(ctx, args)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "export_preset_to_blob",
		Description: "Export an existing preset to a base64-encoded blob.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args sdrangel.PresetIdentifier) (*mcp.CallToolResult, sdrangel.Base64Blob, error) {
		result, err := c.ExportPresetToBlob(ctx, args)
		return nil, result, err
	})
}

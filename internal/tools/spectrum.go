package tools

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/thereisnotime/sdrangel-mcp/internal/sdrangel"
)

func registerSpectrumTools(srv *mcp.Server, c *sdrangel.Client) {
	type DeviceSetIndexArgs struct {
		DeviceSetIndex int `json:"deviceSetIndex"`
	}

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "get_spectrum_settings",
		Description: "Get the spectrum display settings for a device set (FFT size, window, ref level, waterfall, averaging, etc.).",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args DeviceSetIndexArgs) (*mcp.CallToolResult, sdrangel.SpectrumSettings, error) {
		result, err := c.GetSpectrumSettings(ctx, args.DeviceSetIndex)
		return nil, result, err
	})

	type SetSpectrumSettingsArgs struct {
		DeviceSetIndex int                       `json:"deviceSetIndex"`
		Settings       sdrangel.SpectrumSettings `json:"settings"`
	}
	mcp.AddTool(srv, &mcp.Tool{
		Name:        "set_spectrum_settings",
		Description: "Set the spectrum display settings for a device set. Provide deviceSetIndex and a settings object with fields like fftSize, fftWindow, refLevel, powerRange, displayWaterfall, averagingMode, etc.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args SetSpectrumSettingsArgs) (*mcp.CallToolResult, sdrangel.SpectrumSettings, error) {
		result, err := c.SetSpectrumSettings(ctx, args.DeviceSetIndex, args.Settings)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "patch_spectrum_settings",
		Description: "Partially update the spectrum display settings for a device set. Only provided fields are changed. Use for incremental changes like ref level or FFT size.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args SetSpectrumSettingsArgs) (*mcp.CallToolResult, sdrangel.SpectrumSettings, error) {
		result, err := c.PatchSpectrumSettings(ctx, args.DeviceSetIndex, args.Settings)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "start_spectrum_server",
		Description: "Start the spectrum WebSocket server for a device set, which streams FFT data to remote viewers.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args DeviceSetIndexArgs) (*mcp.CallToolResult, sdrangel.SpectrumServer, error) {
		result, err := c.StartSpectrumServer(ctx, args.DeviceSetIndex)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "stop_spectrum_server",
		Description: "Stop the spectrum WebSocket server for a device set.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args DeviceSetIndexArgs) (*mcp.CallToolResult, sdrangel.SpectrumServer, error) {
		result, err := c.StopSpectrumServer(ctx, args.DeviceSetIndex)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "get_spectrum_server_status",
		Description: "Get the run status of the spectrum WebSocket server for a device set.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args DeviceSetIndexArgs) (*mcp.CallToolResult, sdrangel.SpectrumServer, error) {
		result, err := c.GetSpectrumServerStatus(ctx, args.DeviceSetIndex)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "get_spectrum_workspace",
		Description: "Get the workspace index the spectrum display widget is currently assigned to.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args DeviceSetIndexArgs) (*mcp.CallToolResult, sdrangel.WorkspaceInfo, error) {
		result, err := c.GetSpectrumWorkspace(ctx, args.DeviceSetIndex)
		return nil, result, err
	})

	type MoveSpectrumWorkspaceArgs struct {
		DeviceSetIndex int                    `json:"deviceSetIndex"`
		Move           sdrangel.WorkspaceInfo `json:"move"`
	}
	mcp.AddTool(srv, &mcp.Tool{
		Name:        "move_spectrum_to_workspace",
		Description: "Move the spectrum display widget to a different workspace by deviceSetIndex and move.index.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args MoveSpectrumWorkspaceArgs) (*mcp.CallToolResult, sdrangel.SuccessResponse, error) {
		result, err := c.MoveSpectrumToWorkspace(ctx, args.DeviceSetIndex, args.Move)
		return nil, result, err
	})
}

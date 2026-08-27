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
}

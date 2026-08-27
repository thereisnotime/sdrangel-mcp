package tools

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/thereisnotime/sdrangel-mcp/internal/sdrangel"
)

func registerDeviceTools(srv *mcp.Server, c *sdrangel.Client) {
	type SetDeviceArgs struct {
		DeviceSetIndex int                 `json:"deviceSetIndex"`
		Device         sdrangel.DeviceLink `json:"device"`
	}
	mcp.AddTool(srv, &mcp.Tool{
		Name:        "set_device",
		Description: "Load a specific SDR device into a device set. Provide deviceSetIndex and device with fields: deviceHwType (e.g. RTLSDRv2, HackRF), tx (0=Rx, 1=Tx), optionally serial or index.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args SetDeviceArgs) (*mcp.CallToolResult, sdrangel.DeviceDesc, error) {
		result, err := c.SetDevice(ctx, args.DeviceSetIndex, args.Device)
		return nil, result, err
	})

	type DeviceSetIndexArgs struct {
		DeviceSetIndex int `json:"deviceSetIndex"`
	}

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "get_device_settings",
		Description: "Get the current settings for the device in a device set. Returns device type and plugin-specific settings JSON.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args DeviceSetIndexArgs) (*mcp.CallToolResult, sdrangel.DeviceSettings, error) {
		result, err := c.GetDeviceSettings(ctx, args.DeviceSetIndex)
		return nil, result, err
	})

	type SetDeviceSettingsArgs struct {
		DeviceSetIndex int                     `json:"deviceSetIndex"`
		Settings       sdrangel.DeviceSettings `json:"settings"`
	}
	mcp.AddTool(srv, &mcp.Tool{
		Name:        "set_device_settings",
		Description: "Replace all settings for the device in a device set. The settings.settings field is plugin-specific JSON (e.g. RTLSDRSettings, HackRFInputSettings).",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args SetDeviceSettingsArgs) (*mcp.CallToolResult, sdrangel.DeviceSettings, error) {
		result, err := c.SetDeviceSettings(ctx, args.DeviceSetIndex, args.Settings)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "patch_device_settings",
		Description: "Partially update settings for the device in a device set. Only provided fields are changed. Use for incremental changes like frequency tuning.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args SetDeviceSettingsArgs) (*mcp.CallToolResult, sdrangel.DeviceSettings, error) {
		result, err := c.PatchDeviceSettings(ctx, args.DeviceSetIndex, args.Settings)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "start_device",
		Description: "Start the SDR device in a device set (begin acquisition or transmission).",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args DeviceSetIndexArgs) (*mcp.CallToolResult, sdrangel.DeviceState, error) {
		result, err := c.StartDevice(ctx, args.DeviceSetIndex)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "stop_device",
		Description: "Stop the SDR device in a device set.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args DeviceSetIndexArgs) (*mcp.CallToolResult, sdrangel.DeviceState, error) {
		result, err := c.StopDevice(ctx, args.DeviceSetIndex)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "get_device_run_status",
		Description: "Get the current run state (idle, running, error) of the device in a device set.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args DeviceSetIndexArgs) (*mcp.CallToolResult, sdrangel.DeviceState, error) {
		result, err := c.GetDeviceRunStatus(ctx, args.DeviceSetIndex)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "get_device_report",
		Description: "Get the runtime report for the device in a device set (hardware-specific runtime info like actual sample rate, frequency correction, etc.).",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args DeviceSetIndexArgs) (*mcp.CallToolResult, sdrangel.DeviceReport, error) {
		result, err := c.GetDeviceReport(ctx, args.DeviceSetIndex)
		return nil, result, err
	})
}

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
		Description: "Load a specific SDR device into a device set. Provide deviceSetIndex and device with fields: deviceHwType (e.g. RTLSDRv2, HackRF), direction (0=Rx, 1=Tx), and serial and/or index from list_device_plugins to disambiguate multiple identical devices.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args SetDeviceArgs) (*mcp.CallToolResult, sdrangel.DeviceDesc, error) {
		result, err := c.SetDevice(ctx, args.DeviceSetIndex, args.Device)
		return nil, result, err
	})

	type DeviceSetIndexArgs struct {
		DeviceSetIndex int `json:"deviceSetIndex"`
	}

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "get_device_settings",
		Description: "Get the current settings for the device in a device set. Returns deviceHwType, direction, settingsKey (the plugin-specific wire key, e.g. fileInputSettings) and settings (that plugin's settings JSON).",
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
		Description: "Replace all settings for the device in a device set. Call get_device_settings first to learn settings.settingsKey (SDRAngel wraps the plugin's settings object under a plugin-specific wire key, e.g. rtlSdrSettings, hackRFInputSettings — not a generic \"settings\" key) and echo it back along with settings.settings (the plugin-specific JSON).",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args SetDeviceSettingsArgs) (*mcp.CallToolResult, sdrangel.DeviceSettings, error) {
		result, err := c.SetDeviceSettings(ctx, args.DeviceSetIndex, args.Settings)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "patch_device_settings",
		Description: "Partially update settings for the device in a device set. Only provided fields are changed. Use for incremental changes like frequency tuning. Call get_device_settings first to learn settings.settingsKey and echo it back along with the changed fields in settings.settings.",
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
		Description: "Get the runtime report for the device in a device set (hardware-specific runtime info like actual sample rate, frequency correction, etc.). Returns reportKey (the plugin-specific wire key, e.g. fileInputReport) and report (that plugin's report JSON).",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args DeviceSetIndexArgs) (*mcp.CallToolResult, sdrangel.DeviceReport, error) {
		result, err := c.GetDeviceReport(ctx, args.DeviceSetIndex)
		return nil, result, err
	})

	type ExecuteDeviceActionsArgs struct {
		DeviceSetIndex int                    `json:"deviceSetIndex"`
		Actions        sdrangel.DeviceActions `json:"actions"`
	}
	mcp.AddTool(srv, &mcp.Tool{
		Name:        "execute_device_actions",
		Description: "Execute actions on the device in a device set (plugin-specific operations, e.g. a GPS-disciplined device's \"sync\" action). actions.actionsKey is the plugin-specific wire key (e.g. fileInputActions) wrapping actions.actions (the plugin-specific JSON) — not a generic \"actions\" key. Call get_device_settings first to learn the device's plugin naming convention.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args ExecuteDeviceActionsArgs) (*mcp.CallToolResult, sdrangel.SuccessResponse, error) {
		result, err := c.ExecuteDeviceActions(ctx, args.DeviceSetIndex, args.Actions)
		return nil, result, err
	})

	type SubdeviceIndexArgs struct {
		DeviceSetIndex int `json:"deviceSetIndex"`
		SubsystemIndex int `json:"subsystemIndex"`
	}

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "get_subdevice_run_status",
		Description: "Get the current run state (idle, running, error) of one subsystem (Rx or Tx side) of a multi-subsystem MIMO device (e.g. BladeRF2 MIMO, LimeSDR) in a device set. subsystemIndex: 0=Rx, 1=Tx.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args SubdeviceIndexArgs) (*mcp.CallToolResult, sdrangel.DeviceState, error) {
		result, err := c.GetSubdeviceRunStatus(ctx, args.DeviceSetIndex, args.SubsystemIndex)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "start_subdevice",
		Description: "Start one subsystem (Rx or Tx side) of a multi-subsystem MIMO device (e.g. BladeRF2 MIMO, LimeSDR) in a device set. subsystemIndex: 0=Rx, 1=Tx.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args SubdeviceIndexArgs) (*mcp.CallToolResult, sdrangel.DeviceState, error) {
		result, err := c.StartSubdevice(ctx, args.DeviceSetIndex, args.SubsystemIndex)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "stop_subdevice",
		Description: "Stop one subsystem (Rx or Tx side) of a multi-subsystem MIMO device (e.g. BladeRF2 MIMO, LimeSDR) in a device set. subsystemIndex: 0=Rx, 1=Tx.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args SubdeviceIndexArgs) (*mcp.CallToolResult, sdrangel.DeviceState, error) {
		result, err := c.StopSubdevice(ctx, args.DeviceSetIndex, args.SubsystemIndex)
		return nil, result, err
	})
}

package tools

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/thereisnotime/sdrangel-mcp/internal/sdrangel"
)

func registerChannelTools(srv *mcp.Server, c *sdrangel.Client) {
	type AddChannelArgs struct {
		DeviceSetIndex int                 `json:"deviceSetIndex"`
		Channel        sdrangel.ChannelAdd `json:"channel"`
	}
	mcp.AddTool(srv, &mcp.Tool{
		Name:        "add_channel",
		Description: "Add a new channel (demodulator/modulator) to a device set. Provide deviceSetIndex and channel.channelType (e.g. NFMDemod, WFMDemod, AMDemod, SSBDemod).",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args AddChannelArgs) (*mcp.CallToolResult, sdrangel.SuccessResponse, error) {
		result, err := c.AddChannel(ctx, args.DeviceSetIndex, args.Channel)
		return nil, result, err
	})

	type ChannelIndexArgs struct {
		DeviceSetIndex int `json:"deviceSetIndex"`
		ChannelIndex   int `json:"channelIndex"`
	}
	mcp.AddTool(srv, &mcp.Tool{
		Name:        "delete_channel",
		Description: "Delete a channel from a device set by deviceSetIndex and channelIndex.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args ChannelIndexArgs) (*mcp.CallToolResult, sdrangel.SuccessResponse, error) {
		result, err := c.DeleteChannel(ctx, args.DeviceSetIndex, args.ChannelIndex)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "get_channel_settings",
		Description: "Get the current settings for a channel. Returns channelType, direction, settingsKey (the plugin-specific wire key, e.g. NFMDemodSettings, WFMDemodSettings) and settings (that plugin's settings JSON).",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args ChannelIndexArgs) (*mcp.CallToolResult, sdrangel.ChannelSettings, error) {
		result, err := c.GetChannelSettings(ctx, args.DeviceSetIndex, args.ChannelIndex)
		return nil, result, err
	})

	type SetChannelSettingsArgs struct {
		DeviceSetIndex int                      `json:"deviceSetIndex"`
		ChannelIndex   int                      `json:"channelIndex"`
		Settings       sdrangel.ChannelSettings `json:"settings"`
	}
	mcp.AddTool(srv, &mcp.Tool{
		Name:        "set_channel_settings",
		Description: "Replace all settings for a channel. Call get_channel_settings first to learn settings.settingsKey (SDRAngel wraps the plugin's settings object under a plugin-specific wire key, e.g. NFMDemodSettings — not a generic \"settings\" key) and echo it back along with settings.settings (the plugin-specific JSON).",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args SetChannelSettingsArgs) (*mcp.CallToolResult, sdrangel.ChannelSettings, error) {
		result, err := c.SetChannelSettings(ctx, args.DeviceSetIndex, args.ChannelIndex, args.Settings)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "patch_channel_settings",
		Description: "Partially update settings for a channel. Only provided fields are changed. Use for incremental changes like frequency offset or squelch. Call get_channel_settings first to learn settings.settingsKey and echo it back along with the changed fields in settings.settings.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args SetChannelSettingsArgs) (*mcp.CallToolResult, sdrangel.ChannelSettings, error) {
		result, err := c.PatchChannelSettings(ctx, args.DeviceSetIndex, args.ChannelIndex, args.Settings)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "get_channel_report",
		Description: "Get the runtime report for a channel (channel-specific runtime info like signal level, lock status, decoded data). Returns reportKey (the plugin-specific wire key, e.g. NFMDemodReport) and report (that plugin's report JSON).",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args ChannelIndexArgs) (*mcp.CallToolResult, sdrangel.ChannelReport, error) {
		result, err := c.GetChannelReport(ctx, args.DeviceSetIndex, args.ChannelIndex)
		return nil, result, err
	})

	type ExecuteChannelActionsArgs struct {
		DeviceSetIndex int                     `json:"deviceSetIndex"`
		ChannelIndex   int                     `json:"channelIndex"`
		Actions        sdrangel.ChannelActions `json:"actions"`
	}
	mcp.AddTool(srv, &mcp.Tool{
		Name:        "execute_channel_actions",
		Description: "Execute actions on a channel (plugin-specific operations like start/stop recording, send a message, etc.). actions.actionsKey is the plugin-specific wire key (e.g. NFMDemodActions) wrapping actions.actions (the plugin-specific JSON) — not a generic \"actions\" key.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args ExecuteChannelActionsArgs) (*mcp.CallToolResult, sdrangel.SuccessResponse, error) {
		result, err := c.ExecuteChannelActions(ctx, args.DeviceSetIndex, args.ChannelIndex, args.Actions)
		return nil, result, err
	})

	type DeviceSetIndexArgs struct {
		DeviceSetIndex int `json:"deviceSetIndex"`
	}
	mcp.AddTool(srv, &mcp.Tool{
		Name:        "get_channels_report",
		Description: "Get the runtime report for all channels in a device set at once.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args DeviceSetIndexArgs) (*mcp.CallToolResult, sdrangel.ChannelsReport, error) {
		result, err := c.GetChannelsReport(ctx, args.DeviceSetIndex)
		return nil, result, err
	})
}

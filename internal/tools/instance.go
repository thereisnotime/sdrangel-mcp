package tools

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/thereisnotime/sdrangel-mcp/internal/sdrangel"
)

func registerInstanceTools(srv *mcp.Server, c *sdrangel.Client) {
	mcp.AddTool(srv, &mcp.Tool{
		Name:        "get_instance_summary",
		Description: "Get a summary of the SDRAngel instance: app name, version, Qt version, PID, DSP bits, OS, and architecture.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, _ struct{}) (*mcp.CallToolResult, sdrangel.InstanceSummary, error) {
		result, err := c.GetInstanceSummary(ctx)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "stop_instance",
		Description: "Stop the SDRAngel instance. This will shut down the application.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, _ struct{}) (*mcp.CallToolResult, sdrangel.SuccessResponse, error) {
		err := c.StopInstance(ctx)
		if err != nil {
			return nil, sdrangel.SuccessResponse{}, err
		}
		return nil, sdrangel.SuccessResponse{Message: "instance stop requested"}, nil
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "get_instance_config",
		Description: "Get the global SDRAngel instance configuration including preferences and commands.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, _ struct{}) (*mcp.CallToolResult, sdrangel.InstanceConfig, error) {
		result, err := c.GetInstanceConfig(ctx)
		return nil, result, err
	})

	type SetInstanceConfigArgs struct {
		Preferences map[string]interface{} `json:"preferences,omitempty"`
		Commands    map[string]interface{} `json:"commands,omitempty"`
	}
	mcp.AddTool(srv, &mcp.Tool{
		Name:        "set_instance_config",
		Description: "Set the global SDRAngel instance configuration. Accepts preferences and commands as JSON objects.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args SetInstanceConfigArgs) (*mcp.CallToolResult, sdrangel.InstanceConfig, error) {
		cfg := sdrangel.InstanceConfig{}
		if args.Preferences != nil {
			b, err := marshalJSON(args.Preferences)
			if err != nil {
				return nil, sdrangel.InstanceConfig{}, err
			}
			cfg.Preferences = b
		}
		if args.Commands != nil {
			b, err := marshalJSON(args.Commands)
			if err != nil {
				return nil, sdrangel.InstanceConfig{}, err
			}
			cfg.Commands = b
		}
		result, err := c.SetInstanceConfig(ctx, cfg)
		return nil, result, err
	})

	type PatchInstanceConfigArgs struct {
		Preferences map[string]interface{} `json:"preferences,omitempty"`
		Commands    map[string]interface{} `json:"commands,omitempty"`
	}
	mcp.AddTool(srv, &mcp.Tool{
		Name:        "patch_instance_config",
		Description: "Incrementally update the global SDRAngel instance configuration (upsert), unlike set_instance_config which fully replaces it. Presets and Commands if present are added; devices in the working preset are patched or added. Accepts preferences and commands as JSON objects.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args PatchInstanceConfigArgs) (*mcp.CallToolResult, sdrangel.InstanceConfig, error) {
		cfg := sdrangel.InstanceConfig{}
		if args.Preferences != nil {
			b, err := marshalJSON(args.Preferences)
			if err != nil {
				return nil, sdrangel.InstanceConfig{}, err
			}
			cfg.Preferences = b
		}
		if args.Commands != nil {
			b, err := marshalJSON(args.Commands)
			if err != nil {
				return nil, sdrangel.InstanceConfig{}, err
			}
			cfg.Commands = b
		}
		result, err := c.PatchInstanceConfig(ctx, cfg)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "list_device_plugins",
		Description: "List all available SDR device plugins (hardware types) supported by this SDRAngel build.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, _ struct{}) (*mcp.CallToolResult, sdrangel.AvailableDeviceList, error) {
		result, err := c.ListDevicePlugins(ctx)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "list_channel_plugins",
		Description: "List all available channel plugins (demodulators, modulators) supported by this SDRAngel build.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, _ struct{}) (*mcp.CallToolResult, sdrangel.AvailableChannelOrFeatureList, error) {
		result, err := c.ListChannelPlugins(ctx)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "list_feature_plugins",
		Description: "List all available feature plugins supported by this SDRAngel build.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, _ struct{}) (*mcp.CallToolResult, sdrangel.AvailableChannelOrFeatureList, error) {
		result, err := c.ListFeaturePlugins(ctx)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "get_location",
		Description: "Get the GPS location configured in SDRAngel (used for signal calculations like bearing, VOR, ADS-B).",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, _ struct{}) (*mcp.CallToolResult, sdrangel.Location, error) {
		result, err := c.GetLocation(ctx)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "set_location",
		Description: "Set the GPS location in SDRAngel. Latitude and longitude in decimal degrees, altitude in metres.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args sdrangel.Location) (*mcp.CallToolResult, sdrangel.Location, error) {
		result, err := c.SetLocation(ctx, args)
		return nil, result, err
	})
}

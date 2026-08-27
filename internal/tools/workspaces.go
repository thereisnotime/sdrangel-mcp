package tools

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/thereisnotime/sdrangel-mcp/internal/sdrangel"
)

func registerWorkspaceTools(srv *mcp.Server, c *sdrangel.Client) {
	mcp.AddTool(srv, &mcp.Tool{
		Name:        "add_workspace",
		Description: "Add a new empty workspace to the SDRAngel GUI. Returns the new workspace index.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, _ struct{}) (*mcp.CallToolResult, sdrangel.WorkspaceInfo, error) {
		result, err := c.AddWorkspace(ctx)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "delete_workspace",
		Description: "Delete the last empty workspace. The workspace must be empty before deletion.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, _ struct{}) (*mcp.CallToolResult, sdrangel.SuccessResponse, error) {
		result, err := c.DeleteWorkspace(ctx)
		return nil, result, err
	})

	type DeviceSetIndexArgs struct {
		DeviceSetIndex int `json:"deviceSetIndex"`
	}
	mcp.AddTool(srv, &mcp.Tool{
		Name:        "get_device_workspace",
		Description: "Get the workspace index a device set's main widget is currently assigned to.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args DeviceSetIndexArgs) (*mcp.CallToolResult, sdrangel.WorkspaceInfo, error) {
		result, err := c.GetDeviceWorkspace(ctx, args.DeviceSetIndex)
		return nil, result, err
	})

	type MoveDeviceArgs struct {
		DeviceSetIndex int                    `json:"deviceSetIndex"`
		Move           sdrangel.WorkspaceInfo `json:"move"`
	}
	mcp.AddTool(srv, &mcp.Tool{
		Name:        "move_device_to_workspace",
		Description: "Move a device set's main widget to a different workspace by deviceSetIndex and move.index.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args MoveDeviceArgs) (*mcp.CallToolResult, sdrangel.SuccessResponse, error) {
		result, err := c.MoveDeviceToWorkspace(ctx, args.DeviceSetIndex, args.Move)
		return nil, result, err
	})

	type ChannelIndexArgs struct {
		DeviceSetIndex int `json:"deviceSetIndex"`
		ChannelIndex   int `json:"channelIndex"`
	}
	mcp.AddTool(srv, &mcp.Tool{
		Name:        "get_channel_workspace",
		Description: "Get the workspace index a channel widget is currently assigned to.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args ChannelIndexArgs) (*mcp.CallToolResult, sdrangel.WorkspaceInfo, error) {
		result, err := c.GetChannelWorkspace(ctx, args.DeviceSetIndex, args.ChannelIndex)
		return nil, result, err
	})

	type MoveChannelArgs struct {
		DeviceSetIndex int                    `json:"deviceSetIndex"`
		ChannelIndex   int                    `json:"channelIndex"`
		Move           sdrangel.WorkspaceInfo `json:"move"`
	}
	mcp.AddTool(srv, &mcp.Tool{
		Name:        "move_channel_to_workspace",
		Description: "Move a channel widget to a different workspace by deviceSetIndex, channelIndex, and move.index.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args MoveChannelArgs) (*mcp.CallToolResult, sdrangel.SuccessResponse, error) {
		result, err := c.MoveChannelToWorkspace(ctx, args.DeviceSetIndex, args.ChannelIndex, args.Move)
		return nil, result, err
	})

	type FeatureIndexArgs struct {
		FeatureIndex int `json:"featureIndex"`
	}
	mcp.AddTool(srv, &mcp.Tool{
		Name:        "get_feature_workspace",
		Description: "Get the workspace index a feature widget is currently assigned to.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args FeatureIndexArgs) (*mcp.CallToolResult, sdrangel.WorkspaceInfo, error) {
		result, err := c.GetFeatureWorkspace(ctx, args.FeatureIndex)
		return nil, result, err
	})

	type MoveFeatureArgs struct {
		FeatureIndex int                    `json:"featureIndex"`
		Move         sdrangel.WorkspaceInfo `json:"move"`
	}
	mcp.AddTool(srv, &mcp.Tool{
		Name:        "move_feature_to_workspace",
		Description: "Move a feature widget to a different workspace by featureIndex and move.index.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args MoveFeatureArgs) (*mcp.CallToolResult, sdrangel.SuccessResponse, error) {
		result, err := c.MoveFeatureToWorkspace(ctx, args.FeatureIndex, args.Move)
		return nil, result, err
	})
}

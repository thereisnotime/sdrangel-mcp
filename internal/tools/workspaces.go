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

	type MoveDeviceArgs struct {
		DeviceSetIndex int                    `json:"deviceSetIndex"`
		Move           sdrangel.WorkspaceMove `json:"move"`
	}
	mcp.AddTool(srv, &mcp.Tool{
		Name:        "move_device_to_workspace",
		Description: "Move a device set's main widget to a different workspace by deviceSetIndex and move.workspaceIndex.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args MoveDeviceArgs) (*mcp.CallToolResult, sdrangel.SuccessResponse, error) {
		result, err := c.MoveDeviceToWorkspace(ctx, args.DeviceSetIndex, args.Move)
		return nil, result, err
	})

	type MoveChannelArgs struct {
		DeviceSetIndex int                    `json:"deviceSetIndex"`
		ChannelIndex   int                    `json:"channelIndex"`
		Move           sdrangel.WorkspaceMove `json:"move"`
	}
	mcp.AddTool(srv, &mcp.Tool{
		Name:        "move_channel_to_workspace",
		Description: "Move a channel widget to a different workspace by deviceSetIndex, channelIndex, and move.workspaceIndex.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args MoveChannelArgs) (*mcp.CallToolResult, sdrangel.SuccessResponse, error) {
		result, err := c.MoveChannelToWorkspace(ctx, args.DeviceSetIndex, args.ChannelIndex, args.Move)
		return nil, result, err
	})

	type MoveFeatureArgs struct {
		FeatureIndex int                    `json:"featureIndex"`
		Move         sdrangel.WorkspaceMove `json:"move"`
	}
	mcp.AddTool(srv, &mcp.Tool{
		Name:        "move_feature_to_workspace",
		Description: "Move a feature widget to a different workspace by featureIndex and move.workspaceIndex.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args MoveFeatureArgs) (*mcp.CallToolResult, sdrangel.SuccessResponse, error) {
		result, err := c.MoveFeatureToWorkspace(ctx, args.FeatureIndex, args.Move)
		return nil, result, err
	})
}

package tools

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/thereisnotime/sdrangel-mcp/internal/sdrangel"
)

func registerDeviceSetTools(srv *mcp.Server, c *sdrangel.Client) {
	mcp.AddTool(srv, &mcp.Tool{
		Name:        "list_device_sets",
		Description: "List all device sets (Rx/Tx/MIMO tabs) currently open in SDRAngel, including their sampling device and channel count.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, _ struct{}) (*mcp.CallToolResult, sdrangel.DeviceSets, error) {
		result, err := c.ListDeviceSets(ctx)
		return nil, result, err
	})

	type GetDeviceSetArgs struct {
		DeviceSetIndex int `json:"deviceSetIndex"`
	}
	mcp.AddTool(srv, &mcp.Tool{
		Name:        "get_device_set",
		Description: "Get detailed information about a specific device set by its index, including channels.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args GetDeviceSetArgs) (*mcp.CallToolResult, sdrangel.DeviceSetInfo, error) {
		result, err := c.GetDeviceSet(ctx, args.DeviceSetIndex)
		return nil, result, err
	})

	type AddDeviceSetArgs struct {
		Tx int `json:"tx"`
	}
	mcp.AddTool(srv, &mcp.Tool{
		Name:        "add_device_set",
		Description: "Add a new device set. Set tx=0 for an Rx device set, tx=1 for a Tx device set.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args AddDeviceSetArgs) (*mcp.CallToolResult, sdrangel.SuccessResponse, error) {
		result, err := c.AddDeviceSet(ctx, args.Tx)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "remove_device_set",
		Description: "Remove the last device set. The device set must be stopped before removal.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, _ struct{}) (*mcp.CallToolResult, sdrangel.SuccessResponse, error) {
		result, err := c.RemoveDeviceSet(ctx)
		return nil, result, err
	})
}

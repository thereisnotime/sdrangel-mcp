package tools

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/thereisnotime/sdrangel-mcp/internal/sdrangel"
)

func registerLoggingTools(srv *mcp.Server, c *sdrangel.Client) {
	mcp.AddTool(srv, &mcp.Tool{
		Name:        "get_logging",
		Description: "Get the current SDRAngel logging configuration: console level, file level, and log file path.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, _ struct{}) (*mcp.CallToolResult, sdrangel.LoggingInfo, error) {
		result, err := c.GetLogging(ctx)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "set_logging",
		Description: "Set the SDRAngel logging configuration. consoleLevel is required (debug, info, warning, error, fatal). Optionally set fileLevel, fileName, logToFile.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args sdrangel.LoggingInfo) (*mcp.CallToolResult, sdrangel.LoggingInfo, error) {
		result, err := c.SetLogging(ctx, args)
		return nil, result, err
	})
}

package tools

import (
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/thereisnotime/sdrangel-mcp/internal/sdrangel"
)

// New creates a configured MCP server with all SDRAngel tools registered.
func New(c *sdrangel.Client, opts Options) *mcp.Server {
	srv := mcp.NewServer(&mcp.Implementation{Name: "sdrangel-mcp", Version: opts.Version}, nil)
	registerInstanceTools(srv, c)
	registerAudioTools(srv, c)
	registerLoggingTools(srv, c)
	registerPresetTools(srv, c)
	registerConfigurationTools(srv, c)
	registerFeaturePresetTools(srv, c)
	registerDeviceSetTools(srv, c)
	registerSpectrumTools(srv, c)
	registerDeviceTools(srv, c)
	registerChannelTools(srv, c)
	registerFeatureTools(srv, c)
	registerWorkspaceTools(srv, c)
	return srv
}

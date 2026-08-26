package tools

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/thereisnotime/sdrangel-mcp/internal/sdrangel"
)

func registerAudioTools(srv *mcp.Server, c *sdrangel.Client) {
	mcp.AddTool(srv, &mcp.Tool{
		Name:        "get_audio_devices",
		Description: "List all available audio input and output devices known to SDRAngel.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, _ struct{}) (*mcp.CallToolResult, sdrangel.AudioDevices, error) {
		result, err := c.GetAudioDevices(ctx)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "set_audio_input_params",
		Description: "Set parameters for an audio input device by name. Configurable fields: name (required), sampleRate, volume.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args sdrangel.AudioInputDevice) (*mcp.CallToolResult, sdrangel.AudioInputDevice, error) {
		result, err := c.SetAudioInput(ctx, args)
		return nil, result, err
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "set_audio_output_params",
		Description: "Set parameters for an audio output device by name. Configurable fields: name (required), sampleRate, volume, copyToUDP, udpAddress, udpPort.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args sdrangel.AudioOutputDevice) (*mcp.CallToolResult, sdrangel.AudioOutputDevice, error) {
		result, err := c.SetAudioOutput(ctx, args)
		return nil, result, err
	})

	type DeviceNameArgs struct {
		Name string `json:"name"`
	}

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "reset_audio_input",
		Description: "Reset an audio input device to its default (unregistered) parameters.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args DeviceNameArgs) (*mcp.CallToolResult, sdrangel.SuccessResponse, error) {
		err := c.ResetAudioInput(ctx, args.Name)
		if err != nil {
			return nil, sdrangel.SuccessResponse{}, err
		}
		return nil, sdrangel.SuccessResponse{Message: "audio input reset"}, nil
	})

	mcp.AddTool(srv, &mcp.Tool{
		Name:        "reset_audio_output",
		Description: "Reset an audio output device to its default (unregistered) parameters.",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, args DeviceNameArgs) (*mcp.CallToolResult, sdrangel.SuccessResponse, error) {
		err := c.ResetAudioOutput(ctx, args.Name)
		if err != nil {
			return nil, sdrangel.SuccessResponse{}, err
		}
		return nil, sdrangel.SuccessResponse{Message: "audio output reset"}, nil
	})
}

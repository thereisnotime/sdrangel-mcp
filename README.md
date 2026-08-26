# sdrangel-mcp

[![CI](https://github.com/thereisnotime/sdrangel-mcp/actions/workflows/ci.yaml/badge.svg)](https://github.com/thereisnotime/sdrangel-mcp/actions/workflows/ci.yaml)
[![Release](https://github.com/thereisnotime/sdrangel-mcp/actions/workflows/release.yaml/badge.svg)](https://github.com/thereisnotime/sdrangel-mcp/actions/workflows/release.yaml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

An [MCP](https://modelcontextprotocol.io/) server for [SDRAngel](https://github.com/f4exb/sdrangel) — lets an LLM control SDRAngel's REST API to manage devices, channels, features, presets, and more.

## Features

- Full coverage of SDRAngel's REST API (port 8091 by default)
- 62 tools across 11 groups: instance, audio, logging, presets, configurations, device sets, spectrum, devices, channels, features, workspaces
- Runs over stdio — works with Claude Desktop, Claude Code, and any MCP-compatible client
- Single static binary, no runtime dependencies
- Config via environment variables or `.env` file

## Installation

### Download a release

Grab the latest binary from the [releases page](https://github.com/thereisnotime/sdrangel-mcp/releases).

### Build from source

```sh
go install github.com/thereisnotime/sdrangel-mcp/cmd/sdrangel-mcp@latest
```

Or clone and build:

```sh
git clone git@github.com:thereisnotime/sdrangel-mcp.git
cd sdrangel-mcp
just build
```

## Configuration

Copy `.env.example` to `.env` and edit as needed:

```sh
cp .env.example .env
```

| Variable | Default | Description |
|---|---|---|
| `SDRANGEL_BASE_URL` | `http://localhost:8091` | SDRAngel REST API URL |
| `SDRANGEL_TIMEOUT` | `10s` | HTTP request timeout |
| `SDRANGEL_LOG_LEVEL` | `info` | Log level |
| `SDRANGEL_USERNAME` | _(none)_ | HTTP basic auth username |
| `SDRANGEL_PASSWORD` | _(none)_ | HTTP basic auth password |

SDRAngel must be running with its web API enabled (enabled by default, port 8091).

## Usage

### Claude Desktop

Add to `~/Library/Application Support/Claude/claude_desktop_config.json` (macOS) or the equivalent on other platforms:

```json
{
  "mcpServers": {
    "sdrangel": {
      "command": "/path/to/sdrangel-mcp",
      "args": ["serve"],
      "env": {
        "SDRANGEL_BASE_URL": "http://localhost:8091"
      }
    }
  }
}
```

### Command line

```sh
# Start the MCP server on stdio
sdrangel-mcp serve

# List all available tools
sdrangel-mcp tools

# Call a specific tool
sdrangel-mcp call get_instance_summary
sdrangel-mcp call start_device --json '{"deviceSetIndex": 0}'
sdrangel-mcp call add_channel --json '{"deviceSetIndex": 0, "channel": {"channelType": "NFMDemod"}}'
```

### justfile

```sh
just build    # build binary
just test     # run tests
just lint     # run golangci-lint
just check    # tidy + fmt + vet + build + test
just run      # start server on stdio
```

## Tools

| Tool | Description |
|---|---|
| **Instance** | |
| `get_instance_summary` | Get SDRAngel instance info: version, PID, OS, DSP bits |
| `stop_instance` | Stop the SDRAngel application |
| `get_instance_config` | Get global preferences and commands |
| `set_instance_config` | Set global preferences and commands |
| `list_device_plugins` | List available SDR device plugins |
| `list_channel_plugins` | List available channel plugins (demodulators) |
| `list_feature_plugins` | List available feature plugins |
| `get_location` | Get configured GPS location |
| `set_location` | Set GPS location (lat, lon, altitude) |
| **Audio** | |
| `get_audio_devices` | List audio input and output devices |
| `set_audio_input_params` | Configure an audio input device |
| `set_audio_output_params` | Configure an audio output device |
| `reset_audio_input` | Reset audio input to defaults |
| `reset_audio_output` | Reset audio output to defaults |
| **Logging** | |
| `get_logging` | Get logging configuration |
| `set_logging` | Set logging configuration |
| **Presets** | |
| `list_presets` | List all saved presets |
| `load_preset` | Load a preset into a device set |
| `save_preset` | Save device set state into a preset |
| `create_preset` | Create a new preset from device set state |
| `delete_preset` | Delete a saved preset |
| **Configurations** | |
| `list_configurations` | List all saved workspace configurations |
| `load_configuration` | Load a saved configuration |
| `save_configuration` | Save current workspace into a configuration |
| `create_configuration` | Create a new workspace configuration |
| `delete_configuration` | Delete a saved configuration |
| **Device Sets** | |
| `list_device_sets` | List all open device sets |
| `get_device_set` | Get details for a specific device set |
| `add_device_set` | Add a new Rx or Tx device set |
| `remove_device_set` | Remove the last device set |
| **Spectrum** | |
| `get_spectrum_settings` | Get spectrum display settings |
| `set_spectrum_settings` | Set spectrum display settings |
| `start_spectrum_server` | Start the spectrum WebSocket server |
| `stop_spectrum_server` | Stop the spectrum WebSocket server |
| **Devices** | |
| `set_device` | Load an SDR device into a device set |
| `get_device_settings` | Get device settings |
| `set_device_settings` | Replace all device settings |
| `patch_device_settings` | Partially update device settings |
| `start_device` | Start the SDR device |
| `stop_device` | Stop the SDR device |
| `get_device_run_status` | Get device run state |
| `get_device_report` | Get device runtime report |
| **Channels** | |
| `add_channel` | Add a channel (demodulator/modulator) |
| `delete_channel` | Delete a channel |
| `get_channel_settings` | Get channel settings |
| `set_channel_settings` | Replace all channel settings |
| `patch_channel_settings` | Partially update channel settings |
| `get_channel_report` | Get channel runtime report |
| `execute_channel_actions` | Execute channel-specific actions |
| `get_channels_report` | Get report for all channels in a device set |
| **Features** | |
| `list_features` | List active feature instances |
| `add_feature` | Add a feature instance |
| `delete_feature` | Delete a feature instance |
| `get_feature_settings` | Get feature settings |
| `set_feature_settings` | Replace all feature settings |
| `patch_feature_settings` | Partially update feature settings |
| `start_feature` | Start a feature |
| `stop_feature` | Stop a feature |
| `get_feature_run_status` | Get feature run state |
| `get_feature_report` | Get feature runtime report |
| `execute_feature_actions` | Execute feature-specific actions |
| **Workspaces** | |
| `add_workspace` | Add a new workspace |
| `delete_workspace` | Delete the last empty workspace |
| `move_device_to_workspace` | Move a device set to a workspace |
| `move_channel_to_workspace` | Move a channel to a workspace |
| `move_feature_to_workspace` | Move a feature to a workspace |

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md).

## License

[MIT](LICENSE)

# sdrangel-mcp

[![CI](https://github.com/thereisnotime/sdrangel-mcp/actions/workflows/ci.yaml/badge.svg)](https://github.com/thereisnotime/sdrangel-mcp/actions/workflows/ci.yaml)
[![Release](https://github.com/thereisnotime/sdrangel-mcp/actions/workflows/release.yaml/badge.svg)](https://github.com/thereisnotime/sdrangel-mcp/actions/workflows/release.yaml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

An [MCP (Model Context Protocol)](https://modelcontextprotocol.io/) server for [SDRAngel](https://github.com/f4exb/sdrangel), the open-source software-defined radio (SDR) application — lets Claude or any other MCP-compatible LLM control SDRAngel's REST API to tune receivers, manage devices, channels, and features, save presets, and stream spectrum data.

## Documentation

- **[Getting Started](docs/getting-started.md)** — install, configure, and connect to Claude Desktop, Claude Code, or any MCP client
- **[Tools Reference](docs/tools-reference.md)** — every tool, grouped and described in detail
- **[Architecture](docs/architecture.md)** — how the server works, and the plugin-specific wire format SDRAngel's API uses
- **[Examples](docs/examples.md)** — worked prompts and tool-call sequences for common SDR tasks
- **[FAQ](docs/faq.md)** — what this is, what hardware it supports, and how it compares to the raw REST API

## Features

- Full coverage of SDRAngel's REST API (port 8091 by default) — every endpoint in SDRAngel's swagger spec is wrapped
- 93 tools across 12 groups: instance, audio, logging, presets, configurations, feature presets, device sets, spectrum, devices, channels, features, workspaces
- Runs over stdio — a plain MCP server with no client-specific code, so it works identically with Claude Desktop, Claude Code, Cursor, or any other MCP-compatible client (only the config file differs)
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

**Tested against:** SDRAngel 7.27.2. The REST API is fairly stable across versions, but if you hit a mismatch on an older or newer release, please [open an issue](https://github.com/thereisnotime/sdrangel-mcp/issues).

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

### Claude Code

Add to `.mcp.json` in your project root (or `~/.claude.json` for a user-wide server):

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

Or register it without hand-editing JSON:

```sh
claude mcp add sdrangel /path/to/sdrangel-mcp serve --env SDRANGEL_BASE_URL=http://localhost:8091
```

### Other MCP clients

Any client that speaks MCP over stdio works the same way: point it at the `sdrangel-mcp` binary with `serve` as the argument, and set `SDRANGEL_BASE_URL` (or the other environment variables below) if SDRAngel isn't on `localhost:8091`. Consult your client's docs for where it expects the server command to be registered.

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

Full descriptions and usage notes: [docs/tools-reference.md](docs/tools-reference.md).

| Tool | Description |
|---|---|
| **Instance** | |
| `get_instance_summary` | Get SDRAngel instance info: version, PID, OS, DSP bits |
| `stop_instance` | Stop the SDRAngel application |
| `get_instance_config` | Get global preferences and commands |
| `set_instance_config` | Replace global preferences and commands |
| `patch_instance_config` | Incrementally update (upsert) global preferences and commands |
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
| `cleanup_audio_input_devices` | Remove stale registered input device entries |
| `cleanup_audio_output_devices` | Remove stale registered output device entries |
| **Logging** | |
| `get_logging` | Get logging configuration |
| `set_logging` | Set logging configuration |
| **Presets** | |
| `list_presets` | List all saved presets |
| `load_preset` | Load a preset into a device set |
| `save_preset` | Save device set state into a preset |
| `create_preset` | Create a new preset from device set state |
| `delete_preset` | Delete a saved preset |
| `import_preset_from_file` | Import a preset from a server-side file |
| `export_preset_to_file` | Export a preset to a server-side file |
| `import_preset_from_blob` | Import a preset from a base64 blob |
| `export_preset_to_blob` | Export a preset to a base64 blob |
| **Configurations** | |
| `list_configurations` | List all saved workspace configurations |
| `load_configuration` | Load a saved configuration |
| `save_configuration` | Save current workspace into a configuration |
| `create_configuration` | Create a new workspace configuration |
| `delete_configuration` | Delete a saved configuration |
| `import_configuration_from_file` | Import a configuration from a server-side file |
| `export_configuration_to_file` | Export a configuration to a server-side file |
| `import_configuration_from_blob` | Import a configuration from a base64 blob |
| `export_configuration_to_blob` | Export a configuration to a base64 blob |
| **Feature Presets** | |
| `list_feature_presets` | List all saved feature-set presets |
| `delete_feature_preset` | Delete a saved feature-set preset |
| `load_feature_set_preset` | Load a preset into the current feature set |
| `save_feature_set_preset` | Save the feature set into an existing preset |
| `create_feature_set_preset` | Create a new preset from the feature set |
| **Device Sets** | |
| `list_device_sets` | List all open device sets |
| `get_device_set` | Get details for a specific device set |
| `add_device_set` | Add a new Rx or Tx device set |
| `remove_device_set` | Remove the last device set |
| **Spectrum** | |
| `get_spectrum_settings` | Get spectrum display settings |
| `set_spectrum_settings` | Replace all spectrum display settings |
| `patch_spectrum_settings` | Partially update spectrum display settings |
| `start_spectrum_server` | Start the spectrum WebSocket server |
| `stop_spectrum_server` | Stop the spectrum WebSocket server |
| `get_spectrum_server_status` | Get the spectrum WebSocket server run status |
| `get_spectrum_workspace` | Get which workspace the spectrum display is assigned to |
| `move_spectrum_to_workspace` | Move the spectrum display to a workspace |
| **Devices** | |
| `set_device` | Load an SDR device into a device set |
| `get_device_settings` | Get device settings |
| `set_device_settings` | Replace all device settings |
| `patch_device_settings` | Partially update device settings |
| `start_device` | Start the SDR device |
| `stop_device` | Stop the SDR device |
| `get_device_run_status` | Get device run state |
| `get_device_report` | Get device runtime report |
| `execute_device_actions` | Execute device-specific actions |
| `get_subdevice_run_status` | Get run state of one subsystem of a MIMO device |
| `start_subdevice` | Start one subsystem (Rx/Tx side) of a MIMO device |
| `stop_subdevice` | Stop one subsystem (Rx/Tx side) of a MIMO device |
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
| `get_device_workspace` | Get which workspace a device set is assigned to |
| `move_device_to_workspace` | Move a device set to a workspace |
| `get_channel_workspace` | Get which workspace a channel is assigned to |
| `move_channel_to_workspace` | Move a channel to a workspace |
| `get_feature_workspace` | Get which workspace a feature is assigned to |
| `move_feature_to_workspace` | Move a feature to a workspace |

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md).

## License

[MIT](LICENSE)

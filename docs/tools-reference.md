# Tools Reference

sdrangel-mcp exposes 93 MCP tools across 12 groups, covering every endpoint in SDRAngel's REST API. Run `sdrangel-mcp tools` against your own instance for a live, generated listing — this page groups the same tools with fuller descriptions and usage notes.

Several tools follow a **plugin-key wire format**: settings, reports, and actions for devices/channels/features are wrapped under a plugin-specific JSON key (e.g. `fileInputSettings`, `NFMDemodReport`) rather than a generic one. Where a tool's description says "call `get_*` first to learn the key," that's this pattern — see [Architecture: the plugin-key wire format](architecture.md#the-plugin-key-wire-format) for the full explanation before using `set_device_settings`, `patch_channel_settings`, `execute_feature_actions`, or similar tools directly.

## Instance

Global SDRAngel instance state: summary info, preferences, available plugins, and GPS location.

| Tool | Description |
|---|---|
| `get_instance_summary` | Get a summary of the SDRAngel instance: app name, version, Qt version, PID, DSP bits, OS, and architecture. |
| `stop_instance` | Stop the SDRAngel instance. This will shut down the application. |
| `get_instance_config` | Get the global SDRAngel instance configuration including preferences and commands. |
| `set_instance_config` | Set the global SDRAngel instance configuration. Accepts preferences and commands as JSON objects. Full replace. |
| `patch_instance_config` | Incrementally update (upsert) the global SDRAngel instance configuration, unlike `set_instance_config` which fully replaces it. Presets and commands, if present, are added; devices in the working preset are patched or added. |
| `list_device_plugins` | List all available SDR device plugins (hardware types) supported by this SDRAngel build. |
| `list_channel_plugins` | List all available channel plugins (demodulators, modulators) supported by this SDRAngel build. |
| `list_feature_plugins` | List all available feature plugins supported by this SDRAngel build. |
| `get_location` | Get the GPS location configured in SDRAngel (used for signal calculations like bearing, VOR, ADS-B). |
| `set_location` | Set the GPS location in SDRAngel. Latitude and longitude in decimal degrees, altitude in metres. |

## Audio

The audio input/output devices SDRAngel uses for demodulated audio playback and recording.

| Tool | Description |
|---|---|
| `get_audio_devices` | List all available audio input and output devices known to SDRAngel. |
| `set_audio_input_params` | Set parameters for an audio input device by name. Configurable fields: `name` (required), `sampleRate`, `volume`. |
| `set_audio_output_params` | Set parameters for an audio output device by name. Configurable fields: `name` (required), `sampleRate`, `volume`, `copyToUDP`, `udpAddress`, `udpPort`. |
| `reset_audio_input` | Reset an audio input device to its default (unregistered) parameters. |
| `reset_audio_output` | Reset an audio output device to its default (unregistered) parameters. |
| `cleanup_audio_input_devices` | Remove registered parameter entries for audio input devices that are no longer present on the system. |
| `cleanup_audio_output_devices` | Remove registered parameter entries for audio output devices that are no longer present on the system. |

## Logging

| Tool | Description |
|---|---|
| `get_logging` | Get the current SDRAngel logging configuration: console level, file level, and log file path. |
| `set_logging` | Set the SDRAngel logging configuration. `consoleLevel` is required (`debug`, `info`, `warning`, `error`, `fatal`). Optionally set `fileLevel`, `fileName`, `logToFile`. |

## Presets

Presets store a device set's device and channel configuration for a specific frequency/mode, so you can jump back to it later.

| Tool | Description |
|---|---|
| `list_presets` | List all saved presets grouped by group name. |
| `load_preset` | Load a preset into a device set. Specify `deviceSetIndex` and the preset identifier (`groupName`, `name`, `type`). `type`: `R`=Rx, `T`=Tx, `M`=MIMO. |
| `save_preset` | Save the current state of a device set into an existing preset. |
| `create_preset` | Create a new preset from the current state of a device set. |
| `delete_preset` | Delete a saved preset by group name, name, and type. |
| `import_preset_from_file` | Import a preset from a file path, creating a new preset. The file path is resolved on the server's filesystem (the machine running SDRAngel), not the MCP client's. |
| `export_preset_to_file` | Export an existing preset to a file path, on the server's filesystem. |
| `import_preset_from_blob` | Import a preset from a base64-encoded blob, creating a new preset. |
| `export_preset_to_blob` | Export an existing preset to a base64-encoded blob. |

## Configurations

A configuration stores the complete SDRAngel workspace layout — every device set, channel, and feature — as opposed to a single device set's preset.

| Tool | Description |
|---|---|
| `list_configurations` | List all saved configurations grouped by group name. |
| `load_configuration` | Load a saved configuration by `groupName` and `name`. Replaces the current workspace layout. |
| `save_configuration` | Save the current workspace layout into an existing configuration. |
| `create_configuration` | Create a new configuration from the current workspace layout. |
| `delete_configuration` | Delete a saved configuration. |
| `import_configuration_from_file` | Import a configuration from a server-side file path, creating a new configuration. |
| `export_configuration_to_file` | Export an existing configuration to a server-side file path. |
| `import_configuration_from_blob` | Import a configuration from a base64-encoded blob, creating a new configuration. |
| `export_configuration_to_blob` | Export an existing configuration to a base64-encoded blob. |

## Feature Presets

A separate preset catalog, parallel to device presets but for the feature set as a whole (features like Map, VOR Localizer, or ADS-B decoders that aren't tied to a single device set).

| Tool | Description |
|---|---|
| `list_feature_presets` | List all saved feature presets grouped by group name. |
| `delete_feature_preset` | Delete a saved feature preset by group name and description. |
| `load_feature_set_preset` | Load a preset into the current feature set (applies to the whole feature set, not a single device set). |
| `save_feature_set_preset` | Save the current feature set state into an existing feature preset. |
| `create_feature_set_preset` | Create a new feature preset from the current feature set state. |

## Device Sets

A device set is one Rx, Tx, or MIMO "tab" in SDRAngel — the container for one sampling device and its channels.

| Tool | Description |
|---|---|
| `list_device_sets` | List all device sets currently open in SDRAngel, including sampling device and channel count. |
| `get_device_set` | Get detailed information about a specific device set by index, including its channels. |
| `add_device_set` | Add a new device set. `tx=0` for Rx, `tx=1` for Tx. |
| `remove_device_set` | Remove the last device set. It must be stopped first. |

## Spectrum

The FFT/waterfall spectrum display for a device set, and its optional WebSocket streaming server for remote viewers.

| Tool | Description |
|---|---|
| `get_spectrum_settings` | Get the spectrum display settings for a device set (FFT size, window, ref level, waterfall, averaging, etc.). |
| `set_spectrum_settings` | Replace the spectrum display settings for a device set. |
| `patch_spectrum_settings` | Partially update the spectrum display settings. Only provided fields change — use for incremental adjustments like ref level or FFT size. |
| `start_spectrum_server` | Start the spectrum WebSocket server for a device set, streaming FFT data to remote viewers. |
| `stop_spectrum_server` | Stop the spectrum WebSocket server. |
| `get_spectrum_server_status` | Get the run status of the spectrum WebSocket server. |
| `get_spectrum_workspace` | Get which GUI workspace the spectrum display widget is assigned to. |
| `move_spectrum_to_workspace` | Move the spectrum display widget to a different workspace. |

## Devices

The physical (or file/test) sampling device loaded into a device set — tuning, sample rate, gain, and other hardware-specific settings.

| Tool | Description |
|---|---|
| `set_device` | Load a specific SDR device into a device set. Provide `deviceHwType` (e.g. `RTLSDRv2`, `HackRF`), `direction` (0=Rx, 1=Tx), and `serial`/`index` from `list_device_plugins` to disambiguate multiple identical devices. |
| `get_device_settings` | Get the current settings for the device in a device set — returns `deviceHwType`, `direction`, `settingsKey`, and `settings` (see the [plugin-key wire format](architecture.md#the-plugin-key-wire-format)). |
| `set_device_settings` | Replace all device settings. Call `get_device_settings` first to learn `settingsKey`. |
| `patch_device_settings` | Partially update device settings — use for incremental changes like frequency tuning. |
| `start_device` | Start the SDR device (begin acquisition or transmission). |
| `stop_device` | Stop the SDR device. |
| `get_device_run_status` | Get the current run state (idle, running, error) of the device. |
| `get_device_report` | Get the device's runtime report — hardware-specific info like actual sample rate or frequency correction. |
| `execute_device_actions` | Execute plugin-specific actions on the device (e.g. a GPS-disciplined device's "sync" action). |
| `get_subdevice_run_status` | Get run state of one subsystem (Rx or Tx side) of a multi-subsystem MIMO device (e.g. BladeRF2 MIMO, LimeSDR). `subsystemIndex`: 0=Rx, 1=Tx. |
| `start_subdevice` | Start one subsystem of a MIMO device. |
| `stop_subdevice` | Stop one subsystem of a MIMO device. |

## Channels

Demodulators and modulators (NFM, WFM, AM, SSB, and dozens of protocol-specific decoders) attached to a device set.

| Tool | Description |
|---|---|
| `add_channel` | Add a new channel to a device set. Provide `channelType` (e.g. `NFMDemod`, `WFMDemod`, `AMDemod`, `SSBDemod`). |
| `delete_channel` | Delete a channel from a device set. |
| `get_channel_settings` | Get the current settings for a channel — returns `channelType`, `direction`, `settingsKey`, `settings`. |
| `set_channel_settings` | Replace all channel settings. Call `get_channel_settings` first to learn `settingsKey`. |
| `patch_channel_settings` | Partially update channel settings — use for incremental changes like frequency offset or squelch. |
| `get_channel_report` | Get a channel's runtime report — signal level, lock status, decoded data. |
| `execute_channel_actions` | Execute plugin-specific actions on a channel (start/stop recording, send a message, etc.). |
| `get_channels_report` | Get the runtime report for all channels in a device set at once. |

## Features

Standalone functionality not tied to a single device set — maps, VOR/ADS-B decoders that aggregate multiple channels, GPS-disciplined clocks, and more.

| Tool | Description |
|---|---|
| `list_features` | List active feature instances (e.g. Map, ADS-B Demodulator, RDS, VOR Localizer). |
| `add_feature` | Add a new feature instance by `featureType` (e.g. `Map`, `VORLocalizer`, `AISDemod`). Use `list_feature_plugins` to see available types. |
| `delete_feature` | Delete a feature instance. |
| `get_feature_settings` | Get a feature's current settings — returns `featureType`, `settingsKey`, `settings`. |
| `set_feature_settings` | Replace all feature settings. Call `get_feature_settings` first to learn `settingsKey`. |
| `patch_feature_settings` | Partially update feature settings. |
| `start_feature` | Start a feature instance (begin its processing loop). |
| `stop_feature` | Stop a feature instance. |
| `get_feature_run_status` | Get the current run state of a feature instance. |
| `get_feature_report` | Get a feature's runtime report. |
| `execute_feature_actions` | Execute plugin-specific actions on a feature instance. |

## Workspaces

SDRAngel's GUI workspaces — which "screen" a device set, channel, feature, or spectrum widget is docked into.

| Tool | Description |
|---|---|
| `add_workspace` | Add a new empty workspace. Returns its index. |
| `delete_workspace` | Delete the last empty workspace (must be empty first). |
| `get_device_workspace` | Get which workspace a device set's main widget is assigned to. |
| `move_device_to_workspace` | Move a device set's main widget to a different workspace. |
| `get_channel_workspace` | Get which workspace a channel widget is assigned to. |
| `move_channel_to_workspace` | Move a channel widget to a different workspace. |
| `get_feature_workspace` | Get which workspace a feature widget is assigned to. |
| `move_feature_to_workspace` | Move a feature widget to a different workspace. |

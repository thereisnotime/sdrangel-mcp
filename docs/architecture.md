# Architecture

How sdrangel-mcp is built, and the SDRAngel REST API quirks it exists to paper over.

## Project layout

```
cmd/sdrangel-mcp/   CLI entry point (serve / tools / call subcommands)
internal/sdrangel/  HTTP client + types for SDRAngel's REST API
internal/tools/     MCP tool registrations, one file per resource group
internal/version/   Version metadata injected at build time
```

`internal/sdrangel` knows nothing about MCP — it's a plain Go HTTP client for SDRAngel's REST API (`GET`/`POST`/`PUT`/`PATCH`/`DELETE` wrappers around `net/http`, with typed request/response structs). `internal/tools` is the thin MCP layer on top: one `mcp.AddTool` call per operation, each just unmarshaling arguments and calling the matching client method. `cmd/sdrangel-mcp` wires a `sdrangel.Client` and an `mcp.Server` together and runs it over stdio (or exposes `tools`/`call` subcommands for scripting and debugging without an MCP client).

Adding a new tool means: add a client method in `internal/sdrangel/<group>.go`, add any new types to `internal/sdrangel/types.go`, register the tool in `internal/tools/<group>.go`. See [CONTRIBUTING.md](../CONTRIBUTING.md) for the full checklist.

## Why a dedicated server instead of a generic REST-to-MCP bridge

SDRAngel's REST API is large (50+ endpoints), stateful (device sets, channels, and features are addressed by index, and those indices shift as you add/remove them), and has response shapes that vary per hardware/plugin. A generic OpenAPI-to-MCP bridge would either expose the raw REST shape (leaving the LLM to work around plugin-specific quirks itself, badly) or need per-endpoint configuration anyway. A dedicated server lets each tool carry the context an LLM actually needs — descriptions that explain *when* to call it and *what the plugin-specific fields mean*, not just a mechanical parameter list.

## The plugin-key wire format

This is the single most important thing to understand if you're calling `set_*`/`patch_*`/`execute_*` tools directly (an LLM that reads the tool descriptions will generally get this right on its own).

Device, channel, and feature **settings**, **reports**, and **actions** are all plugin-specific — an `RTLSDR` device's settings look nothing like a `HackRF`'s, and an `NFMDemod` channel's settings look nothing like an `SSBDemod`'s. SDRAngel's REST API handles this by wrapping the plugin's payload under a **plugin-specific JSON key**, not a generic one:

```jsonc
// GET /sdrangel/deviceset/0/device/settings for a FileInput device
{
  "deviceHwType": "FileInput",
  "direction": 0,
  "fileInputSettings": { "accelerationFactor": 1, "loop": 1, "...": "..." }
}

// GET /sdrangel/deviceset/0/channel/0/settings for an NFMDemod channel
{
  "channelType": "NFMDemod",
  "direction": 0,
  "NFMDemodSettings": { "afBandwidth": 3000, "rfBandwidth": 12500, "...": "..." }
}
```

Sending a request with a generic `"settings"` key instead of the real one (`fileInputSettings`, `NFMDemodSettings`, ...) gets rejected outright by SDRAngel with `{"message":"Invalid JSON request"}`. The key name isn't derivable from the plugin's type string by a fixed rule either — device keys are lowerCamelCase of the plugin name (`fileInputSettings`), while channel/feature keys keep the plugin's PascalCase name (`NFMDemodSettings`, `MapSettings`) — so it has to be discovered, not guessed.

sdrangel-mcp's `DeviceSettings`/`ChannelSettings`/`FeatureSettings` types (and their `*Report`/`*Actions` counterparts, plus the same pattern for `DeviceActions`) model this directly with a `SettingsKey`/`ReportKey`/`ActionsKey` field and a custom JSON marshaler/unmarshaler (`internal/sdrangel/types.go`):

- **Reading** (`get_device_settings`, `get_channel_report`, ...) parses whatever the real wire key happens to be into `SettingsKey`/`ReportKey`, so you always see the actual key SDRAngel used.
- **Writing** (`set_device_settings`, `patch_channel_settings`, `execute_feature_actions`, ...) re-wraps your payload under that same key when it builds the HTTP request.

The practical rule for any LLM or script using these tools: **call the matching `get_*` tool first**, read back `settingsKey`/`reportKey`, and echo that exact string back on the `set_*`/`patch_*`/`execute_*` call alongside your changes. Every tool description in the [Tools Reference](tools-reference.md) restates this where it applies.

## Other SDRAngel REST API quirks worth knowing

- **Rx/Tx is `direction`, not `tx`, on the wire** for device settings/reports (`0` = Rx, `1` = Tx) — this project's types use `Direction` consistently to match.
- **List endpoints don't always use the field name you'd guess.** `GET /sdrangel/devicesets` returns its array under `deviceSets`, for instance, not `devicesetList`. Every type in `internal/sdrangel/types.go` was checked against a live instance and its response bodies rather than assumed from naming conventions.
- **Some `PATCH`/`PUT`/`POST` endpoints echo back a resource identifier, not a generic success message** — e.g. saving a preset returns the `PresetIdentifier` it just wrote, not `{"message": "ok"}`. sdrangel-mcp's tools return the identifier so a caller (LLM or script) can chain the result into a later call without a second lookup.
- **Not every plugin implements every operation.** A device/channel/feature that doesn't support a report or action returns HTTP 501 from SDRAngel — that's expected behavior surfaced as-is, not an error in this server.

## Testing approach

`internal/sdrangel/types_test.go` and `devices_test.go` pin the plugin-key marshal/unmarshal logic against JSON fixtures captured from a real SDRAngel instance (FileInput, AudioInput, bladeRF2, NFMDemod, Map, and others), so a future change to the wire-format handling can't silently regress back to a generic key SDRAngel would reject. `just check` runs `go vet`, `gofmt -l`, a full build, and `go test -race -shuffle=on`.

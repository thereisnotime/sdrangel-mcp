# sdrangel-mcp

## Build and Test

```sh
just build    # compile binary
just test     # run tests with -race -shuffle=on
just check    # full gate: tidy + fmt-check + vet + build + test
just lint     # golangci-lint
just sast     # gosec SARIF output
just vulncheck # govulncheck
```

## Project Layout

```
cmd/sdrangel-mcp/   CLI entry point (main.go, cli.go, dotenv.go)
internal/sdrangel/  HTTP client + types for SDRAngel REST API
internal/tools/     MCP tool registrations (one file per group)
internal/version/   Version vars injected by GoReleaser ldflags
```

## Non-Obvious Patterns

- `patch_*` tools call PATCH (incremental update); `set_*` tools call PUT (full replace).
- Device, channel, and feature settings/reports/actions are plugin-specific, and SDRAngel wraps the payload under a plugin-specific wire key (e.g. `fileInputSettings`, `NFMDemodSettings`, `MapReport`) instead of a generic `settings`/`report`/`actions` key — sending a generic key gets `"Invalid JSON request"` from the live API. `DeviceSettings`/`ChannelSettings`/`FeatureSettings` (and their `*Report`/`*Actions` counterparts) in `internal/sdrangel/types.go` capture this via `SettingsKey`/`ReportKey`/`ActionsKey` + custom `MarshalJSON`/`UnmarshalJSON` — always call the corresponding `get_*` tool first to learn the key, then echo it back on `set_*`/`patch_*`/`execute_*` calls.
- Device settings/reports use `direction` (not `tx`) as the Rx/Tx field name on the wire — matches `DeviceDesc`/`ChannelSettings`. `DeviceLink.Direction` (used by `set_device`) follows the same convention.
- `json.RawMessage` fields in response types (DeviceSettings, ChannelSettings, etc.) will appear as base64-encoded strings in some MCP clients; pass them as objects in requests.
- The SDRAngel REST API is on port 8091 by default. Enable it in SDRAngel preferences if not running (it's on by default in most builds).
- `mcp.AddTool` (generic) handles arg unmarshaling and output serialization automatically — do not call `json.Marshal` on the output.
- `patchReq` is used instead of `patch` to avoid shadowing the built-in. The name is internal to the sdrangel package.

# Getting Started

This walks through installing sdrangel-mcp, enabling SDRAngel's REST API, and connecting the server to an MCP client. See the [Tools Reference](tools-reference.md) for what you can do once it's running, or [Examples](examples.md) for worked prompts.

## 1. Prerequisites

- **[SDRAngel](https://github.com/f4exb/sdrangel)** installed and able to run — with real SDR hardware (RTL-SDR, HackRF, LimeSDR, BladeRF, Airspy, PlutoSDR, and dozens more) or with a file/test source if you just want to try it out. This project is tested against **SDRAngel 7.27.2**; the REST API is fairly stable across versions, but [open an issue](https://github.com/thereisnotime/sdrangel-mcp/issues) if you hit a mismatch on another release.
- SDRAngel's **web/REST API** enabled. It's on by default in most builds, listening on port `8091`. If it isn't reachable, open SDRAngel's Preferences and confirm the REST API server is enabled and bound to a reachable address.
- An **MCP-compatible client**: Claude Desktop, Claude Code, Cursor, or any other application that speaks MCP over stdio.

Verify the REST API is up before going further:

```sh
curl http://localhost:8091/sdrangel
```

A JSON response with `"appname": "SDRangel"` means you're good to go. A connection error means the REST API isn't enabled or isn't on the port you expect.

## 2. Install sdrangel-mcp

**Download a release binary** from the [releases page](https://github.com/thereisnotime/sdrangel-mcp/releases) — no dependencies, just an executable.

**Or build from source** (requires Go 1.26+):

```sh
git clone git@github.com:thereisnotime/sdrangel-mcp.git
cd sdrangel-mcp
just build   # or: go build -o sdrangel-mcp ./cmd/sdrangel-mcp
```

**Or install via `go install`:**

```sh
go install github.com/thereisnotime/sdrangel-mcp/cmd/sdrangel-mcp@latest
```

## 3. Configure

sdrangel-mcp reads configuration from environment variables, optionally loaded from a `.env` file next to the binary (or pointed to with `--env-file`):

| Variable | Default | Description |
|---|---|---|
| `SDRANGEL_BASE_URL` | `http://localhost:8091` | SDRAngel REST API URL |
| `SDRANGEL_TIMEOUT` | `10s` | HTTP request timeout |
| `SDRANGEL_LOG_LEVEL` | `info` | Log level |
| `SDRANGEL_USERNAME` | _(none)_ | HTTP basic auth username, if SDRAngel's REST API is behind auth |
| `SDRANGEL_PASSWORD` | _(none)_ | HTTP basic auth password |

```sh
cp .env.example .env
# edit .env if SDRAngel isn't on localhost:8091
```

If SDRAngel runs on a different machine (e.g. a headless SDR server on your network), set `SDRANGEL_BASE_URL` to that machine's address — sdrangel-mcp itself can run anywhere that can reach it, including on the same machine as your MCP client.

## 4. Connect to an MCP client

### Claude Desktop

Add to `~/Library/Application Support/Claude/claude_desktop_config.json` (macOS) or the equivalent config path on your platform:

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

Add to `.mcp.json` in your project, or register it directly:

```sh
claude mcp add sdrangel /path/to/sdrangel-mcp serve --env SDRANGEL_BASE_URL=http://localhost:8091
```

### Any other MCP client

sdrangel-mcp is a plain stdio MCP server with nothing Claude-specific in it. Point your client at the binary with `serve` as the argument, the same way you would for any other MCP server — only the client's own config file format differs.

## 5. Verify it works

Without touching an MCP client, you can list and call tools directly from the command line:

```sh
# List every registered tool
sdrangel-mcp tools

# Call a tool directly
sdrangel-mcp call get_instance_summary
sdrangel-mcp call list_device_sets
```

`get_instance_summary` returning SDRAngel's version, PID, and OS confirms the whole chain — binary, REST API connection, and tool registration — is working. From your MCP client, ask something like "what SDRAngel devices are available?" and the model should call `list_device_plugins` on its own.

## Troubleshooting

- **Connection refused / timeout calling any tool** — SDRAngel isn't running, or its REST API isn't enabled, or `SDRANGEL_BASE_URL` doesn't match where it's actually listening.
- **A `set_*`/`patch_*` tool call fails with "Invalid JSON request"** — you likely sent a generic `settings`/`report`/`actions` object instead of using the plugin-specific wire key. See [Architecture: the plugin-key wire format](architecture.md#the-plugin-key-wire-format) — always call the matching `get_*` tool first to learn the key.
- **A tool call errors with an HTTP 501** — that particular device/channel/feature plugin doesn't implement that operation (not every plugin exposes a report, for instance). This comes straight from SDRAngel, not a bug in the server.

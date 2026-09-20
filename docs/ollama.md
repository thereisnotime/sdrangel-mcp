# Running sdrangel-mcp with a local Ollama model

This is an **optional** setup. sdrangel-mcp is a plain stdio MCP server and nothing here is required for normal use with Claude Desktop, Claude Code, or any other MCP client. See [Getting Started](getting-started.md) for the default path.

If you'd rather keep everything local, with no cloud model in the loop, you can drive sdrangel-mcp from [Ollama](https://ollama.com/) using [ollama-mcp-bridge](https://github.com/jonigl/ollama-mcp-bridge). The bridge exposes an Ollama-compatible `/api/chat` endpoint that has your MCP tools attached, so existing Ollama clients work unchanged by pointing at the bridge instead of at Ollama directly.

## Prerequisites

- SDRAngel running with its REST API reachable (default `http://localhost:8091`), as in [Getting Started](getting-started.md#1-prerequisites).
- The `sdrangel-mcp` binary, from the [releases page](https://github.com/thereisnotime/sdrangel-mcp/releases) or built from source.
- [Ollama](https://ollama.com/download) with a tool-calling model pulled, for example `ollama pull qwen3:8b`.
- `ollama-mcp-bridge`, installed with `uvx ollama-mcp-bridge` or `pip install ollama-mcp-bridge`.

## Configure the bridge

Create `mcp-config.json`:

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

This is the same server block every other MCP client uses. Because sdrangel-mcp speaks stdio natively, the bridge launches it directly, with no SSE-to-stdio proxy in between.

## Run it

```sh
ollama serve                                  # if not already running
ollama-mcp-bridge --config mcp-config.json    # listens on :8000 by default
```

Then chat against the bridge exactly as you would against Ollama:

```sh
curl -N -X POST http://localhost:8000/api/chat \
  -H "Content-Type: application/json" \
  -d '{
    "model": "qwen3:8b",
    "messages": [
      {"role": "user", "content": "What SDRAngel version is running, and what device sets are open?"}
    ],
    "stream": true
  }'
```

The bridge forwards the tool list to the model, executes any tool calls against sdrangel-mcp, and folds the results back into the reply.

## Limiting which tools the model sees

sdrangel-mcp registers 93 tools. That is a lot of schema for a small local model to reason over, and small models are noticeably more reliable when given a narrow set. The bridge's `toolFilter` handles both that and basic safety.

**Include mode** keeps a working subset, which is the practical fix for a 1.7b or 4b model that gets lost in the full surface:

```json
{
  "mcpServers": {
    "sdrangel": {
      "command": "/path/to/sdrangel-mcp",
      "args": ["serve"],
      "toolFilter": {
        "mode": "include",
        "tools": [
          "get_instance_summary",
          "list_device_sets",
          "get_device_settings",
          "patch_device_settings",
          "get_channel_settings",
          "patch_channel_settings",
          "start_device",
          "stop_device",
          "list_presets",
          "load_preset"
        ]
      }
    }
  }
}
```

**Exclude mode** is the blunt safety instrument. Unlike Claude Desktop and Claude Code, which prompt before each tool call, a bridge-driven loop executes tool calls unattended, so anything you are not willing to have fire without review should be filtered out here:

```json
"toolFilter": {
  "mode": "exclude",
  "tools": ["stop_instance", "delete_preset", "delete_configuration", "delete_channel", "delete_feature"]
}
```

Tool names must match exactly and are case-sensitive.

## Caveats

- **No confirmation step.** The protocol applies settings immediately, and the bridge has no approval prompt. `toolFilter` in exclude mode is the only gate in this setup. Use it.
- **Transmitting.** If you add Tx device sets (`direction: 1`), an unattended model can key a transmitter. Be deliberate about what you connect and about the licence and regulatory constraints that apply where you are.
- **Model capability.** The plugin-key wire format means most write flows are two steps: call the `get_*` tool to learn the settings key, then echo that key back on the `set_*`/`patch_*` call. See [Architecture](architecture.md#the-plugin-key-wire-format). Models in the 1-2b range often skip the read step and send a generic `settings` key, which SDRAngel rejects with `"Invalid JSON request"`. An 8b or larger tool-calling model handles this considerably better.
- **Latency.** Local inference adds a noticeable delay before tokens appear, particularly when the model is summarising a long tool result such as a device or channel report.

## Related

- [Getting Started](getting-started.md) for the standard MCP client setup.
- [Tools Reference](tools-reference.md) for the full tool list, useful when building a `toolFilter` allow-list.
- [SDRangel driven by Ollama LLM using MCP](https://www.pg540.org/wiki/index.php/SDRAngel_driven_by_Ollama_LLM_using_MCP), an independent Node-RED based experiment by JKO (PG540) that inspired this page.

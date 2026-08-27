# sdrangel-mcp documentation

**sdrangel-mcp** is a [Model Context Protocol (MCP)](https://modelcontextprotocol.io/) server for [SDRAngel](https://github.com/f4exb/sdrangel), the open-source software-defined radio (SDR) application. It exposes SDRAngel's full REST API as MCP tools, so an LLM — Claude, or any other MCP-compatible model — can tune receivers, load demodulators, manage presets, and drive an SDR workflow through natural-language instructions instead of clicking through a GUI.

This directory has the documentation that doesn't fit in the top-level [README](../README.md). Start there for installation; come here for depth.

## Contents

- **[Getting Started](getting-started.md)** — installing sdrangel-mcp, enabling SDRAngel's REST API, connecting it to Claude Desktop, Claude Code, or another MCP client, and verifying the connection.
- **[Tools Reference](tools-reference.md)** — every one of the 93 MCP tools, grouped and described, with the arguments each one expects.
- **[Architecture](architecture.md)** — how the server is built, how it talks to SDRAngel's REST API, and the plugin-specific wire format you need to understand before writing settings/report/action calls.
- **[Examples](examples.md)** — worked prompts and tool-call sequences for common SDR tasks: tuning a receiver, demodulating NOAA APT, streaming spectrum data, saving a preset.
- **[FAQ](faq.md)** — what this project is (and isn't), which SDR hardware it supports (all of it — hardware support comes from SDRAngel, not this server), and how it compares to talking to the REST API directly.

## Quick links

- [SDRAngel](https://github.com/f4exb/sdrangel) — the SDR application this server controls.
- [SDRAngel REST API swagger spec](https://github.com/f4exb/sdrangel/blob/master/swagger/sdrangel/api/swagger/swagger.yaml) — the authoritative source for every endpoint this server wraps.
- [Model Context Protocol](https://modelcontextprotocol.io/) — the open standard for connecting LLMs to tools and data sources.
- [Project README](../README.md) — installation and quick reference.
- [CONTRIBUTING.md](../CONTRIBUTING.md) — how to add a new tool group or submit a change.

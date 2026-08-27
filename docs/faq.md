# FAQ

## What is sdrangel-mcp?

An [MCP (Model Context Protocol)](https://modelcontextprotocol.io/) server that exposes [SDRAngel](https://github.com/f4exb/sdrangel)'s REST API as tools an LLM can call. SDRAngel is an open-source software-defined radio (SDR) application supporting dozens of radio types and hundreds of demodulators/decoders; sdrangel-mcp lets you (or an AI assistant acting on your behalf) control it with natural language instead of the desktop GUI — tuning receivers, loading demodulators, managing device sets, saving presets, streaming spectrum data, and more.

## Is this an official SDRAngel project?

No. sdrangel-mcp is an independent, unofficial client of SDRAngel's public REST API. It doesn't modify SDRAngel or bundle any of its code — it just talks to the same HTTP API the SDRAngel web UI and other REST clients use.

## Which SDR hardware does it support?

Whatever SDRAngel itself supports — sdrangel-mcp doesn't touch hardware directly, it drives SDRAngel through its REST API. That currently includes RTL-SDR, HackRF, LimeSDR, BladeRF (including MIMO configurations), Airspy, AirspyHF, PlutoSDR, USRP, SDRplay, FunCube, and many more, plus file/network/test sources for development without hardware. Call `list_device_plugins` to see exactly what your SDRAngel build has available.

## Does it work with Claude Desktop, Claude Code, or other clients?

Yes — all of them. sdrangel-mcp is a plain stdio MCP server with no client-specific code, so it works identically with any MCP-compatible client: Claude Desktop, Claude Code, Cursor, or your own MCP client implementation. See [Getting Started](getting-started.md) for setup instructions for each.

## Why does a `set_*`/`patch_*` call fail with "Invalid JSON request"?

You're probably sending a generic `settings`/`report`/`actions` payload instead of the plugin-specific one SDRAngel actually expects. Every device, channel, and feature plugin wraps its settings under its own JSON key (e.g. `fileInputSettings`, `NFMDemodSettings`) rather than a shared one — call the matching `get_*` tool first to learn that key (`settingsKey` in the response), then echo it back. See [Architecture: the plugin-key wire format](architecture.md#the-plugin-key-wire-format) for the full explanation.

## Does it require modifying SDRAngel or installing a plugin?

No. SDRAngel's REST API is built in and enabled by default on port 8091 — sdrangel-mcp just needs network access to it. Nothing to install on the SDRAngel side.

## Can SDRAngel run on a different machine than the MCP client?

Yes. Set `SDRANGEL_BASE_URL` to point at wherever SDRAngel's REST API is reachable — a headless SDR server on your network, a Raspberry Pi, a cloud instance, or `localhost` if everything's on one machine. sdrangel-mcp itself is a small, dependency-free binary that can run alongside your MCP client, next to SDRAngel, or anywhere in between.

## How is this different from calling SDRAngel's REST API directly?

You certainly can call the REST API directly — the [swagger spec](https://github.com/f4exb/sdrangel/blob/master/swagger/sdrangel/api/swagger/swagger.yaml) is public. sdrangel-mcp adds: typed Go request/response handling instead of hand-built JSON, the plugin-key wire format resolved for you automatically on reads and made explicit on writes, tool descriptions written for an LLM to reason about (when to call something, what a field means, what to fetch first), and a uniform interface (MCP) that works the same way across every AI assistant that supports it, rather than a bespoke integration per client.

## Is it safe to give an LLM control over my SDR setup?

Consider the same things you would for any tool-using LLM integration: `stop_instance` can shut down SDRAngel entirely, and settings changes are applied immediately without a confirmation step at the protocol level (your MCP client may add its own approval flow — Claude Code and Claude Desktop both prompt before running a tool by default). If you're running SDRAngel as a transmitter (`direction: 1` / Tx device sets), be deliberate about what you connect an LLM to and what license/regulatory constraints apply to transmitting in your jurisdiction.

## Where do I report a bug or request a tool?

Open an issue on [GitHub](https://github.com/thereisnotime/sdrangel-mcp/issues). If you're adding a tool yourself, [CONTRIBUTING.md](../CONTRIBUTING.md) has the checklist — and it's worth cross-checking your addition against SDRAngel's [swagger spec](https://github.com/f4exb/sdrangel/blob/master/swagger/sdrangel/api/swagger/swagger.yaml) directly rather than assuming a REST-conventional field name, since several of SDRAngel's own endpoints don't follow the naming you'd expect (see [Architecture](architecture.md)).

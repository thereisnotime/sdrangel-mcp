# Examples

Worked examples of driving SDRAngel through sdrangel-mcp — as prompts you'd give an LLM, and as the underlying tool calls it would make. Useful both for seeing what's possible and for scripting tool calls directly with `sdrangel-mcp call`.

See the [Tools Reference](tools-reference.md) for every tool's full description, and [Architecture](architecture.md) for why settings/report/action calls need a `settingsKey`/`reportKey`/`actionsKey`.

## Tune an RTL-SDR and listen to an FM broadcast station

**Prompt:** _"Load the RTL-SDR into a new device set and tune it to listen to 101.1 MHz FM."_

An LLM with sdrangel-mcp connected would roughly:

1. `list_device_plugins` — find the RTL-SDR's `deviceHwType`, `serial`, and `index`.
2. `add_device_set` with `{"tx": 0}` — open a new Rx device set.
3. `set_device` with `{"deviceSetIndex": 0, "device": {"deviceHwType": "RTLSDR", "direction": 0, "serial": "..."}}`.
4. `get_device_settings` with `{"deviceSetIndex": 0}` — learn the device's `settingsKey` (e.g. `rtlSdrSettings`).
5. `patch_device_settings` to tune the center frequency:
   ```json
   {
     "deviceSetIndex": 0,
     "settings": {
       "deviceHwType": "RTLSDR",
       "direction": 0,
       "settingsKey": "rtlSdrSettings",
       "settings": { "centerFrequency": 101100000 }
     }
   }
   ```
6. `add_channel` with `{"deviceSetIndex": 0, "channel": {"channelType": "WFMDemod"}}` — add a wideband FM demodulator.
7. `get_channel_settings` with `{"deviceSetIndex": 0, "channelIndex": 0}` — learn the channel's `settingsKey` (`WFMDemodSettings`).
8. `patch_channel_settings` to center the demodulator on the station and set audio volume.
9. `start_device` with `{"deviceSetIndex": 0}`.

## Decode NOAA weather satellite APT imagery

**Prompt:** _"Set up device set 0 to receive NOAA 19 on 137.1 MHz and decode APT."_

1. `patch_device_settings` to tune to `137100000` Hz, as above.
2. `add_channel` with `channelType: "APTDemod"`.
3. `get_channel_settings` / `patch_channel_settings` to configure the APT demodulator's deviation and image output path (`settingsKey`: `APTDemodSettings`).
4. `start_device`.
5. `get_channel_report` periodically to check signal lock and decode progress.

Or, if you've saved this setup before: `load_preset` with `{"deviceSetIndex": 0, "preset": {"groupName": "RTL SDR", "name": "NOAA 19", "type": "R"}}` loads the whole device+channel configuration in one call.

## Stream the spectrum display to a remote viewer

**Prompt:** _"Start streaming the spectrum for device set 0 so I can view it remotely."_

1. `get_spectrum_settings` with `{"deviceSetIndex": 0}` to see the current FFT configuration.
2. `patch_spectrum_settings` to adjust FFT size or reference level if needed.
3. `start_spectrum_server` with `{"deviceSetIndex": 0}` — SDRAngel opens a WebSocket server streaming FFT frames.
4. `get_spectrum_server_status` to confirm it's running and get the port.

## Save the current setup as a preset

**Prompt:** _"Save this device set's configuration as a preset called 'NOAA 19' in the 'RTL SDR' group."_

```json
// create_preset
{
  "deviceSetIndex": 0,
  "preset": { "groupName": "RTL SDR", "name": "NOAA 19", "type": "R" }
}
```

Later, `load_preset` with the same identifier restores the whole device+channel configuration in one call — much faster than re-describing the setup from scratch.

## Query what's currently running

**Prompt:** _"What device sets are open right now, and what's each one doing?"_

1. `list_device_sets` — every open device set, its sampling device, and channel count.
2. For each one, `get_device_run_status` and `get_channels_report` to see live state (running/idle, signal levels, decoded data) without needing to know channel indices in advance.

## Scripting without an LLM

Every tool works from the command line too, which is useful for testing or building your own automation on top of SDRAngel's REST API without writing HTTP client code:

```sh
sdrangel-mcp call list_device_sets
sdrangel-mcp call get_device_settings --json '{"deviceSetIndex": 0}'
sdrangel-mcp call patch_device_settings --json '{
  "deviceSetIndex": 0,
  "settings": {
    "deviceHwType": "RTLSDR",
    "direction": 0,
    "settingsKey": "rtlSdrSettings",
    "settings": {"centerFrequency": 433920000}
  }
}'
```

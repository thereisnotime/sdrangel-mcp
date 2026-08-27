package sdrangel

import (
	"encoding/json"
	"strings"
	"testing"
)

// jsonField unmarshals a JSON object and returns the value at key, failing
// the test if the object doesn't parse or the key is absent.
func jsonField(t *testing.T, data json.RawMessage, key string) any {
	t.Helper()
	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatalf("unmarshal %s: %v", data, err)
	}
	v, ok := m[key]
	if !ok {
		t.Fatalf("key %q missing from %s", key, data)
	}
	return v
}

// Fixtures below are trimmed captures from a live SDRAngel 7.27.2 instance.
// They pin down the actual wire format so a future change to the plugin-key
// marshal/unmarshal logic can't silently regress to the generic
// "settings"/"report"/"actions" key the real API rejects.

const fileInputDeviceSettingsJSON = `{
	"deviceHwType": "FileInput",
	"direction": 0,
	"fileInputSettings": {
		"accelerationFactor": 1,
		"loop": 1,
		"title": "FileInput"
	}
}`

const fileInputDeviceReportJSON = `{
	"deviceHwType": "FileInput",
	"direction": 0,
	"fileInputReport": {
		"sampleRate": 48000,
		"sampleSize": 0
	}
}`

const nfmDemodChannelSettingsJSON = `{
	"NFMDemodSettings": {
		"afBandwidth": 3000,
		"rfBandwidth": 12500
	},
	"channelType": "NFMDemod",
	"direction": 0
}`

const nfmDemodChannelReportJSON = `{
	"NFMDemodReport": {
		"channelPowerDB": -150
	},
	"channelType": "NFMDemod",
	"direction": 0
}`

const mapFeatureSettingsJSON = `{
	"MapSettings": {
		"displayNames": 1,
		"terrain": "Ellipsoid"
	},
	"featureType": "Map"
}`

const deviceSetsJSON = `{
	"deviceSets": [
		{
			"channelcount": 1,
			"samplingDevice": {
				"bandwidth": 4000000,
				"centerFrequency": 105000000,
				"deviceNbStreams": 2,
				"deviceStreamIndex": 0,
				"direction": 0,
				"hwType": "BladeRF2",
				"index": 1,
				"sequence": 0,
				"serial": "2c863439b0684d06a4116dbc79171cd9",
				"state": "running"
			},
			"channels": [
				{"index": 0, "id": "WFMDemod", "title": "WFM Demodulator", "direction": 0}
			]
		}
	],
	"devicesetcount": 1,
	"devicesetfocus": 0
}`

const spectrumSettingsJSON = `{
	"decay": 3,
	"histogramStroke": 55,
	"invertedWaterfall": 0,
	"linear": 1,
	"waterfallShare": 0.7,
	"fftSize": 4096,
	"fpsPeriodMs": 50,
	"useCalibration": 0
}`

const loggingInfoJSON = `{
	"consoleLevel": "debug",
	"dumpToFile": 0
}`

const featureSetJSON = `{
	"featurecount": 1,
	"features": [
		{"id": "Map", "index": 0, "title": "Map", "uid": 1787817090076280}
	]
}`

const audioDevicesJSON = `{
	"inputDevices": [
		{"name": "System default device", "index": -1, "sampleRate": 48000, "volume": 1, "isSystemDefault": 0, "defaultUnregistered": 1}
	],
	"nbInputDevices": 1,
	"nbOutputDevices": 1,
	"outputDevices": [
		{"name": "System default device", "index": -1, "sampleRate": 48000, "copyToUDP": 0, "udpAddress": "127.0.0.1", "udpPort": 9998, "defaultUnregistered": 0}
	]
}`

func TestDeviceSettingsUnmarshalCapturesPluginKey(t *testing.T) {
	var got DeviceSettings
	if err := json.Unmarshal([]byte(fileInputDeviceSettingsJSON), &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if got.DeviceHwType != "FileInput" {
		t.Errorf("DeviceHwType = %q, want FileInput", got.DeviceHwType)
	}
	if got.Direction != 0 {
		t.Errorf("Direction = %d, want 0", got.Direction)
	}
	if got.SettingsKey != "fileInputSettings" {
		t.Errorf("SettingsKey = %q, want fileInputSettings", got.SettingsKey)
	}
	if v := jsonField(t, got.Settings, "accelerationFactor"); v != float64(1) {
		t.Errorf("accelerationFactor = %v, want 1", v)
	}
}

func TestDeviceSettingsMarshalUsesPluginKeyNotGenericSettings(t *testing.T) {
	ds := DeviceSettings{
		DeviceHwType: "FileInput",
		Direction:    0,
		SettingsKey:  "fileInputSettings",
		Settings:     json.RawMessage(`{"accelerationFactor":5}`),
	}
	b, err := json.Marshal(ds)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var m map[string]json.RawMessage
	if err := json.Unmarshal(b, &m); err != nil {
		t.Fatalf("unmarshal marshaled output: %v", err)
	}
	if _, ok := m["settings"]; ok {
		t.Errorf("marshaled output has generic %q key; SDRAngel rejects this with \"Invalid JSON request\": %s", "settings", b)
	}
	if _, ok := m["fileInputSettings"]; !ok {
		t.Errorf("marshaled output missing plugin key %q: %s", "fileInputSettings", b)
	}
	if _, ok := m["direction"]; !ok {
		t.Errorf("marshaled output missing %q key (must not be %q): %s", "direction", "tx", b)
	}
	if _, ok := m["tx"]; ok {
		t.Errorf("marshaled output has stale %q key: %s", "tx", b)
	}
}

func TestDeviceSettingsRoundTrip(t *testing.T) {
	original := DeviceSettings{
		DeviceHwType:    "RTLSDR",
		Direction:       0,
		OriginatorIndex: 2,
		SettingsKey:     "rtlSdrSettings",
		Settings:        json.RawMessage(`{"centerFrequency":433920000}`),
	}
	b, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var roundTripped DeviceSettings
	if err := json.Unmarshal(b, &roundTripped); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if roundTripped.DeviceHwType != original.DeviceHwType ||
		roundTripped.Direction != original.Direction ||
		roundTripped.OriginatorIndex != original.OriginatorIndex ||
		roundTripped.SettingsKey != original.SettingsKey ||
		string(roundTripped.Settings) != string(original.Settings) {
		t.Errorf("round trip mismatch: got %+v, want %+v", roundTripped, original)
	}
}

func TestDeviceSettingsMarshalOmitsZeroOriginatorIndex(t *testing.T) {
	ds := DeviceSettings{DeviceHwType: "FileInput", Direction: 0, SettingsKey: "fileInputSettings", Settings: json.RawMessage(`{}`)}
	b, err := json.Marshal(ds)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if strings.Contains(string(b), "originatorIndex") {
		t.Errorf("expected originatorIndex to be omitted when zero, got %s", b)
	}
}

func TestDeviceReportUnmarshalCapturesPluginKey(t *testing.T) {
	var got DeviceReport
	if err := json.Unmarshal([]byte(fileInputDeviceReportJSON), &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if got.ReportKey != "fileInputReport" {
		t.Errorf("ReportKey = %q, want fileInputReport", got.ReportKey)
	}
	if v := jsonField(t, got.Report, "sampleRate"); v != float64(48000) {
		t.Errorf("sampleRate = %v, want 48000", v)
	}
}

func TestChannelSettingsUnmarshalCapturesPluginKey(t *testing.T) {
	var got ChannelSettings
	if err := json.Unmarshal([]byte(nfmDemodChannelSettingsJSON), &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if got.ChannelType != "NFMDemod" {
		t.Errorf("ChannelType = %q, want NFMDemod", got.ChannelType)
	}
	if got.SettingsKey != "NFMDemodSettings" {
		t.Errorf("SettingsKey = %q, want NFMDemodSettings", got.SettingsKey)
	}
	if v := jsonField(t, got.Settings, "rfBandwidth"); v != float64(12500) {
		t.Errorf("rfBandwidth = %v, want 12500", v)
	}
}

func TestChannelSettingsMarshalUsesPluginKey(t *testing.T) {
	cs := ChannelSettings{ChannelType: "NFMDemod", Direction: 0, SettingsKey: "NFMDemodSettings", Settings: json.RawMessage(`{"squelch":-30}`)}
	b, err := json.Marshal(cs)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var m map[string]json.RawMessage
	if err := json.Unmarshal(b, &m); err != nil {
		t.Fatalf("unmarshal marshaled output: %v", err)
	}
	if _, ok := m["settings"]; ok {
		t.Errorf("marshaled output has generic settings key: %s", b)
	}
	if _, ok := m["NFMDemodSettings"]; !ok {
		t.Errorf("marshaled output missing plugin key NFMDemodSettings: %s", b)
	}
}

func TestChannelReportUnmarshalCapturesPluginKey(t *testing.T) {
	var got ChannelReport
	if err := json.Unmarshal([]byte(nfmDemodChannelReportJSON), &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if got.ReportKey != "NFMDemodReport" {
		t.Errorf("ReportKey = %q, want NFMDemodReport", got.ReportKey)
	}
	if v := jsonField(t, got.Report, "channelPowerDB"); v != float64(-150) {
		t.Errorf("channelPowerDB = %v, want -150", v)
	}
}

func TestChannelActionsMarshalUsesPluginKey(t *testing.T) {
	ca := ChannelActions{ChannelType: "NFMDemod", Direction: 0, ActionsKey: "NFMDemodActions", Actions: json.RawMessage(`{}`)}
	b, err := json.Marshal(ca)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var m map[string]json.RawMessage
	if err := json.Unmarshal(b, &m); err != nil {
		t.Fatalf("unmarshal marshaled output: %v", err)
	}
	if _, ok := m["actions"]; ok {
		t.Errorf("marshaled output has generic actions key: %s", b)
	}
	if _, ok := m["NFMDemodActions"]; !ok {
		t.Errorf("marshaled output missing plugin key NFMDemodActions: %s", b)
	}
}

func TestFeatureSettingsUnmarshalCapturesPluginKey(t *testing.T) {
	var got FeatureSettings
	if err := json.Unmarshal([]byte(mapFeatureSettingsJSON), &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if got.FeatureType != "Map" {
		t.Errorf("FeatureType = %q, want Map", got.FeatureType)
	}
	if got.SettingsKey != "MapSettings" {
		t.Errorf("SettingsKey = %q, want MapSettings", got.SettingsKey)
	}
	if v := jsonField(t, got.Settings, "terrain"); v != "Ellipsoid" {
		t.Errorf("terrain = %v, want Ellipsoid", v)
	}
}

func TestFeatureSettingsMarshalUsesPluginKey(t *testing.T) {
	fs := FeatureSettings{FeatureType: "Map", SettingsKey: "MapSettings", Settings: json.RawMessage(`{"displayNames":0}`)}
	b, err := json.Marshal(fs)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var m map[string]json.RawMessage
	if err := json.Unmarshal(b, &m); err != nil {
		t.Fatalf("unmarshal marshaled output: %v", err)
	}
	if _, ok := m["settings"]; ok {
		t.Errorf("marshaled output has generic settings key: %s", b)
	}
	if _, ok := m["MapSettings"]; !ok {
		t.Errorf("marshaled output missing plugin key MapSettings: %s", b)
	}
}

func TestFeatureActionsMarshalUsesPluginKey(t *testing.T) {
	fa := FeatureActions{FeatureType: "Map", ActionsKey: "MapActions", Actions: json.RawMessage(`{}`)}
	b, err := json.Marshal(fa)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var m map[string]json.RawMessage
	if err := json.Unmarshal(b, &m); err != nil {
		t.Fatalf("unmarshal marshaled output: %v", err)
	}
	if _, ok := m["actions"]; ok {
		t.Errorf("marshaled output has generic actions key: %s", b)
	}
	if _, ok := m["MapActions"]; !ok {
		t.Errorf("marshaled output missing plugin key MapActions: %s", b)
	}
}

func TestDeviceSetsUnmarshalMatchesRealShape(t *testing.T) {
	var got DeviceSets
	if err := json.Unmarshal([]byte(deviceSetsJSON), &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if got.DevicesetCount != 1 {
		t.Errorf("DevicesetCount = %d, want 1", got.DevicesetCount)
	}
	if len(got.DevicesetList) != 1 {
		t.Fatalf("len(DevicesetList) = %d, want 1", len(got.DevicesetList))
	}
	sd := got.DevicesetList[0].SamplingDevice
	if sd.HWType != "BladeRF2" {
		t.Errorf("HWType = %q, want BladeRF2", sd.HWType)
	}
	if sd.CenterFrequency != 105000000 {
		t.Errorf("CenterFrequency = %d, want 105000000", sd.CenterFrequency)
	}
	if sd.Bandwidth != 4000000 {
		t.Errorf("Bandwidth = %d, want 4000000", sd.Bandwidth)
	}
	if sd.Direction != 0 {
		t.Errorf("Direction = %d, want 0", sd.Direction)
	}
	if sd.Index != 1 {
		t.Errorf("Index = %d, want 1", sd.Index)
	}
	if len(got.DevicesetList[0].Channels) != 1 || got.DevicesetList[0].Channels[0].ID != "WFMDemod" {
		t.Errorf("Channels = %+v, want one WFMDemod channel", got.DevicesetList[0].Channels)
	}
}

func TestSpectrumSettingsUnmarshalMatchesRealFieldNames(t *testing.T) {
	var got SpectrumSettings
	if err := json.Unmarshal([]byte(spectrumSettingsJSON), &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if got.Decay != 3 {
		t.Errorf("Decay = %d, want 3", got.Decay)
	}
	if got.HistogramStroke != 55 {
		t.Errorf("HistogramStroke = %d, want 55", got.HistogramStroke)
	}
	if got.Linear != 1 {
		t.Errorf("Linear = %d, want 1", got.Linear)
	}
	if got.WaterfallShare != 0.7 {
		t.Errorf("WaterfallShare = %v, want 0.7", got.WaterfallShare)
	}
	if got.FFTSize != 4096 {
		t.Errorf("FFTSize = %d, want 4096", got.FFTSize)
	}
	if got.FPSPeriodMs != 50 {
		t.Errorf("FPSPeriodMs = %d, want 50", got.FPSPeriodMs)
	}
}

func TestAudioDevicesUnmarshalSplitsInputAndOutputShapes(t *testing.T) {
	var got AudioDevices
	if err := json.Unmarshal([]byte(audioDevicesJSON), &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if len(got.InputDevices) != 1 {
		t.Fatalf("len(InputDevices) = %d, want 1", len(got.InputDevices))
	}
	if got.InputDevices[0].Volume != 1 {
		t.Errorf("InputDevices[0].Volume = %v, want 1", got.InputDevices[0].Volume)
	}
	if len(got.OutputDevices) != 1 {
		t.Fatalf("len(OutputDevices) = %d, want 1", len(got.OutputDevices))
	}
	if got.OutputDevices[0].UDPPort != 9998 {
		t.Errorf("OutputDevices[0].UDPPort = %d, want 9998", got.OutputDevices[0].UDPPort)
	}
	if got.OutputDevices[0].UDPAddress != "127.0.0.1" {
		t.Errorf("OutputDevices[0].UDPAddress = %q, want 127.0.0.1", got.OutputDevices[0].UDPAddress)
	}
}

func TestLoggingInfoUnmarshalUsesDumpToFile(t *testing.T) {
	var got LoggingInfo
	if err := json.Unmarshal([]byte(loggingInfoJSON), &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if got.ConsoleLevel != "debug" {
		t.Errorf("ConsoleLevel = %q, want debug", got.ConsoleLevel)
	}
	if got.DumpToFile != 0 {
		t.Errorf("DumpToFile = %d, want 0", got.DumpToFile)
	}
}

func TestFeatureSetInfoUnmarshalCapturesUID(t *testing.T) {
	var got FeatureSetInfo
	if err := json.Unmarshal([]byte(featureSetJSON), &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if got.FeatureCount != 1 {
		t.Errorf("FeatureCount = %d, want 1", got.FeatureCount)
	}
	if len(got.Features) != 1 {
		t.Fatalf("len(Features) = %d, want 1", len(got.Features))
	}
	f := got.Features[0]
	if f.FeatureID != "Map" || f.Title != "Map" {
		t.Errorf("Features[0] = %+v, want id/title Map", f)
	}
	if f.UID != 1787817090076280 {
		t.Errorf("UID = %d, want 1787817090076280", f.UID)
	}
}

func TestDeviceLinkMarshalUsesDirectionNotTx(t *testing.T) {
	dl := DeviceLink{DeviceHwType: "AudioInput", Direction: 0, Serial: "0"}
	b, err := json.Marshal(dl)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var m map[string]json.RawMessage
	if err := json.Unmarshal(b, &m); err != nil {
		t.Fatalf("unmarshal marshaled output: %v", err)
	}
	if _, ok := m["direction"]; !ok {
		t.Errorf("marshaled output missing direction key: %s", b)
	}
	if _, ok := m["tx"]; ok {
		t.Errorf("marshaled output has stale tx key: %s", b)
	}
}

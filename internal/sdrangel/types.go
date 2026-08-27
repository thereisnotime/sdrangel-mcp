package sdrangel

import "encoding/json"

// splitPluginKey parses a JSON object, returning the raw field map plus the
// single top-level key/value pair not present in known (SDRAngel wraps every
// device/channel/feature settings, report, and actions payload under a
// plugin-specific key such as "fileInputSettings" or "NFMDemodReport" rather
// than a generic "settings"/"report"/"actions" key).
func splitPluginKey(data []byte, known ...string) (raw map[string]json.RawMessage, key string, value json.RawMessage, err error) {
	if err = json.Unmarshal(data, &raw); err != nil {
		return nil, "", nil, err
	}
	isKnown := make(map[string]bool, len(known))
	for _, k := range known {
		isKnown[k] = true
	}
	for k, v := range raw {
		if !isKnown[k] {
			key, value = k, v
			break
		}
	}
	return raw, key, value, nil
}

// marshalWithPluginKey builds a JSON object from known fields plus one
// plugin-specific key/value pair (see splitPluginKey).
func marshalWithPluginKey(fields map[string]json.RawMessage, key string, value json.RawMessage) ([]byte, error) {
	m := make(map[string]json.RawMessage, len(fields)+1)
	for k, v := range fields {
		m[k] = v
	}
	if key != "" {
		m[key] = value
	}
	return json.Marshal(m)
}

func mustMarshal(v any) json.RawMessage {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return b
}

// InstanceSummary holds information about the running SDRAngel instance.
type InstanceSummary struct {
	Appname       string   `json:"appname"`
	Version       string   `json:"version"`
	QtVersion     string   `json:"qtVersion"`
	DspRxBits     int      `json:"dspRxBits"`
	DspTxBits     int      `json:"dspTxBits"`
	PID           int      `json:"pid"`
	Architecture  string   `json:"architecture"`
	OS            string   `json:"os"`
	UserArguments []string `json:"userArguments,omitempty"`
	LogFileName   string   `json:"logFileName,omitempty"`
	ServerAddress string   `json:"serverAddress,omitempty"`
	ServerPort    int      `json:"serverPort,omitempty"`
}

// InstanceConfig holds global preferences and commands.
type InstanceConfig struct {
	Preferences json.RawMessage `json:"preferences,omitempty"`
	Commands    json.RawMessage `json:"commands,omitempty"`
}

// Location holds the GPS location used for signal calculations.
type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Altitude  float64 `json:"altitude"`
}

// AvailableChannelOrFeatureList lists available channel or feature plugins.
type AvailableChannelOrFeatureList struct {
	ChannelCount int          `json:"channelcount,omitempty"`
	Channels     []PluginInfo `json:"channels,omitempty"`
	FeatureCount int          `json:"featurecount,omitempty"`
	Features     []PluginInfo `json:"features,omitempty"`
}

// AvailableDeviceList lists available device plugins.
type AvailableDeviceList struct {
	DeviceCount int               `json:"devicecount,omitempty"`
	Devices     []AvailableDevice `json:"devices,omitempty"`
}

// AvailableDevice describes a single detected device instance (a physical
// device SDRangel can see, not just a supported hardware type).
type AvailableDevice struct {
	DeviceNbStreams int    `json:"deviceNbStreams"`
	DeviceSetIndex  int    `json:"deviceSetIndex"`
	Direction       int    `json:"direction"`
	DisplayedName   string `json:"displayedName"`
	HardwareType    string `json:"hwType,omitempty"`
	Index           int    `json:"index"`
	Sequence        int    `json:"sequence"`
	Serial          string `json:"serial,omitempty"`
}

// PluginInfo describes a single channel or feature plugin.
type PluginInfo struct {
	ID        string `json:"id"`
	IDURI     string `json:"idURI,omitempty"`
	Name      string `json:"name"`
	Version   string `json:"version"`
	Direction int    `json:"direction,omitempty"`
	Index     int    `json:"index,omitempty"`
}

// AudioDevices lists available audio input and output devices.
type AudioDevices struct {
	NbInputDevices  int                     `json:"nbInputDevices"`
	NbOutputDevices int                     `json:"nbOutputDevices"`
	InputDevices    []AudioInputDeviceInfo  `json:"inputDevices,omitempty"`
	OutputDevices   []AudioOutputDeviceInfo `json:"outputDevices,omitempty"`
}

// AudioInputDeviceInfo describes a single available audio input device.
type AudioInputDeviceInfo struct {
	Name                string  `json:"name"`
	Index               int     `json:"index"`
	SampleRate          int     `json:"sampleRate"`
	Volume              float64 `json:"volume,omitempty"`
	IsSystemDefault     int     `json:"isSystemDefault,omitempty"`
	DefaultUnregistered int     `json:"defaultUnregistered,omitempty"`
}

// AudioOutputDeviceInfo describes a single available audio output device.
type AudioOutputDeviceInfo struct {
	Name                string `json:"name"`
	Index               int    `json:"index"`
	SampleRate          int    `json:"sampleRate"`
	DefaultUnregistered int    `json:"defaultUnregistered,omitempty"`
	CopyToUDP           int    `json:"copyToUDP,omitempty"`
	UDPAddress          string `json:"udpAddress,omitempty"`
	UDPPort             int    `json:"udpPort,omitempty"`
	UDPChannelMode      int    `json:"udpChannelMode,omitempty"`
	UDPChannelCodec     int    `json:"udpChannelCodec,omitempty"`
	UDPDecimationFactor int    `json:"udpDecimationFactor,omitempty"`
	UDPUsesRTP          int    `json:"udpUsesRTP,omitempty"`
	RecordToFile        int    `json:"recordToFile,omitempty"`
	RecordSilenceTime   int    `json:"recordSilenceTime,omitempty"`
}

// AudioInputDevice holds parameters for an audio input device.
type AudioInputDevice struct {
	Name       string  `json:"name"`
	SampleRate int     `json:"sampleRate,omitempty"`
	Volume     float64 `json:"volume,omitempty"`
}

// AudioOutputDevice holds parameters for an audio output device.
type AudioOutputDevice struct {
	Name       string  `json:"name"`
	SampleRate int     `json:"sampleRate,omitempty"`
	Volume     float64 `json:"volume,omitempty"`
	CopyToUDP  int     `json:"copyToUDP,omitempty"`
	UDPAddress string  `json:"udpAddress,omitempty"`
	UDPPort    int     `json:"udpPort,omitempty"`
}

// LoggingInfo holds logging configuration.
type LoggingInfo struct {
	ConsoleLevel    string `json:"consoleLevel"`
	FileLevel       string `json:"fileLevel,omitempty"`
	FileName        string `json:"fileName,omitempty"`
	ConsoleMinLevel int    `json:"consoleMinLevel,omitempty"`
	FileMinLevel    int    `json:"fileMinLevel,omitempty"`
	DumpToFile      int    `json:"dumpToFile,omitempty"`
}

// PresetIdentifier uniquely identifies a preset.
type PresetIdentifier struct {
	GroupName       string `json:"groupName"`
	Name            string `json:"name"`
	Type            string `json:"type,omitempty"`
	CenterFrequency int64  `json:"centerFrequency,omitempty"`
}

// PresetTransfer associates a preset with a device set for load/save.
type PresetTransfer struct {
	DeviceSetIndex int              `json:"deviceSetIndex"`
	Preset         PresetIdentifier `json:"preset"`
}

// PresetKeys identifies a preset for deletion.
type PresetKeys struct {
	GroupName string `json:"groupName"`
	Name      string `json:"name"`
	Type      string `json:"type,omitempty"`
}

// Presets is the top-level list of preset groups.
type Presets struct {
	NbGroups int           `json:"nbGroups"`
	Groups   []PresetGroup `json:"groups,omitempty"`
}

// PresetGroup is a named group of presets.
type PresetGroup struct {
	GroupName string             `json:"groupName"`
	NbPresets int                `json:"nbPresets"`
	Presets   []PresetIdentifier `json:"presets,omitempty"`
}

// ConfigurationKeys identifies a configuration.
type ConfigurationKeys struct {
	GroupName string `json:"groupName"`
	Name      string `json:"name"`
}

// Configurations is the top-level list of configuration groups.
type Configurations struct {
	NbGroups int                  `json:"nbGroups"`
	Groups   []ConfigurationGroup `json:"groups,omitempty"`
}

// ConfigurationGroup is a named group of configurations.
type ConfigurationGroup struct {
	GroupName        string              `json:"groupName"`
	NbConfigurations int                 `json:"nbConfigurations"`
	Configurations   []ConfigurationKeys `json:"configurations,omitempty"`
}

// DeviceSets lists all device sets.
type DeviceSets struct {
	DevicesetCount int             `json:"devicesetcount"`
	DevicesetFocus int             `json:"devicesetfocus,omitempty"`
	DevicesetList  []DeviceSetInfo `json:"deviceSets,omitempty"`
}

// DeviceSetInfo describes a single device set.
type DeviceSetInfo struct {
	Index          int           `json:"index,omitempty"`
	SamplingDevice DeviceDesc    `json:"samplingDevice"`
	ChannelCount   int           `json:"channelcount"`
	Channels       []ChannelDesc `json:"channels,omitempty"`
}

// DeviceDesc describes a sampling device.
type DeviceDesc struct {
	ID                string `json:"id,omitempty"`
	Serial            string `json:"serial,omitempty"`
	Sequence          int    `json:"sequence"`
	DeviceNbStreams   int    `json:"deviceNbStreams,omitempty"`
	DeviceStreamIndex int    `json:"deviceStreamIndex,omitempty"`
	HWType            string `json:"hwType,omitempty"`
	Direction         int    `json:"direction"`
	Index             int    `json:"index,omitempty"`
	State             string `json:"state,omitempty"`
	CenterFrequency   int64  `json:"centerFrequency,omitempty"`
	Bandwidth         int64  `json:"bandwidth,omitempty"`
}

// ChannelDesc describes a channel in a device set.
type ChannelDesc struct {
	Index          int    `json:"index"`
	ID             string `json:"id"`
	Title          string `json:"title,omitempty"`
	Direction      int    `json:"direction"`
	DeltaFrequency int64  `json:"deltaFrequency,omitempty"`
}

// DeviceSettings holds device-specific settings. SDRAngel wraps the plugin's
// settings object under a plugin-specific key (e.g. "fileInputSettings",
// "rtlSdrSettings") rather than a generic "settings" key: SettingsKey holds
// that key name (as returned by GetDeviceSettings) and Settings holds its raw
// JSON value. To change settings, call get_device_settings first to learn
// the SettingsKey, then echo it back with SetDeviceSettings/PatchDeviceSettings.
type DeviceSettings struct {
	DeviceHwType    string
	Direction       int
	OriginatorIndex int
	SettingsKey     string
	Settings        json.RawMessage
}

func (d DeviceSettings) MarshalJSON() ([]byte, error) {
	fields := map[string]json.RawMessage{
		"deviceHwType": mustMarshal(d.DeviceHwType),
		"direction":    mustMarshal(d.Direction),
	}
	if d.OriginatorIndex != 0 {
		fields["originatorIndex"] = mustMarshal(d.OriginatorIndex)
	}
	return marshalWithPluginKey(fields, d.SettingsKey, d.Settings)
}

func (d *DeviceSettings) UnmarshalJSON(data []byte) error {
	raw, key, value, err := splitPluginKey(data, "deviceHwType", "direction", "originatorIndex")
	if err != nil {
		return err
	}
	if v, ok := raw["deviceHwType"]; ok {
		if err := json.Unmarshal(v, &d.DeviceHwType); err != nil {
			return err
		}
	}
	if v, ok := raw["direction"]; ok {
		if err := json.Unmarshal(v, &d.Direction); err != nil {
			return err
		}
	}
	if v, ok := raw["originatorIndex"]; ok {
		if err := json.Unmarshal(v, &d.OriginatorIndex); err != nil {
			return err
		}
	}
	d.SettingsKey, d.Settings = key, value
	return nil
}

// DeviceReport holds a device's runtime report. Report is wrapped under a
// plugin-specific key (e.g. "fileInputReport"); ReportKey holds that key name.
type DeviceReport struct {
	DeviceHwType string
	Direction    int
	ReportKey    string
	Report       json.RawMessage
}

func (d DeviceReport) MarshalJSON() ([]byte, error) {
	fields := map[string]json.RawMessage{
		"deviceHwType": mustMarshal(d.DeviceHwType),
		"direction":    mustMarshal(d.Direction),
	}
	return marshalWithPluginKey(fields, d.ReportKey, d.Report)
}

func (d *DeviceReport) UnmarshalJSON(data []byte) error {
	raw, key, value, err := splitPluginKey(data, "deviceHwType", "direction")
	if err != nil {
		return err
	}
	if v, ok := raw["deviceHwType"]; ok {
		if err := json.Unmarshal(v, &d.DeviceHwType); err != nil {
			return err
		}
	}
	if v, ok := raw["direction"]; ok {
		if err := json.Unmarshal(v, &d.Direction); err != nil {
			return err
		}
	}
	d.ReportKey, d.Report = key, value
	return nil
}

// DeviceState holds the run state of a device.
type DeviceState struct {
	State string `json:"state"`
}

// DeviceLink identifies a device to load into a device set.
type DeviceLink struct {
	DeviceHwType string `json:"deviceHwType"`
	Direction    int    `json:"direction"`
	Index        int    `json:"index,omitempty"`
	Serial       string `json:"serial,omitempty"`
}

// ChannelSettings holds channel-specific settings. SDRAngel wraps the
// plugin's settings object under a plugin-specific key (e.g.
// "NFMDemodSettings") rather than a generic "settings" key: SettingsKey
// holds that key name (as returned by GetChannelSettings) and Settings
// holds its raw JSON value. To change settings, call get_channel_settings
// first to learn the SettingsKey, then echo it back with
// SetChannelSettings/PatchChannelSettings.
type ChannelSettings struct {
	ChannelType     string
	Direction       int
	OriginatorIndex int
	SettingsKey     string
	Settings        json.RawMessage
}

func (c ChannelSettings) MarshalJSON() ([]byte, error) {
	fields := map[string]json.RawMessage{
		"channelType": mustMarshal(c.ChannelType),
		"direction":   mustMarshal(c.Direction),
	}
	if c.OriginatorIndex != 0 {
		fields["originatorIndex"] = mustMarshal(c.OriginatorIndex)
	}
	return marshalWithPluginKey(fields, c.SettingsKey, c.Settings)
}

func (c *ChannelSettings) UnmarshalJSON(data []byte) error {
	raw, key, value, err := splitPluginKey(data, "channelType", "direction", "originatorIndex")
	if err != nil {
		return err
	}
	if v, ok := raw["channelType"]; ok {
		if err := json.Unmarshal(v, &c.ChannelType); err != nil {
			return err
		}
	}
	if v, ok := raw["direction"]; ok {
		if err := json.Unmarshal(v, &c.Direction); err != nil {
			return err
		}
	}
	if v, ok := raw["originatorIndex"]; ok {
		if err := json.Unmarshal(v, &c.OriginatorIndex); err != nil {
			return err
		}
	}
	c.SettingsKey, c.Settings = key, value
	return nil
}

// ChannelReport holds a channel's runtime report. Report is wrapped under a
// plugin-specific key (e.g. "NFMDemodReport"); ReportKey holds that key name.
type ChannelReport struct {
	ChannelType string
	Direction   int
	ReportKey   string
	Report      json.RawMessage
}

func (c ChannelReport) MarshalJSON() ([]byte, error) {
	fields := map[string]json.RawMessage{
		"channelType": mustMarshal(c.ChannelType),
		"direction":   mustMarshal(c.Direction),
	}
	return marshalWithPluginKey(fields, c.ReportKey, c.Report)
}

func (c *ChannelReport) UnmarshalJSON(data []byte) error {
	raw, key, value, err := splitPluginKey(data, "channelType", "direction")
	if err != nil {
		return err
	}
	if v, ok := raw["channelType"]; ok {
		if err := json.Unmarshal(v, &c.ChannelType); err != nil {
			return err
		}
	}
	if v, ok := raw["direction"]; ok {
		if err := json.Unmarshal(v, &c.Direction); err != nil {
			return err
		}
	}
	c.ReportKey, c.Report = key, value
	return nil
}

// ChannelActions holds actions to execute on a channel. Actions are wrapped
// under a plugin-specific key (e.g. "NFMDemodActions"); ActionsKey holds
// that key name.
type ChannelActions struct {
	ChannelType string
	Direction   int
	ActionsKey  string
	Actions     json.RawMessage
}

func (c ChannelActions) MarshalJSON() ([]byte, error) {
	fields := map[string]json.RawMessage{
		"channelType": mustMarshal(c.ChannelType),
		"direction":   mustMarshal(c.Direction),
	}
	return marshalWithPluginKey(fields, c.ActionsKey, c.Actions)
}

func (c *ChannelActions) UnmarshalJSON(data []byte) error {
	raw, key, value, err := splitPluginKey(data, "channelType", "direction")
	if err != nil {
		return err
	}
	if v, ok := raw["channelType"]; ok {
		if err := json.Unmarshal(v, &c.ChannelType); err != nil {
			return err
		}
	}
	if v, ok := raw["direction"]; ok {
		if err := json.Unmarshal(v, &c.Direction); err != nil {
			return err
		}
	}
	c.ActionsKey, c.Actions = key, value
	return nil
}

// ChannelsReport holds reports for all channels in a device set.
type ChannelsReport struct {
	DeviceSetIndex int             `json:"deviceSetIndex"`
	ChannelCount   int             `json:"channelcount"`
	Channels       []ChannelReport `json:"channels,omitempty"`
}

// ChannelAdd specifies the type of channel to add.
type ChannelAdd struct {
	ChannelType string `json:"channelType"`
	Tx          int    `json:"tx,omitempty"`
}

// FeatureSetInfo describes the feature set.
type FeatureSetInfo struct {
	FeatureCount int           `json:"featurecount"`
	Features     []FeatureDesc `json:"features,omitempty"`
}

// FeatureDesc describes a single feature instance.
type FeatureDesc struct {
	Index     int    `json:"index"`
	FeatureID string `json:"id"`
	Title     string `json:"title,omitempty"`
	State     string `json:"state,omitempty"`
	UID       int64  `json:"uid,omitempty"`
}

// FeatureSettings holds feature-specific settings. SDRAngel wraps the
// plugin's settings object under a plugin-specific key (e.g. "MapSettings")
// rather than a generic "settings" key: SettingsKey holds that key name (as
// returned by GetFeatureSettings) and Settings holds its raw JSON value. To
// change settings, call get_feature_settings first to learn the
// SettingsKey, then echo it back with
// SetFeatureSettings/PatchFeatureSettings.
type FeatureSettings struct {
	FeatureType     string
	OriginatorIndex int
	SettingsKey     string
	Settings        json.RawMessage
}

func (f FeatureSettings) MarshalJSON() ([]byte, error) {
	fields := map[string]json.RawMessage{
		"featureType": mustMarshal(f.FeatureType),
	}
	if f.OriginatorIndex != 0 {
		fields["originatorIndex"] = mustMarshal(f.OriginatorIndex)
	}
	return marshalWithPluginKey(fields, f.SettingsKey, f.Settings)
}

func (f *FeatureSettings) UnmarshalJSON(data []byte) error {
	raw, key, value, err := splitPluginKey(data, "featureType", "originatorIndex")
	if err != nil {
		return err
	}
	if v, ok := raw["featureType"]; ok {
		if err := json.Unmarshal(v, &f.FeatureType); err != nil {
			return err
		}
	}
	if v, ok := raw["originatorIndex"]; ok {
		if err := json.Unmarshal(v, &f.OriginatorIndex); err != nil {
			return err
		}
	}
	f.SettingsKey, f.Settings = key, value
	return nil
}

// FeatureReport holds a feature's runtime report. Report is wrapped under a
// plugin-specific key (e.g. "MapReport"); ReportKey holds that key name.
type FeatureReport struct {
	FeatureType string
	ReportKey   string
	Report      json.RawMessage
}

func (f FeatureReport) MarshalJSON() ([]byte, error) {
	fields := map[string]json.RawMessage{
		"featureType": mustMarshal(f.FeatureType),
	}
	return marshalWithPluginKey(fields, f.ReportKey, f.Report)
}

func (f *FeatureReport) UnmarshalJSON(data []byte) error {
	raw, key, value, err := splitPluginKey(data, "featureType")
	if err != nil {
		return err
	}
	if v, ok := raw["featureType"]; ok {
		if err := json.Unmarshal(v, &f.FeatureType); err != nil {
			return err
		}
	}
	f.ReportKey, f.Report = key, value
	return nil
}

// FeatureActions holds actions to execute on a feature. Actions are wrapped
// under a plugin-specific key (e.g. "MapActions"); ActionsKey holds that key
// name.
type FeatureActions struct {
	FeatureType string
	ActionsKey  string
	Actions     json.RawMessage
}

func (f FeatureActions) MarshalJSON() ([]byte, error) {
	fields := map[string]json.RawMessage{
		"featureType": mustMarshal(f.FeatureType),
	}
	return marshalWithPluginKey(fields, f.ActionsKey, f.Actions)
}

func (f *FeatureActions) UnmarshalJSON(data []byte) error {
	raw, key, value, err := splitPluginKey(data, "featureType")
	if err != nil {
		return err
	}
	if v, ok := raw["featureType"]; ok {
		if err := json.Unmarshal(v, &f.FeatureType); err != nil {
			return err
		}
	}
	f.ActionsKey, f.Actions = key, value
	return nil
}

// FeatureState holds the run state of a feature.
type FeatureState struct {
	State string `json:"state"`
}

// FeatureAdd specifies the type of feature to add.
type FeatureAdd struct {
	FeatureType string `json:"featureType"`
}

// SpectrumSettings holds spectrum display settings.
type SpectrumSettings struct {
	FFTSize               int     `json:"fftSize,omitempty"`
	FFTOverlap            int     `json:"fftOverlap,omitempty"`
	FFTWindow             int     `json:"fftWindow,omitempty"`
	RefLevel              float64 `json:"refLevel,omitempty"`
	PowerRange            float64 `json:"powerRange,omitempty"`
	DisplayHistogram      int     `json:"displayHistogram,omitempty"`
	Decay                 int     `json:"decay,omitempty"`
	DisplayGrid           int     `json:"displayGrid,omitempty"`
	DisplayGridIntensity  int     `json:"displayGridIntensity,omitempty"`
	DisplayTraceIntensity int     `json:"displayTraceIntensity,omitempty"`
	InvertedWaterfall     int     `json:"invertedWaterfall,omitempty"`
	DecayDivisor          int     `json:"decayDivisor,omitempty"`
	HistogramStroke       int     `json:"histogramStroke,omitempty"`
	DisplayMaxHold        int     `json:"displayMaxHold,omitempty"`
	MaxHoldMultiplier     int     `json:"maxHoldMultiplier,omitempty"`
	DisplayWaterfall      int     `json:"displayWaterfall,omitempty"`
	WaterfallShare        float64 `json:"waterfallShare,omitempty"`
	DisplayCurrent        int     `json:"displayCurrent,omitempty"`
	AveragingMode         int     `json:"averagingMode,omitempty"`
	AveragingValue        int64   `json:"averagingValue,omitempty"`
	Linear                int     `json:"linear,omitempty"`
	SSBFilter             int     `json:"ssb,omitempty"`
	USBFilter             int     `json:"usb,omitempty"`
	MarkersDisplay        int     `json:"markersDisplay,omitempty"`
	Measurement           int     `json:"measurement,omitempty"`
	MeasurementCenter     int64   `json:"measurementCenterFrequency,omitempty"`
	MeasurementBW         int64   `json:"measurementBandwidth,omitempty"`
	FPSPeriodMs           int     `json:"fpsPeriodMs,omitempty"`
	UseCalibration        int     `json:"useCalibration,omitempty"`
	CalibrationInterpMode int     `json:"calibrationInterpMode,omitempty"`
	WSSpectrum            int     `json:"wsSpectrum,omitempty"`
	WSSpectrumAddress     string  `json:"wsSpectrumAddress,omitempty"`
	WSSpectrumPort        int     `json:"wsSpectrumPort,omitempty"`
}

// SpectrumServer holds spectrum websocket server state.
type SpectrumServer struct {
	Run     int    `json:"run"`
	Address string `json:"address,omitempty"`
	Port    int    `json:"port,omitempty"`
}

// WorkspaceInfo holds information about a workspace.
type WorkspaceInfo struct {
	Index int `json:"index"`
}

// SuccessResponse is a generic success message.
type SuccessResponse struct {
	Message string `json:"message,omitempty"`
}

// DeviceActions holds actions to execute on a device. Actions are wrapped
// under a plugin-specific key (e.g. "fileInputActions"); ActionsKey holds
// that key name (following the same convention as ChannelActions/FeatureActions).
type DeviceActions struct {
	DeviceHwType string
	Direction    int
	ActionsKey   string
	Actions      json.RawMessage
}

func (d DeviceActions) MarshalJSON() ([]byte, error) {
	fields := map[string]json.RawMessage{
		"deviceHwType": mustMarshal(d.DeviceHwType),
		"direction":    mustMarshal(d.Direction),
	}
	return marshalWithPluginKey(fields, d.ActionsKey, d.Actions)
}

func (d *DeviceActions) UnmarshalJSON(data []byte) error {
	raw, key, value, err := splitPluginKey(data, "deviceHwType", "direction")
	if err != nil {
		return err
	}
	if v, ok := raw["deviceHwType"]; ok {
		if err := json.Unmarshal(v, &d.DeviceHwType); err != nil {
			return err
		}
	}
	if v, ok := raw["direction"]; ok {
		if err := json.Unmarshal(v, &d.Direction); err != nil {
			return err
		}
	}
	d.ActionsKey, d.Actions = key, value
	return nil
}

// FilePath is a filesystem path, used to import a preset or configuration
// from a file on the server.
type FilePath struct {
	FilePath string `json:"filePath"`
}

// PresetExport specifies which preset to export and the file path to export
// it to (server-side path).
type PresetExport struct {
	FilePath string           `json:"filePath,omitempty"`
	Preset   PresetIdentifier `json:"preset"`
}

// Base64Blob holds a preset or configuration serialized as a base64 blob.
type Base64Blob struct {
	Blob string `json:"blob"`
}

// ConfigurationImportExport specifies the file path and configuration
// identification for configuration import/export.
type ConfigurationImportExport struct {
	FilePath      string            `json:"filePath,omitempty"`
	Configuration ConfigurationKeys `json:"configuration"`
}

// FeaturePresets is the top-level list of feature preset groups.
type FeaturePresets struct {
	NbGroups int                  `json:"nbGroups"`
	Groups   []FeaturePresetGroup `json:"groups,omitempty"`
}

// FeaturePresetGroup is a named group of feature presets.
type FeaturePresetGroup struct {
	GroupName string              `json:"groupName"`
	NbPresets int                 `json:"nbPresets"`
	Presets   []FeaturePresetItem `json:"presets,omitempty"`
}

// FeaturePresetItem describes a single feature preset within a group.
type FeaturePresetItem struct {
	Description string `json:"description"`
}

// FeaturePresetIdentifier uniquely identifies a feature preset.
type FeaturePresetIdentifier struct {
	GroupName   string `json:"groupName"`
	Description string `json:"description"`
}

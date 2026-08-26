package sdrangel

import "encoding/json"

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
	ChannelList []PluginInfo `json:"channelList,omitempty"`
	FeatureList []PluginInfo `json:"featureList,omitempty"`
}

// AvailableDeviceList lists available device plugins.
type AvailableDeviceList struct {
	DeviceList []PluginInfo `json:"deviceList,omitempty"`
}

// PluginInfo describes a single plugin.
type PluginInfo struct {
	DisplayedName string `json:"displayedName"`
	Version       string `json:"version"`
	Copyright     string `json:"copyright,omitempty"`
	HardwareType  string `json:"hwType,omitempty"`
	Tx            int    `json:"tx,omitempty"`
}

// AudioDevices lists available audio input and output devices.
type AudioDevices struct {
	NbInputDevices  int           `json:"nbInputDevices"`
	NbOutputDevices int           `json:"nbOutputDevices"`
	InputDevices    []AudioDevice `json:"audioInputDevice,omitempty"`
	OutputDevices   []AudioDevice `json:"audioOutputDevice,omitempty"`
}

// AudioDevice describes a single audio device.
type AudioDevice struct {
	Name                string  `json:"name"`
	Index               int     `json:"index"`
	SampleRate          int     `json:"sampleRate"`
	Volume              float64 `json:"volume,omitempty"`
	DefaultUnregistered bool    `json:"defaultUnregistered,omitempty"`
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
	LogToFile       int    `json:"logToFile,omitempty"`
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
	GroupName        string             `json:"groupName"`
	NbConfigurations int                `json:"nbConfigurations"`
	Configurations   []ConfigurationKeys `json:"configurations,omitempty"`
}

// DeviceSets lists all device sets.
type DeviceSets struct {
	DevicesetList []DeviceSetInfo `json:"devicesetList,omitempty"`
}

// DeviceSetInfo describes a single device set.
type DeviceSetInfo struct {
	Index          int          `json:"index"`
	SamplingDevice DeviceDesc   `json:"samplingDevice"`
	ChannelCount   int          `json:"channelcount"`
	Channels       []ChannelDesc `json:"channels,omitempty"`
}

// DeviceDesc describes a sampling device.
type DeviceDesc struct {
	ID                string `json:"id"`
	Serial            string `json:"serial,omitempty"`
	Sequence          int    `json:"sequence"`
	DeviceNbStreams    int    `json:"deviceNbStreams,omitempty"`
	DeviceStreamIndex int    `json:"deviceStreamIndex,omitempty"`
	HWType            string `json:"hwType,omitempty"`
	Tx                int    `json:"tx"`
	State             string `json:"state,omitempty"`
}

// ChannelDesc describes a channel in a device set.
type ChannelDesc struct {
	Index          int    `json:"index"`
	ID             string `json:"id"`
	Title          string `json:"title,omitempty"`
	Direction      int    `json:"direction"`
	DeltaFrequency int64  `json:"deltaFrequency,omitempty"`
}

// DeviceSettings holds device-specific settings (plugin payload is opaque).
type DeviceSettings struct {
	DeviceHwType    string          `json:"deviceHwType"`
	Tx              int             `json:"tx"`
	OriginatorIndex int             `json:"originatorIndex,omitempty"`
	Settings        json.RawMessage `json:"settings,omitempty"`
}

// DeviceReport holds a device's runtime report.
type DeviceReport struct {
	DeviceHwType string          `json:"deviceHwType"`
	Tx           int             `json:"tx"`
	Report       json.RawMessage `json:"report,omitempty"`
}

// DeviceState holds the run state of a device.
type DeviceState struct {
	State string `json:"state"`
}

// DeviceLink identifies a device to load into a device set.
type DeviceLink struct {
	DeviceHwType string `json:"deviceHwType"`
	Tx           int    `json:"tx"`
	Index        int    `json:"index,omitempty"`
	Serial       string `json:"serial,omitempty"`
}

// ChannelSettings holds channel-specific settings.
type ChannelSettings struct {
	ChannelType     string          `json:"channelType"`
	Direction       int             `json:"direction"`
	OriginatorIndex int             `json:"originatorIndex,omitempty"`
	Settings        json.RawMessage `json:"settings,omitempty"`
}

// ChannelReport holds a channel's runtime report.
type ChannelReport struct {
	ChannelType string          `json:"channelType"`
	Direction   int             `json:"direction"`
	Report      json.RawMessage `json:"report,omitempty"`
}

// ChannelActions holds actions to execute on a channel.
type ChannelActions struct {
	ChannelType string          `json:"channelType"`
	Direction   int             `json:"direction"`
	Actions     json.RawMessage `json:"actions,omitempty"`
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
}

// FeatureSettings holds feature-specific settings.
type FeatureSettings struct {
	FeatureType     string          `json:"featureType"`
	OriginatorIndex int             `json:"originatorIndex,omitempty"`
	Settings        json.RawMessage `json:"settings,omitempty"`
}

// FeatureReport holds a feature's runtime report.
type FeatureReport struct {
	FeatureType string          `json:"featureType"`
	Report      json.RawMessage `json:"report,omitempty"`
}

// FeatureActions holds actions to execute on a feature.
type FeatureActions struct {
	FeatureType string          `json:"featureType"`
	Actions     json.RawMessage `json:"actions,omitempty"`
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
	FFTSize           int     `json:"fftSize,omitempty"`
	FFTOverlap        int     `json:"fftOverlap,omitempty"`
	FFTWindow         int     `json:"fftWindow,omitempty"`
	RefLevel          float64 `json:"refLevel,omitempty"`
	PowerRange        float64 `json:"powerRange,omitempty"`
	DisplayHistogram  int     `json:"displayHistogram,omitempty"`
	DecayRate         int     `json:"decayRate,omitempty"`
	DisplayGrid       int     `json:"displayGrid,omitempty"`
	Invert            int     `json:"invert,omitempty"`
	DecayDivisor      int     `json:"decayDivisor,omitempty"`
	HistoStroke       int     `json:"histoStroke,omitempty"`
	DisplayMaxHold    int     `json:"displayMaxHold,omitempty"`
	MaxHoldMultiplier int     `json:"maxHoldMultiplier,omitempty"`
	DisplayWaterfall  int     `json:"displayWaterfall,omitempty"`
	WaterfallHeight   int     `json:"waterfallHeight,omitempty"`
	DisplayCurrent    int     `json:"displayCurrent,omitempty"`
	AveragingMode     int     `json:"averagingMode,omitempty"`
	AveragingValue    int64   `json:"averagingValue,omitempty"`
	LinearScale       int     `json:"linearScale,omitempty"`
	SSBFilter         int     `json:"ssb,omitempty"`
	MarkersDisplay    int     `json:"markersDisplay,omitempty"`
	Measurement       int     `json:"measurement,omitempty"`
	MeasurementCenter int64   `json:"measurementCenterFrequency,omitempty"`
	MeasurementBW     int64   `json:"measurementBandwidth,omitempty"`
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

// WorkspaceMove specifies a target workspace index.
type WorkspaceMove struct {
	WorkspaceIndex int `json:"workspaceIndex"`
}

// SuccessResponse is a generic success message.
type SuccessResponse struct {
	Message string `json:"message,omitempty"`
}

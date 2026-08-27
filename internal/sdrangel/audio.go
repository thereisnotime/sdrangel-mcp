package sdrangel

import "context"

func (c *Client) GetAudioDevices(ctx context.Context) (AudioDevices, error) {
	return get[AudioDevices](ctx, c, "/sdrangel/audio")
}

func (c *Client) SetAudioInput(ctx context.Context, params AudioInputDevice) (AudioInputDevice, error) {
	return patchReq[AudioInputDevice](ctx, c, "/sdrangel/audio/input/parameters", params)
}

func (c *Client) SetAudioOutput(ctx context.Context, params AudioOutputDevice) (AudioOutputDevice, error) {
	return patchReq[AudioOutputDevice](ctx, c, "/sdrangel/audio/output/parameters", params)
}

func (c *Client) ResetAudioInput(ctx context.Context, name string) error {
	_, err := del[SuccessResponse](ctx, c, "/sdrangel/audio/input/parameters?device="+name, nil)
	return err
}

func (c *Client) ResetAudioOutput(ctx context.Context, name string) error {
	_, err := del[SuccessResponse](ctx, c, "/sdrangel/audio/output/parameters?device="+name, nil)
	return err
}

func (c *Client) CleanupAudioInputDevices(ctx context.Context) (SuccessResponse, error) {
	return patchReq[SuccessResponse](ctx, c, "/sdrangel/audio/input/cleanup", nil)
}

func (c *Client) CleanupAudioOutputDevices(ctx context.Context) (SuccessResponse, error) {
	return patchReq[SuccessResponse](ctx, c, "/sdrangel/audio/output/cleanup", nil)
}

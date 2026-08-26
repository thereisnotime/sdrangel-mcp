package sdrangel

import (
	"context"
	"fmt"
)

func (c *Client) SetDevice(ctx context.Context, deviceSetIndex int, link DeviceLink) (DeviceDesc, error) {
	return put[DeviceDesc](ctx, c, fmt.Sprintf("/sdrangel/deviceset/%d/device", deviceSetIndex), link)
}

func (c *Client) GetDeviceSettings(ctx context.Context, deviceSetIndex int) (DeviceSettings, error) {
	return get[DeviceSettings](ctx, c, fmt.Sprintf("/sdrangel/deviceset/%d/device/settings", deviceSetIndex))
}

func (c *Client) SetDeviceSettings(ctx context.Context, deviceSetIndex int, settings DeviceSettings) (DeviceSettings, error) {
	return put[DeviceSettings](ctx, c, fmt.Sprintf("/sdrangel/deviceset/%d/device/settings", deviceSetIndex), settings)
}

func (c *Client) PatchDeviceSettings(ctx context.Context, deviceSetIndex int, settings DeviceSettings) (DeviceSettings, error) {
	return patchReq[DeviceSettings](ctx, c, fmt.Sprintf("/sdrangel/deviceset/%d/device/settings", deviceSetIndex), settings)
}

func (c *Client) StartDevice(ctx context.Context, deviceSetIndex int) (DeviceState, error) {
	return post[DeviceState](ctx, c, fmt.Sprintf("/sdrangel/deviceset/%d/device/run", deviceSetIndex), nil)
}

func (c *Client) StopDevice(ctx context.Context, deviceSetIndex int) (DeviceState, error) {
	return del[DeviceState](ctx, c, fmt.Sprintf("/sdrangel/deviceset/%d/device/run", deviceSetIndex), nil)
}

func (c *Client) GetDeviceRunStatus(ctx context.Context, deviceSetIndex int) (DeviceState, error) {
	return get[DeviceState](ctx, c, fmt.Sprintf("/sdrangel/deviceset/%d/device/run", deviceSetIndex))
}

func (c *Client) GetDeviceReport(ctx context.Context, deviceSetIndex int) (DeviceReport, error) {
	return get[DeviceReport](ctx, c, fmt.Sprintf("/sdrangel/deviceset/%d/device/report", deviceSetIndex))
}

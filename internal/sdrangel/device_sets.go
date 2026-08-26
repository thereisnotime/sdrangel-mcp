package sdrangel

import (
	"context"
	"fmt"
)

func (c *Client) ListDeviceSets(ctx context.Context) (DeviceSets, error) {
	return get[DeviceSets](ctx, c, "/sdrangel/devicesets")
}

func (c *Client) GetDeviceSet(ctx context.Context, index int) (DeviceSetInfo, error) {
	return get[DeviceSetInfo](ctx, c, fmt.Sprintf("/sdrangel/deviceset/%d", index))
}

func (c *Client) AddDeviceSet(ctx context.Context, tx int) (SuccessResponse, error) {
	body := map[string]int{"tx": tx}
	return post[SuccessResponse](ctx, c, "/sdrangel/deviceset", body)
}

func (c *Client) RemoveDeviceSet(ctx context.Context) (SuccessResponse, error) {
	return del[SuccessResponse](ctx, c, "/sdrangel/deviceset", nil)
}

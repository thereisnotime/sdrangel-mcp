package sdrangel

import (
	"context"
	"net/http"
)

func (c *Client) GetInstanceSummary(ctx context.Context) (InstanceSummary, error) {
	return get[InstanceSummary](ctx, c, "/sdrangel")
}

func (c *Client) StopInstance(ctx context.Context) error {
	_, _, err := c.do(ctx, http.MethodDelete, "/sdrangel", nil)
	return err
}

func (c *Client) GetInstanceConfig(ctx context.Context) (InstanceConfig, error) {
	return get[InstanceConfig](ctx, c, "/sdrangel/config")
}

func (c *Client) SetInstanceConfig(ctx context.Context, cfg InstanceConfig) (InstanceConfig, error) {
	return put[InstanceConfig](ctx, c, "/sdrangel/config", cfg)
}

func (c *Client) PatchInstanceConfig(ctx context.Context, cfg InstanceConfig) (InstanceConfig, error) {
	return patchReq[InstanceConfig](ctx, c, "/sdrangel/config", cfg)
}

func (c *Client) ListDevicePlugins(ctx context.Context) (AvailableDeviceList, error) {
	return get[AvailableDeviceList](ctx, c, "/sdrangel/devices")
}

func (c *Client) ListChannelPlugins(ctx context.Context) (AvailableChannelOrFeatureList, error) {
	return get[AvailableChannelOrFeatureList](ctx, c, "/sdrangel/channels")
}

func (c *Client) ListFeaturePlugins(ctx context.Context) (AvailableChannelOrFeatureList, error) {
	return get[AvailableChannelOrFeatureList](ctx, c, "/sdrangel/features")
}

func (c *Client) GetLocation(ctx context.Context) (Location, error) {
	return get[Location](ctx, c, "/sdrangel/location")
}

func (c *Client) SetLocation(ctx context.Context, loc Location) (Location, error) {
	return put[Location](ctx, c, "/sdrangel/location", loc)
}

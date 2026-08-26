package sdrangel

import "context"

func (c *Client) ListConfigurations(ctx context.Context) (Configurations, error) {
	return get[Configurations](ctx, c, "/sdrangel/configurations")
}

func (c *Client) LoadConfiguration(ctx context.Context, keys ConfigurationKeys) (SuccessResponse, error) {
	return patchReq[SuccessResponse](ctx, c, "/sdrangel/configuration", keys)
}

func (c *Client) SaveConfiguration(ctx context.Context, keys ConfigurationKeys) (SuccessResponse, error) {
	return put[SuccessResponse](ctx, c, "/sdrangel/configuration", keys)
}

func (c *Client) CreateConfiguration(ctx context.Context, keys ConfigurationKeys) (SuccessResponse, error) {
	return post[SuccessResponse](ctx, c, "/sdrangel/configuration", keys)
}

func (c *Client) DeleteConfiguration(ctx context.Context, keys ConfigurationKeys) (ConfigurationKeys, error) {
	return del[ConfigurationKeys](ctx, c, "/sdrangel/configuration", keys)
}

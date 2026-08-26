package sdrangel

import "context"

func (c *Client) ListPresets(ctx context.Context) (Presets, error) {
	return get[Presets](ctx, c, "/sdrangel/presets")
}

func (c *Client) LoadPreset(ctx context.Context, transfer PresetTransfer) (SuccessResponse, error) {
	return patchReq[SuccessResponse](ctx, c, "/sdrangel/preset", transfer)
}

func (c *Client) SavePreset(ctx context.Context, transfer PresetTransfer) (SuccessResponse, error) {
	return put[SuccessResponse](ctx, c, "/sdrangel/preset", transfer)
}

func (c *Client) CreatePreset(ctx context.Context, transfer PresetTransfer) (SuccessResponse, error) {
	return post[SuccessResponse](ctx, c, "/sdrangel/preset", transfer)
}

func (c *Client) DeletePreset(ctx context.Context, keys PresetKeys) (PresetIdentifier, error) {
	return del[PresetIdentifier](ctx, c, "/sdrangel/preset", keys)
}

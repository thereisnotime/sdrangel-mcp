package sdrangel

import "context"

func (c *Client) ListPresets(ctx context.Context) (Presets, error) {
	return get[Presets](ctx, c, "/sdrangel/presets")
}

func (c *Client) LoadPreset(ctx context.Context, transfer PresetTransfer) (PresetIdentifier, error) {
	return patchReq[PresetIdentifier](ctx, c, "/sdrangel/preset", transfer)
}

func (c *Client) SavePreset(ctx context.Context, transfer PresetTransfer) (PresetIdentifier, error) {
	return put[PresetIdentifier](ctx, c, "/sdrangel/preset", transfer)
}

func (c *Client) CreatePreset(ctx context.Context, transfer PresetTransfer) (PresetIdentifier, error) {
	return post[PresetIdentifier](ctx, c, "/sdrangel/preset", transfer)
}

func (c *Client) DeletePreset(ctx context.Context, keys PresetKeys) (PresetIdentifier, error) {
	return del[PresetIdentifier](ctx, c, "/sdrangel/preset", keys)
}

// ImportPresetFromFile imports a preset from a file path on the server's
// filesystem, creating a new preset.
func (c *Client) ImportPresetFromFile(ctx context.Context, path FilePath) (PresetIdentifier, error) {
	return put[PresetIdentifier](ctx, c, "/sdrangel/preset/file", path)
}

// ExportPresetToFile exports an existing preset to a file path on the
// server's filesystem.
func (c *Client) ExportPresetToFile(ctx context.Context, export PresetExport) (PresetIdentifier, error) {
	return post[PresetIdentifier](ctx, c, "/sdrangel/preset/file", export)
}

// ImportPresetFromBlob deserializes a base64-encoded blob into a new preset.
func (c *Client) ImportPresetFromBlob(ctx context.Context, blob Base64Blob) (PresetIdentifier, error) {
	return put[PresetIdentifier](ctx, c, "/sdrangel/preset/blob", blob)
}

// ExportPresetToBlob serializes an existing preset to a base64-encoded blob.
func (c *Client) ExportPresetToBlob(ctx context.Context, id PresetIdentifier) (Base64Blob, error) {
	return post[Base64Blob](ctx, c, "/sdrangel/preset/blob", id)
}

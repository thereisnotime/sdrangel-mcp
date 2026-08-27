package sdrangel

import "context"

func (c *Client) ListConfigurations(ctx context.Context) (Configurations, error) {
	return get[Configurations](ctx, c, "/sdrangel/configurations")
}

func (c *Client) LoadConfiguration(ctx context.Context, keys ConfigurationKeys) (ConfigurationKeys, error) {
	return patchReq[ConfigurationKeys](ctx, c, "/sdrangel/configuration", keys)
}

func (c *Client) SaveConfiguration(ctx context.Context, keys ConfigurationKeys) (ConfigurationKeys, error) {
	return put[ConfigurationKeys](ctx, c, "/sdrangel/configuration", keys)
}

func (c *Client) CreateConfiguration(ctx context.Context, keys ConfigurationKeys) (ConfigurationKeys, error) {
	return post[ConfigurationKeys](ctx, c, "/sdrangel/configuration", keys)
}

func (c *Client) DeleteConfiguration(ctx context.Context, keys ConfigurationKeys) (ConfigurationKeys, error) {
	return del[ConfigurationKeys](ctx, c, "/sdrangel/configuration", keys)
}

// ImportConfigurationFromFile imports a configuration from a file path on
// the server's filesystem, creating a new configuration.
func (c *Client) ImportConfigurationFromFile(ctx context.Context, path FilePath) (ConfigurationKeys, error) {
	return put[ConfigurationKeys](ctx, c, "/sdrangel/configuration/file", path)
}

// ExportConfigurationToFile exports an existing configuration to a file path
// on the server's filesystem.
func (c *Client) ExportConfigurationToFile(ctx context.Context, export ConfigurationImportExport) (ConfigurationKeys, error) {
	return post[ConfigurationKeys](ctx, c, "/sdrangel/configuration/file", export)
}

// ImportConfigurationFromBlob deserializes a base64-encoded blob into a new
// configuration.
func (c *Client) ImportConfigurationFromBlob(ctx context.Context, blob Base64Blob) (ConfigurationKeys, error) {
	return put[ConfigurationKeys](ctx, c, "/sdrangel/configuration/blob", blob)
}

// ExportConfigurationToBlob serializes an existing configuration to a
// base64-encoded blob.
func (c *Client) ExportConfigurationToBlob(ctx context.Context, keys ConfigurationKeys) (Base64Blob, error) {
	return post[Base64Blob](ctx, c, "/sdrangel/configuration/blob", keys)
}

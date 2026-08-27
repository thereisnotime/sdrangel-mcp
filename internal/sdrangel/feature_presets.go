package sdrangel

import "context"

// ListFeaturePresets lists all feature presets in the instance. Feature
// presets are a top-level catalog distinct from the per-feature-set preset
// resource (see LoadFeatureSetPreset/SaveFeatureSetPreset/CreateFeatureSetPreset).
func (c *Client) ListFeaturePresets(ctx context.Context) (FeaturePresets, error) {
	return get[FeaturePresets](ctx, c, "/sdrangel/featurepresets")
}

// DeleteFeaturePreset deletes a feature preset identified by groupName and description.
func (c *Client) DeleteFeaturePreset(ctx context.Context, id FeaturePresetIdentifier) (FeaturePresetIdentifier, error) {
	return del[FeaturePresetIdentifier](ctx, c, "/sdrangel/featurepreset", id)
}

// LoadFeatureSetPreset loads a preset into the current feature set.
func (c *Client) LoadFeatureSetPreset(ctx context.Context, id FeaturePresetIdentifier) (FeaturePresetIdentifier, error) {
	return patchReq[FeaturePresetIdentifier](ctx, c, "/sdrangel/featureset/preset", id)
}

// SaveFeatureSetPreset saves the current feature set state into an existing preset.
func (c *Client) SaveFeatureSetPreset(ctx context.Context, id FeaturePresetIdentifier) (FeaturePresetIdentifier, error) {
	return put[FeaturePresetIdentifier](ctx, c, "/sdrangel/featureset/preset", id)
}

// CreateFeatureSetPreset creates a new preset from the current feature set state.
//
// NOTE: swagger.yaml literally declares this operation's (POST) response
// schema as PresetIdentifier (groupName/name/type/centerFrequency) while its
// sibling PATCH and PUT operations on the same path both declare
// FeaturePresetIdentifier (groupName/description), and the request body for
// all three verbs is FeaturePresetIdentifier. That PresetIdentifier response
// schema is identical, field-for-field, to the POST response schema of the
// unrelated /sdrangel/preset endpoint, strongly suggesting it was copied
// there by mistake rather than being intentional. FeaturePresetIdentifier is
// used here for consistency with the rest of this resource.
func (c *Client) CreateFeatureSetPreset(ctx context.Context, id FeaturePresetIdentifier) (FeaturePresetIdentifier, error) {
	return post[FeaturePresetIdentifier](ctx, c, "/sdrangel/featureset/preset", id)
}

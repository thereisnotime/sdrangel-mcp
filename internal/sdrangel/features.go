package sdrangel

import (
	"context"
	"fmt"
)

func (c *Client) GetFeatureSet(ctx context.Context) (FeatureSetInfo, error) {
	return get[FeatureSetInfo](ctx, c, "/sdrangel/featureset")
}

func (c *Client) AddFeature(ctx context.Context, f FeatureAdd) (SuccessResponse, error) {
	return post[SuccessResponse](ctx, c, "/sdrangel/featureset/feature", f)
}

func (c *Client) DeleteFeature(ctx context.Context, featureIndex int) (SuccessResponse, error) {
	return del[SuccessResponse](ctx, c, fmt.Sprintf("/sdrangel/featureset/feature/%d", featureIndex), nil)
}

func (c *Client) GetFeatureSettings(ctx context.Context, featureIndex int) (FeatureSettings, error) {
	return get[FeatureSettings](ctx, c, fmt.Sprintf("/sdrangel/featureset/feature/%d/settings", featureIndex))
}

func (c *Client) SetFeatureSettings(ctx context.Context, featureIndex int, settings FeatureSettings) (FeatureSettings, error) {
	return put[FeatureSettings](ctx, c, fmt.Sprintf("/sdrangel/featureset/feature/%d/settings", featureIndex), settings)
}

func (c *Client) PatchFeatureSettings(ctx context.Context, featureIndex int, settings FeatureSettings) (FeatureSettings, error) {
	return patchReq[FeatureSettings](ctx, c, fmt.Sprintf("/sdrangel/featureset/feature/%d/settings", featureIndex), settings)
}

func (c *Client) StartFeature(ctx context.Context, featureIndex int) (FeatureState, error) {
	return post[FeatureState](ctx, c, fmt.Sprintf("/sdrangel/featureset/feature/%d/run", featureIndex), nil)
}

func (c *Client) StopFeature(ctx context.Context, featureIndex int) (FeatureState, error) {
	return del[FeatureState](ctx, c, fmt.Sprintf("/sdrangel/featureset/feature/%d/run", featureIndex), nil)
}

func (c *Client) GetFeatureRunStatus(ctx context.Context, featureIndex int) (FeatureState, error) {
	return get[FeatureState](ctx, c, fmt.Sprintf("/sdrangel/featureset/feature/%d/run", featureIndex))
}

func (c *Client) GetFeatureReport(ctx context.Context, featureIndex int) (FeatureReport, error) {
	return get[FeatureReport](ctx, c, fmt.Sprintf("/sdrangel/featureset/feature/%d/report", featureIndex))
}

func (c *Client) ExecuteFeatureActions(ctx context.Context, featureIndex int, actions FeatureActions) (SuccessResponse, error) {
	return post[SuccessResponse](ctx, c, fmt.Sprintf("/sdrangel/featureset/feature/%d/actions", featureIndex), actions)
}

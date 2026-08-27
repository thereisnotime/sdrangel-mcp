package sdrangel

import (
	"context"
	"fmt"
)

func (c *Client) AddWorkspace(ctx context.Context) (WorkspaceInfo, error) {
	return post[WorkspaceInfo](ctx, c, "/sdrangel/workspace", nil)
}

func (c *Client) DeleteWorkspace(ctx context.Context) (SuccessResponse, error) {
	return del[SuccessResponse](ctx, c, "/sdrangel/workspace", nil)
}

func (c *Client) GetDeviceWorkspace(ctx context.Context, deviceSetIndex int) (WorkspaceInfo, error) {
	return get[WorkspaceInfo](ctx, c, fmt.Sprintf("/sdrangel/deviceset/%d/device/workspace", deviceSetIndex))
}

func (c *Client) MoveDeviceToWorkspace(ctx context.Context, deviceSetIndex int, move WorkspaceInfo) (SuccessResponse, error) {
	return put[SuccessResponse](ctx, c, fmt.Sprintf("/sdrangel/deviceset/%d/device/workspace", deviceSetIndex), move)
}

func (c *Client) GetChannelWorkspace(ctx context.Context, deviceSetIndex, channelIndex int) (WorkspaceInfo, error) {
	return get[WorkspaceInfo](ctx, c, fmt.Sprintf("/sdrangel/deviceset/%d/channel/%d/workspace", deviceSetIndex, channelIndex))
}

func (c *Client) MoveChannelToWorkspace(ctx context.Context, deviceSetIndex, channelIndex int, move WorkspaceInfo) (SuccessResponse, error) {
	return put[SuccessResponse](ctx, c, fmt.Sprintf("/sdrangel/deviceset/%d/channel/%d/workspace", deviceSetIndex, channelIndex), move)
}

func (c *Client) GetFeatureWorkspace(ctx context.Context, featureIndex int) (WorkspaceInfo, error) {
	return get[WorkspaceInfo](ctx, c, fmt.Sprintf("/sdrangel/featureset/feature/%d/workspace", featureIndex))
}

func (c *Client) MoveFeatureToWorkspace(ctx context.Context, featureIndex int, move WorkspaceInfo) (SuccessResponse, error) {
	return put[SuccessResponse](ctx, c, fmt.Sprintf("/sdrangel/featureset/feature/%d/workspace", featureIndex), move)
}

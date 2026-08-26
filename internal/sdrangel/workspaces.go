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

func (c *Client) MoveDeviceToWorkspace(ctx context.Context, deviceSetIndex int, move WorkspaceMove) (SuccessResponse, error) {
	return put[SuccessResponse](ctx, c, fmt.Sprintf("/sdrangel/deviceset/%d/device/workspace", deviceSetIndex), move)
}

func (c *Client) MoveChannelToWorkspace(ctx context.Context, deviceSetIndex, channelIndex int, move WorkspaceMove) (SuccessResponse, error) {
	return put[SuccessResponse](ctx, c, fmt.Sprintf("/sdrangel/deviceset/%d/channel/%d/workspace", deviceSetIndex, channelIndex), move)
}

func (c *Client) MoveFeatureToWorkspace(ctx context.Context, featureIndex int, move WorkspaceMove) (SuccessResponse, error) {
	return put[SuccessResponse](ctx, c, fmt.Sprintf("/sdrangel/featureset/feature/%d/workspace", featureIndex), move)
}

package sdrangel

import (
	"context"
	"fmt"
)

func (c *Client) GetSpectrumSettings(ctx context.Context, deviceSetIndex int) (SpectrumSettings, error) {
	return get[SpectrumSettings](ctx, c, fmt.Sprintf("/sdrangel/deviceset/%d/spectrum/settings", deviceSetIndex))
}

func (c *Client) SetSpectrumSettings(ctx context.Context, deviceSetIndex int, settings SpectrumSettings) (SpectrumSettings, error) {
	return put[SpectrumSettings](ctx, c, fmt.Sprintf("/sdrangel/deviceset/%d/spectrum/settings", deviceSetIndex), settings)
}

func (c *Client) PatchSpectrumSettings(ctx context.Context, deviceSetIndex int, settings SpectrumSettings) (SpectrumSettings, error) {
	return patchReq[SpectrumSettings](ctx, c, fmt.Sprintf("/sdrangel/deviceset/%d/spectrum/settings", deviceSetIndex), settings)
}

func (c *Client) StartSpectrumServer(ctx context.Context, deviceSetIndex int) (SpectrumServer, error) {
	return post[SpectrumServer](ctx, c, fmt.Sprintf("/sdrangel/deviceset/%d/spectrum/server", deviceSetIndex), nil)
}

func (c *Client) StopSpectrumServer(ctx context.Context, deviceSetIndex int) (SpectrumServer, error) {
	return del[SpectrumServer](ctx, c, fmt.Sprintf("/sdrangel/deviceset/%d/spectrum/server", deviceSetIndex), nil)
}

func (c *Client) GetSpectrumServerStatus(ctx context.Context, deviceSetIndex int) (SpectrumServer, error) {
	return get[SpectrumServer](ctx, c, fmt.Sprintf("/sdrangel/deviceset/%d/spectrum/server", deviceSetIndex))
}

func (c *Client) GetSpectrumWorkspace(ctx context.Context, deviceSetIndex int) (WorkspaceInfo, error) {
	return get[WorkspaceInfo](ctx, c, fmt.Sprintf("/sdrangel/deviceset/%d/spectrum/workspace", deviceSetIndex))
}

func (c *Client) MoveSpectrumToWorkspace(ctx context.Context, deviceSetIndex int, move WorkspaceInfo) (SuccessResponse, error) {
	return put[SuccessResponse](ctx, c, fmt.Sprintf("/sdrangel/deviceset/%d/spectrum/workspace", deviceSetIndex), move)
}

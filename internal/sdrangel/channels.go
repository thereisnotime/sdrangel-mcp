package sdrangel

import (
	"context"
	"fmt"
)

func (c *Client) AddChannel(ctx context.Context, deviceSetIndex int, ch ChannelAdd) (SuccessResponse, error) {
	return post[SuccessResponse](ctx, c, fmt.Sprintf("/sdrangel/deviceset/%d/channel", deviceSetIndex), ch)
}

func (c *Client) DeleteChannel(ctx context.Context, deviceSetIndex, channelIndex int) (SuccessResponse, error) {
	return del[SuccessResponse](ctx, c, fmt.Sprintf("/sdrangel/deviceset/%d/channel/%d", deviceSetIndex, channelIndex), nil)
}

func (c *Client) GetChannelSettings(ctx context.Context, deviceSetIndex, channelIndex int) (ChannelSettings, error) {
	return get[ChannelSettings](ctx, c, fmt.Sprintf("/sdrangel/deviceset/%d/channel/%d/settings", deviceSetIndex, channelIndex))
}

func (c *Client) SetChannelSettings(ctx context.Context, deviceSetIndex, channelIndex int, settings ChannelSettings) (ChannelSettings, error) {
	return put[ChannelSettings](ctx, c, fmt.Sprintf("/sdrangel/deviceset/%d/channel/%d/settings", deviceSetIndex, channelIndex), settings)
}

func (c *Client) PatchChannelSettings(ctx context.Context, deviceSetIndex, channelIndex int, settings ChannelSettings) (ChannelSettings, error) {
	return patchReq[ChannelSettings](ctx, c, fmt.Sprintf("/sdrangel/deviceset/%d/channel/%d/settings", deviceSetIndex, channelIndex), settings)
}

func (c *Client) GetChannelReport(ctx context.Context, deviceSetIndex, channelIndex int) (ChannelReport, error) {
	return get[ChannelReport](ctx, c, fmt.Sprintf("/sdrangel/deviceset/%d/channel/%d/report", deviceSetIndex, channelIndex))
}

func (c *Client) ExecuteChannelActions(ctx context.Context, deviceSetIndex, channelIndex int, actions ChannelActions) (SuccessResponse, error) {
	return post[SuccessResponse](ctx, c, fmt.Sprintf("/sdrangel/deviceset/%d/channel/%d/actions", deviceSetIndex, channelIndex), actions)
}

func (c *Client) GetChannelsReport(ctx context.Context, deviceSetIndex int) (ChannelsReport, error) {
	return get[ChannelsReport](ctx, c, fmt.Sprintf("/sdrangel/deviceset/%d/channels/report", deviceSetIndex))
}

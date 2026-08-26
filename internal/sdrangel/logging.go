package sdrangel

import "context"

func (c *Client) GetLogging(ctx context.Context) (LoggingInfo, error) {
	return get[LoggingInfo](ctx, c, "/sdrangel/logging")
}

func (c *Client) SetLogging(ctx context.Context, info LoggingInfo) (LoggingInfo, error) {
	return put[LoggingInfo](ctx, c, "/sdrangel/logging", info)
}

package sdrangel

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
	username   string
	password   string
}

type Options struct {
	BaseURL  string
	Timeout  time.Duration
	Username string
	Password string
}

func New(opts Options) *Client {
	if opts.BaseURL == "" {
		opts.BaseURL = "http://localhost:8091"
	}
	if opts.Timeout == 0 {
		opts.Timeout = 10 * time.Second
	}
	return &Client{
		baseURL:    opts.BaseURL,
		httpClient: &http.Client{Timeout: opts.Timeout},
		username:   opts.Username,
		password:   opts.Password,
	}
}

func (c *Client) do(ctx context.Context, method, path string, body any) ([]byte, int, error) {
	var bodyReader io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, 0, fmt.Errorf("marshal request: %w", err)
		}
		bodyReader = bytes.NewReader(b)
	}
	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+path, bodyReader)
	if err != nil {
		return nil, 0, fmt.Errorf("create request: %w", err)
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	if c.username != "" {
		req.SetBasicAuth(c.username, c.password)
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, 0, fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf("read response: %w", err)
	}
	if resp.StatusCode >= 400 {
		return nil, resp.StatusCode, fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(data))
	}
	return data, resp.StatusCode, nil
}

func get[T any](ctx context.Context, c *Client, path string) (T, error) {
	var zero T
	data, _, err := c.do(ctx, http.MethodGet, path, nil)
	if err != nil {
		return zero, err
	}
	var result T
	if err := json.Unmarshal(data, &result); err != nil {
		return zero, fmt.Errorf("unmarshal: %w", err)
	}
	return result, nil
}

func post[T any](ctx context.Context, c *Client, path string, body any) (T, error) {
	var zero T
	data, _, err := c.do(ctx, http.MethodPost, path, body)
	if err != nil {
		return zero, err
	}
	if len(data) == 0 {
		return zero, nil
	}
	var result T
	if err := json.Unmarshal(data, &result); err != nil {
		return zero, fmt.Errorf("unmarshal: %w", err)
	}
	return result, nil
}

func put[T any](ctx context.Context, c *Client, path string, body any) (T, error) {
	var zero T
	data, _, err := c.do(ctx, http.MethodPut, path, body)
	if err != nil {
		return zero, err
	}
	if len(data) == 0 {
		return zero, nil
	}
	var result T
	if err := json.Unmarshal(data, &result); err != nil {
		return zero, fmt.Errorf("unmarshal: %w", err)
	}
	return result, nil
}

func patchReq[T any](ctx context.Context, c *Client, path string, body any) (T, error) {
	var zero T
	data, _, err := c.do(ctx, http.MethodPatch, path, body)
	if err != nil {
		return zero, err
	}
	if len(data) == 0 {
		return zero, nil
	}
	var result T
	if err := json.Unmarshal(data, &result); err != nil {
		return zero, fmt.Errorf("unmarshal: %w", err)
	}
	return result, nil
}

func del[T any](ctx context.Context, c *Client, path string, body any) (T, error) {
	var zero T
	data, _, err := c.do(ctx, http.MethodDelete, path, body)
	if err != nil {
		return zero, err
	}
	if len(data) == 0 {
		return zero, nil
	}
	var result T
	if err := json.Unmarshal(data, &result); err != nil {
		return zero, fmt.Errorf("unmarshal: %w", err)
	}
	return result, nil
}

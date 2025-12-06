package model

import (
	"bytes"
	"context"
	"net/http"
	"time"
)

type Driver struct {
	baseURL string
	client  *http.Client
}

func NewDriverClient(url string) *Driver {
	return &Driver{
		baseURL: url,
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (d *Driver) ForwardPost(ctx context.Context, path string, body []byte) (*http.Response, error) {
	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, d.baseURL+path, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	return d.client.Do(req)
}

func (d *Driver) ForwardGet(ctx context.Context, path string) (*http.Response, error) {
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, d.baseURL+path, nil)
	return d.client.Do(req)
}

func (d *Driver) ForwardPut(ctx context.Context, path string, body []byte) (*http.Response, error) {
	req, _ := http.NewRequestWithContext(ctx, http.MethodPut, d.baseURL+path, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	return d.client.Do(req)
}

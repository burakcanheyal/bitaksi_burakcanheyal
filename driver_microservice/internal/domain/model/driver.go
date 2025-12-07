package model

import (
	"bytes"
	"context"
	"net/http"
	"time"
)

const InternalAPIKey = "BITAKSI-DB-ACCESS-KEY-5555"

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

	req.Header.Set("X-INTERNAL-KEY", InternalAPIKey)

	return d.client.Do(req)
}

func (d *Driver) ForwardGet(ctx context.Context, path string) (*http.Response, error) {

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, d.baseURL+path, nil)

	// ⭐ Internal API Key ekleme
	req.Header.Set("X-INTERNAL-KEY", InternalAPIKey)

	return d.client.Do(req)
}

func (d *Driver) ForwardPut(ctx context.Context, path string, body []byte) (*http.Response, error) {

	req, _ := http.NewRequestWithContext(ctx, http.MethodPut, d.baseURL+path, bytes.NewReader(body))

	req.Header.Set("Content-Type", "application/json")

	// ⭐ Internal API Key ekleme
	req.Header.Set("X-INTERNAL-KEY", InternalAPIKey)

	return d.client.Do(req)
}

package gtfs

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type feed struct {
	agency   string
	category string
}

type FeedClient struct {
	baseURL string
	client  *http.Client
}

func (c *FeedClient) DownloadFeed(ctx context.Context, f feed) ([]byte, error) {
	p, err := url.JoinPath(c.baseURL, f.agency)
	if err != nil {
		return nil, fmt.Errorf("join path: %w", err)
	}

	u, err := url.Parse(p)
	if err != nil {
		return nil, fmt.Errorf("parse baseURL: %w", err)
	}

	if f.category != "" {
		q := u.Query()
		q.Set("category", f.category)

		u.RawQuery = q.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("construct new http request: %w", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("GET '%s': %w", u, err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET '%s': failed with status code %d", u, resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}

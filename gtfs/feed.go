package gtfs

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Feed struct {
	Agency   string
	Category string
}

type FeedClient struct {
	baseURL string
	client  *http.Client
}

func (c *FeedClient) DownloadFeed(ctx context.Context, f Feed) ([]byte, error) {
	p, err := url.JoinPath(c.baseURL, f.Agency)
	if err != nil {
		return nil, fmt.Errorf("join path: %w", err)
	}

	u, err := url.Parse(p)
	if err != nil {
		return nil, fmt.Errorf("parse baseURL: %w", err)
	}

	if f.Category != "" {
		q := u.Query()
		q.Set("category", f.Category)

		u.RawQuery = q.Encode()
	}

	downloadCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(downloadCtx, http.MethodGet, u.String(), nil)
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

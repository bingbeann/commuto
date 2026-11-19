//go:build integration

package gtfs

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDownloadFeed_ReturnsSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/prasarana", r.URL.Path)
		assert.Equal(t, "rapid-bus-kl", r.URL.Query().Get("category"))

		w.WriteHeader(http.StatusOK)
		_, err := fmt.Fprintf(w, "UEsFBgAAAAAAAAAAAAAAAAAAAAAAAA==")
		require.NoError(t, err)
	}))

	c := &FeedClient{
		baseURL: server.URL,
		client:  &http.Client{},
	}

	b, err := c.DownloadFeed(t.Context(), Feed{"prasarana", "rapid-bus-kl"})

	require.NoError(t, err)
	assert.NotEmpty(t, b)
}

func TestDownloadFeed_ReturnsErr(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))

	c := &FeedClient{
		baseURL: server.URL,
		client:  &http.Client{},
	}

	_, err := c.DownloadFeed(t.Context(), Feed{"prasarana", "rapid-bus-kl"})

	require.Error(t, err)
}

func TestDownloadFeed_WithBadURL_ReturnsErr(t *testing.T) {
	c := &FeedClient{
		baseURL: "htp:/domain.bad",
		client:  &http.Client{},
	}

	_, err := c.DownloadFeed(t.Context(), Feed{"prasarana", "rapid-bus-kl"})

	require.Error(t, err)
}

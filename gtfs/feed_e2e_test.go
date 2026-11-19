//go:build e2e

package gtfs

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDownloadFeed_ToDataGovMy_ReturnsSuccess(t *testing.T) {
	c := &FeedClient{
		baseURL: "https://api.data.gov.my/gtfs-static",
		client:  &http.Client{},
	}

	b, err := c.DownloadFeed(t.Context(), Feed{"prasarana", "rapid-bus-kl"})

	require.NoError(t, err)
	assert.NotEmpty(t, b)
}

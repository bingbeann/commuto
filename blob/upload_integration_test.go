//go:build integration

package blob

import (
	"os"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/transfermanager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/stretchr/testify/require"
)

func TestUpload(t *testing.T) {
	s3Client := newS3Client(t)
	uploader := transfermanager.New(s3Client)
	waiter := s3.NewObjectExistsWaiter(s3Client)

	s := &Service{
		Uploader: uploader,
		Waiter:   waiter,
	}

	err := s.Upload(t.Context(), "gtfs", "tmp", strings.NewReader("this is just for testing"))

	require.NoError(t, err)
}

func newS3Client(t *testing.T) *s3.Client {
	envs := map[string]string{
		"AWS_ENDPOINT_URL_S3":   "http://localhost:8333",
		"AWS_REGION":            "ap-southeast-3",
		"AWS_ACCESS_KEY_ID":     "key",
		"AWS_SECRET_ACCESS_KEY": "secret",
	}
	for k, v := range envs {
		err := os.Setenv(k, v)
		if err != nil {
			t.Fatalf("failed to set env %s: %v", k, err)
		}
	}

	sdkConfig, err := config.LoadDefaultConfig(t.Context())
	require.NoError(t, err)
	client := s3.NewFromConfig(sdkConfig, func(o *s3.Options) {
		o.UsePathStyle = true
		o.RequestChecksumCalculation = aws.RequestChecksumCalculationWhenRequired
	})

	return client
}

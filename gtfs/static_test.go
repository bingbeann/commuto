package gtfs

import (
	"log"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	envs := map[string]string{
		"AWS_ENDPOINT_URL_S3": "http://localhost:8333",
		"AWS_REGION": "ap-southeast-3",
		"AWS_ACCESS_KEY_ID": "key",
		"AWS_SECRET_ACCESS_KEY": "secret",
	}
	for k, v := range envs {
		err := os.Setenv(k, v)
		if err != nil {
			log.Fatalf("failed to set env %s: %v", k, err)
		}
	}

	m.Run()
}

func TestGetStatic(t *testing.T) {
	ctx := t.Context()
	bucket := "gtfs"
	sdkConfig, err := config.LoadDefaultConfig(ctx)
	require.NoError(t, err)
	client := s3.NewFromConfig(sdkConfig, func(o *s3.Options) {
		o.UsePathStyle = true
		o.RequestChecksumCalculation = aws.RequestChecksumCalculationWhenRequired
	})

	n, err := GetStatic(ctx, client, bucket)
	
	require.NoError(t, err)
	assert.Equal(t, 5, n)
}


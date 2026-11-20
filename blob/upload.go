package blob

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/s3/transfermanager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Uploader interface {
	UploadObject(ctx context.Context, input *transfermanager.UploadObjectInput, opts ...func(*transfermanager.Options)) (*transfermanager.UploadObjectOutput, error)
}

type NewObjectExistsWaiter interface {
	Wait(ctx context.Context, params *s3.HeadObjectInput, maxWaitDur time.Duration, optFns ...func(*s3.ObjectExistsWaiterOptions)) error
}

type Service struct {
	Uploader Uploader
	Waiter NewObjectExistsWaiter
}

func (s *Service) Upload(ctx context.Context, bucket string, key string, r io.Reader) error {
	_, err := s.Uploader.UploadObject(ctx, &transfermanager.UploadObjectInput{
		Bucket: new(bucket),
		Key: new(key),
		Body: r,
	})
	if err != nil {
		return fmt.Errorf("upload object '%s' to s3 bucket '%s': %w", key, bucket, err)
	}

	err = s.Waiter.Wait(
		ctx,
		&s3.HeadObjectInput{Bucket: new(bucket), Key: new(key)},
		time.Minute,
	)
	if err != nil {
		return fmt.Errorf("wait for s3 object '%s' to exist: %w", key, err)
	}

	return nil
}

package gtfs

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/s3/transfermanager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type gtfsFeed struct {
	agency string
	category string
}

func GetStatic(ctx context.Context, client *s3.Client, bucket string) (int, error) {
	feeds := []gtfsFeed{
		{"ktmb", ""},
		{"prasarana", "rapid-rail-kl"},
		{"prasarana", "rapid-bus-kl"},
		{"prasarana", "rapid-bus-penang"},
		{"prasarana", "rapid-bus-mrtfeeder"},
	}

	for n, feed := range feeds {
		u := "https://api.data.gov.my/gtfs-static/" + feed.agency + "?category=" + feed.category
	
		resp, err := http.Get(u)
		if err != nil {
			return n, fmt.Errorf("GET '%s': %w", u, err)
		}
		defer func() {
			_ = resp.Body.Close()
		}()

		if resp.StatusCode >= 400 {
			return n, fmt.Errorf("GET '%s': failed with status code %d", u, resp.StatusCode)
		}

		key := fmt.Sprintf("%s_%s.gtfs.zip", feed.agency, feed.category)
	
		uploader := transfermanager.New(client)
		_, err = uploader.UploadObject(ctx, &transfermanager.UploadObjectInput{
			Bucket: new(bucket),
			Key: new(key),
			Body: resp.Body,
		})
		if err != nil {
			return n, fmt.Errorf("s3 put object '%s': %w", key, err)
		}

		err = s3.NewObjectExistsWaiter(client).Wait(
			ctx, 
			&s3.HeadObjectInput{
				Bucket: new(bucket),
				Key: new(key),
			},
			time.Minute,
		)
		if err != nil {
			return n, fmt.Errorf("s3 wait for object '%s' to exist: %w", key, err)
		}
	}

	return len(feeds), nil
}

package s3

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/rs/zerolog/log"
)

// GetObject makes a presigned request that can be used to get an object from a bucket.
// The presigned request is valid for the specified number of seconds.
func (presigner BucketBasics) GetObject(
	ctx context.Context, bucketName string, objectKey string, lifetimeSecs int64,
) (*v4.PresignedHTTPRequest, error) {
	request, err := presigner.PresignClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	}, func(opts *s3.PresignOptions) {
		opts.Expires = time.Duration(lifetimeSecs * int64(time.Second))
	})
	if err != nil {
		log.Info().Msgf("Couldn't get a presigned request to get %v:%v. Here's why: %v\n",
			bucketName, objectKey, err)
	}
	return request, err
}

package s3

import (
	"context"
	"snaptrail/internal/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/rs/zerolog/log"
)

var client *s3.Client

func NewS3ClientFromEnv() {
	conf := config.Get()
	endpoint := conf.S3Endpoint
	// pathStyle := conf.S3PathStyle
	pathStyle := true
	region := conf.S3Region
	accessKey := conf.S3AccessKey
	secretKey := conf.S3SecretKey
	bucketName := conf.S3Bucket

	client = s3.New(s3.Options{
		AppID: "my-application/0.0.1",

		Region:       region,
		BaseEndpoint: aws.String(endpoint),
		UsePathStyle: pathStyle,

		Credentials: credentials.StaticCredentialsProvider{Value: aws.Credentials{
			AccessKeyID:     accessKey,
			SecretAccessKey: secretKey,
		}},
	})

	ctx := context.Background()

	bb := BucketBasics{S3Client: client}

	exists, err := bb.BucketExists(ctx, bucketName)
	if err != nil {
		log.Error().Err(err).Msg("unable to check if bucket exists")
		return
	}
	if !exists {
		err := bb.CreateBucket(ctx, bucketName, region)
		if err != nil {
			log.Fatal().Err(err).Msg("could not create bucket for images")
		}
	}
}

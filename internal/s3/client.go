package s3

import (
	"context"
	"snaptrail/internal/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/rs/zerolog/log"
)

var (
	client        *s3.Client
	presignClient *s3.PresignClient
)

type BucketBasics struct {
	Client        *s3.Client
	PresignClient *s3.PresignClient
}

func NewBucketBasics() BucketBasics {
	return BucketBasics{
		Client:        client,
		PresignClient: presignClient,
	}
}

func NewS3ClientFromEnv() {
	conf := config.Get()
	endpoint := conf.S3Endpoint
	publicHost := "https://" + conf.S3PublicHost
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

	externalClient := s3.New(s3.Options{
		AppID: "my-application/0.0.1",

		Region:       region,
		BaseEndpoint: aws.String(publicHost),
		UsePathStyle: pathStyle,

		Credentials: credentials.StaticCredentialsProvider{Value: aws.Credentials{
			AccessKeyID:     accessKey,
			SecretAccessKey: secretKey,
		}},
	})
	presignClient = s3.NewPresignClient(externalClient)

	ctx := context.Background()

	bb := BucketBasics{Client: client, PresignClient: presignClient}
	log.Info().Msgf("Created the following bucket basics: %v", bb)

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

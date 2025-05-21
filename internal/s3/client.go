package s3

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"snaptrail/internal/config"
	"strings"

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

	testBucket()
}

func testBucket() {
	ctx := context.Background()
	VAR_BUCKET_NAME := fmt.Sprintf("my-bucket-%s", generateRandomHash(8))
	// Create a new Bucket
	_, err := client.CreateBucket(ctx, &s3.CreateBucketInput{
		Bucket: aws.String(VAR_BUCKET_NAME),
	})
	if err != nil {
		log.Err(err).Msg("failed creating bucket")
	}

	log.Printf("Created Bucket: %s", VAR_BUCKET_NAME)

	// Upload a small object
	body := strings.NewReader(strings.TrimSpace(`
# My object

This markdown document is the content of my object.
`))

	_, err = client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(VAR_BUCKET_NAME),
		Key:         aws.String("my-object-e9be8f5c"),
		ContentType: aws.String("text/markdown"),

		Body: body,
	})
	if err != nil {
		log.Err(err).Msg("error putting object")
	}

	log.Info().Msg("Created object")
}

func GetS3() *s3.Client {
	return client
}

func generateRandomHash(length int) string {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		log.Err(err).Msg("failed generating 8 digit hash")
	}
	return hex.EncodeToString(bytes)[:length]
}

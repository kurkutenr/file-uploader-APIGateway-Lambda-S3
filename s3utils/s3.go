package s3utils

import (
	"bytes"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Set this parameters from lambda environment OS.env variables
var apiKey = "<>"
var apiSecretKey = "<>"
var zone = "<>"

// Upload - upload file
func Upload(bucket string, key string, body []byte) (err error) {

	s, err := session.NewSession(&aws.Config{
		Region:      aws.String(zone),
		Credentials: credentials.NewStaticCredentials(apiKey, apiSecretKey, ""),
	})

	s3c := s3.New(s)

	if err != nil {
		panic(err)
	}

	_, err = s3c.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(body),
	})

	return err
}

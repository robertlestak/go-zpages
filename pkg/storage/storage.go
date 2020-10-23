package storage

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

// S3Checker wraps a service connection to S3
type S3Checker struct {
	SVC        *s3.S3
	BucketName *string
}

// Ping checks to see if S3 bucket exists
func (sc *S3Checker) Ping() error {
	input := &s3.HeadBucketInput{Bucket: aws.String(*sc.BucketName)}
	_, err := sc.SVC.HeadBucket(input)
	if err != nil {
		return err
	}
	return nil
}

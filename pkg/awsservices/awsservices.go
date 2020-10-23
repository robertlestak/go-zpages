package awsservices

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rekognition"
)

// RekognitionChecker wraps a service connection to Rekognition
type RekognitionChecker struct {
	SVC *rekognition.Rekognition
}

// Ping checks to see if Rekognition API responds
func (sc *RekognitionChecker) Ping() error {
	input := &rekognition.DescribeProjectsInput{MaxResults: aws.Int64(1)}
	_, err := sc.SVC.DescribeProjects(input)
	if err != nil {
		return err
	}
	return nil
}

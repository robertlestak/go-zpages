package httpchecker

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudfront"
)

// CloudFrontChecker wraps a connection to CloudFront
type CloudFrontChecker struct {
	SVC *cloudfront.CloudFront
	ID  *string
}

// Ping checks to see if CloudFront Distribution exists
func (dc *CloudFrontChecker) Ping() error {
	input := &cloudfront.GetDistributionInput{Id: aws.String(*dc.ID)}
	d, err := dc.SVC.GetDistribution(input)
	if err != nil {
		return err
	}
	ds := *d.Distribution.Status
	if ds != "Deployed" && ds != "InProgress" {
		return fmt.Errorf("distribution status: %s", ds)
	}
	return nil
}

package zpages

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudfront"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/robertlestak/go-zpages/pkg/awsservices"
	"github.com/robertlestak/go-zpages/pkg/storage"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-redis/redis/v8"
	"github.com/robertlestak/go-zpages/pkg/database"
	"github.com/robertlestak/go-zpages/pkg/httpchecker"
)

var (
	// awsSession contains the cached AWSSession
	awsSession *session.Session
)

// AWSSession returns a new AWS Session, using the default
// configured credential provider.
func AWSSession() (*session.Session, error) {
	// if session cached, return from cache
	if awsSession != nil {
		return awsSession, nil
	}
	reg := os.Getenv("AWS_REGION")
	if reg == "" {
		reg = "us-east-1"
	}
	s, err := session.NewSession(&aws.Config{
		Region: aws.String(reg)},
	)
	if err != nil {
		return s, err
	}
	awsSession = s
	return s, nil
}

// Ping pings a HTTP endpoint and returns err
func (d *HTTP) Ping() error {
	c := &httpchecker.Checker{
		Address:     &d.Address,
		Method:      &d.Method,
		Body:        &d.Body,
		StatusCodes: &d.StatusCodes,
	}
	perr := c.Ping()
	if perr != nil {
		return perr
	}
	return nil
}

// Ping pings a CloudFront endpoint and returns err
func (d *CloudFront) Ping() error {
	sess, err := AWSSession()
	if err != nil {
		return err
	}
	c := &httpchecker.CloudFrontChecker{
		ID:  &d.ID,
		SVC: cloudfront.New(sess),
	}
	perr := c.Ping()
	if perr != nil {
		return perr
	}
	return nil
}

// Ping pings a DynamoDB endpoint and returns err
func (d *DynamoDB) Ping() error {
	sess, err := AWSSession()
	if err != nil {
		return err
	}
	c := &database.DynamoChecker{
		TableName: &d.Table,
		SVC:       dynamodb.New(sess),
	}
	perr := c.Ping()
	if perr != nil {
		return perr
	}
	return nil
}

// Ping pings a Elasticsearch endpoint and returns err
func (d *Elasticsearch) Ping() error {
	cfg := elasticsearch.Config{
		Addresses: d.Addresses,
		Username:  d.Username,
		Password:  d.Password,
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return err
	}
	c := &database.ElasticChecker{
		Client: es,
	}
	perr := c.Ping()
	if perr != nil {
		return perr
	}
	return nil
}

// Ping pings a Redis endpoint and returns err
func (d *Redis) Ping() error {
	cl := redis.NewClient(&redis.Options{
		Addr:     d.Address,
		Password: d.Password,
		DB:       d.Database,
	})
	c := &database.RedisChecker{
		Client: cl,
	}
	perr := c.Ping()
	if perr != nil {
		return perr
	}
	return nil
}

// Ping pings a DynamoDB endpoint and returns err
func (d *Rekognition) Ping() error {
	sess, err := AWSSession()
	if err != nil {
		return err
	}
	c := &awsservices.RekognitionChecker{
		SVC: rekognition.New(sess),
	}
	perr := c.Ping()
	if perr != nil {
		return perr
	}
	return nil
}

// Ping pings a S3 endpoint and returns err
func (d *S3) Ping() error {
	sess, err := AWSSession()
	if err != nil {
		return err
	}
	c := &storage.S3Checker{
		SVC:        s3.New(sess),
		BucketName: &d.Bucket,
	}
	perr := c.Ping()
	if perr != nil {
		return perr
	}
	return nil
}

// Ping pings a MySQL endpoint and returns err
func (d *SQL) Ping() error {
	var dataSourceName string
	switch d.Driver {
	case "mysql":
		dataSourceName = fmt.Sprintf(
			"%s:%s@tcp(%s)/%s",
			d.Username,
			d.Password,
			d.Host,
			d.Database,
		)
	case "postgres":
		hs := strings.Split(d.Host, ":")
		h := hs[0]
		p := hs[1]
		dataSourceName = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=require",
			d.Username,
			d.Password,
			h,
			p,
			d.Database,
		)
	}
	c, cerr := database.NewDBChecker(d.Driver, dataSourceName)
	if cerr != nil {
		return cerr
	}
	perr := c.Ping()
	if perr != nil {
		return perr
	}
	return nil
}

func structName(d interface{}) string {
	n := strings.Replace(reflect.TypeOf(d).String(), "*", "", -1)
	n = strings.Replace(n, "zpages.", "", -1)
	return n
}

// Response wraps the response from a Driver
func (d *CloudFront) Response(e error) *Response {
	er := &Response{
		Name: d.Name,
		Type: structName(d),
	}
	if e != nil {
		er.Error = e.Error()
	}
	return er
}

// Response wraps the response from a Driver
func (d *HTTP) Response(e error) *Response {
	er := &Response{
		Name: d.Name,
		Type: structName(d),
	}
	if e != nil {
		er.Error = e.Error()
	}
	return er
}

// Response wraps the response from a Driver
func (d *DynamoDB) Response(e error) *Response {
	er := &Response{
		Name: d.Name,
		Type: structName(d),
	}
	if e != nil {
		er.Error = e.Error()
	}
	return er
}

// Response wraps the response from a Driver
func (d *Elasticsearch) Response(e error) *Response {
	er := &Response{
		Name: d.Name,
		Type: structName(d),
	}
	if e != nil {
		er.Error = e.Error()
	}
	return er
}

// Response wraps the response from a Driver
func (d *SQL) Response(e error) *Response {
	er := &Response{
		Name: d.Name,
		Type: structName(d),
	}
	if e != nil {
		er.Error = e.Error()
	}
	return er
}

// Response wraps the response from a Driver
func (d *Rekognition) Response(e error) *Response {
	er := &Response{
		Name: d.Name,
		Type: structName(d),
	}
	if e != nil {
		er.Error = e.Error()
	}
	return er
}

// Response wraps the response from a Driver
func (d *Redis) Response(e error) *Response {
	er := &Response{
		Name: d.Name,
		Type: structName(d),
	}
	if e != nil {
		er.Error = e.Error()
	}
	return er
}

// Response wraps the response from a Driver
func (d *S3) Response(e error) *Response {
	er := &Response{
		Name: d.Name,
		Type: structName(d),
	}
	if e != nil {
		er.Error = e.Error()
	}
	return er
}

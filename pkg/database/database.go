package database

import (
	"context"
	"database/sql"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql" // mysql
	_ "github.com/lib/pq"              // postgres
)

// DBChecker wraps a DB connection
type DBChecker struct {
	DB *sql.DB
}

// DynamoChecker wraps a service connection to Dynamo
type DynamoChecker struct {
	SVC       *dynamodb.DynamoDB
	TableName *string
}

// ElasticChecker wraps an elasticsearch connection
type ElasticChecker struct {
	Client *elasticsearch.Client
}

// RedisChecker wraps a redis connection
type RedisChecker struct {
	Client *redis.Client
}

// NewDBChecker connects to DB and returns Checker
func NewDBChecker(driverName, dataSourceName string) (*DBChecker, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	return &DBChecker{db}, nil
}

// Ping pings a database to confirm it is available
func (dc *DBChecker) Ping() error {
	err := dc.DB.Ping()
	if err != nil {
		return err
	}
	dc.DB.Close()
	return nil
}

// Ping checks to see if DynamoDB table exists
func (dc *DynamoChecker) Ping() error {
	input := &dynamodb.DescribeTableInput{TableName: aws.String(*dc.TableName)}
	_, err := dc.SVC.DescribeTable(input)
	if err != nil {
		return err
	}
	return nil
}

// Ping checks to see if Elasticsearch is available
func (ec *ElasticChecker) Ping() error {
	_, err := ec.Client.Ping()
	if err != nil {
		return err
	}
	return nil
}

// Ping checks to see if Redis is available
func (dc *RedisChecker) Ping() error {
	_, err := dc.Client.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	dc.Client.Close()
	return nil
}

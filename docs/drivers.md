
# Supported Health Check Drivers

## CloudFront

Requires `cloudfront:GetDistribution` permissions.

`ID`: Distribution ID

## DynamoDB

Requires `dynamodb:DescribeTable` permissions.

`Table`: DynamoDB table name

## Elasticsearch

`Addresses`: slice of `string` addresses to check

`Username`: elasticsearch username

`Password`: elasticsearch password

## HTTP

`Address`: `[proto]://[host]:[port]/[path]`

`Method`: `GET`,`POST`,`PUT`,etc

`Body`: byte array to send

`StatusCodes`: slice of `int`s with valid status codes to expect. ex `[]int{200, 300, 404}`

## Redis

`Address`: redis `host:port`

`Password`: redis password

`Database`: redis database. default `0`

## Rekognition

Requires `rekognition:DescribeProjects` permissions.

## S3

Requires `s3:ListBucket` permissions.

`Bucket`: bucket name to check

## SQL

`Driver`: one of `mysql`, `postgres`

`Host`: database `host:port`

`Database`: database name

`Username`: database username

`Password`: database password

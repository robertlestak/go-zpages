package zpages

// HTTP contains an HTTP probe
type HTTP struct {
	Name        string
	Address     string
	Method      string
	Body        []byte
	StatusCodes []int
}

// CloudFront contains a CloudFront probe
type CloudFront struct {
	Name string
	ID   string
}

// Elasticsearch contains an Elasticsearch probe
type Elasticsearch struct {
	Name      string
	Addresses []string
	Username  string
	Password  string
}

// DynamoDB contains a Dynamo probe
type DynamoDB struct {
	Name  string
	Table string
}

// Redis contains a Redis probe
type Redis struct {
	Name     string
	Address  string
	Password string
	Database int
}

// SQL contains a SQL probe
type SQL struct {
	Name     string
	Driver   string
	Host     string
	Database string
	Username string
	Password string
}

// S3 contains a S3 probe
type S3 struct {
	Name   string
	Bucket string
}

// Rekognition contains a Rekognition probe
type Rekognition struct {
	Name string
}

package database

// PingConfig - Struct to represent the data contained in the ping_config table.
// PingConfig table contains the endpoint to query against and the expected
// status code returned from that endpoint.
type PingConfig struct {
	ID                 int64
	Endpoint           string
	ExpectedStatusCode int
}

// TableName - Gorm convention for determining the table name when executing
// queries
func (PingConfig) TableName() string {
	return "ping_config"
}

package database

// HTTPGetConfig represents the data stored in the http_get_config table.
// id int(11)
// endpoint varchar(255)
// expected_status_code int(11)
type HTTPGetConfig struct {
	ID                 int64
	Endpoint           string
	ExpectedStatusCode int
}

// TableName - Gorm convention for determining the table name when executing
// queries
func (HTTPGetConfig) TableName() string {
	return "http_get_config"
}

// NewHTTPGetConfig constructor for creating a HttpGetConfig
func NewHTTPGetConfig(id int64, endpoint string, expectedStatusCode int) {

}

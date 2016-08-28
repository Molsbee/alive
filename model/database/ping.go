package database

// Ping struct represents the data contained within the ping table and its
// relationships
type Ping struct {
	ID                 int
	Endpoint           string
	ExpectedStatusCode int
}

// TableName - gorm convention for determining the table name when executing
// queries
func (Ping) TableName() string {
	return "ping"
}

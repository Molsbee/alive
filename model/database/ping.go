package database

// Ping struct represents the data contained within the ping table and its
// relationships
type Ping struct {
	ID                 int
	Endpoint           string
	ExpectedStatusCode int
}

// TableName gorm convention for determing table name associated with struct
func (Ping) TableName() string {
	return "ping"
}

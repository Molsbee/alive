package database

import "time"

// PingResponse struct represents the data contained within the ping_response table
type PingResponse struct {
	ID         int
	CreatedTS  time.Time
	Duration   int
	StatusCode int
	PingID     int
}

// TableName - gorm convention for determining the table name when executing
// queries
func (PingResponse) TableName() string {
	return "ping_response"
}

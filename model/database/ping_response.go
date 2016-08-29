package database

import "time"

// PingResponse struct represents the data contained within the ping_response table
type PingResponse struct {
	ID         int64
	CreatedTS  time.Time
	DurationMS int64
	StatusCode int
	PingID     int64
}

// TableName - gorm convention for determining the table name when executing
// queries
func (PingResponse) TableName() string {
	return "ping_response"
}

// NewPingResponse constructor for creating a PingResponse struct without the
// status code field.  It is expected that the user will update the status
// code field before saving the entry
func NewPingResponse(created time.Time, duration int64, statusCode int, pingID int64) *PingResponse {
	return &PingResponse{
		CreatedTS:  created,
		DurationMS: duration,
		StatusCode: statusCode,
		PingID:     pingID,
	}
}

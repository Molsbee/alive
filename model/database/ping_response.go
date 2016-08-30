package database

import "time"

// PingResponse - Struct to represent the data contained in the ping_response
// table.  PingResponse table contains the duration in milliseconds of the call
// and the time the call was executed with the status code received from the
// endpoint and a reference to the PingConfig.
type PingResponse struct {
	ID           int64     `json:"id"`
	CreatedTS    time.Time `json:"createdTS"`
	DurationMS   int64     `json:"durationMS"`
	StatusCode   int       `json:"statusCode"`
	PingConfigID int64     `json:"pingConfigId"`
}

// TableName - Gorm convention for determining the table name when executing
// queries
func (PingResponse) TableName() string {
	return "ping_response"
}

// NewPingResponse constructor for creating a PingResponse struct without the
// status code field.  It is expected that the user will update the status
// code field before saving the entry
func NewPingResponse(created time.Time, duration int64, statusCode int, pingConfigID int64) *PingResponse {
	return &PingResponse{
		CreatedTS:    created,
		DurationMS:   duration,
		StatusCode:   statusCode,
		PingConfigID: pingConfigID,
	}
}

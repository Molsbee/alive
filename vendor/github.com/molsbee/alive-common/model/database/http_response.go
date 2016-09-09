package database

import "time"

// HTTPGetResponse represents the data stored in the http_get_response table.
// id int(11)
// created_ts DATETIME
// duration_ms int(11)
// status_code int(11)
// ping_config_id int(11)
type HTTPGetResponse struct {
	ID              int64
	CreatedTS       time.Time
	DurationMS      int64
	StatusCode      int
	HTTPGetConfigID int64
}

// TableName - Gorm convention for determining the table name when executing
// queries
func (HTTPGetResponse) TableName() string {
	return "http_get_response"
}

// NewHTTPGetResponse constructor for creating a HttpGetResponse
func NewHTTPGetResponse(created time.Time, duration int64, statusCode int, configID int64) *HTTPGetResponse {
	return &HTTPGetResponse{
		CreatedTS:       created,
		DurationMS:      duration,
		StatusCode:      statusCode,
		HTTPGetConfigID: configID,
	}
}

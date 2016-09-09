package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/molsbee/alive-common/model/database"
)

// HTTPGetResponse provides basic repository functions for accessing
// http_get_response table
type HTTPGetResponse struct {
	db *gorm.DB
}

// NewHTTPGetResponse constructor for HTTPGetResponse
func NewHTTPGetResponse(db *gorm.DB) *HTTPGetResponse {
	return &HTTPGetResponse{db: db}
}

// Save creates a new response entrie in http_get_response table
func (r *HTTPGetResponse) Save(response database.HTTPGetResponse) error {
	if err := r.db.Create(&response).Error; err != nil {
		return err
	}

	return nil
}

// FindAllByConfigID returns all entries from http_get_response table based on
// the id provided which corresponds to the http_get_config_id column
func (r *HTTPGetResponse) FindAllByConfigID(id int64) ([]database.HTTPGetResponse, error) {
	responses := []database.HTTPGetResponse{}
	if err := r.db.Where("http_get_config_id=?", id).Find(&responses).Error; err != nil {
		return nil, err
	}

	return responses, nil
}

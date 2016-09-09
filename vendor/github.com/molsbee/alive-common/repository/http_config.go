package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/molsbee/alive-common/model/database"
)

// HTTPGetConfig provides basic repository functions for accessing
// http_get_config table
type HTTPGetConfig struct {
	db *gorm.DB
}

// NewHTTPGetConfig constructor for HTTPGetConfig
func NewHTTPGetConfig(db *gorm.DB) *HTTPGetConfig {
	return &HTTPGetConfig{db: db}
}

// Save creates a new config record in http_get_config table
func (r *HTTPGetConfig) Save(config database.HTTPGetConfig) error {
	if err := r.db.Create(&config).Error; err != nil {
		return err
	}
	return nil
}

// FindAll returns all entries in http_get_config table
func (r *HTTPGetConfig) FindAll() ([]database.HTTPGetConfig, error) {
	configs := []database.HTTPGetConfig{}
	if err := r.db.Find(&configs).Error; err != nil {
		return nil, err
	}

	return configs, nil
}

// FindByID returns a single row from http_get_config table based on the id
// provided
func (r *HTTPGetConfig) FindByID(id int64) (*database.HTTPGetConfig, error) {
	config := database.HTTPGetConfig{}
	if err := r.db.Where("id=?").First(&config).Error; err != nil {
		return nil, err
	}

	return &config, nil
}

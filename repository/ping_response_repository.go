package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/molsbee/alive/model/database"
)

// PingResponseRepository provides basic mechanisms for quering ping_response
// table
type PingResponseRepository struct {
	db *gorm.DB
}

// NewPingResponseRepository convenience method for constructing a ping response
// repository
func NewPingResponseRepository(db *gorm.DB) PingResponseRepository {
	return PingResponseRepository{db: db}
}

// Save creates a database entry for the outcome of ping test
func (pr *PingResponseRepository) Save(response database.PingResponse) error {
	if err := pr.db.Create(&response).Error; err != nil {
		return err
	}

	return nil
}

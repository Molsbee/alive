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

// FindAllByPingConfigID lookup all entries in ping_response table filtered by
// ping_config_id
func (pr *PingResponseRepository) FindAllByPingConfigID(id int64) ([]database.PingResponse, error) {
	pingResponses := []database.PingResponse{}
	if err := pr.db.Where("ping_config_id=?", id).Find(&pingResponses).Error; err != nil {
		return nil, err
	}

	return pingResponses, nil
}

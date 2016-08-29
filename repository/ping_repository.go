package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/molsbee/alive/model/database"
)

// PingRepository provides basic mechanisms for querying ping table
// ping table contains all the configurations of endpoints to call
type PingRepository struct {
	db *gorm.DB
}

// NewPingRepository convenience method for constructing a ping repository
func NewPingRepository(db *gorm.DB) PingRepository {
	return PingRepository{db: db}
}

// Save creates a database recorded for the PingConfig struct
func (p *PingRepository) Save(ping database.PingConfig) error {
	if err := p.db.Create(&ping).Error; err != nil {
		return err
	}

	return nil
}

// FindAll lookup all entries in ping config table
func (p *PingRepository) FindAll() ([]database.PingConfig, error) {
	pings := []database.PingConfig{}
	if err := p.db.Find(&pings).Error; err != nil {
		return nil, err
	}

	return pings, nil
}

// FindByID looks up a single ping config recorded by the id provided
func (p *PingRepository) FindByID(id int64) (*database.PingConfig, error) {
	ping := database.PingConfig{}
	if err := p.db.Where("id=?", id).First(&ping).Error; err != nil {
		return nil, err
	}

	return &ping, nil
}

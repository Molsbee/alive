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

// Save creates a database recorded for the ping struct
func (p *PingRepository) Save(ping database.Ping) error {
	if err := p.db.Create(&ping).Error; err != nil {
		return err
	}

	return nil
}

// FindAll lookup all entries in ping table
func (p *PingRepository) FindAll() ([]database.Ping, error) {
	pings := []database.Ping{}
	if err := p.db.Find(&pings).Error; err != nil {
		return nil, err
	}

	return pings, nil
}

// FindByID looks up a single ping recorded by the id provided
func (p *PingRepository) FindByID(id int) (*database.Ping, error) {
	ping := database.Ping{}
	if err := p.db.Where("id=?", id).First(&ping).Error; err != nil {
		return nil, err
	}

	return &ping, nil
}

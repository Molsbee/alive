package resource

import (
	"encoding/json"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/molsbee/alive/model/database"
	"github.com/molsbee/alive/repository"
)

// PingResource - restful api wrapping retreiving, manipulating and creating
// ping configurations within database
type PingResource struct {
	repository repository.PingRepository
}

// NewPingResource - restful api wrapping retreiving, manipulating and creating
// ping configurations within database
func NewPingResource(db *gorm.DB) PingResource {
	return PingResource{
		repository: repository.NewPingRepository(db),
	}
}

// Get - HandlerFunc for servicing web requests for all ping configurations
func (pr *PingResource) Get(w http.ResponseWriter, r *http.Request) {
	pingConfigs, err := pr.repository.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("ContentType", "application/json")
	json.NewEncoder(w).Encode(pingConfigs)
}

// Create - HandlerFunc for creating ping configuration entries in database
func (pr *PingResource) Create(w http.ResponseWriter, r *http.Request) {
	pingConfig := database.Ping{}
	if err := json.NewDecoder(r.Body).Decode(&pingConfig); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := pr.repository.Save(pingConfig); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}

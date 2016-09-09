package resource

import (
	"encoding/json"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/molsbee/alive-common/model/database"
	"github.com/molsbee/alive-common/repository"
)

// HTTPConfigResource - restful api wrapping retreiving, manipulating and creating
// ping configurations within database
type HTTPConfigResource struct {
	repo *repository.HTTPGetConfig
}

// NewHTTPConfigResource - restful api wrapping retreiving, manipulating and creating
// ping configurations within database
func NewHTTPConfigResource(db *gorm.DB) *HTTPConfigResource {
	return &HTTPConfigResource{
		repo: repository.NewHTTPGetConfig(db),
	}
}

// Get - HandlerFunc for servicing web requests for all ping configurations
func (pr *HTTPConfigResource) Get(w http.ResponseWriter, r *http.Request) {
	httpConfigs, err := pr.repo.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("ContentType", "application/json")
	json.NewEncoder(w).Encode(httpConfigs)
}

// Create - HandlerFunc for creating ping configuration entries in database
func (pr *HTTPConfigResource) Create(w http.ResponseWriter, r *http.Request) {
	httpConfig := database.HTTPGetConfig{}
	if err := json.NewDecoder(r.Body).Decode(&httpConfig); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := pr.repo.Save(httpConfig); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

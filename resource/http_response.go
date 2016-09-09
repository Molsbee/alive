package resource

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/molsbee/alive-common/repository"
)

// HTTPResponseResource - restful api wrapping retreiving ping responses
type HTTPResponseResource struct {
	repo *repository.HTTPGetResponse
}

// NewHTTPResponseResource constructor for creating a new instance of
// PingResponseResource
func NewHTTPResponseResource(db *gorm.DB) *HTTPResponseResource {
	return &HTTPResponseResource{
		repo: repository.NewHTTPGetResponse(db),
	}
}

// Get - HandlerFunc for servicing web request for all ping responses
func (pr *HTTPResponseResource) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	configID, err := strconv.ParseInt(vars["configID"], 10, 64)
	if err != nil {
		http.Error(w, "Config ID required", http.StatusBadRequest)
		return
	}

	responses, err := pr.repo.FindAllByConfigID(configID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responses)
}

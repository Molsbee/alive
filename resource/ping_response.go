package resource

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/molsbee/alive/repository"
)

// PingResponseResource - restful api wrapping retreiving ping responses
type PingResponseResource struct {
	respRep repository.PingResponseRepository
}

// NewPingResponseResource constructor for creating a new instance of
// PingResponseResource
func NewPingResponseResource(db *gorm.DB) PingResponseResource {
	return PingResponseResource{
		respRep: repository.NewPingResponseRepository(db),
	}
}

// Get - HandlerFunc for servicing web request for all ping responses
func (pr *PingResponseResource) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pingConfigID, err := strconv.ParseInt(vars["pingConfigID"], 10, 64)
	if err != nil {
		http.Error(w, "Ping Config ID required", http.StatusBadRequest)
		return
	}

	pingResponses, err := pr.respRep.FindAllByPingConfigID(pingConfigID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pingResponses)
}

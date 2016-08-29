package service

import (
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/molsbee/alive/model/database"
	"github.com/molsbee/alive/repository"
)

type pingService struct {
	pingRepo    repository.PingRepository
	respRepo    repository.PingResponseRepository
	workChannel chan database.PingConfig
}

// StartPingService populates the pingService struct with all the required
// dependencies.  Starts asynchronously polling the ping table and delegates
// ping work requests to the pingWorkers setup within this start section.
func StartPingService(db *gorm.DB) {
	service := pingService{
		pingRepo:    repository.NewPingRepository(db),
		respRepo:    repository.NewPingResponseRepository(db),
		workChannel: make(chan database.PingConfig, 20),
	}

	// Startup ping works in seperate go routines that will poll workChannel
	// for endpoints to ping
	for i := 0; i <= 3; i++ {
		go service.pingWorker(i)
	}

	go func() {
		for {
			pingConfigurations, _ := service.pingRepo.FindAll()
			for _, pingConf := range pingConfigurations {
				service.workChannel <- pingConf
			}
			time.Sleep(time.Minute)
		}
	}()
}

func (ps *pingService) pingWorker(id int) {
	for p := range ps.workChannel {
		start, duration, statusCode := ping(p.Endpoint)
		pingResponse := database.NewPingResponse(start, duration, statusCode, p.ID)
		ps.respRepo.Save(*pingResponse)
	}
}

func ping(endpoint string) (time.Time, int64, int) {
	start := time.Now()
	resp, err := http.Get(endpoint)
	duration := time.Since(start).Nanoseconds() / 1000000
	if err != nil {
		return start, -1, 0
	}
	defer resp.Body.Close()

	return start, duration, resp.StatusCode
}

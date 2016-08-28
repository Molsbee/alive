package service

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/molsbee/alive/repository"
)

type pingService struct {
	repository  repository.PingRepository
	workChannel chan repository.Ping
}

func (ps *pingService) pingWorker(id int) {
	for p := range ps.workChannel {
		resp, err := http.Get(p.Endpoint)
		if err != nil {
			log.Println(err)
		}

		if resp.StatusCode != p.ExpectedStatusCode {
			log.Println("Unexpected status code")
		}

		log.Println("Success")
	}
}

// StartPingService -
func StartPingService(db *gorm.DB) {
	service := pingService{
		repository:  repository.NewPingRepository(db),
		workChannel: make(chan repository.Ping, 20),
	}

	// Startup ping works in seperate go routines that will poll workChannel
	// for endpoints to ping
	for i := 0; i <= 3; i++ {
		go service.pingWorker(i)
	}

	go func() {
		for {
			pingConfigurations, _ := service.repository.FindAll()
			for _, pingConf := range pingConfigurations {
				service.workChannel <- pingConf
			}
			time.Sleep(time.Minute)
		}
	}()
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

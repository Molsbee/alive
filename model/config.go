package model

import "github.com/molsbee/alive/repository"

// Config -
type Config struct {
	Ping []repository.Ping `json:"ping"`
}

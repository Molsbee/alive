package main

import (
	"flag"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/molsbee/alive/resource"
	"github.com/molsbee/alive/service"
)

var databaseURL string

func init() {
	flag.StringVar(&databaseURL, "database", "root@tcp(localhost:3306)/alive?parseTime=true", "-database=root@tcp(localhost:3306)/alive?parseTime=true")
}

func main() {
	log.Printf("Attempting to establish database connection - %s\n", databaseURL)
	db, err := gorm.Open("mysql", databaseURL)
	if err != nil {
		log.Fatal("Unable to establish database connection")
	}
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(20)
	db.LogMode(true)

	service.StartPingService(db)

	router := mux.NewRouter()

	pingResource := resource.NewPingResource(db)
	router.HandleFunc("/configuration/ping", pingResource.Get).Methods("GET")
	router.HandleFunc("/configuration/ping", pingResource.Create).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}

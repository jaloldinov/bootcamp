package main

import (
	"app/api"
	"app/config"
	"app/storage/postgres"
	"log"
	"net/http"
)

func main() {
	cfg := config.Load()

	pgconn, err := postgres.NewConnection(&cfg)

	if err != nil {
		log.Printf("Not Have Connection with Postgres: %+v", err)
	}
	defer pgconn.Close()
	api.NewApi(&cfg, pgconn)

	log.Println("Listening on", cfg.Port)

	err = http.ListenAndServe(":"+cfg.Port, nil)
	if err != nil {
		panic("Server cannot run: " + err.Error())
	}
}

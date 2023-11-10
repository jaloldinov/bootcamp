package main

import (
	"log"
	"net/http"

	"app/api"
	"app/config"
	"app/storage/postgres"
)

func main() {
	cfg := config.Load()

	pgconn, err := postgres.NewConnection(&cfg)

	if err != nil {
		log.Printf("Not Have Connection with Postgres: %+v", err)
	}
	defer pgconn.Close()
	api.NewApi(&cfg, pgconn)

	log.Println("Listening....", cfg.ServerHost+cfg.HTTPPort)

	if err := http.ListenAndServe(cfg.ServerHost+cfg.HTTPPort, nil); err != nil {
		panic("server does not run: " + err.Error())
	}
}

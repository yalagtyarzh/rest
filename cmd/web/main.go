package main

import (
	"log"
	"net/http"

	"github.com/yalagtyarzh/rest/internal/config"
	"github.com/yalagtyarzh/rest/internal/handlers"
	"github.com/yalagtyarzh/rest/internal/repository/dbrepo"
	"github.com/yalagtyarzh/rest/internal/router"
)

func main() {
	log.Println("reading environment")
	cfg := config.GetConfig()

	log.Println("router initializing")
	mux := router.New()

	log.Println("database initializing")
	repo := dbrepo.NewMemStorage()

	h := handlers.NewRepo(repo)
	handlers.NewHandler(h)

	log.Println("Starting!")
	log.Fatal(http.ListenAndServe(cfg.IP+cfg.Port, mux))
}

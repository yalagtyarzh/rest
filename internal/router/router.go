package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	chiware "github.com/go-chi/chi/v5/middleware"

	"github.com/yalagtyarzh/rest/internal/handlers"
	"github.com/yalagtyarzh/rest/internal/middleware"
)

// New creates new router with necessary endpoints and handlers and returns it
func New() http.Handler {
	mux := chi.NewRouter()

	mux.Use(chiware.Recoverer)
	mux.Use(middleware.EventLogger)

	mux.Get("/get-balance", handlers.GetBalance)
	mux.Post("/create-account", handlers.CreateAccount)
	mux.Post("/transfer-money", handlers.TransferMoney)

	return mux
}

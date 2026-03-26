package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rwidjojo/HanchanManager/internal/api/handler"
	"github.com/rwidjojo/HanchanManager/internal/repository"
	"github.com/rwidjojo/HanchanManager/internal/service"
)

// ToDo: if we want to use a repository/db.go then we need to change the code like this
// func NewRouter(db *repository.DB) http.Handler {
func NewRouter(db *pgxpool.Pool) http.Handler {
	r := chi.NewRouter()

	// Global middleware
	r.Use(chimw.StripSlashes)

	// Handlers
	// ToDo: if we want to use a repository/db.go then we need to change the code like this
	// playerRepo := repository.NewPlayerRepo(db.Pool())

	playerRepo := repository.NewPlayerRepo(db)
	playerSvc := service.NewPlayerService(playerRepo)
	playerHandler := handler.NewPlayerHandler(playerSvc)

	// Health check
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	// Players (global — players exist across groups)
	r.Route("/players", func(r chi.Router) {
		r.Post("/", playerHandler.Create)

	})

	return r
}

package api

import (
	"net/http"

	"HanchanManager/internal/api/handler"
	"HanchanManager/internal/repository"
	"HanchanManager/internal/service"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewRouter(db *pgxpool.Pool) http.Handler {
	r := chi.NewRouter()

	// Global middleware
	r.Use(chimw.StripSlashes)

	// Players
	playerRepo := repository.NewPlayerRepo(db)
	playerSvc := service.NewPlayerService(playerRepo)
	playerHandler := handler.NewPlayerHandler(playerSvc)

	// Groups
	groupRepo := repository.NewGroupRepo(db)
	membershipRepo := repository.NewMembershipRepo(db)
	groupSvc := service.NewGroupService(groupRepo, membershipRepo)
	groupHandler := handler.NewGroupHandler(groupSvc)

	// Health check
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	// Players (global — players exist across groups)
	r.Route("/players", func(r chi.Router) {
		r.Post("/", playerHandler.Create)
		r.Get("/", playerHandler.List)
		r.Get("/{id}", playerHandler.GetByID)
	})

	r.Route("/groups", func(r chi.Router) {
		r.Post("/", groupHandler.Create)
		r.Get("/{id}", groupHandler.GetByID)
		r.Post("/{groupID}/players", groupHandler.AddPlayer)
		r.Get("/{groupID}/players", groupHandler.GetPlayers)
	})

	return r
}

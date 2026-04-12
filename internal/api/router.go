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

	// Hanchan
	hanchanRepo := repository.NewHanchanRepo(db)
	hanchanSvc := service.NewHanchanService(hanchanRepo)
	hanchanHandler := handler.NewHanchanHandler(hanchanSvc)

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
		r.Post("/{id}/players", groupHandler.AddPlayer)
		r.Get("/{id}/players", groupHandler.GetPlayers)
	})

	// Hanchans (nested under group for creation, standalone for game ops)
	r.Route("/groups/{groupID}/hanchans", func(r chi.Router) {
		r.Post("/", hanchanHandler.Create)
	})

	r.Route("/hanchans", func(r chi.Router) {
		r.Get("/{id}", hanchanHandler.GetByID)
	})

	return r
}

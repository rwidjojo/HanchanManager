package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"HanchanManager/internal/repository"
	"HanchanManager/internal/service"

	"github.com/go-chi/chi/v5"
)

type PlayerHandler struct {
	svc *service.PlayerService
}

func NewPlayerHandler(svc *service.PlayerService) *PlayerHandler {
	return &PlayerHandler{svc: svc}
}

type createPlayerRequest struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}

func (h *PlayerHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req createPlayerRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	player, err := h.svc.CreatePlayer(r.Context(), req.Username, req.Name)
	if err != nil {
		http.Error(w, "failed to create player", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(player)
}

func (h *PlayerHandler) List(w http.ResponseWriter, r *http.Request) {
	players, err := h.svc.ListPlayers(r.Context())

	if err != nil {
		http.Error(w, "failed to list players", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(players)
}

func (h *PlayerHandler) GetByID(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "invalid player ID", http.StatusBadRequest)
		return
	}

	player, err := h.svc.GetPlayerByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			http.Error(w, "player not found", http.StatusNotFound)
			return
		}

		http.Error(w, "failed to get player", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(player)
}

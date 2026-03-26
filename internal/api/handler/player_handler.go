package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rwidjojo/HanchanManager/internal/service"
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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(player)
}

func (h *PlayerHandler) List(w http.ResponseWriter, r *http.Request) {
	players, err := h.svc.ListPlayers(r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(players)
}

func (h *PlayerHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	player, err := h.svc.GetPlayerByPlayerID(r.Context(), username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(player)
}

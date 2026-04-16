package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"HanchanManager/internal/service"

	"github.com/go-chi/chi/v5"
)

type GroupHandler struct {
	svc *service.GroupService
}

func NewGroupHandler(svc *service.GroupService) *GroupHandler {
	return &GroupHandler{svc: svc}
}

type createGroupRequest struct {
	Code        string  `json:"code"`
	Description *string `json:"description,omitempty"`
}

func (h *GroupHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req createGroupRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	group, err := h.svc.CreateGroup(r.Context(), req.Code, req.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(group)
}

func (h *GroupHandler) GetByID(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Printf("error during conversion: %v\n", err)
		return
	}

	group, err := h.svc.GetGroupByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(group)
}

func (h *GroupHandler) AddPlayer(w http.ResponseWriter, r *http.Request) {

	groupID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Printf("error during conversion: %v\n", err)
		return
	}

	var req struct {
		PlayerID int `json:"player_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	if err := h.svc.AddPlayer(r.Context(), groupID, req.PlayerID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *GroupHandler) GetPlayers(w http.ResponseWriter, r *http.Request) {

	groupID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Printf("error during conversion: %v\n", err)
		return
	}

	players, err := h.svc.GetPlayers(r.Context(), groupID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(players)
}

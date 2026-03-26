package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rwidjojo/HanchanManager/internal/service"
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

func (h *GroupHandler) GetByCode(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")

	group, err := h.svc.GetGroupByCode(r.Context(), code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(group)
}

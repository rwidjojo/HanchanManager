package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"HanchanManager/internal/service"

	"github.com/go-chi/chi/v5"
)

type HanchanHandler struct {
	svc *service.HanchanService
}

func NewHanchanHandler(svc *service.HanchanService) *HanchanHandler {
	return &HanchanHandler{svc: svc}
}

type createHanchanRequest struct {
	Name      *string   `json:"name,omitempty"`
	Date      time.Time `json:"date"`
	BaseScore *int      `json:"base_score,omitempty"`
	Uma       *[]int    `json:"uma,omitempty"`
}

func (h *HanchanHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req createHanchanRequest

	groupID, err := strconv.Atoi(chi.URLParam(r, "groupID"))
	if err != nil {
		fmt.Printf("error during conversion: %v\n", err)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	hanchan, err := h.svc.CreateHanchan(r.Context(), groupID, req.Name, req.Date, req.Uma, req.BaseScore)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(hanchan)
}

func (h *HanchanHandler) GetByID(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Printf("error during conversion: %v\n", err)
		return
	}

	hanchan, err := h.svc.GetHanchanByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(hanchan)
}

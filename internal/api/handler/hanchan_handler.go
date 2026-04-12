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
	Name *string   `json:"name,omitempty"`
	Date time.Time `json:"date"`
	Uma  *[]int    `json:"uma,omitempty"`
}

func (h *HanchanHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req createHanchanRequest
	var uma []int

	groupID, err := strconv.Atoi(chi.URLParam(r, "groupID"))
	if err != nil {
		fmt.Printf("Error during conversion: %v\n", err)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if req.Date.IsZero() {
		req.Date = time.Now()
	}

	if req.Uma == nil {
		uma = []int{15000, 5000, -5000, -15000}
	} else {
		uma = *req.Uma
	}

	if len(uma) != 4 {
		http.Error(w, fmt.Sprintf("Uma must have exactly 4 values, got: %v", uma), http.StatusBadRequest)
		return
	}

	hanchan, err := h.svc.CreateHanchan(r.Context(), groupID, req.Name, req.Date, uma)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(hanchan)
}

func (h *HanchanHandler) GetByID(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Printf("Error during conversion: %v\n", err)
		return
	}

	hanchan, err := h.svc.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(hanchan)
}

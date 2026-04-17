package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"HanchanManager/internal/domain"
	"HanchanManager/internal/service"

	"github.com/go-chi/chi/v5"
)

type HanchanHandler struct {
	svc *service.HanchanService
}

func NewHanchanHandler(svc *service.HanchanService) *HanchanHandler {
	return &HanchanHandler{svc: svc}
}

type playerSeatingRequest struct {
	PlayerID    int    `json:"player_id"`
	InitialSeat string `json:"initial_seat"`
}

type createHanchanRequest struct {
	Name      *string                `json:"name,omitempty"`
	Date      time.Time              `json:"date"`
	BaseScore *int                   `json:"base_score,omitempty"`
	Uma       *[]int                 `json:"uma,omitempty"`
	Seating   []playerSeatingRequest `json:"seating"`
}

func mapSeating(req []playerSeatingRequest) ([]domain.PlayerSeating, error) {
	seating := make([]domain.PlayerSeating, 0, len(req))

	for _, s := range req {
		sw := domain.SeatWind(s.InitialSeat)

		if !sw.IsValid() {
			return nil, fmt.Errorf("invalid seat: %s", s.InitialSeat)
		}

		seating = append(seating, domain.PlayerSeating{
			PlayerID:    s.PlayerID,
			InitialSeat: sw,
		})
	}

	return seating, nil
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

	seating, err := mapSeating(req.Seating)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	input := service.CreateHanchanInput{
		GroupID:   groupID,
		Name:      req.Name,
		Date:      req.Date,
		BaseScore: req.BaseScore,
		Uma:       req.Uma,
		Seating:   seating,
	}

	hanchan, err := h.svc.CreateHanchan(r.Context(), input)
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

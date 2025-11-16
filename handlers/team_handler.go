package handlers

import (
	"encoding/json"
	"net/http"

	"prservice/service"

	"github.com/gorilla/mux"
)

type TeamHandler struct {
	svc *service.TeamService
}

func RegisterTeamRoutes(r *mux.Router, svc *service.TeamService) {
	h := &TeamHandler{svc: svc}
	r.HandleFunc("/team/get", h.GetTeam).Methods("GET")
}

func (h *TeamHandler) GetTeam(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]any{"status": "ok"})
}

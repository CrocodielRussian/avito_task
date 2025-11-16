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
	team := r.PathPrefix("/team").Subrouter()

	team.HandleFunc("/get", h.GetTeam).Methods("GET")
	team.HandleFunc("/add", h.AddTeam).Methods("POST")
}

func (h *TeamHandler) GetTeam(w http.ResponseWriter, r *http.Request) {
	var req setActiveRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	h.svc.GetTeam(r.Context(), req.)
	json.NewEncoder(w).Encode(map[string]any{"status": "ok"})
}

func (h *TeamHandler) AddTeam(w http.ResponseWriter, r *http.Request) {
	h.svc.AddTeam(r.Context())
	json.NewEncoder(w).Encode(map[string]string{"ok": "team add endpoint (skeleton)"})
}

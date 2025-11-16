package handlers

import (
	"encoding/json"
	"net/http"

	"prservice/service"

	"github.com/gorilla/mux"
)

type PRHandler struct {
	svc *service.PRService
}

func RegisterPRRoutes(r *mux.Router, svc *service.PRService) {
	h := &PRHandler{svc: svc}
	pr := r.PathPrefix("/pullRequest").Subrouter()

	pr.HandleFunc("/create", h.Create).Methods("POST")
	pr.HandleFunc("/merge", h.Merge).Methods("POST")
	pr.HandleFunc("/reassign", h.Reassign).Methods("POST")
}

func (h *PRHandler) Create(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]any{"status": "ok"})
}
func (h *PRHandler) Merge(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]any{"status": "ok"})
}
func (h *PRHandler) Reassign(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]any{"status": "ok"})
}

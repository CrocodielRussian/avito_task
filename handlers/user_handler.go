package handlers

import (
	"encoding/json"
	"net/http"

	"prservice/service"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	svc *service.UserService
}

func RegisterUserRoutes(r *mux.Router, svc *service.UserService) {
	h := &UserHandler{svc: svc}
	r.HandleFunc("/users/setIsActive", h.SetActive).Methods("POST")
}

type setActiveRequest struct {
	UserID   string `json:"user_id"`
	IsActive bool   `json:"is_active"`
}

func (h *UserHandler) SetActive(w http.ResponseWriter, r *http.Request) {
	var req setActiveRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.svc.SetActive(r.Context(), req.UserID, req.IsActive)

	json.NewEncoder(w).Encode(map[string]any{"status": "ok"})
}

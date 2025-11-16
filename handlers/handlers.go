package handlers

import (
	"encoding/json"
	"net/http"

	"prservice/service"
)

type Handler struct {
	svc *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{svc: s}
}

func (h *Handler) writeJSON(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	h.writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (h *Handler) AddTeam(w http.ResponseWriter, r *http.Request) {
	var req map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid json"})
		return
	}
	h.writeJSON(w, http.StatusCreated, map[string]string{"ok": "team add endpoint (skeleton)"})
}

func (h *Handler) GetTeam(w http.ResponseWriter, r *http.Request) {
	h.writeJSON(w, http.StatusOK, map[string]string{"note": "get team (skeleton)"})
}

func (h *Handler) SetIsActive(w http.ResponseWriter, r *http.Request) {
	h.writeJSON(w, http.StatusOK, map[string]string{"note": "set isActive (skeleton)"})
}

func (h *Handler) GetReviewsForUser(w http.ResponseWriter, r *http.Request) {
	h.writeJSON(w, http.StatusOK, map[string]string{"note": "get reviews (skeleton)"})
}

func (h *Handler) AddPR(w http.ResponseWriter, r *http.Request) {
	h.writeJSON(w, http.StatusCreated, map[string]string{"note": "create pr (skeleton)"})
}

func (h *Handler) MergePR(w http.ResponseWriter, r *http.Request) {
	h.writeJSON(w, http.StatusOK, map[string]string{"note": "merge pr (skeleton)"})
}

func (h *Handler) ReassignReviewer(w http.ResponseWriter, r *http.Request) {
	h.writeJSON(w, http.StatusOK, map[string]string{"note": "reassign (skeleton)"})
}

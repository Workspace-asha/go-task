package handlers

import (
 "encoding/json"
 "net/http"
 "strconv"

 "github.com/go-chi/chi/v5"
 "task-api/internal/models"
 "task-api/internal/service"
)

type Handler struct {
 svc *service.ProjectService
}

func New(s *service.ProjectService) *Handler {
 return &Handler{s}
}

func (h *Handler) Router() http.Handler {
 r := chi.NewRouter()

 r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("ok"))
 })

 r.Post("/projects", h.Create)
 r.Get("/projects", h.List)

 return r
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
 var p models.Project
 json.NewDecoder(r.Body).Decode(&p)

 if err := h.svc.Create(r.Context(), &p); err != nil {
  http.Error(w, err.Error(), 500)
  return
 }
 json.NewEncoder(w).Encode(p)
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
 page, _ := strconv.Atoi(r.URL.Query().Get("page"))
 if page == 0 {
  page = 1
 }
 limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
 if limit == 0 {
  limit = 10
 }

 res, err := h.svc.List(r.Context(), page, limit)
 if err != nil {
  http.Error(w, err.Error(), 500)
  return
 }
 json.NewEncoder(w).Encode(res)
}
package handlers

import (
 "encoding/json"
 "net/http"
 "strconv"

 "github.com/go-chi/chi/v5"
 "go-task/internal/models"
 "go-task/internal/service"
 "go-task/internal/middleware"

)

type Handler struct {
	svc service.ProjectServiceInterface
}

func New(svc service.ProjectServiceInterface) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Router() http.Handler {
 r := chi.NewRouter()

 r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("ok"))
 })

 r.Post("/login", Login)

 r.Route("/projects", func(r chi.Router) {
 r.Use(middleware.JWTAuth)
 r.Post("/", h.Create)
 r.Get("/", h.List)
 r.Get("/{id}", h.GetByID)
 r.Put("/{id}", h.Update)
 r.Delete("/{id}", h.Delete)
 })

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

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
 id, _ := strconv.Atoi(chi.URLParam(r, "id"))

 p, err := h.svc.GetByID(r.Context(), id)
 if err != nil {
  http.Error(w, "not found", 404)
  return
 }

 json.NewEncoder(w).Encode(p)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
 id, _ := strconv.Atoi(chi.URLParam(r, "id"))

 var p models.Project
 json.NewDecoder(r.Body).Decode(&p)

 if err := h.svc.Update(r.Context(), id, &p); err != nil {
  http.Error(w, err.Error(), 500)
  return
 }

 w.Write([]byte("updated"))
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
 id, _ := strconv.Atoi(chi.URLParam(r, "id"))

 if err := h.svc.Delete(r.Context(), id); err != nil {
  http.Error(w, err.Error(), 500)
  return
 }

 w.Write([]byte("deleted"))
}
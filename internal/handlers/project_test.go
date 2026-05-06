package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"context"

	"go-task/internal/models"

	"github.com/go-chi/chi/v5"
)

// mock service
type mockService struct{}

func (m *mockService) Create(ctx context.Context, p *models.Project) error {
	p.ID = 1
	return nil
}
func (m *mockService) List(ctx context.Context, page, limit int) ([]models.Project, error) {
	return []models.Project{{ID: 1, Name: "Test"}}, nil
}
func (m *mockService) GetByID(ctx context.Context, id int) (*models.Project, error) {
	return &models.Project{ID: id, Name: "Test"}, nil
}
func (m *mockService) Update(ctx context.Context, id int, p *models.Project) error {
	return nil
}
func (m *mockService) Delete(ctx context.Context, id int) error {
	return nil
}

func TestCreateHandler(t *testing.T) {
	svc := &mockService{}
	h := New(svc)

	body := `{"name":"Test","description":"Demo"}`
	req := httptest.NewRequest("POST", "/projects", bytes.NewBuffer([]byte(body)))
	rec := httptest.NewRecorder()

	h.Create(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
}

func TestListHandler(t *testing.T) {
	svc := &mockService{}
	h := New(svc)

	req := httptest.NewRequest("GET", "/projects?page=1&limit=10", nil)
	rec := httptest.NewRecorder()

	h.List(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
}

func TestGetByIDHandler(t *testing.T) {
	svc := &mockService{}
	h := New(svc)

	req := httptest.NewRequest("GET", "/projects/1", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	rec := httptest.NewRecorder()

	h.GetByID(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
}

func TestUpdateHandler(t *testing.T) {
	svc := &mockService{}
	h := New(svc)

	body := `{"name":"Updated","description":"Updated"}`
	req := httptest.NewRequest("PUT", "/projects/1", bytes.NewBuffer([]byte(body)))

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	rec := httptest.NewRecorder()

	h.Update(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
}

func TestDeleteHandler(t *testing.T) {
	svc := &mockService{}
	h := New(svc)

	req := httptest.NewRequest("DELETE", "/projects/1", nil)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	rec := httptest.NewRecorder()

	h.Delete(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
}
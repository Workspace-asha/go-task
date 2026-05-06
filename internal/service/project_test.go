package service

import (
	"context"
	"testing"

	"go-task/internal/models"
)

// mock repo
type mockRepo struct{}

func (m *mockRepo) Create(ctx context.Context, p *models.Project) error {
	p.ID = 1
	return nil
}

func (m *mockRepo) List(ctx context.Context, limit, offset int) ([]models.Project, error) {
	return []models.Project{
		{ID: 1, Name: "Test", Description: "Demo"},
	}, nil
}

func (m *mockRepo) GetByID(ctx context.Context, id int) (*models.Project, error) {
	return &models.Project{ID: id, Name: "Test"}, nil
}

func (m *mockRepo) Update(ctx context.Context, id int, p *models.Project) error {
	return nil
}

func (m *mockRepo) Delete(ctx context.Context, id int) error {
	return nil
}

func TestCreateProject(t *testing.T) {
	repo := &mockRepo{}
	svc := NewProjectService(repo)

	p := &models.Project{Name: "Unit Test"}

	err := svc.Create(context.Background(), p)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if p.ID == 0 {
		t.Fatalf("expected ID to be set")
	}
}

func TestListProjects(t *testing.T) {
	repo := &mockRepo{}
	svc := NewProjectService(repo)

	res, err := svc.List(context.Background(), 1, 10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(res) == 0 {
		t.Fatalf("expected results")
	}
}

func TestGetByID(t *testing.T) {
	repo := &mockRepo{}
	svc := NewProjectService(repo)

	p, err := svc.GetByID(context.Background(), 1)
	if err != nil {
		t.Fatalf("unexpected error")
	}

	if p.ID != 1 {
		t.Fatalf("wrong project")
	}
}

func TestUpdateProject(t *testing.T) {
	repo := &mockRepo{}
	svc := NewProjectService(repo)

	err := svc.Update(context.Background(), 1, &models.Project{Name: "Updated"})
	if err != nil {
		t.Fatalf("unexpected error")
	}
}

func TestDeleteProject(t *testing.T) {
	repo := &mockRepo{}
	svc := NewProjectService(repo)

	err := svc.Delete(context.Background(), 1)
	if err != nil {
		t.Fatalf("unexpected error")
	}
}
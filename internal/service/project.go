package service

import (
 "context"
 "go-task/internal/models"
 "go-task/internal/repository"
)

type ProjectServiceInterface interface {
	Create(ctx context.Context, p *models.Project) error
	List(ctx context.Context, page, limit int) ([]models.Project, error)
	GetByID(ctx context.Context, id int) (*models.Project, error)
	Update(ctx context.Context, id int, p *models.Project) error
	Delete(ctx context.Context, id int) error
}

type ProjectService struct {
	repo repository.ProjectRepository
}

func NewProjectService(r repository.ProjectRepository) *ProjectService {
	return &ProjectService{repo: r}
}

func (s *ProjectService) Create(ctx context.Context, p *models.Project) error {
 return s.repo.Create(ctx, p)
}

func (s *ProjectService) List(ctx context.Context, page, limit int) ([]models.Project, error) {
 offset := (page - 1) * limit
 return s.repo.List(ctx, limit, offset)
}

func (s *ProjectService) GetByID(ctx context.Context, id int) (*models.Project, error) {
 return s.repo.GetByID(ctx, id)
}

func (s *ProjectService) Update(ctx context.Context, id int, p *models.Project) error {
 return s.repo.Update(ctx, id, p)
}

func (s *ProjectService) Delete(ctx context.Context, id int) error {
 return s.repo.Delete(ctx, id)
}
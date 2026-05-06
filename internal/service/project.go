package service

import (
 "context"
 "task-api/internal/models"
 "task-api/internal/repository"
)

type ProjectService struct {
 repo *repository.ProjectRepo
}

func NewProjectService(r *repository.ProjectRepo) *ProjectService {
 return &ProjectService{r}
}

func (s *ProjectService) Create(ctx context.Context, p *models.Project) error {
 return s.repo.Create(ctx, p)
}

func (s *ProjectService) List(ctx context.Context, page, limit int) ([]models.Project, error) {
 offset := (page - 1) * limit
 return s.repo.List(ctx, limit, offset)
}
package repository

import (
 "context"
 "task-api/internal/models"

 "github.com/jackc/pgx/v5/pgxpool"
)

type ProjectRepo struct {
 db *pgxpool.Pool
}

func NewProjectRepo(db *pgxpool.Pool) *ProjectRepo {
 return &ProjectRepo{db}
}

func (r *ProjectRepo) Create(ctx context.Context, p *models.Project) error {
 return r.db.QueryRow(ctx,
  `INSERT INTO projects(name, description) VALUES($1,$2) RETURNING id`,
  p.Name, p.Description,
 ).Scan(&p.ID)
}

func (r *ProjectRepo) List(ctx context.Context, limit, offset int) ([]models.Project, error) {
 rows, err := r.db.Query(ctx,
  `SELECT id,name,description FROM projects LIMIT $1 OFFSET $2`,
  limit, offset,
 )
 if err != nil {
  return nil, err
 }
 defer rows.Close()

 var res []models.Project
 for rows.Next() {
  var p models.Project
  rows.Scan(&p.ID, &p.Name, &p.Description)
  res = append(res, p)
 }
 return res, nil
}
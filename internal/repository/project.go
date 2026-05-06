package repository

import (
 "context"
 "go-task/internal/models"

 "github.com/jackc/pgx/v5/pgxpool"
)

type ProjectRepository interface {
	Create(ctx context.Context, p *models.Project) error
	List(ctx context.Context, limit, offset int) ([]models.Project, error)
	GetByID(ctx context.Context, id int) (*models.Project, error)
	Update(ctx context.Context, id int, p *models.Project) error
	Delete(ctx context.Context, id int) error
}

type ProjectRepo struct {
	db *pgxpool.Pool
}

func NewProjectRepo(db *pgxpool.Pool) *ProjectRepo {
	return &ProjectRepo{db: db}
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

func (r *ProjectRepo) GetByID(ctx context.Context, id int) (*models.Project, error) {
 var p models.Project
 err := r.db.QueryRow(ctx,
  `SELECT id, name, description FROM projects WHERE id=$1`,
  id,
 ).Scan(&p.ID, &p.Name, &p.Description)

 if err != nil {
  return nil, err
 }
 return &p, nil
}

func (r *ProjectRepo) Update(ctx context.Context, id int, p *models.Project) error {
 _, err := r.db.Exec(ctx,
  `UPDATE projects SET name=$1, description=$2 WHERE id=$3`,
  p.Name, p.Description, id,
 )
 return err
}

func (r *ProjectRepo) Delete(ctx context.Context, id int) error {
 _, err := r.db.Exec(ctx,
  `DELETE FROM projects WHERE id=$1`,
  id,
 )
 return err
}
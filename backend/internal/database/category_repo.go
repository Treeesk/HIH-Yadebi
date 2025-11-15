package database

import (
	"context"
	"database/sql"
	"hih-yadebi-backend/internal/models"
)

type CategoryRepository struct {
	DB *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{DB: db}
}

func (r *CategoryRepository) CreateCategory(ctx context.Context, cat *models.Category) error {
	query := `INSERT INTO categories (title) VALUES ($1) RETURNING id`
	return r.DB.QueryRowContext(ctx, query, cat.Title).Scan(&cat.ID)
}

func (r *CategoryRepository) GetAll(ctx context.Context) ([]models.Category, error) {
	rows, err := r.DB.QueryContext(ctx, `SELECT id, title FROM categories`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.Category
	for rows.Next() {
		var c models.Category
		if err := rows.Scan(&c.ID, &c.Title); err != nil {
			return nil, err
		}
		list = append(list, c)
	}
	return list, nil
}

func (r *CategoryRepository) GetCategoryByID(ctx context.Context, id int) (*models.Category, error){
	category := &models.Category{}

	query := `
		SELECT id, title
		FROM categories WHERE id = $1
	`

	row := r.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&category.ID, &category.Title,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return category, err
}

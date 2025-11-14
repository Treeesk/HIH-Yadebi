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
	query := `INSERT INTO categories (name) VALUES ($1) RETURNING id`
	return r.DB.QueryRowContext(ctx, query, cat.Name).Scan(&cat.ID)
}

func (r *CategoryRepository) GetAll(ctx context.Context) ([]models.Category, error) {
	rows, err := r.DB.QueryContext(ctx, `SELECT id, name FROM categories`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.Category
	for rows.Next() {
		var c models.Category
		if err := rows.Scan(&c.ID, &c.Name); err != nil {
			return nil, err
		}
		list = append(list, c)
	}
	return list, nil
}

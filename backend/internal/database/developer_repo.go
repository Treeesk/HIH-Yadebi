package database

import (
	"context"
	"database/sql"
	"hih-yadebi-backend/internal/models"
)

type DeveloperRepository struct {
	DB *sql.DB
}

func NewDeveloperRepository(db *sql.DB) *DeveloperRepository {
	return &DeveloperRepository{DB: db}
}

func (r *DeveloperRepository) CreateDeveloper(ctx context.Context, dev *models.Developer) error {
	query := `
		INSERT INTO developers (name, password_hash, email)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	return r.DB.QueryRowContext(ctx, query,
		dev.Name,
		dev.PasswordHash,
		dev.Email,
	).Scan(&dev.ID)
}

func (r *DeveloperRepository) GetDeveloperByLogin(ctx context.Context, name string) (*models.Developer, error) {
	dev := &models.Developer{}

	query := `
		SELECT id, name, password_hash, email 
		FROM developers WHERE name = $1
	`

	row := r.DB.QueryRowContext(ctx, query, name)
	err := row.Scan(&dev.ID, &dev.Name, &dev.PasswordHash, &dev.Email)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	return dev, err
}

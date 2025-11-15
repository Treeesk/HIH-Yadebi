package database

import (
	"context"
	"database/sql"
	"hih-yadebi-backend/internal/models"
	// "time"
)

type DeveloperRepository struct {
	DB *sql.DB
}

func NewDeveloperRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateDeveloper(ctx context.Context, developer *models.Developer) error {
	query := `
		INSERT INTO users (name, password, email)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	return r.DB.QueryRowContext(ctx, query,
		developer.Name,
		developer.Password,
		developer.Email,
	).Scan(&developer.ID)
}

func (r *UserRepository) GetDeveloperByLogin(ctx context.Context, name string) (*models.Developer, error) {
	developer := &models.Developer{}

	query := `
		SELECT id, name, password, email 
		FROM users WHERE name = $1
	`

	row := r.DB.QueryRowContext(ctx, query, name)
	err := row.Scan(&developer.ID, &developer.Name, &developer.Password, &developer.Email)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	return developer, err
}

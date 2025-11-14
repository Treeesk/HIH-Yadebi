package database

import (
	"context"
	"database/sql"
	"backend/internal/models"
	// "time"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	query := `
		INSERT INTO users (name, password, email)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	return r.DB.QueryRowContext(ctx, query,
		user.Name,
		user.Password,
		user.Email,
	).Scan(&user.ID)
}

func (r *UserRepository) GetUserByLogin(ctx context.Context, name string) (*models.User, error) {
	user := &models.User{}

	query := `
		SELECT id, name, password, email 
		FROM users WHERE name = $1
	`

	row := r.DB.QueryRowContext(ctx, query, name)
	err := row.Scan(&user.ID, &user.name, &user.Password, &user.Email)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	return user, err
}

package database

import (
	"context"
	"database/sql"
	"errors"
	"hih-yadebi-backend/internal/models"
	"time"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) UserExists(ctx context.Context, email, name string) (bool, error) {
	query := `
        SELECT EXISTS(
            SELECT 1 FROM users WHERE email = $1 OR name = $2
        )
    `

	var exists bool
	err := r.DB.QueryRowContext(ctx, query, email, name).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	query := `
        INSERT INTO users (email, name, password, created_at)
        VALUES ($1, $2, $3, $4)
        RETURNING id
    `

	err := r.DB.QueryRowContext(
		ctx,
		query,
		user.Email,
		user.Name,
		user.Password,
		time.Now(),
	).Scan(&user.ID)

	return err
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `
        SELECT id, email, name, password, created_at
        FROM users
        WHERE email = $1
    `

	row := r.DB.QueryRowContext(ctx, query, email)

	user := models.User{}
	err := row.Scan(&user.ID, &user.Email, &user.Name, &user.Password, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

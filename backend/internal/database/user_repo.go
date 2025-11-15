package database

import (
	"context"
	"database/sql"
	// "fmt"
	"hih-yadebi-backend/internal/models"
	// "log"
	"time"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	query := "INSERT INTO users (username, password_hash, created_at) VALUES ($1, $2, $3) RETURNING id, created_at"
	err := r.DB.QueryRowContext(ctx, query, user.Name, user.PasswordHash, time.Now()).Scan(&user.ID, &user.CreatedAt)
	return err
}

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	user := &models.User{}
	query := "SELECT id, username, password_hash, created_at FROM users WHERE username=$1"
	row := r.DB.QueryRowContext(ctx, query, username)
	err := row.Scan(&user.ID, &user.Name, &user.PasswordHash, &user.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return user, err
}

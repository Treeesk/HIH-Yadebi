package database

import (
	"context"
	"database/sql"
	"hih-yadebi-backend/internal/models"
)

type UserAppsRepository struct {
	DB *sql.DB
}

func NewUserAppsRepository(db *sql.DB) *UserAppsRepository {
	return &UserAppsRepository{DB: db}
}

func (r *UserAppsRepository) AddInstalled(ctx context.Context, ua *models.UserApp) error {
	query := `
		INSERT INTO user_apps (user_id, app_id)
		VALUES ($1, $2)
		RETURNING id
	`
	return r.DB.QueryRowContext(ctx, query,
		ua.UserID,
		ua.AppID,
	).Scan(&ua.ID)
}

func (r *UserAppsRepository) GetUserApps(ctx context.Context, userID int) ([]models.UserApp, error) {
	rows, err := r.DB.QueryContext(ctx, `
		SELECT id, user_id, app_id
		FROM user_apps WHERE user_id = $1
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.UserApp
	for rows.Next() {
		var ua models.UserApp
		if err := rows.Scan(&ua.ID, &ua.UserID, &ua.AppID); err != nil {
			return nil, err
		}
		list = append(list, ua)
	}

	return list, nil
}

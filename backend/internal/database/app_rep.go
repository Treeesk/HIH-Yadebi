package database

import (
	"context"
	"database/sql"
	"hih-yadebi-backend/internal/models"
)

type AppRepository struct {
	DB *sql.DB
}

func NewAppRepository(db *sql.DB) *AppRepository {
	return &AppRepository{DB: db}
}

func (r *AppRepository) CreateApp(ctx context.Context, app *models.App) error {
	query := `
		INSERT INTO apps (title, description, size_mb, age_rating, downloads, version, link_apk, developer_id, category_id)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
		RETURNING id
	`

	return r.DB.QueryRowContext(ctx, query,
		app.Title,
		app.Description,
		app.SizeMB,
		app.AgeRating,
		app.Downloads,
		// app.Rating,
		app.Version,
		app.LinkApk,
		app.DeveloperID,
		app.CategoryID,
	).Scan(&app.ID)
}

func (r *AppRepository) GetApp(ctx context.Context, id int) (*models.App, error) {
	app := &models.App{}

	query := `
		SELECT id, title, description, size_mb, age_rating, downloads, version, link_apk, developer_id, category_id
		FROM apps WHERE id = $1
	`

	row := r.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&app.ID, &app.Title, &app.Description, &app.SizeMB,
		&app.AgeRating, &app.Downloads, &app.Version, &app.LinkApk,
		&app.DeveloperID, &app.CategoryID,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return app, err
}

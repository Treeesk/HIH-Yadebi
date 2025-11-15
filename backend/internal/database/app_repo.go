package database

import (
	"context"
	"database/sql"
	"hih-yadebi-backend/internal/models"
	"time"
)

type AppRepository struct {
	DB *sql.DB
}

func NewAppRepository(db *sql.DB) *AppRepository {
	return &AppRepository{DB: db}
}

func (r *AppRepository) CreateApp(ctx context.Context, app *models.App) error {
	query := `
		INSERT INTO apps (
			title, description, size_mb, age_rating, downloads,
			version, link_cloud, developer_id, category_id, created_at
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
		RETURNING id, created_at
	`

	return r.DB.QueryRowContext(ctx, query,
		app.Title,
		app.Description,
		app.SizeMB,
		app.AgeRating,
		app.Downloads,
		app.Version,
		app.LinkCloud,
		app.DeveloperID,
		app.CategoryID,
		time.Now(),
	).Scan(&app.ID, &app.CreatedAt)
}

func (r *AppRepository) GetAppByID(ctx context.Context, id int) (*models.App, error) {
	app := &models.App{}

	query := `
		SELECT id, title, description, size_mb, age_rating, downloads,
		       version, link_cloud, developer_id, category_id, created_at
		FROM apps WHERE id = $1
	`

	row := r.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&app.ID, &app.Title, &app.Description, &app.SizeMB,
		&app.AgeRating, &app.Downloads, &app.Version, &app.LinkCloud, 
		&app.DeveloperID, &app.CategoryID, &app.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return app, err
}

func (r *AppRepository) GetAppsByCategoryID(ctx context.Context, categoryID int) ([]*models.App, error) {
	apps := make([]*models.App, 0)

	query := `
		SELECT id, title, description, size_mb, age_rating, downloads,
		       version, link_cloud, developer_id, category_id, created_at
		FROM apps WHERE category_id = $1
	`

	rows, err := r.DB.QueryContext(ctx, query, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var app models.App
		if err := rows.Scan(
			&app.ID, &app.Title, &app.Description, &app.SizeMB,
			&app.AgeRating, &app.Downloads, &app.Version, &app.LinkCloud,
			&app.DeveloperID, &app.CategoryID, &app.CreatedAt,
		); err != nil {
			return nil, err
		}
		apps = append(apps, &app)
	}

	// Проверка на пустой результат
	if len(apps) == 0 {
		return nil, nil
	}

	return apps, nil
}

func (r *AppRepository) IncrementDownloadCount(ctx context.Context, id int) (error) {
	_, err := r.DB.ExecContext(ctx, `
		UPDATE apps SET downloads = downloads + 1 WHERE id = $1
	`, id)
	return err
} 


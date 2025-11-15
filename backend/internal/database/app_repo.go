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
	query := `INSERT INTO apps (
		title, description, size_mb, age_rating, downloads,
		version, link_apk, developer_id, category_id, created_at
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
		app.LinkApp,
		app.DeveloperID,
		app.CategoryID,
		time.Now(),
	).Scan(&app.ID, &app.CreatedAt)
}

func (r *AppRepository) GetAppByID(ctx context.Context, id int) (*models.App, error) {
	app := &models.App{}

	query := `
		SELECT id, title, description, size_mb, age_rating, downloads,
		       version, link_app, developer_id, category_id, created_at
		FROM apps WHERE id = $1
	`

	row := r.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&app.ID, &app.Title, &app.Description, &app.SizeMB,
		&app.AgeRating, &app.Downloads, &app.Version, &app.LinkApp,
		&app.DeveloperID, &app.CategoryID, &app.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return app, err
}

//	func (r *AppRepository) GetAppByCategoryID(ctx context.Context, category_id int) ([]*models.App, error) {
//		apps := make([]*models.App, 0)
//		query := ""
//		rows, err := r.DB.QueryContext(ctx, query, category_id)
//		if err != nil {
//			if err == sql.ErrNoRows {
//				return nil, nil
//			}
//			return nil, err
//		}
//		for row.Next() {
//			var app models.App
//			err := rows.Scan(&app.ID, &app.Title, &app.CategoryID) // хюйня с маленькой иконкой была
//		}
//	}
func (r *AppRepository) GetPopularApps(limit int) ([]int, error) {

	query := `
        SELECT a.id
        FROM apps a
        LEFT JOIN reviews r ON a.id = r.app_id
        GROUP BY a.id
        ORDER BY 
            COALESCE(AVG(r.score), 0) DESC,  -- 1. Средний рейтинг
            a.downloads DESC,                -- 2. Загрузки
            a.created_at DESC                -- 3. Новые
        LIMIT $1;
    `

	rows, err := r.DB.Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var top []int

	for rows.Next() {
		var top_el int
		if scanErr := rows.Scan(&top_el); scanErr != nil {
			return nil, scanErr
		}
		top = append(top, top_el)
	}

	return top, nil
}
func (r *AppRepository) GetNewApps(limit int) ([]int, error) {

	query := `
        SELECT a.id
        FROM apps a
        LEFT JOIN reviews r ON a.id = r.app_id
        GROUP BY a.id
        ORDER BY 
            a.created_at DESC                -- 3. Новые
        LIMIT $1;
    `

	rows, err := r.DB.Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var top_date []int

	for rows.Next() {
		var top_date_el int
		if scanErr := rows.Scan(&top_date_el); scanErr != nil {
			return nil, scanErr
		}
		top_date = append(top_date, top_date_el)
	}

	return top_date, nil
}

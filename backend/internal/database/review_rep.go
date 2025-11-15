package database

import (
	"context"
	"database/sql"
	"hih-yadebi-backend/internal/models"
)

type ReviewRepository struct {
	DB *sql.DB
}

func NewReviewRepository(db *sql.DB) *ReviewRepository {
	return &ReviewRepository{DB: db}
}

func (r *ReviewRepository) CreateReview(ctx context.Context, review *models.Review) error {
	query := `
		INSERT INTO reviews (user_id, app_id, score, comment)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	return r.DB.QueryRowContext(ctx, query,
		review.UserID,
		review.AppID,
		review.Score,
		review.Comment,
		// review.Helpful,
	).Scan(&review.ID)
}

func (r *ReviewRepository) GetReviewsByApp(ctx context.Context, appID int) ([]models.Review, error) {
	rows, err := r.DB.QueryContext(ctx, `
		SELECT id, user_id, app_id, score, comment
		FROM reviews WHERE app_id = $1
	`, appID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.Review
	for rows.Next() {
		var r models.Review
		if err := rows.Scan(&r.ID, &r.UserID, &r.AppID, &r.Score, &r.Comment); err != nil {
			return nil, err
		}
		list = append(list, r)
	}

	return list, nil
}

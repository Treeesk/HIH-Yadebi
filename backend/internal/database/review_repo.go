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
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at
	`

	return r.DB.QueryRowContext(ctx, query,
		review.UserID,
		review.AppID,
		review.Score,
		review.Comment,
	).Scan(&review.ID, &review.CreatedAt)
}

func (r *ReviewRepository) GetReviewsByApp(ctx context.Context, appID int) ([]models.Review, error) {
	rows, err := r.DB.QueryContext(ctx, `
		SELECT id, user_id, app_id, score, comment, created_at
		FROM reviews WHERE app_id = $1
		ORDER BY created_at DESC
	`, appID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []models.Review
	for rows.Next() {
		var review models.Review
		if err := rows.Scan(
			&review.ID,
			&review.UserID,
			&review.AppID,
			&review.Score,
			&review.Comment,
			&review.CreatedAt,
		); err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}

	return reviews, nil
}

package database

import (
	"database/sql"
)

type AppsRepository struct {
	DB *sql.DB
}

func NewAppsRepository(db *sql.DB) *AppsRepository {
	return &AppsRepository{DB: db}
}

func (r *AppsRepository) GetPopularApps(limit int) ([]int, error) {

	query := `
        SELECT id
        FROM apps
        ORDER BY 
            rating DESC,       
            downloads DESC,     
            created_at DESC     
        LIMIT $1
    `

	rows, err := r.DB.Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ids []int

	for rows.Next() {
		var id int
		if err_rows := rows.Scan(&id); err_rows != nil {
			return nil, err_rows
		}
		ids = append(ids, id)
	}

	return ids, nil
}

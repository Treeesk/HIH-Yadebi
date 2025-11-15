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

	var ids_top []int

	for rows.Next() {
		var id int
		if err_rows := rows.Scan(&id); err_rows != nil {
			return nil, err_rows
		}
		ids_top = append(ids_top, id)
	}

	return ids_top, nil
}
func (r *AppsRepository) GetNewApps(limit int) ([]int, error) {
	query := `
        SELECT id
        FROM apps
        ORDER BY 
            created_at DESC     
        LIMIT $1
    `
	rows, err := r.DB.Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var top_date []int
	for rows.Next() {
		var topdate_el int
		if err_rows := rows.Scan(&topdate_el); err_rows != nil {
			return nil, err_rows
		}
		top_date = append(top_date, topdate_el)
	}
	return top_date, nil
}

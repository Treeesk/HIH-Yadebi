package database

import (
	"backend/internal/config"
	"database/sql"
	"fmt"
	"log"
)

func Connect(cfg *config.ConfigDB) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("postgresql open error: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("postgresql ping error: %w", err)
	}

	// параметры пула соединений
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	log.Println("postgresql successfully connected")
	return db, nil
}

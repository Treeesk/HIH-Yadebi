package main

import (
	"hih-yadebi-backend/internal/config"
	"hih-yadebi-backend/internal/database"
	// "hih-yadebi-backend/internal/handlers"
	"hih-yadebi-backend/internal/routes"
	"log"
)

func main() {
	log.Println("Testing config and database...")

	// Просто загружаем конфиг и подключаемся к БД
	dbConfig := config.LoadDB()
	db := database.Connect(dbConfig)
	database.ApplyMigrations(db, "internal/migrations")
	defer db.Close()
	reviewRepo := database.NewReviewRepository(db)
	userRepo := database.NewUserRepository(db)
	appRepo := database.NewAppRepository(db)
	r := routes.SetupRouter(userRepo, appRepo, reviewRepo)

	log.Println("✅ Everything works!")
	r.Run("0.0.0.0:8080")
}

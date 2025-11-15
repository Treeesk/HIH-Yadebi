package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"hih-yadebi-backend/internal/config"
	"hih-yadebi-backend/internal/database"
	"hih-yadebi-backend/internal/handlers"
)

func main() {
	log.Println("Starting server...")

	// 1. Загружаем конфиг
	dbConfig := config.LoadDB()

	// 2. Подключаемся к базе
	db := database.Connect(dbConfig)
	defer db.Close()

	// 3. Инициализируем репозиторий
	appsRepo := database.NewAppRepository(db)

	// 4. Инициализируем хендлер
	appsHandler := handlers.NewAppsHandler(appsRepo)

	// 5. Gin router
	r := gin.Default()

	// 6. Маршрут для проверки apps_handler

	r.GET("/apps/new", appsHandler.GetNewApp) // получить новые
	r.POST("/apps", appsHandler.Create)       // создать приложение

	r.Run(":8080")

	// 7. Запускаем сервер
}

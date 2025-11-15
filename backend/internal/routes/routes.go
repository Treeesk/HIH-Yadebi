package routes

import (
	"hih-yadebi-backend/internal/database"
	"hih-yadebi-backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userRepo *database.UserRepository, appRepo *database.AppRepository) *gin.Engine {
	router := gin.Default()

	authHandler := handlers.NewAuthHandler(userRepo)
	appsHandler := handlers.NewAppHandler(appRepo)

	// Auth routes
	router.POST("/login", authHandler.Login)
	router.POST("/register", authHandler.Register)

	// Apps routes
	router.GET("/apps/popular", appsHandler.GetPopular)
	router.GET("/apps/new", appsHandler.GetNewApp)
	return router
}

package routes

import (
	"hih-yadebi-backend/internal/database"
	"hih-yadebi-backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userRepo *database.UserRepository, appRepo *database.AppRepository, reviewRepo *database.ReviewRepository) *gin.Engine {
	router := gin.Default()

	appsHandler := handlers.NewAppHandler(appRepo, reviewRepo)

	// Auth routes

	// Apps routes
	router.GET("/apps/:id/reviews", appsHandler.GetAppReviews)
	router.GET("/apps/popular", appsHandler.GetPopular)
	router.GET("/apps/new", appsHandler.GetDateApps)
	return router
}

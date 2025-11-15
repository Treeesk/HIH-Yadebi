package routes

import (
	"github.com/gin-gonic/gin"
	"hih-yadebi-backend/internal/handlers"
)

func SetupRoutes(
	r *gin.Engine,
	authHandler *handlers.AuthHandler,
	storeHandler *handlers.StoreHandler,
	appHandler *handlers.AppHandler,
	profileHandler *handlers.ProfileHandler,
) {

	// =============== AUTH ==================

	r.POST("/login", authHandler.Login)
	r.POST("/register", authHandler.Register)
	r.POST("/register/confirmEmail", authHandler.ConfirmEmail)

	// =============== STORE ================

	r.GET("/mainStorePage", storeHandler.MainStorePage)
	r.GET("/getAppsByCategory", storeHandler.GetAppsByCategoryID)
	r.GET("/getCategories", storeHandler.GetCategories)

	// =============== APPS ==================

	r.GET("/getApp", appHandler.GetAppByID)
	r.GET("/getAppReviews", appHandler.GetAppReviews)
	r.POST("/addReview", appHandler.AddReview)

	// =============== PROFILE ===============

	r.GET("/getProfile", profileHandler.GetProfileByID)
}

func SetupRouter(
	authHandler *handlers.AuthHandler,
	storeHandler *handlers.StoreHandler,
	appHandler *handlers.AppHandler,
	profileHandler *handlers.ProfileHandler,
) *gin.Engine {

	r := gin.Default()

	SetupRoutes(
		r,
		authHandler,
		storeHandler,
		appHandler,
		profileHandler,
	)

	return r
}

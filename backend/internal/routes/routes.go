package routes

import "github.com/gin-gonic/gin"


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
    r.GET("/getAppsByCategory", storeHandler.GetAppsByCategory)
    r.GET("/getCategories", storeHandler.GetCategories)


    // =============== APPS ==================

    r.GET("/getApp", appHandler.GetApp)
    r.GET("/getAppReviews", appHandler.GetAppReviews)
    r.POST("/addReview", appHandler.AddReview)


    // =============== PROFILE ===============

    r.GET("/getProfile", profileHandler.GetProfile)
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

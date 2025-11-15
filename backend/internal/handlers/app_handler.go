package handlers

import (
	"hih-yadebi-backend/internal/database"
	// "time"

	"github.com/gin-gonic/gin"
)

type AppHandler struct {
	AppRepo *database.AppRepository
}

func NewAppHandler(repo *database.AppRepository) *AppHandler {
	return &AppHandler{AppRepo: repo}
}

type ReviewRequest struct {
	AppID   int    `json:"app_id" binding:"required"`
	UserID  int    `json:"user_id" binding:"required"`
	Score   int    `json:"score" binding:"required"`
	Comment string `json:"comment" binding:"required"`
	// CreatedAt time.Time `json:"" binding:"required"`
}

func (h *AppHandler) GetAppByID(c *gin.Context) {

}

func (h *AppHandler) GetAppReviews(c *gin.Context) {

}

func (h *AppHandler) AddReview(c *gin.Context) {

}

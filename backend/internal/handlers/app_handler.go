package handlers

import (
	"hih-yadebi-backend/internal/database"
	"time"

	"github.com/gin-gonic/gin"
)

type AppHandler struct {
	AppRepo *database.AppRepository
}

func NewAppHandler(repo *database.AppRepository) *AppHandler {
	return &AppHandler{AppRepo: repo}
}

type ReviewRequest struct {
	AppID     int
	UserID    int
	Score     int
	Comment   string
	CreatedAt time.Time
}

func (h *AppHandler) GetAppByID(c *gin.Context) {

}

func (h *AppHandler) GetAppReviews(c *gin.Context) {

}

func (h *AppHandler) AddReview(c *gin.Context) {

}

package handlers

import (
	"hih-yadebi-backend/internal/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AppHandler struct {
	AppRepo    *database.AppRepository
	ReviewRepo *database.ReviewRepository
}

func NewAppHandler(repo *database.AppRepository, reviewRepo *database.ReviewRepository) *AppHandler {
	return &AppHandler{AppRepo: repo,
		ReviewRepo: reviewRepo,
	}
}
func (h *AppHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	ctx := c.Request.Context()

	app, err := h.AppRepo.GetAppByID(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
		return
	}

	if app == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "app not found"})
		return
	}

	c.JSON(http.StatusOK, app)
}

func (h *AppHandler) GetPopular(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	limit, _ := strconv.Atoi(limitStr)
	ids, err := h.AppRepo.GetPopularApps(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ids": ids})
}
func (h *AppHandler) GetDateApps(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	limit, _ := strconv.Atoi(limitStr)
	top_date, err := h.AppRepo.GetDateApps(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"top_date": top_date})
}
func (h *AppHandler) GetAppReviews(c *gin.Context) {
	idStr := c.Query("app_id")
	appID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	ctx := c.Request.Context()
	reviews, err := h.ReviewRepo.GetReviewsByApp(ctx, appID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reviews": reviews})
}

package handlers

import (
	"hih-yadebi-backend/internal/database"
	"hih-yadebi-backend/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AppsHandler struct {
	Repo *database.AppRepository
}

func NewAppsHandler(repo *database.AppRepository) *AppsHandler {
	return &AppsHandler{Repo: repo}
}
func (h *AppsHandler) Create(c *gin.Context) {
	var req models.App

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	// вызывать обязательно через контекст
	ctx := c.Request.Context()

	// важно: передаём указатель
	err := h.Repo.CreateApp(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
		return
	}

	// req.ID уже заполнен QueryRowContext
	c.JSON(http.StatusOK, gin.H{
		"status": "created",
		"id":     req.ID,
	})
}

func (h *AppsHandler) GetPopular(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	limit, _ := strconv.Atoi(limitStr)
	ids, err := h.Repo.GetPopularApps(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ids": ids})
}
func (h *AppsHandler) GetNewApp(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	limit, _ := strconv.Atoi(limitStr)
	top_date, err := h.Repo.GetNewApps(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"top_date": top_date})
}

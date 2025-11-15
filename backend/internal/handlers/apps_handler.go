package handlers

import (
	"github.com/gin-gonic/gin"
	"hih-yadebi-backend/internal/database"
	"net/http"
	"strconv"
)

type AppsHandler struct {
	Repo *database.AppsRepository
}

func NewAppsHandler(repo *database.AppsRepository) *AppsHandler {
	return &AppsHandler{Repo: repo}
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

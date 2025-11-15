package handlers

import (
	"hih-yadebi-backend/internal/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AppHandler struct {
	AppRepo *database.AppRepository
}

func NewAppHandler(repo *database.AppRepository) *AppHandler {
	return &AppHandler{AppRepo: repo}
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
func (h *AppHandler) GetNewApp(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	limit, _ := strconv.Atoi(limitStr)
	top_date, err := h.AppRepo.GetNewApps(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"top_date": top_date})
}

package handlers

import (
	"context"
	"hih-yadebi-backend/internal/database"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type StoreHandler struct {
	AppRepo      *database.AppRepository
	CategoryRepo *database.CategoryRepository
}

func NewStoreHandler(appRepo *database.AppRepository, categoryRepo *database.CategoryRepository) *StoreHandler {
	return &StoreHandler{AppRepo: appRepo, CategoryRepo: categoryRepo}
}

func (h *StoreHandler) MainStorePage(c *gin.Context) {

}

func (h *StoreHandler) GetAppsByCategoryID(c *gin.Context) {
	category_id := c.Query("category_id")
	if category_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	category_id_int, err := strconv.Atoi(category_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	apps, err := h.AppRepo.GetAppByCategoryID(ctx, category_id_int)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server-side error"})
		return
	}
	if apps == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "category not found or there are no apps with this category"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"apps": apps})
}

func (h *StoreHandler) GetCategories(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	categories, err := h.CategoryRepo.GetAllCategories(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server-side error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"categories": categories})
}

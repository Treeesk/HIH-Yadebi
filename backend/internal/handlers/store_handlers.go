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
	AppRepo *database.AppRepository
}

func NewStoreHandler(repo *database.AppRepository) *StoreHandler {
	return &StoreHandler{AppRepo: repo}
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
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	apps, err := h.AppRepo.GetAppByCategoryID(ctx, category_id_int)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server-side error"})
		return
	}i
}
func (h *StoreHandler) GetCategories(c *gin.Context) {

}

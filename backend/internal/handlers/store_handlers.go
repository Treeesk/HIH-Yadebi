package handlers

import (
	"hih-yadebi-backend/internal/database"

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

}
func (h *StoreHandler) GetCategories(c *gin.Context) {

}

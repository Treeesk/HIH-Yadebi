package handlers

import (
	"context"
	"hih-yadebi-backend/internal/database"
	"hih-yadebi-backend/internal/models"
	"math/rand"
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
	limitApps := c.DefaultQuery("limit", "10")
	limitApps_int, err := strconv.Atoi(limitApps)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	ctx1, cancel1 := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel1()
	mostPopularApps, err := h.AppRepo.GetMostPopularApps(ctx1, limitApps_int)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server-side error"})
		return
	}
	ctx2, cancel2 := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel2()
	mostNewApps, err := h.AppRepo.GetMostNewApps(ctx2, limitApps_int)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server-side error"})
		return
	}
	respBody := make(map[string][]*models.App)
	respBody["Самые популярные приложения"] = mostPopularApps
	respBody["Новинки"] = mostNewApps
	ctx3, cancel3 := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel3()
	categories, err := h.CategoryRepo.GetAllCategories(ctx3)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server-side error"})
		return
	}
	rand.Shuffle(len(categories), func(i, j int) {
		categories[i], categories[j] = categories[j], categories[i]
	})
	// type categoryWApps struct{
	// 	name string
	// 	apps []*models.App
	// }
	for _, v := range categories {
		ctx4, cancel4 := context.WithTimeout(context.Background(), 5*time.Second)
		cancel4()
		appsByCategory, err := h.AppRepo.GetAppByCategoryID(ctx4, v.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "server-side error"})
			return
		}
		respBody[v.Title] = appsByCategory
	}
	c.JSON(http.StatusOK, gin.H{"categories": respBody})
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

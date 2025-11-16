package handlers

import (
	"context"
	"hih-yadebi-backend/internal/database"
	// "hih-yadebi-backend/internal/models"
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

type StorePageResponse struct {
	Categories []Category `json:"categories"`
}

type Category struct {
	CategoryName string      `json:"category_name"`
	Apps         []LittleApp `json:"apps"`
}

func NewStoreHandler(appRepo *database.AppRepository, categoryRepo *database.CategoryRepository) *StoreHandler {
	return &StoreHandler{AppRepo: appRepo, CategoryRepo: categoryRepo}
}

type LittleApp struct {
	App_id        int    `json:"app_id"`
	App_name      string `json:"app_name"`
	App_category  string `json:"app_category"`
	App_linkCloud string `json:"app_link_cloud"`
}

func (h *StoreHandler) MainStorePage(c *gin.Context) {
	limitApps := c.DefaultQuery("limit", "10")
	limitApps_int, err := strconv.Atoi(limitApps)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	ctx1, cancel1 := context.WithTimeout(context.Background(), 5*time.Second)
	mostPopularApps, err := h.AppRepo.GetMostPopularApps(ctx1, limitApps_int)
	defer cancel1()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server-side error"})
		return
	}
	ctx2, cancel2 := context.WithTimeout(context.Background(), 5*time.Second)
	mostNewApps, err := h.AppRepo.GetMostNewApps(ctx2, limitApps_int)
	defer cancel2()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server-side error"})
		return
	}
	respBody := []Category{}
	ctx3, cancel3 := context.WithTimeout(context.Background(), 5*time.Second)
	categories, err := h.CategoryRepo.GetAllCategories(ctx3)
	defer cancel3()
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
		appsByCategory, err := h.AppRepo.GetAppsByCategoryID(ctx4, v.ID)
		defer cancel4()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "server-side error"})
			return
		}
		littleAppsByCategory := Category{}
		littleAppsByCategory.CategoryName = v.Title
		for _, vv := range appsByCategory {
			littleAppsByCategory.Apps = append(littleAppsByCategory.Apps,
				LittleApp{App_id: vv.ID, App_name: vv.Title, App_category: v.Title, App_linkCloud: vv.LinkCloud})
		}
		respBody = append(respBody, littleAppsByCategory)
	}
	littleMostPopularApps := Category{}
	for _, v := range mostPopularApps {
		categoryTitle, _ := h.CategoryRepo.GetCategoryByID(context.Background(), v.CategoryID)
		littleMostPopularApps.Apps = append(littleMostPopularApps.Apps,
			LittleApp{App_id: v.ID, App_name: v.Title, App_category: categoryTitle.Title, App_linkCloud: v.LinkCloud})
	}
	littleMostPopularApps.CategoryName = "Самые популярные приложения"
	respBody = append(respBody, littleMostPopularApps)
	littleMostNewApps := Category{}
	for _, v := range mostNewApps {
		categoryTitle, _ := h.CategoryRepo.GetCategoryByID(context.Background(), v.CategoryID)
		littleMostNewApps.Apps = append(littleMostNewApps.Apps,
			LittleApp{App_id: v.ID, App_name: v.Title, App_category: categoryTitle.Title, App_linkCloud: v.LinkCloud})
	}
	// respBody["Новинки"] = littleMostNewApps
	littleMostNewApps.CategoryName = "Новинки"
	respBody = append(respBody, littleMostNewApps)
	// respBody = append(respBody, map[string][]*LittleApp{"Новинки": littleMostNewApps})
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
	apps, err := h.AppRepo.GetAppsByCategoryID(ctx, category_id_int)
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

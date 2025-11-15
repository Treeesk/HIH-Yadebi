package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"hih-yadebi-backend/internal/database"
	"hih-yadebi-backend/internal/models"
	"net/http"
	"time"
)

type RegRequest struct {
	Email    string `json:"email" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthHandler struct {
	Users *database.UserRepository
}

func NewAuthHandler(repo *database.UserRepository) *AuthHandler {
	return &AuthHandler{Users: repo}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req RegRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	exists, err := h.Users.UserExists(ctx, req.Email, req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
		return
	}
	if exists {
		c.JSON(http.StatusConflict, gin.H{"error": "user already exists"})
		return
	}

	user := models.User{
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Password,
	}

	err = h.Users.CreateUser(ctx, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"user":   user,
	})
}
func (h *AuthHandler) Login(c *gin.Context) {
	var req RegRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user, _ := h.Users.GetUserByEmail(ctx, req.Email)
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	if user.Password != req.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "wrong password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "login ok",
		"user":   user.Name,
	})
}

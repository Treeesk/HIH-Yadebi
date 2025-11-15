package handlers

import (
	"context"
	"hih-yadebi-backend/internal/database"
	"hih-yadebi-backend/internal/models"
	"hih-yadebi-backend/internal/utils"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	UserRepo *database.UserRepository
}

func NewAuthHandler(userRepo *database.UserRepository) *AuthHandler {
	return &AuthHandler{UserRepo: userRepo}
}

type AuthRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("invalid request: %v\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	existingUser, _ := h.UserRepo.GetUserByUsername(context.Background(), req.Username)
	if existingUser != nil {
		log.Printf("user with this username exists\n")
		c.JSON(http.StatusConflict, gin.H{"error": "user with this username already exists"})
		return
	}
	hash, err := utils.HashPassword(req.Password)
	if err != nil {
		log.Printf("hashing password error: %v\n", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server-side error"})
		return
	}
	user := models.User{
		Name:         req.Username,
		PasswordHash: hash}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = h.UserRepo.CreateUser(ctx, &user)
	if err != nil {
		log.Printf("creating user error: %v\n", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "creating user error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": user.ID, "username": user.Name, "created_at": user.CreatedAt})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("invalid request: %v\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user, err := h.UserRepo.GetUserByUsername(ctx, req.Username)
	if err != nil {
		log.Printf("database error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}
	if user == nil {
		log.Printf("user %v not found\n", req.Username)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	if !utils.CheckPasswordHash(req.Password, user.PasswordHash) {
		log.Printf("wrong password for user %v\n", req.Username)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid password"})
		return
	}

	token, err := utils.GenerateJWT(user.ID, user.Name)
	if err != nil {
		log.Printf("generation token error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "generation token error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"auth_token": token})
	// c.SetCookie("auth_token", token, 3600, "/", "", false, false)
}

func (h *AuthHandler) Profile(c *gin.Context) {
	userID := c.GetInt("user_id")
	username := c.GetString("username")
	c.JSON(http.StatusOK, gin.H{"user_id": userID, "username": username})
}

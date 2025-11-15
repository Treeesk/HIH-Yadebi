package handlers

import (
	"hih-yadebi-backend/internal/database"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	UserRepo *database.UserRepository
}

func NewAuthHandler(repo *database.UserRepository) *AuthHandler {
	return &AuthHandler{UserRepo: repo}
}

type AuthRequest struct {
	email    string
	password string
}

type RegisterRequest struct {
	email    string
	name     string
	password string
}

func (h *AuthHandler) Login(c *gin.Context) {

}

func (h *AuthHandler) Register(c *gin.Context) {

}

func (h *AuthHandler) ConfirmEmail(c *gin.Context) {

}

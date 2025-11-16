package handlers

import (
	"hih-yadebi-backend/internal/database"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	UserRepo *database.UserRepository
}

func NewProfileHandler(repo *database.UserRepository) *ProfileHandler {
	return &ProfileHandler{UserRepo: repo}
}

func (h *ProfileHandler) GetProfileByID(c *gin.Context) {

}

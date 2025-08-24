package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *AuthHandler) Me(c *gin.Context) {
	userID := c.GetString("userID")
	c.JSON(http.StatusOK, gin.H{"user_id": userID})
}

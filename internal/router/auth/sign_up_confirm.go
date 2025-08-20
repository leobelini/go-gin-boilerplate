package auth

import (
	"leobelini/cashly/internal/utils"

	"github.com/gin-gonic/gin"
)

func (h *AuthHandler) SignUpConfirm(c *gin.Context) {
	token := c.Param("token")

	if token == "" {
		utils.HandleValidationError(c, utils.CreateAppError("INVALID_TOKEN", false))
		return
	}

	c.JSON(200, gin.H{"message": "sign up confirmed"})
}

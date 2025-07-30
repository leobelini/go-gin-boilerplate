package auth

import (
	"leobelini/cashly/internal/utils"

	"github.com/gin-gonic/gin"
)

type VerifyAccountRequest struct {
	Token string `json:"token" binding:"required,uuid" `
}

func verifyAccount(c *gin.Context) {
	var req VerifyAccountRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleValidationError(c, err)
		return
	}
}

package auth

import (
	"leobelini/cashly/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *AuthHandler) Me(c *gin.Context) {
	userID := c.GetString("userID")
	ctx := c.Request.Context()

	user, err := h.controllers.User.GetUserById(userID, ctx)
	if err != nil {
		utils.HandleValidationError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": user.ID, "email": user.Email, "name": user.Name})
}

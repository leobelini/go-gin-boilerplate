package auth

import (
	"leobelini/cashly/internal/utils"

	"github.com/gin-gonic/gin"
)

type PasswordRecoveryRequest struct {
	Email string `json:"email" binding:"required,email"`
}

// passwordRecovery godoc
// @Summary      Inicia o processo de recuperação de senha
// @Description  Envia um email com instruções para redefinir a senha
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request  body      PasswordRecoveryRequest  true  "Email do usuário"
// @Success      200      {object}  map[string]string
// @Failure      400      {object}  api.ErrorResponse
// @Router       /auth/password-recovery [post]
func (h *AuthHandler) PasswordRecovery(c *gin.Context) {

	var req PasswordRecoveryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleValidationError(c, err)
		return
	}

	ctx := c.Request.Context()
	err := h.controllers.Auth.PasswordRecovery(req.Email, ctx)
	if err != nil {
		utils.HandleValidationError(c, err)
		return
	}

	c.JSON(200, gin.H{"message": "Password recovery email sent"})
}

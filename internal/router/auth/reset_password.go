package auth

import (
	"leobelini/cashly/internal/utils"

	"github.com/gin-gonic/gin"
)

type ResetPasswordRequest struct {
	Password string `json:"password" binding:"required,min=6"`
}

// resetPassword godoc
// @Summary      Redefine a senha do usuário
// @Description  Redefine a senha do usuário usando o token enviado por email
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        token    path      string                 true  "Token de redefinição de senha"
// @Param        request  body      ResetPasswordRequest   true  "Nova senha do usuário"
// @Success      200      {object}  map[string]string
// @Failure      400      {object}  api.ErrorResponse
// @Router       /auth/reset-password/{token} [put]
func (h *AuthHandler) ResetPassword(c *gin.Context) {

	var req ResetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleError(c, err)
		return
	}

	token := c.Param("token")

	ctx := c.Request.Context()
	err := h.controllers.Auth.ResetPassword(req.Password, token, ctx)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(200, gin.H{"message": "Password reset successful"})
}

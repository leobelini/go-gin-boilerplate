package auth

import (
	"leobelini/cashly/internal/utils"

	"github.com/gin-gonic/gin"
)

// signUpConfirm godoc
// @Summary      Confirma o cadastro de um usuário
// @Description  Confirma o cadastro de um usuário através de um token enviado por email
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        token  path      string  true  "Token de confirmação"
// @Success      200    {object}  map[string]string
// @Failure      400    {object}  api.ErrorResponse
// @Router       /auth/sign-up/confirm/{token} [put]
func (h *AuthHandler) SignUpConfirm(c *gin.Context) {
	token := c.Param("token")

	if token == "" {
		utils.HandleError(c, utils.CreateAppError("INVALID_TOKEN", false))
		return
	}

	if err := h.controllers.Auth.SignUpConfirm(token, c.Request.Context()); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(200, gin.H{"message": "sign up confirmed"})
}

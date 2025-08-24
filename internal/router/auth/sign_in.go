package auth

import (
	"leobelini/cashly/internal/utils"

	"github.com/gin-gonic/gin"
)

type SignInRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type SignInResponse struct {
	Token string `json:"token"`
}

// signIn godoc
// @Summary      Realiza o login do usuário
// @Description  Realiza o login do usuário e retorna um token JWT
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request  body      SignInRequest  true  "Credenciais de login"
// @Success      200      {object}  SignInResponse
// @Failure      400      {object}  api.ErrorResponse
// @Router       /auth/sign-in [post]
func (h *AuthHandler) SignIn(c *gin.Context) {
	var req SignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleError(c, err)
		return
	}

	token, err := h.controllers.Auth.SignIn(req.Email, req.Password, c.Request.Context())
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(200, SignInResponse{Token: token})
}

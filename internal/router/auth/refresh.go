package auth

import (
	"leobelini/cashly/internal/utils"

	"github.com/gin-gonic/gin"
)

// refresh godoc
// @Summary      Refresh do token JWT
// @Description  Gera um novo token JWT e um novo refresh token, baseado no refresh token atual
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      200  {object}  SignInResponse
// @Failure      400  {object}  api.ErrorResponse
// @Router       /auth/refresh [post]
func (h *AuthHandler) Refresh(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	userID, err := utils.ValidateJWT(refreshToken, h.controllers.App.Env.App.JWTSecret)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	token, newRefreshToken, err := utils.GenerateJWT(userID, h.controllers.App.Env.App.JWTSecret)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.SetCookie("refresh_token", newRefreshToken, 7*24*60*60, "/", "", false, true)

	c.JSON(200, SignInResponse{Token: token})

}

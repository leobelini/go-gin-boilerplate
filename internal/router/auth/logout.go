package auth

import "github.com/gin-gonic/gin"

// logout godoc
// @Summary      Logout do usuário
// @Description  Realiza o logout do usuário, removendo o cookie de refresh token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      200  {string}  string	"OK"
// @Failure      400  {object}  api.ErrorResponse
// @Router       /auth/logout [post]
func (h *AuthHandler) Logout(c *gin.Context) {
	c.SetCookie("refresh_token", "", -1, "/", "", false, true)
	c.Status(200)
}

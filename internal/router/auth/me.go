package auth

import (
	"leobelini/cashly/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MeResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

// me godoc
// @Summary      Retorna informações do usuário autenticado
// @Description  Retorna informações do usuário autenticado baseado no token JWT
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      200  {object}  MeResponse
// @Failure      400  {object}  api.ErrorResponse
// @Router       /auth/me [get]
func (h *AuthHandler) Me(c *gin.Context) {
	userID := c.GetString("userID")
	ctx := c.Request.Context()

	user, err := h.controllers.User.GetUserById(userID, ctx)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	resp := MeResponse{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}

	c.JSON(http.StatusOK, resp)
}

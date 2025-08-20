package user

import (
	"leobelini/cashly/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required" validate:"min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// createUser godoc
// @Summary      Cria um usuário
// @Description  Cria um usuário com nome, email e senha
// @Tags         usuários
// @Accept       json
// @Produce      json
// @Param        user  body      CreateUserRequest  true  "Dados do usuário"
// @Success      201   {object}  entity.User
// @Failure      400   {object}  api.ErrorResponse
// @Router       /user [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	ctx := c.Request.Context()

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleValidationError(c, err)
		return
	}

	if err := h.controllers.User.CreateUser(req.Name, req.Email, req.Password, ctx); err != nil {
		utils.HandleValidationError(c, err)
		return
	}

	user, err := h.controllers.User.GetUserByEmail(req.Email, ctx)
	if err != nil {
		utils.HandleValidationError(c, err)
		return
	}

	user.Password = ""
	user.Token = nil

	c.JSON(http.StatusCreated, user)
}

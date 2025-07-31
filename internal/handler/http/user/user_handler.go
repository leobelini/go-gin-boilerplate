package user

import (
	usecase "leobelini/cashly/internal/usecase/user"
	"leobelini/cashly/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required" validate:"min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UserHandler struct {
	userUseCase *usecase.UserUseCase
}

func NewUserHandler(userUseCase *usecase.UserUseCase) *UserHandler {
	return &UserHandler{userUseCase: userUseCase}
}

// createUser godoc
// @Summary      Cria um usuário
// @Description  Cria um usuário com nome, email e senha
// @Tags         usuários
// @Accept       json
// @Produce      json
// @Param        user  body      CreateUserRequest  true  "Dados do usuário"
// @Success      201   {object}  user.User
// @Failure      400   {object}  api.ErrorResponse
// @Router       /user [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleValidationError(c, err)
		return
	}

	if err := h.userUseCase.CreateUser.Execute(c, req.Name, req.Email, req.Password); err != nil {
		utils.HandleValidationError(c, err)
		return
	}

	user, err := h.userUseCase.GetByEmailUser.Execute(c, req.Email)
	if err != nil {
		utils.HandleValidationError(c, err)
		return
	}

	user.Password = ""
	user.Token = ""

	c.JSON(http.StatusCreated, user)

}

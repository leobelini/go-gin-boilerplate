package user

import (
	"leobelini/cashly/internal/types/database"
	"leobelini/cashly/internal/utils"
	"net/http"

	_userController "leobelini/cashly/internal/controller/user"

	"github.com/gin-gonic/gin"
)

// UserRequest representa os dados de entrada
type UserRequest struct {
	Name     string `json:"name" binding:"required" validate:"min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// UserResponse representa os dados de saída
type UserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// createUser godoc
// @Summary      Cria um usuário
// @Description  Cria um usuário com nome, email e senha
// @Tags         usuários
// @Accept       json
// @Produce      json
// @Param        user  body      UserRequest  true  "Dados do usuário"
// @Success      201   {object}  UserResponse
// @Failure      400   {object}  api.ErrorResponse
// @Router       /user [post]
func createUser(c *gin.Context) {
	var req UserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleValidationError(c, err)
		return
	}

	userController := &_userController.UserController{}
	createdUser, err := userController.CreateUser(&database.User{Name: req.Name, Email: req.Email, Password: req.Password})
	if err != nil {
		utils.HandleValidationError(c, err)
		return
	}

	user := UserResponse{ID: createdUser.ID, Name: createdUser.Name, Email: createdUser.Email}

	c.JSON(http.StatusCreated, user)
}

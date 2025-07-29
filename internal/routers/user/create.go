package user

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Validator instance (pode ficar global)
var validate = validator.New()

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

// ErrorResponse padrão para erros
type ErrorResponse struct {
	Message string `json:"message"`
}

// createUser godoc
// @Summary      Cria um usuário
// @Description  Cria um usuário com nome, email e senha
// @Tags         usuários
// @Accept       json
// @Produce      json
// @Param        user  body      UserRequest  true  "Dados do usuário"
// @Success      201   {object}  UserResponse
// @Failure      400   {object}  ErrorResponse
// @Router       /user [post]
func createUser(c *gin.Context) {
	var req UserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]string, len(ve))
			for i, fe := range ve {
				out[i] = fmt.Sprintf("Campo '%s' falhou na validação '%s'", fe.Field(), fe.Tag())
			}
			c.JSON(http.StatusBadRequest, gin.H{"errors": out})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Sucesso
	user := UserResponse{
		ID:    "123",
		Name:  req.Name,
		Email: req.Email,
	}

	c.JSON(http.StatusCreated, user)
}

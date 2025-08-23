package user

import (
	"context"
	"leobelini/cashly/internal/entity"
	"leobelini/cashly/internal/model"
	"leobelini/cashly/internal/utils"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (c *UserController) CreateUser(name, email, password string, ctx context.Context) error {

	return c.app.Model.Transaction(func(tx *model.Model) error {
		// Verifica se já existe
		userExists, err := tx.User.GetByEmailUser(email, ctx)
		if err != nil {
			return err
		}
		if userExists != (entity.User{}) {
			return utils.CreateAppError("USER_ALREADY_EXISTS", false)
		}

		// Criptografa a senha
		bytesPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		token := uuid.New().String()
		user := entity.User{
			ID:              uuid.New().String(),
			Name:            name,
			Email:           email,
			CreatedAt:       time.Now(),
			Password:        string(bytesPassword),
			AccountVerified: false,
			Token:           &token,
		}

		// Cria usuário no banco
		if err := tx.User.CreateUser(user, ctx); err != nil {
			return err
		}

		// Envia para a fila – se falhar, retorna erro para cancelar a transação
		if err := c.app.Job.RecoveryPassword(ctx, user.Email, user.Name, *user.Token); err != nil {
			return err // rollback será disparado
		}

		return nil
	})

}

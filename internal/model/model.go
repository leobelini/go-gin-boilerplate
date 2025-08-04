package model

import (
	"leobelini/cashly/internal/model/user"

	"gorm.io/gorm"
)

type Model struct {
	db   *gorm.DB
	User *user.UserModel
}

func LoadModels(dbGorm *gorm.DB) *Model {
	return &Model{User: user.NewUserModel(dbGorm), db: dbGorm}
}

func (m *Model) Transaction(fc func(tx *Model) error) error {
	return m.db.Transaction(func(tx *gorm.DB) error {
		// Criar nova instância do Model usando a transação
		txModel := &Model{
			db:   tx,
			User: user.NewUserModel(tx),
		}
		return fc(txModel)
	})
}

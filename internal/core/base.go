package core

import (
	"leobelini/cashly/internal/core/app"
	"leobelini/cashly/internal/core/dto"
)

type BaseApp struct {
	Database *app.Database
	Env      *dto.DtoEnvApp
}

func NewBaseApp() *BaseApp {
	env, err := app.LoadEnv()
	if err != nil {
		panic(err)
	}

	dataBase := app.NewDatabase(env)

	return &BaseApp{Database: dataBase, Env: env}
}

package core

import (
	"leobelini/cashly/internal/core/app"
	"leobelini/cashly/internal/core/dto"

	"github.com/hibiken/asynq"
)

type BaseApp struct {
	Database *app.Database
	Env      *dto.DtoEnvApp
	Job      *asynq.Client
}

func NewBaseApp() *BaseApp {
	env, err := app.LoadEnv()
	if err != nil {
		panic(err)
	}

	dataBase := app.NewDatabase(env)
	job := app.NewJob(env)
	job.StartJob()

	return &BaseApp{Database: dataBase, Env: env, Job: job.Client}
}

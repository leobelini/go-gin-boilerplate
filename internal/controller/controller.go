package controller

import (
	"leobelini/cashly/internal/controller/auth"
	"leobelini/cashly/internal/controller/user"
	"leobelini/cashly/internal/core/dto"
	internalDto "leobelini/cashly/internal/dto"
	"leobelini/cashly/internal/model"
	"leobelini/cashly/internal/queue/job"
)

type Controller struct {
	User *user.UserController
	Auth *auth.AuthController
	App  *internalDto.DtoApp
}

func NewController(model *model.Model, job *job.Job, env *dto.DtoEnvApp) *Controller {

	app := &internalDto.DtoApp{Env: env, Job: job, Model: model}

	userC := user.NewUserController(app)
	authC := auth.NewAuthController(app)

	return &Controller{
		User: userC,
		Auth: authC,
		App:  app,
	}
}

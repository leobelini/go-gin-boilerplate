package controller

import (
	"leobelini/cashly/internal/controller/user"
	"leobelini/cashly/internal/core/dto"
	internalDto "leobelini/cashly/internal/dto"
	"leobelini/cashly/internal/model"
	"leobelini/cashly/internal/queue/job"
)

type Controller struct {
	User *user.UserController
	App  *internalDto.DtoApp
}

func NewController(model *model.Model, job *job.Job, env *dto.DtoEnvApp) *Controller {

	app := &internalDto.DtoApp{Env: env, Job: job}

	return &Controller{
		User: user.NewUserController(model, job, env),
		App:  app,
	}
}

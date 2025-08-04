package controller

import (
	"leobelini/cashly/internal/controller/user"
	"leobelini/cashly/internal/core/dto"
	"leobelini/cashly/internal/model"
	"leobelini/cashly/internal/queue/job"
)

type Controller struct {
	User *user.UserController
}

func NewController(model *model.Model, job *job.Job, env *dto.DtoEnvApp) *Controller {
	return &Controller{
		User: user.NewUserController(model, job, env),
	}
}

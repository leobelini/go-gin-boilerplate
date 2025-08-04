package user

import (
	"leobelini/cashly/internal/core/dto"
	"leobelini/cashly/internal/model"
	"leobelini/cashly/internal/queue/job"
)

type UserController struct {
	model *model.Model
	job   *job.Job
	env   *dto.DtoEnvApp
}

func NewUserController(model *model.Model, job *job.Job, env *dto.DtoEnvApp) *UserController {
	return &UserController{model: model, job: job, env: env}
}

package auth

import (
	"leobelini/cashly/internal/core/dto"
	"leobelini/cashly/internal/model"
	"leobelini/cashly/internal/queue/job"
)

type AuthController struct {
	model *model.Model
	job   *job.Job
	env   *dto.DtoEnvApp
}

func NewAuthController(model *model.Model, job *job.Job, env *dto.DtoEnvApp) *AuthController {
	return &AuthController{model: model, job: job, env: env}
}

package dto

import (
	"leobelini/cashly/internal/core/dto"
	"leobelini/cashly/internal/model"
	"leobelini/cashly/internal/queue/job"
)

type DtoApp struct {
	Env   *dto.DtoEnvApp
	Job   *job.Job
	Model *model.Model
}

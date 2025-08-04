package app

import (
	"leobelini/cashly/internal/core/dto"
	"strconv"

	"github.com/hibiken/asynq"
)

type Job struct {
	env    *dto.DtoEnvApp
	Client *asynq.Client
}

func NewJob(env *dto.DtoEnvApp) *Job {
	return &Job{env: env}
}

func (j *Job) StartJob() {
	addr := j.env.Redis.Host + ":" + strconv.Itoa(j.env.Redis.Port)

	j.Client = asynq.NewClient(asynq.RedisClientOpt{Addr: addr})
}

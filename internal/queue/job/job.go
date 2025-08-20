package job

import (
	"github.com/hibiken/asynq"
)

type Job struct {
	SendConfirmationEmailRegister *SendConfirmationEmailRegisterDispatcher
}

func NewJob(redisClient *asynq.Client) *Job {
	return &Job{SendConfirmationEmailRegister: SendConfirmationEmailRegisterQueue(redisClient)}
}

package job

import (
	confirmemailregister "leobelini/cashly/internal/job/confirm_email_register"

	"github.com/hibiken/asynq"
)

type Job struct {
	SendConfirmationEmailRegister *confirmemailregister.SendConfirmationEmailRegisterDispatcher
}

func NewJob(redisClient *asynq.Client) *Job {
	return &Job{SendConfirmationEmailRegister: confirmemailregister.NewTaskDispatcher(redisClient)}
}

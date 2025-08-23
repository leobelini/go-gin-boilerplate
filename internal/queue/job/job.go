package job

import (
	"github.com/hibiken/asynq"
)

type Job struct{ client *asynq.Client }

func NewJob(redisClient *asynq.Client) *Job {
	return &Job{client: redisClient}
}

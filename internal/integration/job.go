package integration

import (
	"leobelini/cashly/config"

	"github.com/hibiken/asynq"
)

func StartJob() *asynq.Client {
	config.LoadRedisEnv()
	env := config.GetRedisEnv()

	addr := env.Host + ":" + env.Port

	return asynq.NewClient(asynq.RedisClientOpt{Addr: addr})
}

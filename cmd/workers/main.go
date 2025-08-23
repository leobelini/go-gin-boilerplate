package main

import (
	"fmt"
	"leobelini/cashly/internal/controller"
	"leobelini/cashly/internal/core"
	"leobelini/cashly/internal/model"
	"leobelini/cashly/internal/queue/job"
	"leobelini/cashly/internal/queue/worker"

	"github.com/hibiken/asynq"
)

func main() {
	baseApp := core.NewBaseApp()
	env := baseApp.Env
	dataBase := baseApp.Database
	job := job.NewJob(baseApp.Job)

	addr := fmt.Sprintf("%s:%d", env.Redis.Host, env.Redis.Port)

	srv := asynq.NewServer(asynq.RedisClientOpt{Addr: addr},
		asynq.Config{Concurrency: 10})

	mux := asynq.NewServeMux()

	models := model.LoadModels(dataBase.Db)
	controllers := controller.NewController(models, job, env)
	worker.RegisterWorkers(mux, controllers)

	if err := srv.Run(mux); err != nil {
		panic(err)
	}
}

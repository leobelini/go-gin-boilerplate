package main

import (
	"fmt"
	"leobelini/cashly/internal/core"
	confirmemailregister "leobelini/cashly/internal/job/confirm_email_register"

	"github.com/hibiken/asynq"
)

func main() {
	baseApp := core.NewBaseApp()
	env := baseApp.Env

	addr := fmt.Sprintf("%s:%d", env.Redis.Host, env.Redis.Port)

	srv := asynq.NewServer(asynq.RedisClientOpt{Addr: addr},
		asynq.Config{Concurrency: 10})

	mux := asynq.NewServeMux()

	mux.HandleFunc(confirmemailregister.TypeSendConfirmationEmail, confirmemailregister.NewWorker)
	fmt.Println("Worker ", confirmemailregister.TypeSendConfirmationEmail, " registered")

	if err := srv.Run(mux); err != nil {
		panic(err)
	}
	fmt.Println("Server started")

}

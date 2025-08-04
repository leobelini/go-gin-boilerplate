package worker

import (
	"fmt"
	"leobelini/cashly/internal/controller"
	"leobelini/cashly/internal/queue/dto"

	"github.com/hibiken/asynq"
)

func RegisterWorkers(mux *asynq.ServeMux, controllers *controller.Controller) {

	mux.HandleFunc(dto.TypeSendConfirmationEmail, SendConfirmationEmailRegisterWorker(controllers).Worker)
	fmt.Println("Worker ", dto.TypeSendConfirmationEmail, " registered")

}

package worker

import (
	"fmt"
	"leobelini/cashly/internal/controller"
	"leobelini/cashly/internal/queue/dto"

	"github.com/hibiken/asynq"
)

type Worker struct {
	controllers *controller.Controller
}

func RegisterWorkers(mux *asynq.ServeMux, controllers *controller.Controller) {

	workers := &Worker{controllers: controllers}

	mux.HandleFunc(dto.TypeSendConfirmationEmail, workers.SendConfirmationEmailRegisterWorker)
	fmt.Println("Worker ", dto.TypeSendConfirmationEmail, " registered")

}

package worker

import (
	"fmt"
	"leobelini/cashly/internal/controller"
	internalDto "leobelini/cashly/internal/dto"
	"leobelini/cashly/internal/queue/dto"

	"github.com/hibiken/asynq"
)

type Worker struct {
	controllers *controller.Controller
	app         *internalDto.DtoApp
}

func RegisterWorkers(mux *asynq.ServeMux, controllers *controller.Controller) {

	workers := &Worker{controllers: controllers, app: controllers.App}

	mux.HandleFunc(dto.TypeSendConfirmationEmail, workers.SendConfirmationEmailRegisterWorker)
	fmt.Println("Worker ", dto.TypeSendConfirmationEmail, " registered")

	mux.HandleFunc(dto.TypeSendRecoveryPasswordEmail, workers.RecoveryPassword)
	fmt.Println("Worker ", dto.TypeSendRecoveryPasswordEmail, " registered")

}

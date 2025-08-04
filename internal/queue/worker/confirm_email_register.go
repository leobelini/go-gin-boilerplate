package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"leobelini/cashly/internal/controller"
	"leobelini/cashly/internal/queue/dto"

	"github.com/hibiken/asynq"
)

type SendConfirmationEmailRegisterDispatcher struct {
	controllers *controller.Controller
}

func SendConfirmationEmailRegisterWorker(controllers *controller.Controller) *SendConfirmationEmailRegisterDispatcher {
	return &SendConfirmationEmailRegisterDispatcher{controllers: controllers}
}

func (w *SendConfirmationEmailRegisterDispatcher) Worker(ctx context.Context, t *asynq.Task) error {
	var payload dto.SendConfirmationEmailRegisterPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return err
	}

	fmt.Printf("📨 Enviando e-mail de boas-vindas para %s <%s>\n", payload.Name, payload.Email)

	return nil
}

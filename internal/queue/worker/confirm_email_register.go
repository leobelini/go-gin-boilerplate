package worker

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"leobelini/cashly/internal/queue/dto"
	"leobelini/cashly/internal/utils"

	"github.com/hibiken/asynq"
)

type PayloadSendConfirmationEmailRegister struct {
	Name     string
	AppName  string
	TokenUrl string
}

//go:embed templates/send_confirmation_email_register.mjml
var templateSendConfirmationEmailRegister string

func (w *Worker) SendConfirmationEmailRegisterWorker(ctx context.Context, t *asynq.Task) error {
	var payload dto.SendConfirmationEmailRegisterPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return err
	}

	fmt.Printf("📨 Enviando e-mail de boas-vindas para %s <%s>\n", payload.Name, payload.Email)

	payloadTemplate := PayloadSendConfirmationEmailRegister{
		Name:     payload.Name,
		AppName:  w.app.Env.App.Name,
		TokenUrl: fmt.Sprintf("%s/sign-up/%s", w.app.Env.App.URL, payload.Token),
	}

	err := utils.SendEmail(templateSendConfirmationEmailRegister, payloadTemplate, w.app.Env, utils.SendEmailParams{To: payload.Email, Ctx: ctx, Subject: "Confirmação de Cadastro"})
	if err != nil {
		return fmt.Errorf("erro ao enviar e-mail: %w", err)
	}

	return nil
}

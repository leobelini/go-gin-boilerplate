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

type PayloadRecoveryPassword struct {
	Name     string
	AppName  string
	TokenUrl string
}

//go:embed templates/send_recovery_password_email.mjml
var templateRecoveryPassword string

func (w *Worker) RecoveryPassword(ctx context.Context, t *asynq.Task) error {
	var payload dto.SendConfirmationEmailRegisterPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return err
	}

	fmt.Printf("📨 Enviando e-mail de recuperação de senha para %s <%s>\n", payload.Name, payload.Email)

	payloadTemplate := PayloadSendConfirmationEmailRegister{
		Name:     payload.Name,
		AppName:  w.app.Env.App.Name,
		TokenUrl: fmt.Sprintf("%s/password-recovery/%s", w.app.Env.App.URL, payload.Token),
	}

	err := utils.SendEmail(templateRecoveryPassword, payloadTemplate, w.app.Env, utils.SendEmailParams{To: payload.Email, Ctx: ctx, Subject: "Recuperação de senha"})
	if err != nil {
		return fmt.Errorf("erro ao enviar e-mail: %w", err)
	}

	fmt.Printf("✅ E-mail de recuperação de senha enviado para %s <%s>\n", payload.Name, payload.Email)

	return nil
}

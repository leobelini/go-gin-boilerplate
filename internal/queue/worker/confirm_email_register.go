package worker

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"html/template"
	"leobelini/cashly/internal/queue/dto"
	"net/smtp"

	"github.com/Boostport/mjml-go"
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

	tpl, err := template.New("email").Parse(templateSendConfirmationEmailRegister)
	if err != nil {
		return fmt.Errorf("erro ao parsear template MJML: %w", err)
	}

	payloadTemplate := PayloadSendConfirmationEmailRegister{
		Name:     payload.Name,
		AppName:  w.controllers.Env.App.Name,
		TokenUrl: fmt.Sprintf("%s/confirm?token=%s", w.controllers.Env.App.URL, payload.Token),
	}

	var mjmlBuffer bytes.Buffer
	if err := tpl.Execute(&mjmlBuffer, payloadTemplate); err != nil {
		return fmt.Errorf("erro ao executar template MJML: %w", err)
	}

	// 2️⃣ Converte MJML para HTML
	html, err := mjml.ToHTML(ctx, mjmlBuffer.String(), mjml.WithMinify(true))
	if err != nil {
		return fmt.Errorf("erro ao converter MJML para HTML: %w", err)
	}

	// 3️⃣ Envia o e-mail usando Mailpit (localhost:1025)
	from := "no-reply@cashly.com"
	to := payload.Email
	msg := []byte(fmt.Sprintf(
		"From: %s\r\n"+
			"To: %s\r\n"+
			"Subject: Confirmação de Cadastro\r\n"+
			"MIME-Version: 1.0\r\n"+
			"Content-Type: text/html; charset=UTF-8\r\n"+
			"\r\n%s",
		from, to, html,
	))

	addr := fmt.Sprintf("%s:%d", w.controllers.Env.Smtp.Host, w.controllers.Env.Smtp.Port)
	if err := smtp.SendMail(addr, nil, from, []string{to}, msg); err != nil {
		return fmt.Errorf("erro ao enviar e-mail: %w", err)
	}

	return nil
}

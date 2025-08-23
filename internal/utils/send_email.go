package utils

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"leobelini/cashly/internal/core/dto"
	"net/smtp"

	"github.com/Boostport/mjml-go"
)

type SendEmailParams struct {
	To      string
	Subject string
	Ctx     context.Context
}

func SendEmail(temp string, payload interface{}, env *dto.DtoEnvApp, params SendEmailParams) error {

	// 1️⃣ Renderiza o template MJML )
	tpl, err := template.New("email").Parse(temp)
	if err != nil {
		return err
	}

	var mjmlBuffer bytes.Buffer
	if err := tpl.Execute(&mjmlBuffer, payload); err != nil {
		return err
	}

	// 2️⃣ Converte MJML para HTML
	html, err := mjml.ToHTML(params.Ctx, mjmlBuffer.String(), mjml.WithMinify(true))
	if err != nil {
		return err
	}

	fromHeader := fmt.Sprintf("Cashly <%s>", env.Smtp.From)

	msg := []byte(fmt.Sprintf(
		"From: %s\r\n"+
			"To: %s\r\n"+
			"Subject: %s\r\n"+
			"MIME-Version: 1.0\r\n"+
			"Content-Type: text/html; charset=UTF-8\r\n"+
			"\r\n%s",
		fromHeader, params.To, params.Subject, html,
	))

	addr := fmt.Sprintf("%s:%d", env.Smtp.Host, env.Smtp.Port)
	err = smtp.SendMail(addr, nil, env.Smtp.From, []string{params.To}, msg)
	if err != nil {
		return err
	}

	return nil
}

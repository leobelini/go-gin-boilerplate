package job

import (
	"context"
	"encoding/json"
	"leobelini/cashly/internal/queue/dto"

	"github.com/hibiken/asynq"
)

func (d *Job) RecoveryPassword(ctx context.Context, email, name, token string) error {
	payload, err := json.Marshal(dto.SendRecoveryPasswordPayload{
		Token: token,
		Email: email,
		Name:  name,
	})
	if err != nil {
		return err
	}

	task := asynq.NewTask(dto.TypeSendRecoveryPasswordEmail, payload)
	_, err = d.client.EnqueueContext(ctx, task)
	return err
}

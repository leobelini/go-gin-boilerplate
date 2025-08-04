package job

import (
	"context"
	"encoding/json"
	"leobelini/cashly/internal/queue/dto"

	"github.com/hibiken/asynq"
)

type SendConfirmationEmailRegisterDispatcher struct {
	client *asynq.Client
}

func SendConfirmationEmailRegisterQueue(redisClient *asynq.Client) *SendConfirmationEmailRegisterDispatcher {
	return &SendConfirmationEmailRegisterDispatcher{client: redisClient}
}

func (d *SendConfirmationEmailRegisterDispatcher) AddQueue(ctx context.Context, email, name, token string) error {
	payload, err := json.Marshal(dto.SendConfirmationEmailRegisterPayload{
		Token: token,
		Email: email,
		Name:  name,
	})
	if err != nil {
		return err
	}

	task := asynq.NewTask(dto.TypeSendConfirmationEmail, payload)
	_, err = d.client.EnqueueContext(ctx, task)
	return err
}

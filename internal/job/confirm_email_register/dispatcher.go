package confirmemailregister

import (
	"context"
	"encoding/json"

	"github.com/hibiken/asynq"
)

type SendConfirmationEmailRegisterDispatcher struct {
	client *asynq.Client
}

func NewTaskDispatcher(redisClient *asynq.Client) *SendConfirmationEmailRegisterDispatcher {
	return &SendConfirmationEmailRegisterDispatcher{client: redisClient}
}

func (d *SendConfirmationEmailRegisterDispatcher) SendConfirmationEmailRegister(ctx context.Context, email, name, token string) error {
	payload, err := json.Marshal(Payload{
		Token: token,
		Email: email,
		Name:  name,
	})
	if err != nil {
		return err
	}

	task := asynq.NewTask(TypeSendConfirmationEmail, payload)
	_, err = d.client.EnqueueContext(ctx, task)
	return err
}

package confirmemailregister

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
)

func NewWorker(ctx context.Context, t *asynq.Task) error {
	var payload Payload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return err
	}

	fmt.Printf("📨 Enviando e-mail de boas-vindas para %s <%s>\n", payload.Name, payload.Email)

	return nil
}

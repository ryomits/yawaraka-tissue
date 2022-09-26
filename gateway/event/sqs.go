package event

import "context"

type SQS interface {
	Receive(ctx context.Context, f func(key, data string) error) (int, error)
	SendMessage(ctx context.Context, key, msg string) error
	BatchSendMessage(ctx context.Context, key string, messages []string) error
}

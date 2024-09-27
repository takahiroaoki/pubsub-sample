package handler

import "context"

type Handler[Message any] interface {
	HandleMessage(ctx context.Context, msg *Message) error
	validate(ctx context.Context, msg *Message) error
}

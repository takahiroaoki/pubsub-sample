package infra

import (
	"context"
	"subscriber/handler"

	"cloud.google.com/go/pubsub"
)

type Subscriber[Message any] interface {
	Receive(ctx context.Context, msgHandler handler.Handler[Message]) error
	handleError(err error, msg *pubsub.Message)
}

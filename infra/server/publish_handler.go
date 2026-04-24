package server

import (
	"context"
	"net/http"
	"pubsub-sample/model"
)

type publisher interface {
	Publish(ctx context.Context, msg model.Something) (string, error)
}

type publishHandler struct {
	p publisher
}

func (h *publishHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	h.p.Publish(ctx, model.Something{})
}

func NewPublishHandler(p publisher) http.Handler {
	return &publishHandler{
		p: p,
	}
}

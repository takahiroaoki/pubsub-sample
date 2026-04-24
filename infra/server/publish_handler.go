package server

import (
	"context"
	"net/http"
	"pubsub-sample/model"
	"pubsub-sample/util"
)

type publisher interface {
	Publish(ctx context.Context, st model.Something) (string, error)
}

type publishHandler struct {
	p publisher
}

func (h *publishHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if _, err := h.p.Publish(ctx, model.Something{}); err != nil {
		util.ErrorLog(err.Error())
	}
	util.InfoLog("publish success")
}

func NewPublishHandler(p publisher) http.Handler {
	return &publishHandler{
		p: p,
	}
}

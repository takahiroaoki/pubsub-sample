package publisher

import (
	"context"
	"pubsub-sample/model"
	"pubsub-sample/util"

	"cloud.google.com/go/pubsub"
)

type publisher struct {
	topic *pubsub.Topic
}

func (p *publisher) Publish(ctx context.Context, msg model.Something) (string, error) {
	util.InfoLog("published!")
	return "", nil
}

func NewPublisher(projectID, topicID string) (p *publisher, closeFunc func() error, err error) {
	return nil, func() error {
		return nil
	}, nil
}

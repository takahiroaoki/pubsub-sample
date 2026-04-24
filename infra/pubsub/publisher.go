package publisher

import (
	"context"
	"errors"
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
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return nil, nil, err
	}

	topic := client.Topic(topicID)
	exists, err := topic.Exists(ctx)
	if err != nil {
		return nil, nil, err
	}
	if !exists {
		return nil, nil, errors.New("topic does not exist")
	}
	topic.EnableMessageOrdering = true

	return &publisher{
		topic: topic,
	}, client.Close, nil
}

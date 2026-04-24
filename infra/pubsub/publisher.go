package pubsubclient

import (
	"context"
	"encoding/json"
	"errors"
	"pubsub-sample/model"
	"time"

	"cloud.google.com/go/pubsub"
)

type publisher struct {
	topic *pubsub.Topic
}

func (p *publisher) Publish(ctx context.Context, st model.Something) (string, error) {
	data, err := json.Marshal(st)
	if err != nil {
		return "", err
	}

	orderingKey := time.Now().String()
	srvID, err := p.topic.Publish(ctx, &pubsub.Message{
		OrderingKey: orderingKey,
		Data:        data,
	}).Get(ctx)
	if err != nil {
		p.topic.ResumePublish(orderingKey)
		return srvID, err
	}
	return srvID, nil
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

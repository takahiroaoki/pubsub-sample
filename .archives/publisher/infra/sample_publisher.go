package infra

import (
	"context"
	"encoding/json"
	"time"

	"cloud.google.com/go/pubsub"
)

type samplePublisherImpl struct {
	topic *pubsub.Topic
}

func (spi *samplePublisherImpl) Publish(msg sampleMessage) (string, error) {
	ctx := context.Background()
	orderingKey := time.Now().String()

	data, err := json.Marshal(struct {
		Word string `json:"word"`
	}{
		Word: msg.word,
	})
	if err != nil {
		return "", err
	}

	srvID, err := spi.topic.Publish(ctx, &pubsub.Message{
		Data:        data,
		OrderingKey: orderingKey,
	}).Get(ctx)
	if err != nil {
		spi.topic.ResumePublish(orderingKey)
		return srvID, err
	}
	return srvID, nil
}

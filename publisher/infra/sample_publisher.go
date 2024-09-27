package infra

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"cloud.google.com/go/pubsub"
)

type SamplePublisher struct {
	topic *pubsub.Topic
}

type SampleMessage struct {
	word string
}

func NewSampleMessage(word string) SampleMessage {
	return SampleMessage{
		word: word,
	}
}

func (sp *SamplePublisher) Publish(msg SampleMessage) (string, error) {
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

	srvID, err := sp.topic.Publish(ctx, &pubsub.Message{
		Data:        data,
		OrderingKey: orderingKey,
	}).Get(ctx)
	if err != nil {
		sp.topic.ResumePublish(orderingKey)
		return srvID, err
	}
	return srvID, nil
}

type SamplePublisherConfig struct {
	projectID string
	topicID   string
}

func NewSamplePublisherConfig(projectID string, topicID string) SamplePublisherConfig {
	return SamplePublisherConfig{
		projectID: projectID,
		topicID:   topicID,
	}
}

func NewSamplePublisher(config SamplePublisherConfig) (Publisher[SampleMessage], func() error, error) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, config.projectID)
	if err != nil {
		return nil, nil, err
	}

	topic := client.Topic(config.topicID)
	exists, err := topic.Exists(ctx)
	if err != nil {
		return nil, nil, err
	}
	if !exists {
		return nil, nil, errors.New("topic does not exist")
	}
	topic.EnableMessageOrdering = true

	return &SamplePublisher{
		topic: topic,
	}, client.Close, nil
}

package infra

import (
	"context"
	"errors"

	"cloud.google.com/go/pubsub"
)

type Publisher[Message any] interface {
	Publish(msg Message) (serverID string, err error)
}

type publisherConfig struct {
	projectID string
	topicID   string
}

func NewPublisherConfig(projectID string, topicID string) publisherConfig {
	return publisherConfig{
		projectID: projectID,
		topicID:   topicID,
	}
}

/*
SamplePublisher
*/
type sampleMessage struct {
	word string
}

func NewSampleMessage(word string) sampleMessage {
	return sampleMessage{
		word: word,
	}
}

func NewSamplePublisher(config publisherConfig) (Publisher[sampleMessage], func() error, error) {
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

	return &samplePublisherImpl{
		topic: topic,
	}, client.Close, nil
}

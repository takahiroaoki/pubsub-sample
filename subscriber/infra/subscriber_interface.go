package infra

import (
	"context"
	"fmt"
	"log"
	"subscriber/handler"

	"cloud.google.com/go/pubsub"
)

type Subscriber[Handler any] interface {
	Receive(ctx context.Context, msgHandler Handler) error
	HasDeadLetterSubscription() bool
	ReceiveDeadLetter(ctx context.Context, msgHandler Handler) error
}

type subscriberConfig struct {
	projectID        string
	subscriptionID   string
	dlSubscriptionID string
}

func NewSubscriberConfig(projectID string, subscriptionID string, dlSubscriptionID string) subscriberConfig {
	return subscriberConfig{
		projectID:        projectID,
		subscriptionID:   subscriptionID,
		dlSubscriptionID: dlSubscriptionID,
	}
}

func startSubscription(ctx context.Context, client *pubsub.Client, subscriptionID string) (*pubsub.Subscription, error) {
	subscription := client.Subscription(subscriptionID)
	exists, err := subscription.Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("subscription (%v) does not exist", subscriptionID)
	}
	return subscription, nil
}

/*
SampleSubscriber
*/
func NewSampleSubscriber(ctx context.Context, config subscriberConfig) (Subscriber[handler.SampleHandler], func(), error) {
	client, err := pubsub.NewClient(ctx, config.projectID)
	if err != nil {
		return nil, nil, err
	}

	closeFunc := func() {
		if err := client.Close(); err != nil {
			log.Println("falied to close subscription client")
		}
	}

	sampleSubscriber := &sampleSubscriberImpl{}

	subscription, err := startSubscription(ctx, client, config.subscriptionID)
	if err != nil {
		return nil, nil, err
	}
	sampleSubscriber.subscription = subscription

	if len(config.dlSubscriptionID) > 0 {
		dlSubscription, err := startSubscription(ctx, client, config.dlSubscriptionID)
		if err != nil {
			return nil, nil, err
		}
		sampleSubscriber.dlSubscription = dlSubscription
	}

	return sampleSubscriber, closeFunc, nil
}

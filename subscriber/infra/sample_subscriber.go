package infra

import (
	"context"
	"subscriber/handler"

	"cloud.google.com/go/pubsub"
)

type SampleSubscriber struct {
	subscription *pubsub.Subscription
}

func (ss *SampleSubscriber) Receive(ctx context.Context, msgHandler handler.Handler[handler.SampleMessage]) error {
	if err := ss.subscription.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		if msg.DeliveryAttempt != nil {
			// TODO
		}

		if err := msgHandler.HandleMessage(ctx, handler.NewSampleMessage()); err != nil {
			ss.handleError(err, msg)
		}
	}); err != nil {
		// TODO
	}
	return nil
}

func (ss *SampleSubscriber) handleError(err error, msg *pubsub.Message) {
	// TODO
}

type SampleSubscriberConfig struct {
	projectID      string
	subscriptionID string
}

func NewSampleSubscriberConfig(projectID string, subscriptionID string) SampleSubscriberConfig {
	return SampleSubscriberConfig{
		projectID:      projectID,
		subscriptionID: subscriptionID,
	}
}

func NewSampleSubscriber(ctx context.Context, config SampleSubscriberConfig) (Subscriber[handler.SampleMessage], func(), error) {
	client, err := pubsub.NewClient(ctx, config.projectID)
	if err != nil {
		return nil, nil, err
	}

	closeFunc := func() {
		if err := client.Close(); err != nil {
			// TODO:
		}
	}

	return &SampleSubscriber{
		subscription: client.Subscription(config.subscriptionID),
	}, closeFunc, nil
}

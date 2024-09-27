package infra

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"subscriber/handler"

	"cloud.google.com/go/pubsub"
)

type SampleSubscriber struct {
	subscription *pubsub.Subscription
}

func (ss *SampleSubscriber) Receive(ctx context.Context, msgHandler handler.Handler[handler.SampleMessage]) error {
	if err := ss.subscription.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		if msg.DeliveryAttempt != nil {
			log.Printf("DeliveryAttempt: %v", *msg.DeliveryAttempt)
		}

		decoded := struct {
			Word string `json:"word"`
		}{}
		if err := json.Unmarshal(msg.Data, &decoded); err != nil {
			log.Println("Error on unmarshal message")
			msg.Ack()
		}
		if err := msgHandler.HandleMessage(ctx, handler.NewSampleMessage(decoded.Word)); err != nil {
			log.Println(err)
			ss.handleError(err, msg)
		}
	}); err != nil {
		return err
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
			log.Println("falied to close subscription client")
		}
	}

	subscription := client.Subscription(config.subscriptionID)
	exists, err := subscription.Exists(ctx)
	if err != nil {
		return nil, nil, err
	}
	if !exists {
		return nil, nil, errors.New("subscription does not exist")
	}

	return &SampleSubscriber{
		subscription: subscription,
	}, closeFunc, nil
}

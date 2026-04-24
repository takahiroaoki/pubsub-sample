package pubsubclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"pubsub-sample/model"
	"pubsub-sample/util"

	"cloud.google.com/go/pubsub"
)

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

type subscriber struct {
	subscription   *pubsub.Subscription
	dlSubscription *pubsub.Subscription
	msgHandler     msgHandler
	dlMsgHandler   msgHandler
}

type msgHandler interface {
	Handle(ctx context.Context, st model.Something) error
}

func (s *subscriber) Receive(ctx context.Context) error {
	if s == nil {
		return errors.New("*subscriber is nil")
	}

	if err := s.subscription.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		if msg.DeliveryAttempt != nil {
			util.InfoLog(fmt.Sprintf("delivery-attempt: %v", *msg.DeliveryAttempt))
		}

		decoded := model.Something{}
		if err := json.Unmarshal(msg.Data, &decoded); err != nil {
			util.ErrorLog(fmt.Sprintf("unmarshal message: %v", err))
			// it is useless to retry the msssage that cannot be unmarshalled
			msg.Ack()
		}
		if err := s.msgHandler.Handle(ctx, decoded); err != nil {
			util.WarnLog(fmt.Sprintf("handle message: %v", err))
			// make pubsub retry
			msg.Nack()
		}
		msg.Ack()
	}); err != nil {
		return err
	}
	return nil
}

func (s *subscriber) ReceiveDeadLetter(ctx context.Context) error {
	if s == nil {
		return errors.New("*subscriber is nil")
	}

	if err := s.dlSubscription.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		// it is useless to retry the dead letter message
		msg.Ack()

		decoded := model.Something{}
		if err := json.Unmarshal(msg.Data, &decoded); err != nil {
			util.ErrorLog((fmt.Sprintf("unmarshal dead letter message: %v", err)))
			return
		}
		if err := s.dlMsgHandler.Handle(ctx, decoded); err != nil {
			util.ErrorLog(fmt.Sprintf("handle dead letter message: %v", err))
		}
	}); err != nil {
		return err
	}
	return nil
}

func NewSubscriber(projectID, subscriptionID, dlSubscriptionID string, msgHandler, dlMsgHandler msgHandler) (s *subscriber, closeFunc func() error, err error) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return nil, nil, err
	}

	closeFunc = func() error {
		if err := client.Close(); err != nil {
			util.InfoLog(fmt.Sprintf("falied to close subscription client: %v", err))
			return err
		}
		return nil
	}

	subscription, err := startSubscription(ctx, client, subscriptionID)
	if err != nil {
		return nil, nil, err
	}
	dlSubscription, err := startSubscription(ctx, client, dlSubscriptionID)
	if err != nil {
		return nil, nil, err
	}
	subscriber := &subscriber{
		subscription:   subscription,
		dlSubscription: dlSubscription,
		msgHandler:     msgHandler,
		dlMsgHandler:   dlMsgHandler,
	}

	return subscriber, closeFunc, nil
}

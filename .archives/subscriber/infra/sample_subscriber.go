package infra

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"subscriber/handler"

	"cloud.google.com/go/pubsub"
)

type sampleSubscriberImpl struct {
	subscription   *pubsub.Subscription
	dlSubscription *pubsub.Subscription
}

func (ss *sampleSubscriberImpl) Receive(ctx context.Context, msgHandler handler.Handler[handler.SampleMessage]) error {
	if ss == nil {
		return errors.New("*sampleSubscriberImpl is nil")
	}

	if err := ss.subscription.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		if msg.DeliveryAttempt != nil {
			log.Printf("[INFO] delivery-attempt: %v", *msg.DeliveryAttempt)
		}

		decoded := struct {
			Word string `json:"word"`
		}{}
		if err := json.Unmarshal(msg.Data, &decoded); err != nil {
			log.Printf("[ERROR] unmarshal message: %v", err)
			// it is useless to retry the msssage that cannot be unmarshalled
			msg.Ack()
		}
		if err := msgHandler.HandleMessage(ctx, handler.NewSampleMessage(decoded.Word)); err != nil {
			log.Printf("[WARN] handle message: %v", err)
			// make pubsub retry
			msg.Nack()
		}
		msg.Ack()
	}); err != nil {
		return err
	}
	return nil
}

func (ss *sampleSubscriberImpl) HasDeadLetterSubscription() bool {
	return true
}

func (ss *sampleSubscriberImpl) ReceiveDeadLetter(ctx context.Context, msgHandler handler.Handler[handler.SampleMessage]) error {
	if ss == nil {
		return errors.New("*sampleSubscriberImpl is nil")
	}

	if err := ss.dlSubscription.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		// it is useless to retry the dead letter message
		msg.Ack()

		decoded := struct {
			Word string `json:"word"`
		}{}
		if err := json.Unmarshal(msg.Data, &decoded); err != nil {
			log.Printf("[ERROR] unmarshal dead letter message: %v", err)
			return
		}
		if err := msgHandler.HandleDeadLetterMessage(ctx, handler.NewSampleMessage(decoded.Word)); err != nil {
			log.Printf("[ERROR] handle dead letter message: %v", err)
		}
	}); err != nil {
		return err
	}
	return nil
}

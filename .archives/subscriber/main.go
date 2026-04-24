package main

import (
	"context"
	"log"
	"subscriber/config"
	"subscriber/handler"
	"subscriber/infra"
	"sync"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	pubsubConfig := config.NewPubSubConfig()

	sampleHandler := handler.NewSampleHandler()
	sampleSubscriber, stopFunc, err := infra.NewSampleSubscriber(ctx, infra.NewSubscriberConfig(pubsubConfig.ProjectID(), pubsubConfig.SubscriptionID(), pubsubConfig.DeadLetterSubscriptionID()))
	if err != nil {
		log.Fatal(err)
	}
	defer stopFunc()

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := sampleSubscriber.Receive(ctx, sampleHandler); err != nil {
			log.Printf("[ERROR] pull message: %v", err)
			cancel()
		}
	}()

	if sampleSubscriber.HasDeadLetterSubscription() {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := sampleSubscriber.ReceiveDeadLetter(ctx, sampleHandler); err != nil {
				log.Printf("[ERROR] pull dead letter message: %v", err)
				cancel()
			}
		}()
	}

	wg.Wait()
}

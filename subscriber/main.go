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
	sampleSubscriber, stopFunc, err := infra.NewSampleSubscriber(ctx, infra.NewSampleSubscriberConfig(pubsubConfig.ProjectID(), pubsubConfig.SubscriptionID()))
	if err != nil {
		log.Fatal(err)
	}
	defer stopFunc()

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := sampleSubscriber.Receive(ctx, sampleHandler); err != nil {
			cancel()
		}
	}()

	wg.Wait()
}

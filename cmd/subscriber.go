package cmd

import (
	"context"
	"fmt"
	"pubsub-sample/config"
	pubsubclient "pubsub-sample/infra/pubsub"
	"pubsub-sample/usecase"
	"pubsub-sample/util"
	"sync"

	"github.com/spf13/cobra"
)

func newSubscriberCmd() *cobra.Command {
	subscriberCmd := &cobra.Command{
		Use: "subscriber",
		RunE: func(cmd *cobra.Command, args []string) error {
			subscriberConfig := config.NewSubscriberConfig()
			stUsecase := usecase.NewSomethingUsecase()
			stRecUsecase := usecase.NewSomethingRecoveryUsecase()
			subscriber, closeFunc, err := pubsubclient.NewSubscriber(
				subscriberConfig.ProjectID(),
				subscriberConfig.SubscriptionID(),
				subscriberConfig.DlSubscriptionID(),
				stUsecase,
				stRecUsecase,
			)
			defer closeFunc()
			if err != nil {
				util.FatalLog(fmt.Sprintf("NewSubscriber: %v", err))
			}

			wg := &sync.WaitGroup{}
			wg.Add(1)
			ctx, cancel := context.WithCancel(context.Background())
			go func() {
				defer wg.Done()
				if err := subscriber.Receive(ctx); err != nil {
					util.ErrorLog(fmt.Sprintf("pull message: %v", err))
					cancel()
				}
			}()

			wg.Add(1)
			go func() {
				defer wg.Done()
				if err := subscriber.ReceiveDeadLetter(ctx); err != nil {
					util.ErrorLog(fmt.Sprintf("pull dead letter message: %v", err))
					cancel()
				}
			}()

			wg.Wait()
			return nil
		},
	}
	return subscriberCmd
}

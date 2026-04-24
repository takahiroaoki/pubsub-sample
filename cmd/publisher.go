package cmd

import (
	"fmt"
	"pubsub-sample/config"
	publisher "pubsub-sample/infra/pubsub"
	"pubsub-sample/util"

	"github.com/spf13/cobra"
)

func newPublisherCmd() *cobra.Command {
	publisherCmd := &cobra.Command{
		Use: "publisher",
		RunE: func(cmd *cobra.Command, args []string) error {
			publisherConfig := config.NewPublisherConfig()
			publisher, closeFunc, err := publisher.NewPublisher(
				publisherConfig.ProjectID(),
				publisherConfig.TopicID(),
			)
			defer closeFunc()
			if err != nil {
				util.FatalLog(fmt.Sprintf("[NewPublisher] %v", err))
			}
			util.InfoLog(
				fmt.Sprintf("publisher called: %v", publisher),
			)
			return nil
		},
	}
	return publisherCmd
}

package cmd

import (
	"fmt"
	"pubsub-sample/config"
	"pubsub-sample/util"

	"github.com/spf13/cobra"
)

func newPublisherCmd() *cobra.Command {
	publisherCmd := &cobra.Command{
		Use: "publisher",
		RunE: func(cmd *cobra.Command, args []string) error {
			publisherConfig := config.NewPublisherConfig()
			util.InfoLog(
				fmt.Sprintf("publisher called: %v", publisherConfig),
			)
			return nil
		},
	}
	return publisherCmd
}

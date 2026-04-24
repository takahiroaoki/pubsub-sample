package cmd

import (
	"pubsub-sample/util"

	"github.com/spf13/cobra"
)

func newPublisherCmd() *cobra.Command {
	publisherCmd := &cobra.Command{
		Use: "publisher",
		RunE: func(cmd *cobra.Command, args []string) error {
			util.InfoLog("publisher Called!")
			return nil
		},
	}
	return publisherCmd
}

package cmd

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"pubsub-sample/config"
	pubsubclient "pubsub-sample/infra/pubsub"
	"pubsub-sample/infra/server"
	"pubsub-sample/util"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

func newPublisherCmd() *cobra.Command {
	publisherCmd := &cobra.Command{
		Use: "publisher",
		RunE: func(cmd *cobra.Command, args []string) error {
			publisherConfig := config.NewPublisherConfig()
			publisher, closeFunc, err := pubsubclient.NewPublisher(
				publisherConfig.ProjectID(),
				publisherConfig.TopicID(),
			)
			defer closeFunc()
			if err != nil {
				util.FatalLog(fmt.Sprintf("NewPublisher: %v", err))
			}

			srv := &http.Server{
				Addr: ":8080",
			}
			http.Handle("/publish", server.NewPublishHandler(publisher))

			idleConnsClosed := make(chan struct{})
			go func() {
				stopReq := make(chan os.Signal, 1)
				signal.Notify(stopReq, syscall.SIGINT, syscall.SIGTERM)
				<-stopReq

				shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
				defer cancel()

				util.InfoLog("HTTP server is shutting down...")
				if err := srv.Shutdown(shutdownCtx); err != nil {
					util.ErrorLog(fmt.Sprintf("HTTP server Shutdown: %v", err))
				}
				close(idleConnsClosed)
			}()

			util.InfoLog(fmt.Sprintf("Starting HTTP server on %s", srv.Addr))
			if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
				util.ErrorLog(fmt.Sprintf("HTTP server ListenAndServe: %v", err))
			}

			<-idleConnsClosed
			util.InfoLog("HTTP server stopped")
			return nil
		},
	}
	return publisherCmd
}

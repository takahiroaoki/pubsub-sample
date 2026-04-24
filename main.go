package main

import (
	"fmt"
	"pubsub-sample/cmd"
	"pubsub-sample/util"
)

func main() {
	if err := cmd.NewRootCmd().Execute(); err != nil {
		util.FatalLog(fmt.Sprintf("Failed to execute the command. Error: %v", err))
	}
}

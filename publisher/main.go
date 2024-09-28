package main

import (
	"log"
	"publisher/config"
	"publisher/infra"
)

func main() {
	pubsubConfig := config.NewPubSubConfig()
	samplePublisherConfig := infra.NewSamplePublisherConfig(pubsubConfig.ProjectID(), pubsubConfig.TopicID())
	samplePublisher, closeFunc, err := infra.NewSamplePublisher(samplePublisherConfig)
	if err != nil {
		log.Fatalln(err)
	}
	defer closeFunc()

	srvID, err := samplePublisher.Publish(infra.NewSampleMessage("test"))
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("ServerID: %s was successfully published!", srvID)
}

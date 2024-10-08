package main

import (
	"flag"
	"log"
	"publisher/config"
	"publisher/infra"
)

func main() {
	f := flag.String("msg", "undefined", "payload of pubsub message")
	flag.Parse()

	pubsubConfig := config.NewPubSubConfig()
	publisherConfig := infra.NewPublisherConfig(pubsubConfig.ProjectID(), pubsubConfig.TopicID())
	samplePublisher, closeFunc, err := infra.NewSamplePublisher(publisherConfig)
	if err != nil {
		log.Fatalln(err)
	}
	defer closeFunc()

	srvID, err := samplePublisher.Publish(infra.NewSampleMessage(*f))
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("ServerID: %s was successfully published!", srvID)
}

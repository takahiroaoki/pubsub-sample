package config

import (
	"os"
)

var env *envVars

func init() {
	env = &envVars{
		pubsubProjectID:                os.Getenv("PUBSUB_PROJECT_ID"),
		pubsubTopicID:                  os.Getenv("PUBSUB_TOPIC_ID"),
		pubsubDeadLetterTopicID:        os.Getenv("PUBSUB_DEAD_LETTER_TOPIC_ID"),
		pubsubSubscriptionID:           os.Getenv("PUBSUB_SUBSCRIPTION_ID"),
		pubsubDeadLetterSubscriptionID: os.Getenv("PUBSUB_DEAD_LETTER_SUBSCRIPTION_ID"),
	}
}

type envVars struct {
	/* About PubSub */
	pubsubProjectID                string
	pubsubTopicID                  string
	pubsubDeadLetterTopicID        string
	pubsubSubscriptionID           string
	pubsubDeadLetterSubscriptionID string
}

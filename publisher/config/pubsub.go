package config

type PubSubConfig struct {
	projectID                string
	topicID                  string
	deadLetterTopicID        string
	subscriptionID           string
	deadLetterSubscriptionID string
}

func (psc PubSubConfig) ProjectID() string {
	return psc.projectID
}

func (psc PubSubConfig) TopicID() string {
	return psc.topicID
}

func (psc PubSubConfig) DeadLetterTopicID() string {
	return psc.deadLetterTopicID
}

func (psc PubSubConfig) SubscriptionID() string {
	return psc.subscriptionID
}

func (psc PubSubConfig) DeadLetterSubscriptionID() string {
	return psc.deadLetterSubscriptionID
}

func NewPubSubConfig() PubSubConfig {
	return PubSubConfig{
		projectID:                env.pubsubProjectID,
		topicID:                  env.pubsubTopicID,
		deadLetterTopicID:        env.pubsubDeadLetterTopicID,
		subscriptionID:           env.pubsubSubscriptionID,
		deadLetterSubscriptionID: env.pubsubDeadLetterSubscriptionID,
	}
}

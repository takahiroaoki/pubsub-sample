package config

type SubscriberConfig struct {
	projectID        string
	subscriptionID   string
	dlSubscriptionID string
}

func (c SubscriberConfig) ProjectID() string {
	return c.projectID
}

func (c SubscriberConfig) SubscriptionID() string {
	return c.subscriptionID
}

func (c SubscriberConfig) DlSubscriptionID() string {
	return c.dlSubscriptionID
}

func NewSubscriberConfig() SubscriberConfig {
	return SubscriberConfig{
		projectID:        env.pubsubProjectID,
		subscriptionID:   env.pubsubSubscriptionID,
		dlSubscriptionID: env.pubsubDeadLetterSubscriptionID,
	}
}

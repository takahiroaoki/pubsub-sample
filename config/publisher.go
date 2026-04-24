package config

type PublisherConfig struct {
	projectID string
	topicID   string
}

func (c PublisherConfig) ProjectID() string {
	return c.projectID
}

func (c PublisherConfig) TopicID() string {
	return c.topicID
}

func NewPublisherConfig() PublisherConfig {
	return PublisherConfig{
		projectID: env.pubsubProjectID,
		topicID:   env.pubsubTopicID,
	}
}

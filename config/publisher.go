package config

type PublisherConfig struct {
	projectID string
	topicID   string
}

func (psc PublisherConfig) ProjectID() string {
	return psc.projectID
}

func (psc PublisherConfig) TopicID() string {
	return psc.topicID
}

func NewPublisherConfig() PublisherConfig {
	return PublisherConfig{
		projectID: env.pubsubProjectID,
		topicID:   env.pubsubTopicID,
	}
}

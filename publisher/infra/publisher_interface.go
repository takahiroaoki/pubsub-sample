package infra

type Publisher[Message any] interface {
	Publish(msg Message) (serverID string, err error)
}

package handler

import "context"

type Handler[Message any] interface {
	HandleMessage(ctx context.Context, msg *Message) error
	HandleDeadLetterMessage(ctx context.Context, msg *Message) error
	validate(ctx context.Context, msg *Message) error
}

/*
SampleHandler
*/
type sampleMessage struct {
	word string
}

func NewSampleMessage(word string) *sampleMessage {
	return &sampleMessage{
		word: word,
	}
}

type SampleHandler interface {
	Handler[sampleMessage]
}

func NewSampleHandler() SampleHandler {
	return &sampleHandlerImpl{}
}

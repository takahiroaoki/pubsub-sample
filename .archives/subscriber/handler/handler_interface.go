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
type SampleMessage struct {
	word string
}

func NewSampleMessage(word string) *SampleMessage {
	return &SampleMessage{
		word: word,
	}
}

func NewSampleHandler() Handler[SampleMessage] {
	return &sampleHandlerImpl{}
}

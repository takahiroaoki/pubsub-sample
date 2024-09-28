package handler

import (
	"context"
	"log"
)

type SampleMessage struct {
	word string
}

func NewSampleMessage(word string) *SampleMessage {
	return &SampleMessage{
		word: word,
	}
}

type SampleHandler struct{}

func (sh *SampleHandler) HandleMessage(ctx context.Context, msg *SampleMessage) error {
	log.Printf("> %s", msg.word)
	return nil
}

func (sh *SampleHandler) validate(ctx context.Context, msg *SampleMessage) error {
	return nil
}

func NewSampleHandler() Handler[SampleMessage] {
	return &SampleHandler{}
}

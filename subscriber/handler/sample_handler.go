package handler

import "context"

type SampleMessage struct{}

func NewSampleMessage() *SampleMessage {
	return &SampleMessage{}
}

type SampleHandler struct{}

func (sh *SampleHandler) HandleMessage(ctx context.Context, msg *SampleMessage) error {
	return nil
}

func (sh *SampleHandler) validate(ctx context.Context, msg *SampleMessage) error {
	return nil
}

func NewSampleHandler() Handler[SampleMessage] {
	return &SampleHandler{}
}

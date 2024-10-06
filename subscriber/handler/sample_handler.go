package handler

import (
	"context"
	"errors"
	"log"
)

type sampleHandlerImpl struct{}

func (shi *sampleHandlerImpl) HandleMessage(ctx context.Context, msg *sampleMessage) error {
	if shi == nil {
		return errors.New("*sampleHandlerImpl is nil")
	}
	log.Printf("[INFO] > %s", msg.word)
	return nil
}

func (shi *sampleHandlerImpl) HandleDeadLetterMessage(ctx context.Context, msg *sampleMessage) error {
	if shi == nil {
		return errors.New("*sampleHandlerImpl is nil")
	}
	log.Printf("[INFO] > %s", msg.word)
	// basically, need error log to alert the dead letter message
	return nil
}

func (shi *sampleHandlerImpl) validate(ctx context.Context, msg *sampleMessage) error {
	if shi == nil {
		return errors.New("*sampleHandlerImpl is nil")
	}
	return nil
}

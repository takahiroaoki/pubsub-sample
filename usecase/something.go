package usecase

import (
	"context"
	"fmt"
	"pubsub-sample/model"
	"pubsub-sample/util"
)

type somethingUsecase struct{}

func (u *somethingUsecase) Handle(ctx context.Context, st model.Something) error {
	util.InfoLog(fmt.Sprintf("%v", st))
	return nil
}

func NewSomethingUsecase() *somethingUsecase {
	return &somethingUsecase{}
}

type somethingRecoveryUsecase struct{}

func (u *somethingRecoveryUsecase) Handle(ctx context.Context, st model.Something) error {
	util.InfoLog(fmt.Sprintf("%v", st))
	return nil
}

func NewSomethingRecoveryUsecase() *somethingUsecase {
	return &somethingUsecase{}
}

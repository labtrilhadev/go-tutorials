package usecase

import "context"

type AnyUseCase struct{}

func NewAnyUseCase() AnyUseCase {
	return AnyUseCase{}
}

func (u *AnyUseCase) Perform(ctx context.Context) string {
	return "Hello!"
}

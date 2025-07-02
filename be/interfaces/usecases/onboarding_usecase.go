package usecase_interfaces

import (
	"context"
	"github.com/RandySteven/Idolized/entities/payloads/requests"
)

type OnboardingUsecase interface {
	RegisterUser(ctx context.Context, request *requests.OnboardingRequest)
}

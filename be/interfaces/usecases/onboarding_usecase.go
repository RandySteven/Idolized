package usecase_interfaces

import (
	"context"
	"github.com/RandySteven/Idolized/apperror/apperror"
	"github.com/RandySteven/Idolized/entities/payloads/requests"
	"github.com/RandySteven/Idolized/entities/payloads/responses"
)

type OnboardingUsecase interface {
	RegisterUser(ctx context.Context, request *requests.RegisterRequest) (result *responses.RegisterResponse, customErr *apperror.CustomError)
	LoginUser(ctx context.Context, request *requests.LoginRequest) (result *responses.LoginResponse, customErr *apperror.CustomError)
	GetOnboarded(ctx context.Context) (result *responses.OnboardUserResponse, customErr *apperror.CustomError)
}

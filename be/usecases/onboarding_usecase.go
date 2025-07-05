package usecases

import (
	"context"
	"github.com/RandySteven/Idolized/apperror/apperror"
	"github.com/RandySteven/Idolized/entities/payloads/requests"
	"github.com/RandySteven/Idolized/entities/payloads/responses"
	repository_interfaces "github.com/RandySteven/Idolized/interfaces/repositories"
	usecase_interfaces "github.com/RandySteven/Idolized/interfaces/usecases"
)

type (
	onboardingUsecase struct {
		onboardMethod    map[string]interface{}
		userRepository   repository_interfaces.UserRepository
		talentRepository repository_interfaces.TalentRepository
		groupRepository  repository_interfaces.GroupRepository
	}
)

func (o *onboardingUsecase) RegisterUser(ctx context.Context, request *requests.RegisterRequest) (result *responses.RegisterResponse, customErr *apperror.CustomError) {
	//TODO implement me
	panic("implement me")
}

func (o *onboardingUsecase) LoginUser(ctx context.Context, request *requests.LoginRequest) (result *responses.LoginResponse, customErr *apperror.CustomError) {
	//TODO implement me
	panic("implement me")
}

func (o *onboardingUsecase) GetOnboarded(ctx context.Context) (result *responses.OnboardUserResponse, customErr *apperror.CustomError) {
	//TODO implement me
	panic("implement me")
}

func (o *onboardingUsecase) userRegister(ctx context.Context) (result *responses.OnboardUserResponse, customErr *apperror.CustomError) {
	return
}

func (o *onboardingUsecase) talentRegister(ctx context.Context) (result *responses.OnboardUserResponse, customErr *apperror.CustomError) {
	return
}

func (o *onboardingUsecase) groupRegister(ctx context.Context) (result *responses.OnboardUserResponse, customErr *apperror.CustomError) {
	return
}

var _ usecase_interfaces.OnboardingUsecase = &onboardingUsecase{}

func newOnboardingUsecase() *onboardingUsecase {
	o := &onboardingUsecase{}
	onboardMethod := make(map[string]interface{})

	onboardMethod["user"] = o.userRegister
	onboardMethod["talent"] = o.talentRegister
	onboardMethod["group"] = o.groupRegister

	o.onboardMethod = onboardMethod

	return o

}

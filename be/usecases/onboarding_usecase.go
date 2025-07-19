package usecases

import (
	"context"
	"fmt"
	"github.com/RandySteven/Idolized/apperror/apperror"
	"github.com/RandySteven/Idolized/entities/models"
	"github.com/RandySteven/Idolized/entities/payloads/requests"
	"github.com/RandySteven/Idolized/entities/payloads/responses"
	"github.com/RandySteven/Idolized/enums"
	repository_interfaces "github.com/RandySteven/Idolized/interfaces/repositories"
	usecase_interfaces "github.com/RandySteven/Idolized/interfaces/usecases"
	"github.com/RandySteven/Idolized/utils"
	"github.com/google/uuid"
	"time"
)

const (
	userRequestKey   = `register-request-USER`
	talentRequestKey = `register-request-TALENT`
	groupRequestKey  = `register-request-GROUP`
)

type (
	onboardFunc       func(ctx context.Context) (result *responses.RegisterResponse, customErr *apperror.CustomError)
	onboardingUsecase struct {
		onboardMethod     map[string]onboardFunc
		accountRepository repository_interfaces.AccountRepository
		userRepository    repository_interfaces.UserRepository
		talentRepository  repository_interfaces.TalentRepository
		groupRepository   repository_interfaces.GroupRepository
	}
)

func (o *onboardingUsecase) RegisterUser(ctx context.Context, request *requests.RegisterRequest) (result *responses.RegisterResponse, customErr *apperror.CustomError) {

	key := fmt.Sprintf(`register-request-%s`, request.Type)
	ctx2 := context.WithValue(ctx, key, request)

	_, customErr = o.onboardMethod[request.Type](ctx2)
	if customErr != nil {
		return nil, customErr
	}
	return
}

func (o *onboardingUsecase) LoginUser(ctx context.Context, request *requests.LoginRequest) (result *responses.LoginResponse, customErr *apperror.CustomError) {
	return
}

func (o *onboardingUsecase) GetOnboarded(ctx context.Context) (result *responses.OnboardUserResponse, customErr *apperror.CustomError) {
	return
}

func (o *onboardingUsecase) userRegister(ctx context.Context) (result *responses.RegisterResponse, customErr *apperror.CustomError) {
	userRegister := ctx.Value(userRequestKey).(*requests.RegisterRequest)
	onboardRequest := userRegister.OnboardingRequest.(requests.OnboardingUser)
	account, err := o.accountRepository.Save(ctx, &models.Account{
		AccountID:       utils.AccountIDGenerate(userRegister.Type),
		ParentAccountID: ``,
		AccountRole:     userRegister.Type,
	})

	requestId := ctx.Value(enums.RequestID).(string)

	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to register account`, err)
	}

	user, err := o.userRepository.Save(ctx, &models.User{
		AccountID:         account.AccountID,
		UserID:            uuid.NewString(),
		Email:             onboardRequest.Email,
		Password:          utils.HashPassword(onboardRequest.Password),
		DateOfBirth:       fmt.Sprintf("%s-%s-%s", onboardRequest.Year, onboardRequest.Month, onboardRequest.Date),
		FullName:          fmt.Sprintf("%s %s", onboardRequest.FirstName, onboardRequest.LastName),
		UserName:          onboardRequest.UserName,
		Gender:            onboardRequest.Gender,
		Status:            ``,
		ProfilePictureURL: ``,
	})
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to create user`, err)
	}

	return &responses.RegisterResponse{
		RequestID:   requestId,
		UserID:      user.UserID,
		Status:      user.Status,
		VerifiedURL: utils.ReturnVerifyToken(requestId),
		CreatedAt:   time.Now(),
	}, nil
}

func (o *onboardingUsecase) talentRegister(ctx context.Context) (result *responses.RegisterResponse, customErr *apperror.CustomError) {
	return
}

func (o *onboardingUsecase) groupRegister(ctx context.Context) (result *responses.RegisterResponse, customErr *apperror.CustomError) {
	return
}

var _ usecase_interfaces.OnboardingUsecase = &onboardingUsecase{}

func newOnboardingUsecase() *onboardingUsecase {
	o := &onboardingUsecase{}
	onboardMethod := make(map[string]onboardFunc)

	onboardMethod["user"] = o.userRegister
	onboardMethod["talent"] = o.talentRegister
	onboardMethod["group"] = o.groupRegister

	o.onboardMethod = onboardMethod

	return o

}

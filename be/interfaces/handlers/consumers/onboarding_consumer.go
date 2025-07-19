package consumers

import "context"

type OnboardingConsumer interface {
	UserRegistration(ctx context.Context) (err error)
}

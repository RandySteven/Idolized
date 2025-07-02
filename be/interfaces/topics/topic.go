package topics_interfaces

import (
	"context"
)

type Topic interface {
	RunnerConsumer(handler func(string)) error
	WriteMessage(ctx context.Context, value string) (err error)
}

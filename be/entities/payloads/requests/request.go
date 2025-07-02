package requests

type Request struct {
	IdempotencyKey string `json:"idempotency_key"`
}

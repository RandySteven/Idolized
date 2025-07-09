package utils

import (
	"context"
	"fmt"
	"log"
	"time"
)

func Retry(ctx context.Context, maxRetries int, fn func(ctx context.Context) error) error {
	numbOfRetry := 0

	for numbOfRetry < maxRetries {
		err := fn(ctx)
		if err == nil {
			return nil
		}

		numbOfRetry++
		log.Printf("Retry %d/%d failed: %v", numbOfRetry, maxRetries, err)

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(6 * time.Second):
		}
	}

	log.Println("⚠️ Max retries reached")
	return fmt.Errorf("max retries reached")
}

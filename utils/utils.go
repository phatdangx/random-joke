package utils

import (
	"fmt"
	"math/rand"
	"time"
)

type RetryFunc func() error

func RetryWithExponentialBackoff(retryFunc RetryFunc, maxRetries int, baseDelay int) error {
	for i := 0; i < maxRetries; i++ {
		err := retryFunc()
		if err == nil {
			return nil
		}

		if i < maxRetries-1 {
			delay := time.Duration(baseDelay*(1<<i)) * time.Millisecond
			time.Sleep(delay + time.Duration(rand.Intn(100))*time.Millisecond)
		} else {
			return fmt.Errorf("max retries reached: %w", err)
		}
	}
	return nil
}

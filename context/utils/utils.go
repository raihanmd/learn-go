package utils

import (
	"context"
	"time"
)

func CreateCounter(ctx context.Context) chan int {
	dest := make(chan int)

	go func() {
		defer close(dest)
		counter := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				dest <- counter
				counter++
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	return dest
}

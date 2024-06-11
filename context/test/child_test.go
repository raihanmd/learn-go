package test

import (
	"context"
	"context/utils"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestChildContext(t *testing.T) {
	A := context.Background()

	B := context.WithValue(A, "b", "B")
	C := context.WithValue(A, "c", "C")

	D := context.WithValue(B, "d", "D")
	E := context.WithValue(B, "e", "E")

	F := context.WithValue(C, "f", "F")

	for _, val := range []context.Context{A, B, C, D, E, F} {
		t.Log(val)
	}

	assert.Equal(t, "E", E.Value("e").(string))
	assert.Equal(t, "B", D.Value("b").(string))
	assert.Equal(t, "C", F.Value("c").(string))
	assert.Nil(t, F.Value("b"))
}

// ? Goroutine Leak e.g
func TestContextCancel(t *testing.T) {
	t.Log("Total Goroutines:", runtime.NumGoroutine())

	ctx, cancel := context.WithCancel(context.Background())
	dest := utils.CreateCounter(ctx)
	for val := range dest {
		t.Log(val)
		if val == 10 {
			break
		}
	}
	cancel()

	time.Sleep(300 * time.Millisecond)

	t.Log("Total Goroutines:", runtime.NumGoroutine())
	assert.Equal(t, 2, runtime.NumGoroutine())
}

func TestContextTimeout(t *testing.T) {
	t.Log("Total Goroutines:", runtime.NumGoroutine())

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	dest := utils.CreateCounter(ctx)
	for val := range dest {
		t.Log(val)
	}

	t.Log("Total Goroutines:", runtime.NumGoroutine())
	assert.Equal(t, 2, runtime.NumGoroutine())
}

// ? Deadline will persist end the ctx at declared time
func TestContextDeadline(t *testing.T) {
	t.Log("Total Goroutines:", runtime.NumGoroutine())

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	defer cancel()
	dest := utils.CreateCounter(ctx)
	for val := range dest {
		t.Log(val)
	}

	t.Log("Total Goroutines:", runtime.NumGoroutine())
	assert.Equal(t, 2, runtime.NumGoroutine())
}

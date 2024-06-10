# Go Project with Goroutines and Unit Tests

## Welcome

Welcome to the Go project demonstrating the use of goroutines with unit tests! This project shows how to write and test concurrent code in Go using goroutines and the `testing` package.

## Introduction to Goroutines

Goroutines are a lightweight way to achieve concurrency in Go. They allow functions or methods to run concurrently with other functions or methods, making your programs more efficient and responsive.

## Example Goroutine with Unit Test in Go

Here is a simple example of using a goroutine in Go and how to write a unit test for it:

### Example Function with Goroutine

Let's say we have a function that increments a counter in a concurrent manner:

```go
// counter.go
package counter

import (
    "sync"
)

// Counter is a simple counter struct
type Counter struct {
    mu    sync.Mutex
    count int
}

// Increment increments the counter
func (c *Counter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.count++
}

// Value returns the current count
func (c *Counter) Value() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.count
}

// IncrementConcurrently increments the counter n times concurrently
func (c *Counter) IncrementConcurrently(n int) {
    var wg sync.WaitGroup
    for i := 0; i < n; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            c.Increment()
        }()
    }
    wg.Wait()
}
```

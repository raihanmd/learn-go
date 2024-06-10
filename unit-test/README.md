# Go Project with Unit Tests

## Welcome

Welcome to the Go project with unit tests! This project demonstrates how to write and run unit tests in the Go programming language.

## Introduction to Unit Testing

Unit testing is an essential part of software development that allows us to verify that our code works as expected. In Go, we can use the built-in `testing` package to write and run unit tests.

## Example Unit Test in Go

Here is a simple example of a function and its unit test in Go:

### Example Function

Let's say we have a function to add two numbers:

```go
// utils.go
package utils

// Add adds two numbers
func Add(a, b int) int {
    return a + b
}
```

```go
// utils_test.go
package utils

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5

    if result != expected {
        t.Errorf("Add(2, 3) = %d; expected %d", result, expected)
    }
}
```

run with `go test -v (path specify or nothing for root directory)`
